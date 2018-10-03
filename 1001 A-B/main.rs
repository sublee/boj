// https://www.acmicpc.net/problem/1001
use std::io::{self, Read};

// 두 정수 A와 B를 입력받은 다음,
// A-B를 출력하는 프로그램을 작성하시오.
fn main() {
    // 입력:
    // 첫째 줄에 A와 B가 주어진다. (0 < A, B < 10)
    let mut buffer = String::new();
    io::stdin().read_to_string(&mut buffer);

    let mut iter = buffer.split_whitespace();
    let a = iter.next().unwrap();
    let b = iter.next().unwrap();

    // 출력:
    // 첫째 줄에 A-B를 출력한다.
    let a32 = a.parse::<i32>().unwrap();
    let b32 = b.parse::<i32>().unwrap();

    println!("{}", a32 - b32);
}
