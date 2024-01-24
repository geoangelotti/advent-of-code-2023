const zip = (a: RegExpMatchArray[], b:RegExpMatchArray[]) => a.map((k, i) => [k, b[i]]);

export function processPart1(input: string): string {
	const numbers = new RegExp(/[0-9]+/g);
	const lines = input.split("\n");
	const times = Array.from(lines[0].matchAll(numbers)); 
	const distances = Array.from(lines[1].matchAll(numbers)); 
	const zipped = zip(times, distances);
	return zipped.map(item => {
		const [[time], [distance]] = item.map(i => i.map(j => Number(j)));
		let minimumTimeHeld = Array.from(Array(time+1).keys()).find(timeHeld => timeHeld * (time - timeHeld) > distance);
		return time - 2 * Number(minimumTimeHeld) + 1
	}).reduce((acc, c) => acc * c, 1).toString();
};

export function processPart2(input: string): string {
	const numbers = new RegExp(/[0-9]+/g);
	const lines = input.split("\n");
	const time = Number(Array.from(lines[0].matchAll(numbers)).join("")); 
	const distance = Number(Array.from(lines[1].matchAll(numbers)).join("")); 
	const minimumTimeHeld = Array.from(Array(time+1).keys()).find(timeHeld => timeHeld * (time - timeHeld) > distance);
	return (time - 2 * Number(minimumTimeHeld) + 1).toString();
};