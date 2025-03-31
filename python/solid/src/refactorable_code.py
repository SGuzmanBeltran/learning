# A badly designed e-commerce system with multiple SOLID violations


from solid.src.database import Database
from solid.src.services.notification_service import NotificationService


class OnlineStore:
    def __init__(self):
        self.orders: list = []
        self.payment_processor = PaymentProcessor()
        self.notification_service = NotificationService()
        self.db = Database()

    # Order Processing
    def create_order(self, user_id, items):
        user = self.get_user_details(user_id)
        if not user:
            print("User not found.")
            return

        total_amount = 0
        order_items = []

        # Calculate order total and create order items
        for item in items:
            product = self.get_product_details(item["product_id"])
            if not product:
                print(f"Product {item['product_id']} not found.")
                continue

            if product["stock_quantity"] < item["quantity"]:
                print(f"Not enough stock for {product['name']}.")
                continue

            item_total = product["price"] * item["quantity"]
            total_amount += item_total

            order_items.append(
                {
                    "product": product,
                    "quantity": item["quantity"],
                    "item_total": item_total,
                }
            )

            # Update stock
            self.update_product_stock(
                product["id"], product["stock_quantity"] - item["quantity"]
            )

        # Apply tax
        tax = total_amount * 0.1  # 10% tax
        total_amount += tax

        # Create order
        import time
        import datetime

        order = {
            "id": f"ORD-{int(time.time())}",
            "user_id": user_id,
            "items": order_items,
            "total_amount": total_amount,
            "tax": tax,
            "status": "pending",
            "created_at": datetime.datetime.now(),
        }

        # Process payment
        payment_result = self.payment_processor.process_payment(
            user["payment_info"], total_amount
        )

        if payment_result["success"]:
            order["status"] = "paid"

            # Save order to database
            self.orders.append(order)
            self.db.save_order(order)

            # Send confirmation email
            self.notification_service.send_email(
                user["email"],
                "Order Confirmation",
                f"Dear {user['name']}, your order ({order['id']}) has been processed successfully.",
            )

            print(f"Order {order['id']} created and paid successfully.")
            return order
        else:
            print(f"Payment failed: {payment_result['error']}")
            return None

    # Reporting
    def generate_sales_report(self, start_date, end_date):
        # Filter orders by date range
        filtered_orders = [
            order
            for order in self.orders
            if start_date <= order["created_at"] <= end_date
        ]

        # Calculate total sales
        total_sales = sum(order["total_amount"] for order in filtered_orders)

        # Group sales by product category
        sales_by_category = {}
        for order in filtered_orders:
            for item in order["items"]:
                category = item["product"]["category"]
                if category not in sales_by_category:
                    sales_by_category[category] = 0
                sales_by_category[category] += item["item_total"]

        # Format report
        report = {
            "start_date": start_date,
            "end_date": end_date,
            "total_sales": total_sales,
            "order_count": len(filtered_orders),
            "sales_by_category": sales_by_category,
        }

        # Print report
        print("=== SALES REPORT ===")
        print(
            f"Period: {start_date.strftime('%Y-%m-%d')} to {end_date.strftime('%Y-%m-%d')}"
        )
        print(f"Total Sales: ${total_sales:.2f}")
        print(f"Order Count: {len(filtered_orders)}")
        print("Sales by Category:")
        for category, amount in sales_by_category.items():
            print(f"  {category}: ${amount:.2f}")

        # Save report to database
        self.db.save_report(report)

        return report


class PaymentProcessor:
    def process_payment(self, payment_info, amount):
        print(f"Processing payment of ${amount:.2f} using {payment_info['type']}...")

        # Simulate payment processing
        if payment_info["type"] == "credit_card":
            # Validate credit card (very basic validation)
            if len(payment_info["card_number"]) != 16:
                return {"success": False, "error": "Invalid card number"}

            # Connect to payment gateway (simulated)
            print("Connecting to credit card payment gateway...")
            # Process payment logic would go here

            # Mock successful response
            return {"success": True}
        elif payment_info["type"] == "paypal":
            # Connect to PayPal API (simulated)
            print("Connecting to PayPal API...")
            # Process payment logic would go here

            # Mock successful response
            return {"success": True}
        else:
            return {"success": False, "error": "Unsupported payment method"}
