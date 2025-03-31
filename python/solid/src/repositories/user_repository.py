from solid.src.models.user import User
from solid.src.refactorable_code import Database


class UserRepository:
    def __init__(self, database: Database):
        self.users: list[User] = []
        self.db = database

    def add_user(self, user: User) -> User:
        self.users.append(user)
        self.db.save_user(user)
        return user

    def get_user_details(self, user_id: int) -> User | None:
        return next((u for u in self.users if u.id == user_id), None)
