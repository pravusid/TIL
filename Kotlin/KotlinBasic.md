# Kotlin

> Kotlin in Action을 읽으며 발췌, 요약 및 기타내용 추가하였음

Kotlin은 IDE를 개발하는 Jetbrain에서 JVM호환 언어로 개발중이다.
현재 VM없이 네이티브로 작동하는 Kotlin을 개발중이다.

Kotlin은 개발팀 대부분이 살고있는 St.petersburg 근처의 섬 이름이다.

Android Studio 3.0부터 Java와 함께 정식지원 언어로 등록되었다.

## Kotlin 특성

- Open Source 프로젝트이다
- Java 코드와의 상호운용성을 중시함: 심지어 Java 코드와 섞어 쓸 수도 있음
- 따라서 서버, 안드로이드등 자바가 실행되는 모든곳에 적용 가능하다.
- JavaScript로 컴파일 할 수 있다.
- Native Kotlin을 개발중이다.
- 정적타입 지정언어이다, 그러나 타입추론(Type Inference)을 지원한다.
- 자바와 마찬가지로 클래스, 인터페이스, 제네릭스를 지원한다.
- 자바와 다르게 NULLABLE TYPE을 지원한다 (NPE여부 검사 가능)
- 함수형 언어의 특징을 일부 지원한다.

### 함수형 프로그래밍의 핵심개념

- 함수가 일급시민(first-class)이다: 함수를 변수에 할당하고, 인자로 다른 함수에 전달하거나 반환값으로 받을 수 있다.
- 불변성(Immutablility): 함수형언어에서는 일단 만들어지고 나면 바뀌지 않는 불변객체를 사용한다.
- no side-effect: 함수가 다른객체의 상태를 변경하지 않으며 함수 외부와 상호작용하지 않는 순수함수를 사용한다.

함수형 프로그래밍의 장범

- 명령형(imperative) 코드에 비해 간결하다
- 다중 스레드를 사용해도 안전하다: 불변객체를 사용하기 때문
- 테스트하기 용이하다: 부수효과 없는 함수를 사용하기 때문

### Kotlin에서 지원하는 함수형 프로그래밍

- 함수가 다른함수를 파라미터로 받거나 함수가 새로운 함수를 반환할 수 있다.
- 람다식을 지원한다
- 불변객체를 간편하게 만들 수 있는 문법을 제공한다.
- Kotlin 표준라이브러리는 객체와 컬렉션을 함수형으로 다룰 수 있는 API를 제공한다.

### 자바와 상호 호완 운용

Kotlin을 컴파일하면 class 파일을 생성하지만 실행을 위해서는 Kotlin runtime library에 의존하므로
Kotlin runtime library로 함께 배포해야 한다.

## Kotlin 기초

### 함수와 변수

Hello World! 예제

```kotlin
fun main(args: Array<String>) {
  println("Hello, world!")
}
```

- 함수선언: `fun`
- 타입은 이름 뒤에 명시
- 함수를 최상위에 정의가능: 함수가 반드시 클래스안에 존재할 필요 없음.
- 배열도 일반클래스로 처리: 배열처리를 위한 문법 `Object[]`이 따로 존재하지 않는다.
- `System.out.println` 대신 `println` 이라는 wrapper를 제공한다.
- 줄 끝에 세미콜론(`;`)을 붙이지 않아도 된다.

#### 함수

블록이 본문인 함수

```kotlin
fun max(a: Int, b: Int): Int {
  return if (a > b) a else b
}
```

위의 함수를 풀어쓰면 다음과 같다

```text
fun 함수명(파라미터): 반환타입 {
  함수본문
}
```

또한 Kotlin에서는 if가 statement(문)이 아닌 expression(식)으로 연산자 처럼 사용될 수 있다.

따라서 위의 함수를 보다 간결하게 표현하면 아래와 같다 (식이 본문인 함수)

```kt
fun max(a: Int, b, Int): Int = if (a > b) a else b
```

식이 본문인 함수는 타입추론을 사용하여 반환값을 생략할 수 있다. (블록이 본문인 함수는 해당하지 않음)

```kt
fun max(a: Int, b: Int) = if (a > b) a else b
```

#### 변수

Kotlin에서는 타입지정을 생략하는 경우가 흔하므로 타입을 이름 뒤쪽에 작성하도록 되어있다.

초기값이 있으면 타입생략이 가능하지만, 초기화 하지않고 변수를 선언하기 위해서는 타입을 반드시 명시해야한다.

```kt
//가능
val answer = 42

// 불가능
val answer
answer = 42
```

변수 선언은 두 가지 키워드를 사용하여 할 수 있다.

##### val (value): 변경불가능한(immutable) 참조를 저장하는 변수, 자바의 final에 해당

기본적으로 불변변수인 val 키워드로 선언하는 것이 좋다.
그러나 val 참조자체는 불변이라도 참조가 가리키는 객체의 내부 값은 변경가능하다.

val은 블록내에서 정확히 **한 번만 초기화** 되어야 한다.

##### var (variable): 변경가능한(mutable) 참조를 저장하는 변수

var 키워드를 사용하면 변수 값을 변경할 수 있지만 타입은 바뀌지 않는다. (형 변환을 해야 함)

#### 문자열 템플릿

```kt
fun main (args: Array<String>) {
  val name = if (args.size > 0) args[0] else "Kotlin"
  println("Hello, ${name}!")
}
```

변수를 문자열 내부에서 사용하려면 `${}` 기호를 추가하면 된다.
> String의 경우 괄호(`{}`)를 붙이지 않아도 되지만, 식(expression)이나 한글이 붙는경우 error가 발생할 수 있다.

`$`기호를 사용하려면 escape 기호를 사용해야한다. (`\$`)

### 클래스와 프로퍼티

```java
public class Person {
  private final String name;

  public Person(String name) {
    this.name = name;
  }

  public String getName() {
    return name;
  }
}
```

위의 자바로 작성한 내용과 아래의 Kotlin 코드는 동일한 기능을 수행한다.

```kt
class Person(val name: String)
```

Kotlin에서는 프로퍼티를 간결하게 기술할 수 있는 구문을 제공한다. 또한 Kotlin의 기본 접근제한자는 public이므로 생략가능하다.

자바에서 멤버필드의 접근제한자는 캡슐화를 위해 private으로 두고 접근자(getter) 메소드를 사용한다.

```kt
class Person(
  val name: String,
  var isMarried: Boolean
)
```

- val로 명시된 프로퍼티는 읽기전용으로 public getter를 생성한다.
- var로 명시된 프로퍼티는 public getter, public setter를 생성한다.

그러나 접근자 사용은 property 직접 접근형태의 문법을 사용한다. (실제 직접접근은 아님)
getter 뿐만 아니라 setter의 경우도 동일하다.

```kt
val person = Person("Bob", true)
println(person.name)
println(person.isMarried)
```

#### 커스텀 접근자

접근자를 직접 작성할 수도 있다.

```kt
class Rectangle(val Height: Int, val width: Int) {
  val isSquare: Boolean
    get() {
      return height == width
    }
}
```

접근자 사용은 동일하다. `println(rectangle.isSquare)`

### enum과 when

#### enum

```kt
enum class Color {
  RED, GREEN, BLUE
}
```

kotlin에서 enum은 soft keyword로 class 키워드 앞에 붙는다.

프로퍼티와 메소드가 있는 enum 클래스를 선언할 수도 있다.

```kt
enum class Color(
  val red: Int, val green: Int, val blue: Int
) {
  RED(255, 0, 0),
  GREEN(0,255,0),
  BLUE(0, 0, 255); // Kotlin에서 유일하게 세미콜론 생략이 불가능하다.

  fun rgb() = (red * 256 + green) * 256 + blue
}
```

#### when

when은 switch문에 해당하는 문법이다.

```kt
fun getMnemonic(color: Color) =
  when (color) {
    Color.RED -> "Richard"
    Color.GREEN -> "Gave"
    Color.BLUE -> "Battle"
  }
```

분기안에서 여러 값을 사용할 수도 있다.

```kt
fun getMnemonic(color: Color) =
  when (color) {
    Color.RED, Color.GREEN -> "Richard"
    Color.BLUE -> "Battle"
  }
```

자바의 switch와 다르게 각 분기에 break;를 넣지 않아도 되며 분기조건에 임의의 객체를 허용한다.

```kt
fun mix(c1: Color, c2: Color) =
  when (setOf(c1, c2)) {
    setOf(RED, YELLOW) -> ORANGE
    setOf(YELLOW, BLUE) -> GREEN
    setOf(BLUE, VIOLET) -> INDIGO
    else -> throw Exception("color exception")
  }
```

객체를 통한 분기검사는 동등성(equility)비교를 통해 이루어진다.

when은 인자 없이 사용할 수도 있다. 이 경우 if / elseif 문과 유사한 형식이 된다.

```kt
fun mix(c1: Color, c2: Color) =
  when {
    (c1 == RED && c2 == YELLOW) ||
    (c1 == YELLOW && c2 == RED) -> ORANGE
    (c1 == YELLOW && c2 == BLUE) ||
    (c1 == BLUE && c2 == YELLOW) -> GREEN
    (c1 == BLUE && c2 == VIOLET) ||
    (c1 == VIOLET && c2 == BLUE) -> INDIGO
    else -> throw Exception("color exception")
  }
```

### 스마트 캐스트: 타입검사 + 타입캐스트

```kt
interface Expr
class Num(val value: Int): Expr
class Sum(val left: Expr, val right: Expr): Expr
```

`(1 + 2) + 4`라는 식을 `Sum(Sum(Num(1), Num(2), Num(4)))`라는 구조의 객체로 생성할 수 있다.
계산을 위해서 위의 내용을 실제로 구현하면

```kt
fun eval(e: Expr): Int =
  when (e) {
    // 스마트 캐스트가 적용된 상태
    is Num -> {
      println("num: ${e.value})
      e.value // 블록의 마지막 식이므로 e의 타입이 Num이면 e.value가 반환된다.
    }
    is Sum -> {
      val left = eval(e.left)
      val right = eval(e.right)
      println("sum: ${left} + ${right}")
      eval(e.right) + eval(e.left) // 다음식이 반환된다
    }
    else ->
      throw IllegalArgumentException("Unknown expression")
  }
```

Kotlin은 `is`를 사용해 변수 타입을 검사하며 이는 다른언어의 `instanceOf`와 유사하다.
또한 Kotlin에서는 `is`를 통해 확인한 변수는 컴파일러가 캐스팅해준다.

### 이터레이션: while, for

#### while

자바와 동일

#### for

반복의 범위를 지정하기 위해서 범위 연산자 `..`를 사용한다.
<https://kotlinlang.org/docs/reference/ranges.html>

`val oneToTen = 1..10`의 경우 1과 10을 포함하는(닫힌 구간) 10회의 반복 범위를 의미한다.

```kt
for (i in 100 downTo 1 step 2) {
  print(i)
}
```

반복문 연산자 (1 to 10)

- `i in 1..10`: `1 <= i <= 10`
- `i in 1 until 10`: `1 <= i < 10`
- `i in 10 downTo 1`: `100 >= i >= 1`
- `step n`: n씩 건너뛴다 (기본값: 정방향 1, 역방향 -1)

#### map에 대한 iteration

```kt
val binaryReps = TreeMap<Char, String>()

// assignment
for (c in 'A'..'F') { // 문자열도 순서대로 범위지정이 가능하다
  val binary = Integer.toBinaryString(c.toInt())
  binaryReps[c] = binary // 맵에 값 추가
}

// iteration
for ( (letter, binary) in binaryReps ) {
  println("${letter} = ${binary})
}
```

#### in으로 원소검사

`java.lang.Comparable` 인터페이스를 구현 클래스라면 해당 클래스의 인스턴스 객체를 사용해 범위를 만들 수 있다.

```kt
// 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
fun isLetter(c: Char) = c in 'a'..'z' || c in 'A'..'Z'
fun isNotDigit(c: Char) = c !in '0'..'9'
```

`in` 및 `!in` 연산자를 `when`에서 사용할 수 있다.

```kt
fun recognize(c: Char) = when(c) {
  in '0'..'9' ->
    "it's a digit!"
  in 'a'..'z', in 'A'..'Z' ->
    "it's a letter!"
  else ->
    "I have no idea"
}
```

### 예외처리

자바와 달리 코틀린의 `throw`는 식이므로 다른식에 포함될 수 있다.

```kt
val percentage =
  if (number in 0..100)
    number
  else
    throw IlleagalArumentExcption("error!")
```

try catch finally는 자바와 유사하지만 method 뒤의 `throws Exception` 작성여부가 다르다
자바와 다르게 kotlin은 checked exception과 unchecked exception을 구분하지 않는다.

이는 자바 프로그래머들이 체크예외를 사용하는 방식을 고려한 것이다.
자바는 체크 예외처리를 강제하지만 많은 프로그래머들이 의미 없이 예외를 다시 던지거나,
예외를 잡되 처리하지는 않고 그냥 무시하는 코드를 작성하는 경우가 흔한 문제가 있다.

```kt
fun readNumber(reader: BufferedReader): Int? { //
  try {
    val line = reader.readLine()
    return Integer.parseInt(line)
  } catch(e: NumberFormatException) {
    return null
  } finally {
    reader.close()
  }
}
```

try를 식으로 사용할 수도 있다.

```kt
fun readNumber(reader: BufferedReader) {
  val number = try {
    Integer.parseInt(reader.readLine())
  } catch (e: NumberFormatException) {
    null // 예외발생시 null 값이 반환된다.
  }
  println(number)
}
```

## 함수 정의와 호출
