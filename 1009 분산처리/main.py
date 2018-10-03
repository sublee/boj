import sys
for x in range(int(sys.stdin.readline())):
    a, b = sys.stdin.readline().split()
    a, b = int(a), int(b)
    print(1 if b == 0 else (a ** (b % 4 + 4) % 10) or 10)
