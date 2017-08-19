# Vue instance

## Vue 인스턴스 소개

  컴포넌트 사용 예시
  ```html
  <div id="app-7">
    <ol>
      <!--
        이제 각 todo-item 에 todo 객체를 제공합니다.
        화면에 나오므로, 각 항목의 컨텐츠는 동적으로 바뀔 수 있습니다.
        또한 각 구성 요소에 "키"를 제공해야합니다 (나중에 설명 됨).
       -->
      <todo-item
        v-for="item in groceryList"
        v-bind:todo="item"
        v-bind:key="item.id">
      </todo-item>
    </ol>
  </div>
  ```
  ```js
    Vue.component('todo-item', {
    props: ['todo'],
    template: '<li>{{ todo.text }}</li>'
  })
  var app7 = new Vue({
    el: '#app-7',
    data: {
      groceryList: [
        { id: 0, text: 'Vegetables' },
        { id: 1, text: 'Cheese' },
        { id: 2, text: 'Whatever else humans are supposed to eat' }
      ]
    }
  })
  ```

  The `<style>` tag is used to specify component CSS. By default, all style information is available application wide. If we want to scope CSS to just the current component, we can add the scoped attribute to the style tag like so - `<style scoped>`.

## Vue 인스턴스 생성자

모든 Vue vm은 Vue 생성자 함수로 root Vue 인스턴스를 생성하여 부트스트래핑됩니다.
  ```js
  var vm = new Vue({
    // 옵션
  })
  ```

Vue 생성자는 미리 정의 된 옵션으로 재사용 가능한 컴포넌트 생성자를 생성하도록 확장 될 수 있습니다
  ```js
  var MyComponent = Vue.extend({
    // 옵션 확장
  })
  // `MyComponent`의 모든 인스턴스는
  // 미리 정의된 확장 옵션과 함께 생성됩니다.
  var myComponentInstance = new MyComponent()
  ```

각 Vue 인스턴스는 data 객체에 있는 모든 속성을 프록시 처리 합니다.
  ```js
  var data = { a: 1 }
  var vm = new Vue({
    data: data
  })
  vm.a === data.a // -> true
  // 속성 설정은 원본 데이터에도 영향을 미칩니다.
  vm.a = 2
  data.a // -> 2
  ```

Vue 인스턴스는 데이터 속성 외에도 유용한 인스턴스 속성 및 메소드를 제공합니다. 이 프로퍼티들과 메소드들은 $ 접두사로 프록시 데이터 속성과 구별됩니다.
  ```js
  var data = { a: 1 }
  var vm = new Vue({
    el: '#example',
    data: data
  })
  vm.$data === data // -> true
  vm.$el === document.getElementById('example') // -> true
  // $watch 는 인스턴스 메소드 입니다.
  vm.$watch('a', function (newVal, oldVal) {
    // `vm.a`가 변경되면 호출 됩니다.
  })
  ```

## [인스턴트 라이프사이클 훅](https://kr.vuejs.org/v2/guide/instance.html#%EB%9D%BC%EC%9D%B4%ED%94%84%EC%82%AC%EC%9D%B4%ED%81%B4-%EB%8B%A4%EC%9D%B4%EC%96%B4%EA%B7%B8%EB%9E%A8)

created, mounted, updated, destroyed 가 존재
  ```js
  var vm = new Vue({
    data: {
      a: 1
    },
    created: function () {
      // `this` 는 vm 인스턴스를 가리킵니다.
      console.log('a is: ' + this.a)
    }
  })
  // -> "a is: 1"
  ```
