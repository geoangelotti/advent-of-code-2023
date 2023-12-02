pub fn process_part_1(input: &str) -> String {
    input
        .lines()
        .into_iter()
        .map(|line| {
            let values = line
                .chars()
                .filter(|char| char.is_ascii_digit())
                .map(|char| char as u32 - '0' as u32)
                .collect::<Vec<u32>>();
            match values.len() {
                0 => 0,
                _ => values.first().unwrap() * 10 + values.last().unwrap(),
            }
        })
        .sum::<u32>()
        .to_string()
}

pub fn process_part_2(input: &str) -> String {
    process_part_1(input)
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT1: &str = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";

    const INPUT2: &str = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";

    #[test]
    fn part_1_works() {
        let result = process_part_1(INPUT1);
        assert_eq!(result, 142.to_string());
    }

    #[test]
    fn part_2_works() {
        let result = process_part_2(INPUT2);
        assert_eq!(result, 281.to_string());
    }
}
