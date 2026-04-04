import grpc
from concurrent import futures
import pb.embedding_pb2 as embedding_pb2
import pb.embedding_pb2_grpc as embedding_pb2_grpc

class EmbedderServicer(embedding_pb2_grpc.EmbedderServicer):
    def Embed(self, request, context):
        return embedding_pb2.VectorResponse(
            vector=[0.1, 0.2, 0.3]
        )
    
def server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    embedding_pb2_grpc.add_EmbedderServicer_to_server(EmbedderServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    server()

