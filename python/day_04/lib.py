import re
import math


def return_points(line: str) -> int:
    parts = line.split(": ")[1].split(" | ")
    winning_numbers = set([x for x in re.split(r" +", parts[0]) if len(x) > 0])
    played_numbers = set([x for x in re.split(r" +", parts[1]) if len(x) > 0])
    return int(math.pow(2, played_numbers.intersection(winning_numbers).__len__() -1))

def process_part1(input: str)-> str:
    return str(sum([return_points(line) for line in input.split("\n")]))
