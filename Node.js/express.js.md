# Express.js

## 소개

## 설치

`yarn add express`

### eslint 설치

`yarn add --dev eslint`

`eslint --init`

### babel 설치

`yarn add --dev "eslint-babel babel-cli babel-core babel-preset-es2015 babel-preset-stage-0`

`packages.json`

```json
  "scripts": {
      "dev": "nodemon -w src --exec \"babel-node src --presets es2015,stage-0\"",
      "build": "babel src -s -D -d dist --presets es2015,stage-0",
      "start": "node dist",
      "prestart": "npm run -s build",
      "test": "eslint src"
    }
  }
}
```

`.eslintrc.js`

```js
  extends: 'airbnb-base',
  parserOptions: {
    parser: 'babel-eslint',
    ecmaVersion: 6,
    sourceType: 'module'
  },
  env: {
    'node': true
  },
  rules: {
    'no-console': 0,
    'no-unused-vars': 1
  }
```

`.babelrc`

```text
{
  "presets": [
    ["es2015", { "modules": false }],
    ["stage-0"]
  ],
  "env": {
    "test": {
      "presets": [
        ["env"]
      ]
    }
  }
}
```

## template engine

### pug

`yarn add pug`
