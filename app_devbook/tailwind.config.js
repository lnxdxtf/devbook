/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,vue}",
    "./node_modules/flowbite/**/*.js"
  ],
  darkMode: "media",
  theme: {
    extend: {
      colors: {
        devbook: {
          light: {
            shades: "#F3F3F2", // Use this color as the background for your dark-on-light designs, or the text color of an inverted design.
            accent: "#B87846", // Accent colors can be used to bring attention to design elements by contrasting with the rest of the palette.
          },
          dark: {
            shades: "#2A2A31", // Use as the text color for dark-on-light designs, or as the background for inverted designs. 
            accent: "#98553E", // Another accent color to consider. Not all colors have to be used - sometimes a simple color scheme works best.
          },
          // main: "#767472", // This color should be eye-catching but not harsh. It can be liberally applied to your layout as its main identity
          main: {
            one: "#6a4c93",
            two: "#ff595e",
          },

          input: {
            light: "#D4D4D8",
            dark: "#52525B",
          }
        }
      }
    },
  },
  plugins: [
    require('flowbite/plugin')
  ],
}