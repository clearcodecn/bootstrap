/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./web/templates/**/*.gohtml",
    "./web/static/*.js",
    "./web/static/*.css",
    "./node_modules/flowbite/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("flowbite/plugin"), require("flowbite-typography")],
  safelist: [
    "w-64",
    "w-*",
    "w-8/12",
    "rounded-l-lg",
    "rounded-r-lg",
    "bg-gray-200",
    "grid-cols-4",
    "grid-cols-7",
    "h-6",
    "leading-6",
    "h-9",
    "leading-9",
    "shadow-lg",
  ],
};

// <script src="../path/to/flowbite/dist/flowbite.min.js"></script>
