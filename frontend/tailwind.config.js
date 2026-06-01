/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans:['Inter', 'Arial', 'Helvetica', 'sans-serif']
      },
    },
  },
  plugins: [],
}
