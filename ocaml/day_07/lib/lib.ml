module Utils = struct
  let read_file filename =
    let channel = open_in filename in
    let rec read_lines accumulator =
      try
        let line = input_line channel in
        read_lines (accumulator ^ line ^ "\n")
      with End_of_file ->
        close_in channel;
        accumulator
    in
    read_lines ""

  let split_into_lines input_string = Str.split (Str.regexp "\n") input_string
end

module Day_07 = struct
  type hand_type =
    | FiveOfAKind
    | FourOfAKind
    | FullHouse
    | ThreeOfAKind
    | TwoPair
    | OnePair
    | HighCard

  let get_strength = function
    | FiveOfAKind -> 6
    | FourOfAKind -> 5
    | FullHouse -> 4
    | ThreeOfAKind -> 3
    | TwoPair -> 2
    | OnePair -> 1
    | HighCard -> 0

  open Hashtbl

  let character_frequencies s =
    let frequencies = Hashtbl.create 10 in
    String.iter
      (fun c ->
        let current_count =
          try Hashtbl.find frequencies c with Not_found -> 0
        in
        Hashtbl.replace frequencies c (current_count + 1))
      s;
    frequencies

  let get_hand_strength hand =
    let frequencies = character_frequencies hand in
    let frequencies = of_seq (to_seq_values frequencies) in
    let sorted = List.sort compare frequencies in
    let joined = String.concat "" (List.map string_of_int sorted) in
    let () = print_endline (joined ^ hand) in
    match joined with
    | "5" -> Some FiveOfAKind
    | "14" -> Some FourOfAKind
    | "23" -> Some FullHouse
    | "113" -> Some ThreeOfAKind
    | "122" -> Some TwoPair
    | "1112" -> Some OnePair
    | "11111" -> Some HighCard
    | _ -> None

  let parse line =
    let parts = Str.split (Str.regexp " ") line in
    ( Option.get (get_hand_strength (List.hd parts)),
      int_of_string (List.hd (List.tl parts)) )

  let process_part_1 input =
    let lines = input |> Utils.split_into_lines in
    let hands = List.map parse lines in
    let nums = List.map (fun (_, num) -> num) hands in
    let () = print_endline (String.concat "" (List.map string_of_int nums)) in
    List.fold_left ( + ) 0 nums |> string_of_int
end
