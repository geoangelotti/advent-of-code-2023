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

  let hand_type_to_int = function
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

  let string_to_hand_type hand =
    let hand =
      match hand with
      | "5" -> Some FiveOfAKind
      | "14" -> Some FourOfAKind
      | "23" -> Some FullHouse
      | "113" -> Some ThreeOfAKind
      | "122" -> Some TwoPair
      | "1112" -> Some OnePair
      | "11111" -> Some HighCard
      | "" -> Some FiveOfAKind
      | _ -> None
    in
    Option.get hand

  let compare_hands hand1 hand2 =
    let (h1_type, h1_score), _ = hand1 in
    let (h2_type, h2_score), _ = hand2 in
    let h1_strength = hand_type_to_int h1_type in
    let h2_strength = hand_type_to_int h2_type in
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

  let card_to_score card joker_score =
    match card with
    | 'A' -> 14
    | 'K' -> 13
    | 'Q' -> 12
    | 'J' -> joker_score
    | 'T' -> 10
    | _ -> int_of_char card - int_of_char '0'

  let get_hand_strength hand =
    let chars = string_to_char_list hand in
    let scores = List.map (fun card -> card_to_score card 11) chars in
    let counts = character_counts hand in
    let counts = List.of_seq (Hashtbl.to_seq_values counts) in
    let sorted = List.sort compare counts in
    let joined = String.concat "" (List.map string_of_int sorted) in
    let hand_type = string_to_hand_type joined in
    (hand_type, scores)

  let get_joker_hand_strength hand =
    let counts = character_counts hand in
    let jokers = Hashtbl.find_opt counts 'J' in
    match jokers with
    | None -> get_hand_strength hand
    | Some jokers ->
        let chars = string_to_char_list hand in
        let scores = List.map (fun card -> card_to_score card 0) chars in
        let filtered = Str.global_replace (Str.regexp "J") "" hand in
        let counts = character_counts filtered in
        let counts = List.of_seq (Hashtbl.to_seq counts) in
        let sorted = List.sort (fun (_, ai) (_, bi) -> ai - bi) counts in
        let rec add_joker list joker =
          match list with
          | [] -> []
          | [ (c, i) ] -> (c, i + joker) :: add_joker [] joker
          | hd :: tl -> hd :: add_joker tl joker
        in
        let jokered = add_joker sorted jokers in
        let joined =
          String.concat "" (List.map (fun (_, i) -> string_of_int i) jokered)
        in
        let hand_type = string_to_hand_type joined in
        (hand_type, scores)

  let parse line get_hand_strength =
    let parts = Str.split (Str.regexp " ") line in
    (get_hand_strength (List.hd parts), int_of_string (List.hd (List.tl parts)))

  let enumerate_list list =
    let rec enumerate_list_helper acc n list =
      match list with
      | [] -> acc
      | hd :: tl -> enumerate_list_helper (acc @ [ (n, hd) ]) (n + 1) tl
    in
    enumerate_list_helper [] 0 list

  let process input get_hand_strength =
    let lines = Str.split (Str.regexp "\n") input in
    let hands = List.map (fun line -> parse line get_hand_strength) lines in
    let sorted = List.sort compare_hands hands in
    let enumerated = enumerate_list sorted in
    let take_bid (i, (_, bid)) = (i + 1) * bid in
    let bids = List.map take_bid enumerated in
    List.fold_left ( + ) 0 bids |> string_of_int

  let process_part_1 input = process input get_hand_strength
  let process_part_2 input = process input get_joker_hand_strength
end
