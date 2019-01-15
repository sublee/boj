from itertools import permutations

_, m = input().split()
m = int(m)

nums = [int(x) for x in input().split()]
nums.sort()

for seq in permutations(nums, m):
    print(' '.join(str(x) for x in seq))
