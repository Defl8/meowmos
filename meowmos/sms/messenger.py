from typing import override
from .sms_handler import SMSHandler


class Messenger(SMSHandler):
    def __init__(self, acc_sid: str, auth_token: str, send_number: str) -> None:
        super().__init__(acc_sid, auth_token, send_number)

    @override
    def send(self, receive_number: str, msg_body: str) -> None:
        print(f"Sending msg to '{receive_number}'")
        super().send(receive_number, msg_body)
