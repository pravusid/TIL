# Rust Enums & Pattern Matching

열거자는 다양한 언어에서 지원하지만 조금씩다르다.
러스트의 결거자는 함수형 언어의 대수식(algebraic) 데이터 타입에 더 가깝다.

## 열거자 정의하기

```rs
enum IpAddrKind {
  V4,
  V6,
}
```

이렇게 정의한 `IpAddrKind` 열거자는 코드에서 사용할 수 있는 하나의 타입으로 간주된다.

### 열거자의 값

`IpAddrKind` 열거자의 각 값을 표현하는 인스턴스는 다음과 같이 생성한다.

```rs
let four = IpAddrKind::V4;
let six = IpAddrKind::V6;
```

데이터를 열거자의 열거값에 직접 넣을 수 있다.

```rs
enum IpAddr {
  V4(String),
  V6(String),
}

let home = IpAddr::V4(String::from("127.0.0.1"));
let loopback = IpAddr:V6(String::from("::1"));
```

구조체 대신 열거자를 사용함으로써 얻을 수 있는 또 다른 이점은 열거자에 나열된 각각의 값은 서로 다른 타입과 다른 수의 연관데이터를 보유할 수 있다는 점이다.

```rs
enum IpAddr {
  V4(u8, u8, u8, u8),
  V6(String),
}

let home = IpAddr(127, 0, 0, 1);
```

IP 주소를 저장하교 표현하는 것은 일반적인 작업이므로 이미 표준 라이브러리에 정의되어 있다.

```rs
struct Ipv4Addr {
  // ...
}

struct Ipv6Addr {
  // ...
}

enum IpAddr {
  V4(Ipv4Addr),
  V6(Ipv6Addr),
}
```

더 다양한 타입을 이용해서 열거자의 값을 정의할 수 있다

```rs
Quit,
Move { x: i32, y: i32 },
Write(String),
ChangeColor(i32, i32, i32),
```

이를 통해 구조체와 달리 여러종류의 메시지를 매개변수로 전달받는 함수의 매개변수 타입을 쉽게 정의할 수 있다.

`impl` 블록을 이용해 열거자에도 메서드를 정의할 수 있다.

```rs
impl Message {
  fn call(&self) {
    // body
  }
}

let m = Message::Write(String::from("hello"));
m.call();
```

### Option 열거자를 Null값 대신 사용할 때의 장점

표준 라이브러리가 제공하는 `Option` 열거자는 다양한 곳에서 사용할 수 있다.

러스트는 다른 언어가 가지고 있는 `null`이라는 값의 개념이 없다.

널 값의 문제점은 널 값을 널이 아닌 값처럼 사용하려고 하면 에러가 발생한다는 점이다.

사실 문제는 널의 개념이 아니라 그 구현에 있다.
그래서 러스트는 널 값이라는 개념이 없지만, 어떤 값의 존재여부를 표현하는 열거자를 정의하고 있다.
이 열거자가 바로 `Option<T>`이며 표준 라이브러리에 다음과 같이 정의되어 있다.

```rs
enum Option<T> {
  Some(T),
  None,
}
```

`Option<T>` 열거자는 매우 유용하며 심지어 프렐류드에 포함되어 있다.
따라서 이 열거자는 물론 그 안에 열거된 값도 명시적으로 범위로 가져올 필요 없다.
즉, `Option::` 접두어 없이도 `Some`이나 `None`값을 직접 사용할 수 있다.

다음 예제는 숫자 타입과 문자열 타입을 저장하는 `Option` 열거자를 사용하는 방법이다.

```rs
let some_number = Some(5);
let some_string = Some("a string");

let absent_number: Option<i32> = None;
```

`Some` 대신 `None` 값을 사용하려면 컴파일러에게 열거자의 타입이 무엇인지를 알려줘야 한다.

`Option<T>`와 `T`는 다른 타입이기 때문에 컴파일러는 유효한 값이 명확히 존재할 때는 `Option<T>` 값을 사용하는 것을 허락하지 않는다.

따라서 `T` 타입에 대한 작업을 실행하기 전 `Option<T>` 타입을 `T` 타입으로 변환해야 한다.

## match 흐름 제어 연산자

통상 `Option<T>` 값을 사용하려면 열거자에 나열된 개별 값들을 처리할 코드를 작성해야 한다.

match 표현식은 이런 코드를 쉽게 작성할 수 있는 흐름 제어 연산자이다.

패턴은 리터럴, 변수 이름, 와일드카드를 비롯해 다양한 값으로 구성할 수 있다.

다음은 동전을 받아 동전 종류를 파악한 후 동전의 가치를 센트로 반환하는 함수를 구현한 코드이다.

```rs
enum Coin {
  Penny,
  Nickle,
  Dime,
  Quarter,
}

fn value_in_cents(coin: Coin) -> u32 {
  match coin {
    Coin::Penny => {
      println!("Penny")
      1
    },
    Coin::Nickle => 5,
    Coin::Dime => 10,
    Coin:: Quarter => 25,
  }
}
```

`if`문의 표현식은 반드시 분리언 값을 리턴해야하지만, `match` 연산자의 표현식은 어떤 타입의 값도 사용할 수 있다.

패턴매칭에서 여러 줄의 코드를 실행하고자 한다면 중괄호를 사용해야 한다.

### 값을 바인딩하는 패턴

`match` 표현식이 갖는 또 다른 장점은 패턴에 일치하는 값 일부를 바인딩할 수 있다는 점이다.

```rs
#[derive(Debug)]
enum UsState {
  Alabama, Alaska, // ...
}

enum Coin {
  Penny,
  Nickle,
  Dime,
  Quarter(UsState),
}

fn value_in_cents(coin: Coin) -> u32 {
  match coin {
    Coin::Penny => 1,
    Coin::Nickle => 5,
    Coin::Dime => 10,
    Coin:: Quarter(state) => {
      println!("State: {}", state);
      25
    },
  }
}
```

### `Option<T>`를 이용한 매칭

```rs
fn plus_one(x: Option<i32>) -> Option<i32> {
  match x {
    None => None,
    Some(i) => Some(i + 1),
  }
}

let five = Some(5);
let six = plus_one(five);
let none = plus_one(None);
```

### match는 반드시 모든 경우를 처리해야 한다

다음 예제의 `plus_one` 함수는 버그를 가지고 있으며 컴파일조차 되지 않는다.

```rs
fn plus_one(x: Option<i32>) -> Option<i32> {
  match x {
    Some(i) => Some(i + 1),
  }
}
```

이 예제에서는 `None` 값에 해당하는 경우를 처리하지 않기 때문에 코드가 컴파일 되지 않는다.

### 자리지정자 `_`

러스트는 모든 경우를 다 처리하고 싶지 않을 때 사용하는 패턴도 제공한다.
이처럼 별도의 처리가 필요하지 않을 때는 `_` 자리지정자로 대체하면 된다.

```rs
let some_u8_value = 0u8;

match some_u8_value {
  1 => println!("one"),
  3 => println!("three"),
  5 => println!("five"),
  7 => println!("seven"),
  _ => (),
}
```

`_` 패턴은 모든 값에 일치함을 의미한다.
따라서 `match` 표현식의 마지막에 추가하면 앞에 나열했던 패턴에 일치하지 않는 나머지 모든 패턴에 일치하게 된다.

`()`는 단순한 유닛값이며 이 경우 아무일도 일어나지 않는다.

하지만 `match` 표현식은 한 가지 경우만 처리해야 할 때 사용하기는 장황하다.
이런 경우를 대비해 러스트는 `if let` 구문을 제공한다.

## if let을 이용한 간결한 흐름 제어

`if let` 문법은 여러 경우 중 한가지만 처리하고, 나머지는 고려하고 싶지 않을 때 간편하게 처리하는 문법이다.

```rs
if let Some(3) = some_u8_value {
  println!("three");
}
```

`if let` 문법은 주어진 값에 대해 하나의 패턴만 검사하고 나머지 값은 무시하기 위한 `match` 표현식의 문법설탕이다.

`if let` 구문은 `else` 구문을 포함할 수 있는데 이는 `match` 표현식의 `_` 패턴에 해당한다.

```rs
let mut count = 0;

if let Coin::Quarter(state) => coin {
  println!("State: {}", state);
} else {
  count += 1;
}
```
