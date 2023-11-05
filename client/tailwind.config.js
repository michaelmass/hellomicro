const plugin = require("tailwindcss/plugin");

module.exports = {
  mode: "jit",
  darkMode: "class",
  purge: {
    mode: "all",
    preserveHtmlElements: false,
    safelist: [],
    content: ["./src/**/*.{js,jsx,ts,tsx}", "./**/*.html"],
  },
  theme: {
    fontFamily: {
      sans: ["Poppins", "Helvetica", "sans-serif"],
      display: ["Poppins", "Helvetica", "sans-serif"],
      body: ["Poppins", "Helvetica", "sans-serif"],
      html: ["Poppins", "Helvetica", "sans-serif"],
    },
    boxShadow: {
      md:
        "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)",
      lg:
        "0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)",
      subheader: "0px 10px 30px 0px rgb(82 63 105 / 8%)",
      aside: "rgba(0, 0, 0, 0.1) 0px 4px 12px",
      DEFAULT: "0px 0px 30px 0px rgb(82 63 105 / 5%)",
    },
    svg: (theme) => ({
      ...theme("colors"),
    }),
    extend: {
      colors: {
        primary: {
          light: "#9CA5F8",
          DEFAULT: "#424B9E",
          dark: "#242D80",
        },
        secondary: {
          DEFAULT: "#232A34",
        },
      },
    },
  },
  variants: {
    extend: {
      backgroundColor: ["even"],
    },
  },
  plugins: [
    plugin(({ addUtilities, theme }) => {
      const colors = theme("svg", {});

      const utilities = Object.keys(colors).map((color) =>
        typeof colors[color] === "string"
          ? {
            [`.svg-${color} svg g [fill]`]: {
              fill: colors[color],
            },
          }
          : Object.keys(colors[color])
            .map((extension) =>
              extension === "DEFAULT"
                ? {
                  [`.svg-${color} svg g [fill]`]: {
                    fill: colors[color][extension],
                  },
                }
                : {
                  [`.svg-${color}-${extension} svg g [fill]`]: {
                    fill: colors[color][extension],
                  },
                }
            )
            .filter((value) => value)
      );

      addUtilities(utilities, ["hover", "group-hover"]);
    }),
  ],
};
