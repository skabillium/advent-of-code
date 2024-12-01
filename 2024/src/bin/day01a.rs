// Description: https://adventofcode.com/2024/day/1

fn main() {
    let input = include_str!("../../inputs/day01/day01.txt");
    let result = solve(input);
    println!("Total Difference: {}", result);
}

fn solve(input: &str) -> String {
    let lines: Vec<String> = input.lines().map(str::to_string).collect();

    let mut col1: Vec<i32> = Vec::with_capacity(lines.len() as usize);
    let mut col2: Vec<i32> = Vec::with_capacity(lines.len() as usize);

    for line in lines.iter() {
        let split: Vec<String> = line.split_whitespace().map(str::to_string).collect();
        let first: i32 = split.first().unwrap().parse().unwrap();
        let second: i32 = split.last().unwrap().parse().unwrap();

        col1.push(first);
        col2.push(second);
    }

    col1.sort();
    col2.sort();

    let mut total_difference = 0;
    for i in 0..col1.len() {
        total_difference += (col1[i] - col2[i]).abs();
    }

    total_difference.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_result() {
        let input = include_str!("../../inputs/day01/day01.txt");
        let expected = "1590491";
        let result = solve(input);
        assert_eq!(result, expected);
    }
}
