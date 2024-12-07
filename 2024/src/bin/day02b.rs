// Description: https://adventofcode.com/2024/day/2

fn main() {
    let input = include_str!("../../inputs/day02/day02.txt");
    let result = solve(input);
    println!("{}", result);
}

fn solve(input: &str) -> String {
    let lines = input.lines().map(str::to_string);

    let mut total_safe = 0;
    'lines_loop: for line in lines {
        let report: Vec<i32> = line.split(" ").map(|s| s.parse::<i32>().unwrap()).collect();

        if is_report_safe(&report) {
            total_safe += 1;
            continue;
        }

        for i in 0..report.len() {
            let mut permutation = report.clone();
            permutation.remove(i);

            if is_report_safe(&permutation) {
                total_safe += 1;
                continue 'lines_loop;
            }
        }
    }

    total_safe.to_string()
}

fn is_report_safe(report: &Vec<i32>) -> bool {
    for i in 1..report.len() - 1 {
        let prev: i32 = report[i - 1];
        let curr: i32 = report[i];
        let next: i32 = report[i + 1];

        let diff_prev = prev - curr;
        let diff_next = curr - next;

        if diff_prev * diff_next < 0
            || diff_prev == 0
            || diff_prev.abs() > 3
            || diff_next == 0
            || diff_next.abs() > 3
        {
            return false;
        }
    }

    return true;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_result() {
        let input = include_str!("../../inputs/day02/day02.txt");
        let expected = "589";
        let result = solve(input);
        assert_eq!(result, expected);
    }
}
