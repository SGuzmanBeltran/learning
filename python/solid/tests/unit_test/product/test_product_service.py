from solid.src.refactorable_code import Database
from solid.src.repositories.product_repository import ProductRepository
from solid.src.services.product_service import ProductService
import pytest

from solid.tests.unit_test.fakes.fake_notification_service import (
    FakeNotificationService,
)  # type: ignore


@pytest.fixture
def fake_product_repository():
    return ProductRepository(Database())


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


def test_update_product_notification(
    fake_product_repository, fake_notification_service
):
    product_service = ProductService(fake_product_repository, fake_notification_service)
    product_service.add_product(
        1, "Test Product", 10, "Test Category", "Test Description", 100
    )
    product_service.update_product_stock(1, 5)
    assert fake_notification_service.send_email_count == 1
    assert fake_notification_service.send_sms_count == 0


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
