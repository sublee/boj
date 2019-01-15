from itertools import combinations

_, m = input().split()
m = int(m)

nums = [int(x) for x in input().split()]
nums.sort()

for seq in combinations(nums, m):
    print(' '.join(str(x) for x in seq))
