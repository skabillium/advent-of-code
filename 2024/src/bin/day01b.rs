// Description: https://adventofcode.com/2024/day/1#part2
use std::collections::HashMap;

fn main() {
    let input = include_str!("../../inputs/day01/day01.txt");
    let result = solve(input);
    println!("Similarity Score: {}", result);
}

fn solve(input: &str) -> String {
    let lines: Vec<String> = input.lines().map(str::to_string).collect();

    let mut col1: Vec<i32> = Vec::with_capacity(lines.len() as usize);

    let mut frequencies: HashMap<String, i32> = HashMap::new();

    for line in lines.iter() {
        let split: Vec<String> = line.split_whitespace().map(str::to_string).collect();
        let first: i32 = split.first().unwrap().parse().unwrap();
        let second = split.last().unwrap();

        col1.push(first);

        let frequency = frequencies.get(second);
        match frequency {
            Some(f) => {
                frequencies.insert(second.clone(), f + 1);
            }
            None => {
                frequencies.insert(second.clone(), 1);
            }
        }
    }

    let mut similarity_score = 0;
    for num in col1.iter() {
        match frequencies.get(&num.to_string()) {
            Some(n) => similarity_score += num * n,
            _ => (),
        }
    }

    similarity_score.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_result() {
        let input = include_str!("../../inputs/day01/day01.txt");
        let expected = "22588371";
        let result = solve(input);
        assert_eq!(result, expected);
    }
}
