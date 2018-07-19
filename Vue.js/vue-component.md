# Vue.js Component

<https://vuejs.org/v2/guide/components.html>

## 컴포넌트 생성 및 사용

### 컴포넌트 정의

Vue 생성자는 미리 정의 된 옵션으로 재사용 가능한 컴포넌트 생성자를 생성하도록 확장 가능하다.

```js
var MyComponent = Vue.extend({
  // 옵션 확장
})
// `MyComponent`의 모든 인스턴스는
// 미리 정의된 확장 옵션과 함께 생성됩니다.
var myComponentInstance = new MyComponent()
```

또한 전역 컴포넌트를 등록하려면, `Vue.component(tagName, options)`를 사용할 수 있다
Global registration은 Root Vue instance(with new Vue)가 생성되기 이전에 수행된다.

```js
// Define a new component called button-counter
Vue.component('button-counter', {
  data: function () {
    return {
      count: 0
    }
  },
  template: '<button v-on:click="count++">You clicked me {{ count }} times.</button>'
})
```

```html
<div id="components-demo">
  <button-counter></button-counter>
</div>
```

```js
new Vue({ el: '#components-demo' })
```

### 컴포넌트 재사용

컴포넌트를 여러번 선언하면 독립적인 scope를 갖는 컴포넌트가 여러개 생성된다.

```html
<div id="components-demo">
  <button-counter></button-counter>
  <button-counter></button-counter>
  <button-counter></button-counter>
</div>
```

### 컴포넌트와 data

컴포넌트에서 사용되는 data는 반드시 함수여야 한다. => data를 함수의 반환 값으로 처리 (객체 생성 후 반환)

data를 함수 반환 값으로 사용하지 않으면 서로의 객체를 공유하여 고유한 내부상태를 가질 수 없다.

### 컴포넌트 작성

컴포넌트는 부모-자식 관계에서 일반적으로 사용하기 위해 작성한다.
부모 컴포넌트는 자체 템플릿에서 자식 컴포넌트를 사용할 수 있으며, 서로 메시지를 주고 받아야 할 필요가 있다.

부모 컴포넌트는 자식에게 `props`로 데이터를 전달하고, 자식 컴포넌트는 `events`를 통해 부모에게 메시지를 보낸다.

### A single root element

컴포넌트의 root element는 하나여야 한다.

```html
<!-- root element가 두개인 경우 오류가 발생한다 -->
<h3>{{ title }}</h3>
<div v-html="content"></div>

<!-- 다음과 같이 구성해야 한다 -->
<div class="blog-post">
  <h3>{{ title }}</h3>
  <div v-html="content"></div>
</div>
```

### 컴포넌트 이름 규약

컴포넌트 이름을 정의할 때 두 가지 방식이 있다

kebab-case 사용

```js
Vue.component('my-component-name', { /* ... */ })
```

kebab-case로 이름을 지을 때, 참조하는 custom element역시 kebab-case로 사용해야 한다: `<my-component-name>`

PascalCase 사용

```js
Vue.component('MyComponentName', { /* ... */ })
```

PascalCase로 이름을 지으면, 참조하는 custom element를 사용할 때 kebab-case 및 PascalCase 모두 사용할 수 있다.
즉, `<my-component-name>`, `<MyComponentName>` 두 방법이 가능하다. 그러나 DOM에서 직접 접근할 때는 kebab-case 이름만 유효하다.

### 지역(local) 컴포넌트 등록

전역 등록은 많은경우 이상적/효율적이지 않다.
이런경우 컴포넌트를 JavaScript 객체로 선언하고 지역적으로 사용할 수 있다.

```js
var ComponentA = { /* ... */ }
var ComponentB = { /* ... */ }
var ComponentC = { /* ... */ }

// 컴포넌트를 Vue 인스턴스의 components option에 선언한다

new Vue({
  el: '#app'
  components: {
    'component-a': ComponentA,
    'component-b': ComponentB
  }
})
```

Key는 custom element의 이름이 되고 value는 component object이다.

로컬 컴포넌트는 하위 컴포넌트에서도 가능하다. 컴포넌트 A를 컴포넌트B의 subcomponent로 선언할 수 있다.

```js
var ComponentA = { /* ... */ }

var ComponentB = {
  components: {
    'component-a': ComponentA
  },
  // ...
}
```

컴포넌트 선언에서 key-value 없이 `ComponentA`만 있다면 `ComponentA: ComponentA`을 줄여쓴 것과 동일하다

### 모듈 시스템

#### 모듈 시스템에서 컴포넌트 로컬 등록

Babel 이나 Webpack을 이용한 모듈 시스템에서는 컴포넌트 디렉토리를 만들어 각 컴포넌트를 하나의 파일로 관리하기를 권장한다.

컴포넌트를 사용하려면 local 등록할 컴포넌트를 import하여 사용할 수 있다.

```js
import ComponentA from './ComponentA'
import ComponentC from './ComponentC'

export default {
  components: {
    ComponentA,
    ComponentC
  },
  // ...
}
```

위의 코드에 따르면, ComponentA 와 ComponentC는 ComponentB의 template에서 사용될 수 있다.

#### Base Components를 자동으로 전역 등록

여러 컴포넌트에서 자주/공통적으로 사용되는 컴포넌트를 등록해야 할 필요가 있다.
결과적으로 많은 컴포넌트 리스트가 base components에 나열될 것이다

```js
import BaseButton from './BaseButton.vue'
import BaseIcon from './BaseIcon.vue'
import BaseInput from './BaseInput.vue'

export default {
  components: {
    BaseButton,
    BaseIcon,
    BaseInput
  }
}
```

상대적으로 템플릿에서는 적은 markup을 작성한다

```html
<BaseInput
  v-model="searchText"
  @keydown.enter="search" />
<BaseButton @click="search">
  <BaseIcon name="search" />
</BaseButton>
```

만약 Webpack을 (Vue CLI 3+, 내부적으로 Webpack이 작동함) 사용한다면 예제 형식의 common base components를 작성하면 된다.
다음은 base components를 globally import하는 예제이다: `src/main.js`

```js
import Vue from 'vue'
import upperFirst from 'lodash/upperFirst'
import camelCase from 'lodash/camelCase'

const requireComponent = require.context(
  // The relative path of the components folder
  './components',
  // Whether or not to look in subfolders
  false,
  // The regular expression used to match base component filenames
  /Base[A-Z]\w+\.(vue|js)$/
)

requireComponent.keys().forEach(fileName => {
  // Get component config
  const componentConfig = requireComponent(fileName)

  // Get PascalCase name of component
  const componentName = upperFirst(
    camelCase(
      // Strip the leading `./` and extension from the filename
      fileName.replace(/^\.\/(.*)\.\w+$/, '$1')
    )
  )

  // Register component globally
  Vue.component(
    componentName,
    // Look for the component options on `.default`, which will
    // exist if the component was exported with `export default`,
    // otherwise fall back to module's root.
    componentConfig.default || componentConfig
  )
})
```

## Props

### Props 사용

모든 컴포넌트 인스턴스에는 scope가 있으므로, 하위 컴포넌트에서 상위 컴포넌트 데이터를 직접 참조 할 수 없다.
prop는 상위 컴포넌트의 데이터를 하위 컴포넌트로 전달하기위한 사용자 지정 특성이다.

상위 컴포넌트에서는 하위 컴포넌트 사용 선언시 props를 정의하고,
하위 컴포넌트에서는 props 옵션을 사용하여 상위 컴포넌트로 부터 받을 props를 선언해야 한다.

```js
Vue.component('child', {
  // props 정의
  props: ['message'],
  // vm의 this.message로 사용가능
  template: '<span>{{ message }}</span>'
})
```

상위 컴포넌트에서는 다음과 같이 props를 전달해야 한다.

```js
<child message="안녕하세요!"></child>
```

### Prop Casing (camelCase vs kebab-case)

HTML 속성은 대소문자 구분을 하지 않으므로 Prop을 사용할 경우 camelCased 이름은 kebab-cased로 바꿔 사용한다

```js
Vue.component('blog-post', {
  // camelCase in JavaScript
  props: ['postTitle'],
  template: '<h3>{{ postTitle }}</h3>'
})
```

```html
<!-- kebab-case in HTML -->
<blog-post post-title="hello!"></blog-post>
```

만약 string templates을 사용한다면(vue instance에 template 기재시) casing은 적용되지 않는다.

### Prop Types & Validation

prop를 받는 곳에서 prop type및 요구사항을 지정할 수 있다.
문자열 배열을 통해 여러 type을 정의하거나, validator를 사용할 수 있다.

```js
Vue.component('example', {
  props: {
    // 기본 타입 확인 (`null` 은 어떤 타입이든 가능
    propA: Number,
    // 여러 타입이 가능하면 배열로 지정
    propB: [String, Number],
    // 문자열이며 not null
    propC: {
      type: String,
      required: true
    },
    // 숫자이며 입력이 없다면 기본 값 100
    propD: {
      type: Number,
      default: 100
    },
    // 객체/배열의 기본값은 함수 반환값으로 정의
    propE: {
      type: Object,
      default: function () {
        return { message: 'hello' }
      }
    },
    // 사용자 정의 유효성 검사 가능
    propF: {
      validator: function (value) {
        return value > 10
      }
    }
  }
})
```

type은 다음 네이티브 생성자 중 하나를 사용할 수 있다

- `String`
- `Number`
- `Boolean`
- `Function`
- `Object`
- `Array`
- `Symbol`

또한, `type` 은 커스텀 생성자 함수가 될 수 있고, `assertion`은 `instanceof` 로 이루어짐

props 검증이 실패하면 Vue는 콘솔에서 경고를 출력한다(개발 빌드를 사용하는 경우).

> props는 컴포넌트 인스턴스가 생성되기 전에 검증되기 때문에, default 또는 validator 함수 내에서 data, computed 또는 methods와 같은 인스턴스 속성을 사용할 수 없다.

### 동적 Props

v-bind를 사용하여 부모의 데이터와 props를 동적으로 바인딩 할 수 있다.
상위 컴포넌트에서 데이터가 업데이트 될 때마다 하위 컴포넌트로 전달된다.

```html
<div>
  <input v-model="parentMsg">
  <br>
  <child v-bind:my-message="parentMsg"></child>
</div>
```

단축구문(`:`)을 사용할 수도 있다: `<child :my-message="parentMsg"></child>`

객체의 모든 속성을 prop로 전달하려면, 인자없이 bind를 사용한다. (v-bind:prop-name 대신 v-bind)

```js
todo: {
  text: 'Learn Vue',
  isComplete: false
}
```

부모 템플릿에서 prop와 바인딩한다

```html
<todo-item v-bind="todo"></todo-item>
```

위의 구문은 아래와 같은 기능을 수행한다

```html
<todo-item
  v-bind:text="todo.text"
  v-bind:is-complete="todo.isComplete">
</todo-item>
```

### 리터럴 vs. 동적

```html
<!-- 이것은 일반 문자열 "1"을 전달 -->
<comp some-prop="1"></comp>
```

리터럴 prop는 문자열로 데이터를 전달한다.
숫자를 전달하려면 값이 JavaScript 표현식으로 평가되도록 v-bind를 사용해야한다.

```html
<!-- 이것은 실제 숫자로 전달 -->
<comp v-bind:some-prop="1"></comp>
```

### 단방향 데이터 흐름

모든 prop는 하위 속성과 상위 속성 사이의 단방향 바인딩을 형성한다.
상위에서 하위로 데이터 변경은 전달 되지만, 반대는 가능하지 않다.
이는 하위 컴포넌트가 부모의 상태를 변경하여 앱의 데이터 흐름을 추론하기 어렵게 만드는 것을 방지하기 위함이다.

#### prop을 직접 변경 하지 않고 사용하기 위한 방법

prop는 초기 값을 전달 하는데만 사용되며 하위 컴포넌트는 이후에 이를 로컬 데이터 속성으로 사용하기만 하는 경우

```js
// case1: prop의 초기 값을 초기 값으로 사용하는 로컬 데이터 속성을 정의
props: ['initialCounter'],
data: function () {
  return { counter: this.initialCounter }
}

// case2: prop 값으로 부터 계산된 속성을 정의
props: ['size'],
computed: {
  normalizedSize: function () {
    return this.size.trim().toLowerCase()
  }
}
```

자바 스크립트의 객체와 배열은 참조로 전달되므로,
prop가 배열이나 객체인 경우 하위 객체 또는 배열 자체를 부모 상태로 변경하면 부모 상태에 영향을 준다.

## Props가 아닌 속성

명시적인 props는 하위 컴포넌트에 정보를 전달하는데 적절하지만,
컴포넌트를 라이브러리로 만드는 경우 컴포넌트가 사용되는 상황을 항상 예측할 수는 없다.

따라서 컴포넌트의 루트 요소에 추가되는 임의의 속성을 허용해야 한다.

예를 들어, data-3d-date-picker 속성을 요구하는 bs-date-input 컴포넌트를 사용하고 있다면, 속성을 컴포넌트 인스턴스에 추가 할 수 있다.

`<bs-date-input data-3d-date-picker="true"></bs-date-input>`

`data-3d-date-picker="true"`속성은 `bs-date-input`의 루트 엘리먼트에 자동으로 추가 될 것이다.

### 존재하는 속성 교체/병합

`bs-date-input`의 템플릿이 다음과 같다면

```html
<input type="date" class="form-control">
```

date picker 플러그인을 변경하려면 다음과 같이 특정 클래스를 변경해야 한다.

```html
<bs-date-input data-3d-date-picker="true" class="date-picker-theme-dark"></bs-date-input>
```

이 경우 class에 대한 두 개의 서로 다른 값이 충돌할 수 있다.

- 템플릿의 컴포넌트에 의해 설정된 `form-control`
- 부모 컴포넌트에서 전달된 `date-picker-theme-dark`

대부분의 속성의 경우 기존 값을 새로운 값이 대체된다.

다른 속성과 다르게 class와 style 속성은 두 값이 합쳐져서 최종 값인 `form-control date-picker-theme-dark` 이 된다.

### 속성(attribute) 상속 해제

컴포넌트의 Root element에서 속성을 상속하기를 원하지 않는다면 옵션에서 설정할 수 있다.

```js
Vue.component('my-component', {
  inheritAttrs: false,
  // ...
})
```

위 옵션은 속성 이름과 값이 컴포넌트로 전달되는 `$attrs` instance property와 사용하기 좋다

```js
Vue.component('base-input', {
  inheritAttrs: false,
  props: ['label', 'value'],
  template: `
    <label>
      {{ label }}
      <input v-bind="$attrs"
        v-bind:value="value"
        v-on:input="$emit('input', $event.target.value)">
    </label>
  `
})
```

```html
<base-input
  v-model="username"
  class="username-input"
  placeholder="Enter your username">
</base-input>
```

위의 패턴은 Base components를 원래 존재하는 HTML element처럼 사용가능하게 한다

## v-on을 이용한 사용자 지정 이벤트

모든 Vue 인스턴스는 다음과 같은 이벤트 인터페이스를 구현한다.

- `$on(eventName)`을 사용하여 이벤트를 감지한다
- `$emit(eventName)`을 사용하여 이벤트를 발생시킨다

Vue의 이벤트 시스템은 브라우저의 `EventTarget API`와 별개로, `$on` 과 `$emit` 는 `addEventListener` 와 `dispatchEvent`의 별칭이 아니다.

- 부모 컴포넌트는 자식 컴포넌트를 템플릿에서 직접 `v-on` 을 사용하여 자식 컴포넌트에서 보내진 이벤트를 청취할 수 있다
- `$on`은 자식에서 호출한 이벤트는 감지 하지 않으므로, `v-on`을 **반드시 템플릿에 지정해야 한다**

```html
<div id="counter-event-example">
  <p>{{ total }}</p>
  <button-counter v-on:increment="incrementTotal"></button-counter>
  <button-counter v-on:increment="incrementTotal"></button-counter>
</div>
```

하위 컴포넌트는 외부와 완전히 분리 된다는 점에 유의해야 한다.

```js
// 부모 컴포넌트
new Vue({
  el: '#counter-event-example',
  data: {
    total: 0
  },
  methods: {
    incrementTotal: function () {
      this.total += 1
    }
  }
})

// 자식 컴포넌트
Vue.component('button-counter', {
  template: '<button v-on:click="incrementCounter">{{ counter }}</button>',
  data: function () {
    return {
      counter: 0
    }
  },
  methods: {
    incrementCounter: function () {
      this.counter += 1
      this.$emit('increment')
    }
  },
})
```

이벤트를 발생시킬 때 인자도 함께 전달할 수 있다

```html
<button v-on:click="$emit('enlarge-text', 0.1)">
  Enlarge text
</button>
```

부모 컴포넌트에서 인자를 `$event` 객체로 받아 inline 사용할 수 있다.

```html
<blog-post
  ...
  v-on:enlarge-text="postFontSize += $event">
</blog-post>
```

이벤트 핸들러가 메소드라면, 파라미터로 해당 인자를 받아 사용할 수 있다.

```html
<blog-post
  ...
  v-on:enlarge-text="onEnlargeText">
</blog-post>
```

```js
methods: {
  onEnlargeText: function (enlargeAmount) {
    this.postFontSize += enlargeAmount
  }
}
```

### 이벤트 이름

components나 props와 다르게 이벤트 이름은 casing을 지원하지 않는다.
대신 emitted event는 event listener에서 확인 하는 값고 정확히 동일하게 작성되어야 한다.

```js
this.$emit('myEvent')
```

camelCase로 발생시킨 이벤트는 kebab-cased 버전의 리스너에서 작동하지 않는다.

```html
<my-component v-on:my-event="doSomething"></my-component>
```

이벤트 이름은 JavaScript 변수나 프로퍼티로 사용되지 않으므로 camelCase나 PascalCase로 작성될 필요가 없다.
게다가 v-on event listeners는 DOM templates 내에 작성되어 자동으로 소문자변환이 된다.
따라서 **이벤트 이름은 항상 kebab-case로 작성**하는 것이 좋다.

### 사용자 정의 이벤트를 사용하여 폼 입력 컴포넌트 만들기

사용자 정의 이벤트는 v-model 에서 작동하는 사용자 정의 입력을 만드는데에도 사용할 수 있다

```html
<input v-model="searchText">
```

위 문장은 아래와 같다.

```html
<input
  v-bind:value="searchText"
  v-on:input="searchText = $event.target.value">

<!-- 컴포넌트와 함께 사용하면 다음과 같이 간단해진다 -->
<custom-input
  v-bind:value="searchText"
  v-on:input="searchText = $event">
</custom-input>
```

따라서 v-model을 사용하는 컴포넌트는

- value 속성과 value prop을 binding이 필요하다
- 입력이 이루어지면 정의한 input 이벤트가 발생한다

간단한 예제를 살펴보자

```html
<currency-input v-model="price"></currency-input>
```

```js
Vue.component('currency-input', {
  template: '\
    <span>\
      $\
      <input\
        ref="input"\
        v-bind:value="value"\
        v-on:input="updateValue($event.target.value)">\
    </span>\
  ',
  props: ['value'],
  methods: {
    // 값을 직접 업데이트하는 대신 메소드를 통해 입력값을 가공할 수 있다
    updateValue: function (value) {
      var formattedValue = value
        .trim()
        .slice(0,
          value.indexOf('.') === -1 ? value.length : value.indexOf('.') + 3)
      if (formattedValue !== value) {
        this.$refs.input.value = formattedValue
      }
      // 입력 이벤트를 통해 숫자 값을 내보낸다
      this.$emit('input', Number(formattedValue))
    }
  }
})
```

### 컴포넌트의 v-model 사용자 정의

기본적으로 컴포넌트의 v-model은 value를 보조 변수로 사용하고 input을 이벤트로 사용하지만
체크 박스와 라디오 버튼과 같은 일부 입력 타입은 다른 목적으로 value 속성을 사용할 수 있다.

model 옵션을 사용하여 다음 경우에 충돌을 피할 수 있다.

```js
Vue.component('base-checkbox', {
  model: {
    prop: 'checked',
    event: 'change'
  },
  props: {
    checked: Boolean
  },
  template: `
    <input
      type="checkbox"
      v-bind:checked="checked"
      v-on:change="$emit('change', $event.target.checked)">
  `
})
```

```html
<base-checkbox v-model="lovingVue"></base-checkbox>
```

`lovingVue`의 값은 `checked` prop으로 전달될 것이다
`<base-checkbox>`가 새로운 값에 의해 change event를 발생시키면 `lovingVue` prop은 업데이트 된다.

### 컴포넌트에 네이티브 이벤트 바인딩

컴포넌트의 **루트 엘리먼트**에서 네이티브 이벤트를 수신하려는 경우 `v-on:{event}.native` 수식자를 사용한다.

`<my-component v-on:click.native="doTheThing"></my-component>`

아래의 `<base-input>` 컴포넌트의 root element는 `<label>` 이다.
이런경우 부모 컴포넌트에서 `.native` 리스너는 에러없이 작동하지 않는다.

```html
<label>
  {{ label }}
  <input
    v-bind="$attrs"
    v-bind:value="value"
    v-on:input="$emit('input', $event.target.value)"
  >
</label>
```

이러한 문제가 발생했을 때 컴포넌트에서 사용된 리스너 객체를 포함하고 있는 `$listeners` property를 사용할 수 있다.
`v-on="$listeners"`를 사용하여 컴포넌트에 포함된 특정한 자식 요소까지 포함한 이벤트 리스너를 보낼 수 있다.

```json
{
  focus: function (event) { /* ... */ }
  input: function (value) { /* ... */ },
}
```

`<input>` 같이 v-model과 함께 사용하고 싶은 요소의 경우,
아래의 `inputListeners()`와 같은 리스너를 위한 새로운 computed property를 작성할 수 있다.

```js
Vue.component('base-input', {
  inheritAttrs: false,
  props: ['label', 'value'],
  computed: {
    inputListeners() {
      var vm = this
      // 비어있는 객체에 인수 객체들을 합쳐서 새 객체를 할당한다
      return Object.assign({},
        // 부모 객체의 모든 리스너를 불러와 합친다
        this.$listeners,
        // 특정 리스너를 override 할 사용자 정의 리스너를 추가한다
        {
          // 컴포넌트가 v-model과 작동하게 한다
          input: (event) => {
            vm.$emit('input', event.target.value)
          }
        }
      )
    }
  },
  template: `
    <label>
      {{ label }}
      <input
        v-bind="$attrs"
        v-bind:value="value"
        v-on="inputListeners">
    </label>
  `
})
```

`<base-input>` component는 평범한 `<input>` element 처럼 사용될 수 있다.
`.native` 수정자 없이 모든 attributes 와 listeners가 동작할 것이다.

### 비 부모-자식간 통신

부모-자식간이 아닌 컴포넌트 사이에서 통신할 필요가 있을 때가 있다.
이 때는 비어있는 Vue 인스턴스를 중앙 이벤트 버스로 사용할 수 있습니다.

```js
var bus = new Vue()

// 컴포넌트 A의 메소드
bus.$emit('id-selected', 1)

// 컴포넌트 B의 created 훅
bus.$on('id-selected', function (id) {
  // ...
})
```

복잡한 경우에는 Vuex 같은 전용 상태 관리 라이브러리를 사용하는 것이 좋다.

### `.sync` 수식어

일부 경우에 속성에 “양방향 바인딩”이 필요할 수 있다. Vue1 버전에 있던 `.sync` 수식어와 동일하다.
자식 컴포넌트가 `.sync`를 가지는 속성을 변경하면 부모의 값에도 반영된다.

하지만 단방향 데이터 흐름이 아니기 때문에 장기적으로 유지보수에 문제가 생긴다.
부모 상태에 영향을 미치므로, `.sync`를 사용할 때는 코드를 더욱 일관성있고 명백하게 만들어야합니다.

`.sync`는 `v-on`에서 자동으로 확장되는 syntax sugar이다.

`<comp :foo.sync="bar"></comp>`

위의 코드는 아래와 같다

`<comp :foo="bar" @update:foo="val => bar = val"></comp>`

하위 컴포넌트가 foo를 갱신하려면 명시적으로 이벤트를 보내야한다: `this.$emit('update:foo', newValue)`

## 슬롯을 사용한 콘텐츠 배포

컴포넌트를 사용할 때 다음과 같이 컴포넌트를 구성하는 것이 좋다.

```html
<app>
  <app-header></app-header>
  <app-footer></app-footer>
</app>
```

- `<app>` 컴포넌트는 `<app>`이 사용하는 컴포넌트에 의해 어떤 컨텐츠를 받을지 결정된다
- `<app>` 컴포넌트에는 자체 템플릿이 있을 가능성이 크다

위의 경우 부모 “content”와 컴포넌트의 자체 템플릿을 섞는 방법이 필요하다.
Vue.js는 원본 콘텐츠의 배포판 역할을하기 위해 특수한 `<slot>` 엘리먼트를 사용한다.

### 범위 컴파일

```html
<child-component>
  {{ message }}
</child-component>
```

message는 부모 데이터에 바인딩 되어야 한다.

> 상위 템플릿의 모든 내용은 상위 범위로 컴파일된다. 하위 템플릿의 모든 내용은 하위 범위에서 컴파일된다. 분산된 콘텐츠는 상위 범위에서 컴파일된다.

### 단일 슬롯

하위 컴포넌트 템플릿에 최소한 하나의 `<slot>` 콘텐츠가 포함되어 있지 않으면 부모 콘텐츠가 삭제된다.
속성이 없는 슬롯이 하나 뿐인 경우 전체 내용 조각이 DOM의 해당 위치에 삽입되어 슬롯 자체를 대체한다.

`<slot>` 태그 안에 있는 내용은 대체 콘텐츠로 간주되어,
하위 범위에서 컴파일되며 부모 컴포넌트의 자식 호출 엘리먼트 내부가 비어 있고 삽입할 콘텐츠가 없는 경우에만 표시된다.

```html
<div>
  <h2>나는 자식 컴포넌트의 제목입니다</h2>
  <slot>
    제공된 컨텐츠가 없는 경우에만 보실 수 있습니다.
  </slot>
</div>

<!-- 컴포넌트를 사용하는 부모 -->
<div>
  <h1>나는 부모 컴포넌트의 제목입니다</h1>
  <my-component>
    <p>이것은 원본 컨텐츠 입니다.</p>
    <p>이것은 원본 중 추가 컨텐츠 입니다</p>
  </my-component>
</div>

<!-- 부모 컴포넌트는 아래처럼 렌더링 됨 -->
<div>
  <h1>나는 부모 컴포넌트의 제목입니다</h1>
  <div>
    <h2>나는 자식 컴포넌트의 제목 입니다</h2>
    <p>이것은 원본 컨텐츠 입니다.</p>
    <p>이것은 원본 중 추가 컨텐츠 입니다</p>
  </div>
</div>
```

### 이름을 가지는 슬롯

`<slot>` 엘리먼트는 특별한 속성인 `name` 을 가지고 있다. 이름이 다른 슬롯이 여러 개 있을 수 있다.

명명되지 않은 기본 슬롯은 일치하지 않는 콘텐츠의 대리자 역할을 한다. 기본 슬롯이 없으면 일치하지 않는 콘텐츠가 삭제됩니다.

```html
<div class="container">
  <header>
    <slot name="header"></slot>
  </header>
  <main>
    <slot></slot>
  </main>
  <footer>
    <slot name="footer"></slot>
  </footer>
</div>

<!-- 부모 컴포넌트 -->
<app-layout>
  <h1 slot="header">여기에 페이지 제목이 위치합니다</h1>
  <p>메인 컨텐츠의 단락입니다.</p>
  <p>하나 더 있습니다.</p>
  <p slot="footer">여기에 연락처 정보입니다.</p>
</app-layout>

<!-- 아래와 같이 렌더링 된다 -->
<div class="container">
  <header>
    <h1>여기에 페이지 제목이 위치합니다</h1>
  </header>
  <main>
    <p>메인 컨텐츠의 단락입니다.</p>
    <p>하나 더 있습니다.</p>
  </main>
  <footer>
    <p>여기에 연락처 정보입니다.</p>
  </footer>
</div>
```

### 범위를 가지는 슬롯

범위가 지정된 슬롯은 이미 렌더링 된 엘리먼트가 아닌,
데이터를 전달할 수 있는 재사용 가능한 템플릿으로 작동하는 특별한 유형의 슬롯이다.

prop을 컴포넌트에게 전달하는 것처럼, 하위 컴포넌트에서 데이터를 슬롯에 전달하면 된다.

```html
<div class="child">
  <slot text="hello from child"></slot>
</div>
```

부모 컴포넌트에서는, slot-scope 속성을 가진 `<template>` 엘리먼트가 있어야 한다.
slot-scope의 값은 자식으로부터 전달 된 props 객체를 담고있는 임시 변수의 이름이다.

```html
<div class="parent">
  <child>
    <template slot-scope="props">
      <span>hello from parent</span>
      <span>{{ props.text }}</span>
    </template>
  </child>
</div>

<!-- 위를 렌더링하면 다음과 같이 출력된다 -->
<div class="parent">
  <div class="child">
    <span>hello from parent</span>
    <span>hello from child</span>
  </div>
</div>
```

> slot-scope 는 `<template>` 뿐 아니라 컴포넌트나 엘리먼트에서도 사용할 수 있다

slot-scope 일반적인 사용 사례는 사용자 정의 list component 이다.

```html
<my-awesome-list :items="items">
  <!-- scoped slot도 이름을 가질 수 있다 -->
  <li slot="item" slot-scope="props" class="my-fancy-item">
    {{ props.text }}
  </li>
</my-awesome-list>

<!-- 리스트 컴포넌트의 템플릿 -->
<ul>
  <slot name="item" v-for="item in items" :text="item.text">
    <!-- 내용 -->
  </slot>
</ul>
```

### 구조분해

`slot-scope` 값은 함수의 인수 위치에 사용되는 유효한 JavaScript 표현식이다.

```html
<child>
  <span slot-scope="{ text }">{{ text }}</span>
</child>
```

## 동적 컴포넌트 / 비동기 컴포넌트

같은 마운트 포인트에서 `<component>`의 `is` 속성에 동적으로 바인드 하여 동적 컴포넌트 전환이 가능하다.

```js
var vm = new Vue({
  el: '#example',
  data: {
    currentView: 'home'
  },
  components: {
    home: { /* ... */ },
    posts: { /* ... */ },
    archive: { /* ... */ }
  }
})
```

```html
<component v-bind:is="currentView">
  <!-- vm.currentView가 변경되면 컴포넌트가 변경된다 -->
</component>
```

### keep-alive

트랜지션된 컴포넌트를 메모리에 유지하여 상태를 보존하거나 다시 렌더링하지 않도록하려면
동적 컴포넌트를 `<keep-alive>` 엘리먼트로 래핑하면 된다.

```html
<keep-alive>
  <component :is="currentView">
    <!-- 비활성화 된 컴포넌트는 캐시됨 -->
  </component>
</keep-alive>
```

> `<keep-alive>`는 컴포넌트의 name 옵션을 통해서 또는 local/global registration을 통해서, 이름을 가진 컴포넌트만 적용 가능하다.

### 비동기 컴포넌트

Vue에서는 비동기식 팩토리 함수로 컴포넌트를 정의 할 수 있다.
컴포넌트가 실제로 렌더링되어야 할 때만 팩토리 기능을 트리거하고 이후의 리렌더링을 위해 결과를 캐시한다.

```js
Vue.component('async-example', function (resolve, reject) {
  setTimeout(function () {
    // 컴포넌트 정의를 resolve 콜백에 전달함
    resolve({
      template: '<div>I am async!</div>'
    })
  }, 1000)
})
```

팩토리 함수는 resolve 콜백을 받아서 서버에서 컴포넌트 정의를 가져 왔을 때 호출된다.
권장되는 접근법 중 하나는 Webpack의 코드 분할 기능과 함께 비동기 컴포넌트를 사용하는 것이다.

```js
Vue.component('async-webpack-example', function (resolve) {
  // 이 특별한 require 구문은 Webpack이 Ajax 요청을 통해 로드되는 번들로 작성된 코드를 자동으로 분리하도록 지시한다.
  require(['./my-async-component'], resolve)
})

// factory 함수에서 Promise를 반환한다
Vue.component(
  'async-webpack-example',
  // `import` 함수는 `Promise`를 반환함
  () => import('./my-async-component')
)

// local 컴포넌트를 사용하는 경우, Promise를 반환하는 함수로 선언할 수 있다
new Vue({
  // ...
  components: {
    'my-component': () => import('./my-async-component')
  }
})
```

### 고급 비동기 컴포넌트

2.3 버전부터 비동기 컴포넌트 팩토리는 다음 형태의 객체를 반환할 수 있다

```js
const AsyncComp = () => ({
  // 불러올 컴포넌트이다. 반드시 Promise여야 한다
  component: import('./MyComp.vue'),
  // 컴포넌트를 비동기로 불러오는 동안 사용될 컴포넌트
  loading: LoadingComp,
  // 실패했을 경우 사용하는 컴포넌트
  error: ErrorComp,
  // 로딩 컴포넌트를 보여주기전 지연하는 정도. 기본값: 200ms.
  delay: 200,
  // timeout 된다면 error 컴포넌트가 보여진다: 기본값: Infinity
  timeout: 3000
})
```

라우트 컴포넌트에서 위의 문법을 사용하려면 vue-router 2.4.0+ 버전을 사용해야 한다

## Edge Case 다루기

다음 상황은 드물게 발생하지만 손해를 발생시키거나 위험한 오류일 수도 있다.
이러한 문제들은 Vue.js의 rule을 약간 우회해서 해결해야 할 수도 있다.

### 엘리먼트와 컴포넌트 접근

많은경우 다른 컴포넌트 인스턴스에 접근하거나 DOM 엘리먼트를 직접 조작하는것을 피하는게 좋다.
그러나 이런 행동이 적합한 경우도 있다.

#### Root Instance에 접근하기

Vue 인스턴스의 모든 subcomponent는 `$root` 프로퍼티로 root instance에 접근할 수 있다.

```js
// The root Vue instance
new Vue({
  data: {
    foo: 1
  },
  computed: {
    bar: function () { /* ... */ }
  },
  methods: {
    baz: function () { /* ... */ }
  }
})
```

모든 서브-컴포넌트들은 마치 글로벌 상태저장소에 접근하듯 루트 인스턴트에 접근할 수 있다.

```js
// Get root data
this.$root.foo

// Set root data
this.$root.foo = 2

// Access root computed properties
this.$root.bar

// Call root methods
this.$root.baz()
```

많지 않은 컴포넌트로 이루어진 작은 크기의 앱에서 이 방법은 편리하지만, 이 패턴은 중규모 이상의 App에서는 적절하지 않다.
규모가 큰 App에서는 Vuex와 같은 상태 관리 라이브러리를 쓰는 것이 좋다

#### 부모 컴포넌트의 인스턴스에 접근하기

`$root` 프로퍼티와 비슷하게 `$parent` 자식 인스턴스에서 부모 인스턴스에 접근할 때 사용될 수 있다.
이 방식은 prop으로 데이터를 보내는 방식의 게으른 대체제가 될 수 있다.

많은 경우 부모 인스턴스에 직접 접근하는 것은 구조 이해와 디버깅에 어려움을 불러온다.
특히 부모 인스턴스에 가변 데이터가 있을때 변이가 어디에서 시작했는지 찾기 어렵기 때문에 더욱 그러하다.

그러나 이런 패턴도 적절하게 쓰일 수 있는 경우가 있다.
예를 들어 HTML로 렌더링 되는 것이아닌 JavaScript API와 상호작용만 하는 추상적인 컴포넌트의 경우 그러하다.

```html
<google-map>
  <google-map-markers v-bind:places="iceCreamShops"></google-map-markers>
</google-map>
```

`<google-map>` 컴포넌트는 모든 서브-컴포넌트가 접근해야 하는 지도 프로퍼티로 정의될 수 있다.
이런 경우  `<google-map-markers>`는 마커를 설정하기 위해서 `$parent.getMap`과 같은 방법으로 부모 컴포넌트에 접근하고 싶을 수 있다.

그러나 이러한 패턴으로 만들어진 컴포넌트들은 본질적으로 취약한 구조를 지닌다.
예를 들어 `<google-map-region>` 컴포넌트를 추가할 때 `<google-map-markers>` 내부에 넣어 보이게 할 수 있다.

```html
<google-map>
  <google-map-region v-bind:shape="cityBoundaries">
    <google-map-markers v-bind:places="iceCreamShops"></google-map-markers>
  </google-map-region>
</google-map>
```

그러면 내부의 `<google-map-markers>`는 아래와 같이 부모 컴포넌트를 찾아야 한다.

`var map = this.$parent.map || this.$parent.$parent.map`

이런 경우 dependency injection을 사용하는 것이 권장된다

#### 자식 컴포넌트 인스턴스와 엘리먼트에 접근하기

props와 event의 존재에도 불구하고, JavaScript를 통해 자식 컴포넌트에 직접 접근하고 싶을 때가 있다.
직접 접근하기 위해서 자식 컴포넌트에 ref 속성을 사용해서 참조id를 할당하려 할 것이다.

```html
<base-input ref="usernameInput"></base-input>
```

`<base-input>` 인스턴스에 접근하기 위해서 참조를 정의한 컴포넌트를 다음과 같이 사용할 수 있다: `this.$refs.usernameInput`

이런방법은 부모에서 자식의 `input`에 focus가 필요한 경우 유용할 수 있다.
이런 경우 `<base-input>` 컴포넌트는 내부에 구체적인 엘리먼트에 접근을 제공하기 위해 비슷한 방법으로 `ref`를 사용한다.

```html
<input ref="input">
```

부모 컴포넌트가 사용할 메소드를 다음과 같이 작성할 수 있다

```js
methods: {
  // Used to focus the input from the parent
  focus: function () {
    this.$refs.input.focus()
  }
}
```

다음 방법으로 부모 컴포넌트는 `<base-input>` 내부의 `input`에 focus를 작동시킬 수 있다.

```js
this.$refs.usernameInput.focus()
```

`ref`가 v-for와 함께 사용되면, 자식 컴포넌트들의 참조를 포함하는 배열을 반환한다.

`$refs`는 컴포넌트가 렌더링 된후 존재하며, 반응형이 아니다.
따라서 템플릿이나 계산된 속성에서 `$refs`의 사용을 피해야 한다.

#### Dependency Injection

앞에서 부모 컴포넌트 인스턴스에 다음과 같은 방식으로 접근하였다.

```html
<google-map>
  <google-map-region v-bind:shape="cityBoundaries">
    <google-map-markers v-bind:places="iceCreamShops"></google-map-markers>
  </google-map-region>
</google-map>
```

이 컴포넌트에서 모든 `<google-map>`의 자식 컴포넌트는 어떤 지도와 상호작용 하는지 알기 위해 `getMap` 메소드에 접근해야 한다.
불행히도, `$parent` 속성을 사용해서는 중첩 단계가 커질때 제대로 확장되지 않는다.
이럴때 `provide`와 `inject`라는 두 새로운 인스턴스 옵션을 사용하는 **dependency injection**이 유용하다.

`provide` 옵션은 자식 컴포넌트에게 제공할 데이터/메소드를 명시할 수 있게 한다.
이 경우 `<google-map>` 내부의 `getMap` 메소드는

```js
provide: function () {
  return {
    getMap: this.getMap
  }
}
```

그러고 나서 아무 자식 컴포넌트에서나 `inject` 옵션으로 인스턴스에 추가하고 싶은 특정한 프로퍼티를 받을 수 있다.

```js
inject: ['getMap']
```

`$parent`를 사용하지 않고, DI를 사용할 때의 이점은 `<google-map>`의 전체 인스턴스를 노출하지 않고,
어떠한 자식 컴포넌트에서도 `getMap` 메소드에 접근할 수 있다는 것이다.
이는 자식 컴포넌트가 의존하는 것을 변경하거나 삭제할 수 있다는 두려움 없이, 보다 안전하게 컴포넌트를 개발할수 있게 해준다
컴포넌트간의 인터페이스는 props를 사용하는 것 처럼 명확히 정의된다.

사실 dependency injection을 먼 거리의 props 처럼 여길지도 모른다.

- 조상 컴포넌트는 어떤 자손 컴포넌트가 자신이 제공하는 프로퍼티를 사용할지 몰라도 된다
- 자손 컴포넌트는 주입받은 프로퍼티가 어디에서 온것인지 몰라도 된다

그러나 dependency injection의 좋지 않은 면도 있다. DI는 App의 컴포넌트들을 현재 구성된 방식으로 결합하여 리팩토링을 어렵게 만든다.

또한 `provided` 프로퍼티는 반응형이 아니다.
그것은 의도한 것인데, 왜냐하면 그것들을 사용해 중앙 데이터 저장소 범위를 만드는 것은 `$root`를 같은 목적으로 쓰는것 만큼 나쁜 방법이기 때문이다.

만약 App에서 공유하고 싶은 특정한 프로퍼티가 있다면, 또는 조상 컴포넌트 내부의 provided data를 업데이트 하고 싶다면,
Vuex와 같은 실제 상태를 관리하는 솔루션을 사용하고 싶은 징후이다.

### Programmatic Event Listeners

So far, you’ve seen uses of $emit, listened to with v-on, but Vue instances also offer other methods in its events interface. We can:

    Listen for an event with $on(eventName, eventHandler)
    Listen for an event only once with $once(eventName, eventHandler)
    Stop listening for an event with $off(eventName, eventHandler)

You normally won’t have to use these, but they’re available for cases when you need to manually listen for events on a component instance. They can also be useful as a code organization tool. For example, you may often see this pattern for integrating a 3rd-party library:

// Attach the datepicker to an input once
// it's mounted to the DOM.
mounted: function () {
  // Pikaday is a 3rd-party datepicker library
  this.picker = new Pikaday({
    field: this.$refs.input,
    format: 'YYYY-MM-DD'
  })
},
// Right before the component is destroyed,
// also destroy the datepicker.
beforeDestroy: function () {
  this.picker.destroy()
}

This has two potential issues:

    It requires saving the picker to the component instance, when it’s possible that only lifecycle hooks need access to it. This isn’t terrible, but it could be considered clutter.
    Our setup code is kept separate from our cleanup code, making it more difficult to programmatically clean up anything we set up.

You could resolve both issues with a programmatic listener:

mounted: function () {
  var picker = new Pikaday({
    field: this.$refs.input,
    format: 'YYYY-MM-DD'
  })

  this.$once('hook:beforeDestroy', function () {
    picker.destroy()
  })
}

Using this strategy, we could even use Pikaday with several input elements, with each new instance automatically cleaning up after itself:

mounted: function () {
  this.attachDatepicker('startDateInput')
  this.attachDatepicker('endDateInput')
},
methods: {
  attachDatepicker: function (refName) {
    var picker = new Pikaday({
      field: this.$refs[refName],
      format: 'YYYY-MM-DD'
    })

    this.$once('hook:beforeDestroy', function () {
      picker.destroy()
    })
  }
}

See this fiddle for the full code. Note, however, that if you find yourself having to do a lot of setup and cleanup within a single component, the best solution will usually be to create more modular components. In this case, we’d recommend creating a reusable <input-datepicker> component.

To learn more about programmatic listeners, check out the API for Events Instance Methods.

Note that Vue’s event system is different from the browser’s EventTarget API. Though they work similarly, $emit, $on, and $off are not aliases for dispatchEvent, addEventListener, and removeEventListener.
Circular References
Recursive Components

Components can recursively invoke themselves in their own template. However, they can only do so with the name option:

name: 'unique-name-of-my-component'

When you register a component globally using Vue.component, the global ID is automatically set as the component’s name option.

Vue.component('unique-name-of-my-component', {
  // ...
})

If you’re not careful, recursive components can also lead to infinite loops:

name: 'stack-overflow',
template: '<div><stack-overflow></stack-overflow></div>'

A component like the above will result in a “max stack size exceeded” error, so make sure recursive invocation is conditional (i.e. uses a v-if that will eventually be false).
Circular References Between Components

Let’s say you’re building a file directory tree, like in Finder or File Explorer. You might have a tree-folder component with this template:

<p>
  <span>{{ folder.name }}</span>
  <tree-folder-contents :children="folder.children"/>
</p>

Then a tree-folder-contents component with this template:

<ul>
  <li v-for="child in children">
    <tree-folder v-if="child.children" :folder="child"/>
    <span v-else>{{ child.name }}</span>
  </li>
</ul>

When you look closely, you’ll see that these components will actually be each other’s descendent and ancestor in the render tree - a paradox! When registering components globally with Vue.component, this paradox is resolved for you automatically. If that’s you, you can stop reading here.

However, if you’re requiring/importing components using a module system, e.g. via Webpack or Browserify, you’ll get an error:

Failed to mount component: template or render function not defined.

To explain what’s happening, let’s call our components A and B. The module system sees that it needs A, but first A needs B, but B needs A, but A needs B, etc. It’s stuck in a loop, not knowing how to fully resolve either component without first resolving the other. To fix this, we need to give the module system a point at which it can say, “A needs B eventually, but there’s no need to resolve B first.”

In our case, let’s make that point the tree-folder component. We know the child that creates the paradox is the tree-folder-contents component, so we’ll wait until the beforeCreate lifecycle hook to register it:

beforeCreate: function () {
  this.$options.components.TreeFolderContents = require('./tree-folder-contents.vue').default
}

Or alternatively, you could use Webpack’s asynchronous import when you register the component locally:

components: {
  TreeFolderContents: () => import('./tree-folder-contents.vue')
}

Problem solved!
Alternate Template Definitions
Inline Templates

When the inline-template special attribute is present on a child component, the component will use its inner content as its template, rather than treating it as distributed content. This allows more flexible template-authoring.

<my-component inline-template>
  <div>
    <p>These are compiled as the component's own template.</p>
    <p>Not parent's transclusion content.</p>
  </div>
</my-component>

However, inline-template makes the scope of your templates harder to reason about. As a best practice, prefer defining templates inside the component using the template option or in a <template> element in a .vue file.
X-Templates

Another way to define templates is inside of a script element with the type text/x-template, then referencing the template by an id. For example:

<script type="text/x-template" id="hello-world-template">
  <p>Hello hello hello</p>
</script>

Vue.component('hello-world', {
  template: '#hello-world-template'
})

These can be useful for demos with large templates or in extremely small applications, but should otherwise be avoided, because they separate templates from the rest of the component definition.
Controlling Updates

Thanks to Vue’s Reactivity system, it always knows when to update (if you use it correctly). There are edge cases, however, when you might want to force an update, despite the fact that no reactive data has changed. Then there are other cases when you might want to prevent unnecessary updates.
Forcing an Update

If you find yourself needing to force an update in Vue, in 99.99% of cases, you’ve made a mistake somewhere.

You may not have accounted for change detection caveats with arrays or objects, or you may be relying on state that isn’t tracked by Vue’s reactivity system, e.g. with data.

However, if you’ve ruled out the above and find yourself in this extremely rare situation of having to manually force an update, you can do so with $forceUpdate.
Cheap Static Components with v-once

Rendering plain HTML elements is very fast in Vue, but sometimes you might have a component that contains a lot of static content. In these cases, you can ensure that it’s only evaluated once and then cached by adding the v-once directive to the root element, like this:

Vue.component('terms-of-service', {
  template: `
    <div v-once>
      <h1>Terms of Service</h1>
      ... a lot of static content ...
    </div>
  `
})

Once again, try not to overuse this pattern. While convenient in those rare cases when you have to render a lot of static content, it’s simply not necessary unless you actually notice slow rendering – plus, it could cause a lot of confusion later. For example, imagine another developer who’s not familiar with v-once or simply misses it in the template. They might spend hours trying to figure out why the template isn’t updating correctly.

## 기타

### 재사용 가능한 컴포넌트 제작하기

Vue 컴포넌트API는 prop, 이벤트, 슬롯의 세 부분으로 나누어 진다

- Props 는 외부 환경을 데이터를 컴포넌트로 전달한다
- 컴포넌트는 이벤트를 통해서만 외부 환경에서 사이드이펙트를 발생시킬 수 있다
- 슬롯을 사용하면 외부 환경에서 추가 콘텐츠가 포함된 컴포넌트를 작성할 수 있다

```html
<my-component :foo="baz" :bar="qux" @event-a="doThis" @event-b="doThat">
  <img slot="icon" src="...">
  <p slot="main-text">Hello!</p>
</my-component>
```

### 재귀 컴포넌트

컴포넌트는 자신의 템플릿에서 재귀적으로 호출할 수 있습니다. 그러나, 그들은 name 옵션으로만 가능합니다.

name: 'unique-name-of-my-component'

Vue.component를 사용하여 컴포넌트를 전역적으로 등록하면, 글로벌 ID가 컴포넌트의 name 옵션으로 자동 설정됩니다.

Vue.component('unique-name-of-my-component', {
  // ...
})

주의하지 않으면 재귀적 컴포넌트로 인해 무한 루프가 발생할 수도 있습니다.

name: 'stack-overflow',
template: '<div><stack-overflow></stack-overflow></div>'

위와 같은 컴포넌트는 “최대 스택 크기 초과” 오류가 발생하므로 재귀 호출이 조건부 (즉, 마지막에 false가 될 v-if를 사용하세요)인지 확인하십시오.

### 컴포넌트 사이의 순환 참조

Finder나 파일 탐색기와 같이 파일 디렉토리 트리를 작성한다고 가정해 보겠습니다. 이 템플릿을 가지고 tree-folder 컴포넌트를 가질 수 있습니다.

<p>
  <span>{{ folder.name }}</span>
  <tree-folder-contents :children="folder.children"/>
</p>

그런 다음이 템플릿이 있는 tree-folder-contents 컴포넌트 :

<ul>
  <li v-for="child in children">
    <tree-folder v-if="child.children" :folder="child"/>
    <span v-else>{{ child.name }}</span>
  </li>
</ul>

자세히 살펴보면이 컴포넌트가 실제로 렌더링 트리에서 서로의 자식 및 조상인 패러독스라는 것을 알 수 있습니다! Vue.component를 이용해 전역으로 컴포넌트 등록할 때, 이 패러독스는 자동으로 해결됩니다. 그런 경우에 처해있으면 한번 읽어보세요.

그러나 모듈 시스템 을 사용하여 컴포넌트를 필요로하거나 가져오는 경우. Webpack 또는 Browserify를 통해 오류가 발생합니다.

컴포넌트를 마운트하지 못했습니다 : 템플릿 또는 렌더링 함수가 정의되지 않았습니다.

무슨 일이 일어나고 있는지 설명하기 위해 모듈 A와 B를 호출 할 것입니다. 모듈 시스템은 A가 필요합니다 하지만 A는 B를 우선적으로 필요로 합니다 게다가 B는 A를 필요로 하는 것을 알 수 있습니다. 먼저 서로 다른 컴포넌트를 해결하지 않고 두 컴포넌트를 완전히 해결하는 방법을 알지 못합니다. 이를 해결하려면 모듈 시스템에 “A는 B를 필요로 하나 B를 먼저 해결할 필요가 없습니다.”라고 말할 수있는 지점을 제공해야합니다.

여기에서는 tree-folder 컴포넌트로 삼을 것입니다. 패러독스를 만드는 자식은 tree-folder-contents 컴포넌트이므로, beforeCreate 라이프 사이클 훅이 등록 될 때까지 기다릴 것입니다.

beforeCreate: function () {
  this.$options.components.TreeFolderContents = require('./tree-folder-contents.vue')
}

문제가 해결되었습니다!

### 인라인 템플릿

하위 컴포넌트에 inline-template 이라는 특수한 속성이 존재할 때, 컴포넌트는 그 내용을 분산 된 내용으로 취급하지 않고 템플릿으로 사용합니다. 따라서 보다 유연한 템플릿 작성이 가능합니다.

<my-component inline-template>
  <div>
    <p>이것은 컴포넌트의 자체 템플릿으로 컴파일됩니다.</p>
    <p>부모가 만들어낸 내용이 아닙니다.</p>
  </div>
</my-component>

그러나, inline-template 은 템플릿의 범위를 추론하기 더 어렵게 만듭니다. 가장 좋은 방법은 template 옵션을 사용하거나.vue 파일의template 엘리먼트를 사용하여 컴포넌트 내부에 템플릿을 정의하는 것입니다.

### X-Templates

템플리트를 정의하는 또 다른 방법은 text/x-template 유형의 스크립트 엘리먼트 내부에 ID로 템플릿을 참조하는 것입니다. 예:

<script type="text/x-template" id="hello-world-template">
  <p>Hello hello hello</p>
</script>

Vue.component('hello-world', {
  template: '#hello-world-template'
})

이 기능은 큰 템플릿이나 매우 작은 응용 프로그램의 데모에는 유용 할 수 있지만 템플릿을 나머지 컴포넌트 정의와 분리하기 때문에 피해야합니다.

### v-once를 이용한 비용이 적게드는 정적 컴포넌트

일반 HTML 엘리먼트를 렌더링하는 것은 Vue에서 매우 빠르지만 가끔 정적 콘텐츠가 많이 포함 된 컴포넌트가 있을 수 있습니다. 이런 경우,v-once 디렉티브를 루트 엘리먼트에 추가함으로써 캐시가 한번만 실행되도록 할 수 있습니다.

Vue.component('terms-of-service', {
  template: '\
    <div v-once>\
      <h1>Terms of Service</h1>\
      ... a lot of static content ...\
    </div>\
  '
})
