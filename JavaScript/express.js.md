# Express.js

## 설치

### Express & middleware

* express: js server Framework
* body-paser: middleware for body parsing
* morgan: middleware for logging
* rimraf: rm -rf 명령을 위한 library

* dotenv: 환경변수 불러오기 위한 library
* live reload를 위한 nodemon

## 배포

배포 후 환경변수를 production으로 변경해야 함

`NODE_ENV=production node app.js`

앱 내부에서 `process.env.NODE_ENV` 값에 할당되어 express 배포시 최적화 처리됨

## 설정

### `packages.json`

```json
{
  "name": "express-babel",
  /* ... */
  "main": "dist/index.js",
  "scripts": {
    "prestart": "npm run -s build",
    "start": "node -r dotenv/config dist/index.js",
    "dev": "nodemon src/index.js --exec \"node -r dotenv/config -r babel-register\"",
    "clean": "rimraf dist",
    "build": "npm run clean && mkdir -p dist && babel src -s -D -d dist",
    "test": "jest --watch",
    "lint": "esw -w src test"
  },
  "dependencies": {
    "babel-cli": "^6.26.0",
    "babel-plugin-transform-class-properties": "^6.24.1",
    "babel-plugin-transform-object-rest-spread": "^6.26.0",
    "babel-preset-env": "^1.6.1",
    "body-parser": "^1.18.2",
    "express": "^4.16.2",
    "global": "^4.3.2",
    "mongoose": "^5.0.13",
    "morgan": "^1.9.0",
    "pug": "^2.0.0-beta11",
    "rimraf": "^2.6.2"
  },
  "devDependencies": {
    "babel-eslint": "^8.0.3",
    "babel-jest": "^21.2.0",
    "babel-register": "^6.26.0",
    "dotenv": "^4.0.0",
    "eslint": "^4.19.1",
    "eslint-config-airbnb-base": "^12.1.0",
    "eslint-plugin-import": "^2.10.0",
    "eslint-plugin-jest": "^21.3.2",
    "eslint-watch": "^3.1.4",
    "jest": "^21.2.1",
    "nodemon": "^1.12.1",
    "supertest": "^3.0.0"
  },
  "jest": {
    "testEnvironment": "node"
  }
}
```

### `.eslintrc.js`

```js
module.exports = {
  parser: "babel-eslint",
  plugins: [
    "import", "jest"
  ],
  parserOptions: {
    ecmaVersion: 2017,
    sourceType: "module"
  },
  env: {
    node: true,
    jest: true
  },
  extends: [
    "airbnb-base"
  ],
  rules: {
    "jest/no-focused-tests": 2,
    "jest/no-identical-title": 2
  }
};
```

### `.babelrc`

```json
{
  "presets": [
    [
      "env",
      {
        "targets": {
          "node": "current"
        }
      }
    ]
  ],
  "plugins": [
    "transform-object-rest-spread",
    "transform-class-properties"
  ]
}
```

## CODE

### `index.js`

```js
import app from './app';

const { PORT } = process.env;
app.listen(PORT, () => console.log(`Listening on port ${PORT}`)); // eslint-disable-line no-console
```

### `.env`

```text
PORT=8080
```

### `app.js`

```js
import express from 'express';
import path from 'path';
import logger from 'morgan';
import bodyParser from 'body-parser';
import routes from './routes';
import api from './api';
import database from './model/connector';

const app = express();
app.disable('x-powered-by');

// setup template engine
app.set('views', path.join(__dirname, '../views'));
app.set('view engine', 'pug');

app.use(logger('dev', {
  skip: () => app.get('env') === 'test',
}));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, '../public')));

// Database Connection
// database.connect();

// Routes
app.use('/', routes);
app.use('/api', api);

// Catch 404 and forward to error handler
app.use((req, res, next) => {
  const err = new Error('Not Found');
  err.status = 404;
  next(err);
});

// Error handler
app.use((err, req, res, next) => { // eslint-disable-line no-unused-vars
  res
    .status(err.status || 500)
    .render('error', {
      message: err.message,
    });
});

export default app;
```

### `routes.js`

```js
import { Router } from 'express';

const routes = Router();

/**
 * GET home page
 */
routes.get('/', (req, res) => {
  res.render('index', { title: 'Express Babel' });
});

/**
 * GET /list
 *
 * This is a sample route demonstrating
 * a simple approach to error handling and testing
 * the global error handler. You most certainly want to
 * create different/better error handlers depending on
 * your use case.
 */
routes.get('/list', (req, res, next) => {
  const { title } = req.query;

  if (title == null || title === '') {
    // You probably want to set the response HTTP status to 400 Bad Request
    // or 422 Unprocessable Entity instead of the default 500 of
    // the global error handler (e.g check out https://github.com/kbariotis/throw.js).
    // This is just for demo purposes.
    next(new Error('The "title" parameter is required'));
    return;
  }

  res.render('index', { title });
});

export default routes;
```

## TEST

### `routes.test.js`

```js
import request from 'supertest';
import app from '../src/app.js';

describe('GET /', () => {
  it('should render properly', async () => {
    await request(app).get('/').expect(200);
  });
});

describe('GET /list', () => {
  it('should render properly with valid parameters', async () => {
    await request(app)
      .get('/list')
      .query({ title: 'List title' })
      .expect(200);
  });

  it('should error without a valid parameter', async () => {
    await request(app).get('/list').expect(500);
  });
});

describe('GET /404', () => {
  it('should return 404 for non-existent URLs', async () => {
    await request(app).get('/404').expect(404);
    await request(app).get('/notfound').expect(404);
  });
});
```
