from solid.src.models.user import User
from solid.src.repositories.user_repository import UserRepository
from solid.src.services.notification_service import NotificationService
import bcrypt  # type: ignore


class UserService:
    def __init__(
        self, user_repository: UserRepository, notification_service: NotificationService
    ):
        self.user_repository = user_repository
        self.notification_service = notification_service

    def register_user(
        self,
        id: int,
        name: str,
        email: str,
        password: str,
        address: dict,
        payment_info: dict,
    ) -> None:
        # Hash the password (very insecure way)
        hashed_password = self.hash_password(password)

        user = User(
            id=id,
            name=name,
            email=email,
            password=hashed_password,
            address=address,
            payment_info=payment_info,
        )

        user_added = self.user_repository.add_user(user)

        # Send welcome email
        self.notification_service.send_email(
            user_added.email,
            "Welcome to our store!",
            f"Dear {user_added.name}, thank you for registering!",
        )

        print(f"User {user_added.name} registered successfully.")

    def get_user_details(self, user_id: int) -> User | None:
        return self.user_repository.get_user_details(user_id)

    def hash_password(self, password: str) -> str:
        # Generate a salt
        salt = bcrypt.gensalt()
        # Hash the password
        hashed_password = bcrypt.hashpw(password.encode("utf-8"), salt)
        return hashed_password.decode("utf-8")
