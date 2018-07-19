# Vue.js 템플릿

Vue.js는 렌더링된 DOM을 기본 Vue 인스턴스 데이터에 선언적으로 바인딩 할 수있는 HTML 기반 템플릿 구문을 사용.
내부적으로 Vue는 템플릿을 가상 DOM 렌더링 함수로 컴파일 함. ( render() 함수 )

## 보간법

데이터 바인딩의 가장 기본 형태는 “Mustache” 구문(이중 중괄호: `{{ }}`)

```html
<span>메시지: {{ msg }}</span>
```

HTML 출력을 위해서는 v-html 디렉티브 사용 (취약점 문제로 거의 사용안함)

```html
<div v-html="rawHtml"></div>
```

속성 값 출력을 위해서 v-bind 디렉티브 사용

```html
<div v-bind:id="dynamicId"></div>
<button v-bind:disabled="isButtonDisabled">Button</button>
```

실제로 Vue.js는 모든 데이터 바인딩 내에서 JavaScript 표현식의 모든 기능을 지원

```html
{{ number + 1 }}
{{ ok ? 'YES' : 'NO' }}
{{ message.split('').reverse().join('') }}
<div v-bind:id="'list-' + id"></div>
```

## 디렉티브

디렉티브는 v- 접두사가 있는 특수 속성이다. v-for를 제외한 디렉티브 속성 값은 단일 JavaScript 표현식이 된다.
디렉티브로 연결된 표현식의 값이 변경되면, 그 값을 반응적으로 DOM에 적용한다.

## 전달인자 (v-bind)

일부 HTML 태그에 대응되는 디렉티브는 “전달인자”를 사용할 수 있다.

```html
<a v-bind:href="url"></a>
<div v-bind:id="dynamicId"></div>
```

boolean 속성의 경우 내용이 있는 경우, true인 경우 true 이지만 false는 여러 case로 가능하다.

```html
<button v-bind:disabled="isButtonDisabled">Button</button>
```

`null`, `undefined`, or `false`의 경우 `disabled` 속성은 태그에 포함되지 않을 것이다.

v-bind를 줄여서 (`:`)를 사용할 수 있다.

```html
<!-- 전체 구문 -->
<a v-bind:href="url"></a>
<!-- 약어 -->
<a :href="url"></a>
```

## v-bind: class / style

### HTML 클래스 바인딩

#### 클래스: 객체 구문

인라인 스타일링을 위해서 데이터 바인딩을 사용한다. 하지만 여러 속성을 string 연결로 처리하는 것은 성가신 일이다.
이런 이유로 class와 style을 위해 사용할 수 있는 v-bind의 특수기능을 사용할 수 있다.

객체를 `v-bind:class`로 보내서 클래스를 동적으로 추가/제거 할 수 있다.

```html
<!-- active 클래스는 isActive의 T/F 여부에 따라서 활성화 될 것이다 -->
<div v-bind:class="{ active: isActive }"></div>

여러개의 객체를 binding해서 여러개의 필드를 사용할 수 있고,
v-bind:class 디렉티브는 일반 class 속성과 공존할 수 있다 (:class가 class를 덮어쓰지 않음)
<div class="static" v-bind:class="{ active: isActive, 'text-danger': hasError }"></div>
```

클래스에 binding 되는 data는 다음과 같다면

```js
data: {
  isActive: true,
  hasError: false
}
```

아래와 같이 렌더링 된다

```html
<div class="static active"></div>
<!-- 해당 태그는 data 변화에 따라 동적으로 변한다. 만약 hasError가 true로 바뀐다면 클래스는 "static active text-danger" -->
```

바인딩 된 객체는 인라인일 필요는 없다 (풀어쓰지 않아도 됨), 아래도 위와 같은 결과로 렌더링 된다.

```html
<div v-bind:class="classObject"></div>
```

```js
data: {
  classObject: {
    active: true,
    'text-danger': false
  }
}
```

또한 객체를 반환하는 계산된 속성에도 바인딩 할 수 있다.

```html
<div v-bind:class="classObject"></div>
```

```js
data: {
  isActive: true,
  error: null
},
computed: {
  classObject: function () {
    return {
      active: this.isActive && !this.error,
      'text-danger': this.error && this.error.type === 'fatal'
    }
  }
}
```

#### 클래스: 배열 구문

배열을 v-bind:class 에 전달하여 클래스 목록을 지정할 수 있다

```html
<div v-bind:class="[activeClass, errorClass]"></div>
```

```js
data: {
  activeClass: 'active',
  errorClass: 'text-danger'
}
```

아래와 같이 렌더링 될 것이다

```html
<div class="active text-danger"></div>
```

목록에 있는 클래스를 조건부 토글하려면 삼항 연산자를 이용할 수 있다
그러나 여러 조건이 있는 경우 장황므로 배열 구문 내에서 객체 구문을 사용할 수 있다.

```html
<div v-bind:class="[isActive ? activeClass : '', errorClass]"></div>

<!-- 객체구문 사용 -->
<div v-bind:class="[{ active: isActive }, errorClass]"></div>
```

#### 컴포넌트와 함께 사용

사용자 정의 컴포넌트로 class 속성을 사용하면, 클래스는 컴포넌트의 루트 엘리먼트에 추가 되며 기존 클래스는 덮어쓰지 않는다.

```js
Vue.component('my-component', {
  template: '<p class="foo bar">Hi</p>'
})
```

```html
<my-component class="baz boo"></my-component>
```

아래와 같이 렌더링 된다

```html
<p class="foo bar baz boo">Hi</p>
```

클래스 바인딩도 동일하게 사용할 수 있다

```html
<my-component v-bind:class="{ active: isActive }"></my-component>
```

isActive가 true라면 아래와 같이 렌더링 된다

```html
<p class="foo bar active">Hi</p>
```

### 인라인 스타일 바인딩

#### 스타일: 객체 구문

v-bind:style 객체 구문은 매우 직설적으로, 거의 CSS 처럼 보이지만 JavaScript 객체이다.
속성 이름에 따옴표를 사용한 camelCase와 kebab-case를 사용할 수 있다.

```html
<div v-bind:style="{ color: activeColor, fontSize: fontSize + 'px' }"></div>
```

```js
data: {
  activeColor: 'red',
  fontSize: 30
}
```

스타일 객체에 직접 바인딩 하거나 계산된 속성과 함께 사용할 수도 있다

```html
<div v-bind:style="styleObject"></div>
```

```js
data: {
  styleObject: {
    color: 'red',
    fontSize: '13px'
  }
}
```

#### 스타일: 배열 구문

배열 구문은 같은 엘리먼트에 여러 개의 스타일 객체를 사용할 수 있게 한다

```html
<div v-bind:style="[baseStyles, overridingStyles]"></div>
```

#### 자동 접두사

`v-bind:style` 에 브라우저 벤더 접두어가 필요한 CSS 속성 (예: transform)을 사용하면,
Vue는 자동으로 해당 접두어를 감지하여 스타일을 적용한다.

#### 다중 값 제공

2.3 버전 부터 스타일 속성에 접두사가 있는 여러 값을 배열로 전달할 수 있다.

```html
<div v-bind:style="{ display: ['-webkit-box', '-ms-flexbox', 'flex'] }"></div>
```

브라우저가 지원하는 배열의 마지막 값만 렌더링한다.
위의 경우 flexbox의 접두어가 붙지않은 버전을 지원하는 브라우저에 대해 `display : flex`를 렌더링한다.

## 이벤트 수신자: v-on(@) 디렉티브

DOM 이벤트를 수신하는 v-on 디렉티브 예시이다. `<a v-on:click="doSomething">`

```html
<!-- 전체 구문 -->
<a v-on:click="doSomething"></a>
<!-- 약어 -->
<a @click="doSomething"></a>
```

### 인라인 핸들러

핸들러에에서 인라인으로 인자를 사용할 수 있다. 인자로 이벤트 자체를 넘겨줄 수도 있다.

```html
<button v-on:click="warn('Form cannot be submitted yet.', $event)">
  Submit
</button>
```

```js
// ...
methods: {
  warn: function (message, event) {
    // now we have access to the native event
    if (event) event.preventDefault()
    alert(message)
  }
}
```

### v-on(@) 디렉티브 modifier

- `.stop`: `event.stopPropagation()`: 이벤트에 부모태그가 반응(bubble-up)하는 것을 막는다
- `.prevent`: `event.preventDefault()`: 해당 이벤트만 실행하고 태그의 다른기능(`form action`, `a herf` ...) 실행을 막는다
- `.capture`
- `.self`
- `.once`
- `.passive`

> `.passive` 와 `.prevent`는 함께 사용하면 안된다. `.prevent`가 작동하지 않는다.

### key-modifier

```html
<!-- `vm.submit()`은 `keyCode` 13에 해당하는 버튼이 눌러졌을 때 실행됨 -->
<input v-on:keyup.13="submit">

<!-- 키코드 매칭 편의를 위해서 alias를 제공함 -->
<input v-on:keyup.enter="submit">
<input @keyup.enter="submit">
```

alias가 제공되는 modifer 목록

- `.enter`
- `.tab`
- `.delete` (“Delete”, “Backspace” 키 둘다 반응함)
- `.esc`
- `.space`
- `.up`
- `.down`
- `.left`
- `.right`

global 설정으로 사용자 정의 key modifier alias를 설정할 수 있다

```js
// `v-on:keyup.f1` 설정
Vue.config.keyCodes.f1 = 112
```

#### Automatic Key Modifiers

kebab-case로 `KeyboardEvent.key` 형식에 맞춰 직접 유효한 Key 이름을 입력하여 사용할 수 있다.

```html
<input @keyup.page-down="onPageDown">
```

위의 경우 핸들러는 `$event.key === 'PageDown'`인 경우에 호출될 것이다.

#### System Modifier Keys

키보드, 마우스 이벤트리스너와 함께 반응하는 시스템 키 modifier를 사용할 수 있다.

- `.ctrl`
- `.alt`
- `.shift`
- `.meta`: 매킨토시(command key (⌘)), 윈도우즈(windows key (⊞))

```html
<!-- Alt + C -->
<input @keyup.alt.67="clear">

<!-- Ctrl + Click -->
<div @click.ctrl="doSomething">Do something</div>
```

> 해당키만 단독으로 사용된 것을 감지하기 위해서는 `.exact`를 이용한다.

```html
<!-- Ctrl키와 함께 Alt 나 Shift 키가 눌러져도 실행됨 -->
<button @click.ctrl="onClick">A</button>

<!-- Ctrl 키만 눌러진 상태에서 실행됨 -->
<button @click.ctrl.exact="onCtrlClick">A</button>

<!-- system modifier가 하나도 눌러지지 않았을 때만 실행됨 -->
<button @click.exact="onClick">A</button>
```

#### Mouse Button Modifiers

- `.left`
- `.right`
- `.middle`

특정한 마우스 버튼에만 핸들러가 이벤트를 발생시킨다

## 조건부 렌더링: v-if

v-if 디렉티브를 사용하여 조건부 블록을 작성할 수 있다.

```html
<h1 v-if="ok">Yes</h1>
```

### 조건부 블록에 v-if 사용

`v-if`는 디렉티브기 때문에 하나의 엘리먼트에 사용할 수 있다.
하나 이상의 엘리먼트를 대상으로 동시에 조건부 블록을 만드려면 래퍼 역할을 하는 `<template>` 엘리먼트에 `v-if`를 사용할 수 있다.
최종 결과에는 `<template>` 엘리먼트가 렌더링 되지 않는다.

```html
<template v-if="ok">
  <h1>Title</h1>
  <p>Paragraph 1</p>
  <p>Paragraph 2</p>
</template>
```

### v-else

v-else 디렉티브를 사용하여 v-if에 대한 “else 블록”을 나타낼 수 있다.
v-else 엘리먼트는 v-if 엘리먼트 또는 v-else-if 엘리먼트 바로 뒤에 있어야 한다.

```html
<div v-if="Math.random() > 0.5">
  이제 나를 볼 수 있어요
</div>
<div v-else>
  이제는 안보입니다
</div>
```

### v-else-if

v-else-if는 v-if에 대한 “else if 블록” 역할을 하며, 여러개를 사용할 수 있습니다.
v-else와 마찬가지로, v-else-if 엘리먼트는 v-if 또는 v-else-if 엘리먼트 바로 뒤에 와야 한다.

```html
<div v-if="type === 'A'">
  A
</div>
<div v-else-if="type === 'B'">
  B
</div>
<div v-else-if="type === 'C'">
  C
</div>
<div v-else>
  Not A/B/C
</div>
```

### key를 이용한 재사용 가능한 엘리먼트 제어

Vue는 엘리먼트를 매번 렌더링 하지 않고 캐싱하여 사용하는 경우가 있다.
Vue를 빠르게 만드는데 도움이 되는 것 이외에 몇 가지 이점이 있다.

```html
<template v-if="loginType === 'username'">
  <label>사용자 이름</label>
  <input placeholder="사용자 이름을 입력하세요">
</template>
<template v-else>
  <label>이메일</label>
  <input placeholder="이메일 주소를 입력하세요">
</template>
```

위 코드에서 loginType을 바꾸어도 input value는 지워지지 않는다.
두 템플릿 모두 같은 요소를 사용하므로 `<input>`은 대체되지 않고 단지 label 내용과 placeholder만 변경된다.

이것은 항상 바람직하지는 않기 때문에, 유일한 값으로 key 속성을 추가하여 재사용하지 않음을 알리는 방법이 있다.

```html
<template v-if="loginType === 'username'">
  <label>사용자 이름</label>
  <input placeholder="사용자 이름을 입력하세요" key="username-input">
</template>
<template v-else>
  <label>이메일</label>
  <input placeholder="이메일 주소를 입력하세요" key="email-input">
</template>
```

이제 트랜지션 할 때마다 입력이 처음부터 렌더링된다.
`<label>` 엘리먼트는 key 속성이 없기 때문에 여전히 재사용 된다.

### v-show

엘리먼트를 조건부로 표시하기 위한 또 다른 옵션은 v-show 디렉티브이다.
v-show는 단순히 엘리먼트에 display CSS 속성을 토글하므로 엘리먼트는 항상 렌더링 이후 DOM에 남아있다.

```html
<h1 v-show="ok">안녕하세요!</h1>
```

v-show는 `<template>` 구문을 지원하지 않으며 `v-else`와도 작동하지 않는다.

### v-if vs v-show

v-if는 조건부 블럭 안의 이벤트 리스너와 자식 컴포넌트가 토글하는 동안 적절하게 제거되고 다시 만들어지기 때문에 실제 조건부 렌더링이다.

v-if는 초기 렌더링에서 조건이 거짓인 경우 아무것도 하지 않고, 조건 블록이 처음으로 참이 될 때 까지 렌더링 되지 않는다.

v-show는 CSS 기반 토글만으로 초기 조건에 관계 없이 엘리먼트가 항상 렌더링 된다.

따라서 v-if는 토글 비용이 높고 v-show는 초기 렌더링 비용이 더 높다.
자주 토글되는 element라면 v-show를, 런타임 시 조건이 바뀌지 않으면 v-if를 사용하는 것이 좋다.

### v-if 와 v-for

v-if와 함께 사용하는 경우, v-for는 v-if보다 높은 우선순위를 갖는다.

## 리스트 렌더링: v-for

### v-for로 엘리먼트에 배열 매핑하기

v-for 디렉티브는 배열을 대상으로 리스트를 렌더링한다. v-for 디렉티브는 `item in items(원본 배열)` 형태의 문법을 사용한다.

```html
<ul id="example-1">
  <li v-for="item in items">
    {{ item.message }}
  </li>
</ul>
```

```js
var example1 = new Vue({
  el: '#example-1',
  data: {
    items: [
      { message: 'Foo' },
      { message: 'Bar' }
    ]
  }
})
```

v-for 블록은 부모 범위 속성에 대한 모든 권한이 있으며, 현재 항목의 인덱스를 두 번째 전달인자로 받을 수 있다.

```html
<ul id="example-2">
  <li v-for="(item, index) in items">
    {{ parentMessage }} - {{ index }} - {{ item.message }}
  </li>
</ul>
```

```js
var example2 = new Vue({
  el: '#example-2',
  data: {
    parentMessage: 'Parent',
    items: [
      { message: 'Foo' },
      { message: 'Bar' }
    ]
  }
})
```

`in` 대신에 `of`를 구분자로 사용할 수 있다: `<div v-for="item of items"></div>`

### v-for와 객체

v-for를 사용하여 객체의 values를 출력할 수 있다.

```html
<ul id="v-for-object" class="demo">
  <li v-for="value in object">
    {{ value }}
  </li>
</ul>
```

```js
new Vue({
  el: '#v-for-object',
  data: {
    object: {
      firstName: 'John',
      lastName: 'Doe',
      age: 30
    }
  }
})
```

두 번째 인자로 key를 받을 수 있다.

```html
<div v-for="(value, key) in object">
  {{ key }}: {{ value }}
</div>
```

index는 세번째 인자로 받을 수 있다.

```html
<div v-for="(value, key, index) in object">
  {{ index }}. {{ key }} : {{ value }}
</div>
```

iteration 순서는 `Object.keys()`의 순서에 따라 결정되며, 결과는 JavaScript 엔진에 따라 다르다.

### key

v-for에서 렌더링된 엘리먼트 목록을 갱신할 때 기본적으로 “in-place patch” 전략을 사용한다.

데이터 항목의 순서가 변경된 경우 변경된 순서와 일치하도록 DOM 요소를 이동하는 대신,
각 요소를 적절한 위치에 패치하고 인덱스를 통해 렌더링할 내용을 반영하는지 확인한다.

목록의 출력 결과가 하위 컴포넌트 상태 또는 임시 DOM 상태(예: 폼 input)에 의존하지 않는 경우에는 기본전략이 적합하지 않을 수 있다.

각 노드의 id를 추적하여 기존 엘리먼트를 재사용하고 재정렬할 수 있도록 힌트를 제공하려면 각 항목에 고유 key 속성을 제공해야 한다.
key에 적합한 것은 데이터의 id이다. 속성처럼 작동하기 때문에 v-bind를 사용하여 동적 값에 바인딩 한다.

```html
<div v-for="item in items" :key="item.id">
  <!-- content -->
</div>
```

DOM 내용이 단순하지 않거나 의도적인 성능 향상을 위해 기본 동작에 의존하지 않는한, v-for에 key를 추가하는 것이 권장된다.

key는 Vue가 노드를 식별하는 일반적인 메커니즘이기 때문에, v-for 디렉티브가 아닌 일반적인 용도로 사용할 수 있다.

### 배열 변경 감지

#### 변이 메소드

Vue는 감시중인 배열의 변이 메소드를 래핑하여 뷰 갱신을 trigger한다.

- `push()`
- `pop()`
- `shift()`
- `unshift()`
- `splice()`
- `sort()`
- `reverse()`

items 배열에서 변이 메소드를 호출할 수 있다: `example1.items.push({ message: 'Baz' })`

#### 배열 대체

변이 메소드는 호출된 원본 배열을 변형한다.

변형을 하지 않는 방법도 있다. `filter()`, `concat()` 와 `slice()`를 사용하면 원본 배열을 변형하지 않고 항상 새 배열을 반환한다.

```js
example1.items = example1.items.filter(function (item) {
  return item.message.match(/Foo/)
})
```

새로운 배열을 반환하더라도 Vue가 기존 DOM을 버리고 전체 목록을 다시 렌더링 하지 않고, 배열을 겹치는 객체가 포함된 다른 배열로 대체한다.

#### 주의 사항

JavaScript의 제한으로 인해 Vue는 배열에 대해 다음과 같은 변경 사항을 감지할 수 없다.

- 인덱스로 배열에 있는 항목을 직접 설정하는 경우: `vm.items[indexOfItem] = newValue`
- 배열 길이를 수정하는 경우: `vm.items.length = newLength`

주의 사항 1번을 회피하기 위해 다음 두 경우 모두 `vm.items[indexOfItem] = newValue` 와 동일하게 수행하며,
반응형 시스템에서도 상태 변경을 트리거 한다.

```js
// Vue.set
Vue.set(example1.items, indexOfItem, newValue)

// Array.prototype.splice
example1.items.splice(indexOfItem, 1, newValue)
```

주의 사항 중 2번을 회피하기 위해 `splice`를 사용해야 한다.

```js
example1.items.splice(newLength)
```

### 객체 변경 감지에 관한 주의사항

JavaScript의 한계로 Vue는 속성 추가 및 삭제를 감지하지 못한다.

```js
var vm = new Vue({
  data: {
    a: 1
  }
})
// `vm.a` 는 반응형입니다.

vm.b = 2
// `vm.b` 는 반응형이 아닙니다.
```

Vue는 이미 만들어진 인스턴스의 루트레벨에서 새로운 반응형 속성을 동적으로 추가하는 것을 허용하지 않는다.

그러나 Vue.set(object, key, value) 메소드를 사용하여 중첩된 객체에 반응형 속성을 추가할 수 있다.

```js
var vm = new Vue({
  data: {
    userProfile: {
      name: 'Anika'
    }
  }
})

// userProfile 객체에 새로운 속성 age를 추가
Vue.set(vm.userProfile, 'age', 27)

// 인스턴스 메소드인 vm.$set 사용가능. 이는 전역 Vue.set의 별칭임
vm.$set(this.userProfile, 'age', 27)
```

`Object.assign()`이나 `_.extend()`를 사용해 기존의 객체에 새 속성을 할당할 수도 있다.
이 경우 두 객체의 속성을 사용해 새 객체를 만들어야 한다.

```js
Object.assign(this.userProfile, {
  age: 27,
  favoriteColor: 'Vue Green'
})
```

새로운 반응형 속성을 추가한다.

```js
this.userProfile = Object.assign({}, this.userProfile, {
  age: 27,
  favoriteColor: 'Vue Green'
})
```

### 필터링 / 정렬 된 결과 표시하기

원본 데이터를 실제로 변경하지 않고 배열의 필터링 또는 정렬된 버전을 표시해야 할 필요가 있다.
이 경우 필터링 된 배열이나 정렬된 배열을 반환하는 계산된 속성을 만들어서 사용한다.

```html
<li v-for="n in evenNumbers">{{ n }}</li>
```

```js
data: {
  numbers: [ 1, 2, 3, 4, 5 ]
},
computed: {
  evenNumbers: function () {
    return this.numbers.filter(function (number) {
      return number % 2 === 0
    })
  }
}
```

계산된 속성을 실행할 수 없는 상황(예: 중첩 된 v-for 루프 내부)에서는 다음 방법을 사용한다.

```html
<li v-for="n in even(numbers)">{{ n }}</li>
```

```js
data: {
  numbers: [ 1, 2, 3, 4, 5 ]
},
methods: {
  even: function (numbers) {
    return numbers.filter(function (number) {
      return number % 2 === 0
    })
  }
}
```

### Range v-for

v-for에서 숫자 배열을 사용할 수 있다.

```html
<div>
  <span v-for="n in 10">{{ n }} </span>
</div>
```

### v-for 템플릿

v-if와 마찬가지로, `<template>`태그를 사용해 여러 엘리먼트의 블럭을 렌더링 할 수 있다.

```html
<ul>
  <template v-for="item in items">
    <li>{{ item.msg }}</li>
    <li class="divider"></li>
  </template>
</ul>
```

### v-for 와 v-if

동일한 노드에 v-for와 v-if가 있다면 v-for가 높은 우선순위를 갖는다.
즉, v-if는 루프가 반복될 때마다 실행된다. _일부_ 항목만 렌더링 하려는 경우 유용하다.

```html
<li v-for="todo in todos" v-if="!todo.isComplete">
  {{ todo }}
</li>
```

실행을 조건부로 하는 것이 목적이라면 래퍼 엘리먼트(또는 `<template>`)을 사용하여야 한다(상위노드에서 조건 렌더링)

```html
<ul v-if="todos.length">
  <li v-for="todo in todos">
    {{ todo }}
  </li>
</ul>
<p v-else>No todos left!</p>
```

### v-for 와 컴포넌트

v-for를 사용자 정의 컴포넌트에 직접 사용할 수 있다.

```html
<my-component v-for="item in items" :key="item.id"></my-component>
```

그러나 컴포넌트는 자체 범위를 갖기 때문에 반복할 데이터를 컴포넌트로 전달하려면 props도 사용해야 한다.

```html
<my-component
  v-for="(item, index) in items"
  v-bind:item="item"
  v-bind:index="index"
  v-bind:key="item.id">
</my-component>
```

컴포넌트에 item을 자동 주입하지 않는 이유는 컴포넌트가 v-for와 결합도를 낮추어 컴포넌트를 재사용하기 위함이다.

할 일 목록 예제를 보자

```html
<div id="todo-list-example">
  <input v-model="newTodoText"
    v-on:keyup.enter="addNewTodo"
    placeholder="Add a todo">
  <ul>
    <li is="todo-item"
      v-for="(todo, index) in todos"
      v-bind:key="todo.id"
      v-bind:title="todo.title"
      v-on:remove="todos.splice(index, 1)">
    </li>
  </ul>
</div>
```

`is="todo-item"` 속성은 `<todo-item>`과 같은 일을 하지만 `<li>` 엘리먼트는 `<ul>` 안에서만 유효하므로,
잠재적인 브라우저의 구문 분석 오류를 해결할 수 있다.

```js
Vue.component('todo-item', {
  template: '\
    <li>\
      {{ title }}\
      <button v-on:click="$emit(\'remove\')">X</button>\
    </li>\
  ',
  props: ['title']
})

new Vue({
  el: '#todo-list-example',
  data: {
    newTodoText: '',
    todos: [
      {
        id: 1,
        title: 'Do the dishes',
      },
      {
        id: 2,
        title: 'Take out the trash',
      },
    ],
    nextTodoId: 3
  },
  methods: {
    addNewTodo: function () {
      this.todos.push({
        id: this.nextTodoId++,
        title: this.newTodoText
      })
      this.newTodoText = ''
    }
  }
})
```

## 폼 입력 바인딩 (v-model)

`v-model` 디렉티브를 사용하여 폼 input과 textarea 엘리먼트에 양방향 데이터 바인딩을 생성할 수 있다.

`v-model`은 기본적으로 사용자 입력 이벤트에 대한 데이터를 업데이트하는 “syntax sugar”이며 일부 경우에 특별히 주의 해야함

> v-model은 모든 form 엘리먼트의 초기 value와 checked 그리고 selected 속성을 무시하며 항상 Vue 인스턴스 데이터를 원본 소스로 취급한다. 따라서 컴포넌트의 data 옵션 안에 있는 JavaScript에서 초기값을 선언해야 한다.

### v-model 기본 사용법

```html
<input v-model="message" placeholder="여기를 수정해보세요">
<p>메시지: {{ message }}</p>
```

여러 개의 체크박스는 같은 배열을 바인딩 할 수 있다.

```html
<div id='example-3'>
  <input type="checkbox" id="jack" value="Jack" v-model="checkedNames">
  <label for="jack">Jack</label>
  <input type="checkbox" id="john" value="John" v-model="checkedNames">
  <label for="john">John</label>
  <input type="checkbox" id="mike" value="Mike" v-model="checkedNames">
  <label for="mike">Mike</label>
  <br>
  <span>체크한 이름: {{ checkedNames }}</span>
</div>
```

라디오 버튼 및 셀렉트는 단일 객체가 반환된다.

```html
<input type="radio" id="one" value="One" v-model="picked">
<label for="one">One</label>
<br>
<input type="radio" id="two" value="Two" v-model="picked">
<label for="two">Two</label>
<br>
<span>선택: {{ picked }}</span>
```

> v-model 표현식의 초기 값이 어떤 옵션에도 없으면, `<select>` 엘리먼트는 “선택없음” 상태로 렌더링 된다.. iOS에서는 이 경우 변경 이벤트가 발생하지 않아 사용자가 첫 번째 항목을 선택할 수 없게되므로, 위 예제처럼 사용하지 않는 옵션에 빈 값을 넣는 것이 좋다.

v-for를 이용한 동적 옵션 렌더링 및 text/value 분리

```html
<select v-model="selected">
  <option v-for="option in options" v-bind:value="option.value">
    {{ option.text }}
  </option>
</select>
<span>Selected: {{ selected }}</span>
```

```js
new Vue({
  el: '...',
  data: {
    selected: 'A',
    options: [
      { text: 'One', value: 'A' },
      { text: 'Two', value: 'B' },
      { text: 'Three', value: 'C' }
    ]
  }
})
```

### v-model 수식어

`.lazy`: 기본적으로, v-model은 각 입력 이벤트 후 입력과 데이터를 동기화한다. 수식어로 데이터 변경이후 동기화 할 수 있다.

```html
<!-- "input" 대신 "change" 이후에 동기화 됩니다. -->
<input v-model.lazy="msg" >
```

`.number`: 사용자 입력이 자동으로 숫자로 형변환 되기를 원하면, v-model이 관리하는 input에 number 수식어를 추가한다.

```html
<input v-model.number="age" type="number">
```

`.trim`: v-model이 관리하는 input을 자동으로 trim 하기 원하면, trim 수정자를 추가한다.

```html
<input v-model.trim="msg">
```

## 필터

Vue.js에서는 텍스트 formatting을 위해서 필터를 사용할 수 있다. 필터는 Mustache 보간과 v-bind 표현식 두 곳에서 사용할 수 있다.
필터는 JavaScript 표현식의 끝에 `|` 기호를 추가해서 사용한다.

```html
<!-- Mustaches 사용시 -->
{{ message | capitalize }}
<!-- v-bind 사용시 -->
<div v-bind:id="rawId | formatId"></div>
```

필터 함수는 항상 표현식의 값을 첫번째 인자로 받는다.

```js
new Vue({
  // ...
  filters: {
    capitalize(value) {
      if (!value) return '';
      value = value.toString();
      return value.charAt(0).toUpperCase() + value.slice(1);
    }
  }
})
```

필터는 체이닝 가능하다. `{{ message | filterA | filterB }}`

필터는 JavaScript 함수로 선언되므로 인자를 받을 수 있다. `{{ message | filterA('arg1', arg2) }}`
