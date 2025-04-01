import pytest

from solid.src.database import Database
from solid.src.models.user import PaymentInfo
from solid.src.repositories.product_repository import ProductRepository
from solid.src.repositories.user_repository import UserRepository
from solid.src.services.order_service import OrderService
from solid.src.services.payment_service import PaymentService
from solid.src.services.product_service import ProductService
from solid.src.repositories.order_repository import OrderRepository
from solid.src.services.user_service import UserService
from solid.tests.unit_test.fakes.fake_notification_service import (
    FakeNotificationService,
)


@pytest.fixture
def fake_order_repository():
    return OrderRepository(Database())


@pytest.fixture
def fake_notification_service():
    return FakeNotificationService()


@pytest.fixture
def fake_user_service(fake_notification_service):
    return UserService(UserRepository(Database()), fake_notification_service)


@pytest.fixture
def fake_product_service(fake_notification_service):
    return ProductService(ProductRepository(Database()), fake_notification_service)


@pytest.fixture
def fake_payment_service():
    return PaymentService()


def test_create_order(
    fake_order_repository,
    fake_user_service,
    fake_product_service,
    fake_notification_service,
    fake_payment_service,
):
    order_service = OrderService(
        fake_order_repository,
        fake_notification_service,
        fake_product_service,
        fake_user_service,
        fake_payment_service,
    )
    fake_user_service.register_user(
        id=1,
        name="John Doe",
        email="john.doe@example.com",
        password="password",
        address="123 Main St",
        payment_info=PaymentInfo(type="credit_card", card_number="1234567890123456"),
    )
    fake_product_service.add_product(
        id=1,
        name="Product 1",
        price=10,
        category="Category 1",
        description="Description 1",
        stock_quantity=10,
    )
    order_service.create_order(1, [{"product_id": 1, "quantity": 1}])
    assert order_service.order_repository.orders[0].user_id == 1
    assert order_service.order_repository.orders[0].items[0].product.id == 1
    assert order_service.order_repository.orders[0].items[0].quantity == 1
    assert order_service.order_repository.orders[0].items[0].price == 10
    assert order_service.order_repository.orders[0].status == "paid"
    assert order_service.order_repository.orders[0].created_at is not None
