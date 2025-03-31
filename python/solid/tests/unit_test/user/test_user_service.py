import bcrypt
import pytest

from solid.src.database import Database
from solid.src.repositories.user_repository import UserRepository
from solid.src.services.user_service import UserService
from solid.tests.unit_test.fakes.fake_notification_service import (
    FakeNotificationService,
)


@pytest.fixture
def fake_user_repository():
    return UserRepository(Database())


@pytest.fixture
def fake_notification_service():
    return FakeNotificationService()


def test_register_user(fake_user_repository, fake_notification_service):
    user_service = UserService(fake_user_repository, fake_notification_service)
    user_service.register_user(
        1, "John Doe", "john@example.com", "password123", "123 Main St", "1234567890"
    )
    assert len(user_service.user_repository.users) == 1


def test_get_user_details(fake_user_repository, fake_notification_service):
    user_service = UserService(fake_user_repository, fake_notification_service)
    user_service.register_user(
        1, "John Doe", "john@example.com", "password123", "123 Main St", "1234567890"
    )
    user = user_service.get_user_details(1)
    assert user is not None
    assert user.name == "John Doe"
    assert user.email == "john@example.com"
    assert user.address == "123 Main St"
    assert user.payment_info == "1234567890"


def test_hash_password(fake_user_repository, fake_notification_service):
    user_service = UserService(fake_user_repository, fake_notification_service)
    hashed_password = user_service.hash_password("password123")
    assert hashed_password is not None
    assert len(hashed_password) > 0
    assert bcrypt.checkpw(
        "password123".encode("utf-8"), hashed_password.encode("utf-8")
    )
