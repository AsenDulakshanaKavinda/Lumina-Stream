
import os
from pathlib import Path
from hydra import initialize_config_dir, compose
from omegaconf import DictConfig
from dotenv import load_dotenv; load_dotenv()

# Resolve the absolute path to the 'configs' directory relative to this file.
CONFIG_PATH = Path(__file__).resolve().parents[1] / "configs"

if not CONFIG_PATH.exists():
    raise FileNotFoundError(f"Config directory not found at {CONFIG_PATH}")

def load_config() -> DictConfig:
    """
    Initializes Hydra and composes the project configuration.

    This function locates the configuration directory, initializes the 
    Hydra context, add the configuration directory to Hydra's search path, 
    and composes the configuration based on the environment variable `ENV`. The default environment is set to "dev".

    Returns:
        DictConfig: A dictionary-like object containing the merged configurations.

    Raises:
        RuntimeError: If Hydra fails to initialize or the config file is missing/invalid.
    """

    config_env = os.getenv("ENV", "dev")
    config_name = f"config.{config_env}"

    config_path = str(CONFIG_PATH)
    try:
        with initialize_config_dir(version_base=None, config_dir=config_path):
            config = compose(config_name=config_name)
        return config
    except Exception as e:
        raise RuntimeError(f"Error while loading configurations: {str(e)}")

config = load_config()