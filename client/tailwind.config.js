/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './views/**/*.html',
  ],
  theme: {
    extend: {
      backgroundImage: {
        'register-login': "url('/static/images/background.jpg')",
        'gradient-image': "url('/static/images/background-home.jpg')",
      },
      height: {
        '6/7': '85.7143%', // 6/7 is approximately 85.7143%
        '9/10': '90%',
      },
      bottom: {
        '26': '6.5rem',
      },
    },
  },
  plugins: [],
}

