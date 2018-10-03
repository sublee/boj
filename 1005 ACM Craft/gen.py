"""
Generates an input text.

   $ python3.6 gen.py 4 4
   1
   4 4
   1 30 30 10
   3 4
   1 3
   1 4
   1 2
   4

"""
import itertools
import random
import sys

n = int(sys.argv[1])
k = int(sys.argv[2])

deps = list(itertools.combinations(range(1, n+1), 2))
k = min(k, len(deps))

print(1)
print(n, k)
for x in range(n):
    print(random.choice([1, 10, 20, 30, 100]), end=' ')
print()

random.shuffle(deps)

for x in range(k):
    print(deps[x][0], deps[x][1])

print(n)
