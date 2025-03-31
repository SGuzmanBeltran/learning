from solid.src.models.product import Product
from solid.src.models.user import User


class Database:
    def save_product(self, product: Product) -> bool:
        print(f"Saving product to database: {product.name}")
        # In a real implementation, this would save to a database
        return True

    def update_product(self, product: Product) -> bool:
        print(f"Updating product in database: {product.name}")
        # In a real implementation, this would update a database record
        return True

    def save_user(self, user: User) -> bool:
        print(f"Saving user to database: {user.name}")
        # In a real implementation, this would save to a database
        return True

    def save_order(self, order):
        print(f"Saving order to database: {order['id']}")
        # In a real implementation, this would save to a database
        return True

    def save_report(self, report):
        print(
            f"Saving report to database for period: {report['start_date'].strftime('%Y-%m-%d')} to {report['end_date'].strftime('%Y-%m-%d')}"
        )
        # In a real implementation, this would save to a database
        return True
