import re
import math
from typing import List


def return_points(line: str) -> int:
    parts = line.split(": ")[1].split(" | ")
    winning_numbers = set([x for x in re.split(r" +", parts[0]) if len(x) > 0])
    played_numbers = set([x for x in re.split(r" +", parts[1]) if len(x) > 0])
    return int(math.pow(2, played_numbers.intersection(winning_numbers).__len__() -1))

def calculate_points(input: str) -> List[int]:
    return [return_points(line) for line in input.split("\n")]

def process_part1(input: str)-> str:
    return str(sum(calculate_points(input)))

def process_part2(input: str) ->str:
    points = calculate_points(input)
    list_of_points = [[point] for point in points ]
    for i in range(points.__len__()):
        for point in list_of_points[i]:
            cards = int (math.log2(point)+1) if point else 0
            for j in range(i+1, i+1+cards):
                if j < points.__len__():
                    list_of_points[j].append(points[j])        
    return str(sum([points.__len__() for points in list_of_points]))

