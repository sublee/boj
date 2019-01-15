from itertools import product


n, m = input().split()
n, m = int(n), int(m)


for seq in product(range(1, n+1), repeat=m):
    print(' '.join(str(x) for x in seq))
