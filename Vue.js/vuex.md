# Vuex

## Vuex란

Vue.js 애플리케이션에 대한 상태 관리 패턴 / 라이브러리

모든 컴포넌트에 대해서 중앙 집중식 저장소 역할을 하고 예측가능한 방식으로 상태변경

React의 Flux 구현체인 Redux와 유사하나 차이점 존재

## Vuex store

store는 App 상태를 보유하고있는 컨테이너

Vue store는 반응형이다. Vue 컴포넌트는 저장소의 상태가 변경되면 업데이트 된다.

저장소의 상태를 직접 변경할 수 없고(immutable) 명시적인 커밋을 이용하여 변이한다.
이렇게 하면 모든 상태에 대한 추적이 가능한 기록이 남는다.

### 저장소 사용

```js
Vue.use(Vuex);
const store = new Vuex.Store({
  state: {
    count: 0
  },
  mutations: {
    increment (state) {
      state.count++;용
    },
  },
});
```

state 객체에 접근하여 store.commit 메소드로 상태변경을 일으킬 수 있다.

```js
this.$store.commit('increment'); // root instance에서 주입된 store 호출 후 commit
```

## 구조

### Getters

저장소 state를 가져와 변환하는 작업을 여러곳에서 동일하게 해야한다면,
저장소 getter를 사용해서 변환한 값을 호출 가능하다.

getters 정의

```js
const store = new Vuex.Store({
  state: {
    todos: [
      { id: 1, text: '...', done: true },
      { id: 2, text: '...', done: false },
    ]
  },
  getters: {
    doneTodos: state => {
      return state.todos.filter(todo => todo.done);
    },
    doneTodosCount: (state, getters) => { // 두번 째 인자로 다른 getter를 받을 수 있음
      return getters.doneTodos.length;
    },
  }
})
```

컴포넌트에서 getters 사용

```js
  computed: {
    doneTodosCount() {
      return this.$store.getters.doneTodosCount;
    }
  }
```

#### 컴포넌트 내부에서 mapGetters 사용

mapGetters는 저장소 getter를 local computed에 매핑함

```js
import { mapGetters } from 'vuex';

export default {
  // ...
  computed: {
    // getter를 다른이름으로 매핑
    ...mapGetters({
      doneCount: 'doneTodosCount'
    }),
    // getter를 같은 이름으로 매핑해서 사용
    ...mapGetters([
      'doneTodosCount',
    ])
  },
  // ...
}
```

### Mutations

store에서 변이를 사용해서 실제로 상태를 변경할 수 있다.
변이 핸들러 함수는 동기적이어야 한다. (비동기X)

```js
const store = new Vuex.Store({
  state: {
    count: 1
  },
  mutations: {
    increment (state) {
      state.count++;
    }
  }
})
```

변이를 할 때 직접 핸들러를 호출 할수 없고 commit을 호출해야 한다.

```js
this.$store.commit('increment');
```

변이에 대해서 payload라고 하는 commit에 대한 추가 전달인자를 사용할 수 있다.

```js
// ...
mutations: {
  increment (state, payload) {
    state.count += payload.amount;
  }
}
```

commit을 호출할 때 추가전달인자도 함께 명시하면 된다.

```js
this.$store.commit('increment', {
  amount: 10
});
```

### 컴포넌트 내부에서 mapMutations 사용

```js
import { mapMutations } from 'vuex';

export default {
  // ...
  methods: {
    ...mapMutations([
      'increment' // this.$store.commit('increment')를 this.increment()에 매핑
    ]),
    // 다른이름으로 매핑
    ...mapMutations([
      add: 'increment'
    ])
    ...
  }
}
```

### Actions

액션은 변이와 유사하지만 다음과 같은 차이가 있다.

- 상태를 변이시키는 대신 액션으로 변이에 대한 커밋을 한다.
- 임의의 비동기 작업이 포함될 수 있다.

```js
const store = new Vuex.Store({
  state: {
    count: 0
  },
  mutations: {
    increment (state) {
      state.count++;
    }
  },
  actions: {
    increment (context) {
      context.commit('increment');
    }
  }
});
```

액션 핸들러는 store 인스턴스의 methods / properties 세트를 포함한 컨텍스트 객체를 받는다.
따라서 `context.commit`을 호출하여 mutations commit이나 `context.state`, `context.getters`를
통해서 상태와 getters에 접근할 수 있다.

코드를 단순화 하기 위해서 (특히 `commit`을 여러번 호출하는 경우) 전달인자 분해를 사용한다.

```js
actions: {
  increment ({ commit }) {
    commit('increment');
  }
}
```

액션 사용은 dispatch 메소드 사용

```js
this.$store.dispatch('increment`);
```

`commit`이 아니라 `dispatch`를 사용하는 것은 비동기 작업을 염두에 둔 것

```js
actions: {
  incrementAsync ({ commit, state }) { // store 인스턴스 내부 컨텍스트를 분리해 가져올 수 있다.
    setTimeout(() => {
      commit('increment');
    })
  }
}
```

`action` 역시 payload를 지원한다.

```js
this.$store.dispatch('incrementAsync', {
  amount: 10
})
```

#### 컴포넌트 내부에서 mapActions 사용

```js
import { mapActions } from 'vuex';

export default {
  //...
  methods: {
    ...mapActions([
      'increment' // this.$store.dispatch('increment')를 this.increment()에 매핑
    ]),
    // 다른이름으로 매핑
    ...mapActions({
      add: 'increment'
    })
  }
}
```

#### 비동기 액션 구성 (promise)

`store.dispatch`는 트리거 된 액션 핸들러에 의해 반환된 Promise를 처리할 수 있으며 Promise를 반환한다.

```js
actions: {
  actionA ({ commit }) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        commit('someMutation');
        resolve();
      }, 1000)
    })
  },
  // 액션 내부에서 액션 사용
  actionB ({ dispatch, commit }) {
    return dispatch('actionA').then(() => {
      commit('someOtherMutation');
    })
  }
}
```

다음과 같이 사용

```js
store.dispatch('actionA').then(() => {
  // ...
});
```

`async/await` 사용시

```js
// getData() 및 getOtherData()가 Promise를 반환할 시
actions: {
  async actionA ({ commit }) {
    commit('gotData', await getData());
  },
  async actionB ({ dispatch, commit }) {
    await dispatch('actionA'); // actionA가 끝나기를 기다린다
    commit('gotOtherData', await getOtherData());
  }
}
```

### State

#### 컴포넌트 내부에서 mapState 사용

컴포넌트가 여러 저장소 상태 속성이나 getter를 사용할 경우 반복선언을 줄여 쓰기 위해 사용

```js
import { mapState } from 'vuex';

export default {
  computed: mapState({
    count: state => state.count,
    countAlias: 'count', // state => state.count 와 같은 효과
    countPlusLocalstate (state) {
      return state.count + this.localCount;
    },
  })
}
```

this.count를 store.state.count에 매핑,
spread연산자를 사용하면 로컬의 computed 속성과 함께 사용가능

```js
// ...
  computed: {
    localComputed () { /* ... */ },
    ...mapState([
      'count'
    ])
  }
// ...
```

## 모듈화

Vuex는 저장소를 모듈로 나눌 수 있다. 각 모듈은 자체 상태, 변이, 액션, 게터 및 중첩된 모듈을 포함할 수 있다.

```js
const moduleA = {
  state: { ... },
  mutations: { ... },
  actions: { ... },
  getters: { ... }
}

const moduleB = {
  state: { ... },
  mutations: { ... },
  actions: { ... },
  getters: { ... }
}

const store = new Vuex.Store({
  modules: {
    a: moduleA,
    b: moduleB
  }
})

store.state.a; // moduleA 의 state
store.state.b; // moduleB 의 state
```

모듈의 `mutations` 와 `getters` 내부에서 첫 번째 전달인자는 **모듈의 지역상태** 가 된다.

```js
const moduleA = {
  state: {
    count: 0
  },
  mutations: {
    increment (state) {
      // state는 지역 모듈 상태
      state.count++;
    }
  },
  getters: {
    doubleCount (state) {
      return state.count * 2;
    }
  }
}
```

모듈 `getters` 내부에서 `rootState`는 세 번째 전달인자로 전해진다

```js
const mouduleA = {
  //...
  getters: {
    sumWithRootCount (state, getters, rootstate) {
      return state.count + rootState.count;
    }
  }
}
```

`actions`의 `context.state`는 지역상태이고 `context.rootState`는 전역 상태이다.

```js
// ...
actions: {
  incrementIfOddOnRootSum ({ state, commit, rootState }) {
    if ((state.count + rootState.count) % 2 === 1) {
      commit('increment');
    }
  }
}
```

## 네임스페이스

### 네임스페이스를 상수로 등록

```js
// mutation-types.js
export const SOME_MUTATION = 'SOME_MUTATION';

// store.js
import Vuex form 'vuex';
import { SOME_MUTATION } from './mutation-types';

const store = new Vuex.Store({
  state: { ... },
  mutations: {
    [SOME_MUTATION] (state) {
      /* ... */
    }
  }
});
```

### 모듈 네임스페이스

모듈내의 action, mutation, getter는 전역 네임스페이스에 등록된다.
여러 모듈 사용시 이름충돌을 피하기 위해서 접두사 또는 접미사를 붙여 모듈 자신의 네임스페이스를 지정할 수 있다.

```js
// types.js
// getter, action, mutation의 이름을 상수로 정의하고 모듈이름 `todos` 접두어를 붙인다.
export const DONE_COUNT = 'todos/DONE_COUNT';
export const FETCH_ALL = 'todos/FETCH_ALL';
export const TOGGLE_DONE = 'todos/TOGGLE_DONE';

// modules/todos.js
import * from types from '../types';

const todosModule = {
  state: {
    todos: []
  },
  getters: {
    [types.DONE_COUNT] (state) {
      // ...
    }
  },
  actions: {
    [types.FETCH_ALL] (context, payload) {
      // ...
    }
  },
  mutations: {
    [types.TOGGLE_DONE] (state, payload) {
      // ...
    }
  }
}
```

## 동적 모듈등록

`store.registerModule` 메소드로 저장소가 생상 된 후에 모듈등록 가능

```js
store.registerModule('myModule', {
  // ...
})
```

이 때 모듈 상태는 `store.state.myModule`으로 접근가능

`store.unregisterModule(moduleName)`을 사용하여 동적으로 등록된 모듈을 제거할 수 있다.
이 방법으로는 저장소 생성시 선언된 정적모듈은 제거할 수 없다.
