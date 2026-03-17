# utils.py
import logging
import os
from typing import Optional
from datetime import datetime

def configure_logging(log_level: str, log_file: str) -> None:
    """Configure the logging module to output logs to a file."""
    logging.basicConfig(
        level=getattr(logging, log_level.upper()),
        format="%(asctime)s - %(levelname)s - %(message)s",
        handlers=[
            logging.FileHandler(log_file),
            logging.StreamHandler()
        ]
    )

def get_current_datetime() -> str:
    """Return the current date and time in the 'YYYY-MM-DD HH:MM:SS' format."""
    return datetime.now().strftime("%Y-%m-%d %H:%M:%S")

def get_user_home_directory() -> str:
    """Return the path to the user's home directory."""
    return os.path.expanduser("~")

def is_valid_email(email: str) -> bool:
    """Check if the provided email address is valid."""
    import re
    email_regex = r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
    return bool(re.match(email_regex, email))

def get_random_password(length: int = 12) -> str:
    """Generate a random password of the specified length."""
    import secrets
    return ''.join(secrets.choice('abcdefghijklmnopqrstuvwxyz0123456789') for _ in range(length))

def get_uuid() -> str:
    """Generate a random UUID."""
    import uuid
    return str(uuid.uuid4())

def get_environment_variable(var_name: str) -> Optional[str]:
    """Return the value of the environment variable with the specified name."""
    return os.environ.get(var_name)