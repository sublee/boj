from itertools import combinations


n, m = input().split()
n, m = int(n), int(m)


for seq in combinations(range(1, n+1), m):
    print(' '.join(str(x) for x in seq))
