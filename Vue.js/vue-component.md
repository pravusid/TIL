# Vue.js Component

<https://kr.vuejs.org/v2/guide/components.html>

Vue 생성자는 미리 정의 된 옵션으로 재사용 가능한 컴포넌트 생성자를 생성하도록 확장 가능하다.

```js
var MyComponent = Vue.extend({
  // 옵션 확장
})
// `MyComponent`의 모든 인스턴스는
// 미리 정의된 확장 옵션과 함께 생성됩니다.
var myComponentInstance = new MyComponent()
```
