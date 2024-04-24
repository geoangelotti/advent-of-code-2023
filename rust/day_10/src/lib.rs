pub fn process_part_1(input: &str) -> String {
    input.into()
}

#[cfg(test)]
mod tests {
    use rstest::rstest;

    use super::*;

    #[rstest]
    #[case(
        ".....
.S-7.
.|.|.
.L-J.
.....",
        "4"
    )]
    #[case(
        "..F7.
.FJ|.
SJ.L7
|F--J
LJ...",
        "4"
    )]
    #[test]
    fn test_process(#[case] input: &str, #[case] output: &str) {
        let result = process_part_1(input);
        assert_eq!(result, output);
    }
}
