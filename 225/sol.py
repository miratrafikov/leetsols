import queue

class MyStack:

    def __init__(self):
        self.q = []

    def push(self, x: int) -> None:
        self.q.append(x)

    def pop(self) -> int:
        element = None
        for i in range(len(self.q) - 1):
            element = self.q.pop(0)
            self.push(element)
        return self.q.pop(0)

    def top(self) -> int:
        element = None
        for i in range(len(self.q)):
            element = self.q.pop(0)
            self.push(element)
        return element

    def empty(self) -> bool:
        if len(self.q) == 0:
            return True
        return False


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()
