# Vue.js 반응형에 대해 깊이 알아보기

Vue의 가장 큰 특징 중 하나는 눈에는 보이지 않는 반응형 시스템이다.
모델은 Plain JavaScript 객체이고 수정하면 화면이 갱신된다.
상태 관리를 간단하고 직관적으로 만들어주지만 common gotcha를 피하기 위해 어떻게 작동하는지 이해하는 것도 중요하다.

## 변경 내용을 추적하는 방법

Vue 인스턴스에 Plain JavaScript 객체를 data 옵션으로 전달하면,
Vue는 모든 프로퍼티를 `Object.defineProperty`를 사용하여 `getter / setter`로 변환한다.
이는 Vue가 ES5를 사용할 수 없는 IE8 이하를 지원하지 않는 이유이다.

`getter / setter`는 사용자에게는 보이지 않으나 프로퍼티에 액세스 하거나 수정할 때 Vue가 종속성 추적 및 변경 알림을 수행할 수 있게한다.
한 가지 주의사항은 변환된 데이터 객체가 기록될 때 브라우저가 `getter / setter` 형식을 다르게 처리하므로,
친숙한 인터페이스를 사용하기 위해 `vue-devtools`를 설치하는 것이 좋다.

모든 컴포넌트 인스턴스에는 컴포넌트가 종속적으로 렌더링되는 동안 “수정”된 모든 프로퍼티를 기록하는 watcher 인스턴스가 대응된다.
나중에 종속적인 `setter`가 트리거 되면 watcher에 알리고 컴포넌트가 다시 렌더링 된다.

![사이클](https://raw.githubusercontent.com/pravusid/TIL/master/Vue.js/img/reactivity.png)

## 변경 감지 경고

모던 JavaScript의 한계 (그리고 `Object.observe`의 포기)로 인해 Vue는 프로퍼티의 추가/제거를 감지할 수 없다.
Vue는 인스턴스 초기화 중에 `getter / setter` 변환 프로세스를 수행하기 때문에,
data 객체에 프로퍼티가 있어야 Vue가 이를 반응형으로 변환할 수 있다.

```js
var vm = new Vue({
  data: {
    a: 1
  }
})
// `vm.a` 은 이제 반응적입니다.

vm.b = 2
// `vm.b` 은 이제 반응적이지 않습니다.
```

Vue는 이미 만들어진 인스턴스에 새로운 루트 수준의 반응형 프로퍼티를 동적으로 추가하는 것을 허용하지 않는다.
그러나 `Vue.set(object, key, value)` 메소드를 사용하여 중첩된 객체에 반응성 속성을 추가 할 수 있다.

```js
Vue.set(vm.someObject, 'b', 2)
```

Vm.$set 인스턴스 메소드를 사용할 수도 있다. 이 메소드는 전역 `Vue.set`의 alias 이다.

```js
this.$set(this.someObject, 'b', 2)
```

때로는 `Object.assign()` 또는 `_.extend()`를 사용하여 기존 객체에 많은 프로퍼티를 할당 할 수 있다.
그러나 객체에 추가 된 프로퍼티는 변경 내용을 트리거하지 않습니다.
이 경우 원본 객체와 mixin 객체의 프로퍼티를 함께 사용하여 새 객체를 만든다.

```js
// `Object.assign(this.someObject, { a: 1, b: 2 })` 대신에
this.someObject = Object.assign({}, this.someObject, { a: 1, b: 2 })
```

## 반응형 속성 선언하기

Vue는 루트 수준의 반응형 프로퍼티를 동적으로 추가 할 수 없으므로,
모든 루트 수준의 반응형 데이터 프로퍼티를 빈 값으로라도 초기에 선언하여 Vue 인스턴스를 초기화해야 한다.

```js
var vm = new Vue({
  data: {
    // 빈 값으로 메시지를 선언 합니다.
    message: ''
  },
  template: '<div>{{ message }}</div>'
})
// 나중에 `message`를 설정합니다.
vm.message = 'Hello!'
```

data 옵션에 message를 선언하지 않으면 Vue는 렌더 함수가 존재하지 않는 프로퍼티에 접근하려고 한다는 경고를 한다.

이 제한 사항에는 기술적인 이유가 있다.
종속성 추적 시스템에서 edge cases 부류를 제거하고 Vue 인스턴스를 유형 검사 시스템으로 더 멋지게 만든다.
그러나 코드 유지관리 측면에서도 중요한 고려 사항이 있다.
data 객체는 컴포넌트 상태에 대한 스키마와 같은 역할을 한다.
모든 반응형 프로퍼티를 미리 선언하면 나중에 다시 방문하거나 다른 개발자가 읽을 때 구성 요소 코드를 더 쉽게 이해할 수 있다.

## 비동기 갱신 큐

Vue는 DOM 업데이트를 비동기로 한다.
데이터 변경이 발견 될 때마다 큐를 열고 같은 이벤트 루프에서 발생하는 모든 데이터 변경을 버퍼에 담는다.
같은 Watcher가 여러 번 발생하면 대기열에서 한 번만 푸시된다. 버퍼링을 통한 중복의 제거는 불필요한 계산과 DOM 조작을 피하는 데 있어 중요하다.
그 다음, 이벤트 루프 “tick”에서 Vue는 대기열을 비우고 실제 (이미 중복이 제거된) 작업을 수행한다.
내부적으로 Vue는 비동기 큐를 위해 네이티브 `Promise.then`과 `MessageChannel`를 시도하고 `setTimeout(fn, 0)`으로 돌아간다.

예를 들어, `vm.someData = 'new value'`를 설정하면, 컴포넌트는 즉시 재렌더링되지 않는다.
큐가 flush 될 때 다음 “tick” 에서 업데이트 된다.
대개의 경우 이 작업을 신경 쓸 필요는 없지만 업데이트 후 DOM 상태에 의존하는 작업을 수행하려는 경우 까다로울 수 있다.
Vue.js는 일반적으로 개발자가 “데이터 중심” 방식으로 생각하고 DOM을 직접 만지지 않도록 권장하지만 때로는 직접 건드려야 하는 경우도 있다.
Vue.js가 데이터 변경 후 DOM 업데이트를 마칠 때까지 기다리려면 데이터가 변경된 직후에 `Vue.nextTick(callback))`을 사용할 수 있다.
콜백은 DOM이 업데이트 된 후에 호출된다.

```html
<div id="example">{{ message }}</div>
```

```js
var vm = new Vue({
  el: '#example',
  data: {
    message: '123'
  }
})
vm.message = 'new message' // 데이터  변경
vm.$el.textContent === 'new message' // false
Vue.nextTick(function () {
  vm.$el.textContent === 'new message' // true
})
```

또한 `vm.$nextTick()` 인스턴스 메소드가 있다.
이는 내부 컴포넌트들에 특히 유용한데, 전역 Vue가 필요없고 콜백의 this 컨텍스트가 자동으로 현재 Vue 인스턴스에 바인드될 것이기 때문이다.

```js
Vue.component('example', {
  template: '<span>{{ message }}</span>',
  data: function () {
    return {
      message: '갱신 안됨'
    }
  },
  methods: {
    updateMessage: function () {
      this.message = '갱신됨'
      console.log(this.$el.textContent) // => '갱신 안됨'
      this.$nextTick(function () {
        console.log(this.$el.textContent) // => '갱신됨'
      })
    }
  }
})
```
