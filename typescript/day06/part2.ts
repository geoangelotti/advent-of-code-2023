import { processPart2 } from './index';

const input = await Bun.file("input.txt").text();
console.log(processPart2(input));