import grpc
from concurrent import futures

from pb import embedding_pb2_grpc, embedding_pb2
from utils import config, get_logger

from .model import embedding_model

log = get_logger(__file__)
address = config.environments.grpc_config.server_address


class EmbedderServicer(embedding_pb2_grpc.EmbedderServicer):
    def Embed(self, request, context):
        try:
            log.info("Received a request")
            return embedding_pb2.VectorResponse(
                vector=embedding_model.encode(request.text)
            )
        except Exception as e:
            log.error("Error while responding to the request.")
            raise RuntimeError(f"Error while responding to the request: {str(e)}")
    

def server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    embedding_pb2_grpc.add_EmbedderServicer_to_server(EmbedderServicer(), server)
    server.add_insecure_port(address)
    server.start()
    log.info("Server is running")
    server.wait_for_termination()


