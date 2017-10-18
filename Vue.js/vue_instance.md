# Vue instance

## Vue 인스턴스 옵션

### Vue 인스턴스 생성자

모든 Vue vm은 Vue 생성자 함수로 root Vue 인스턴스를 생성하여 부트스트래핑된다.

```js
var vm = new Vue({
  // 옵션
})
```

Vue 생성자는 미리 정의 된 옵션으로 재사용 가능한 컴포넌트 생성자를 생성하도록 확장 가능하다.

```js
var MyComponent = Vue.extend({
  // 옵션 확장
})
// `MyComponent`의 모든 인스턴스는
// 미리 정의된 확장 옵션과 함께 생성됩니다.
var myComponentInstance = new MyComponent()
```

### data 객체

각각의 Vue 인스턴스는 data 객체의 모든 속성을 프록시 처리한다.
데이터 속성 외에도 유용한 인스턴스 속성 및 메소드가 있고, 이 프로퍼티들과 메소드들은 $ 접두사로 프록시 데이터 속성과 구별하여 호출가능하다.

```js
var data = { a: 1 };
var vm = new Vue({
  el: '#app',
  data: data
});

vm.a === data.a; // true
vm.$data === data; // true
vm.$el === document.getElementById('app'); // true
// $watch 는 인스턴스 메소드 입니다.
vm.$watch('a', function (newVal, oldVal) {
  // `vm.a`가 변경되면 호출 됩니다.
});
```

### computed (계산된 속성)

표현식은 유지보수성을 위해서 단순한 연산에만 사용하고, 복잡한 결과도출을 위해서는 계산된 속성을 사용해야 한다.

계산된 속성은 종속성(아래의 경우 message) 중 일부가 변경된 경우에만 다시 계산된다.
따라서 이후에 reversedMessage에 접근하면 캐싱된 값을 반환받는다. 만약 캐싱호출을 원하지 않는다면 reversedMessage() 메소드를 호출하면 된다.

```js
var vm = new Vue({
  el: '#app',
  data() {
    return {
      message: '안녕하세요'
    };
  },
  computed: {
    // 계산된 getter
    reversedMessage() {
      // this는 vm 인스턴스
      return this.message.split('').reverse().join('');
    }
  }
});
```

```html
<div id="app">
  <p>원본: {{ message }}</p>
  <p>계산된 값: {{ reversedMessage }}</p>
</div>
```

계산된 속성은 필요한 경우 setter를 사용할 수도 있다.

```js
// ...
  computed: {
    fullName: {
      get() {
        return this.firstName + ' ' + this.lastName;
      },
      set(newValue) {
        var names = newValue.split(' ');
        this.firstName = names[0];
        this.lastName = names[names.length - 1];
      }
    }
  }
// ...
vm.fullName = 'Sand Park';
// 실행하면 vm.firstName === 'Sand', vm.lastName === 'Park'
```

### watch (감시된 속성)

Vue는 인스턴스 데이터변경을 관찰하고 이에 반응하는 보다 일반적인 속성감시방법을 제공한다.
하지만 watch 콜백보다는 계산된 속성을 사용하는 것이 더 좋다.

```js
var vm = new Vue({
  data() {
    return {
      firstName: 'Foo',
      lastName: 'Bar',
      fullName: 'Foo Bar'
    };
  },
  watch: {
    firstName(val) {
      this.fullName = val + ' ' + this.lastName;
    },
    lastName(val) {
      this.fullName = this.firstName + ' ' + val;
    }
  }
});
```

감시된 속성을 사용하면 데이터 변화를 적용하기 위해 이미 존재하는 속성에 접근 (this) 해야 한다. 계산된 속성을 이용하면 코드를 줄일 수 있다.

```js
var vm = new Vue({
  data() {
    return {
      firstName: 'Foo',
      lastName: 'Bar'
    };
  },
  computed: {
    fullName() {
      return this.firstName + ' ' + this.lastName;
    }
  }
});
```

## [인스턴트 라이프사이클 훅](https://kr.vuejs.org/v2/guide/instance.html#%EB%9D%BC%EC%9D%B4%ED%94%84%EC%82%AC%EC%9D%B4%ED%81%B4-%EB%8B%A4%EC%9D%B4%EC%96%B4%EA%B7%B8%EB%9E%A8)

created, mounted, updated, destroyed ...

1. new Vue()
1. `beforeCreate()`
1. Initialize Data & Events
1. Instace created
1. `created()`
1. Compile template or el's template
1. `beforeMount()`
1. replace el with compiled template
1. Mounted to DOM
1. `mounted()`
1. dataChanged
1. `beforeUpdate()`
1. re-render DOM
1. `updated()`
1. `beforeDestory()`
1. destroyed
1. `destroyed()`

```js
var vm = new Vue({
  data() {
    return {
      a: 1
    };
  },
  created() {
    // `this` 는 vm 인스턴스를 가리킵니다.
    console.log('a is: ' + this.a)
  }
})
// -> "a is: 1"
```
