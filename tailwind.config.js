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
        secondary: '#22b378',
        secondarygray: '#fbfbfb',
        secondarygraydarker: '#eee',
        secondaryblue: '#339DFF',
      },
    },
  },
  variants: {
    extend: {
      borderColor: ['focus'],
    },
  },
  plugins: [],
};
