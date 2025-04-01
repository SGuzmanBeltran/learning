class PaymentService:
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
