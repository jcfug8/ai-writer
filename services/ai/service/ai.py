from concurrent import futures
import multiprocessing
import logging

import entities_pb2_grpc as grpcPb
import entities_pb2 as pb
import grpc
import json
import predictor

SMALL_TYPE = "SMALL_TYPE"
LARGE_TYPE = "LARGE_TYPE"


def predictor_runner(model_name, pipe):
    p = predictor.Predictor(model_name=model_name)
    pipe.send(True)
    while 1:
        req_pipe, seed, t = pipe.recv()
        logging.info("received generation request from " +
                     model_name + " predictor")
        try:
            res = None
            if t == SMALL_TYPE:
                res = p.simple(text=seed)
            elif t == LARGE_TYPE:
                res = p.large(text=seed)
            else:
                raise Exception("unknown prediction type")
            req_pipe.send((res, 1))
        except Exception as e:
            req_pipe.send((e, -1))


class AIServicer(grpcPb.AIServicer):
    """Provides methods that implement functionality of ai server."""
    predictors_pipes = []
    predictors_proccesses = []

    def __init__(self, m_names):
        for m_name in m_names:
            p_pipe, c_pipe = multiprocessing.Pipe()
            p = multiprocessing.Process(
                target=predictor_runner,
                args=(m_name, c_pipe)
            )
            p.start()
            p_pipe.recv()
            logging.info("Predictor " + m_name + " Started")
            self.predictors_pipes.append(p_pipe)
            self.predictors_proccesses.append(p)

    def GetSimpleGeneration(self, request, context):
        msg = None
        try:
            genreKey = 0
            logging.info("genre " + request.genre)
            if request.genre != 0:
                genreKey = request.genre - 1

            p_pipe, c_pipe = multiprocessing.Pipe()
            self.predictors_pipes[genreKey].send(
                (c_pipe, request.seed_text, SMALL_TYPE)
            )
            msg, status_code = p_pipe.recv()
            if status_code == -1:
                raise msg
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
        logging.info("generation successful " + str(msg))
        return pb.GetSimpleGenerationReply(messages=msg)

    def GetLargeGeneration(self, request, context):
        msg = None
        try:
            genreKey = 0
            logging.info("genre " + request.genre)
            if request.genre != 0:
                genreKey = request.genre - 1

            p_pipe, c_pipe = multiprocessing.Pipe()
            self.predictors_pipes[genreKey].send(
                (c_pipe, request.seed_text, LARGE_TYPE)
            )
            msg, status_code = p_pipe.recv()
            if status_code == -1:
                raise msg
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
        logging.info("generation successful " + str(msg))
        return pb.GetLargeGenerationReply(message=msg)


def serve():
    logging.info("Createing AIServicer")
    servicer = AIServicer([
        "agatha_christie",
        "charles_dickens",
    ])

    logging.info("Starting AI Service")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    grpcPb.add_AIServicer_to_server(servicer, server)
    server.add_insecure_port('[::]:50051')
    server.start()
    logging.info("Started AI Service")
    server.wait_for_termination()


if __name__ == '__main__':
    logging.info("Loading basic logging config")
    logging.basicConfig()
    logging.info("Setting logging level")
    root = logging.getLogger()
    root.setLevel(0)
    logging.info("About to start AI Service")
    serve()
