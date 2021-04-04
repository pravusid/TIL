# Rust Struct

## 구조체의 정의와 인스턴스 생성

구조체는 튜플과 유사하다.
하지만 튜플과는 달리 각 데이터에 별개의 이름을 부여해서 더 분명한 표현을 할 수 있다.

구조체를 정의하려면 `struct` 키워드 다음에 이름을 지정해주면 된다.
그 후 중괄호 안에 구조체가 저장할 데이터 타입과 이름을 의미하는 필드(`field`)를 나열하면 된다.

```rs
struct User {
  username: String,
  email: String,
  sign_in_count: u64,
  active: bool,
}
```

구조체를 정의한 후 이를 사용하려면 각 필드에 저장할 값을 명시해서 구조체의 인스턴스를 생성해야 한다.

구조체에서 값을 읽으려면 마침표(`.`)를 사용하면 된다.
만일 인스턴스가 가변 인스턴스라면 마침표를 이용해 특정 필드에 새로운 값을 대입할 수도 있다.

```rs
let user1 = User {
  email: String::from("someone@test.com"),
  username: String::from("someuser"),
  active: true,
  sign_in_count: 1,
}

user1.email = String::from("newone@test.com");
```

러스트는 구조체의 몇몇 필드만 따로 가변 데이터로 표현하는 것을 지원하지 않으므로,
구조체의 새로운 인스턴스를 생성하여 반환하는 함수를 사용할 수 있다.

```rs
fn build_user(email: String, username: String) -> User {
  User {
    email: email,
    username: username,
    active: true,
    sign_in_count: 1,
  }
}
```

이 경우 `email`, `username` 필드가 반복되는 문제가 있는데 이를 해결할 방법이 있다.

### 필드초기화 단축문법

위의 함수에서 매개변수이름과 구조체 필드이름이 같은데, 필드초기화 단축문법을 이용하면 한 번만 이름을 사용해도 된다.

```rs
fn build_user(email: String, username: String) -> User {
  User {
    email,
    username,
    active: true,
    sign_in_count: 1,
  }
}
```

### 기존 인스턴스로부터 새 인스턴스 생성

이미 존재하는 인스턴스에서 몇 개의 필드 값만 수정한 새 구조체 인스턴스가 필요할 때가 있다.
이 경우 구조체 갱신문법(struct update syntax)를 사용하면 편리하다.

```rs
let user2 = User {
  email: String::from("another@test.com"),
  username: String::from("anotheruser"),
  active: user1.active,
  sign_in_count: user1.sign_in_count,
}
```

구조체 갱신 문법을 이용하면 더 적은 코드로 같은 결과를 얻을 수 있다.

`..` 문법은 명시적으로 나열하지 않은 나머지 필드에 대해서는 기존 인스턴스의 필드 값을 사용하라는 의미이다.

```rs
let user2 = User {
  email: String::from("another@test.com"),
  username: String::from("anotheruser"),
  ..user1,
}
```

### 이름 없는 필드를 가진 튜플 구조체로 다른 타입 생성

튜플과 유사하게 생긴 구조체를 튜플 구조체(tuple struct)라고 한다.

튜플 구조체는 구조체에는 이름을 부여하지만, 필드에는 이름을 부여하지 않고 타입만 지정한다.

튜플 구조체를 정의하려면 `struct` 키워드와 구조체의 이름, 그리고 튜플안에서 사용할 타입들을 차례로 나열하면 된다.

```rs
struct Color(i32, i32, i32);

let black = Color(0, 0, 0);
```

### 필드가 없는 유사 유닛 구조체

러스트에서는 필드가 하나도 없는 구조체를 선언할 수도 있다.

이런 구조체를 unit-like structs 라고 한다. 이 구조체는 유닛타입 (`()`)과 유사하게 동작하기 때문이다.

### 구조체 데이터의 소유권

앞에서 정의한 `User` 구조체를 보면 문자열 슬라이스 타입이 아니라 `String` 타입을 사용하고 있다.
그 이유는 구조체가 데이터의 소유권을 갖게 함으로써 유효한 머위 내에 존재하는 동안 데이터도 유효할 수 있도록 하기 위함이다.

구조체에 다른 변수가 소유한 데이터의 참조를 저장할 수도 있지만, 그러려면 lifetimes 기능을 사용해야 한다.

## 메서드 문법

메서드(method)는 함수와 유사하다.

하지만 메서드는 함수와 달리 구조체의 컨텍스트안에 정의하며 첫번째 매개변수는 항상 메서드를 호출할 구조체의 인스턴스를 표현하는 `self`이다.

### 메서드 정의하기

```rs
#[derive(Debug)]
struct Rectangle {
  widht: u32,
  height: u32,
}

impl Rectangle {
  fn area(&self) -> u32 {
    self.width * self.height
  }
}

fn main() {
  let rect = Rectangle { width: 30, height: 50 };
  println!("면적 {}", rect1.area());
}
```

`Rectangle` 타이바의 컨텍스트 안에 메소드를 정의하려면 `impl` 블록을 이용하면 된다.

메서드는 `self`에 대한 소유권을 갖거나, 불변 인스턴스를 대여하거나, 다른 매개변수들 처럼 `self`의 가변 인스턴스를 대여할 수도 있다.

### 메서드 호출 연산자

C, C++ 에서는 메서드를 호출할 때 서로 다른 연산자를 사용한다.

객체의 메서드를 직접호출할 때는 `.` 연산자를 사용하고, 객체의 포인터를 이용해 메서드를 호출할 때는 `->` 연산자를 이용하거나 포인터를 역참조한다.
즉, 객체가 포인터일 때는 `object->something()`은 `(*object).something()`과 유사하다.

러스트에는 `->` 연산자에 해당하는 연산자가 없다. 대신 러스트는 자동 참조 및 역참조기능을 제공한다.

`object.something()`과 같이 메서드를 호추랗면 러스트는 메서드의 시그니처에 따라 자동으로 `object` 변수에 `&`, `&mut` 또는 `*`를 추가한다.

따라서 아래의 코드는 완전히 같다.

```rs
p1.distance(&p2);
(&p1).distance(&p2);
```

이 자동참조 기능은 메서드가 수신자인 `self`의 타입을 명확하게 선언하고 있기 때문에 동작한다.
메서드의 수신자와 이름으로 메서드가 값을 읽는지(`&self`), 값을 쓰는지(`&mut self`), 소비하는지(`self`)를 정확하게 알아낼 수 있다.

### 더 많은 매개변수를 갖는 메서드

Rectangle 구조체의 인스턴스에 또 다른 인스턴스를 전달해서 첫 번째 Rectangle 인스턴스의 면적이 두 번째 인스턴스의 면적을 완전히 포함할 수 있으면 true를 반환하는 코드를 작성해보자.

```rs
fn main() {
  let rect1 = Rectangle { width: 30, height: 50 };
  let rect2 = Rectangle { width: 10, height: 40 };
  let rect3 = Rectangle { width: 60, height: 45 };

  println("rect1 can hold rect2 {}", rect1.can_hold(&rect2));
  println("rect1 can hold rect3 {}", rect1.can_hold(&rect3));
}

impl Rectangle {
  // ...

  fn can_hold(&self, other: &Rectangle) -> bool {
    self.width > other.width && self.height > other.height
  }
}
```

### 연관 함수

`impl` 블록의 다른 기능은 `self` 매개변수를 사용하지 않는 다른 함수도 정의할 수 있다는 점이다.

이런 함수들은 연관 함수(associated functions)라고 한다.
이 함수들은 구조체 인스턴스를 직접 전달받지 않기 때문에 메서드가 아니라 함수다.

`String::from` 함수가 연관함수의 예다.

연관함수는 구조체의 새로운 인스턴스를 반환하는 생성자를 구현할 때 자주 사용한다.

연관함수를 호출하려면 구조체의 이름과 `::` 문법을 사용하면된다.

### 여러개의 impl 블록

각 구조체는 여러 개의 `impl` 블록을 선언할 수 있다.
예를 들어, 각 메소드를 별개의 `impl` 블록에 선언하여도 완전히 똑같이 동작한다.

```rs
impl Rectangle {
  fn area(&self) -> u32 {
    self.width * self.height
  }
}

impl Rectangle {
  fn can_hold(&self, other: &Rectangle) -> bool {
    self.width > other.width && self.height > other.height
  }
}
```
