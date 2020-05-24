import sys
from typing import List


def build(values: List[int]) -> List[int]:
    # O(n)
    n = len(values)
    st: List[int] = [0] * (n*4)

    def f(vertex: int, lo: int, hi: int) -> None:
        if lo == hi:
            st[vertex] = values[lo]
            return

        mid = (lo+hi) // 2
        child_l = vertex*2+1
        child_r = vertex*2+2

        f(child_l, lo, mid)
        f(child_r, mid+1, hi)

        st[vertex] = st[child_l] + st[child_r]

    f(0, lo=0, hi=n-1)
    return st


def update(st: List[int], values: List[int], index: int, value: int) -> None:
    # O(logn)
    diff = value - values[index]
    values[index] = value

    def f(st_index: int, lo: int, hi: int) -> None:
        if index < lo or index > hi:
            return

        st[st_index] += diff

        if lo != hi:
            mid = (lo+hi) // 2
            f(st_index*2+1, lo, mid)
            f(st_index*2+2, mid+1, hi)

    n = len(values)
    f(0, lo=0, hi=n-1)


def summate(st: List[int], values: List[int], start: int, end: int) -> int:
    # O(logn)
    def f(vertex: int, lo: int, hi: int) -> int:
        if start > hi or end < lo:
            return 0

        if start <= lo and end >= hi:
            return st[vertex]

        mid = (lo+hi) // 2
        return f(vertex*2+1, lo, mid) + f(vertex*2+2, mid+1, hi)

    n = len(values)
    return f(0, lo=0, hi=n-1)


def main():
    n, m, k = map(int, sys.stdin.readline().split())

    values = []
    for _ in range(n):
        x = int(sys.stdin.readline())
        values.append(x)

    st = build(values)

    for _ in range(m+k):
        a, b, c = map(int, sys.stdin.readline().split())

        if a == 1:
            update(st, values, index=b-1, value=c)
        elif a == 2:
            print(summate(st, values, start=b-1, end=c-1))


if __name__ == '__main__':
    main()
