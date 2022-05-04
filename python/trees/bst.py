from dataclasses import dataclass
from typing import Optional


@dataclass
class Node:
    """A component of a binary search tree"""

    key: int
    left: Optional["Node"] = None
    right: Optional["Node"] = None

    def insert(self, key: int) -> None:
        if key < self.key:
            if self.left:
                self.left.insert(key)
            else:
                self.left = Node(key)
        else:
            if self.right:
                self.right.insert(key)
            else:
                self.right = Node(key)

    def search(self, key: int) -> bool:
        if key == self.key:
            return True

        if key < self.key and self.left:
            return self.left.search(key)
        elif key > self.key and self.right:
            return self.right.search(key)

        return False

    def delete(self, key: int) -> None:
        """TODO: Remove the node by a key from the BST"""


def main():
    bst = Node(100)
    print("Initial BST:", bst)

    bst.insert(52)
    bst.insert(203)
    bst.insert(19)
    bst.insert(76)
    bst.insert(150)
    bst.insert(310)
    bst.insert(7)
    bst.insert(24)
    bst.insert(88)
    bst.insert(276)
    print("BST after adding a few leaves:", bst)

    if bst.search(76):
        print("Node with a key = 76 has been found in BST")
    else:
        print("Node with a key = 76 has not been found in BST")

    if bst.search(400):
        print("Node with a key = 400 has been found in BST")
    else:
        print("Node with a key = 400 has not been found in BST")


if __name__ == "__main__":
    main()
