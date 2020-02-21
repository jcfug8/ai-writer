from concurrent import futures
import logging

import entities_pb2_grpc as grpcPb
import entities_pb2 as pb
import grpc
import json
import os
import numpy as np
import tensorflow as tf

import model
import sample
import encoder


class AIServicer(grpcPb.AIServicer):
    """Provides methods that implement functionality of ai server."""

    def __init__(
        self,
        model_name='124M',
        seed=None,
        nsamples=1,
        batch_size=1,
        length=None,
        temperature=1,
        top_k=0,
        top_p=0.0,
    ):
        """
        Interactively run the model
        :model_name=117M : String, which model to use
        :seed=None : Integer seed for random number generators, fix seed to reproduce
        results
        :nsamples=1 : Number of samples to return total
        :batch_size=1 : Number of batches (only affects speed/memory).  Must divide nsamples.
        :length=None : Number of tokens in generated text, if None (default), is
        determined by model hyperparameters
        :temperature=1 : Float value controlling randomness in boltzmann
        distribution. Lower temperature results in less random completions. As the
        temperature approaches zero, the model will become deterministic and
        repetitive. Higher temperature results in more random completions.
        :top_k=0 : Integer value controlling diversity. 1 means only 1 word is
        considered for each step (token), resulting in deterministic completions,
        while 40 means 40 words are considered at each step. 0 (default) is a
        special setting meaning no restrictions. 40 generally is a good value.
        :top_p=0.0 : Float value controlling diversity. Implements nucleus sampling,
        overriding top_k if set to a value > 0. A good setting is 0.9.
        """

        self.model_name = model_name
        self.seed = seed
        self.nsamples = nsamples
        self.batch_size = batch_size
        self.length = length
        self.temperature = temperature
        self.top_k = top_k
        self.top_p = top_p
        self.sess = tf.Session(graph=tf.Graph())

        # initialization functions
        with self.sess.graph.as_default():
            self.check_samples_and_size()
            self.get_encoder()
            self.override_hparams()
            self.init_context()
            self.set_seed()
            self.restore_checkpoint()

    def check_samples_and_size(self):
        """Ensure that nsamples is divisable by batch size"""
        logging.info("ensuring nsamples is divisable by batch size")
        if self.batch_size is None:
            self.batch_size = 1
        assert self.nsamples % self.batch_size == 0

    def get_encoder(self):
        logging.info("getting model's encoder")
        self.enc = encoder.get_encoder(self.model_name)

    def override_hparams(self):
        """Load model h params and ensure generation lenth isn't too large"""
        logging.info("setting up hparams")
        self.hparams = model.default_hparams()
        with open(os.path.join('models', self.model_name, 'hparams.json')) as f:
            logging.info("loading hparams.json")
            self.hparams.override_from_dict(json.load(f))

        logging.info("setting up generation length")
        if self.length is None:
            logging.info("length was None")
            self.length = self.hparams.n_ctx // 2
        elif self.length > self.hparams.n_ctx:
            raise ValueError(
                "Can't get samples longer than window size: %s" % self.hparams.n_ctx)

    def init_context(self):
        """Create the context"""
        logging.info("getting context")
        self.context = tf.placeholder(tf.int32, [self.batch_size, None])

        logging.info("setting up output")
        self.output = sample.sample_sequence(
            hparams=self.hparams, length=self.length,
            context=self.context,
            batch_size=self.batch_size,
            temperature=self.temperature, top_k=self.top_k, top_p=self.top_p
        )

    def set_seed(self):
        """Set up seed"""
        logging.info("setting up seed")
        np.random.seed(self.seed)
        tf.set_random_seed(self.seed)

    def restore_checkpoint(self):
        """Restore Checkpoint"""
        logging.info("creating saver")
        saver = tf.train.Saver()

        logging.info("loading checkpoint")
        ckpt = tf.train.latest_checkpoint(
            os.path.join('models', self.model_name))

        logging.info("restoring session")
        saver.restore(self.sess, ckpt)

    def __del__(self):
        self.sess.close()

    def GetSimpleGeneration(self, request, context):
        generated_text = None
        try:
            # pass
            generated_text = self.interact_model(text=request.seed_text)
        except Exception as e:
            logging.info("Error happened - " + str(e))
            context.abort(
                grpc.StatusCode.INVALID_ARGUMENT,
                json.dumps({
                    "status_text": "InternalServerError",
                    "satus": 400,
                    "messages": [str(e)]
                })
            )
            return
        logging.info("generation successful " + str(generated_text))
        return pb.GetSimpleGenerationReply(messages=generated_text)

    """Tensorflow Functions"""

    def interact_model(self, text=""):
        with self.sess.graph.as_default():
            context_tokens = self.enc.encode(text)
            generated = []
            for j in range(self.nsamples // self.batch_size):
                logging.info("outer loop " + str(j))
                out = self.sess.run(self.output, feed_dict={
                    self.context: [context_tokens for _ in range(self.batch_size)]
                })[:, len(context_tokens):]
                for i in range(self.batch_size):
                    logging.info("inner loop " + str(i))
                    generated.append(self.enc.decode(out[i]))
        return generated


def serve():
    logging.info("Createing AIServicer")
    servicer = AIServicer(length=10, nsamples=3)

    logging.info("Starting AI Service")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    grpcPb.add_AIServicer_to_server(servicer, server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.info("Loading basic logging config")
    logging.basicConfig()
    logging.info("Setting logging level")
    root = logging.getLogger()
    root.setLevel(0)
    logging.info("About to start AI Service")
    serve()
