# Redux

## FLUX

MVC패턴에서 model은 view로 데이터를 보내 렌더링을 한다

사용자와 App의 상호작용은 view를 통해 일어나기 때문에 사용자의 입력에 따라 view가 model을 업데이트 해야할 때도 있다.
업데이트는 비동기적으로 일어날 수 있고 1:N으로 일어날 수도 있다.

문제를 해결하기 위해 데이터를 단방향으로만 보내는 Flux 아키텍처를 만들어낸다.

(Action) -> Dispatcher -> Store -> View -> (Action) -> Dispatcher

### 액션 생성자

App의 상태를 변경하거나 view를 업데이트하려면 액션을 생성해야 한다.

액션 생성자는 시스템에 정의된 액션인 type과 payload를 포함한 액션을 생성한다.

액션 생성자가 메시지를 생성한 뒤 dispatcher로 넘겨준다

### 디스패쳐

dispatcher는 callback이 등록되어 있는 곳이다. 액션이 넘어오면 여러 스토어에 액션을 보낸다.

동기적으로 액션을 처리하고 스토어사이에 의존성이 있다면 순서에따라 처리한다.

Flux의 dispatcher는 다른아키텍쳐들과 다르게 액션타입과 관계없이 등록된 모든 스토어로 보내고, 스토어는 액션을 받은뒤 처리여부를 결정한다.

### 스토어

모든 상태변경은 store에 의해서 결정되어야 한다.

store에는 setter가 없으므로 상태변경을 위한 요청을 store에 직접보낼 수 없다. 즉 액션생성자, 디스패처를 통해 이루어져야한다.

스토어 내부에서 swtich statement를 사용해 처리할 액션과 무시할 액션을 결정하고, 처리할 액션에 따라 상태를 변경한다.

스토어에서 상태를 변경하고 나면 controller view에 change event를 내보낸다.

### 컨트롤러 뷰

controller view는 store와 view사이의 중간관리자 역할을 한다.
상태가 변경되었을 때 store가 사실을 controller view에 알려주면 controller view는 자신 아래에 있는 모든 view에게 새로운 상태를 넘긴다.

## Redux 소개

Flux 아키텍쳐를 구현한 라이브러리

### Redux 원칙

1. Single Source of Truth : 어플리케이션의 state를 위해 단 하나의 store를 사용

1. State is Read-only : App에서 store의 state를 직접변경할수 없고 변경을 위해서는 action이 dispatch되어야 함

1. Changes are made with pure Function : action객체를 처리하는 함수를 reducer라고 부른다.
  reducer는 정보를 받아서 state를 어떻게 업데이트할지 정의한다.
  reducer는 '순수 함수'로 작성되어야 함 : 비동기 처리, 네트워크 및 DB접근, 인수변경, API사용(Math.random()...) => 불가

### Redux 구조

Flux를 구현한 Redux는 일부 구조가 Flux와 다르다

#### Redux 액션 생성자

Redux 액션 생성자는 dispatcher로 액션을 보내지 않고 포맷을 바꾼 뒤 액션을 돌려준다.

#### Redux 스토어

Redux는 하나의 스토어만 존재한다. 스토어는 state tree 전체를 유지하는 역할을 수행한다.
액션에 대한 상태변화는 reducer에게 일임 한다. dispatcher의 역할 수행이라고 볼 수 있다.

#### reducers

스토어는 액션이 어떤 상태변화를 만드는지 리듀서에게 물어본다. 루트 리듀서는 상태를 나누어 각 리듀서에게 할당한다.
하위 리듀서는 넘겨받은 상태변화를 복사본에 적용한다. 리듀서는 계층이 존재하고 필요한 만큼의 깊이를 가질 수 있다.

상태객체는 직접 변경되지 않고 각각의 상태조각 복사본이 변경되고 다시 하나의 새로운 state객체로 합쳐진다.
하위 리듀서로 부터 변경된 상태를 받아 루트 리듀서가 상태를 업데이트 하고 이를 다시 스토어로 보낸다.
스토어는 이 객체를 새로운 App의 state로 만든다.

#### Smart and dumb components

컨트롤러 뷰와 일반 뷰의 개념과 비슷한 컨셉이다.

smart components는 액션처리를 책임진다. 액션 처리가 필요한 경우 props를 통해 dumb comp에 함수를 보낸다.

dumb components는 액션에 직접 의존성을 갖지 않는다. 모든 액션을 props로 넘겨받기 때문이다.
즉 다른 로직을 가진 App에서 재사용될 수 있다.

#### 뷰레이어 바인딩

스토어와 뷰를 연결하기 위해 뷰 레이어 바인딩이 필요하다. react-redux가 그 역할을 수행한다.

뷰 레이어 바인딩은 세 가지의 컨셉을 갖고 있다.

1. provider component : 컴포넌트 트리를 감싸는 컴포넌트.
  `connect()`를 이용해 루트 컴포넌트 아래 컴포넌트가 스토어에 연결되게 해준다.

1. `connect()` : react-redux의 함수이다.
  `connect()`를 이용해서 컴포넌트를 감싸주면 selector를 이용해서 필요한 연결을 만들어, 컴포넌트가 App 상태 업데이트를 받을 수 있다.

1. selector : App 상태안의 어느 부분이 컴포넌트의 props로 필요한 것인지 지정하는 함수이다.

#### 루트 컴포넌트

Redux에서 루트 컴포넌트는 초기화 과정에서 스토어를 생성하고 어떤 리듀서를 사용할지 알려주며 뷰레이어 바인딩과 뷰를 불러온다.

### 작동순서

#### 준비단계

1. 루트 컴포넌트는 `createStore()`를 이용해 스토어를 생성하고 무슨 리듀서를 사용할지 알려준다.
  `combineReducers()`를 이용해서 리듀서를 하나로 묶는다.

1. 루트 컴포넌트는 provider component로 컴포넌트를 감싸고 스토어를 연결한다.

    provider component 내의 smart component는 `connect()`로 스토어에 연결된다.

1. smart component는 `bindActionCreators()`로 액션 콜백을 준비한다. 액션은 포맷이 바뀐 뒤 자동적으로 dumb component로 보내진다.

#### 데이터 흐름

1. 뷰가 액션을 요청한다. 액션 생성자가 포맷 변환 뒤 반환한다.

1. `bindActionCreators()`를 통해 자동으로 액션이 보내진다. 준비단계에서 사용되지 않았다면 뷰가 직접 액션을 보낸다.

1. 스토어가 액션을 받아서 현재 App의 state tree와 액션을 루트 리듀서에게 보낸다.

1. 루트 리듀서는 state tree를 조각으로 나눈 뒤 서브리듀서로 넘겨준다.

1. 서브리듀서는 받은 상태조각의 복사본에 액션을 적용한다. 루트 리듀서에게 변경된 복사본 조각을 보낸다.

1. 루트 리듀서는 받은 상태조각을 모아 만든 state tree를 스토어에 반환한다. 스토어는 옛 상태트리를 새 것으로 바꾼다.

1. 스토어는 뷰레이어 바인딩에게 App 상태가 변경되었다는 것을 알린다.

1. 뷰 레이어 바인딩은 스토어에게 새로운 상태를 요청한다.

1. 뷰 레이어 바인딩은 뷰에게 화면을 업데이트 하도록 요청한다.
