# utils.py

import logging
import os
import uuid
import json
import datetime
import pytz

from user_dashboard.config import Config
from user_dashboard.models import User, Session

def create_user_session(user_id):
    """Creates a new session for the given user."""
    try:
        session = Session(
            user_id=user_id,
            session_id=uuid.uuid4(),
            created_at=datetime.datetime.now(tz=pytz.utc)
        )
        session.save()
        return session
    except Exception as e:
        logging.error(f"Failed to create user session: {e}")
        return None

def get_user(user_id):
    """Retrieves a user by their ID."""
    try:
        return User.get(id=user_id)
    except User.DoesNotExist:
        return None

def is_maintenance_mode():
    """Checks if the application is in maintenance mode."""
    return Config.MAINTENANCE_MODE

def get_current_user():
    """Gets the current user."""
    # NOTE: This is a placeholder implementation. Please implement actual logic to retrieve the current user.
    return User.get(id=1)  # Change this to a real implementation