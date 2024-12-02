// Description: https://adventofcode.com/2024/day/2

fn main() {
    let input = include_str!("../../inputs/day02/day02.txt");
    let result = solve(input);
    println!("{}", result);
}

fn solve(input: &str) -> String {
    let lines = input.lines().map(str::to_string);

    let mut total_safe = 0;
    'line_loop: for line in lines {
        let split: Vec<String> = line.split(" ").map(str::to_string).collect();
        for i in 1..split.len() - 1 {
            let prev: i32 = split[i - 1].parse().unwrap();
            let curr: i32 = split[i].parse().unwrap();
            let next: i32 = split[i + 1].parse().unwrap();

            let diff_prev = prev - curr;
            let diff_next = curr - next;

            if diff_prev * diff_next < 0
                || diff_prev == 0
                || diff_prev.abs() > 3
                || diff_next == 0
                || diff_next.abs() > 3
            {
                continue 'line_loop;
            }
        }

        total_safe += 1;
    }

    total_safe.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_result() {
        let input = include_str!("../../inputs/day02/day02.txt");
        let expected = "549";
        let result = solve(input);
        assert_eq!(result, expected);
    }
}
