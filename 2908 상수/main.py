import sys
a, _, b = sys.stdin.read().strip().partition(' ')
a, b = a[::-1], b[::-1]
print(max(a, b))
