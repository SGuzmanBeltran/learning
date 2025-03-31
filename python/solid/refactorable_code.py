# A badly designed e-commerce system with multiple SOLID violations


class OnlineStore:
    def __init__(self):
        self.products = []
        self.users = []
        self.orders = []
        self.payment_processor = PaymentProcessor()
        self.notification_service = NotificationService()
        self.db = Database()

    # Product Management
    def add_product(self, id, name, price, category, description, stock_quantity):
        product = {
            "id": id,
            "name": name,
            "price": price,
            "category": category,
            "description": description,
            "stock_quantity": stock_quantity,
        }
        self.products.append(product)
        self.db.save_product(product)
        print(f"Product {name} added to inventory.")

    def update_product_stock(self, product_id, new_quantity):
        product = next((p for p in self.products if p["id"] == product_id), None)
        if product:
            product["stock_quantity"] = new_quantity
            self.db.update_product(product)
            print(f"Product {product['name']} stock updated to {new_quantity}.")

            # Send notification if stock is low
            if new_quantity < 10:
                self.notification_service.send_email(
                    "admin@store.com",
                    f"Low stock alert for {product['name']}",
                    f"The stock for {product['name']} is low ({new_quantity} remaining).",
                )

    def get_product_details(self, product_id):
        return next((p for p in self.products if p["id"] == product_id), None)

    # User Management
    def register_user(self, id, name, email, password, address, payment_info):
        # Hash the password (very insecure way)
        hashed_password = password[::-1]

        user = {
            "id": id,
            "name": name,
            "email": email,
            "password": hashed_password,
            "address": address,
            "payment_info": payment_info,
            "is_admin": False,
        }

        self.users.append(user)
        self.db.save_user(user)

        # Send welcome email
        self.notification_service.send_email(
            email, "Welcome to our store!", f"Dear {name}, thank you for registering!"
        )

        print(f"User {name} registered successfully.")

    def get_user_details(self, user_id):
        return next((u for u in self.users if u["id"] == user_id), None)

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


class NotificationService:
    def send_email(self, to, subject, body):
        print(f"Sending email to {to}")
        print(f"Subject: {subject}")
        print(f"Body: {body}")
        # In a real implementation, this would connect to an email service
        return True

    def send_sms(self, to, message):
        print(f"Sending SMS to {to}: {message}")
        # In a real implementation, this would connect to an SMS service
        return True


class Database:
    def save_product(self, product):
        print(f"Saving product to database: {product['name']}")
        # In a real implementation, this would save to a database
        return True

    def update_product(self, product):
        print(f"Updating product in database: {product['name']}")
        # In a real implementation, this would update a database record
        return True

    def save_user(self, user):
        print(f"Saving user to database: {user['name']}")
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


# Example usage
if __name__ == "__main__":
    import datetime

    store = OnlineStore()

    # Add products
    store.add_product(
        1, "Laptop", 999.99, "Electronics", "Powerful laptop for developers", 20
    )
    store.add_product(
        2, "Headphones", 149.99, "Electronics", "Noise-cancelling headphones", 30
    )

    # Register a user
    store.register_user(
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

    # Create an order
    store.create_order(
        1, [{"product_id": 1, "quantity": 1}, {"product_id": 2, "quantity": 2}]
    )

    # Generate a sales report
    start_date = datetime.datetime(2023, 1, 1)
    end_date = datetime.datetime(2023, 12, 31)
    store.generate_sales_report(start_date, end_date)
