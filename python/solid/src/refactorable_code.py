# A badly designed e-commerce system with multiple SOLID violations


from solid.src.database import Database
from solid.src.services.notification_service import NotificationService


class OnlineStore:
    def __init__(self):
        self.notification_service = NotificationService()
        self.db = Database()

    # Order Processing

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
