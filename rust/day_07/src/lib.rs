use itertools::Itertools;
use nom::{
    bytes::complete::tag,
    character::complete::{self, alphanumeric1, newline},
    multi::separated_list1,
    sequence::tuple,
    IResult, Parser,
};
use nom_supreme::ParserExt;
use std::cmp::Ordering::{self, Equal, Greater, Less};
use std::{collections::BTreeMap, iter::zip};

#[derive(Debug, PartialEq, PartialOrd, Eq)]
enum Type {
    FiveOfAKind = 6,
    FourOfAKind = 5,
    FullHouse = 4,
    ThreeOfAKind = 3,
    TwoPair = 2,
    OnePair = 1,
    HighCard = 0,
}

#[derive(Debug, PartialEq, Eq)]
struct Hand<'a> {
    type_strength: Type,
    cards: &'a str,
    bid: u64,
}

fn my_cmp(a: &Hand, b: &Hand) -> Ordering {
    if a.type_strength > b.type_strength {
        return Less;
    }
    if a.type_strength < b.type_strength {
        return Greater;
    }
    let view = zip(a.cards.chars(), b.cards.chars());
    for (a, b) in view {
        if a != b {
            return match (a, b) {
                ('A', _) => Greater,
                (_, 'A') => Less,
                ('K', _) => Greater,
                (_, 'K') => Less,
                ('Q', _) => Greater,
                (_, 'Q') => Less,
                ('J', _) => Greater,
                (_, 'J') => Less,
                ('T', _) => Greater,
                (_, 'T') => Less,
                (a, b) => a.cmp(&b),
            };
        }
    }
    Equal
}

impl<'a> Ord for Hand<'a> {
    fn cmp(&self, other: &Self) -> Ordering {
        my_cmp(self, other)
    }
}

impl<'a> PartialOrd for Hand<'a> {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(my_cmp(self, other))
    }
}

fn get_type_strength(cards: &str) -> Option<Type> {
    let frequency: BTreeMap<char, u32> = cards
        .chars()
        .sorted()
        .group_by(|&x| x)
        .into_iter()
        .map(|(k, v)| (k, v.count() as u32))
        .collect();
    match frequency.len() {
        1 => Some(Type::FiveOfAKind),
        2 => {
            if let Some(_) = frequency.into_iter().find(|(_, v)| *v == 4) {
                return Some(Type::FourOfAKind);
            } else {
                return Some(Type::FullHouse);
            }
        }
        3 => {
            if let Some(_) = frequency.into_iter().find(|(_, v)| *v == 3) {
                return Some(Type::ThreeOfAKind);
            } else {
                return Some(Type::TwoPair);
            }
        }
        4 => Some(Type::OnePair),
        5 => Some(Type::HighCard),
        _ => None,
    }
}

fn parse_hand(input: &str) -> IResult<&str, Hand> {
    let (input, (cards, bid)) = tuple((alphanumeric1, complete::u64.preceded_by(tag(" "))))(input)?;
    Ok((
        input,
        Hand {
            type_strength: get_type_strength(cards).unwrap(),
            cards,
            bid,
        },
    ))
}

fn parse_hands(input: &str) -> IResult<&str, Vec<Hand>> {
    let (input, hands) = separated_list1(newline, parse_hand).parse(input)?;
    Ok((input, hands))
}

pub fn process_part_1(input: &str) -> String {
    let (_, mut hands) = parse_hands(input).unwrap();
    hands.sort();
    hands
        .iter()
        .enumerate()
        .map(|(i, h)| (i + 1) as u64 * h.bid)
        //.collect::<Vec<u64>>()
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
    fn it_works() {
        let result = process_part_1(INPUT);
        assert_eq!(result, "6440");
    }
}
