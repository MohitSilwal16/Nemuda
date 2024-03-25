/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './views/**/*.{html,templ}',
  ],
  theme: {
    extend: {
      backgroundImage:{
        'register-login':"url('/static/images/background.jpg')",
      }
    },
  },
  plugins: [],
}

