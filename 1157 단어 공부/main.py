from collections import Counter
import sys

counter = Counter()
word = sys.stdin.read()

for x in word:
    if x.isalpha():
        counter[x.upper()] += 1

top2 = counter.most_common(2)

if len(top2) == 2 and top2[0][1] == top2[1][1]:
    print('?')
else:
    print(top2[0][0])
