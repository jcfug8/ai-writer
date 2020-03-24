import logging

import os
import numpy as np
import tensorflow as tf

import model
import sample
import encoder
import json


class Predictor():

    def __init__(
        self,
        model_name='124M',
        seed=None,
        simple_nsamples=3,
        batch_size=1,
        simple_length=10,
        large_length=100,
        temperature=1,
        top_k=0,
        top_p=0.0,
    ):
        """
        Interactively run the model
        :model_name=117M : String, which model to use
        :seed=None : Integer seed for random number generators, fix seed to reproduce
        results
        :simple_nsamples=1 : Number of samples to return total
        :batch_size=1 : Number of batches (only affects speed/memory).  Must divide simple_nsamples.
        :simple_length=None : Number of tokens in generated text, if None (default), is
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
        self.simple_nsamples = simple_nsamples
        self.batch_size = batch_size
        self.simple_length = simple_length
        self.large_length = large_length
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
        """Ensure that simple_nsamples is divisable by batch size"""
        logging.info("ensuring simple_nsamples is divisable by batch size")
        if self.batch_size is None:
            self.batch_size = 1
        assert self.simple_nsamples % self.batch_size == 0

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

        logging.info("setting up generation simple_length")
        if self.simple_length is None:
            logging.info("simple_length was None")
            self.simple_length = self.hparams.n_ctx // 2
        elif self.simple_length > self.hparams.n_ctx:
            raise ValueError(
                "Can't get samples longer than window size: %s" % self.hparams.n_ctx)

        logging.info("setting up generation large_length")
        if self.large_length is None:
            logging.info("large_length was None")
            self.large_length = self.hparams.n_ctx // 2
        elif self.large_length > self.hparams.n_ctx:
            raise ValueError(
                "Can't get samples longer than window size: %s" % self.hparams.n_ctx)

    def init_context(self):
        """Create the context"""
        logging.info("getting context")
        self.context = tf.placeholder(tf.int32, [self.batch_size, None])

        logging.info("setting up simple output")
        self.simple_output = sample.sample_sequence(
            hparams=self.hparams, length=self.simple_length,
            context=self.context,
            batch_size=self.batch_size,
            temperature=self.temperature, top_k=self.top_k, top_p=self.top_p
        )

        logging.info("setting up large output")
        self.large_output = sample.sample_sequence(
            hparams=self.hparams, length=self.large_length,
            context=self.context,
            batch_size=1,
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

    """Tensorflow Functions"""

    def simple(self, text=""):
        with self.sess.graph.as_default():
            context_tokens = self.enc.encode(text)
            generated = []
            for j in range(self.simple_nsamples // self.batch_size):
                logging.info("outer loop " + str(j))
                out = self.sess.run(self.simple_output, feed_dict={
                    self.context: [context_tokens for _ in range(self.batch_size)]
                })[:, len(context_tokens):]
                for i in range(self.batch_size):
                    logging.info("inner loop " + str(i))
                    generated.append(self.enc.decode(out[i]))
        return generated

    def large(self, text=""):
        with self.sess.graph.as_default():
            context_tokens = self.enc.encode(text)
            logging.info("generating large")
            out = self.sess.run(self.large_output, feed_dict={
                self.context: [context_tokens]
            })[:, len(context_tokens):]
        return self.enc.decode(out[0])
