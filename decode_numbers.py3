import sys
import math
import string
import doctest
import itertools


DEBUG = True


def count(start, line):
    cutline = line[start:]
    if len(cutline) == 1:
        return 1
    elif len(cutline) == 2:
        if int(cutline) <= 26:
            return 2
        else:
            return 1
    else:
        if int(cutline[:2]) <= 26:
            return count(start+1, line) + count(start+2, line)
        else:
            return count(start+1, line)


def solve(line):
    return count(0, line)

if len(sys.argv) > 2 and sys.argv[2] == "--secrettests":
    doctest.testmod()
    sys.exit()

test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    print(solve(test.strip()))

test_cases.close()
