from solid.models.product import Product
from solid.refactorable_code import NotificationService
from solid.repositories.product_repository import ProductRepository  # type: ignore


class ProductService:
    def __init__(
        self,
        product_repository: ProductRepository,
        notification_service: NotificationService,
    ):
        self.product_repository = product_repository
        self.notification_service = notification_service

    def add_product(
        self,
        id: int,
        name: str,
        price: float,
        category: str,
        description: str,
        stock_quantity: int,
    ) -> None:
        product = Product(
            id=id,
            name=name,
            price=price,
            category=category,
            description=description,
            stock_quantity=stock_quantity,
        )
        result = self.product_repository.add_product(product)
        if result:
            print(f"Product {name} added to inventory.")
        else:
            print(f"Failed to add product {name} to inventory.")

    def update_product_stock(self, product_id: str, new_quantity: int) -> None:
        updated_product = self.product_repository.update_product(
            product_id, new_quantity
        )

        if updated_product and updated_product.stock_quantity < 10:
            self.notification_service.send_email(
                "admin@store.com",
                f"Low stock alert for {updated_product.name}",
                f"The stock for {updated_product.name} is low ({updated_product.stock_quantity} remaining).",
            )

        if not updated_product:
            print(f"Failed to update product {product_id} stock.")

    def get_product_details(self, product_id: str) -> Product | None:
        return self.product_repository.get_product_details(product_id)
