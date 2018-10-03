// https://www.acmicpc.net/problem/1004
var fs = require('fs')

var input = fs.readFileSync('/dev/stdin').toString().split('\n')

// 입력의 첫 줄에는 테스트 케이스의 개수 T가 주어진다.
var t = parseInt(input.shift())

// 그 다음 줄부터 각각의 테스트케이스에 대해
for (var i = 0; i < t; i++) {
  // 첫째 줄에 출발점 (x1, y1)과 도착점 (x2, y2)이 주어진다.
  var points = input.shift().split(' ')

  var x1 = parseInt(points[0])
  var y1 = parseInt(points[1])
  var x2 = parseInt(points[2])
  var y2 = parseInt(points[3])

  delete points

  // 두 번째 줄에는 행성계의 개수 n이 주어지며
  var n = parseInt(input.shift())

  // 세 번째 줄부터 n줄에 걸쳐 행성계의 중점과 반지름 (cx, cy, r)이 주어진다.
  var planets = []
  for (var j = 0; j < n; j++) {
    var planet = input.shift().split(' ')

    var cx = parseInt(planet[0])
    var cy = parseInt(planet[1])
    var r  = parseInt(planet[2])

    planets.push([cx, cy, r])
  }

  solve(x1, y1, x2, y2, planets)
}

// 각 테스트 케이스에 대해 어린 왕자가 거쳐야 할
// 최소의 행성계 진입/이탈 회수를 출력한다.
function solve(x1, y1, x2, y2, planets) {
  var counter = []

  planets.forEach(function (planet) {
    var cx = planet[0]
    var cy = planet[1]
    var r  = planet[2]

    if (isInCircle(x1, y1, cx, cy, r) ^ isInCircle(x2, y2, cx, cy, r)) {
      counter.push(1)
    }
  })

  console.log(counter.length)
}

function isInCircle(x, y, cx, cy, r) {
  var dSq = Math.pow(x-cx, 2) + Math.pow(y-cy, 2)
  return dSq <= Math.pow(r, 2)
}
