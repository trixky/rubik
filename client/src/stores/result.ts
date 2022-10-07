import { writable } from 'svelte/store';
import SanitizeInput from '../sanitizers/input';

function createResult() {
	const { subscribe, set } = writable(<string[]>[]);

	return {
		subscribe,
		set,
		setFromString: (input: string): boolean => {
			try {
				set(SanitizeInput(input));
			} catch {
				return false;
			}

			return true;
		},
		reset: () => set(<string[]>[])
	};
}

const result = createResult();

export default result;
