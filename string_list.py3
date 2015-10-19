import sys
import math
import string
import doctest
import itertools


DEBUG = True


def solve(n, s):
    chars = [x for x in s]
    args = [chars] * n
    perms = itertools.product(*args)
    #print(chars)

    perms = set(perms)
    out = []
    for p in perms:
        #print("appending", p)
        out.append(''.join([x for x in p]))
    out.sort()
    return ','.join(out)

if len(sys.argv) > 2 and sys.argv[2] == "--secrettests":
    doctest.testmod()
    sys.exit()

test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    test = test.strip().split(",")
    n = int(test[0].strip())
    print(solve(n, test[1]))

test_cases.close()
