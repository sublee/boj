#include <iostream>
#include <string>

void solve(const int &n) {
  int zero_q[3] = {0};
  int one_q[3]  = {0};

  for (int i = 0; i <= n; i++) {
    // shift
    zero_q[2] = zero_q[1];
    zero_q[1] = zero_q[0];
    zero_q[0] = 0;

    one_q[2] = one_q[1];
    one_q[1] = one_q[0];
    one_q[0] = 0;

    // cumulate
    switch (i) {
      case 0:
        zero_q[0] = 1;
        break;
      case 1:
        one_q[0] = 1;
        break;
      default:
        zero_q[0] = zero_q[1] + zero_q[2];
        one_q[0]  = one_q[1]  + one_q[2];
        break;
    }
  }

  // 각 테스트 케이스마다 0이 출력되는 횟수와 1이 출력되는 횟수를
  // 공백으로 구분해서 출력한다.
  int zero = zero_q[0];
  int one  = one_q[0];
  std::cout << zero << ' ' << one << std::endl;
}

int main() {
  std::string line;

  // 첫째 줄에 테스트 케이스의 개수 T가 주어진다.
  std::getline(std::cin, line);
  int t = std::stoi(line);

  // 각 테스트 케이스는 한 줄로 이루어져 있고, N이 주어진다.
  // N은 40보다 작거나 같은 자연수 또는 0이다.
  for (int i = 0; i < t; i++) {
    std::getline(std::cin, line);
    int n = std::stoi(line);

    solve(n);
  }

  return 0;
}
