module.exports = {
  "env": {
      "node": true,
      "commonjs": true,
      "es2021": true,
      "jest": true
  },
  "extends": "eslint:recommended",
  "overrides": [
      {
          "env": {
              "node": true
          },
          "files": [
              ".eslintrc.{js,cjs}"
          ],
          "parserOptions": {
              "sourceType": "script"
          }
      }
  ],
  "parserOptions": {
      "ecmaVersion": "latest"
  },
  "rules": {
      "indent": ["error", 2],
      "linebreak-style": ["error", "unix"],
      "quotes": ["error", "single"],
      "semi": ["error", "always"],
      "no-unused-vars": ["warn"],
      "no-console": ["warn", { "allow": ["warn", "error", "info"] }]
  }
}
