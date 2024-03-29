use itertools::{Itertools, Position};
use nom::{
    bytes::complete::tag,
    character::complete::{self, alphanumeric1, newline},
    multi::separated_list1,
    sequence::tuple,
    IResult, Parser,
};
use nom_supreme::ParserExt;
use std::ops::Deref;

#[derive(Debug, Clone, Copy)]
enum HandType {
    FiveOfAKind = 6,
    FourOfAKind = 5,
    FullHouse = 4,
    ThreeOfAKind = 3,
    TwoPair = 2,
    OnePair = 1,
    HighCard = 0,
}

fn get_hand_strength(cards: &str) -> (HandType, (u32, u32, u32, u32, u32)) {
    use HandType::*;
    let counts = cards.chars().counts().values().sorted().join("");
    let hand_type = match counts.deref() {
        "5" => Some(FiveOfAKind),
        "14" => Some(FourOfAKind),
        "23" => Some(FullHouse),
        "113" => Some(ThreeOfAKind),
        "122" => Some(TwoPair),
        "1112" => Some(OnePair),
        "11111" => Some(HighCard),
        _ => None,
    }
    .unwrap();
    let card_scores = cards
        .chars()
        .map(|card| match card {
            'A' => 14,
            'K' => 13,
            'Q' => 12,
            'J' => 11,
            'T' => 10,
            value => value.to_digit(10).unwrap(),
        })
        .collect_tuple()
        .unwrap();
    (hand_type, card_scores)
}

fn get_joker_hand_strength(cards: &str) -> (HandType, (u32, u32, u32, u32, u32)) {
    use HandType::*;
    let counts = cards.chars().counts();
    let values = if let Some(joker_count) = counts.get(&'J') {
        if *joker_count == 5 {
            "5".to_string()
        } else {
            counts
                .iter()
                .filter_map(|(key, value)| (key != &'J').then_some(value))
                .sorted()
                .with_position()
                .map(|(position, value)| match position {
                    Position::Last | Position::Only => value + joker_count,
                    _ => *value,
                })
                .join("")
        }
    } else {
        counts.values().sorted().join("")
    };

    let hand_type = match values.deref() {
        "5" => FiveOfAKind,
        "14" => FourOfAKind,
        "23" => FullHouse,
        "113" => ThreeOfAKind,
        "122" => TwoPair,
        "1112" => OnePair,
        "11111" => HighCard,
        value => panic!("should never happen. Encountered `{}`", value),
    };
    let card_scores = cards
        .chars()
        .map(|card| match card {
            'A' => 14,
            'K' => 13,
            'Q' => 12,
            'J' => 1,
            'T' => 10,
            value => value.to_digit(10).unwrap(),
        })
        .collect_tuple()
        .unwrap();
    (hand_type, card_scores)
}

fn parse_hand(input: &str) -> IResult<&str, (&str, u64)> {
    let (input, (cards, bid)) = tuple((alphanumeric1, complete::u64.preceded_by(tag(" "))))(input)?;
    Ok((input, (cards, bid)))
}

fn parse_hands(input: &str) -> IResult<&str, Vec<(&str, u64)>> {
    let (input, hands) = separated_list1(newline, parse_hand).parse(input)?;
    Ok((input, hands))
}

pub fn process_part_1(input: &str) -> String {
    let (_, hands) = parse_hands(input).unwrap();
    hands
        .iter()
        .map(|(hand, bid)| (hand, bid, get_hand_strength(hand)))
        .sorted_by_key(|(_, _, x)| (x.0 as u8, x.1))
        .enumerate()
        .map(|(i, (_, bid, _))| (i as u64 + 1) * bid)
        .sum::<u64>()
        .to_string()
}

pub fn process_part_2(input: &str) -> String {
    let (_, hands) = parse_hands(input).unwrap();
    hands
        .iter()
        .map(|(hand, bid)| (hand, bid, get_joker_hand_strength(hand)))
        .sorted_by_key(|(_, _, x)| (x.0 as u8, x.1))
        .enumerate()
        .map(|(i, (_, bid, _))| (i as u64 + 1) * bid)
        .sum::<u64>()
        .to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483";

    #[test]
    fn part_1_works() {
        let result = process_part_1(INPUT);
        assert_eq!(result, "6440");
    }

    #[test]
    fn part_2_works() {
        let result = process_part_2(INPUT);
        assert_eq!(result, "5905");
    }
}
