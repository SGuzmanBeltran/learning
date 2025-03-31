from solid.src.models.product import Product
from solid.src.refactorable_code import Database
import pytest

from solid.src.repositories.product_repository import ProductRepository  # type: ignore


@pytest.fixture
def fake_database():
    return Database()


def test_add_product(fake_database):
    product_repository = ProductRepository(fake_database)
    product_repository.add_product(
        Product(1, "Test Product", 10, "Test Category", "Test Description", 100)
    )
    assert len(product_repository.products) == 1
    assert product_repository.products[0].id == 1
    assert product_repository.products[0].name == "Test Product"
    assert product_repository.products[0].price == 10
    assert product_repository.products[0].category == "Test Category"
    assert product_repository.products[0].description == "Test Description"


def test_update_product(fake_database):
    product_repository = ProductRepository(fake_database)
    product_repository.add_product(
        Product(1, "Test Product", 10, "Test Category", "Test Description", 100)
    )
    updated_product = product_repository.update_product(1, 50)
    assert updated_product is not None
    assert updated_product.stock_quantity == 50
    assert product_repository.products[0].stock_quantity == 50


def test_get_product_details(fake_database):
    product_repository = ProductRepository(fake_database)
    product_repository.add_product(
        Product(1, "Test Product", 10, "Test Category", "Test Description", 100)
    )
    product = product_repository.get_product_details(1)
    assert product is not None
    assert product.id == 1
    assert product.name == "Test Product"
    assert product.price == 10
    assert product.category == "Test Category"
    assert product.description == "Test Description"
    assert product.stock_quantity == 100
