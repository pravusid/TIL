# Node.js Event Loop

- <https://nodejs.org/en/learn/asynchronous-work/event-loop-timers-and-nexttick>
- <https://nodejs.org/en/learn/asynchronous-work/dont-block-the-event-loop>
- <https://evan-moon.github.io/2019/08/01/nodejs-event-loop-workflow/>
- <https://www.korecmblog.com/node-js-event-loop/>
- <https://dev.to/altamashali/nodejs-event-loop-in-action-d5o>

## Node.js 이벤트 루프

> 이벤트 루프는 가능하다면 언제나 시스템 커널에 작업을 떠넘겨서(libuv) Node.js가 논 블로킹 I/O 작업을 수행하도록 해줍니다.(JavaScript가 싱글 스레드임에도 불구하고)

### 이벤트 루프 개요

다음은 이벤트 루프의 단계이다

```txt
   ┌───────────────────────────┐
┌─>│           timers          │
│  └─────────────┬─────────────┘
│  ┌─────────────┴─────────────┐
│  │     pending callbacks     │
│  └─────────────┬─────────────┘
│  ┌─────────────┴─────────────┐
│  │       idle, prepare       │
│  └─────────────┬─────────────┘      ┌───────────────┐
│  ┌─────────────┴─────────────┐      │   incoming    │
│  │           poll            │<─────┤  connections, │
│  └─────────────┬─────────────┘      │   data, etc.  │
│  ┌─────────────┴─────────────┐      └───────────────┘
│  │           check           │
│  └─────────────┬─────────────┘
│  ┌─────────────┴─────────────┐
└──┤      close callbacks      │
   └───────────────────────────┘
```

- 각 단계마다 실행할 콜백의 FIFO 큐(큐가 아닌 다른 자료구조일수도 있음)가 존재(idle, prepare 제외)한다
- 큐의 작업을 모두 실행하거나, 콜백 제한(timers 단계에서는 없음)에 이르면 다음단계로 이동한다
- 큐에서 실행한 작업이 또 다른 작업을 추가하거나, poll 단계에서 처리된 새로운 이벤트가 커널에 의해 큐에 추가될 수 있다
- 이벤트 루프 진입점은 어디일까?
  - 여러 자료에서 timers 페이즈부터 시작하는 것으로 설명하고 있다
  - 그러나, 스크립트를 실행하면 이벤트 루프부터 생성한 뒤 스크립트를 읽기 때문에 이벤트 루프는 poll 페이즈에서 대기할 것이라 생각된다
  - 실제로 `setTimeout(() => { console.log('1') }, 0)` 다음 `Promise.resolve().then(() => console.log('2'))` 순서로 실행하면 `2, 1`이 출력됨을 확인할 수 있다
  - 즉, timers 페이즈 전 micro task가 실행되었다는 의미이다

### 이벤트 루프 단계

#### timers

- 이 페이즈가 가지고 있는 큐(min-heap 자료구조사용)에는 setTimeout이나 setInterval 같은 타이머들의 콜백을 저장하게 된다
- 타이머는 페이즈가 돌아왔을 때 실행조건이 맞다면 실행된다 (반드시 지정한 시간에 실행되는 것은 아니다; delay는 최소보장 대기시간임)

#### Pending i/o callback phase

- 이전 이벤트 루프에서 실행된 작업들의 콜백이 실행 대기 중인지, 즉 pending_queue에 들어와 있는지를 확인함
- 만약 실행 대기 중이라면 pending_queue가 비거나 시스템의 실행 한도 초과에 도달할 때까지 대기하고 있던 콜백들을 실행

#### Idle, Prepare phase

- Idle 페이즈는 매 Tick마다 실행된다
- Prepare 페이즈는 매 폴링(Polling)때마다 실행된다
- Node.js 관리 목적의 페이즈이다

#### Poll phase

- I/O 이벤트 콜백을 watch_queue에서 관리한다
- watch_queue는 작업 순서와 관계없이 우선 완료되는 작업을 처리하기 위해 File Descriptor 정보를 가지고 있다
- 만약 watch_queue(Poll phase가 가지고 있는 큐)가 비어있지 않다면, 큐가 비거나 시스템 최대 실행 한도에 다다를 때까지 동기적으로 모든 콜백을 실행한다
- 일단 watch_queue가 비어있다면, Node.js는 곧바로 다음 페이즈로 넘어가는 것이 아니라 약간 대기시간을 가지게 된다
  - check_queue, pending_queue, closing_callbacks_queue 순서로 작업이 있는지 확인하고 작업이 있는 단계로 진행한다
  - timers_queue에 작업이 있다면 다음 실행가능한 타이머의 시간까지 대기한다
  - 확인한 모든 큐에 작업이 없으면 대기한다

#### Check phase

- setImmediate()는 이벤트 루프의 별도 단계에서 실행되는 특수한 타이머이고 check 페이즈에서 실행된다
- setImmediate()는 poll 단계가 완료된 후 콜백 실행을 스케줄링하는데 libuv API를 사용한다

#### Close callbacks

- `close` 이벤트의 콜백이 실행된다
- 모든 작업이 끝나면 (또는 실행할 작업이 없으면), 다시 timers 페이즈로 진행한다

### NextTick, MicroTask

[node v11 이후](https://github.com/nodejs/node/issues/22257) 현재실행하고 있는 작업이 끝나면 즉시 실행된다 (크롬 브라우저 구현과 같은 순서가 되었다)

#### nextTickQueue

- process.nextTick()으로 실행하는 nextTickQueue는 이벤트 루프의 일부가 아니다
- 페이즈와는 다르게 시스템 실행한도 초과에 영향을 받지 않으므로 I/O starvation을 유발할 수 있다

#### microTaskQueue

- 비동기작업 관리를 위해서 ECMA에선 PromiseJobs라는 내부 큐(internal queue)를 명시하고 있으며, V8 엔진에선 이를 '마이크로태스크 큐(microtask queue)' 라고 부른다
- nextTickQueue 다음 실행된다
- 이행된 프로미스의 핸들러가 microTaskQueue에 들어간다
- 실행할 것이 아무것도 남아있지 않을 때만 microTaskQueue에 있는 작업이 실행되기 시작한다
- microTaskQueue는 먼저 들어온 작업을 먼저 실행한다 (FIFO, first-in-first-out)

## Web Browser Event Loop

- <https://dev.to/jasmin/difference-between-the-event-loop-in-browser-and-node-js-1113>
- <https://blog.insiderattack.net/javascript-event-loop-vs-node-js-event-loop-aea2b1b85f5c>
- <https://ko.javascript.info/event-loop>
- <https://ko.javascript.info/microtask-queue>

이벤트 루프는 ECMA 표준에 명시되어 있지 않으므로 런타임에 따라 구현이 다르다
