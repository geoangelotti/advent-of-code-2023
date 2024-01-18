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

  let character_counts s =
    let counts = Hashtbl.create 5 in
    String.iter
      (fun c ->
        let current_count = try Hashtbl.find counts c with Not_found -> 0 in
        Hashtbl.replace counts c (current_count + 1))
      s;
    counts

  let list_to_tuple list =
    let opt =
      match list with [ a; b; c; d; e ] -> Some (a, b, c, d, e) | _ -> None
    in
    Option.get opt

  let hand_type hand =
    let hand =
      match hand with
      | "5" -> Some FiveOfAKind
      | "14" -> Some FourOfAKind
      | "23" -> Some FullHouse
      | "113" -> Some ThreeOfAKind
      | "122" -> Some TwoPair
      | "1112" -> Some OnePair
      | "11111" -> Some HighCard
      | _ -> None
    in
    Option.get hand

  let compare_hands hand1 hand2 =
    let (h1_type, h1_score), _ = hand1 in
    let (h2_type, h2_score), _ = hand2 in
    let h1_strength = get_strength h1_type in
    let h2_strength = get_strength h2_type in
    let compare_scores sc1 sc2 =
      match (sc1, sc2) with
      | [], [] -> 0
      | [], _ -> 0
      | _, [] -> 0
      | hd1 :: tl1, hd2 :: tl2 ->
          if hd1 <> hd2 then hd1 - hd2 else compare tl1 tl2
    in
    if h1_strength <> h2_strength then h1_strength - h2_strength
    else compare_scores h1_score h2_score

  let string_to_char_list s = s |> String.to_seq |> List.of_seq

  let card_to_score card =
    match card with
    | 'A' -> 14
    | 'K' -> 13
    | 'Q' -> 12
    | 'J' -> 11
    | 'T' -> 10
    | _ -> int_of_char card - int_of_char '0'

  let get_hand_strength hand =
    let counts = character_counts hand in
    let counts = List.of_seq (Hashtbl.to_seq_values counts) in
    let sorted = List.sort compare counts in
    let joined = String.concat "" (List.map string_of_int sorted) in
    let hand_type = hand_type joined in
    let chars = string_to_char_list hand in
    let scores = List.map card_to_score chars in
    (hand_type, scores)

  let parse line =
    let parts = Str.split (Str.regexp " ") line in
    (get_hand_strength (List.hd parts), int_of_string (List.hd (List.tl parts)))

  let enumerate_list list =
    let rec enumerate_list_helper n list =
      match list with
      | [] -> []
      | hd :: tl -> (n, hd) :: enumerate_list_helper (n + 1) tl
    in
    enumerate_list_helper 0 list

  let process_part_1 input =
    let lines = input |> Utils.split_into_lines in
    let hands = List.map parse lines in
    let sorted = List.sort compare_hands hands in
    let enumerated = enumerate_list sorted in
    let take_bid (i, (_, bid)) = (i + 1) * bid in
    let bids = List.map take_bid enumerated in
    List.fold_left ( + ) 0 bids |> string_of_int
end
