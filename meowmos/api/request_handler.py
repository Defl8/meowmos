import requests
from enum import Enum

type EndpointLike = str


class RequestType(Enum):
    GET = 0
    POST = 1
    PUT = 2
    DELETE = 3


class RequestHandler:
    def __init__(self, base_url: str, **endpoints: RequestType) -> None:
        self.base_url: str = base_url
        self.endpoints: dict[EndpointLike, RequestType] = endpoints

    def get(
        self, url: str, params: dict[str, str | int] | None = None
    ) -> requests.Response:
        """Generic GET method request method.

        Params:
            url(str): The url to make the request to.
            params(dict[str, str | int] | None): Optional parameters to send with the request.

        Returns:
            The request response object.
        """

        resp: requests.Response = requests.get(url, params=params)
        if not resp.ok:
            resp.raise_for_status()

        return resp
