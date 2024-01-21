use itertools::{Itertools, Position};

fn process_history(mut nums: Vec<i64>) -> i64 {
    let mut end_numbers: Vec<i64> = vec![];
    loop {
        if nums.iter().all(|n| n == &0) {
            break;
        }
        nums = nums
            .iter()
            .tuple_windows::<(&i64, &i64)>()
            .with_position()
            .map(|(position, (left, right))| {
                match position {
                    Position::Last | Position::Only => {
                        end_numbers.push(*right);
                    }
                    _ => {}
                };
                right - left
            })
            .collect::<Vec<i64>>();
    }
    end_numbers.iter().sum()
}

pub fn process_part_1(input: &str) -> String {
    input
        .lines()
        .map(|line| {
            let nums = line
                .split_whitespace()
                .map(|num| num.parse::<i64>().unwrap())
                .collect::<Vec<i64>>();
            process_history(nums)
        })
        .sum::<i64>()
        .to_string()
}

pub fn process_part_2(input: &str) -> String {
    input
        .lines()
        .map(|line| {
            let nums = line
                .split_whitespace()
                .map(|num| num.parse::<i64>().unwrap())
                .rev()
                .collect::<Vec<i64>>();
            process_history(nums)
        })
        .sum::<i64>()
        .to_string()
}

#[cfg(test)]
mod tests {
    const INPUT: &str = "0 3 6 9 12 15
    1 3 6 10 15 21
    10 13 16 21 30 45";

    use super::*;

    #[test]
    fn part_1_works() {
        let result = process_part_1(INPUT);
        assert_eq!(result, "114");
    }

    #[test]
    fn part_2_works() {
        let result = process_part_2(INPUT);
        assert_eq!(result, "2");
    }
}
