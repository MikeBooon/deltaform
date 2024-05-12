/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				'base-content': '#000000',
				base: '#FFFFFF',
				primary: '#540D6E',
				secondary: '#EE4266',
				accent: '#FFD23F'
			},
			fontFamily: {
				header: 'Arvo, serif',
				body: 'Barlow, sans-serif'
			}
		}
	},
	plugins: []
};
