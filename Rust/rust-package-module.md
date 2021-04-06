# Rust Package, Module

프로젝트가 커지면 코드를 여러 개의 파일과 모듈로 나누어 관리해야 한다.

범위(scope)란 일종의 중첩된 컨텍스트로, 특정 범위안에 작성된 코드는 그 범위 내에서만 유효한 여러 이름을 사용한다.

러스트는 코드의 구조를 관리하기 위한 기능을 제공한다.
이러한 기능을 합쳐서 모듈 시스템이라고 부르며 다음 개념들이 포함된다.

- 패키지: 크레이트를 빌드, 테스트, 공유할 수 있는 카고의 기능
- 크레이트: 라이브러리나 실행파일을 생성하는 모듈의 트리
- 모듈과 use: 코드의 구조와 범위, 그리고 경로의 접근성을 제어하는 기능
- 경로(path): 구조체, 함수 혹은 모듈 등의 이름을 결정하는 방식

## 패키지와 크레이트

크레이트는 하나의 바이너리 혹은 라이브러리이다.
크레이트 루트(root)는 러스트 컴파일러가 컴파일을 시작해서 크레이트의 루트 모듈을 만들어내는 소스파일이다.

패키지는 이 크레이트를 빌드하는 방법을 서술하는 `Cargo.toml` 파일을 갖는다.

패키지에 포함할 수 있는 아이템을 결정하는 데는 몇 가지 규칙이 있다.

- 패키지는 하나 이상의 라이브러리 크레이트를 포함하거나 아예 포함하지 않을 수 있다.
- 또한 바이너리 크레이트도 원하는 만큼 포함할 수 있지만, 최소한 하나의 (라이브러리 or 바이너리)크레이트는 포함해야 한다.

`cargo new` 명령을 실행하면 `Cargo.toml` 파일을 생성하며 패키지를 만들어 낸다.

이 경우 `src/main.rs` 파일이 패키지와 같은 이름을 갖는 바이너리 크레이트 루트가 된다.

마찬가지로 `src/lib.rs` 파일이 있으면 패키지와 같은 이름의 라이브러리 크레이트를 포함한다고 판단하며 해당파일을 크레이트 루트로 인식한다.

어떤 패키지에 `src/main.rs` 파일과 `src/lib.rs` 파일이 모두 있다면 라이브러리와 바이너리 크레이트를 모두 가진다는 것이며 두 크레이트 이름 모두 패키지 이름과 같다.

패키지의 `src/bin/` 디렉토리에 파일을 분산해서 여러개의 바이너리 크레이트를 추가할 수도 있는데, 이때 각 디렉토리의 파일은 별도의 바이너리 크레이트가 된다.

## 모듈을 이용한 범위와 접근성 제어

모듈(module)은 크레이트의 코드를 그룹화해서 가독성과 재사용성을 향상하는 방법이다.
또한 외부의 코드가 사용할 수 있는지(public), 사용할 수 없는지(private)를 결정하기도 한다.

레스토랑 예제를 살펴보자.

레스토랑 구조는 접객(front of house)과 지원(back of house)로 구분한다.

먼저 `cargo new --lib restaurant` 명령으로 라이브러리를 생성하고, 다음 내용으로 `src/lib.rs` 파일을 작성한다.

```rs
mod front_of_house {
  mod hosting {
    fn add_to_waitlist() {}

    fn seat_at_table() {}
  }

  mod serving {
    fn take_order() {}

    fn serve_order() {}

    fn take_payment() {}
  }
}
```

먼저 `mod` 키워드를 사용해 모듈을 정의한 뒤 모듈의 이름을 지정한다. 모듈 안에는 다른 모듈을 정의할 수 있다.

`src/main.rs`와 `src/lib.rs` 파일을 크레이트 루트라고 부른다고 했다.
그 이유는 이 두파일의 콘덴츠가 `crate`라는 이름의 모듈로 구성되며, 이 모듈은 module tree라고 부르는 크레이트의 모듈구조에서 루트역할을 하기 때문이다.

## 경로를 이용해 모듈 트리의 아이템 참조

경로는 크게 두 가지 형태이다

- absolute path: 크레이트 이름이나 `crate` 리터럴을 이용해 크레이트 루트부터 시작하는 경로
- relative apth: 현재 모듈로부터 시작하며 `self`, `super` 혹은 현재 모듈의 식별자를 이용한다

절대 및 상대 경로는 하나 혹은 그 이상의 식별자로 구성되며 각 식별자는 이중 콜론(`::`)으로 구분한다.

```rs
mod front_of_house {
  mod hosting {
    fn add_to_waitlist() {}
  }
}

pub fn eat_at_restaurant() {
  // 절대경로
  crate::front_of_house::hosting::add_to_waitlist();

  // 상대경로
  front_of_house::hosting::add_to_waitlist();
}
```

러스트에서 접근성이 동작하는 방식은 모든 아이템(함수, 메소드, 구조체, 열거자, 모듈, 상수 ...)은 비공개이고,
부모 모듈의 아이템들은 자식 모듈의 비공개 아이템을 사용할 수 없지만 자식 모듈의 아이템은 부모 모듈의 아이템을 사용할 수 있다.

### pub 키워드로 경로 공개

앞의 코드에서 `hosting` 모듈과 하위의 아이템은 비공개이므로 `pub` 키워드를 사용해야 한다.

```rs
// ...
  pub mod hosting {
    pub fn add_to_waitlist() {}
  }
// ...
```

### super로 시작하는 상대경로

상대경로는 `super` 키워드를 이용해 부모 모듈로부터 시작할 수도 있다

`super` 키워드를 사용하면 나중에 코드를 다른 모듈로 이동해도 수정해야 할 코드를 최소화 할 수 있다.

### 구조체와 열거자 공개

`pub` 키워드를 사용했을 때 구조체는 공개되지만 구조체의 필드는 여전히 비공개 상태이다.

```rs
mod back_of_house {
  pub struct Breakfast {
    pub toast: String,
    seasonal_fruit: String,
  }
}
```

반면 열거자를 공개하면 모든 열거값은 공개된다

```rs
mod back_of_house {
  pub enum Appetizer {
    Soup,
    Salad,
  }
}

pub fn eat_at_restaurant() {
  let order1 = back_of_house::Appetizer::Soup;
  let order2 = back_of_house::Appetizer::Salad;
}
```

## `use` 키워드로 경로를 범위로 가져오기

`use` 키워드와 경로를 추가하면 마치 해당 모듈을 크레이트 루트에 정의한 것처럼 그 범위에서 유효한 이름이된다.

`use` 키워드에 상대경로를 지정하려면 현재 범위의 이름부터 시작하는 대신 `self` 키워드를 이용한 경로를 사용해야 한다.

```rs
mod front_of_house {
  pub mod hosting {
    pub fn add_to_waitlist() {}
  }
}

use self::front_of_house::hosting;

pub fn eat_at_restaurant() {
  hosting::add_to_waitlist();
}
```

### 관용적인 경로 사용하기

일반적으로 경로 불러오기는 모듈수준까지 이루어진다.

구조체 열거자 혹은 다른 아이템은 전체 경로를 사용하는 것이 관례이다.

## `as` 키워드로 새로운 이름 부여

```rs
use std::fmt::Result;
use std::io::Result as IoResult;

fn function1() -> Result {
  //
}

fn function2() -> IoResult<()> {
  //
}
```

### `pub use` 키워드로 이름 다시 내보내기

`use` 키워드를 이용해 범위로 이름을 가져오면 이 이름은 새 범위에서 비공개 이름이 된다.

`pub use` 키워드로 아이템을 현재 범위로 가져올 뿐만 아니라 다른 범위로 내보낼 수 있다.

```rs
mod front_of_house {
  pub mod hosting {
    pub fn add_to_waitlist() {}
  }
}

pub use crate::front_of_house::hosting;
```

### 외부 패키지의 사용

프로젝트에서 `rand` 패키지를 사용하려면 `Cargo.toml` 파일에 다음 코드를 추가한다

```toml
[dependencies]
rand = "0.5.5"
```

의존성이 추가되면 카고는 의존성 패키지를 `https://crates.io/`에서 내려받는다

### 중첩 경로로 `use` 목록 정리

```rs
use std::io;
use std::cmp::Ordering;
```

중첩된 경로를 이용해 위 아이템을 한 줄의 코드로 가져올 수 있다

```rs
use std::{io, cmp::Ordering};
```

### glob 연산자

어떤 경로의 공개 아이템을 모두 현재 범위로 가져오려면 glob 연산자 `*`를 사용해서 경로를 지정한다

```rs
use std::collections::*;
```

## 모듈을 다른 파일로 분리하기

모듈의 크기가 커지면 별도의 파일로 분리하는 것이 좋다.

`src/lib.rs`

```rs
mod front_of_house;

pub use crate::front_of_house::hosting;

pub fn eat_at_restaurant() {
  hosting::add_to_waitlist();
}
```

`src/front_of_house.rs`

모듈의 이름만 선언되어 있으면 본문을 모듈이름과 같은 파일에서 가져온다

```rs
pub mod hosting;
```

`src/front_of_house/hosting.rs`

```rs
pub fn add_to_waitlist() {}
```
