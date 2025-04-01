class Order:
    def __init__(self, id, user_id, items, total_amount, tax, status, created_at):
        self.id = id
        self.user_id = user_id
        self.items = items
        self.total_amount = total_amount
        self.tax = tax
        self.status = status
        self.created_at = created_at
