class FakeNotificationService:
    def __init__(self):
        self.send_email_count = 0
        self.send_sms_count = 0

    def send_email(self, to: str, subject: str, body: str) -> bool:
        self.send_email_count += 1
        return True

    def send_sms(self, to: str, message: str) -> bool:
        self.send_sms_count += 1
        return True
