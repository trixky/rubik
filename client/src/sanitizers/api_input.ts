import { instructions } from '../stores/static/instructions';
import type Result from '../models/result';
import type ResultInstruction from '../models/result_instruction';

export default function (input: string): Result {
	const processed_input = input.trim().toUpperCase();
	const time_token = processed_input.split(' ', 1)[0];
	const time = parseInt(time_token);

	const groups = processed_input
		.slice(time_token.length)
		.trim()
		.split(',')
		.map((group) =>
			group
				.trim()
				.split(' ')
				.filter((instruction) => instruction.length > 0)
		);

	if (groups.length != 4) {
		throw 'need 4 groups';
	}

	if (isNaN(time)) {
		throw 'time is corrupted';
	}

	const result_instructions = <ResultInstruction[]>[];

	groups.forEach((group, index) => {
		group.forEach((instruction) => {
			if (!instructions.includes(instruction)) {
				throw 'invalid instruction: ' + instruction;
			}
			result_instructions.push(<ResultInstruction>{
				instruction,
				group: index + 1
			});
		});
	});

	return <Result>{ instructions: result_instructions, time };
}
