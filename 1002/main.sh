#!/usr/lib/env bash
# https://www.acmicpc.net/problem/1002

# NOTE(sublee): We cannot use bc, seq in acmicpc.net Bash.

main() {
  # 첫째 줄에 테스트 케이스의 개수 T가 주어진다.
  read T

  # 각 테스트 케이스는 다음과 같이 구성되어있다. 한 줄에 x1, y1, r1, x2, y2,
  # r2가 주어진다. x1, y1, x2, y2는 -10,000보다 크거나 같고, 10,000보다 작거나
  # 같은 정수이고, r1, r2는 10,000보다 작거나 같은 자연수이다.
  i=1
  while [[ "$i" -le "$T" ]]; do

    read -a nums
    solve "${nums[@]}"

    let i+=1
  done
}

solve() {
  # 각 테스트 케이스마다 류재명이 있을 수 있는 위치의 수를 출력한다.
  # 만약 류재명이 있을 수 있는 위치의 개수가 무한대일 경우에는 -1을 출력한다.
  local x1="$1"
  local y1="$2"
  local r1="$3"

  local x2="$4"
  local y2="$5"
  local r2="$6"

  # the same point, same radius.
  if [[ "$x1" -eq "$x2" ]] && [[ "$y1" -eq "$y2" ]] && [[ "$r1" -eq "$r2" ]]
  then
    echo -1
    return
  fi

  # https://mathbang.net/101
  let "dx=x2-x1"
  let "dy=y2-y1"
  let "d_sq=dx*dx + dy*dy"

  let "a=r2-r1"
  let "a_sq=a*a"

  let "b=r2+r1"
  let "b_sq=b*b"

  if [[ "$d_sq" -eq "$a_sq" ]] || [[ "$d_sq" -eq "$b_sq" ]]; then
    # d = r-r' or r+r', 1 intersection
    echo 1
  elif [[ "$d_sq" -gt "$a_sq" ]] && [[ "$d_sq" -lt "$b_sq" ]]; then
    # r-r' < d < r+r', 2 intersections
    echo 2
  else
    # otherwise, no intersections
    echo 0
  fi
}

main
