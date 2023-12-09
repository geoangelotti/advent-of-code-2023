import lib


fn main() raises:
    let file = open("input.txt", "\n").read()
    print(lib.process_part1(file))
