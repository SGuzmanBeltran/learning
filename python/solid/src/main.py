from solid.src.database import Database
from solid.src.refactorable_code import NotificationService, OnlineStore
from solid.src.repositories.order_repository import OrderRepository
from solid.src.repositories.product_repository import ProductRepository
from solid.src.repositories.user_repository import UserRepository
from solid.src.services.order_service import OrderService
from solid.src.services.payment_service import PaymentService
from solid.src.services.product_service import ProductService
from solid.src.services.user_service import UserService


if __name__ == "__main__":
    import datetime

    store = OnlineStore()

    database = Database()
    notification_service = NotificationService()

    product_repository = ProductRepository(database)
    product_service = ProductService(product_repository, notification_service)
    # Add products
    product_service.add_product(
        1, "Laptop", 999.99, "Electronics", "Powerful laptop for developers", 20
    )
    product_service.add_product(
        2, "Headphones", 149.99, "Electronics", "Noise-cancelling headphones", 30
    )

    user_repository = UserRepository(database)
    user_service = UserService(user_repository, notification_service)
    # Register a user
    user_service.register_user(
        1,
        "John Doe",
        "john@example.com",
        "password123",
        {"street": "123 Main St", "city": "Anytown", "zip_code": "12345"},
        {
            "type": "credit_card",
            "card_number": "1234567890123456",
            "expiry_date": "12/24",
            "cvv": "123",
        },
    )

    order_repository = OrderRepository(database)
    payment_service = PaymentService()
    # Create an order
    order_service = OrderService(
        order_repository,
        notification_service,
        product_service,
        user_service,
        payment_service,
    )
    order_service.create_order(
        1, [{"product_id": 1, "quantity": 1}, {"product_id": 2, "quantity": 2}]
    )

    # Generate a sales report
    # start_date = datetime.datetime(2023, 1, 1)
    # end_date = datetime.datetime(2023, 12, 31)
    # store.generate_sales_report(start_date, end_date)
