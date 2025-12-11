from abc import ABC


class SMSHandler(ABC):
    """Abstract base to describe sms handler capable objects."""

    def __init__(self, acc_sid: str, auth_token: str, send_number: str) -> None:
        self.acc_sid: str = acc_sid
        self.auth_token: str = auth_token
        self.send_number: str = send_number

    def send(self) -> None: ...

    def receive(self) -> str: ...
