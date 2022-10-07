import { instructions } from '../stores/static/instructions';

export default function (input: string): string[] {
	const tokens = input.trim().toUpperCase().split(' ');

	if (tokens.every((token) => instructions.includes(token))) {
		return tokens;
	}

	throw 'input not valid';
}
