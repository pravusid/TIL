# Redux

## Flux

MVC패턴에서 model은 view로 데이터를 보내 렌더링을 한다

사용자와 App의 상호작용은 view를 통해 일어나기 때문에 사용자의 입력에 따라 view가 model을 업데이트 해야할 때도 있다.
업데이트는 비동기적으로 일어날 수 있고 1:N으로 일어날 수도 있다.

문제를 해결하기 위해 데이터를 단방향으로만 보내는 Flux 아키텍처를 만들어낸다.

(Action) -> Dispatcher -> Store -> View -> (Action) -> Dispatcher

### 액션 생성자

App의 상태를 변경하거나 view를 업데이트하려면 액션을 생성해야 한다.

액션 생성자는 시스템에 정의된 액션인 type과 payload를 포함한 액션을 생성한다.

액션 생성자가 메시지를 생성한 뒤 dispatcher로 넘겨준다

### 디스패쳐

dispatcher는 callback이 등록되어 있는 곳이다. 액션이 넘어오면 여러 스토어에 액션을 보낸다.

동기적으로 액션을 처리하고 스토어사이에 의존성이 있다면 순서에따라 처리한다.

Flux의 dispatcher는 다른아키텍쳐들과 다르게 액션타입과 관계없이 등록된 모든 스토어로 보내고, 스토어는 액션을 받은뒤 처리여부를 결정한다.

### 스토어

모든 상태변경은 store에 의해서 결정되어야 한다.

store에는 setter가 없으므로 상태변경을 위한 요청을 store에 직접보낼 수 없다. 즉 액션생성자, 디스패처를 통해 이루어져야한다.

스토어 내부에서 swtich statement를 사용해 처리할 액션과 무시할 액션을 결정하고, 처리할 액션에 따라 상태를 변경한다.

스토어에서 상태를 변경하고 나면 controller view에 change event를 내보낸다.

### 컨트롤러 뷰

controller view는 store와 view사이의 중간관리자 역할을 한다.

상태가 변경되었을 때 store가 사실을 controller view에 알려주면 controller view는 자신 아래에 있는 모든 view에게 새로운 상태를 넘긴다.

## Redux 소개

<https://redux.js.org/introduction/getting-started>

Flux, Elm, immutable.js, baobab, RxJS 등의 영향을 받아 만들어짐

### Redux 원칙

- Single Source of Truth
  - 어플리케이션의 state를 위해 단 하나의 store(상태트리)를 사용

- State is Read-only
  - app에서 store의 state를 직접 변경할 수 없고 변경을 위해서는 action을 emit 하여 dispatch되어야 함
  - view나 네트워크 콜백이 직접적으로 상태를 변경하지 않는다, 대신 상태를 변경하려는 의도를 전달한다
  - 모든 변경은 중앙에 집중되어 있고 변경은 순차적으로 이루어지므로 어더한 경쟁 상태도 없다

- Changes are made with pure Function
  - action객체를 처리하는 함수를 reducer라고 부른다
  - reducer는 정보를 받아서 state를 어떻게 업데이트할지 정의한다(이전상태를 받아 다음 상태를 반환)
  - reducer는 '순수 함수'로 작성되어야 함 : 비동기 처리, 네트워크 및 DB접근, 인수변경, API사용(Math.random()...) => 불가

### Actions

액션은 app에서 store로 데이터를 보내는 정보집합이다. 액션은 store에 제공되는 유일한 정보 원천이다.

`store.dispatch()`를 사용하여 store로 데이터를 보낸다

```js
const ADD_TODO = 'ADD_TODO'

{
  type: ADD_TODO,
  text: 'Build my first Redux app'
}
```

액션은 POJO이다. 액션은 작업 유형을 나타내는 `type` 프로퍼티가 있어야 한다. 일반적으로 타입은 문자열 상수로 정의해야 한다.

> 타입별 액션을 별도 파일로 분리하지 않아도 된다. 또한 소규모 프로젝트의 경우 타입에 문자열 리터럴을 사용하는 것이 더 쉬울 수 있다.
> [큰 규모의 프로젝트에서 보일러 플레이트 줄이기](https://redux.js.org/recipes/reducing-boilerplate)를 참고

타입을 제외한 액션 객체 구조는 작성하는 사람에 달려있다. 참고: [Flux 표준 액션](https://github.com/acdlite/flux-standard-action)

사용자가 할 일을 완료한 것으로 처리하는 액션 타입을 추가한다.

```js
{
  type: TOGGLE_TODO,
  index: 5
}
```

각 작업에서 가능한 최소 데이터를 전달하는 것이 좋다(전체 할 일 객체보다는 참조를 위한 index 전달)

마지막으로 현재 보이는 할일을 변경하기 위한 액션 타입을 추가한다

```js
{
  type: SET_VISIBILITY_FILTER,
  filter: SHOW_COMPLETED
}
```

#### Action Creators

action creators는 액션을 만드는 기능이다.

Redux에서 action creators는 단순히 action을 반환한다

```js
function addTodo(text) {
  return {
    type: ADD_TODO,
    text
  };
}
```

전통적인 Flux에서 action creators는 종종 호출되었을 때 dispatch를 트리거한다

```js
function addTodoWithDispatch(text) {
  const action = {
    type: ADD_TODO,
    text
  };
  dispatch(action);
}
```

Redux에서는 그렇게 처리하지 않고, 대신 action creators의 반환값을 `dispathch()` 함수에 넘긴다

```js
dispatch(addTodo(text));
dispatch(completeTodo(index));
```

대신 bound action creator를 만들어서 dispatch를 동시에 처리 할 수 있다

```js
const boundAddTodo = text => dispatch(addTodo(text));
const boundCompleteTodo = index => dispatch(completeTodo(index));

boundAddTodo(text);
boundCompleteTodo(index);
```

`dispatch()` 함수는 `store.dispatch()`로 상점에서 직접 접근할 수 있지만,
react-redux의 `connect()` 같은 helper를 사용하여 접근할 가능성이 높다.

`bindActionCreators()`를 사용하여 여러 action creators를 `dispatch()` 함수에 자동 바인딩 할 수 있다.

action creators는 비동기일 수 있으며 side effect가 존재할 수 있다.

#### actions.js

```js
/*
 * action types
 */

export const ADD_TODO = "ADD_TODO";
export const TOGGLE_TODO = "TOGGLE_TODO";
export const SET_VISIBILITY_FILTER = "SET_VISIBILITY_FILTER";

/*
 * other constants
 */

export const VisibilityFilters = {
  SHOW_ALL: "SHOW_ALL",
  SHOW_COMPLETED: "SHOW_COMPLETED",
  SHOW_ACTIVE: "SHOW_ACTIVE"
};

/*
 * action creators
 */

export function addTodo(text) {
  return { type: ADD_TODO, text };
}

export function toggleTodo(index) {
  return { type: TOGGLE_TODO, index };
}

export function setVisibilityFilter(filter) {
  return { type: SET_VISIBILITY_FILTER, filter };
}
```

### Reducers

reducer는 store에 전송한 action에 대응하여 app의 상태가 어덯게 변경되는지 지정한다.
action은 무엇이 일어났는지만 설명하지, app의 상태가 어덯게 변하는지를 설명하지는 않는다.

#### Designing the State

Redux에서는 모든 app 상태가 단일 객체로 저장된다.

할 일 app의 경우 두 가지 다른 내용을 저장하려 한다

- 현재 선택된 visibility filter
- 실제 할 일의 목록

상태트리에 일부 데이터와 UI 상태를 저장하는 경우가 있지만, 데이터로부터 UI 상태는 분리하는 것이 좋다.

```js
{
  visibilityFilter: 'SHOW_ALL',
  todos: [
    {
      text: 'Consider using Redux',
      completed: true
    },
    {
      text: 'Keep all state in a single tree',
      completed: false
    }
  ]
}
```

> 복잡한 app에서는 서로 다른 엔티티가 각자를 참조할 것이다. 따라서 상태를 가능한 중첩없이 정규화 하는 것이 좋다.
> key로 ID를 사용하여 모든 객체를 저장하고 다른 엔티티를 참조할 때 ID를 사용한다. app의 상태를 데이터베이스처럼 생각한다.

#### Handling Actions

상태 객체처럼 보이는 것을 준비했으므로 reducer를 작성해보자.

reducer는 이전 상태와 action으로 다음 상태를 반환하는 순수 함수이다.

```js
(previousState, action) => newState;
```

reducer라고 불리는 이유는 `Array.prototype.reduce(reducer[, initialValue])`에 전달될 타입의 함수이기 때문이다.
reducer가 순수함수인 것이 매우 중요하다. reducer 내부에서 절대 하면 안될 것들이 있다.

- 인자를 변형하는(mutate) 것
- 부수효과(side-effect)를 발생시키는 것(API call, routing transtions...)
- non-pure 함수 호출 (`Date.now()`, `Math.random()`)

초기 상태를 지정하여 시작해보자.
redux는 최초 reducer를 `undefined` 상태와 함께 호출된다.

```js
import { VisibilityFilters } from "./actions";

const initialState = {
  visibilityFilter: VisibilityFilters.SHOW_ALL,
  todos: []
};

function todoApp(state = initialState, action) {
  // 지금은 액션을 다루지 않고 주어진 상태를 반환함
  return state;
}
```

이제는 `SET_VISIBILITY_FILTER`를 처리해보자. 할 일은 `visibilityFilter` 상태를 변경하는 것 뿐이다.

```js
import { SET_VISIBILITY_FILTER, VisibilityFilters } from "./actions";

// ...

function todoApp(state = initialState, action) {
  switch (action.type) {
    case SET_VISIBILITY_FILTER:
      return Object.assign({}, state, {
        visibilityFilter: action.filter
      });
    default:
      return state;
  }
}
```

다음이 중요하다

- **상태를 변경하지 않는다**: `Object.assign()`을 이용해서 복사본을 만든다. 대신 `{ ...state, ...newState }`를 사용할 수도 있다.
- **기본 타입으로 이전 상태를 반환한다**: 알수 없는 action을 대상으로 이전 상태를 반환한다

> `switch` 문은 실제 보일러 플레이트가 아니다. Flux의 진짜 보일러 플레이트는 개념이다.
> (emit 하고 update, store에 등록하고 이를 dispatch하는 일, store가 객체가 되는 일)
> Redux는 이를 event emitter 대신 pure reducer로 해결한다.
> 불행히도 많은 곳에서 `switch`문의 사용여부에 따라 프레임워크를 선택한다.
> `switch`를 사용하지 않으려면 [보일러플레이트 줄이기](https://redux.js.org/recipes/reducing-boilerplate#reducers)대로 handler map을 허용하는 `createReducer` 함수를 사용한다.

#### Handling More Actions

처리해야 할 작업이 두 가지 더 있다.
`SET_VISIBILITY_FILTER`와 마찬가지로 `ADD_TODO` 및 `TOGGLE_TODO` action을 가져온 다음 reducer를 확장하여 이를 처리한다.

```js
import {
  ADD_TODO,
  TOGGLE_TODO,
  SET_VISIBILITY_FILTER,
  VisibilityFilters
} from "./actions";

// ...

function todoApp(state = initialState, action) {
  switch (action.type) {
    case SET_VISIBILITY_FILTER:
      return Object.assign({}, state, {
        visibilityFilter: action.filter
      });
    case ADD_TODO:
      return Object.assign({}, state, {
        todos: [
          ...state.todos,
          {
            text: action.text,
            completed: false
          }
        ]
      });
    case TOGGLE_TODO:
      return Object.assign({}, state, {
        todos: state.todos.map((todo, index) => {
          if (index === action.index) {
            return Object.assign({}, todo, {
              completed: !todo.completed
            });
          }
          return todo;
        })
      });
    default:
      return state;
  }
}
```

`TOGGLE_TODO`에서 특정한 item을 mutation에 의존하지 않고 갱신하려 하려고 하므로 인덱스 항목을 제외하고 동일한 항목으로 새 배열을 만들어야 한다.
이러한 작업을 자주 처리하는 경우 immutable 같은 라이브러리를 사용하는 것이 좋다.

값 복제를 하지 않은 상태에서 state 내부에 어떠한 처리도 해서는 안된다.

#### Splitting Reducers

위의 verbose한 코드를 보기 쉽게 만드는 방법이 있을까?

코드에서 `todos`와 `visibilityFilter`는 완전히 독립적으로 갱신되는 것으로 보인다.

```js
import { combineReducers } from "redux";
import {
  ADD_TODO,
  TOGGLE_TODO,
  SET_VISIBILITY_FILTER,
  VisibilityFilters
} from "./actions";
const { SHOW_ALL } = VisibilityFilters;

function visibilityFilter(state = SHOW_ALL, action) {
  switch (action.type) {
    case SET_VISIBILITY_FILTER:
      return action.filter;
    default:
      return state;
  }
}

function todos(state = [], action) {
  switch (action.type) {
    case ADD_TODO:
      return [
        ...state,
        {
          text: action.text,
          completed: false
        }
      ];
    case TOGGLE_TODO:
      return state.map((todo, index) => {
        if (index === action.index) {
          return Object.assign({}, todo, {
            completed: !todo.completed
          });
        }
        return todo;
      });
    default:
      return state;
  }
}

const todoApp = combineReducers({
  visibilityFilter,
  todos
});

export default todoApp;
```

분리한 각각의 reducer는 global state의 각 부분을 관리한다.
app이 더 커지면 각 reducer는 파일로 분할하여 독립적으로 유지하고 다른 데이터 도메인을 관리하게 할 수 있다.

마지막으로 Redux는 `combineReducers()`라는 유틸리티를 제공한다.
유틸리티의 도움을 받아 여러 reducer는 하나로 결합할 수 있다.

### Store

Store는 actions와 reducers를 하나로 묶는 객체이다. store는 다음의 책임을 따른다.

- 애플리케이션 상태를 보유한다
- `getState()`를 통해 state에 접근할 수 있다
- `dispatch(action)`을 통해 상태를 갱신할 수 있다
- `subscribe(listener)`를 통해 리스너를 등록한다
- `subscribe(listener)`가 밯놘함 함수를 통해 리스너 등록을 해제한다

Redux app에는 단일 저장소만 있다는 사실에 유의해야 한다.
데이터 처리 로직을 분리하려는 경우 여러개의 store 대신 reducer 합성을 사용해야 한다.

reducer가 있으면 store를 쉽게 만들 수 있다.
이전 섹션에서 `combineReducers()`를 사용하여 어러 reducer를 하나로 결합하였는데, 이를 `createStore()`에 전달한다.

```js
import { createStore } from "redux";
import todoApp from "./reducers";
const store = createStore(todoApp);
```

선택적으로 초기 state를 `createStore()`의 두 번째 인자값으로 지정할 수 있다.

```js
const store = createStore(todoApp, window.STATE_FROM_SERVER);
```

#### Dispatching Actions

사용예시

```js
import {
  addTodo,
  toggleTodo,
  setVisibilityFilter,
  VisibilityFilters
} from "./actions";

// Log the initial state
console.log(store.getState());

// Every time the state changes, log it
// Note that subscribe() returns a function for unregistering the listener
const unsubscribe = store.subscribe(() => console.log(store.getState()));

// Dispatch some actions
store.dispatch(addTodo("Learn about actions"));
store.dispatch(addTodo("Learn about reducers"));
store.dispatch(addTodo("Learn about store"));
store.dispatch(toggleTodo(0));
store.dispatch(toggleTodo(1));
store.dispatch(setVisibilityFilter(VisibilityFilters.SHOW_COMPLETED));

// Stop listening to state updates
unsubscribe();
```

### Data flow

Redux 아키텍처는 엄격한 단방향 데이터 흐름을 중심으로 진행된다

이는 app의 모든 데이터가 동일한 수명 주기 패턴을 따르므로 app의 로직을 예측가능하고 이해하기 쉽게 만든다.
또한 데이터 정규화를 권장하므로 서로 알지 못하는 여러개의 독립적인 동일 데이터의 사본이 생성되지 않는다.

Redux 앱의 데이터 수명주기는 다음 4단계를 따른다

#### `store.dispatch(action)` 호출

action은 무엇이 일어났는지를 기술하는 plain javascript object 이다. 예를 들면,

```js
{ type: 'LIKE_ARTICLE', articleId: 42 }
{ type: 'FETCH_USER_SUCCESS', response: { id: 3, name: 'Mary' } }
{ type: 'ADD_TODO', text: 'Read the Redux docs.' }
```

action을 아주 간단한 뉴스 조각으로 생각해보자.

컴포넌트, XHR 콜백, 예약된 간격 같은, 앱의 어느곳에서나 `store.dispatch(action)`를 호출할 수 있다.

#### 리덕스 store는 사용자가 제공한 reducer funtion을 호출

store는 현재 상태트리와 action이라는 두 가지 인자를 reducer에 전달한다.
예를 들어, 할 일 앱에서 root reducer는 다음과 같은 것을 받을 것이다.

```js
// 현재 애플리케이션 상태(선택된 필터에서 해야 할 일 목록)
let previousState = {
  visibleTodoFilter: 'SHOW_ALL',
  todos: [
    {
      text: 'Read the docs.',
      complete: false
    }
  ]
}

// 수행될 action(할 일에 추가)
let action = {
  type: 'ADD_TODO',
  text: 'Understand the flow.'
}

// reducer는 다음 애플리케이션 상태를 반환한다
let nextState = todoApp(previousState, action)
```

reducer는 순수 함수이다. 오로지 다음 상태만 계산한다. 따라서 완벽히 예측 가능해야 한다.
동일한 입력으로 여러 번 호출하여도 동일한 출력이 생성된다.

side-effect가 있는 API 호출, 라우터 전환 같은 행위를 수행해서는 안되며, 이는 action이 dispatch 되기전 실행되어야 한다.

#### root reducer는 여러 reducer의 출력을 단일 상태 트리로 결합한다

root reducer를 구성하는 방법은 전적으로 사용자에게 달려있다.
Redux는 각각의 상태 트리 분기를 관리하는 reducer로 분할 할 수 있도록 `combineReducers()` helper 함수를 제공한다.

`combineReducers()` 작동 방식은 다음과 같다.
두 개의 reducer가 있다. 하나는 할 일 목록용이고 다른 하나는 필터 선택에 관한 것이다.

```js
function todos(state = [], action) {
  // Somehow calculate it...
  return nextState
}

function visibleTodoFilter(state = 'SHOW_ALL', action) {
  // Somehow calculate it...
  return nextState
}

let todoApp = combineReducers({
  todos,
  visibleTodoFilter
})
```

action을 발생시키면, `combineReducers`는 두 개의 reducer를 호출하는 `todoApp`을 반환한다.

```js
let nextTodos = todos(state.todos, action)
let nextVisibleTodoFilter = visibleTodoFilter(state.visibleTodoFilter, action)
```

그런 다음 두 결과 집합을 단일 상태트리로 결합한다

```js
return {
  todos: nextTodos,
  visibleTodoFilter: nextVisibleTodoFilter
}
```

`combineReducers()`는 편리한 helper이지만 반드시 사용하지 않아도 되고, 각자의 root reducer를 작성할 수 있다.

#### Redux store는 root reducer의 반환한 전체 상태트리를 저장한다

새로운 트리는 앱의 다음 상태이다.
이제는 `store.subscribe(listener)`에 등록된 모든 리스너가 호출된다.
리스너는 `store.getState()`를 호출하여 현재 상태를 얻는다.

이제 새로운 상태를 반영하도록 UI를 업데이트 할 수 있다.
만약 `react-redux`와 같은 바인딩을 사용하는 경우 `component.setState(newState)`가 호출된다.

### React Redux
