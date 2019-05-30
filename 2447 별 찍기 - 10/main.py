from __future__ import print_function


def f(n, i, j, disp):
    if n == 1:
        disp[i][j] = 1
        return

    k = n // 3
    f(k, i, j, disp)
    f(k, i, j+k, disp)
    f(k, i, j+k+k, disp)
    f(k, i+k, j, disp)
    f(k, i+k, j+k+k, disp)
    f(k, i+k+k, j, disp)
    f(k, i+k+k, j+k, disp)
    f(k, i+k+k, j+k+k, disp)


n = int(input())
disp = [[0] * n for _ in range(n)]
f(n, 0, 0, disp)


for line in disp:
    for x in line:
        if x:
            print('*', end='')
        else:
            print(' ', end='')
    print()
