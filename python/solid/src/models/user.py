class PaymentInfo:
    def __init__(self, type: str, card_number: str):
        self.type = type
        self.card_number = card_number


class User:
    def __init__(
        self,
        id: int,
        name: str,
        email: str,
        password: str,
        address: dict,
        payment_info: PaymentInfo,
        is_admin: bool = False,
    ):
        self.id = id
        self.name = name
        self.email = email
        self.password = password
        self.address = address
        self.payment_info = payment_info
        self.is_admin = is_admin
