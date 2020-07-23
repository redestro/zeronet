module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: ["plugin:vue/essential", "@vue/airbnb"],
  parserOptions: {
    parser: "babel-eslint"
  },
  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "warn",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "warn",
    semi: ["error", "always"],
    "no-trailing-spaces": "error",
    "object-curly-spacing": [1, "always"],
    "no-dupe-class-members": 1,
    "no-iterator": 1,
    "comma-dangle": ["error", "never"],
    "one-var": [1, "never"],
    "no-nested-ternary": 0,
    "no-restricted-syntax": [1, "WithStatement"],
    "no-duplicate-imports": 1,
    "no-unneeded-ternary": 1,
    "no-case-declarations": 1,
    "dot-notation": 1,
    "no-useless-constructor": 1,
    eqeqeq: 1,
    "object-shorthand": 1,
    "prefer-rest-params": 1,
    "no-array-constructor": 1,
    "array-callback-return": 0,
    "no-param-reassign": 1,
    "prefer-arrow-callback": 1,
    "no-useless-concat": 1,
    "indent": 0,
    "prefer-template": 1,
    "comma-spacing": 0,
    "quote-props": 0,
    "key-spacing": 0,
    "arrow-parens": 0,
    "arrow-body-style": 0,
    "quotes": 0,
    "no-extra-bind": 0,
    "no-var": 1,
    "eol-last": 0,
    "no-unused-vars": 0,
    "prefer-destructuring": 0,
    "no-else-return": 0,
    "vars-on-top": 0,
    "no-underscore-dangle": 0,
    "no-alert": 0,
    "no-lone-blocks": 0,
    "no-plusplus": 0,
    "max-len": ["error", { code: 120 }]
  }
};
