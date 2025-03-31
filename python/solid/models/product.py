class Product:
    def __init__(
        self,
        id: int,
        name: str,
        price: float,
        category: str,
        description: str,
        stock_quantity: int,
    ):
        self.id = id
        self.name = name
        self.price = price
        self.category = category
        self.description = description
        self.stock_quantity = stock_quantity
