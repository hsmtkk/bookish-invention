use std::convert::TryInto;

pub fn collatz_steps(n: i64) -> i64 {
    return collatz_numbers(n).len().try_into().unwrap();
}

#[test]
fn test_collatz_steps(){
    let want = 9;
    let got = collatz_steps(12);
    assert_eq!(want,got);
}

pub fn collatz_numbers(n: i64) -> Vec<i64> {
    let mut m = n;
    let mut result: Vec<i64> = Vec::new();
    loop {
        if m % 2 == 0 {
            m /= 2;
        } else {
            m = 3 * m + 1;
        }
        result.push(m);
        if m == 1 {
            break;
        }
    }
    return result;
}

#[test]
fn test_collatz_numbers(){
    let want = vec![6, 3, 10, 5, 16, 8, 4, 2, 1];
    let got = collatz_numbers(12);
    assert_eq!(want,got);
}