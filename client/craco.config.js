const CracoSwcPlugin = require('craco-swc')
const tailwindcss = require('tailwindcss')
const autoprefixer = require('autoprefixer')

module.exports = {
  style: {
    postcss: {
      plugins: [tailwindcss, autoprefixer],
    },
  },
  plugins: [{ plugin: CracoSwcPlugin }],
}
