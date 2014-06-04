import sys
import math


D_EDGES = 6


def how_many_bats(L, d, n, bats):
    segments = [D_EDGES - d] + bats + [L - D_EDGES + d]
    count = 0
    for ix, s_end in enumerate(segments[1:]):
        s_start = segments[ix]
        l = s_end - s_start
        count += math.floor(l / d) - 1
    return count 


test_cases = open(sys.argv[1], 'r')
for test in test_cases:
    test = test.strip().split()
    L = int(test[0])
    d = int(test[1])
    n = int(test[2])
    bats = test[3:]
    bats = lis(map(int, bats))
    bats.sort()
    print(how_many_bats(L, d, n, bats))

test_cases.close()
