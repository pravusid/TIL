# Vue Router

Vue 라우터는 아래의 기능을 포함하고 있다

- 중첩된 라우트/뷰 매핑
- 모듈화된, 컴포넌트 기반의 라우터 설정
- 라우터 파라미터, 쿼리, 와일드카드
- Vue.js의 트랜지션 시스템을 이용한 트랜지션 효과
- 세밀한 네비게이션 컨트롤
- active CSS 클래스를 자동으로 추가해주는 링크
- HTML5 히스토리 모드 또는 해시 모드 (IE9에서 자동으로 폴백)
- 사용자 정의 가능한 스크롤 동작

## 시작하기

- router-link 컴포넌트를 사용하여 네비게이션 적용
- 구체적인 속성은 `to` prop을 이용함
- 기본적으로 `<router-link>`는 `<a>` 태그로 렌더링된다

```html
<div id="app">
  <h1>Hello App!</h1>
  <p>
    <router-link to="/foo">Go to Foo</router-link>
    <router-link to="/bar">Go to Bar</router-link>
  </p>
  <router-view>
    <!-- 현재 라우트에 맞는 컴포넌트가 렌더링되는 영역. -->
  </router-view>
</div>
```

`main.js`

```js
import router from './routes';

new Vue({
  // ...
  router,
  // ...
}).$mount('#app');
```

`routes/index.js`

```js
import Vue from 'vue';
import Router from 'vue-router';

import HelloWorld from '../components/HelloWorld.vue';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld,
    },
  ],
});

export default router;

```

`this.$router`는 정확히 `router`와 동일하다.
`this.$router`를 사용하는 이유는 라우터를 조작해야하는 할때 라우터 객체를 가져오지 않고 접근하기 위함이다.

`<router-link>`는 현재 라우트와 일치할 때 자동으로 `.router-link-active` 클래스가 추가된다.

```js
// Home.vue
export default {
  computed: {
    username () {
      return this.$route.params.username
    }
  },
  methods: {
    goBack () {
      window.history.length > 1 ? this.$router.go(-1) : this.$router.push('/')
    }
  }
}
```

## 동적 라우트 매칭

동일 컴포넌트에 패턴으로 주어진 라우트를 매칭해야 하는 경우가 있다.
이 경우 경로에서 동적 세그먼트를 사용하면 된다.

```js
const User = {
  template: '<div>User</div>'
}

const router = new VueRouter({
  routes: [
    // 동적 세그먼트는 앞에 콜론을 붙인다
    { path: '/user/:id', component: User }
  ]
})
```

이제 `/user/foo`와 `/user/bar`은 같은 경로에 매핑된다.

라우트가 일치하면 동적 세그먼트의 값은 모든 컴포넌트에서 `this.$route.params`로 표시된다.

```js
const User = {
  template: '<div>User {{ $route.params.id }}</div>'
}
```

| 패턴 | 일치하는 경로 | `$route.params` |
| --- | --- | --- |
| /user/:username | /user/evan | { username: 'evan' } |
| /user/:username/post/:post_id | /user/evan/post/123 | { username: 'evan', post_id: 123 } |

`$route` 객체는 `$route.params`, `$route.query`, `$route.hash` 등의 정보를 제공한다.

### Params 변경 사항에 반응하기

Params와 함께 라우트를 사용하면 `/user/foo`에서 `/user/bar`로 이동할 때 컴포넌트 인스턴스가 재사용 되므로 주의해야 한다.

즉, 컴포넌트의 라이프 사이클 훅이 호출되지 않음을 의미한다.
동일한 컴포넌트의 params 변경 사항에 반응하려면 `$route` 객체를 감시하면 된다.

```js
const User = {
  template: '...',
  watch: {
    '$route' (to, from) {
      // 경로 변경에 반응하여...
    }
  }
}
```

버전 2.2에서 소개된 beforeRouteUpdate 가드를 사용할 수도 있다

```js
const User = {
  template: '...',
  beforeRouteUpdate (to, from, next) {
    // react to route changes...
    next();
  }
}
```

### 고급 매칭 패턴

vue-router는 라우트 매칭 엔진으로 path-to-regexp를 사용한다.
따라서 선택적 동적 세그먼트, 0개 이상의 요구 사항, 커스텀 정규식 패턴과 같은 여러 고급 매칭 패턴을 지원한다.

### 매칭 우선순위

동일한 URL이 여러 라우트와 일치하는 경우 일치하는 라우트 정의의 순서에 따라 우선 순위가 결정된다.
즉, 경로가 더 먼저 정의 될수록 우선 순위가 높아진다.

## 중첩된 라우트

vue-router를 통해 중첩된 라우트 구성을 사용하여 간단히 관계를 표현할 수 있다.

```text
/user/foo/profile                     /user/foo/posts
+------------------+                  +-----------------+
| User             |                  | User            |
| +--------------+ |                  | +-------------+ |
| | Profile      | |  +------------>  | | Posts       | |
| |              | |                  | |             | |
| +--------------+ |                  | +-------------+ |
+------------------+                  +-----------------+
```

위와 같은 관계에 대응하는 코드는 다음과 같다

```html
<div id="app">
  <p>
    <router-link to="/user/foo">/user/foo</router-link>
    <router-link to="/user/foo/profile">/user/foo/profile</router-link>
    <router-link to="/user/foo/posts">/user/foo/posts</router-link>
  </p>
  <router-view></router-view>
</div>
```

중첩 outlet에 컴포넌트를 렌더링하려면 children을 사용한다.

```js
const router = new VueRouter({
  routes: [
    { path: '/user/:id', component: User,
      children: [
        // UserHome will be rendered inside User's <router-view>
        // when /user/:id is matched
        { path: '', component: UserHome },

        // UserProfile will be rendered inside User's <router-view>
        // when /user/:id/profile is matched
        { path: 'profile', component: UserProfile },

        // UserPosts will be rendered inside User's <router-view>
        // when /user/:id/posts is matched
        { path: 'posts', component: UserPosts }
      ]
    }
  ]
})
```

## 프로그래밍 방식 네비게이션

### router.push(location, onComplete?, onAbort?)

다른 URL로 이동하려면 router.push를 사용한다

> Vue 인스턴스 내부에서 라우터 인스턴스에 $router로 액세스 할수 있다. (`this.$router.push`)

`push`는 새로운 항목을 히스토리 스택에 넣기 때문에 사용자가 브라우저의 뒤로 가기 버튼을 클릭하면 이전 URL로 이동하게된다.

`<router-link :to="...">`를 클릭하면 router.push(...)를 호출하는 것과 같다.

```js
// 리터럴 string
router.push('home')

// object
router.push({ path: 'home' })

// 이름을 가지는 라우트
router.push({ name: 'user', params: { userId: 123 }})

// 쿼리와 함께 사용, 결과는 /register?plan=private
router.push({ path: 'register', query: { plan: 'private' }})
```

2.2 버전이후 선택적으로 `router.push` 또는 `router.replace`에 두번째와 세번째 전달인자로 `onComplete`와 `onAbort` 콜백을 제공한다.
콜백은 탐색이 성공적으로 완료되거나(모든 비동기 훅이 해결된 후) 또는 중단(현재 탐색이 완료되기 전에 동일한 경로로 이동하거나 다른 경로 이동)될 때 호출된다.

### router.replace(location)

`router.push`와 같은 역할을 하지만 새로운 히스토리 항목에 추가하지 않고 탐색한다(현재 항목을 대체함))

`<router-link :to="..." replace>`: `router.replace(...)`

### router.go(n)

`window.history.go(n)`와 비슷하게 히스토리 스택에서 정수를 매개 변수로 앞 또는 뒤로 이동하는 단계를 나타낸다.

```js
// 한 단계 앞으로 간다. history.forward()와 같다.
router.go(1)

// 한 단계 뒤로 감. history.back()와 같다
router.go(-1)

// 3 단계 앞으로 간다.
router.go(3)

// 지정한 만큼의 기록이 없으면 자동으로 실패
router.go(-100)
router.go(100)
```

### History 조작

router.push, router.replace 및 router.go는
window.history.pushState,window.history.replaceState 및 window.history.go를 모방한 형태이다.

vue-router 네비게이션 메소드(push,replace,go)는 모든 라우터 모드(history,hash 및abstract)에서 일관되게 작동한다.

## 이름을 가지는 라우트

Router 인스턴스를 생성하는 동안 routes 옵션에 라우트 이름을 지정할 수 있다.

```js
const router = new VueRouter({
  routes: [
    {
      path: '/user/:userId',
      name: 'user',
      component: User
    }
  ]
})
```

이름을 가진 라우트에 링크하려면, 객체를 `router-link`, 컴포넌트의 `to prop`로 전달할 수 있다.

- `<router-link :to="{ name: 'user', params: { userId: 123 }}">User</router-link>`
- `router.push({ name: 'user', params: { userId: 123 }})`

두 경우 모두 라우터는 `/user/123` 경로로 이동한다.

## 이름을 가지는 뷰

때로는 여러 개의 뷰를 중첩하지 않고 동시에 표시해야 하는 경우가 있다.

이름이 없는 router-view는 이름으로 default가 주어진다.

```html
<router-view class="view one"></router-view>
<router-view class="view two" name="a"></router-view>
<router-view class="view three" name="b"></router-view>
```

이 경우 동일한 라우트에 대해 여러 컴포넌트가 필요하다.

```js
const router = new VueRouter({
  routes: [
    {
      path: '/',
      components: {
        default: Foo,
        a: Bar,
        b: Baz
      }
    }
  ]
})
```

## 리다이렉트와 별칭

### 리다이렉트

/a에서 /b로 리디렉션하려면

```js
const router = new VueRouter({
  routes: [
    { path: '/a', redirect: '/b' }
  ]
})
```

리디렉션은 이름이 지정된 라우트를 지정할 수도 있다

```js
const router = new VueRouter({
  routes: [
    { path: '/a', redirect: { name: 'foo' }}
  ]
})
```

동적 리디렉션을 위한 함수를 사용할 수도 있다

```js
const router = new VueRouter({
  routes: [
    { path: '/a', redirect: to => {
      // 함수는 인수로 대상 라우트를 받는다
      // 여기서 path/location 반환함
    }}
  ]
})
```

### 별칭

/a의 별칭은 /b는 사용자가 /b를 방문했을 때 URL은 /b을 유지하지만 사용자가 /a를 방문한 것처럼 매칭한다

```js
const router = new VueRouter({
  routes: [
    { path: '/a', component: A, alias: '/b' }
  ]
})
```

## 라우트 컴포넌트에 속성 전달

컴포넌트에서 `$route`를 사용하면 특정 URL에 의존하는 컴포넌트와 라우트와 강결합이 발생한다.

컴포넌트와 라우터 속성을 분리하려면 다음과 같이 할 수 있다.

$route에 의존성 추가

```js
const User = {
  template: '<div>User {{ $route.params.id }}</div>'
}
const router = new VueRouter({
  routes: [
    { path: '/user/:id', component: User }
  ]
})
```

속성에 의존성 해제

```js
const User = {
  props: ['id'],
  template: '<div>User {{ id }}</div>'
}
const router = new VueRouter({
  routes: [
    { path: '/user/:id', component: User, props: true },
  ]
})
```

이를 통해 컴포넌트 재사용 및 테스트가 용이해진다.

### Boolean 모드

라우터 props 옵션을 true로 설정하면 route.params가 컴포넌트 props로 설정된다.

### 객체 모드

라우터 props옵션이 객체일때 컴포넌트 props에 그대로 전달된다. props가 정적일 때 유용하다.

```js
const router = new VueRouter({
  routes: [
    { path: '/promotion/from-newsletter', component: Promotion, props: { newsletterPopup: false } }
  ]
})
```

### 함수 모드

props를 반환하는 함수를 만들 수 있다.
라우트 변경시에만 평가되므로 props 함수는 상태를 저장하지 않아야 한다.

```js
const router = new VueRouter({
  routes: [
    { path: '/search', component: SearchUser, props: (route) => ({ query: route.query.q }) }
  ]
})
// /search?q=vue는 {query: "vue"}를 SearchUser 컴포넌트에 전달한다.
```

## HTML5 히스토리 모드

vue-router의 기본 모드는 `hash mode` 이다.
URL 해시를 사용하여 전체 URL을 시뮬레이션하므로 URL이 변경될 때 페이지가 새로고침 되지 않는다.

해시모드 대신 라우터의 history 모드 를 사용할 수 있다.
`history.pushState` API를 활용하여 페이지를 다시 로드하지 않고도 URL 탐색을 할 수 있다.

```js
const router = new VueRouter({
  mode: 'history',
  routes: [...]
})
```

히스토리 모드를 사용할 때 SPA의 경우 사용자가 직접 `http://oursite.com/user/id` 에 접속하면 404 오류가 발생한다.
문제를 해결하려면 서버에 간단하게 포괄적인 대체 경로를 추가하여 SPA에 접근하게 해야 한다.

### 서버 설정 예제

#### Apache

```text
<IfModule mod_rewrite.c>
  RewriteEngine On
  RewriteBase /
  RewriteRule ^index\.html$ - [L]
  RewriteCond %{REQUEST_FILENAME} !-f
  RewriteCond %{REQUEST_FILENAME} !-d
  RewriteRule . /index.html [L]
</IfModule>
```

#### nginx

```text
location / {
  try_files $uri $uri/ /index.html;
}
```

#### Native Node.js

```js
const http = require("http")
const fs = require("fs")
const httpPort = 80

http.createServer((req, res) => {
  fs.readFile("index.htm", "utf-8", (err, content) => {
    if (err) {
      console.log('We cannot open "index.htm" file.')
    }

    res.writeHead(200, {
      "Content-Type": "text/html; charset=utf-8"
    })

    res.end(content)
  })
}).listen(httpPort, () => {
  console.log("Server listening on: http://localhost:%s", httpPort)
})
```

#### Express와 Node.js

Express의 경우 `connect-history-api-fallback` 미들웨어를 고려할 수 있다.

#### Internet Information Services (IIS)

```xml
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
 <system.webServer>
   <rewrite>
     <rules>
       <rule name="Handle History Mode and custom 404/500" stopProcessing="true">
         <match url="(.*)" />
         <conditions logicalGrouping="MatchAll">
           <add input="{REQUEST_FILENAME}" matchType="IsFile" negate="true" />
           <add input="{REQUEST_FILENAME}" matchType="IsDirectory" negate="true" />
         </conditions>
         <action type="Rewrite" url="index.html" />
       </rule>
     </rules>
   </rewrite>
     <httpErrors>
         <remove statusCode="404" subStatusCode="-1" />
         <remove statusCode="500" subStatusCode="-1" />
         <error statusCode="404" path="/survey/notfound" responseMode="ExecuteURL" />
         <error statusCode="500" path="/survey/error" responseMode="ExecuteURL" />
     </httpErrors>
     <modules runAllManagedModulesForAllRequests="true"/>
 </system.webServer>
</configuration>
```

#### 주의 사항

모든 발견되지 않은 경로가 이제 index.html 파일을 제공하므로, 실제로 존재하지 않는 경로접근에도 404에러를 반환하지 않는다.
이 문제를 해결하려면 Vue 앱에서 catch-all 라우트를 구현하여 404 페이지를 표시해야한다.

```js
const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '*', component: NotFoundComponent }
  ]
})
```

## 네비게이션 가드

라우트 탐색 프로세스에 연결하는 방법에는 전역, 라우트별 또는 컴포넌트가 있다.

Params 또는 쿼리를 변경시에는 네비게이션 가드가 실행되지 않는다.
이경우 `$route` 객체를 감시하거나 컴포넌트 가드의 `beforeRouteUpdate`를 사용하여야 한다.

모든 가드 기능은 세 가지 전달인자를 받는다

- `to`: 라우트: 이동할 대상 라우트 객체
- `from`: 라우트: 현재 라우트로 오기 전 라우트객체
- `next()`: 파이프라인의 다음 훅으로 이동한다 (항상 next 함수를 호출해야 한다)
  - `next(false)`: 현재 네비게이션을 중단, 브라우저 URL이 변경되면(수동으로 변경) from경로의 URL로 재설정된다.
  - `next('/')` / `next({ path: '/' })`: 다른 위치로 리디렉션, 현재 네비게이션이 중단되고 새 네비게이션이 시작된다.
  - `next(error)`: (2.4.0 이후) next에 전달된 인자가 Error의 인스턴스이면 탐색이 중단되고 `router.onError()`를 이용해 등록된 콜백에 에러가 전달됨

### 전역 가드

`router.beforeEach`를 사용하여 전역등록을 할 수 있다.

```js
const router = new VueRouter({ ... })

router.beforeEach((to, from, next) => {
  // ...
})
```

네비게이션이 트리거될 때마다 가드가 모든 경우에 실행된다.
가드는 비동기식으로 실행 될 수 있으며 네비게이션은 모든 훅이 해결되기 전까지 보류중 으로 간주된다.

### Global Resolve Guards

버전 2.5 이후로 `router.beforeResolve`를 사용하여 글로벌 가드를 등록 할 수 있으며, 이는 `router.beforeEach`와 유사하다.

### Global After Hooks

가드와 달리 이 훅은 next 함수를 인자로 받지 못해 네비게이션에 영향을 줄 수 없다

```js
router.afterEach((to, from) => {
  // ...
})
```

### 라우트 별 가드

`beforeEnter` 가드를 라우트의 설정 객체에 직접 정의 할 수 있다

```js
const router = new VueRouter({
  routes: [
    {
      path: '/foo',
      component: Foo,
      beforeEnter: (to, from, next) => {
        // ...
      }
    }
  ]
})
```

### 컴포넌트 내부 가드

다음을 사용하여 라우트 컴포넌트안에 라우트 네비게이션 가드를 직접 정의 할 수 있다

- `beforeRouteEnter`
- `beforeRouteUpdate`
- `beforeRouteLeave`

```js
const Foo = {
  template: `...`,
  beforeRouteEnter (to, from, next) {
    // 이 컴포넌트를 렌더링하는 라우트 앞에 호출된다
    // 이 가드가 호출 될 때 컴포넌트 인스턴스가 생성되지 않았기 때문에 `this`로 접근할 수 없다
  },
  beforeRouteLeave (to, from, next) {
    // 이 컴포넌트를 렌더링하는 라우트가 이전으로 네비게이션 될 때 호출된다
    // `this` 컴포넌트 인스턴스에 접근 할 수 있다
  }
}
```

`beforeRouteEnter` 가드는 네비게이션이 확인되기 전에 호출되어서 컴포넌트 인스턴스의 this에 접근하지 못한다.
그러나 콜백을 next에 전달하여 인스턴스에 액세스 할 수 있다.
네비게이션이 확인되고 컴포넌트 인스턴스가 콜백에 전달인자로 전달 될 때 콜백이 호출된다.

```js
beforeRouteEnter (to, from, next) {
  next(vm => {
    // `vm`을 통한 컴포넌트 인스턴스 접근
  })
}
```

`beforeRouteLeave` 안에서는 `this`에 직접 접근 할 수 있다.
leave 가드는 일반적으로 사용자가 저장하지 않은 편집 내용을 두고 실수로 라우트를 떠나는 것을 방지하는데 사용된다.
탐색은 `next(false)`를 호출하여 취소할 수 있다.

### 전체 네비게이션 시나리오

- 네비게이션이 트리거됨
- 비활성화될 컴포넌트에서 가드를 호출
- 전역 beforeEach 가드 호출
- 재사용되는 컴포넌트에서 beforeRouteUpdate 가드 호출 (2.2 이상)
- 라우트 설정에서 beforeEnter 호출
- 비동기 라우트 컴포넌트 해결
- 활성화된 컴포넌트에서 beforeRouteEnter 호출
- 전역 beforeResolve 가드 호출 (2.5이상)
- 네비게이션 완료
- 전역 afterEach 훅 호출
- DOM 갱신 트리거 됨
- 인스턴스화 된 인스턴스들의 beforeRouteEnter가드에서 next에 전달 된 콜백을 호출

## 라우트 메타 필드

라우트를 정의 할 때 meta 필드를 포함시킬 수 있다

```js
const router = new VueRouter({
  routes: [
    {
      path: '/foo',
      component: Foo,
      children: [
        {
          path: 'bar',
          component: Bar,
          // 메타 필드
          meta: { requiresAuth: true }
        }
      ]
    }
  ]
})
```

routes 설정의 각 라우트 객체를 라우트 레코드라고 한다.
라우트 레코드는 중첩 될 수 있으므로 라우트가 일치하면 둘 이상의 라우트 레코드와 잠재적으로 일치 할 수 있다.

예를 들어, 위의 라우트 구성에서 URL `/foo/bar`는 상위 라우트 레코드와 하위 라우트 레코드 모두와 일치한다.

라우트와 일치하는 모든 라우트 레코드는 `$route` 객체(그리고 네비게이션 가드의 라우트 객체)에 `$route.matched` 배열로 노출된다.
따라서 `$route.matched`를 반복하여 라우트 레코드의 메타 필드를 검사 할 필요가 있다.

```js
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // 이 라우트는 인증여부를 확인하며 인증하지 않은 경우 로그인 페이지로 리디렉션한다
    if (!auth.loggedIn()) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    next() // 반드시 next()를 호출해야 한다
  }
})
```

## 트랜지션

`<router-view>`는 본질적으로 동적인 컴포넌트이기 때문에
`<transition>` 컴포넌트를 사용하는 것과 같은 방식으로 트랜지션 효과를 적용할 수 있다.

```html
<transition>
  <router-view></router-view>
</transition>
```

### 라우트 별 트랜지션

위의 사용법은 모든 라우트에 대해 동일한 트랜지션을 적용한다.
각 라우트의 컴포넌트가 서로 다른 트랜지션을 갖도록 하려면 각 라우트 컴포넌트 내에서 다른 이름으로 `<transition>`을 사용해야 한다.

```js
const Foo = {
  template: `
    <transition name="slide">
      <div class="foo">...</div>
    </transition>
  `
}

const Bar = {
  template: `
    <transition name="fade">
      <div class="bar">...</div>
    </transition>
  `
}
```

### 라우트 기반 동적 트랜지션

대상 라우트와 현재 라우트 간의 관계를 기반으로 동적으로 사용할 트랜지션을 결정할 수도 있다

```html
<!-- 동적 트랜지션을 위한 name을 지정합니다. -->
<transition :name="transitionName">
  <router-view></router-view>
</transition>
```

부모 구성 요소에서 `$route`를 보고 사용할 트랜지션을 결정한다

```js
watch: {
  '$route' (to, from) {
    const toDepth = to.path.split('/').length
    const fromDepth = from.path.split('/').length
    this.transitionName = toDepth < fromDepth ? 'slide-right' : 'slide-left'
  }
}
```

## 데이터 가져오기

라우트가 활성화될 때 서버에서 데이터를 가져와야 하는 경우 두 가지 방법을 사용할 수 있다.

- 탐색 후 가져오기: 먼저 탐색하고 컴포넌트의 라이프 사이클 훅에서 데이터를 가져오며, 데이터를 가져오는 동안 로드 상태를 표시한다
- 탐색하기 전에 가져오기: 라우트 가드에서 경로를 탐색하기 전에 데이터를 가져오고 그 후에 탐색을 수행한다

### 탐색 후 가져오기

들어오는 컴포넌트를 즉시 탐색하고 렌더링하며 컴포넌트의 `created` 훅에서 데이터를 가져오게된다.
네트워크를 통해 데이터를 가져 오는 동안 로드 상태를 표시 할 수 있다.

```js
<template>
  <div class="post">
    <div class="loading" v-if="loading">
      Loading...
    </div>

    <div v-if="error" class="error">
      {{ error }}
    </div>

    <div v-if="post" class="content">
      <h2>{{ post.title }}</h2>
      <p>{{ post.body }}</p>
    </div>
  </div>
</template>

export default {
  data () {
    return {
      loading: false,
      post: null,
      error: null
    }
  },
  created () {
    // 뷰가 생성되고 데이터가 이미 감시 되고 있을 때 데이터를 가져온다
    this.fetchData()
  },
  watch: {
    // 라우트가 변경되면 메소드를 다시 호출한다
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      this.error = this.post = null
      this.loading = true
      // `getPost`를 데이터 가져오기 위한 유틸리티/API 래퍼로 변경한다
      getPost(this.$route.params.id, (err, post) => {
        this.loading = false
        if (err) {
          this.error = err.toString()
        } else {
          this.post = post
        }
      })
    }
  }
}
```

### 탐색하기 전에 가져오기

이 접근 방식을 사용하면 새 경로로 이동하기 전에 데이터를 가져온다.
들어오는 컴포넌트에서 `beforeRouteEnter` 가드에서 데이터를 가져올 수 있으며 fetch가 완료되면 `next`를 호출하면 된다.

```js
export default {
  data () {
    return {
      post: null,
      error: null
    }
  },
  beforeRouteEnter (to, from, next) {
    getPost(to.params.id, (err, post) => {
      next(vm => vm.setData(err, post))
    })
  },
  watch: {
    // 라우트가 변경되면 메소드를 다시 호출한다
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      this.error = this.post = null
      this.loading = true
      // `getPost`를 데이터 fetch 유틸리티/API 래퍼로 바꾼다
      getPost(this.$route.params.id, (err, post) => {
        this.loading = false
        if (err) {
          this.error = err.toString()
        } else {
          this.post = post
        }
      })
    }
  }
}
```

- 다음 뷰에 대한 리소스를 가져 오는 동안 사용자는 현재 뷰를 유지한다.
- 따라서 데이터를 가져 오는 동안 진행률을 표시하는 것이 좋다.
- 데이터 가져오기가 실패하면 일종의 메시지를 표시해야 한다.

## 스크롤 동작 (HTML5 히스토리 모드에서만 작동)

클라이언트 측 라우팅으로 새로운 경로로 이동할 때, 라우트 탐색에서 사용자 정의 스크롤 동작을 할 수 있다.

라우터 인스턴스를 생성 할 때 scrollBehavior 함수를 정의할 수 있다

```js
const router = new VueRouter({
  routes: [...],
  scrollBehavior (to, from, savedPosition) {
    // 원하는 위치로 돌아가기
  }
})
```

scrollBehavior 함수는 `to`와 `from` 라우트 객체를 받는다.
`savedPosition` 인자는 브라우저의 뒤/앞으로 버튼으로 트리거되는 `popstate` 네비게이션인 경우에만 사용할 수 있다.

함수가 반환하는 스크롤 위치 객체는 다음과 같은 형태이다

- `{ x: number, y: number }`
- `{ selector: string, offset? : { x: number, y: number }}` (offset은 2.6 이상만 지원)

```js
// 라우트 네비게이션에 대해 페이지가 맨 위로 스크롤된다
scrollBehavior (to, from, savedPosition) {
  return { x: 0, y: 0 }
}
```

`savedPosition`을 반환하면 뒤로/앞으로 버튼으로 탐색할 때 네이티브와 같은 동작이 발생한다

```js
scrollBehavior (to, from, savedPosition) {
  if (savedPosition) {
    return savedPosition
  } else {
    return { x: 0, y: 0 }
  }
}
```

**anchor로 스크롤** 동작을 시뮬레이트하려면 다음과 같다

```js
scrollBehavior (to, from, savedPosition) {
  if (to.hash) {
    return {
      selector: to.hash
      // , offset: { x: 0, y: 10 }
    }
  }
}
```

## 지연된 로딩

번들러를 이용하여 앱을 제작할 때 JavaScript 번들이 상당히 커져 페이지로드 시간에 영향을 줄 수 있다.
이 때, 각 라우트의 컴포넌트를 별도의 단위로 분할하고 경로를 방문할 때 로드하는 것이 효율적이다.

Vue의 비동기 컴포넌트와 Webpack의 코드 분할을 함께 사용한다.

- 비동기 컴포넌트는 Promise를 반환하는 팩토리 함수로 정의할 수 있다 (컴포넌트가 resolve 되어야함)
  - `const Foo = () => Promise.resolve({ /* 컴포넌트 정의 */ })`
- Webpack에서 dynamic import를 사용하여 코드 분할 포인트를 지정할 수 있다
  - `import('./Foo.vue') // returns a Promise`

```js
const Foo = () => import('./Foo.vue')

// 라우터 설정에서 변경점은 없음
const router = new VueRouter({
  routes: [
    { path: '/foo', component: Foo }
  ]
})
```

### 같은 묶음으로 컴포넌트 그룹화하기

동일한 라우트 아래에 중첩된 모든 컴포넌트를 동일한 비동기 묶음으로 그룹화 할 수 있다.
이름을 가진 묶음이라는 특수 주석 문법으로 이름을 제공하면 된다.

```js
const Foo = () => import(/* webpackChunkName: "group-foo" */ './Foo.vue')
const Bar = () => import(/* webpackChunkName: "group-foo" */ './Bar.vue')
const Baz = () => import(/* webpackChunkName: "group-foo" */ './Baz.vue')
```
