import dotenv
from environment.dotenv import get_env_vars


def main() -> None:
    loaded: bool = dotenv.load_dotenv()
    if not loaded:
        raise ValueError(".env file could not be loaded.")

    print(".env file loaded successfully.")





if __name__ == "__main__":
    main()
