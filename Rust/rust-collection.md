# Rust Collection

러스트에서 빈번하게 활용하는 세 가지 컬렉션을 알아보자

- vector: 연속된 일련의 값을 저장한다
- string: 문자(character)의 컬렉션이다
- hashmap: 키-값을 저장하는 map의 구현체이다

## 벡터(vector)

벡터를 이용하면 하나 이상의 값을 데이터 구조에 담을 수 있으며 모든 값은 메모리상에 연속으로 저장된다

벡터는 같은 타입의 값만 저장할 수 있다.

### 새로운 벡터 생성

```rs
let v: Vec<i32> = Vec::new();
```

벡터에 초기값이 있으면 타입을 유추할 수 있으므로 일반적으로는 `vec!` 매크로가 제공된다

```rs
let v = vec![1, 2, 3];
```

### 벡터 수정

벡터에 값을 추가하려면 `push` 메서드를 사용한다

```rs
let mut v = Vec::new();
v.push(1)
v.push(2)
```

### 벡터 해제

벡터 역시 범위를 벗어날 때 `drop` 메서드가 호출된다

```rs
{
  let v = vec![1, 2, 3];

  // v 사용가능
} // v 메모리가 해제됨
```

### 벡터로부터 값 읽기

벡터에 저장된 값을 참조하는 방법은 크게 두 가지이다 (인덱스 문법, get 메서드)

```rs
let v = vec![1, 2, 3, 4, 5];

let third: &i32 = &v[2];
println!("세 번째: {}", third);

match v.get(2) {
  Some(third) => println!("세 번째: {}", third),
  None => println!("none"),
}
```

벡터 참조와 인덱스로 접근할 수 있다. 이 경우 참조가 존재하지 않으면 패닉이 발생한다.

`get` 메서드를 이용하면 `Option<&T>` 타입의 값을 얻을 수 있다.

유효한 참조값을 얻게되면 참조가 계속해서 유효할 수 있도록 대여값 검사가 실행된다.

```rs
let mut v = vec![1, 2, 3, 4];

let first = &v[0];
v.push(5);
```

그러나 코드는 컴파일 오류가 발생한다

벡터는 연속한 메모리영역에 데이터를 보관하므로 벡터에 값을 추가하는 도중 벡터에 여유공간이 없다면 메모리 재할당이 필요하기 때문이다.

### 벡터에 저장된 값 순회

`for` 루프를 이용해서 벡터에 저장된 값에 대한 참조를 얻을 수 있다

```rs
let v = vec![1, 2, 3, 4, 5];

for i in &v {
  println!("{}", i);
}
```

가변 참조를 얻어서 값을 변경할 수도 있다

```rs
let mut v = vec![1, 2, 3, 4, 5];

for i in &mut v {
  *i += 50;
}
```

가변 참조가 가리키는 값을 변경하려면 역참조 연산자(`*`)를 이용해서 변수 `i`에 저장된 값을 가져와야 한다

### 열거자를 이용해 여러 타입 저장하기

열거자의 열거값을 벡터에 저장하면 여러 타입을 저장할 수 있다

```rs
enum SpreadsheetCell {
  Int(i32),
  Float(f64),
  Text(String),
}

let row = vec![
  SpreadsheetCell::Int(3),
  SpreadsheetCell::Text(String::from("blue")),
  SpreadsheetCell::Float(10.10),
]
```

러스트는 컴파일 시점에 벡터에 어떤 타입의 값이 저장되는지 알고 있으므로 어느 정도의 힙 메모리가 필요한지도 알고있다.

## String 타입에 UTF-8 형식의 텍스트 저장

문자열은 생각보다 복잡한 데이터 구조이며 UTF-8 형식으로 저장된다

### 문자열이란 무엇일까

러스트는 언어명세에서 오직 한 가지 문자열 슬라이스인 `str` 타입만을 지원한다.

문자열 슬라이스는 어딘가에 UTF-8 형식으로 인코딩되어 저장된 문자열에 대한 참조이다.
반면 `String` 타입은 언어 명세에 정의된 것이 아니라 러스트 표준 라이브러리가 제공하는 타입이다.

러스트의 문자열이란 둘 중 하나를 의미하는 것이 아니라 `String` 타입과 문자열 슬라이스 `&str` 타입을 동시에 의미한다.

러스트의 표준 라이브러리는 `OsString`, `OsStr`, `CString`, `Cstr`과 같은 다른 종류의 문자열 타입도 제공한다.

### 새 문자열 생성하기

`String` 타입은 `Vec<T>` 타입이 지원하는 대부분의 작업을 지원한다

`new` 함수를 사용해서 새 문자열을 생성할 수 있다

```rs
let mut s = String::new();
```

문자열 리터럴의 `to_string` 메서드를 이용해서 문자열을 생성할 수 있다

```rs
let s = "문자열 초기값".to_string();
```

문자열 리터럴을 이용해 `String` 타입을 생성하려면 `String::from` 함수를 이용해도 되며, `to_string`과 똑같이 동작한다

```rs
let s = String::from("문자열 초기값");
```

### 문자열 수정하기

문자열의 크기는 늘어날 수 있으며 벡터처럼 더 많은 데이터를 넣으면 내용도 변경할 수 있다

`+` 연산자나 `format!` 매크로를 이용해 문자열을 연결할 수도 있다

#### `push_str`, `push` 메서드를 이용해 문자열 연결

`push_str` 메서드는 문자열 슬라이스를 `String`에 연결한다

```rs
let mut s = String::from("foo");
s.push_str("bar");
```

`push_str` 메서드는 문자열 슬라이스의 소유권을 가질 필요가 없으므로 이후 매개변수를 재호출 할 수 있다.

```rs
let mut s1 = String::from("foo");
let s2 = "bar";
s1.push_str("bar");
println!("s2: {}", s2);
```

`push` 메서드는 하나의 문자(character)를 매개변수로 받아 `String`에 추가한다.

```rs
let s = String::from("lo");
s.push('l');
```

#### `+` 연산자나 `format!` 매크로를 이용한 연결

```rs
let s1 = String::from("hello, ");
let s2 = String::from("world!");
let s3 = s1 + &s2; // s1은 메모리가 해제되어 더 이상 사용할 수 없음
// s2는 유효함
```

`+` 연산자를 사용했을 때 호출되는 메서드의 시그니처는 다음의 형태이다

내용을 살펴보면, s1의 소유권을 확보한 뒤 s2 참조를 매개변수로 받아 강제 역참조한 뒤 값을 복사해서 덧붙이고 결과값에 대한 소유권을 반환한다

```rs
fn add(self, s: &str) -> String
```

복잡한 문자열 결합에는 `format!` 매크로가 더 적합하다

```rs
let s1 = String::from("tic");
let s2 = String::from("tac");
let s3 = String::from("toe");

let s = format!("{}-{}-{}", s1, s2, s3);
```

### 문자열의 인덱스

러스트에서는 인덱스를 이용해 `String` 값의 일부에 접근하려고 하면 오류가 발생한다.

#### 문자열의 내부

`String`은 `Vec<u8>` 타입을 한번 감싼 타입니다.

```rs
let len = String::from("Hola").len();
```

이 경우 변수 `len`의 값은 4가 된다. 즉 벡터에 저장된 문자열 'Hola'의 길이가 4byte라는 뜻이다.

```rs
let len = String::from("안녕하세요").len();
```

러스트에서 이 문자열의 길이는 15이다. UTF-8 형식으로 인코딩하면 15byte를 사용한다.

그래서 문자열의 바이트에 인덱스로 접근하면 올바른 유니코드 스칼라값을 가져오지 못할 수 있다.

```rs
let hello = String::from("안녕하세요");
let answer = &hello[0];
```

이 문자열을 UTF-8 형식으로 인코딩하면 '안'의 첫 번째 바이트는 236이고 두 번째 바이트는 149이다.
그래서 answer 변수값은 236이다. 그러나 236은 우리가 얻으려는 문자가 아니다.

이런 문제를 피하기 위해서 러스트는 미리 실수를 차단한다.

#### 바이트와 스칼라값, 그래핌 클러스트

러스트 관점에서 볼 때 문자열은 크게 바이트(bytes), 스칼라값(scalar values), 그리고 그래핌 클러스터(grapheme clusters == letter) 세 가지로 구분한다.

예를 들어, 한글 '안녕하세요'는 다음과 같이 `u8` 값들의 벡터에 저장된다

`[236, 149, 136, 235, 133, 149, 237, 149, 152, 236, 132, 184, 236, 154, 148]`

이 값을 러스트의 `char` 타입인 유니코드 스칼라 값으로 표현하면 다음과 같다

`['안', '녕', '하', '세', '요']`

이 벡터에는 다섯 개의 `char` 값이 저장되어 있다.
같은 데이터를 그래핌 클러스터로 표현하면 다음과 같다.

`["안", "녕", "하", "세", "요"]`

러스트가 `String` 타입에서 인덱스 사용을 지원하지 않는 마지막 이유는 인덱스 처리에는 항상 일정한 시간(`O(1)`)이 소요되어야 하지만
유효한 문자를 파악하기 위해서 처음부터 스캔해야 하므로 `String` 타입에 대해서는 일정한 성능을 보장할 수 없어서이다.

### 문자열 슬라이스 하기

문자열 인덱싱 작업은 상황에 따라 결과 타입이 하나의 바이트, 문자, 그래핌 클러스터 혹은 하나의 문자열 슬라이스가 될 수 있으므로 좋은 선택이 아니다.

인덱스를 이용해 문자열 슬라이스를 생성하려는 의미를 명확히 하려면 `[]` 기호에 하나의 숫자를 인덱스로 전달하는 대신 특정 바이트 범위를 지정해야 한다.

```rs
let hello = "안녕하세요";
let s = &hello[0..3];
```

만일 `&hello[0..1]` 범위를 호출하면 런타임에 패닉이 발생한다.

> thread 'main' panicked at 'byte index 1 is not a char boundary; ...

### 문자열을 순회하는 메서드

개별 유니코드 스칼라값을 조작해야 한다면 가장 좋은 방법은 `chars` 메서드를 사용하는 것이다.

'안녕하세요' 문자열애 대해 `chars` 메서드를 호출하면 문자열을 다섯 개의 `char` 타입 값으로 분리하고 개별문자를 순회할 수 있다.

```rs
for c in "안녕하세요".chars() {
  println!("{}", c);
}
```

다음 결과가 출력된다

```txt
안
녕
하
세
요
```

`bytes` 메서드는 문자열의 각 바이트를 반환할 때 사용한다.

## 키와 값을 저장하는 해시맵

`HashMap<K, V>` 타입은 `K` 타입의 키에 `V` 타입의 값을 매핑하여 저장한다.

이 과정에서 해시함수를 통한 키와 값을 저장하는 규칙을 사용한다.

여러 프로그래밍 언어가 유사한 형태의 데이터 구조를 제공하는데 해시, 맵, 객체, 해시테이블, 딕셔너리, 연관배열등의 다른 이름을 사용한다.

### 새로운 해시 맵 생성하기

`new` 함수로 빈 해시 맵을 생성할 수 있고, `insert` 메서드를 통해 새로운 키와 값을 추가할 수 있다.

```rs
use std::collections::HashMap;

let mut scores = HashMap::new();

scores.insert(String::from("blue"), 10);
scores.insert(String::from("yellow"), 50);
```

표준 라이브러리는 해시 맵을 생성하는 내장 매크로를 제공하지 않는다.

벡터와 마찬가지로 해시 맵은 데이터를 힙 메모리에 저장한다.
또한 벡터처럼 해시 맵 역시 모든 키와 모든 값의 타입이 같아야 한다.

해시 맵을 생성하는 다른 방법은 키와 값을 가지고 있는 튜플의 벡터에 대해 `collect` 메서드를 호출하는 방식이다.

```rs
use std::collections::HashMap;

let teams = vec![String::from("blue"), String::from("yellow")];
let initial_scores = vec![10, 50];

let scores: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();
```

`collect` 메서드는 여러가지 데이터 구조를 생성할 수 있으므로 그중에 어떤 타입을 생성할 것인지를 명시하기 위해 `HashMap<_, _>` 타입 애노테이션이 필요하다.

### 해시 맵과 소유권

`i32`처럼 `Copy` 트레이트를 구현하는 타입은 캆들이 해시 맵으로 복사된다.
반면, `String`처럼 값을 소유하는 타입은 값이 해시 맵으로 이동하며, 해시 맵이 그 값들의 소유권을 갖게된다.

```rs
use std::collections::HashMap

let field_name = String::from("Favorite color");
let field_value = String::from("블루");

let mut map = HashMap::new();
map.insert(field_name, field_value);
// field_name, field_value 변수는 여기서부터 유효하지 않음
```

만일 해시 맵에 값으 참조를 추가하면 그 값은 해시맵으로 이동하지 않는다.
다만, 참조가 가리키는 값은 해시 맵이 유효한 범위애 있는 동안 함께 유효해야 한다.

### 해시 맵의 값에 접근하기

```rs
use std::collections::HashMap;

let mut scores = HashMap::new();

scores.insert(String::from("blue"), 10);
scores.insert(String::from("yellow"), 50);

let team_name = String::from("blue");
let score = scores.get(&team_name);
```

예제에서 `score` 변수는 `blue`에 연결된 값을 갖게 되며 타입은 `Some(&10)`이다.

`for` 루프를 이용하면 벡터와 마찬가지로 해시 맵에서 기와 값의 쌍을 순회할 수 있다.

```rs
// ...

for (key, value) in &scores {
  println!("{}: {}", key, value);
}
```

### 해시 맵 수정하기

해시 맵의 각 키에는 하나의 값만 할당할 수 있다.

해시 맵의 데이터를 수정할 때는 키에 이미 값이 할당된 경우를 어떻게 처리할지 결정해야 한다.

#### 값 덮어쓰기

```rs
use std::collections::HashMap;

let mut scores = HashMap::new();

scores.insert(String::from("blue"), 10);
scores.insert(String::from("blue"), 20);
```

#### 키에 값이 할당되어 있지 않을 때만 추가

해시 맵은 값의 할당 여부를 확인할 키를 매개변수로 사용하는 `entry`라는 API를 제공한다.

`entry` 메서드의 반환값은 값이 존재하는지 알려주는 `Entry` 열거자이다.

```rs
use std::collections::HashMap;

let mut scores = HashMap::new();

scores.insert(String::from("blue"), 10);
scores.entry(String::from("blue")).or_insert(50);
// blue: 10
```

#### 기존 값에 따라 값 수정

키의 값을 확인한 후 현재 값에 따라 새 값으로 수정할 때가 있다

```rs
use std::collections::HashMap;

let text = "hello world wonderful world";

let mut map = HashMap::new();

for word in text.split_whitespace() {
  let count = map.entry(word).or_insert(0);
  *count += 1;
}

// { "world": 2, "hello" 1, "wonderful": 1 }
```

`or_insert` 메서드는 키에 할당된 값에 대한 가변 참조(`&mut V`)를 반환한다.

### 해시 함수

기본 해시 함수 대신 `BuildHasher` 트레이트를 구현하는 다른 해시 제공자(hasher)를 사용할 수 있다.
