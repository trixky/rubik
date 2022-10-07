/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			gridTemplateColumns: {
				7: 'repeat(7, minmax(0, 1fr))',
				8: 'repeat(8, minmax(0, 1fr))'
			},
			gridTemplateRows: {
				7: 'repeat(7, minmax(0, 1fr))',
				8: 'repeat(8, minmax(0, 1fr))'
			}
		}
	},
	plugins: []
};
