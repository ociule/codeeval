import sys
import math
import string
import doctest
import itertools


N_OPS = 3
SOL = 42

def space_len(n):
    """
    >>> space_len(5)
    121
    >>> space_len(1)
    1
    """
    return sum([N_OPS ** i for i in range(n)])

def solve(numbers):
    perms = list(itertools.permutations(numbers, 5))
    for perm in perms:
        if sums42(list(perm)):
            return 'YES' 
    return 'NO'

def sums42(numbers):
    """
    >>> sums42([44, 6, 1, 49, 47])
    'NO'
    >>> sums42([34, 1, 49, 2, 21])
    'YES'
    >>> sums42([31, 38, 27, 51, 18])
    'NO'
    >>> sums42([42, 0, 0, 0, 0])
    'YES'
    >>> sums42([38, 1, 1, 1, 1])
    'YES'
    >>> sums42([60, 1, 3, 5, 20])
    'YES'
    """
    space = [0] * space_len(len(numbers))
    space[0] = numbers[0]

    start = 0
    end = 1

    curr_sol = end
    for current_num in range(len(numbers) - 1):
        n2 = numbers[current_num + 1]
        last = (current_num == len(numbers) - 2)

        for num in space[start:end]:
            space[curr_sol] = num + n2
            curr_sol += 1
            space[curr_sol] = num - n2
            curr_sol += 1
            space[curr_sol] = num * n2
            curr_sol += 1
        start = end
        end = curr_sol 
        #print space, start, end
        #print space[start:end]
        if last and 42 in space[start:end]:
            return True 
        
    
    return False


if len(sys.argv) > 2 and sys.argv[2] == "--secrettests":
    DEBUG = True
    doctest.testmod()
    sys.exit()

test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    test = test.strip()
    numbers = [int(s) for s in test.split()]
    print(solve(numbers))

test_cases.close()
