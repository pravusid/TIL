# React 입문

## React 소개

UI를 만들기 위한 자바스크립트 라이브러리 (View 영역)

## React 프로젝트 생성

- `sudo npm install -g webpack webpack-dev-server`
  - webpack : 브라우저 위에서 import를 할수있게 해주고 자바스크립트 파일을 합쳐준다
  - webpack-dev-server : static 파일 웹서버를 열수있고 hot-loader를 통해 코드가 수정 될 때마다 자동으로 리로드

- node project 생성 `npm init`
  - `package.json` : 프로젝트 정보

- React 설치 `npm install --save react react-dom`

- 개발 의존모듈 설치
  - `npm install --save-dev babel-core babel-loader babel-preset-es2015 babel-preset-react`
  - `npm insatll --save-dev react-hot-loader webpack webpack-dev-server`

- webpack 설정 : root에 `webpack.config.js`

- `/public/index.html` 파일 작성

- `/src/components/App.js` 파일 작성
  - `import React from 'react';`(es6) == `var React = require('react');`
  - `export default App;`(es6) == `module.export = App;`

- `/src/index.js` 파일 작성
  ```js
  import React from 'react';
  import ReactDOM from 'react-dom';
  import App from './components/App';
  const rootElement = document.getElementById('root');
  ReactDOM.render(<App/>, rootElement);
  ```

- 개발서버 실행스크립트 수정
  - `package.json` 파일의 `scripts`에 해당내용 추가 : `"dev-server": "webpack-dev-server"`

- react-hot-loader 적용 `/webpack.config.js`
  - 3.0 beta로 업데이트 `npm install --save-dev react-hot-loader@next`

  - `webpack.config.js` 수정
  ```json
  var webpack = require('webpack');

  module.exports = {
    entry: ['react-hot-loader/patch', './src/index.js'],

    output: {
      path: __dirname + '/public/',
      filename: 'bundle.js'
    },

    devServer: {
      hot: true,
      inline: true,
      host: '0.0.0.0',
      port: 4000,
      contentBase: __dirname + '/public/',
    },

    module: {
      loaders: [{
        test: /.js$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
        query: {
          cacheDirectory: true,
          presets: ['es2015', 'react'],
          plugins: ["react-hot-loader/babel"]
        }
      }]
    },

    plugins: [
      new webpack.HotModuleReplacementPlugin()
    ]
  };
  ```

  - `/src/index.js` 수정
  ```js
  import React from 'react';
  import ReactDOM from 'react-dom';
  import { AppContainer } from 'react-hot-loader';
  import App from './components/App';

  ReactDOM.render(
    <AppContainer>
      <App/>
    </AppContainer>,
    document.getElementById('root')
  );

  // Hot Module Replacement API
  if (module.hot) {
    module.hot.accept('./components/App', () => {
      const NextApp = require('./components/App').default;
      ReactDOM.render(
        <AppContainer>
          <NextApp/>
        </AppContainer>
        ,
        document.getElementById('root')
      );
    });
  }
  ```

- 서버 실행 `npm run dev-server`

### create-react-app 사용

- 설치 `npm install -g create-react-app`

- App 생성 `create-react-app hello-world`
  - react 프로젝트를 생성하고 의존성 패키지들을 설치함 (`/node_modules/react-scripts` 경로에)

- webpack-dev-server 실행 `npm start`

- `npm run eject` 명령을 통해 react-scripts를 제거하고 다시 프로젝트에 풀어서 생성 -> 프로젝트 설정가능

- react-hot-loader 적용

## React 시작

### 클래스 생성

`extends React.Component` React 컴포넌트를 상속하는 클래스 생성

### JSX

XML-like syntax를 native JavaScript로 변경
  ```js
  class Codelab extends React.Component {
    render() {
      let text = 'Hi I am codelab';
      let style = {
        backgroundColor: 'aqua'
      }
      return (
        <div style={style}>{text}</div>
      );
    }
  }

  class App extends React.Component {
    render() {
      return (
        <Codelab/>
      );
    }
  }

  ReactDOM.render(<App/>, document.getElementById('root'));
  ```
  ```html
  <div id="root"></div>
  ```

- Container Element (div...)로 감싸야 한다
- JSX 안에서 JavaScript를 표현하기 위해서는 {} 으로 wrapping
- style 선언은 camelcase로 (background-color => backgroundColor)

### props

- 컴포넌트 내부의 immutable data
- JSX 내부에 { this.props.propsName }
- 컴포넌트를 사용 할 때, <> 괄호 안에 propsName = "value"
- this.props.children은 기본적으로 갖고있는 props로 `<Cpnt>`값`</Cpnt>`

  ```js
  class Codelab extends React.Component {
    render() {
      return (
        <div>
          <h1>Hello {this.props.name}</h1>
          <div>{this.props.children}</div>
        </div>
      );
    }
  }

  class App extends React.Component {
    render() {
      return (
        <Codelab name="velopert">{this.props.children}</Codelab>
      );
    }
  }

  ReactDOM.render(<App name="velopert">I am your child</App>, document.getElementById('root'));
  ```
  ```html
  <div id="root"></div>
  ```

### state

- 유동적 데이터
- JSX 내부에 {this.state.stateName}
- 생성자에서 초기값 설정 필수 `this.state={}`, 초기값 설정 한 번만 사용
- 값을 수정할 때에는 `this.setState({...})`

  ```js
  class Counter extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        value: 0
      };
      this.handleClick = this.handleClick.bind(this);
    }

    handleClick() {
      this.setState({
        value:this.state.value + 1
      });
    }

    render() {
      return (
        <div>
          <h2>{this.state.value}</h2>
          <button onClick={this.handleClick}>Press Me</button>
        </div>
      );
    }
  }

  class App extends React.Component {
    render() {
      return (
        <Codelab/>
      );
    }
  }

  ReactDOM.render(<App/>, document.getElementById('root'));
  ```
  ```html
  <div id="root"></div>
  ```

### component

map() 함수 사용
  ```js
  let numbers = [1,2,3,4,5];
  let processed = numbers.map((num) => {
    return num*num;
  });
  ```

컴포넌트 정의
  ```js
  class Contactinfo extends React.Component {
    render() {
      return (
        <div>{this.props.contact.name}{this.props.contact.phone}</div>
      );
    }
  }

  class Contact extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        contactData: [
          {name:'Abet',phone:'010-0000-0001'},
          {name:'Betty',phone:'010-0000-0002'},
          {name:'Charlie',phone:'010-0000-0003'},
          {name:'David',phone:'010-0000-0004'}
        ]
      };
    }

    render() {
      const mapToComponent = (data) => {
        return data.map((contact, i) => {
          return (<ContactInfo contact={contact} key={i}/>);
        });
      };
      return (
        <div>
          {mapToComponent(this.state.contactData)}
        </div>
      );
    }
  }

  class App extends React.Component {
    render() {
      return (
        <Contact/>
      );
    }
  }

  ReactDOM.render(<App></App>, document.getElementById('root'));
  ```
  ```html
  <div id="root"></div>
  ```