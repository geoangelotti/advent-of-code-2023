from python import Python


fn process_part1(input: String) raises -> String:
    let lines = input.split("\n")
    for i in range(lines.__len__()):
        let line = lines[i]
        let parts = line.split(": ")
        let card = parts[1].split(" | ")
        let winningNumbers = card[0].split(" ")
    return ""
