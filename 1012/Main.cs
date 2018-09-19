using System;

namespace Acmicpc
{
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
                var m = int.Parse(parts[0]);
                var n = int.Parse(parts[1]);
                var k = int.Parse(parts[2]);

                // 그 다음 K줄에는 배추의 위치 X(0 ≤ X ≤ M-1), Y(0 ≤ Y ≤ N-1)가 주어진다.
                int[,] a = new int[m, n];
                for (int j = 0; j < k; j++)
                {
                    parts = Console.ReadLine().Split(new Char[] { ' ' }, 2);
                    var x = int.Parse(parts[0]);
                    var y = int.Parse(parts[1]);
                    a[x, y] = 1;
                }

                Console.WriteLine(Solve(a, m, n));
            }
        }

        static int Solve(int[,] a, int m, int n)
        {
            int[,] b = new int[m, n];
            bool[,] v = new bool[m, n];

            int count = 0;

            for (int i = 0; i < m; i++)
            {
                for (int j = 0; j < n; j++)
                {
                    if (v[i, j])
                    {
                        continue;
                    }

                    if (a[i, j] == 1)
                    {
                        Visit(a, m, n, i, j, v);
                        count++;
                    }
                }
            }

            return count;
        }

        static void Visit(int[,] a, int m, int n, int i, int j, bool[,] v)
        {
            if (v[i, j])
            {
                return;
            }

            v[i, j] = true;

            if (a[i, j] != 1) {
                return;
            }

            if (i > 0)
            {
                Visit(a, m, n, i-1, j, v);
            }
            if (i < m-1)
            {
                Visit(a, m, n, i+1, j, v);
            }
            if (j > 0)
            {
                Visit(a, m, n, i, j-1, v);
            }
            if (j < n-1)
            {
                Visit(a, m, n, i, j+1, v);
            }
        }
    }
}
