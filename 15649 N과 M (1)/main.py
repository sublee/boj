from itertools import permutations


n, m = input().split()
n, m = int(n), int(m)


for seq in permutations(range(1, n+1), m):
    print(' '.join(str(x) for x in seq))
