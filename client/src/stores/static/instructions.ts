export const moves = ['F', 'R', 'U', 'B', 'L', 'D'];
export const colors = ['yellow', '', 'red', 'blue', 'green', 'orange'];
export const modifiers = ['', "'", '2'];
export const instructions = moves.flatMap((move) => modifiers.map((modifier) => move + modifier));
export const physical_instructions = instructions.filter(
	(instruction) => !instruction.endsWith('2')
);
