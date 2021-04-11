# Rust Error Handling

러스트는 에러를 크게 회복 가능한(recoverable) 에러와 회복 불가능한(unrecoverable) 에러로 구분한다.

회복 불가능한 에러는 배열의 범위를 벗어나는 메모리에 대한 접근처럼 항상 버그 가능성을 내포하고 있다.

대부분 언어는 두 가지 에러를 구분하지 않고 예외(exception)같은 메커니즘을 이용해 처리한다.

러스트에 예외라는 개념은 없다. 대신 회복 가능한 에러를 표현하는 `Result<T, E>` 타입과
회복 불가능한 에러가 발생한 프로그램의 실행을 `panic!` 매크로를 지원한다.

## panic! 매크로를 이용한 회복 불가능한 에러 처리

간혹 코드에 문제가 생겼는데 처리할 방법이 전혀 없는 경우가 있다.

이때 `panic!` 매크로를 사용하면 프로그램은 실패 메시지를 출력하고 스택을 모두 정리한 다음 종료한다.

패닉 매크로를 호출해보자

```rs
fn main() {
  panic!("crash and burn");
}
```

### 패닉이 발생했을 때 스택을 풀어주거나(unwind) 취소하기

기본적으로 패닉이 발생하면 러스트는 스택을 역순으로 순회하면서 각 함수에 전달되었던 데이터를 정리하기 시작한다.

또 다른 방법은 스택을 즉시 취소해서 스택을 정리하지 않고 애플리케이션을 종료하는 방법이 있다.
이 경우 프로그램이 사용하던 메모리는 운영체제가 정리해야 한다.

만일 프로그램의 바이너리 파일의 크기를 최대한 작게해야 한다면 `Cargo.toml` 파일의 `[profile]` 섹션에 `panic = 'abort'`를 추가해서
스택을 풀어주지 않고 취소하게 할 수 있다.

```toml
[profile.release]
panic = 'abort'
```

### panic! 역추적 사용하기

코드에서 매크로를 직접 호출하는 것이 아니라 라이브러리 안에서 `panic!` 매크로가 호출되는 예제를 보자.

```rs
fn main() {
  let v = vec![1, 2, 3];
  v[99];
}
```

이 경우 러스트는 패닉을 일으킨다.

C 언어의 경우 원치 않던 값이라도 해당하는 위치의 메모리에 저장된 값을 반환한다.
이를 buffer overread 라고 하며 공격자가 읽어서는 안되는 값을 읽게 만드는 보안 취약점이 되기 쉽다.

`RUST_BACKTRACE` 환경변수를 이용하면 에러를 자세히 역추적할 수 있다.

```sh
RUST_BACKTRACE=1 cargo run
```

그러면 자세한 정보들이 출력되는데 `cargo build`, `cargo run` 명령을 `--release` 플래그 없이 실행해서 debug symbols 활성화 상태여야 한다.

## Result 타입으로 에러 처리하기

대부분의 에러는 프로그램을 완전히 종료해야 할 정도로 치명적이지는 않다.

`Result` 열거자는 `Ok`와 `Err` 열거값이 정의되어 있다.

```rs
enum Result<T, E> {
  Ok(T),
  Err(E),
}
```

`Result` 타입을 반환하는 함수를 호출해보자

```rs
use std::fs::File;

fn main() {
  let f = File::open("hello.txt");

  let f = match f {
    Ok(file) => file,
    Err(error) => {
      panic!("파일 열기 실패: {:?}", error);
    }
  }
}
```

`Option` 열거자와 마찬가지로 `Result` 열거자와 열거값도 프렐류드에 의해 자동으로 임포트된 상태이므로 `Result::`를 명시할 필요가 없다.

### match 표현식으로 여러 종류의 에러 처리하기

무조건 `panic`을 호출하지 않고 실패 원인에 따라 다르게 동작하도록 수정해보자.

```rs
use std::fs::File;

fn main() {
  let f = File::open("hello.txt");

  let f = match f {
    Ok(file) => file,
    Err(ref error) => match error.kind() {
      ErrorKind::NotFound => match File::create("hello.txt") {
        Ok(fc) => fc,
        Err(e) => panic!("파일을 생성하지 못했습니다: {:?}", e)
      },
      other_err => panic!("파일을 열지 못했습니다: {:?}", other_err),
    },
  };
}
```

### 에러 발생 시 패닉을 발생하는 더 빠른 방법: unwrap, expect

`match` 표현식을 사용하는 방법은 의도대로 잘 동작하지만, 코드가 길어지며 의도를 정확하게 표현하지 못한다.

그래서 `Result<T, E>` 타입은 여러가지 헬퍼 메서드를 제공한다.

`unwrap` 메서드는 `match` 표현식과 정확히 같은 동작을 하는 단축 메서드이다.

`Result` 타입의 값이 `Ok` 열거값이라면 `unwrap` 메서드는 `Ok` 열거값에 저장된 값을 반환한다.
만일 `Err` 열거값이라면 `panic!` 매크로를 호출한다.

`expect` 메서드는 `unwrap` 메서드와 비슷하지만 `panic!` 매크로에 에러 메시지를 전달해준다.

```rs
use std::fs::File;

fn main() {
  let f = File::open("hello.txt").expect("파일을 열 수 없습니다!");
}
```

### 에러 전파하기

실패할 가능성이 있는 함수를 호출하는 함수를 작성할 때는 에러를 함수안에서 처리하지 않고
호출하는 코드에 에러를 반환해서 호출자가 에러를 처리하게 할 수 있다.

```rs
use std::io;
use std::io::Read;
use std::fs::File;

fn read_username_from_file() -> Result<String, io::Error> {
  let f = File::open("hello.txt");

  let mut f = match f {
    Ok(file) => File,
    Err(e) => return Err(e),
  };

  let mut s = String::new();

  match f.read_to_string(&mut s) {
    Ok(_) => Ok(s),
    Err(e) => Err(e),
  }
}
```

에러가 발생했을 때 `panic!` 매크로를 호출하는 것이 아니라,
함수의 실행을 조기 중단하고 `File::open` 메서드가 리턴한 에러값을 이 함수의 에러값으로 호출자에게 반환한다.

러스트에서는 이처럼 에러를 전파하는 것이 일반적이므로 이 작업을 쉽게 처리할 수 있도록 `?` 연산자를 제공한다.

#### `?` 연산자를 이용해 에러를 더 쉽게 전파하기

```rs
use std::io;
use std::io::Read;
use std::io::File;

fn read_username_from_file() -> Result<String, io::Error> {
  let mut f = File::open("hello.txt")?;
  let mut s = String::new();

  f.read_to_string(&mut s)?;
  Ok(s);
}
```

`?` 연산자는 에러처리를 위해 `match` 연산자를 사용했을 때와 같은 방식으로 동작한다.

`Result` 값이 `Ok`면 `Ok` 열거값에 저장된 값이 반환되며, `Err`이면 `Err` 값이 반환되어 에러가 전파된다.

`match` 표현식과 `?` 연산자 동작에는 한 가지 차이점이 있다.

`?` 연산자의 경우 에러값은 `from` 함수를 이용해 전달된다.

이 함수는 표준 라이브러리에 정의된 `From` 트레이트에 선언되어 있으며 에러를 어떤 한 타입에서 다른 타입으로 변환한다.
`?` 연산자가 `from` 함수를 호출하면 전달된 에러타입은 현재 함수의 반환타입에 정의된 에러 타입으로 변환된다.

각 에러타입이 자신을 반환에러타입으로 변환하는 `from` 함수를 구현하면 `?` 연산자가 그 변환을 자동으로 수행한다.

`?` 연산자를 이용하면 코드를 더 간결하게 유지할 수 있다.
게다가 `?` 연산자 다음에 메서드를 연결해서 호출하면 코드를 더 짧게 작성할 수 있다.

```rs
use std::io;
use std::io::Read;
use std::fs::File;

fn read_username_from_file() -> Result<String, io::Error> {
  let mut s = String::new();

  File::open("hello.txt")?.read_to_string(&mut s)?;
  Ok(s);
}
```

러스트는 이 작업을 더 편리하게 실행할 수 있는 `fs::read_to_string` 함수를 제공한다

```rs
use std::io;
use std::fs;

fn read_username_from_file() -> Result<String, io::Error> {
  fs::read_to_string("hello.txt")?;
}
```

#### `?` 연산자는 Result 타입을 반환하는 함수에서만 사용할 수 있다

`?` 연산자는 `Result` 타입을 반환하는 함수에 대해서만 사용할 수 있다.

`()` 타입을 반환하는 `main` 함수에 `?` 연산자를 사용하면 어떤일이 벌어지는지 보자

```rs
use std::fs::File;

fn main() {
  let f = File::open("hello.txt");
}
```

이 코드를 컴파일 하면 다음과 같은 에러메시지를 보게 된다

```sh
error[E0277]: the '?' operator can only be used in a function that returns 'Result' or 'Option'
(or another type that implements 'std::ops::try')
```

`Result` 타입을 반환하지 않는 함수안에서 `Result` 타입을 반환하는 다른 함수를 호출할 때는
호출자에게 에러를 전파할 가능성이 있는 `?` 연산자 보다는 `match` 표현식이나 `Result` 타입의 메서드 중 하나를 사용해서 `Result`를 처리해야 한다.

`main` 함수는 특별한 함수여서 이 함수가 반환할 수 있는 값의 타입에 제한이 있다.
`main` 함수가 반환할 수 있는 타입 중 하나는 `()`이며 `Result<T, E>` 타입을 반환할 수도 있다.

```rs
use std::error::Error;
use std::fs::File;

fn main() -> Result<(), Box<dyn Error>> {
  let f = File::open("hello.txt")?;
  Ok(())
}
```

`Box<dyn Error>` 타입은 모든 종류의 에러를 의미하는 트레이트 객체(trait object)라고 부르는 타입이다.

`main` 함수가 이 타입을 반환하면 `?` 연산자를 이용할 수 있다.

## 패닉에 빠질 것인가 말 것인가

드물지만 `Result` 타입을 반환하는 것보다는 그냥 패닉을 발생시키는 것이 더 적잘할 때도 있다.

### 예제, 프로토타입 코드, 그리고 테스트

`unwrap`과 `expect` 메서드는 실제 에러를 어떻게 처리해야 할 것인지 결정하기 앞서 프로토타이핑 해보는 상황에서 유용하다.

### 컴파일러보다 개발자가 더 많은 정보를 가진 경우

호출하는 작업이 특정 상황에서는 실패할 가능성이 없어도, 통상적으로는 작업이 실패할 수 있으므로
절대 `Err`값이 반환되지 않는 것을 확신하는 상황에서도 `unwrap` 메서드를 호출해주는 것이 좋다.

```rs
use std::new::IpAddr;
let home: IpAddr = "127.0.0.1".parse().unwrap();
```

### 에러 처리를 위한 조언

코드가 결국 잘못된 상태가 될 상황이라면 패닉을 발생시키는 것도 나쁜 선택은 아니다.
여기서 잘못된 상태라는 것은 어떤 가설이나 보장, 계약 혹은 불변이어야 할 것들이 깨진 상황을 말한다.

- 잘못된 상태는 원래 기대했던 동작이 어쩌다 실패하는 상황을 말하는 것이 아니다
- 어느 지점 이후의 코드는 프로그램이 절대 잘못된 상태에 놓이지 않아야만 정상적으로 동작한다
- 이 정보를 사용 중인 타입으로 표현할 방법이 없다

하지만 어떤 이유로 작업이 실패할 수 있다면 `panic!` 매크로 대신 `Result` 타입을 반환하는 것이 낫다.

parser에 잘못된 형식의 데이터가 전달되는 등의 상황에서 `Result` 타입을 반환해서
작업이 실패할 가능성이 있음을 명확히하고 잘못된 상태를 호출자에게 전파하여 호출자가 문제를 어떻게 해결하게 할 것인지 결정하게 하는 것이 좋다.

코드가 어떤 값에 대한 작업을 수행할 때는 그 값이 유효한지를 반드시 먼저 검사한 후 패닉을 발생시켜야 한다.
그 이유는 안정성 때문이며, 유효하지 않은 데이터를 기반으로 작업을 실행하면 코드가 취약점에 노출될 수 있다.
표준 라이브러리가 유효한 범위를 벗어나는 메모리를 접근할 때 패닉을 발생시키는 것도 그 이유이다.

함수가 패닉을 유발할 때는 반드시 API문서에 해당 함수에 관해 설명해야 한다.

### 유효성 검사를 위한 커스텀 타입

러스트의 타입 시스템이 유효한 값의 전달을 보장한다는 사실을 활용하기 위해 유효성 검사를 위한 커스텀 타입을 생성하는 방법에 대해 알아보자.

유효성 검사를 수행하는 방법 중 하나는 `u32` 타입 대신 `i32` 타입을 사용해서 음수를 입력해도 처리할 수 있도록 한 다음,
입력된 숫자가 범위 안의 값인지를 검사하는 것이다.

```rs
loop {
  // ...

  let guess: i32 = match guess.trim().parse() {
    Ok(num) => num,
    Err(_) => continue,
  }

  if (guess < 1 || guess > 100) {
    println!("!에서 100 사이의 값을 입력해 주세요");
    continue;
  }

  match guess.cmp(&secret_number);
  
  // ...
}
```

이 경우 `guess` 변수와 `secret_number` 변수를 비교하는 코드가 실행될 때 `guess` 변수값은 반드시 1에서 100 사이의 값이 된다.
이 방법은 이상적인 해결책이 아니다. 이 조건을 만족해야 하는 함수가 많아지면 `if`문을 이용해서 매번 검사하는 것은 비효율적이다.

더 나은 방법은 새로운 타입을 생성하고 이 타입의 인스턴스를 생성하는 함수에 유효성 검사 코드를 작성하는 것이다.

예제는 `Guess` 타입을 정의하고 값의 범위가 1에서 100사이인 경우에만 타입의 인스턴스를 생성하는 `new` 함수를 구현한 코드이다.

```rs
pub struct Guess {
  value: i32,
}

impl Guess {
  pub fn new(value: i32) -> Guess {
    if (value < 1 || value > 100) {
      panic!("값은 1과 100 사이의 값이어야 합니다");
    }

    Guess {
      value
    }
  }

  pub fn value(&self) -> i32 {
    self.value
  }
}
```

이 경우 `Guess:new` 함수가 패닉을 발생시키는 조건은 반드시 공개된 API문서에 설명되어야 한ㄷ.
