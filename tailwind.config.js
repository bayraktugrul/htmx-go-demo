/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        'internal/templates/*.templ',
    ],
    theme: {},
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
    ]
}