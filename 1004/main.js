const fs = require('fs')

const input = fs.readFileSync('/dev/stdin').toString().split('\n')

// 입력의 첫 줄에는 테스트 케이스의 개수 T가 주어진다.
let t = parseInt(input.shift())

// 그 다음 줄부터 각각의 테스트케이스에 대해
for (let i = 0; i < t; i++) {
  // 첫째 줄에 출발점 (x1, y1)과 도착점 (x2, y2)이 주어진다.
  let points = input.shift().split(' ')

  let x1 = parseInt(points[0])
  let y1 = parseInt(points[1])
  let x2 = parseInt(points[2])
  let y2 = parseInt(points[3])

  delete points

  // 두 번째 줄에는 행성계의 개수 n이 주어지며
  let n = parseInt(input.shift())

  // 세 번째 줄부터 n줄에 걸쳐 행성계의 중점과 반지름 (cx, cy, r)이 주어진다.
  let planets = []
  for (let j = 0; j < n; j++) {
    let planet = input.shift().split(' ')

    let cx = parseInt(planet[0])
    let cy = parseInt(planet[1])
    let r  = parseInt(planet[2])

    planets.push([cx, cy, r])
  }

  solve(x1, y1, x2, y2, planets)
}

function solve(x1, y1, x2, y2, planets) {
}
