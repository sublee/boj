from itertools import combinations_with_replacement

_, m = input().split()
m = int(m)

nums = [int(x) for x in input().split()]
nums.sort()

for seq in combinations_with_replacement(nums, m):
    print(' '.join(str(x) for x in seq))
