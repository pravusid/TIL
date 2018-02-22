# Vue.js 템플릿

Vue.js는 렌더링된 DOM을 기본 Vue 인스턴스 데이터에 선언적으로 바인딩 할 수있는 HTML 기반 템플릿 구문을 사용.
내부적으로 Vue는 템플릿을 가상 DOM 렌더링 함수로 컴파일 함. ( render() 함수 )

## 보간법

데이터 바인딩의 가장 기본 형태는 “Mustache” 구문(이중 중괄호)

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

디렉티브는 v- 접두사가 있는 특수 속성이다. v-for를 제외한 디렉티브 속성 값은 단일 JavaScript 표현식이 된다. 디렉티브로 연결된 표현식의 값이 변경되면, 그 값을 반응적으로 DOM에 적용한다.

### 전달인자 (v-bind, v-on)

일부 HTML 태그에 대응되는 디렉티브는 “전달인자”를 사용할 수 있다. `<a v-bind:href="url"></a>`

DOM 이벤트를 수신하는 v-on 디렉티브 예시이다. `<a v-on:click="doSomething">`

#### 약어

v-bind 약어 (:)

```html
<!-- 전체 구문 -->
<a v-bind:href="url"></a>
<!-- 약어 -->
<a :href="url"></a>
```

v-on 약어 (@)

```html
<!-- 전체 구문 -->
<a v-on:click="doSomething"></a>
<!-- 약어 -->
<a @click="doSomething"></a>
```

### v-if

속성 값이 true이면 디렉티브 하위값을 출력한다. `<p v-if="seen">이제 나를 볼 수 있어요</p>`

### 수식어

수식어는 점으로 표시되는 특수 접미사로, 디렉티브의 바인딩 세부방식을 나타낸다. 예를 들어, .prevent 수식어는 event.preventDefault()를 호출하도록 v-on 디렉티브에게 알려준다.

`<form v-on:submit.prevent="onSubmit"></form>`

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

```html
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
