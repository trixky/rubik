// generate a random number between a minimum and a maximum
export function getRangedRandomNumber(min: number, max: number): number {
	return Math.floor(min + Math.random() * max);
}
