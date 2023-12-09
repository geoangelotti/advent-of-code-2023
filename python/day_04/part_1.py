from lib import process_part1

def main():
    with open("input.txt", "r") as file:
        print(process_part1(file.read()))


if __name__ == "__main__":
    main()