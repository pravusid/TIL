# Rust Generic, Trait and Lifetimes

## 함수로부터 중복 제거하기

제네릭 문법에 앞서 제네릭 타입을 사용하지 않고 함수로부터 중복된 코드를 제거하는 방법을 알아보자.

> 리스트로 부터 가장 큰 숫자 찾기

```rs
fn main() {
  let number_list = vec![34, 50, 25, 100, 65];

  let mut largest = number_list[0];

  for number in number_list {
    if number > largest {
      largest = number;
    }
  }

  println!("가장 큰 숫자: {}", largest);
}
```

두 개의 리스트에서 큰 수를 찾으려면 정수의 리스트를 매개변수로 전달 받아 작업을 실행하는 함수를 정의하면 된다

```rs
fn largest(list: &[i32]) -> i32 {
  let mut largest = list[0];

  for &item in list.iter() {
    if item > largest {
      largest = item;
    }
  }

  largest
}

fn main() {
  let number_list = vec![34, 50, 25, 100, 65];

  let result = largest(&number_list);
  println!("가장 큰 숫자: {}", result);

  let number_list = vec![341, 502, 215, 1500, 625];

  let result = largest(&number_list);
  println!("가장 큰 숫자: {}", result);
}
```

## 제네릭 데이터 타입

제네릭은 여러 구체화된 타입을 사용할 수 있는 함수 시그니처나 구조체 같은 아이템을 정의할 때 사용한다.

### 함수 정의에서 사용하기

일반적으로 제네릭을 사용하는 함수를 정의할 때는 특정한 타입의 매개변수와 반환타입을 사용하는 함수의 시그니처에 사용한다.

> 시그니처의 이름과 타입만 다르고 동작은 같은 두 함수

```rs
fn largest_i32(list: &[i32]) -> i32 {
  let mut largest = list[0];

  for &item in list.iter() {
    if item > largest {
      largest = item;
    }
  }

  largest
}

fn largest_char(list: &[char]) -> char {
  let mut largest = list[0];

  for &item in list.iter() {
    if item > largest {
      largest = item;
    }
  }

  largest
}
```

제네릭 `largest` 함수를 정의하려면 함수의 이름과 매개변수 목록 사이의 꺾쇠괄호(<>)에 타입 이름을 명시하면 된다

`fn largest<T>(list: &[T]) -> T`

이 함수는 어떤 타입 `T`로 일반화 한 함수이다

```rs
fn largest<T>(list: &[T]) -> T {
  let mut largest = list[0];

  for &item in list.iter() {
    if item > largest {
      largest = item;
    }
  }

  largest
}
```

그러나 위 코드를 컴파일 하면 에러가 발생한다

```txt
binary operation '>' cannot be applied to type 'T'
note: 'T' might need a bound for 'std::cmp::PartialOrd'
```

함수 내부에서 비교연산을 하므로 `std::cmp::PartialOrd` 트레이트를 구현할 것을 요구한다.

### 구조체 정의에서 사용하기

구조체의 필드에도 `<>` 구문을 이용해 제네릭 타입 매개변수를 사용할 수 있다

```rs
struct Point<T> {
  x: T,
  y: T,
}
```

구조체의 필드를 각기 다른 타입으로 선언하고 싶다면 다중 제네릭 타입 매개변수를 사용하면 된다

```rs
struct Point<T, U> {
  x: T,
  y: U,
}
```

### 열거자 정의에서 사용하기

구조체와 마찬가지로 열거값에 제네릭 데이터 타입을 사용하는 열거자도 정의할 수 있다.

```rs
enum Result<T, E> {
  Ok(T),
  Err(E),
}
```

### 메서드 정의에서 사용하기

구조체나 열거자의 메서드에도 제네릭 타입을 활용할 수 있다.

```rs
struct Point<T> {
  x: T,
  y: T,
}

impl<T> Point<T> {
  fn x(&self) -> &T {
    &self.x
  }
}

fn main() {
  let p = Point { x: 5, y: 10 };
  println!("p.x = {}", p.x());
}
```

`impl` 키워드 다음에 타입 `T`를 지정하면 제네릭 타입임을 인식한다.

모든 타입을 사용할 수 있는 제네릭 인스턴스 대신 `Point<f32>` 처럼 특정 타입의 인스턴스에만 적용할 메서드를 구현할 수도 있다

```rs
impl Point<f32> {
  fn distance_from_origin(&self) -> f32 {
    (self.x.powi(2) + self.y.powi(2)).sqrt()
  }
}
```

`f32`가 아닌 다른 `Point<T>` 인스턴스는 이 메서드를 사용할 수 없다.

구조체의 정의에 사용하는 제네릭 타입 매개변수는 구조체의 메서드 시그니처에 사용한 것과 항상 같지는 않다.

```rs
struct Point<T, U> {
  x: T,
  y: U,
}

impl<T, U> Point<T, U> {
  fn mixup<V, W>(self, other: Point<V, W>) -> Point<T, W> {
    Point {
      x: self.x,
      y: other.y,
    }
  }
}

fn main() {
  let p1 = Point { x: 5, y: 10.4 };
  let p2 = Point { x: "hello", y: 'c' };

  let p3 = p1.mixup(p2);
  println!("p3.x = {}, p3.y = {}", p3.x, p3.y);
}
```

### 제네릭의 성능

러스트에서 제네릭 타입을 사용한다고 해서 구체화된 타입을 사용할 때보다 성능이 떨어지지 않는다.

러스트는 컴파일 시점에 제네릭을 사용하는 코드를 단일화(Monomorphzation)한다.
단일화는 컴파일 시점에 제네릭 코드를 실제로 사용하는 구체화된 타입으로 변환하는 과정이다.

## 트레이트: 공유 가능한 행위를 정의하는 방법
