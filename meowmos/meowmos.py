import dotenv
from environment.dotenv import get_env_vars

from sms.messenger import Messenger


def main() -> None:
    loaded: bool = dotenv.load_dotenv()
    if not loaded:
        raise ValueError(".env file could not be loaded.")

    print(".env file loaded successfully.")

    sid, token, send_num = get_env_vars("TWILIO_ACC_SID", "TWILIO_AUTH_TOKEN", "SEND_NUMBER")

    messenger: Messenger = Messenger(sid, token, send_num)
    messenger.send("+17802434269", "this is a test")




if __name__ == "__main__":
    main()
