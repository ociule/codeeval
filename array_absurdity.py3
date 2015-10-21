import sys
import math
import string
import doctest
import itertools


DEBUG = True


def solve(n, numbers):
    s = sum(numbers)
    cs = (n - 1) * (n - 2) / 2
    return int(s - cs)

if len(sys.argv) > 2 and sys.argv[2] == "--secrettests":
    doctest.testmod()
    sys.exit()

test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    test = test.strip().split(";")
    n = int(test[0].strip())
    numbers = [int(s) for s in test[1].strip().split(",")]
    print(solve(n, numbers))

test_cases.close()
