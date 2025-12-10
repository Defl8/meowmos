from typing import Protocol


class SMSHandler(Protocol):
    def send(self) -> None: ...

    def receive(self) -> str: ...

