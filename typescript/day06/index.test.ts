import { expect, test } from "bun:test";
import { processPart1, processPart2 } from "./index";

const INPUT =`Time:      7  15   30
Distance:  9  40  200`;

test("part 1", () => {
  expect(processPart1(INPUT)).toBe("288");
});

test("part 2", () => {
  expect(processPart2(INPUT)).toBe("71503");
});