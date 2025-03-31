class User:
    def __init__(
        self,
        id: int,
        name: str,
        email: str,
        password: str,
        address: dict,
        payment_info: dict,
        is_admin: bool = False,
    ):
        self.id = id
        self.name = name
        self.email = email
        self.password = password
        self.address = address
        self.payment_info = payment_info
        self.is_admin = is_admin
