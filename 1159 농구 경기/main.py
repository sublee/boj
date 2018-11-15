from collections import Counter

# Count first letters during accepting inputs.
counter = Counter()

n = int(input())
for _ in range(n):
    name = input()
    counter[name[0]] += 1

# Collect common first letters.
letters = []
for l, c in counter.most_common():
    if c < 5:
        break
    letters.append(l)

# Print them in alphabetical order.
if letters:
    letters.sort()
    print(''.join(letters))
else:
    print('PREDAJA')
