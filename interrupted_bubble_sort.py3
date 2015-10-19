import sys
import math
import string
import doctest
import itertools


DEBUG = True


def solve(numbers, iter_):
    iter_ = min(iter_, len(numbers))
    nIters = 0
    while nIters < iter_:
        currentPos = 0
        while currentPos < len(numbers) - 1:
            if numbers[currentPos] > numbers[currentPos+1]:
                numbers[currentPos], numbers[currentPos + 1] = numbers[currentPos + 1], numbers[currentPos]
            currentPos += 1
        nIters += 1
    return " ".join(map(str, numbers))

if len(sys.argv) > 2 and sys.argv[2] == "--secrettests":
    doctest.testmod()
    sys.exit()

test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    test = test.strip().split("|")
    numbers = [int(s) for s in test[0].strip().split()]
    iter_ = int(test[1].strip())
    print(solve(numbers, iter_))

test_cases.close()
