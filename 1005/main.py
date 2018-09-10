from collections import deque
import sys


def parse_n_n(s):
    x, y = s.split(maxsplit=1)
    return int(x), int(y)


def main(f=sys.stdin):
    # 첫째 줄에는 테스트케이스의 개수 T가 주어진다.
    num_tests = int(f.readline())
    for x in range(num_tests):
        # 각 테스트 케이스는 다음과 같이 주어진다.
        # 첫째 줄에 건물의 개수 N 과
        # 건물간의 건설순서규칙의 총 개수 K이 주어진다.
        num_units, num_deps = parse_n_n(f.readline())

        # 둘째 줄에는 각 건물당 건설에 걸리는 시간 D가
        # 공백을 사이로 주어진다.
        delays_s = f.readline()
        delays_ss = delays_s.split(maxsplit=num_units)
        delays = [None] + [int(d) for d in delays_ss]

        # 셋째 줄부터 K+2줄까지 건설순서 X Y가 주어진다.
        # (이는 건물 X를 지은 다음에 건물 Y를 짓는 것이 가능하다는 의미이다)
        deps = []
        for y in range(num_deps):
            i, j = parse_n_n(f.readline())
            deps.append((i, j))

        # 마지막 줄에는 백준이가 승리하기 위해
        # 건설해야 할 건물의 번호 W가 주어진다.
        the_unit = int(f.readline())

        print(solve(delays, deps, the_unit))


def solve(delays, deps, the_unit):
    # organize deps for easy to use
    refs = [None] + [set() for d in delays]
    indegrees = [0] * len(delays)

    for i, j in deps:
        refs[i].add(j)
        indegrees[j] += 1

    # topological sort
    search_q = deque()
    topology = []

    while len(topology) < len(indegrees)-1:
        for i, x in enumerate(indegrees[1:], start=1):
            if x == 0:
                search_q.append(i)

        while search_q:
            i = search_q.popleft()
            topology.append(i)
            for j in refs[i]:
                indegrees[j] -= 1
                if indegrees[j] == 0:
                    search_q.append(j)

    # analyze delays from dependencies
    dep_delays = [0] * len(delays)
    for i in topology:
        for j in refs[i]:
            dep_delays[j] = max(dep_delays[j], delays[i] + dep_delays[i])

    return dep_delays[the_unit] + delays[the_unit]


if __name__ == '__main__':
    main()
