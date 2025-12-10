import os
from dotenv import load_dotenv


def get_env_var(var_name: str) -> str:
    """Fetch an evironment variable value.

    Params:
        var_name(str): The environment variable name.

    Returns:
        The environment variable value as a string.
    """
    VAR: str | None = os.getenv(var_name)
    if VAR is None:
        raise TypeError("SID could not be found.")
    return VAR


def get_env_vars(*args: str) -> tuple[str, ...]:
    """Get the requested the event environment variables.

    Params:
        args(str): Argument collection containing the desired environment variables.

    Returns:
        A tuple of the values of the provided environment variables as strings.
    """
    return tuple(get_env_var(arg) for arg in args)
