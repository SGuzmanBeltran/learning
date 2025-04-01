from solid.src.database import Database
from solid.src.models.order import Order


class OrderRepository:
    def __init__(self, database: Database):
        self.database = database
        self.orders: list[Order] = []

    def save_order(self, order: Order):
        self.orders.append(order)
        self.database.save_order(order)
