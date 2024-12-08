use regex::Regex;

fn main() {
    let input = include_str!("../../inputs/day03/day03.txt");
    let result = solve(input);
    println!("{result}");
}

fn solve(input: &str) -> String {
    let instruction_regex = Regex::new(r"mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)").unwrap();
    let mul_regex = Regex::new(r"mul\(([0-9]+),([0-9]+)\)").unwrap();

    let mut sum = 0;
    let mut is_do = true;
    instruction_regex.captures_iter(input).for_each(|cap| {
        let (instruction, _): (&str, [&str; 0]) = cap.extract();
        match instruction {
            "do()" => is_do = true,
            "don't()" => is_do = false,
            _ => {
                if !is_do {
                    return;
                }

                let (_, [l, r]) = mul_regex
                    .captures_iter(instruction)
                    .peekable()
                    .peek()
                    .unwrap()
                    .extract();

                let left = l.parse::<i32>().unwrap();
                let right = r.parse::<i32>().unwrap();
                sum += left * right;
            }
        }
    });

    sum.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solve() {
        let input = include_str!("../../inputs/day03/day03.txt");
        let expected = "75920122";
        let result = solve(input);
        assert_eq!(result, expected);
    }
}
