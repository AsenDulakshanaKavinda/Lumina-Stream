from utils import config, get_logger
from app.service import server

log = get_logger(__file__)

def main():
    server()


if __name__ == "__main__":
    main()
