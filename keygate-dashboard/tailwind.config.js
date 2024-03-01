/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      transitionProperty: {
        multiple:
          "width , height , backgroundColor , border-radius, border, color, border-color",
      },
    },
  },
  plugins: [],
};
