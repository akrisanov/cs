from dataclasses import dataclass
from typing import Any, Optional


@dataclass
class Node:
    data: Any
    next: Optional["Node"] = None


@dataclass
class LinkedList:
    head: Optional[Node] = None
    tail: Optional[Node] = None
    length: int = 0

    def __str__(self) -> str:
        to_print = self.head

        output = []

        for _ in range(0, self.length):
            output.append(f'"{to_print.data}"')
            to_print = to_print.next

        return "{" + ", ".join(output) + "}"

    def push(self, data: Any) -> None:
        new_node = Node(data)
        if self.head:
            self.tail.next = new_node
        else:
            self.head = new_node
        self.tail = new_node
        self.length += 1

    def get(self, index) -> Optional[Any]:
        node = self._find(index)
        if node:
            return node.data
        return None

    def pop(self) -> Optional[Any]:
        return self.delete(self.length - 1)

    def delete(self, index) -> Optional[Any]:
        if self.length < 1 or index < 0 or index > self.length - 1:
            return None

        to_remove = None

        if index == 0:
            if self.head:
                to_remove = self.head.data
                self.head = self.head.next
        else:
            prev = self._find(index)
            if prev:
                to_remove = prev.next.data
                prev.next = prev.next.next

        self.length -= 1

        return to_remove

    def _find(self, index) -> Optional[Node]:
        if index < 0 or index >= self.length:
            return None

        current = self.head
        for _ in range(0, index - 1):
            current = current.next
        return current


def main():
    ll = LinkedList()
    print("Empty linked list:", ll)

    ll.push("Hello")
    ll.push(",")
    ll.push("World")
    ll.push("!")
    print("A few items have been pushed into the list:", ll)

    print(f'The following item has been removed from the end of the list: "{ll.pop()}"')
    print(f'The following item has been removed from the 1st index: "{ll.delete(1)}"')

    print(ll)

    ll.delete(-1)
    ll.delete(100)
    print(ll)

    print(f'The following item has been removed from the zero index: "{ll.delete(0)}"')
    print(ll)
    print(f'The following item has been removed from the zero index: "{ll.delete(0)}"')
    print(ll)


if __name__ == "__main__":
    main()
