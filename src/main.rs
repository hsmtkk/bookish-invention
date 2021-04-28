mod collatz;

fn main() {
    for i in 10..100 {
        let ns = collatz::collatz_numbers(i);
        println!("{:?}", ns);
    }
}
