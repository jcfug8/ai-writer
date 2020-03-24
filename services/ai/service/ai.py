from concurrent import futures
import logging

import entities_pb2_grpc as grpcPb
import entities_pb2 as pb
import grpc
import json
import predictor


class AIServicer(grpcPb.AIServicer):
    """Provides methods that implement functionality of ai server."""
    predictors = {}

    def __init__(self):
        self.predictors["general"] = predictor.Predictor()

    def GetSimpleGeneration(self, request, context):
        generated_text = None
        try:
            # pass
            generated_text = self.predictors["general"].simple(
                text=request.seed_text
            )
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

    def GetLargeGeneration(self, request, context):
        generated_text = None
        try:
            # pass
            generated_text = self.predictors["general"].large(
                text=request.seed_text
            )
        except Exception as e:
            logging.info("Error happened - " + str(e))
            context.abort(
                grpc.StatusCode.INVALID_ARGUMENT,
                json.dumps({
                    "status_text": "InternalServerError",
                    "satus": 400,
                    "message": str(e)
                })
            )
            return
        logging.info("generation successful " + str(generated_text))
        return pb.GetLargeGenerationReply(message=generated_text)


def serve():
    logging.info("Createing AIServicer")
    servicer = AIServicer()

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
