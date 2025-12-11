from .handler import SMSHandler


class Messenger(SMSHandler):
    def __init__(self, acc_sid: str, auth_token: str, send_number: str) -> None:
        super().__init__(acc_sid, auth_token, send_number)
