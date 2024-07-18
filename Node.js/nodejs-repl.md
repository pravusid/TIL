# NodeJs REPL

Read-Eval-Print-Loop

<https://nodejs.org/api/repl.html>

## Commands

<https://nodejs.org/api/repl.html#repl_commands_and_special_keys>

## Custom REPL

<https://www.stefanjudis.com/today-i-learned/how-to-create-your-own-node-js-repl/>

`index.js`

```js
// index.js
const repl = require('repl');

// define available methods and state
const state = {
  printSomething() {
    console.log("That's awesome!");
  }
};

const myRepl = repl.start("custom repl > ");

Object.assign(myRepl.context, state);
```

실행

```bash
$ node index.js

custom repl > printSomething()
# That's awesome!
```
