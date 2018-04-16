# 함수 정의와 호출

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

## 파라미터

### 파라미터 이름

위에서 만든 함수 호출시 가독성을 위해서 인자의 이름을 명시할 수 있다.
인자를 명시하고 나면 그 뒤에 오는 모든인자의 이름을 명시해야 한다.

```kt
joinToString(collection, separator = " ", prefix = " ", postfix = ".")
```

### 디폴트 파라미터

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

## 최상위 함수와 최상위 프로퍼티

### 최상위 함수

코틀린에서는 특정한 클래스에 포함할 필요가 없는 메소드 (유틸성 메소드)를 클래스 밖에서 선언할 수 있다.

```kt
@file:JvmName("StringFunctions") // 클래스 이름을 파일명과 다르게 할 때 어노테이션을 사용할 수 있다.
package strings
fun joinToString(...): String { ... }
```

이 경우 바이너리 코드로 컴파일되면 소스가 명시된 파일명으로 생성된 클래스의 static method로 작동한다.

### 최상위 프로퍼티

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

## 확장 함수

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

### 자바에서 확장 함수 호출

내부적으로 확장 함수는 수신객체를 첫 번째 인자로 받는 정적메소드이다.
그래서 확장함수를 호출해도 다른 어댑터 객체나 실행시점 부가비용이 들지 않는다.

따라서 자바에서 확장함수를 사용하기 위해서는 단지 정적 메소드를 호추하면서 첫 번째 인자로 수신객체를 넘기면 된다.

확장함수가 `StringUtil.kt` 파일에 정의 되어있다면

```java
char c = StringUtilKt.lastChar("Java");
```

### 확장함수로 유틸리티 함수 정의

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

### 확장함수는 오버라이드할 수 없다

확장함수는 클래스의 일부가 아니기 때문에 (static 이므로) 오버라이드 할 수 없다.

## 확장 프로퍼티

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

## 컬렉션 처리

추가된 코틀린 컬렉션 특성

- vararg 키워드를 사용하면 호출 시 인자 개수가 달라질 수 있는 함수를 정의할 수 있다.
- 중위(infix) 함수 호출 구문을 사용하면 인자가 하나뿐인 메소드를 간편하게 호출할 수 있다.
- 구조분해선언(destructuring declaration)을 사용하면 복합적인 값을 분해해서 여러 변수에 나눠 담을 수 있다.

### 가변인자 함수: 파라미터 개수가 달라질 수 있는 함수를 정의하는 방법

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

### 중위 호출과 구조 분해 선언

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

## 문자열과 정규식 다루기

코틀린 문자열은 자바 문자열과 같기 때문에 코틀린 문자열을 임의의 자바 메소드에 넘겨도 되며,
자바 코드에서 받은 문자열을 아무 코틀린 표준 라이브러리 함수에 전달해도 문제없다.

### 문자열 나누기

자바 String의 split 메소드는 입력받는 조건으로 정규식을 사용한다.

코틀린에서는 다른 조합의 파라미터를 받는 split 확장함수를 제공한다.
정규식을 파라미터로 받는 함수는 String이 아닌 Regex타입의 값을 받는다.

```kt
>>> println("12.345-6.A".split("\\.|-".toRegex()))
[12, 345, 6, A]

>>> println("12.345-6.A.split(".", "-"))
[12, 345, 6, A]
```

### 삼중 따옴표

삼중 따옴표는 여러줄에 걸쳐 문자열을 사용할 수 있고 어떤문자도 역슬래시를 사용하여 escape 할 필요가 없다.

삼중 따옴표를 사용하여 여러줄에 걸쳐 문자열을 사용할 때 들여쓰기를 하려면 TrimMargin(`|`)을 사용하면 된다.

```kt
al withoutMargin1 = """ABC
                |123
                |456""".trimMargin()
println(withoutMargin1) // ABC\n123\n456

val withoutMargin2 = """
    #XYZ
    #foo
    #bar
""".trimMargin("#")
println(withoutMargin2) // XYZ\nfoo\nbar
```

## 로컬함수와 확장

자바에서 메소드 추출 리팩토링을 적용하여 각 부분을 재활용할 수 있지만,
그렇게 하면 클래스안에 작은 메소드간의 관계를 파악하기 어려워져 코드를 이해하기 어려워질 수 있다.

리팩토링간 추출한 메소드를 별도의 내부클레스 안에 넣을 수도 있지만 코틀린에서는 보다 좋은 해법이 있다.
코틀린에서는 함수에서 추출한 함수를 원 함수 내부에 중첩시킬 수 있다.

중복된 코드가 있는 예제

```kt
class User(val id: Int, val name: String, val address: String)

fun saveUser(user: User) {
  if (user.name.isEmpty()) {
    throw IllegalArgumentException(
      "Can't save user ${user.id}: empty Name"
    )
  }

  if (user.address.isEmpty()) {
    throw IllegalArgumentException(
      "Can't save user ${user.id}: empty Address"
    )
  }

  // user를 데이터베이스에 저장
}
```

로컬함수를 사용하여 위 코드 중복 줄여보자.
또한 로컬함수는 자신이 속한 바깥 함수의 모든 파라미터와 변수를 사용할 수 있다.
이를 이용해서 위 코드를 고쳐보자

```kt
class User(val id: Int, val name: String, val address: String)

fun saveUser(user: User) {
  //필드를 검증하는 로컬함수를 정의한다
  fun validate(value: String, fieldName: String) {
    if (value.isEmpty()) {
      throw IllegalArgumentException(
         "Can't save user ${user.id}: empty ${fieldName}"
      )
    }
  }

  validate(user, user.name, "Name")
  validate(user, user.address, "Address")

  // user를 데이터베이스에 저장
}
```

물론 검증 로직을 User 클래스를 확장한 함수로 만들 수도 있다.

```kt
class User(val id: Int, val name: String, val address: String)

fun User.validateBeforeSave() {
  fun validate(value: String, fieldName: String) {
    if (value.isEmpty()) {
      throw IllegalArgumentException(
         "Can't save user ${user.id}: empty ${fieldName}"
      )
    }
  }

  validate(user, user.name, "Name")
  validate(user, user.address, "Address")
}

fun saveUser(user: User) {
  user.validateBeforeSave()
  // user를 데이터베이스에 저장
}
```
