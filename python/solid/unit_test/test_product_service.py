from solid.src.models.product import Product
from solid.src.services.product_service import ProductService
import pytest  # type: ignore


class FakeProductRepository:
    def __init__(self):
        self.products = []

    def add_product(self, product: Product) -> bool:
        self.products.append(product)
        return True

    def update_product(self, product_id: str, new_quantity: int) -> Product | None:
        product = next((p for p in self.products if p.id == product_id), None)
        if product:
            product.stock_quantity = new_quantity
        return product

    def get_product_details(self, product_id: str) -> Product | None:
        return next((p for p in self.products if p.id == product_id), None)


class FakeNotificationService:
    def send_email(self, to: str, subject: str, body: str) -> bool:
        return True

    def send_sms(self, to: str, message: str) -> bool:
        return True


@pytest.fixture
def fake_product_repository():
    return FakeProductRepository()


@pytest.fixture
def fake_notification_service():
    return FakeNotificationService()


def test_add_product(fake_product_repository, fake_notification_service):
    product_service = ProductService(fake_product_repository, fake_notification_service)
    product_service.add_product(
        1, "Test Product", 10, "Test Category", "Test Description", 100
    )
    assert len(product_service.product_repository.products) == 1


def test_update_product_stock(fake_product_repository, fake_notification_service):
    product_service = ProductService(fake_product_repository, fake_notification_service)
    product_service.add_product(
        1, "Test Product", 10, "Test Category", "Test Description", 100
    )
    product_service.update_product_stock(1, 50)
    updated_product = product_service.update_product_stock(2, 100)
    assert product_service.product_repository.products[0].stock_quantity == 50
    assert updated_product is None


def test_get_product_details(fake_product_repository):
    product_service = ProductService(fake_product_repository, fake_notification_service)
    product_service.add_product(
        1, "Test Product", 10, "Test Category", "Test Description", 100
    )
    product = product_service.get_product_details(1)
    assert product is not None
    assert product.id == 1
    assert product.name == "Test Product"
    assert product.price == 10
    assert product.category == "Test Category"
    assert product.description == "Test Description"
    assert product.stock_quantity == 100
