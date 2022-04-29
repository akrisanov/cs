from optparse import Option

import math
from typing import List, Optional


def binary_search(key: int, array: List[int]) -> Optional[int]:
    """Search for an id in an array. Return an index of the id if found."""
    if len(array) == 0:
        return None

    middle = math.floor(len(array) / 2)

    if key == array[middle]:
        return middle

    if key < array[middle]:
        return binary_search(key, array[:middle])

    if id > array[middle]:
        return binary_search(key, array[middle + 1 :])


if __name__ == "__main__":
    array = [10, 3, 45, 4, 0, -1, 2, 77, 2, 3]
    array.sort()

    key = 2
    result = binary_search(key, array)

    if result:
        print(f"{key} has been found at the index {result} in the array of {array}")
    else:
        print(f"{key} is not found in the array of {array}")
