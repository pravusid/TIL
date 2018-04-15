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

함수예제

```kt
fun <T> joinToString(
  collection: Collection<T>,
  separator: String,
  prefix: String,
  postfix: String
): String {
  val result = StringBuilder(prefix)
  for ((index, element) in collection.withIndex()) {
    if (index > 0) result.append(separator)
    result.append(element)
  }
  result.append(postfix)
  return result.toString()
}
```

### 파라미터

#### 파라미터 이름

위에서 만든 함수 호출시 가독성을 위해서 인자의 이름을 명시할 수 있다.
인자를 명시하고 나면 그 뒤에 오는 모든인자의 이름을 명시해야 한다.

```kt
joinToString(collection, separator = " ", prefix = " ", postfix = ".")
```

#### 디폴트 파라미터

코틀린에서는 디폴트 파라미터를 이용하여 불필요한 Method Overloading을 피할 수 있다.

```kt
fun <T> joinToString(
  collection: Collection<T>,
  separator: String = ", ",
  prefix: String = "",
  postfix: String = ""
): String
```

collection과 separator 호출시 뒤의 인수 두개는 생략가능

`joinToString(list, "; ")`

중간 인자를 생략하고 뒤의 인수를 호출한다면 인자 이름을 붙여서 호출해야한다.

`joinToString(list, postfix = ";", prefix="# ")`

> 만약 자바에서 코틀린 함수를 호출 할 때 디폴트 파라미터로 오버로딩한 함수를 호출하고 싶다면 `@JvmOverloads` 코틀린 함수 어노테이션을 사용하면 된다.

### 최상위 함수와 최상위 프로퍼티

#### 최상위 함수

코틀린에서는 특정한 클래스에 포함할 필요가 없는 메소드 (유틸성 메소드)를 클래스 밖에서 선언할 수 있다.

```kt
@file:JvmName("StringFunctions") // 클래스 이름을 파일명과 다르게 할 때 어노테이션을 사용할 수 있다.
package strings
fun joinToString(...): String { ... }
```

이 경우 바이너리 코드로 컴파일되면 소스가 명시된 파일명으로 생성된 클래스의 static method로 작동한다.

#### 최상위 프로퍼티

프로퍼티 역시 최상위 수준에 놓을 수 있다.

```kt
var opCount = 0
fun performOperation() {
  opCount++
}
```

정적필드에 해당하는 프로퍼티도 최상위에 놓을 수 있다

`val UNIX_LINE_SEPARATOR = "\n"`

val로 선언한 변수의 경우 getter로 접근해야 하기 때문에 상수로 사용하기에 자연스럽지 않다.

`const val UNIX_LINE_SEPARATOR = "\n"`

const 변경자로 자바의 `public static final` 필드로 컴파일 할 수 있다.
단 자바의 primary type과 String type의 프로퍼티만 const로 지정할 수 있다.

### 확장 함수

확장함수는 어떤 클래스의 멤버 메소드인것처럼 호출할 수 있지만 그 클래스 밖에 선언된 함수이다.
기존 자바 API를 재작성하지 않고도 코틀린이 제공하는 여러 기능을 사용하는 용도로 활용될 수 있다.

확장함수를 만드려면 추가하려는 함수 이름 앞에 그 함수가 확장할 클래스의 이름을 덧붙이기만 하면 된다.
클래스 이름을 receiver type 이라 부르며 확장함수가 호출되는 대상이 되는 값을 receiver object라 부른다.

```kt
package strings

fun String.lastChar(): Char = this.get(this.length -1)
```

함수를 호출하는 구문은 다른 일반 클래스 멤버를 호출하는 구문과 같다.

`println("Kotlin".lastChar())`

Java뿐만 아니라 Groovy등 다른 JVM언어로 작성된 클래스도 확장할 수 있다.

확장함수 내부에서 일반적인 인스턴스 메소드의 내부와 마찬가지로 receiver object의 메소드나 프로퍼티를 바로 사용할 수 있다.
하지만 확장함수가 캡슐화를 깨지는 않기 때문에 클래스 내부의 private, protected 멤버는 사용할 수 없다.

확장함수를 정의했다고 해도 자동으로 모든 소스코드에서 그 함수를 사용할 수 있지는 않다.
확장함수를 사용하기 위해서는 임포트를 해야 한다. 코틀린에서는 클래르르 임포트할 때와 동일한 구문으로 개별 함수를 불러올 수 있다.

```kt
import strings.lastChar

val c = "Kotlin".lastChar()
```

`as` 키워드를 사용하면 임포트한 클래스나 함수를 다른이름으로 부를 수 있다.

```kt
import strings.lastChar as last

val c = "Kotlin".last()
```

#### 자바에서 확장 함수 호출

내부적으로 확장 함수는 수신객체를 첫 번째 인자로 받는 정적메소드이다.
그래서 확장함수를 호출해도 다른 어댑터 객체나 실행시점 부가비용이 들지 않는다.

따라서 자바에서 확장함수를 사용하기 위해서는 단지 정적 메소드를 호추하면서 첫 번째 인자로 수신객체를 넘기면 된다.

확장함수가 `StringUtil.kt` 파일에 정의 되어있다면

```java
char c = StringUtilKt.lastChar("Java");
```

#### 확장함수로 유틸리티 함수 정의

joinToString 함수 예제

```kt
fun <T> Collection<T>.joinToString(
  separator: String = ", ",
  prefix: String = "",
  postfix: String = ""
): String {
  val result = StringBuilder(prefix)

  for ((index, element) in this.withIndex()) {
    if (index > 0) result.append(separator)
    result.append(element)
  }
  result.append(postfix)
  return result.toString()
}
```

실행시 결과

```kt
>>> val list = listOf(1, 2, 3)
>>> println(list.joinToString(separator="; ", prefix = "(", postfix = ")"))
(1; 2; 3)
```

확장함수는 적적 메소드 호출에 대한 syntatic sugar일 뿐이다.
그래서 클래스가 아닌 더 구체적인 타입을 수신객체 타입으로 지정할 수 있다.

문자열의 컬렉션에 대해서만 호출할 수 있는 join 함수를 정의하고 싶다면 다음과 같이하면 된다.

```kt
fun Collection<String>.join(
  separator: String = ", ",
  prefix: String = "",
  postfix: Strig = ""
) = joinToString(separator, prefix, postfix)

>>> println(listOf("one", "tow", "eight").join(" "))
one two eight
```

위의 함수를 String이 아닌 객체의 리스트에 대해 호출할 수 없다.

```kt
>>> listof(1, 2, 8).join()
Error: Type mismatch: inferred type is List<Int> but Collection<String> was expected.
```

#### 확장함수는 오버라이드할 수 없다

확장함수는 클래스의 일부가 아니기 때문에 (static 이므로) 오버라이드 할 수 없다.

### 확장 프로퍼티

확장프로퍼티를 사용하면 기존 클래스 객체에 대한 프로퍼티 형식의 구문으로 사용할 수 있는 API를 추가할 수 있다.
프로퍼티라는 이름으로 불리기는 하지만 상태를 저장할 적절한 방법이 없기 때문에 (기존 클래스의 인스턴스에 필드를 추가할 방법은 없으므로)
실제로 확장 프로퍼티는 아무 상태도 가질 수 없지만, 확장 프로퍼티 문법으로 더 짧은 코드륵 작성할 수 있다.

앞에서 정의한 lastChar라는 함수를 프로퍼티로 변경해 보자.

```kt
val String.lastChar: Char
  get() = get(length - 1)
```

확장함수의 경우와 마찬가지로 확장 프로퍼티도 일반적인 프로퍼티와 같은데, 단지 수신 객체 클래스가 추가됐을 뿐이다.
뒷받침 필드가 없어 기본 게터 구현을 제공할 수 없으므로 getter는 반드시 정의해야 한다.
마찬가지로 초기화 코드에서 계산한 값을 담을 장소가 없으므로 초기화 코드도 쓸 수 없다.

```kt
var StringBuilder.lastChar: Char
  get() = get(length - 1)
  set(value: Char) {
    this.setCharAt(length - 1, value)
  }
```

사용법은 멤버 프로퍼티 사용법과 동일하다

```kt
>>> println("Kotlin".lastChar)
n
>>> val sb = StringBuilder("Kotlin?")
>>> sb.lastChar = '!'
```

자바에서 확장 프로퍼티를 사용하고 싶다면 항상 `StringUtilKt.getLastChar("Java");` 처럼
게터나 세터를 명시적으로 호출해야 한다.

### 컬렉션 처리

추가된 코틀린 컬렉션 특성

- vararg 키워드를 사용하면 호출 시 인자 개수가 달라질 수 있는 함수를 정의할 수 있다.
- 중위(infix) 함수 호출 구문을 사용하면 인자가 하나뿐인 메소드를 간편하게 호출할 수 있다.
- 구조분해선언(destructuring declaration)을 사용하면 복합적인 값을 분해해서 여러 변수에 나눠 담을 수 있다.

#### 가변인자 함수: 파라미터 개수가 달라질 수 있는 함수를 정의하는 방법

리스트를 생성하는 함수를 호출할 때 원하는 만큼 원소르 전달할 수 있다.

`val list = listOf(2, 3, 4, 7, 11)`

라이브러리에서 이 함수의 정의를 보면 다음과 같다

`fun listOf<T>(vararg values: T): List<T> { ... }`

자바에도 가변길이 인자 (...)가 존재한다.
가변길이 인자는 메소드를 호출할 때 원하는 개수만큼 값을 인자로 넘기면 자바 컴파일러가 배열에 그 값들을 넣어주는 기능이다.

이미 배열에 들어있는 원소를 가변 길이 인자로 넘길 때도 코틀린과 자바 구문이 다르다.
자바에서는 배열을 그냥 넘기면 되지만 코틀린에서는 배열을 명시적으로 풀어서 배열의 각 원소가 인자로 전달되게 해야한다.
기술적으로는 스프레드 연산자가 그런 작업을 해준다. 하지만 실제로는 전달하려는 배열앞에 `*`를 붙이기만 하면된다.

```kt
fun main(args: Array<String>) {
  val list = listOf("args: ", *args)
  pringln(list)
}
```

#### 중위 호출과 구조 분해 선언

맵을 만드려면 mapOf 함수를 사용한다

`val map = mapOf(1 to "one", 7 to "seven", 53 to "fifty-three")`

`to`라는 단어는 코틀린 키워드가 아니라 중위호출(infix call)이라는 특별한 방식으로 to라는 일반 메소드를 호출한 것이다.

중위 호출시에는 수신 객체와 유일한 메소드 인자 사이에 메소드 이름을 넣는다.

```kt
1.to("one") // to 메소드를 일반적인 방법으로 호출
1 to "one" // to 메소드를 중위 호출 방식으로 호출
```

인자가 하나뿐인 일반 메소드나 인자가 하나뿐인 확장함수에 중위호출을 사용할 수 있다.
함수를 중위 호출에 사용하게 허용하고 싶으면 infix 변경자를 함수 선언 앞에 추가해야한다.

`infix fun Any.to(other: Any) = Pair(this, other)`

위의 to 함수는 Pair의 인스턴스를 반환한다. Pair는 코틀린 표준 라이브러리 클래스로 두 원소로 이루어진 순서쌍을 표현한다.

위의 함수와 Pair의 내용으로 두 변수를 초기화 할 수 있다.

`val (number, name) = 1 to "one"`

이런 기능을 구조 분해 선언이라고 부른다.

Pair 인스턴스 외에도 구조분해를 적용할 수 있다.
예를 들어 key와 value라는 두 변수를 맵의 원소를 사용해 초기화 할 수 있다.

루프문에서도 구조 분해 선언을 활용할 수 있다.

```kt
for ((index, element) in collection.withIndex()) {
  println("${index}: ${element"}")
}
```

to 함수는 확장 함수이다. to를 사용하면 타입과 관계없이 임의의 순서쌍을 만들수 있는데 이는 to의 수신객체가 제네릭 하다는 것이다.
mapOf 함수의 선언을 살펴보면 다음과 같다

`fun<K, V> mapOf(vararg values: Pair<K, V>): Map<K, V>`

### 문자열과 정규식 다루기
