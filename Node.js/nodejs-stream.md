# Stream in Node.js

- <https://nodejs.org/api/stream.html>
- <https://github.com/FEDevelopers/tech.description/wiki/Node.js-Stream-당신이-알아야할-모든-것>

## 종류

- `Readable`: 추상화된 읽기 가능한 데이터 stream
- `Writable`: 추상화된 쓰기 가능한 데이터 stream
- `Duplex`: 읽기/쓰기 모두 가능한 데이터 stream
- `Transform`: `Writable` stream을 입력받아 `Readable` stream을 출력하는 stream

HTTP의 경우 (Readable)`http.ImcomingMessage` -> (Writable)`http.ServerResponse`

## 사용

### pipe

`Readable stream`으로부터 `Duplex` 혹은 `Transform` 혹은 `Writable stream`으로 `pipe`를 사용할 수 있다.

```js
readableStream.pipe(transformStream1).pipe(transformStream2).pipe(writable);
```

<https://nodejs.org/api/stream.html#readablepipedestination-options>

> - Pipe에서는 기본적으로 소스인 Readable stream에서 **'end' 이벤트**가 발생하면, Writable stream의 `end()`가 호출되어 더 이상 쓰기가 작동하지 않는다
> - 이 기본기능을 비활성화 하려면 `{ end: false }` 옵션을 사용하면되고, destination Writable stream은 열려있는 상태가 유지된다
> - 중요한 주의사항은, Readable stream 처리중 오류가 발생하면 Writable destination은 자동으로 닫히지 않는다는 것이다 (메모리 누수 방지를 위해 오류발생시 수동으로 닫아야 함)

#### pipe & backpressuring

- <https://nodejs.org/en/learn/modules/backpressuring-in-streams>
- <https://nodejs.org/api/stream.html#api-for-stream-implementers>
- [Do transform streams have to handle backpressure as well?](https://github.com/nodejs/help/issues/2695)

pipeline & backpressuring 예제

```js
import { createReadStream } from 'node:fs';
import { createInterface } from 'node:readline/promises';
import { Writable } from 'node:stream';
import { pipeline } from 'node:stream/promises';

const result = [];

await pipeline(
  createReadStream('./input.log'),
  (input) => createInterface({ input }),
  new Writable({
    objectMode: true,
    highWaterMark: 500,
    writev: async ([{ chunk }], cb) => {
      if (!chunk) {
        return cb();
      }

      const pos = chunk.indexOf(' web: {');
      try {
        const o = JSON.parse(chunk.slice(pos + 6));
        result.push(o);
        cb();
      } catch (e) {
        // cb(e); // throw error
        cb(); // noop
      }
    },
  })
);

console.log(result);
```

### event

모든 stream은 `EventEmitter`의 인스턴스이므로 event handler를 사용하여 stream을 처리할 수 있다.

```js
readable.pipe(writable);
```

위의 코드는 다음과 같을 것이다.

```js
readable.on('data', (chunk) => {
  writable.write(chunk);
});

readable.on('end', () => {
  writable.end();
});
```

- event on Readable stream

  - `data`: consumer에게 chunk가 전송될 때
  - `end`: 스트림에 더이상 데이터가 없을 때
  - `error`
  - `close`
  - `readable`

- event on Writable stream

  - `drain`: Writable stream이 추가적인 데이터를 수신할 수 있을 때
  - `finish`: 모든 데이터가 flush될 때
  - `error`
  - `close`
  - `pipe`/`unpipe`

### Paused / Flowing

- Readable stream은 기본적으로 Paused 상태이나 Flowing으로 상태를 변경할 수 있다.
- Paused Readable stream은 `read()` 메소드로 처리할 수 있다.
- Flowing Readable stream은 event handler에 의해서 처리되고 consumer가 없다면 데이터는 사라진다.
- 두 상태간의 변경은 `resume()`(to Flowing) / `pause()`(to Paused) 메소드를 사용하면 된다.
- `pipe` 메소드를 사용한다면 두 종류의 Readable stream을 모두 처리가능하다.

### Streams Promises API

<https://nodejs.org/api/stream.html#streams-promises-api>

> Added in: v15.0.0

- `stream.pipeline(source[, ...transforms], destination[, options])`
- `stream.finished(stream[, options])`

```js
import { pipeline } from 'node:stream/promises';
import { createReadStream, createWriteStream } from 'node:fs';
import { createGzip } from 'node:zlib';

await pipeline(createReadStream('archive.tar'), createGzip(), createWriteStream('archive.tar.gz'));
console.log('Pipeline succeeded.');
```

### Utility Consumers

> `Readable` stream을 변환 하는데 사용하는 유틸리티 함수이다
>
> --<https://nodejs.org/api/webstreams.html#utility-consumers>

```ts
declare module 'stream/consumers' {
  import { Blob as NodeBlob } from 'node:buffer';
  import { Readable } from 'node:stream';
  function buffer(stream: NodeJS.ReadableStream | Readable | AsyncIterable<any>): Promise<Buffer>;
  function text(stream: NodeJS.ReadableStream | Readable | AsyncIterable<any>): Promise<string>;
  function arrayBuffer(stream: NodeJS.ReadableStream | Readable | AsyncIterable<any>): Promise<ArrayBuffer>;
  function blob(stream: NodeJS.ReadableStream | Readable | AsyncIterable<any>): Promise<NodeBlob>;
  function json(stream: NodeJS.ReadableStream | Readable | AsyncIterable<any>): Promise<unknown>;
}
```

다음 방법으로 Readable stream을 Uint8Array(or Buffer) 타입의 데이터로 변환할 수 있다
([[#Stream <-> Buffer 변환]], Stream to Buffer 기능과 동일)

```js
const uint8Arr = new Uint8Array(await arrayBuffer(readable));
```

## 생성

### Writable stream

```js
const { Writable } = require('stream');

const outStream = new Writable({
  write(chunk, encoding, cb) {
    console.log(chunk.toString());
    cb();
  },
});

process.stdin.pipe(outStream);
```

- chunk: 데이터, 보통은 버퍼
- encoding: 인코딩
- cb: 콜백 (완료 후 호출)

위의 구현은 일종의 `stdout`이다.

```js
process.stdin.pipe(process.stdout);
```

### Readable stream

```js
const inStream = new Readable({
  read(size) {
    this.push(String.fromCharCode(this.currentCharCode++));
    if (this.currentCharCode > 90) {
      this.push(null); // 더 이상 데이터 없음
    }
  },
});
inStream.currentCharCode = 65;

inStream.pipe(process.stdout);
```

읽기 스트림을 읽는 동안 `read` 메소드가 실행된다. 즉, 데이터가 소비될 때 Push 한다.

### Duplex

```js
const { Duplex } = require('stream');

const inoutStream = new Duplex({
  write(chunk, encoding, callback) {
    console.log(chunk.toString());
    callback();
  },

  read(size) {
    this.push(String.fromCharCode(this.currentCharCode++));

    if (this.currentCharCode > 90) {
      this.push(null);
    }
  },
});
inoutStream.currentCharCode = 65;

process.stdin.pipe(inoutStream).pipe(process.stdout);
```

- `process.stdin.pipe(inoutStream)`: 키보드 입력을 받아서 화면에 출력 (echo)
- `inoutStream.pipe(process.stdout)`: 알파벳 입력을 받아서 화면에 출력

Duplex의 읽기/쓰기 stream은 독립적으로 동작하며 단지 둘을 합쳐놓았을 뿐이다.

### Transform

Transform stream에서는 `read`나 `write` 메소드 역할을 동시에 수행하는 `transform` 메소드를 구현한다.

```js
const { Transform } = require('stream');

const toUpperCase = new Transform({
  transform(chunk, encoding, callback) {
    this.push(chunk.toString().toUpperCase());
    callback();
  },
});

process.stdin.pipe(toUpperCase).pipe(process.stdout);
```

기본적으로 스트림은 버퍼(Buffer)나 문자열(String) 값을 소비하지만, `objectMode` 옵션으로 JavaScript 객체를 사용할 수 있다.

```js
const { Transform } = require('stream');

const commaSplitter = new Transform({
  readableObjectMode: true,
  transform(chunk, encoding, callback) {
    this.push(chunk.toString().trim().split(','));
    callback();
  },
});

const arrayToObject = new Transform({
  readableObjectMode: true, // 내보낼 때
  writableObjectMode: true, // 받을 때
  transform(chunk, encoding, callback) {
    const obj = {};
    for (let i = 0; i < chunk.length; i += 2) {
      obj[chunk[i]] = chunk[i + 1];
    }
    this.push(obj);
    callback();
  },
});

const objectToString = new Transform({
  writableObjectMode: true,
  transform(chunk, encoding, callback) {
    this.push(JSON.stringify(chunk) + '\n');
    callback();
    // 콜백의 두번째 인자로 push할 수도 있다
    // callback(null, JSON.stringify(chunk) + "\n");
  },
});

process.stdin.pipe(commaSplitter).pipe(arrayToObject).pipe(objectToString).pipe(process.stdout);
```

## Examples

### 파일을 읽어서 압축하는 도중 상태를 출력하며, 완료시 "Done" 메시지를 출력하는 예제

```js
const fs = require('fs');
const zlib = require('zlib');
const file = process.argv[2];

fs.createReadStream(file)
  .pipe(zlib.createGzip())
  .on('data', () => process.stdout.write('.'))
  .pipe(fs.createWriteStream(file + '.gz'))
  .on('finish', () => console.log('Done'));
```

event handler를 등록하는 대신 Transform을 사용할 수 있다

```js
// ...
const { Transform } = require('stream');

fs.createReadStream(file)
  .pipe(zlib.createGzip())
  .pipe(
    new Transform({
      transform(chunk, encoding, callback) {
        process.stdout.write('.');
        callback(null, chunk);
      },
    })
  )
  .pipe(fs.createWriteStream(file + '.zz'))
  .on('finish', () => console.log('Done'));
```

### Stream <-> Buffer 변환

Stream to Buffer

```js
function streamToBuffer(stream) {
  return new Promise((resolve, reject) => {
    const buffers = [];
    stream.on('error', reject);
    stream.on('data', (data) => buffers.push(data));
    stream.on('end', () => resolve(Buffer.concat(buffers)));
  });
}
```

Buffer to Stream

```js
let Readable = require('stream').Readable;

function bufferToStream(buffer) {
  const stream = new Readable();
  stream.push(buffer);
  stream.push(null);
  return stream;
}
```

Buffer to Stream with PassThrough

```js
let PassThrough = require('stream').PassThrough;

function bufferToStream(buffer) {
  const stream = new PassThrough();
  stream.end(buffer);
  return stream;
}
```
