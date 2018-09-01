# Reusability & Composition

## Mixin

### 기초

Mixin은 Vue 컴포넌트를 재사용하기 위한 기능이다.

mixin 객체는 모든 구성 요소 옵션을 포함할 수 있다.
컴포넌트에 mixin을 사용하면 해당 mixin의 모든 옵션이 컴포넌트의 고유 옵션에 혼합된다.

```js
// mixin 객체 생성
var myMixin = {
  created() {
    this.hello()
  },
  methods: {
    hello() {
      console.log('hello from mixin!')
    }
  }
}

// mixin을 사용할 컴포넌트 정의
var Component = Vue.extend({
  mixins: [myMixin]
})

var component = new Component() // => "hello from mixin!"
```

### 옵션 병합

mixin과 컴포넌트에 중첩 옵션이 포함되어 있으면 **적절한 전략**에 따라 병합된다.

같은 이름을 가진 훅 함수의 경우, 배열에 들어가 mixin의 훅부터 순서대로 호출된다

```js
var mixin = {
  created() {
    console.log('mixin hook called')
  }
}

new Vue({
  mixins: [mixin],
  created() {
    console.log('component hook called')
  }
})

// => "mixin hook called"
// => "component hook called"
```

methods, components, directives 처럼 객체 값을 요구하는 옵션은 같은 객체에 병합된다.
이러한 객체에 충돌하는 키가 있을 경우 컴포넌트의 옵션이 우선순위를 갖습니다.

data의 경우도 충돌하는 속성이 있는 경우 컴포넌트의 값이 우선순위를 갖는다.

```js
var mixin = {
  methods: {
    foo: function () {
      console.log('foo')
    },
    conflicting: function () {
      console.log('from mixin')
    }
  }
}

var vm = new Vue({
  mixins: [mixin],
  methods: {
    bar: function () {
      console.log('bar')
    },
    conflicting: function () {
      console.log('from self')
    }
  }
})

vm.foo() // => "foo"
vm.bar() // => "bar"
vm.conflicting() // => "from self"
```

`Vue.extend()`에서도 같은 병합 전략이 사용된다.

### 전역 Mixin

mixin은 전역으로 적용할 수 있으며 전역 mixin은 이후에 생성된 모든 Vue 인스턴스에 영향을 미친다

```js
// `myOption` 사용자 정의 옵션을 위한 핸들러 주입
Vue.mixin({
  created() {
    var myOption = this.$options.myOption
    if (myOption) {
      console.log(myOption)
    }
  }
})

new Vue({
  myOption: 'hello!'
})
// => "hello!"
```

전역 mixin은 생성된 모든 단일 Vue 인스턴스(3rd-party 컴포넌트를 포함)에 영향을 주기 때문에 신중하게 사용해야 한다.

### 사용자 정의 옵션 병합 전략

커스텀 로직을 사용해 커스텀 옵션을 병합하려면, `Vue.config.optionMergeStrategies`에 함수를 정의하면된다.

```js
Vue.config.optionMergeStrategies.myOption = (toVal, fromVal) => {
  // return 병합된 값
}
```

대부분의 객체 기반 옵션에서 methods에서 사용한 것과 같은 전략을 간단하게 사용할 수 있다

```js
var strategies = Vue.config.optionMergeStrategies
strategies.myOption = strategies.methods
```

고급 예제는 Vuex의 1.x 병합 전략으로 살펴볼 수 있다

```js
const merge = Vue.config.optionMergeStrategies.computed
Vue.config.optionMergeStrategies.vuex = (toVal, fromVal) => {
  if (!toVal) return fromVal
  if (!fromVal) return toVal
  return {
    getters: merge(toVal.getters, fromVal.getters),
    state: merge(toVal.state, fromVal.state),
    actions: merge(toVal.actions, fromVal.actions)
  }
}
```

## 사용자 지정 디렉티브

### 시작

Vue는 기본 디렉티브(v-model과 v-show ...) 외에도 사용자 정의 디렉티브를 등록할 수 있다.

Vue 2에서 코드 재사용 및 추상화는 컴포넌트 단위에서 이루어지지만
일반 엘리먼트에 하위 수준의 DOM 액세스가 필요한 경우가 있을 수 있으며 이 경우 사용자 지정 디렉티브를 사용할 수도 있다.

input 엘리먼트와 focusing에 대한 작업을 수행하는 디렉티브 예제를 보자

페이지가 로드되면 해당 엘리먼트는 포커스를 얻는다.
이 페이지를 방문한 이후 다른것을 클릭하지 않았다면 input 엘리먼트에 포커스가 되어 있어야한다.

```js
// 전역 사용자 정의 디렉티브 v-focus 등록
Vue.directive('focus', {
  // 바인딩 된 엘리먼트가 DOM에 삽입되었을 때...
  inserted: function (el) {
    // 엘리먼트에 포커스를 줍니다
    el.focus()
  }
})
```

컴포넌트의 `directives` 옵션으로 디렉티브를 로컬범위로 등록할 수 있다

```js
directives: {
  focus: {
    // 디렉티브 정의
    inserted: function (el) {
      el.focus()
    }
  }
}
```

템플릿에서는 모든 요소에서 새로운 v-focus 속성을 사용할 수 있다: `<input v-focus>`

### 훅 함수

디렉티브 정의 객체는 여러가지 훅 함수와 함께 제공될 수 있다

- bind: 디렉티브가 처음 엘리먼트에 바인딩 될 때 한번만 호출됨
- inserted: 바인딩 된 엘리먼트가 부모 노드에 삽입 되었을 때 호출됨 (이것은 부모 노드 존재를 보장하며 반드시 document내에 있는 것은 아님)
- update: 포함하는 컴포넌트가 업데이트 된 후 호출되지만 자식이 업데이트 되기 전일 가능성이 있다. 디렉티브의 값과 관계없이 바인딩의 현재 값과 이전 값을 비교하여 불필요한 업데이트를 건너 뛸 수 있다.
- componentUpdated: 포함하고 있는 컴포넌트와 그 자식들 이 업데이트 된 후에 호출됨
- unbind: 디렉티브가 엘리먼트로부터 언바인딩된 경우에만 한번 호출됨

### 디렉티브 훅 전달인자

디렉티브 훅은 다음을 전달인자로 사용할 수 있다

- el: 디렉티브가 바인딩된 엘리먼트. 이 것을 사용하면 DOM 조작을 할 수 있다
- binding: 바인딩은 아래의 속성을 포함한다
  - name: 디렉티브 이름, `v-` 접두사가 없다
  - value: 디렉티브에서 전달받은 값. 예를 들어 `v-my-directive="1 + 1"`인 경우 value는 `2` 이다
  - expression: 표현식 문자열. 예를 들어 `v-my-directive="1 + 1"`이면, 표현식은 `"1 + 1"` 이다
  - oldValue: 이전 값. `update`와 `componentUpdated`에서만 사용할 수 있다. 이를 통해 값이 변경되었는지 확인할 수 있다
  - arg: 디렉티브의 전달인자, 있는 경우에만 존재합니다. 예를 들어 `v-my-directive:foo` 이면 `"foo"` 이다
  - modifiers: 포함된 수식어가 있다면 확인한다. 예를 들어 `v-my-directive.foo.bar`이면, 수식어 객체는 `{ foo: true, bar: true }` 이다
- vnode: Vue 컴파일러가 만든 가상노드. `VNode API` 참조
- oldVnode: 이전의 가상노드. `update`와 `componentUpdated`에서만 사용할 수 있다

모든 전달인자는 읽기 전용으로 사용하여야 한다.
훅을 통해 이 정보들을 전달하는 경우, 엘리먼트의 dataset을 이용하면 된다.

```html
<div id="hook-arguments-example" v-demo:foo.a.b="message"></div>
```

```js
Vue.directive('demo', {
  bind: function (el, binding, vnode) {
    var s = JSON.stringify
    el.innerHTML =
      'name: '       + s(binding.name) + '<br>' +
      'value: '      + s(binding.value) + '<br>' +
      'expression: ' + s(binding.expression) + '<br>' +
      'argument: '   + s(binding.arg) + '<br>' +
      'modifiers: '  + s(binding.modifiers) + '<br>' +
      'vnode keys: ' + Object.keys(vnode).join(', ')
  }
})

new Vue({
  el: '#hook-arguments-example',
  data: {
    message: 'hello!'
  }
})
```

### 함수 약어(shorthand)

많은 경우에, bind와 update에서 같은 동작이 이루어지길 원할 수 있다

```js
Vue.directive('color-swatch', function (el, binding) {
  el.style.backgroundColor = binding.value
})
```

### 객체 리터럴

디렉티브에 여러 값이 필요한 경우, JavaScript 객체 리터럴을 전달할 수 있다.
디렉티브는 유효한 JavaScript 표현식을 사용할 수 있다.

```html
<div v-demo="{ color: 'white', text: 'hello!' }"></div>
```

```js
Vue.directive('demo', function (el, binding) {
  console.log(binding.value.color) // => "white"
  console.log(binding.value.text)  // => "hello!"
})
```

## Render Functions & JSX

### 기본

Vue는 HTML 작성시 템플릿을 사용할 것을 권장한다.
그러나 완전한 JavaScript가 필요한 경우 템플릿대용으로 컴파일러에 가까운 render 함수를 사용할 수 있다.

링크를 포함한 헤더를 생성한다고 가정한 예제를 보자

```html
<h1>
  <a name="hello-world" href="#hello-world">
    Hello world!
  </a>
</h1>
```

위의 HTML을 생성하기 위해 다음의 컴포넌트 인터페이스가 필요할 것이다

```html
<anchored-heading :level="1">Hello world!</anchored-heading>
```

level prop를 기반으로 제목을 생성하는 컴포넌트이다

```html
<script type="text/x-template" id="anchored-heading-template">
  <h1 v-if="level === 1">
    <slot></slot>
  </h1>
  <h2 v-else-if="level === 2">
    <slot></slot>
  </h2>
  <h3 v-else-if="level === 3">
    <slot></slot>
  </h3>
  <h4 v-else-if="level === 4">
    <slot></slot>
  </h4>
  <h5 v-else-if="level === 5">
    <slot></slot>
  </h5>
  <h6 v-else-if="level === 6">
    <slot></slot>
  </h6>
</script>
```

```js
Vue.component('anchored-heading', {
  template: '#anchored-heading-template',
  props: {
    level: {
      type: Number,
      required: true
    }
  }
})
```

이 템플릿은 장황할 뿐만 아니라 모든 헤딩 수준에 대해 `<slot> </slot>`을 중복으로 가지고 있으며
앵커 엘리먼트를 추가 할 때도 똑같이 해야 하는 문제점이 있다.

이제 render 함수로 다시 작성해보자

```js
Vue.component('anchored-heading', {
  render: function (createElement) {
    return createElement(
      'h' + this.level,   // 태그 이름
      this.$slots.default // 자식의 배열
    )
  },
  props: {
    level: {
      type: Number,
      required: true
    }
  }
})
```

이 경우 anchored-heading 안에 Hello world!와 같이 slot 속성 없이 자식을 전달할 때
그 자식들은 `$slots.default` 에있는 컴포넌트 인스턴스에 저장된다는 것을 알아야한다.

### 노드, 트리, 그리고 버추얼 DOM

render 함수를 알아보기 전에 브라우저 작동 방식을 알아야한다

```html
<div>
  <h1>My title</h1>
  Some text content
</div>
```

브라우저가 이 코드를 읽게 되면, 모든 내용을 추적하기 “DOM 노드” 트리를 만든다.

모든 엘리먼트는 노드이다. 각 텍스트도 주석도 노드이다. 그리고 각 노드는 자식을 가질 수 있다.

노드를 효율적으로 갱신하는 것은 어렵지만 수동으로 할 필요는 없다. 템플릿에서 Vue가 페이지에서 수정하기 원하는 HTML만 지정하면 된다.

```html
<h1>{{ blogTitle }}</h1>
```

또는 render 함수에서

```js
render: function (createElement) {
  return createElement('h1', this.blogTitle)
}
```

두가지 경우 모두 Vue는 페이지및 blogTitle을 자동으로 갱신한다.

#### 버추얼 DOM

Vue는 실제 DOM에 필요한 변경사항을 추적하기 위해 virtual DOM을 만든다.

```js
return createElement('h1', this.blogTitle)
```

createElement의 반환값은 실제 DOM 엘리먼트와 정확하게 일치하지는 않는다.
반환값인 createNodeDescription는 Vue에게 자식노드에 대한 설명을 포함하여 페이지에서 렌더링해야하는 노드의 종류를 설명하는 정보를 포함한다.

이런 노드 명세를 가상노드라고 부른다. “버추얼 DOM”은 Vue 컴포넌트 트리로 만들어진 가상노트 트리이다.

### createElement 전달인자

createElement 함수에서 템플릿 기능을 사용하는 방법을 살펴보자

```js
// @returns {VNode}
createElement(
  // {String | Object | Function}
  // HTML 태그 이름, 컴포넌트 옵션 또는 비동기 함수 중 하나를 반환한다 (필수 사항)
  'div',

  // {Object}
  // 템플릿에서 사용할 속성에 해당하는 데이터 객체 (선택 사항)
  {
    // 다음 섹션 참고
  },

  // {String | Array}
  // VNode 자식들. `createElement()`를 사용해 만들거나, 간단히 문자열을 사용해 'text VNodes'를 얻을 수 있다 (선택사항)
  [
    'Some text comes first.',
    createElement('h1', 'A headline'),
    createElement(MyComponent, {
      props: {
        someProp: 'foobar'
      }
    })
  ]
)
```

#### 데이터 객체 깊이 알아 보기

`v-bind:class` 와 `v-bind:style`이 템플릿에서 특별한 처리를 하는 것과 비슷하게, VNode 데이터 객체에 최상위 필드가 있다.

이 객체는 `innerHTML`과 같은 DOM 속성뿐 아니라 일반적인 HTML 속성도 바인딩 할 수 있게 한다.
(`v-html` 디렉티브를 대신해 사용할 수 있다)

```js
{
  // `v-bind:class` 와 같음
  'class': {
    foo: true,
    bar: false
  },
  // `v-bind:style` 와 같음
  style: {
    color: 'red',
    fontSize: '14px'
  },
  // 일반 HTML 속성
  attrs: {
    id: 'foo'
  },
  // 컴포넌트 props
  props: {
    myProp: 'bar'
  },
  // DOM 속성
  domProps: {
    innerHTML: 'baz'
  },
  // `v-on:keyup.enter`와 같은 수식어가 지원되지 않으나 이벤트 핸들러는 `on` 아래에 중첩된다
  // 수동으로 핸들러에서 keyCode를 확인해야 한다
  on: {
    click: this.clickHandler
  },
  // 컴포넌트 전용.
  // `vm.$emit`를 사용하여 컴포넌트에서 발생하는 이벤트가 아닌 기본 이벤트를 받을 수 있게 한다
  nativeOn: {
    click: this.nativeClickHandler
  },
  // 사용자 지정 디렉티브.
  // Vue는 이를 관리하기 때문에 바인딩의 oldValue는 설정할 수 없다
  directives: [
    {
      name: 'my-custom-directive',
      value: '2',
      expression: '1 + 1',
      arg: 'foo',
      modifiers: {
        bar: true
      }
    }
  ],
  // 범위 지정 슬롯. 형식은 { name: props => VNode | Array<VNode> }
  scopedSlots: {
    default: props => createElement('span', props.text)
  },
  // 이 컴포넌트가 다른 컴포넌트의 자식인 경우, 슬롯의 이름
  slot: 'name-of-slot',
  // 기타 최고 레벨 속성
  key: 'myKey',
  ref: 'myRef'
}
```

#### 전체 예제

```js
var getChildrenTextContent = function (children) {
  return children.map(function (node) {
    return node.children
      ? getChildrenTextContent(node.children)
      : node.text
  }).join('')
}

Vue.component('anchored-heading', {
  render: function (createElement) {
    // kebabCase id를 만듭니다.
    var headingId = getChildrenTextContent(this.$slots.default)
      .toLowerCase()
      .replace(/\W+/g, '-')
      .replace(/(^\-|\-$)/g, '')

    return createElement(
      'h' + this.level,
      [
        createElement('a', {
          attrs: {
            name: headingId,
            href: '#' + headingId
          }
        }, this.$slots.default)
      ]
    )
  },
  props: {
    level: {
      type: Number,
      required: true
    }
  }
})
```

#### 제약사항

컴포넌트 트리의 **모든 VNode는 고유 해야 한다**.

렌더링 함수가 유효하지 않은 예제를 보자

```js
render: function (createElement) {
  var myParagraphVNode = createElement('p', 'hi')
  return createElement('div', [
    // Vnode가 중복됨!
    myParagraphVNode, myParagraphVNode
  ])
}
```

같은 엘리먼트 / 컴포넌트를 여러 번 복제하려는 경우 팩토리 기능을 사용하여 여러 번 반복 할 수 있다.

예를 들어, 다음 렌더링 함수는 20개의 같은 p태그를 완벽하게 렌더링하는 방법이다.

```js
render: function (createElement) {
  return createElement('div',
    Array.apply(null, { length: 20 }).map(function () {
      return createElement('p', 'hi')
    })
  )
}
```

### 템플릿 기능을 일반 JavaScript로 변경하기

#### v-if 와 v-for

어디든 JavaScript를 사용할 수 있는 환경이면 Vue 렌더링 함수는 한가지 방법만을 제공하지는 않는다.

예를 들어, v-if와 v-for를 사용하는 템플릿에서

```html
<ul v-if="items.length">
  <li v-for="item in items">{{ item.name }}</li>
</ul>
<p v-else>No items found.</p>
```

이것은 render 함수에서 `if / else` 와 `map`을 사용하여 재작성 할 수 있다

```js
render: function (createElement) {
  if (this.items.length) {
    return createElement('ul', this.items.map(function (item) {
      return createElement('li', item.name)
    }))
  } else {
    return createElement('p', 'No items found.')
  }
}
```

#### v-model

렌더 함수에는 직접적으로 v-model에 대응되는 것이 없으므로 직접 구현해야 한다.

```js
render: function (createElement) {
  var self = this
  return createElement('input', {
    domProps: {
      value: self.value
    },
    on: {
      input: function (event) {
        self.value = event.target.value
        self.$emit('input', event.target.value)
      }
    }
  })
}
```

#### 이벤트 및 키 수식어

`.passive`, `.capture` 및 `.once` 이벤트 수식어를 위해 Vue는 `on`과 함께 사용할 수 있는 접두사를 제공한다

| 수식어 | 접두어 |
| ----- | ----- |
| `.passive` | `&` |
| `.capture` | `!` |
| `.once` | `~` |
| `.capture.once` 또는 `.once.capture` | `~!` |

##### 예제

```js
on: {
  '!click': this.doThisInCapturingMode,
  '~keyup': this.doThisOnce,
  `~!mouseover`: this.doThisOnceInCapturingMode
}
```

다른 모든 이벤트 및 키 수식어의 경우 처리기에서 이벤트 메서드를 간단하게 사용할 수 있으므로 고유한 접두사는 필요하지 않음

| 수식어 | 동등한 핸들러 |
| ----- | ----- |
| `.stop` | `event.stopPropagation()` |
| `.prevent` | `event.preventDefault()` |
| `.self` | `if (event.target !== event.currentTarget) return` |
| 키: `.enter`, `.13` | `if (event.keyCode !== 13) return` (13 대신 다른 키코드 사용가능) |
| Modifiers Keys: `.ctrl`, `.alt`, `.shift`, `.meta` | `if (!event.ctrlKey) return` (ctrlKey를 altKey, shiftKey 또는 metaKey로 변경) |

다음은 위의 수식어를 사용한 예제이다

```js
on: {
  keyup: function (event) {
    // 이벤트를 내보내는 요소가 이벤트가 바인딩 된 요소가 아닌 경우 중단
    if (event.target !== event.currentTarget) return
    // 키보드에서 뗀 키가 Enter키 (13)이 아니며 Shift키가 동시에 눌러지지 않은 경우 중단
    if (!event.shiftKey || event.keyCode !== 13) return
    // 전파를 멈춘다
    event.stopPropagation()
    // 엘리먼트 기본 동작을 방지한다
    event.preventDefault()
    // ...
  }
}
```

#### Slots

`this.$slots`에서 정적 슬롯 내용을 VNodes의 배열로 접근할 수 있다

```js
render: function (createElement) {
  // `<div><slot></slot></div>`
  return createElement('div', this.$slots.default)
}
```

특정 범위를 가지는 슬롯 `this.$scopedSlots`에서 VNode를 반환하는 함수로 접근할 수 있다

```js
render: function (createElement) {
  // `<div><slot :text="msg"></slot></div>`
  return createElement('div', [
    this.$scopedSlots.default({
      text: this.msg
    })
  ])
}
```

범위 함수 슬롯을 렌더링 함수를 사용하여 하위 컴포넌트로 전달하려면 VNode 데이터에서 `scopedSlots` 필드를 사용한다

```js
render (createElement) {
  return createElement('div', [
    createElement('child', {
      // 데이터 객체의 `scopedSlots`를 다음 형식으로 전달합니다
      // { name: props => VNode | Array<VNode> }
      scopedSlots: {
        default: function (props) {
          return createElement('span', props.text)
        }
      }
    })
  ])
}
```

### JSX

render 함수를 많이 작성하면 다음과 같이 작성하는 것이 불편할 것이다

```js
createElement(
  'anchored-heading', {
    props: {
      level: 1
    }
  }, [
    createElement('span', 'Hello'),
    ' world!'
  ]
)
```

템플릿 버전이 아래 처럼 너무 간단한 경우에 특히 더 그럴 것

```html
<anchored-heading :level="1">
  <span>Hello</span> world!
</anchored-heading>
```

Vue와 JSX를 함께 사용하기 위해 Babel plugin를 이용할 수 있다

```js
import AnchoredHeading from './AnchoredHeading.vue'

new Vue({
  el: '#demo',
  render (h) {
    return (
      <AnchoredHeading level={1}>
        <span>Hello</span> world!
      </AnchoredHeading>
    )
  }
})
```

`createElement`를 별칭 `h`로 이용하는 것은 Vue 생태계의 convention이다.
사용하는 범위에서 h를 사용할 수 없다면, 앱은 오류를 발생시킨다.

### 함수형 컴포넌트

앞에 작성한 anchored heading component는 props를 가진 단순 기능일 뿐이다
어떤 상태도 없고 전달된 상태를 감시하며 라이프사이클 관련 메소드도 없다.

이와 같은 경우, 컴포넌트를 함수형으로 표현할 수 있다.
즉, 컴포넌트가 상태가 없고(data 없음) 인스턴스화 되지 않은 경우(this 컨텍스트가 없음)를 말한다.

함수형 컴포넌트는 다음과 같다

```js
Vue.component('my-component', {
  functional: true,
  // 인스턴스의 부족함을 보완하기 위해 2번째에 컨텍스트 인수가 제공됨
  render: function (createElement, context) {
    // ...
  },
  // Props는 선택사항
  props: {
    // ...
  }
})
```

> 주의 : 2.3.0 이전 버전에서, 함수형 컴포넌트에서 props을 받아들이려면 props 옵션이 필요하다. 2.3.0 이상에서는 props 옵션을 생략할 수 있으며, 컴포넌트 노드에서 발견된 모든 속성은 암시적으로 props으로 추출된다.

2.5.0+ 이후로, 싱글 파일 컴포넌트를 사용하는 경우, 템플릿 기반의 함수형 컴포넌트를 정의할 수 있다

```html
<template functional>
</template>
```

컴포넌트에서 필요한 모든 내용은 context로 전달된다

- props: 전달받은 props에 대한 객체
- children: VNode 자식의 배열
- slots: 슬롯 객체를 반환하는 함수
- data: 컴포넌트에 전달된 전체 데이터 객체
- parent: 상위 컴포넌트에 대한 참조
- listeners: (2.3.0+) 부모에게 등록된 이벤트 리스너를 가진 객체. `data.on`의 alias
- injections: (2.3.0+) inject 옵션을 사용하면 resolved injection을 가진다

`functional:true`를 추가한 후 anchored heading component의 렌더 함수를 업데이트 하는 것은
context 전달인자를 추가하고 `this.$slots.default`를 `context.children`으로 갱신한 다음
`this.level`을 `context.props.level`로 갱신하는 과정으로 이루어진다.

함수형 컴포넌트는 단지 함수일 뿐이므로 렌더링에 들어가는 비용이 적지만,
Vue 개발자 도구의 컴포넌트 트리에서 함수형 컴포넌트를 볼 수 없다.

함수형 컴포넌트는 래퍼 컴포넌트로도 매우 유용하다

```js
var EmptyList = { /* ... */ }
var TableList = { /* ... */ }
var OrderedList = { /* ... */ }
var UnorderedList = { /* ... */ }

Vue.component('smart-list', {
  functional: true,
  render: function (createElement, context) {
    function appropriateListComponent () {
      var items = context.props.items

      if (items.length === 0)           return EmptyList
      if (typeof items[0] === 'object') return TableList
      if (context.props.isOrdered)      return OrderedList

      return UnorderedList
    }

    return createElement(
      appropriateListComponent(),
      context.data,
      context.children
    )
  },
  props: {
    items: {
      type: Array,
      required: true
    },
    isOrdered: Boolean
  }
})
```

#### slots() vs children

어떤 경우에는 `slots().default`는 `children`과 같다.

하지만 아래의 컴포넌트의 경우 `children`은 두 개의 단락을 제공할 것이고 `slots().default`는 오직 두 번째 단락을 반환한다.

따라서 `children`과 `slots()`을 모두 사용하여 컴포넌트가 슬롯 시스템에 대해 알고 있는지 확인하거나,
단순하게 `children`을 전달하여 다른 컴포넌트에 책임을 위임할지 선택할 수 있다.

```html
<my-functional-component>
  <p slot="foo">
    first
  </p>
  <p>second</p>
</my-functional-component>
```
