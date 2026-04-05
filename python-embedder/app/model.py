from pathlib import Path

from sentence_transformers import SentenceTransformer

from utils import get_logger, config


log = get_logger(__file__)


class Embeddings:
    """
    A class to handle the embedding model.

    Attributes:
        model_config: The configuration for the embedding model.

    Methods:        
        - download_embedding_model: Download the embedding model and save it to the specified path.
        - loading_embedding_model: Load the embedding model from the specified path.
    
    """

    def __init__(self):
        self.model_config = config.environments.embedding_model


    def download_embedding_model(self) -> None:
        """
        Download the embedding model and save it to the specified path.
        """
        try:
            log.info("Downloading embedding model")

            model = SentenceTransformer(
                model_name_or_path = self.model_config.model_name,
            )
            log.info(f"Saving model")
            model.save(self.model_config.model_path)

        except Exception as e:
            log.error("Error while downloading embedding model")
            raise RuntimeError(f"Error while downloading embedding model: {str(e)}")
        
        
    def loading_embedding_model(self) -> SentenceTransformer:
        """
        Load the embedding model from the specified path.

        args:            
            model_path: The path to the embedding model.
        returns:         
            The loaded embedding model.
        exceptions:
            RuntimeError: If there is an error while loading the embedding model.
        """

        try:
            log.info("Loading embedding model")
            model = SentenceTransformer(
                model_name_or_path = self.model_config.model_path,
                local_files_only = True
            )
            return model

        except Exception as e:
            log.error("Error while loading embedding model")
            raise RuntimeError(f"Error while loading embedding model: {str(e)}")


# - Initialize the embedding model
embedding_model = Embeddings().loading_embedding_model()
