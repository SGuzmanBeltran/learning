# Kata: https://kata-log.rocks/banking-kata


from datetime import datetime
from dataclasses import dataclass


@dataclass
class History:
    date: datetime
    amount: int
    balance: int

    def __str__(self) -> str:
        formatted_date = self.date.strftime("%d.%m.%Y")
        return f"{formatted_date}    {self.amount}   {self.balance}"


class Account:
    def __init__(self) -> None:
        self.balance = 0
        self.history: list[History] = []

    def deposit(self, quantity: int) -> None:
        self.balance += quantity
        history = History(datetime.now(), quantity, self.balance)
        self.history.append(history)

    def withdraw(self, quantity: int) -> None:
        if self.balance - quantity < 0:
            return
        self.balance -= quantity
        history = History(datetime.now(), quantity, self.balance)
        self.history.append(history)

    def printStatement(self) -> str:
        result = "Date     Amount      Balance \n"
        for history in self.history:
            result += f"{history} \n"
        return result


account = Account()
account.deposit(500)
account.withdraw(100)
print(account.printStatement())
