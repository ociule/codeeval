import sys, doctest
from collections import defaultdict


def levenshtein(a, b):
    """Calculates the Levenshtein distance between a and b.
    From http://hetland.org/coding/python/levenshtein.py
    >>> levenshtein("kitten", "sitting")
    3
    >>> levenshtein("kitten", "sitten")
    1
    >>> levenshtein("sittin", "sitting")
    1
    """
    n, m = len(a), len(b)
    if n > m:
        # Make sure n <= m, to use O(min(n,m)) space
        a,b = b,a
        n,m = m,n
        
    current = range(n+1)
    for i in range(1,m+1):
        previous, current = current, [i]+[0]*n
        for j in range(1,n+1):
            add, delete = previous[j]+1, current[j-1]+1
            change = previous[j-1]
            if a[j-1] != b[i-1]:
                change = change + 1
            current[j] = min(add, delete, change)
            
    return current[n]


def main():
    test_cases = open(sys.argv[1], 'r')
    word = sys.argv[2]

    still_tests = True
    for test in test_cases:
        test = test.strip()
        if still_tests:
            if test == "END OF INPUT":
                still_tests = False
                continue
        else:
            if levenshtein(test, word) == 1:
                print test
        

    test_cases.close()

if len(sys.argv) == 1:
    doctest.testmod()
else:
    main()
