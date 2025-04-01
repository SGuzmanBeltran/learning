from solid.src.models.order import Order, OrderItem
from solid.src.models.product import Product
from solid.src.repositories.order_repository import OrderRepository
from solid.src.services.notification_service import NotificationService
from solid.src.services.payment_service import PaymentService
from solid.src.services.product_service import ProductService
from solid.src.services.user_service import UserService
import time
import datetime


class OrderService:
    def __init__(
        self,
        order_repository: OrderRepository,
        notification_service: NotificationService,
        product_service: ProductService,
        user_service: UserService,
        payment_service: PaymentService,
    ):
        self.order_repository = order_repository
        self.notification_service = notification_service
        self.product_service = product_service
        self.user_service = user_service
        self.payment_service = payment_service

    def create_order(self, user_id: int, items: list[dict]) -> Order | None:
        user = self.user_service.get_user_details(user_id)
        if not user:
            print("User not found.")
            return None

        total_amount, order_items = self.calculate_order_total(items)

        # Apply tax
        tax = total_amount * 0.1  # 10% tax
        total_amount += tax

        # Create order
        order = Order(
            id=f"ORD-{int(time.time())}",
            user_id=user_id,
            items=order_items,
            total_amount=total_amount,
            tax=tax,
            status="pending",
            created_at=datetime.datetime.now(),
        )

        # Process payment
        payment_result = self.payment_service.process_payment(
            user.payment_info, total_amount
        )

        if payment_result["success"]:
            order.status = "paid"

            # Save order to database
            self.order_repository.save_order(order)

            # Send confirmation email
            self.notification_service.send_email(
                user.email,
                "Order Confirmation",
                f"Dear {user.name}, your order ({order.id}) has been processed successfully.",
            )

            print(f"Order {order.id} created and paid successfully.")
            return order
        else:
            print(f"Payment failed: {payment_result['error']}")
            return None

    def calculate_order_total(self, items: list[dict]) -> tuple[float, list[OrderItem]]:
        total_amount = 0
        order_items: list[OrderItem] = []
        # Calculate order total and create order items
        for item in items:
            product = self.product_service.get_product_details(item["product_id"])
            if not product:
                print(f"Product {item['product_id']} not found.")
                continue

            if product.stock_quantity < item["quantity"]:
                print(f"Not enough stock for {product.name}.")
                continue

            item_total = product.price * item["quantity"]
            total_amount += item_total
            self.add_order_item(order_items, product, item, item_total)
        return total_amount, order_items

    def add_order_item(
        self,
        order_items: list[OrderItem],
        product: Product,
        item: dict,
        item_total: float,
    ) -> None:
        order_items.append(
            OrderItem(
                product=product,
                quantity=item["quantity"],
                price=item_total,
            )
        )

        # Update stock
        self.product_service.update_product_stock(
            product.id, product.stock_quantity - item["quantity"]
        )
