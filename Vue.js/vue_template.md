# Vue.js 템플릿

내부적으로 Vue는 템플릿을 가상 DOM 렌더링 함수로 컴파일 합니다

## 보간법

데이터 바인딩의 가장 기본 형태는 “Mustache” 구문(이중 중괄호)
```html
<span>메시지: {{ msg }}</span>
```

HTML 출력을 위해서는 v-html 디렉티브 사용
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

디렉티브는 v- 접두사가 있는 특수 속성입니다. 디렉티브 속성 값은 단일 JavaScript 표현식 이 됩니다. (나중에 설명할 v-for는 예외입니다.) 디렉티브의 역할은 표현식의 값이 변경될 때 사이드이펙트를 반응적으로 DOM에 적용하는 것

### v-if

`<p v-if="seen">이제 나를 볼 수 있어요</p>`

### 전달인자

일부 디렉티브는 콜론으로 표시되는 “전달인자”를 사용할 수 있습니다. 예를 들어, v-bind 디렉티브는 반응적으로 HTML 속성을 갱신하는데 사용됩니다.
`<a v-bind:href="url"></a>`

또 다른 예로 DOM 이벤트를 수신하는 v-on 디렉티브입니다.
`<a v-on:click="doSomething">`

### 수식어

수식어는 점으로 표시되는 특수 접미사로, 디렉티브를 특별한 방법으로 바인딩 해야 함을 나타냅니다. 예를 들어, .prevent 수식어는 트리거된 이벤트에서 event.preventDefault()를 호출하도록 v-on 디렉티브에게 알려줍니다.
`<form v-on:submit.prevent="onSubmit"></form>`

## 필터

Vue.js에서는 일반 텍스트 서식을 적용할 때 사용할 수 있는 필터를 정의할 수 있습니다. 필터는 Mustache 보간과 v-bind 표현식 두 곳에서 사용할 수 있습니다. 필터는 JavaScript 표현식의 끝에 추가해야 하며 “파이프” 기호로 표시됩니다.
```html
<!-- Mustaches 사용시 -->
{{ message | capitalize }}
<!-- v-bind 사용시 -->
<div v-bind:id="rawId | formatId"></div>
```

필터 함수는 항상 표현식의 값을 첫번째 전달 인자로 받습니다.
```html
new Vue({
  // ...
  filters: {
    capitalize: function (value) {
      if (!value) return ''
      value = value.toString()
      return value.charAt(0).toUpperCase() + value.slice(1)
    }
  }
})
```

필터는 체이닝 가능합니다.
`{{ message | filterA | filterB }}`

필터는 JavaScript 함수이므로 전달인자를 사용할 수 있습니다.
`{{ message | filterA('arg1', arg2) }}`

## 약어

v-bind 약어
```html
<!-- 전체 구문 -->
<a v-bind:href="url"></a>
<!-- 약어 -->
<a :href="url"></a>
```

v-on 약어
```html
<!-- 전체 구문 -->
<a v-on:click="doSomething"></a>
<!-- 약어 -->
<a @click="doSomething"></a>
```
