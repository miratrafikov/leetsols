class Node:
    def __init__(self, val, next = None):
        self.val = val
        self.next = next


class MyLinkedList:
    
    def __init__(self):
        self.head = None
        self.length = 0
    
    def get(self, index: int) -> int:
        if index > self.length - 1:
            return -1
        curr = self.head
        for i in range(index):
            curr = curr.next
        return curr.val

    def addAtHead(self, val: int) -> None:
        if not self.head:
            self.head = Node(val)
        else:
            self.head = Node(val, self.head)
        self.length += 1

    def addAtTail(self, val: int) -> None:
        if not self.head:
            self.head = Node(val)
        else:
            curr = self.head
            while curr.next:
                curr = curr.next
            curr.next = Node(val)
        self.length += 1
            

    def addAtIndex(self, index: int, val: int) -> None:
        if index > self.length:
            return
        if index == 0:
            self.head = Node(val, self.head)
        else:
            prev = None
            curr = self.head
            for i in range(index):
                prev = curr
                curr = curr.next
            prev.next = Node(val)
            prev.next.next = curr
        self.length += 1

    def deleteAtIndex(self, index: int) -> None:
        if index > self.length - 1:
            return
        if index == 0:
            self.head = self.head.next
        else:
            prev = None
            curr = self.head
            for i in range(index):
                prev = curr
                curr = curr.next
            prev.next = curr.next
        self.length -= 1

# Your MyLinkedList object will be instantiated and called as such:
# obj = MyLinkedList()
# param_1 = obj.get(index)
# obj.addAtHead(val)
# obj.addAtTail(val)
# obj.addAtIndex(index,val)
# obj.deleteAtIndex(index)
