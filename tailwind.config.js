module.exports = {
  purge: [
    './**/*.html',
    './**/*.js',
    './**/*.go',
    './**/*.templ',
  ],
  theme: {
    extend: {
      colors: {
        primary: '#21263c',
        primarydarker: '#181c2e', //darker
        secondary: '#21263c',
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
