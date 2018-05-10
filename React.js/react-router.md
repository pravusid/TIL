# React-router v4

## 소개

React는 view만 담당하는 라이브러리이다. 주소에 따라 다른 view를 연결하기 위해서 react-router 라이브러리를 사용한다.

## 사용하기

### 설치

`yarn add react-router-dom` : 브라우저에서 사용되는 react-router

### import

`src/App.js`

```js
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Home from './routes/Home';
import About from './routes/About';
import Posts from './routes/Posts';
import NotFound from './routes/NotFound';

import Header from './components/Header';

const App = () => {
  return (
    <Router>
      <div>
        <Header/>>
        <Switch>
          <Route exact path="/" component={Home}/>
          <Route path="/about/:username" component={About}/>
          <Route path="/posts" component={Posts}>
          <Route component={NotFound}>
        </Switch>
      </div>
    </Router>
  );
};
export default App;
```

### route로 연결되는 components

1. `src/routes` 폴더 생성
1. `src/routes/Home.js` 컴포넌트 생성 / rsc 생성
1. `src/routes/About.js` 컴포넌트 생성 / rsc 생성

  ```js
  import React from 'react';

  const About = ({match}) => {
    return (
      <div>
        {match.params.username} 소개
      </div>
    );
  };
  export defaut About
  ```

### components

1. `src/components/Header.css` 생성
1. `src/components/Header.js` 생성

  ```js
  import React from 'react';
  import { Link } from 'react-router-dom';
  import './Header.css';

  const Header = () => {
    return (
      <div className="header">
        <Link to="/" className="item">홈</Link>
        <Link to="/about/pravusid" className="item">소개</Link>
        <Link to="/posts" className="item">포스트</Link>
      </div>
    );
  };
  export default Header;
  ```

### 라우트 내부의 라우트

`src/routes/Posts.js`

```js
import React from 'react';
import { Route, Link } from 'react-router-dom';

// 내부 컴포넌트
const Post = ({match}) => {
  return (
    <div>
      <h2>
        {match.params.title}
      </h2>
    </div>
  );
};

const Posts = () => {
  return (
    <div>
      <h1>포스트</h1>
      <Link to="/posts/react">React</Link>
      <Link to="/posts/redux">Redux</Link>
      <Route path="/posts/:title" component={Post}/>>
    </div>
  );
};

export default Posts;
```

### 선택된 링크만 스타일 활성화

1. `Link` => `NavLink`
1. `.item.active` => `.item.active, .item.active`
1. `activeClassName="active"`

### Redirect

- `src/routes/Login.js` 생성

- `src/routes/MyPage.js` 생성

  ```js
  import { Redirect } from 'react-router-dom';

    return (
      <div>
        { !logged && <Redirect to="/login"/> }
      </div>
    );
  ```

### Query Parameter

`src/routes/Search.js`

```js
import React from 'react';
const Search = ({location}) => {
  return (
    <div>
      { new URLSearchParams(location.search).get('keyword')} 검색
    </div>
  );
};
export default Search;
```
