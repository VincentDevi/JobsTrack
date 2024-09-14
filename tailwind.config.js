/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["internal/view/**/*.templ"],
	theme: {
		extend: {
			colors: {
				'black': '#000000',
				'white': '#FFFFFF',
				'beige': '#DAD7CD',
				'vlight-green': '#A3B18A',
				'light-green': '#588157',
				'green': '#3A5A40',
				'dark-green': '#344E41'
			},
			borderColor: {
				'black': '#000000',
				'white': '#FFFFFF',
				'beige': '#DAD7CD',
				'vlight-green': '#A3B18A',
				'light-green': '#588157',
				'green': '#3A5A40',
				'dark-green': '#344E41'
			},
		}

	},
	plugins: [],
}
