import type ResultInstruction from './result_instruction';

export default interface Result {
	instructions: ResultInstruction[];
	time: number;
}
