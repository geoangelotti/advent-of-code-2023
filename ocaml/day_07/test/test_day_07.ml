open OUnit2
open Day_07.Lib.Day_07

let input = {|32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483|}

let process_part_1_test _ =
  let result = process_part_1 input in
  assert_equal ~printer:(fun x -> x) "6440" result

let process_part_2_test _ =
  let result = process_part_2 input in
  assert_equal ~printer:(fun x -> x) "5905" result

let suite =
  "Test Suite"
  >::: [
         "process_part_1" >:: process_part_1_test;
         "process_part_2" >:: process_part_2_test;
       ]

let () = run_test_tt_main suite
