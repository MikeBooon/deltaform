/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				'base-content': '#FFFFFF',
				base: '#282943',
				primary: '#F8B85D',
				secondary: '#E46B3F',
				accent: '#C6412A'
			},
			fontFamily: {
				header: 'Arvo, serif',
				body: 'Barlow, sans-serif'
			}
		}
	},
	plugins: []
};
