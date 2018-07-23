# 함수형 프로그래밍

## 함수형 프로그래밍 패러다임

- 명령형 프로그래밍
  - modifying mutable variables
  - using assignments
  - control structures such as if-then-else, loops, brek, continue, return
- 함수형 프로그래밍
- 논리형 프로그래밍

객체지향프로그래밍은 여러 패러다임의 교차점에 있음 (명령형의 특성을 많이 가지고 있다)

### 함수형 프로그래밍 특징

명령형 프로그래밍 -> 폰 노이만 병목 문제

- one or more data types
- operations on these types
- laws that describe the relationships between values and operations
- Normally, a theory does not describe mutations

함수를 추상화하고 함수를 합성하기 위한 방법

- 좁은의미로 FP는 mutable variable, 할당, 루프 같은 명령형 제어문이 없다.
- 넓은 의미로 함수에 집중한다
- 특히 함수는 변수처럼 생산되고 소비되고 합성되는 값이 된다: 함수는 first-class citizens
  - 함수는 다른 함수 내부와 같은 어느곳에서나 정의된다
  - 다른 값들과 마찬가지로 함수도 함수의 파라미터와 반환값으로 전달된다
  - 다른값들과 마찬가지로 함수를 합성하는 연산자가 존재한다
- 함수형 프로그래밍은 모듈화 되고 병렬처리에 유리하다.

### 함수형 언어의 종류

엄격한 의미의 함수형 언어

- Pure Lisp, XSLT, XPath, Xquery, FP
- Haskell (without I/O Monad or Unsafe PerformIO)

넓은 의미의 함수형 언어

- Lisp, Scheme, Racket, Clojure
- SML, Ocaml, F#
- Haskell (full language)
- Scala
- Smalltalk, Ruby

## Elements of Programming

### REPL (Read-Eval-Print Loop)

`sbt console`

### Evaluation

non-primitive 표현식은 다음과 같이 평가한다

1. 가장 왼쪽 연산자를 가져온다
2. 피연산자 평가 (왼쪽에서 오른쪽으로)
3. 피연산자에 연산자를 적용합니다.

산술 연산 평가 예시

```scala
(2 * pi) * radius
(2 * 3.14159) * radius
6.28318 * radius
6.28318 * 10
62.8318
```

### Parameter Evaluation

매개변수가 있는 함수도 연산자와 비슷한 방식으로 평가를 적용함

```scala
sumOfSquares(3, 2+2)
sumOfSquares(3, 4)
square(3) + square(4)
3 * 3 + square(4)
9 + square(4)
9 + 4 * 4
9 + 16
25
```

### The substitution model

표현식을 값으로 줄이는 일

Side effect가 없는 모든 표현식에 적용될 수 있다.

> 반환 값의 `c++`와 같은 식은 side effect를 가진다

Substitution model은 함수형 프로그래밍의 기초가 되는 lambda-calculus로 공식화 될 수 있다.

모든 표현식이 값으로 줄일 수 있지는 않다: 유한한 단계가 있는 경우만 가능(무한루프X)

### 평가 전략의 변경

인터프리터는 fuction application을 다시 쓰기 전 함수 인자를 값으로 줄인다.

대신 함수를 줄이지 않은 인자에 적용할 수 있다.

```scala
sumOfSquares(3, 2+2)
square(3) + square(2+2)
3 * 3 + square(2+2)
9 + square(2+2)
9 + (2+2) * (2+2)
9 + 4 * (2+2)
9 + 4 * 4
25
```

### Call by Value, Call by Name

값에 의한 호출: 인자의 값이 모두 평가된 후 파라미터로 전달된다.

이름에 의한 호출: 파라미터로 표현식이 전달되어 호출될 때마다 식이 평가됨

스칼라에서 Call by Name으로 파라미터를 정의할 때는 파라미터 타입 앞에 `=>`를 붙인다.

만약 줄여야할 표현식이 순수 함수로 이루어져 있고 종료점이 있다면 call by value와 call by name의 최종결과는 같다.

Call by value는 함수 인자가 한번만 평가된다는 측면에서 장점이 있다.

Call by name은 대응하는 파라미터가 함수 본문에서 사용되지 않을 때 인자가 평가되지 않는다는 측면에서 장점이 있다.

```scala
def test(x: Int, y: Int) = x * x

test(7, 2*4)

// call by value
test(7, 8)
7 * 7
49

// call by name
7 * 7
49
```

## Evaluation Strategies and Termination

표현식의 Call by Value 평가가 종료된다면 표현식의 Call by Name 평가도 종료된다.

역은 성립하지 않는다.

CBN은 끝나고 CBV는 끝나지 않는 경우

```scala
def first(x: Int, y: Int) = x

// consider the expression
first(1, loop)
```

## Conditional and Value Definitions

### Conditional Expressions

`if-else`: It is used for expressions not statements

Boolean 표현식은 자바와 같다.

## Blocks and Lexical Scope

### 중첩 함수

작은 함수로 많이 쪼개는 것은 좋은 함수형 프로그래밍 스타일이다

그러나 함수이름(예제참고)인 `sqrtIter`, `improve`, `isGoodEnough`는 `sqrt` 구현에만 사용된다.
보통 우리는 사용자가 이러한 함수에 직접 접근하기를 원하지 않는다.

네임스페이스 오염을 막기위해서 보조 함수들을 `sqrt`함수의 내부로 넣을 수 있다.

```scala
def sqrt(x: Double) = {
  def sqrtIter(guess: Double): Double =
    if (isGoodEnough(guess)) guess
    else sqrtIter(improve(guess))

  def improve(guess: Double) =
    (guess + x / guess) / 2

  def isGoodEnough(guess: Double) =
    abs(square(guess) - x) / x < 0.001

  sqrtIter(1.0)
}
```

### Semicolons and infix operators

세미콜론 `;`을 사용하면 문장의 종료를 알린다.

문장의 마지막에 연산자 (`+` ...)가 나온다면 다음줄에 이어지는 문장이 있다는 뜻이다.

## Tail Recursion / Tail Call

Loop 기능, 깊은 재귀 체인을 피할 수 있다

### 최대 공약수

유클리드 알고리즘을 이용한 최대공약수 구하기

```scala
def gcd(a: Int, b: Int): Int =
  if (b == 0) else gcd(b, a % b)

// 아래와 같이 평가된다
gcd(14, 21)
if (21 == 0) 14 else gcd(21, 14 % 21)
if (false) 14 else gcd(21, 14 % 21)
gcd(21, 14 % 21)
gcd(21, 14)
if (14 == 0) 21 else gcd(14, 21 % 14)
gcd(14, 21 % 14)
gcd(14, 7)
...
gcd(7, 0)
if (0 == 0) 7 else gcd(0, 7 % 0)
7
```

### 팩토리얼

```scala
def factorial(n: Int): Int =
  if (n == 0) 1 else n * factorial(n - 1)

// 아래와 같이 평가된다
factorial(4)
if (4 == 0) 1 else 4 * factorial(4 - 1)
4 * factorial(3)
4 * factorial(3 * factorial(2))
4 * factorial(3 * factorial(2 * factorial(1)))
4 * factorial(3 * factorial(2 * factorial(1 * factorial(0))))
4 * factorial(3 * factorial(2 * factorial(1 * 1)))
...
120
```

팩토리얼을 꼬리 재귀방식으로 쓰면

```scala
def tailRecursionFactorial(n: Int): Int = {
  def loop(acc: Int, n: Int): Int =
    if (n == 0) acc
    else loop(acc * n, n - 1)

  loop(1, n)
}
```

### tail recursion 확인

`@tailrec` 애노테이션을 사용하여 꼬리재귀로 사용가능하고 그대로 작성되었는지 컴파일러가 확인
