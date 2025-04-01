import datetime
import pytest

from solid.src.database import Database
from solid.src.models.order import Order, OrderItem
from solid.src.models.product import Product
from solid.src.repositories.order_repository import OrderRepository


@pytest.fixture
def fake_database():
    return Database()


def test_save_order(fake_database):
    order_repository = OrderRepository(fake_database)

    order_repository.save_order(
        Order(
            id="1",
            user_id=1,
            items=[
                OrderItem(
                    Product(
                        1, "Test Product", 10, "Test Category", "Test Description", 100
                    ),
                    1,
                    10,
                )
            ],
            total_amount=10,
            tax=1,
            status="pending",
            created_at=datetime.datetime.now(),
        )
    )

    assert len(order_repository.orders) == 1
    assert order_repository.orders[0].id == "1"
    assert order_repository.orders[0].user_id == 1
    assert order_repository.orders[0].items[0].product.name == "Test Product"
    assert order_repository.orders[0].total_amount == 10
    assert order_repository.orders[0].status == "pending"
    assert order_repository.orders[0].created_at is not None
