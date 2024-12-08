use regex::Regex;

fn main() {
    let input = include_str!("../../inputs/day03/day03.txt");
    let result = solve(input);
    println!("{result}");
}

fn solve(input: &str) -> String {
    let regx = Regex::new(r"mul\(([0-9]+),([0-9]+)\)").unwrap();

    let mut sum = 0;
    regx.captures_iter(input).for_each(|m| {
        let (_, [a, b]) = m.extract();
        let first = a.parse::<i32>().unwrap();
        let second = b.parse::<i32>().unwrap();

        sum += first * second;
    });

    sum.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solve() {
        let input = include_str!("../../inputs/day03/day03.txt");
        let expected = "156388521";
        let result = solve(input);
        assert_eq!(result, expected);
    }
}
