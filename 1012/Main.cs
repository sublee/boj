using System;
using System.Collections.Generic;

namespace Acmicpc
{
    using Vector2 = Tuple<int, int>;

    class Acmicpc1012
    {
        static void Main(string[] args)
        {
            // 입력의 첫 줄에는 테스트 케이스의 개수 T가 주어진다.
            var t = int.Parse(Console.ReadLine());

            for (int i = 0; i < t; i++)
            {
                // 그 다음 줄부터 각각의 테스트 케이스에 대해 첫째 줄에는
                // 배추를 심은 배추밭의 가로길이 M(1 ≤ M ≤ 50)과 세로길이 N(1 ≤ N ≤ 50),
                // 그리고 배추가 심어져 있는 위치의 개수 K(1 ≤ K ≤ 2500)이 주어진다.
                var parts = Console.ReadLine().Split(new Char[] { ' ' }, 3);
                var k = int.Parse(parts[2]);

                // 그 다음 K줄에는 배추의 위치 X(0 ≤ X ≤ M-1), Y(0 ≤ Y ≤ N-1)가 주어진다.
                var positions = new HashSet<Vector2>(k);
                for (int j = 0; j < k; j++)
                {
                    parts = Console.ReadLine().Split(new Char[] { ' ' }, 2);
                    var x = int.Parse(parts[0]);
                    var y = int.Parse(parts[1]);
                    positions.Add(new Vector2(x, y));
                }

                Console.WriteLine(Solve(positions));
            }
        }

        static int Solve(HashSet<Vector2> positions)
        {
            var visited = new HashSet<Vector2>(positions.Count);

            int count = 0;

            foreach (var position in positions)
            {
                if (visited.Contains(position))
                {
                    continue;
                }
                BFS(position, positions, visited);
                count++;
            }

            return count;
        }

        static void BFS(Vector2 position, HashSet<Vector2> positions, HashSet<Vector2> visited)
        {
            var q = new Queue<Vector2>();
            q.Enqueue(position);

            while (q.Count != 0)
            {
                position = q.Dequeue();

                if (visited.Contains(position) || !positions.Contains(position))
                {
                    continue;
                }

                visited.Add(position);

                q.Enqueue(new Vector2(position.Item1 + 1, position.Item2));
                q.Enqueue(new Vector2(position.Item1 - 1, position.Item2));
                q.Enqueue(new Vector2(position.Item1, position.Item2 + 1));
                q.Enqueue(new Vector2(position.Item1, position.Item2 - 1));
            }
        }
    }
}
