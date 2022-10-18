import { writable } from 'svelte/store';
import SanitizeApiInput from '../sanitizers/api_input';
import type Result from '../models/result';

function createResult() {
	const { subscribe, set } = writable(<Result>{ instructions: [], time: 0 });

	return {
		subscribe,
		set,
		setFromString: (input: string): boolean => {
			try {
				set(SanitizeApiInput(input));
			} catch (err) {
				console.log(err);
				return false;
			}

			return true;
		},
		reset: () => set(<Result>{ instructions: [], time: 0 })
	};
}

const result = createResult();

export default result;
