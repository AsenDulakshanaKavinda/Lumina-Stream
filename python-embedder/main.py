from utils import config, get_logger
from app.service import server
from app import Embeddings

log = get_logger(__file__)

def download_embedding_model():
    emb = Embeddings()
    emb.download_embedding_model()


def main():
    server()
    # download_embedding_model()


if __name__ == "__main__":
    main()
