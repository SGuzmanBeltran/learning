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
