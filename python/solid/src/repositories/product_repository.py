from solid.src.database import Database
from solid.src.models.product import Product


class ProductRepository:
    def __init__(self, database: Database):
        self.products: list[Product] = []
        self.db = database

    def add_product(self, product: Product) -> bool:
        self.products.append(product)
        return self.db.save_product(product)

    def update_product(self, product_id: int, new_quantity: int) -> Product | None:
        product = next((p for p in self.products if p.id == product_id), None)
        if product:
            product.stock_quantity = new_quantity
            self.db.update_product(product)
            print(f"Product {product.name} stock updated to {new_quantity}.")
        return product

    def get_product_details(self, product_id: int) -> Product | None:
        return next((p for p in self.products if p.id == product_id), None)
