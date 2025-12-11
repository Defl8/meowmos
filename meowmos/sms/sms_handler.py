from twilio.rest import Client

from abc import ABC


class SMSHandler(ABC):
    """Abstract base to describe sms handler capable objects."""

    def __init__(self, acc_sid: str, auth_token: str, send_number: str) -> None:
        self.acc_sid: str = acc_sid
        self.auth_token: str = auth_token
        self.send_number: str = send_number
        self.twilio_client: Client = Client(acc_sid, auth_token)

    def send(self, receive_number: str, msg_body: str) -> None:
        _ = self.twilio_client.messages.create(
            to=receive_number, from_=self.send_number, body=msg_body
        )

    def receive(self) -> str: ...
