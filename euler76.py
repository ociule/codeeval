# Solution to Euler 76


target = 10000
ways = [0] * (target + 1)
ways[0] = 1

for i in range(1, target):
    for j in range(i, target + 1):
        ways[j] += ways[j - i]

print ways.pop()
