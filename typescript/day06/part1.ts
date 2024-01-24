import { processPart1 } from './index';

const input = await Bun.file("input.txt").text();
console.log(processPart1(input));