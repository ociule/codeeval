import sys
import math
import string
import doctest


LETTERS = string.ascii_uppercase


def solve(n):
    """
    >>> solve(1)
    'A'
    >>> solve(26)
    'Z'
    >>> solve(27)
    'AA'
    >>> solve(28)
    'AB'
    >>> solve(52)
    'AZ'
    >>> solve(676)
    'YZ'
    >>> solve(677)
    'ZA'
    >>> solve(3702)
    'ELJ'
    """
    #if n <= 26:
    #    return LETTERS[n - 1]
    #if n <= 676:
    #    fd = n // 26
    #    ld = n % 26
    #    return LETTERS[fd - 1] + LETTERS[ld - 1]
    out = []
    while n > 0:

        r = n % 26
        if r == 0:
            pass
            out.append('Z')
            n = n // 26 - 1
        else:
            out.append(LETTERS[r - 1])
            n = n // 26
    out.reverse()
    return ''.join(out)

if len(sys.argv) > 2 and sys.argv[2] == "--secrettests":
    DEBUG = True
    doctest.testmod()
    sys.exit()

test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    test = test.strip()
    n = int(test)
    print(solve(n))

test_cases.close()

