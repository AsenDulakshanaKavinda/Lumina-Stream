from utils import config, get_logger

log = get_logger(__file__)

def main():
    log.info("Hello from python-embedder!")
    print(config)


if __name__ == "__main__":
    main()
