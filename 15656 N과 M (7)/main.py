from itertools import product

_, m = input().split()
m = int(m)

nums = [int(x) for x in input().split()]
nums.sort()

for seq in product(nums, repeat=m):
    print(' '.join(str(x) for x in seq))
