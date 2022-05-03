from typing import Any, Dict, Optional


class ArrayList:
    """Implement an array list without using Python list objects."""

    def __init__(self, data: Optional[Dict[int, Any]] = None, length: int = 0) -> None:
        if not data:
            self._data = {}
        self.length = length

    def __str__(self) -> str:
        elements = map(str, self._data.values())
        elements = ", ".join(elements)
        return f"ArrayList of size {self.length}: [{elements}]"

    def push(self, value) -> None:
        """Add an item to the end of the array"""
        self._data[self.length] = value
        self.length += 1

    def pop(self) -> Optional[Any]:
        """
        Remove the last item in the array and returns it.
        Note, we do not shrink the dict here, nor return a new ArrayList.
        """
        return self.delete(self.length - 1)

    def get(self, index: int) -> Optional[int]:
        """Return that item from the array"""
        if self.length > 0:
            return self._data[index]
        return None

    def delete(self, index: int) -> Optional[Any]:
        """Remove item from the array and collapse the array"""
        last_index = self.length - 1

        if self.length < 1 or index > self.length - 1:
            return None

        removed = self._data[index]

        for i in range(index, self.length - 1):
            self._data[i] = self._data[i + 1]

        del self._data[last_index]
        self.length -= 1

        return removed


def main():
    arrlist = ArrayList()
    print("Initial ", arrlist)

    arrlist.push(15)
    print(arrlist)

    arrlist.push(24)
    arrlist.push(15)
    arrlist.push(19)
    arrlist.push(28)
    arrlist.push(26)
    print(arrlist)

    print("Remove the last element in the array: ", arrlist.pop())
    print(arrlist)

    print("Remove 2nd element in the array: ", arrlist.get(1))
    arrlist.delete(1)
    print(arrlist)


if __name__ == "__main__":
    main()
