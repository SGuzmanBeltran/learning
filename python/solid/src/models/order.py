import datetime
from solid.src.models.product import Product


class OrderItem:
    def __init__(self, product: Product, quantity: int, price: float):
        self.product = product
        self.quantity = quantity
        self.price = price


class Order:
    def __init__(
        self,
        id: str,
        user_id: int,
        items: list[OrderItem],
        total_amount: float,
        tax: float,
        status: str,
        created_at: datetime.datetime,
    ):
        self.id = id
        self.user_id = user_id
        self.items = items
        self.total_amount = total_amount
        self.tax = tax
        self.status = status
        self.created_at = created_at
