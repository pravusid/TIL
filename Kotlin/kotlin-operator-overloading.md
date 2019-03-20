# 연산자 오버로딩

## 산술 연산자 오버로딩

자바에서는 원시 타입에 대해서만 산술 연산자를 사용할 수 있고, 추가로 `String`에 대해 `+` 연산자를 사용할 수 있다.
그러나 다른 클래스에서도 산술 연산자가 유용한 경우가 있다.
예를 들어 `BigInteger` 클래스를 다룬다면 `add` 메소드를 호출하기보다는 `+` 연산을 사용하는 편이 낫다.
컬렉션에 원소를 추가하는 경우에도 `+=` 연산자를 사용할 수 있으면 좋다.

```kotlin
data class Point(val x: Int, val y: Int) {
  operator fun plus(other: Point): Point {
    return Point(x + other.x, y + other.y)
  }
}

>>> val p1 = Point(10, 20)
>>> val p2 = Point(30, 40)
>>> println(p1 + p2)
Point(x=40, y=60)
```

`plus` 함수 앞에 `operator` 키워드를 붙여야한다. 연산자를 오버로딩 하는 함수 앞에는 반드시 `operator`가 있어야한다.

연산자를 멤버 함수로 만드는 대신 확장 함수로 정의할 수도 있다.

### 오버로딩 가능한 이항 산술 연산자

| 식 | 함수이름 |
| --- | --- |
| `*` | `times` |
| `/` | `div` |
| `%` | `rem` |
| `+` | `plus` |
| `-` | `minus` |

연산자 우선순위는 언제나 표준 숫자 타입에 대한 연산자 우선순위와 같다.

> 자바를 코틀린에서 호출하는 경우에는 함수 이름이 코틀린의 관례에 맞기만 하면 항상 연산자 식을 사용해 그 함수를 호출할 수 있다. 만약 자바 클래스에 원하는 연산자 기능을 제공하는 메소드가 이미 있지만 이름만 다르다면, 관례에 맞는 이름을 가진 확장함수를 작성하고 연산을 기존 자바 메소드에 위임하면 된다.

코틀린 연산자가 자동으로 교환법칙을 지원하지는 않는다. 따라서 교환법칙에 맞는 식에 대응하는 연산자 함수를 더 정의해야 한다.

연산자 함수의 반환타입이 두 피연산자 중 하나와 일치해야 하는것도 아니다.

```kotlin
operator fun Char.times(count: Int): String {
  return toString().repeat(count)
}

>>> println('a' * 3)
aaa
```

일반함수와 마찬가지로 `operator` 함수도 오버로딩할 수 있다.
이름은 같지만 파라미터 타입이 서로 다른 연산자 함수를 여럿 만들수 있다.

코틀린은 표준 숫자 타입에 대해 비트 연산자를 정의하지 않는다. 따라서 커스텀 타입에서 비트 연산자를 정의할 수도 없다.
대신 중위 연산자 표기법을 지원하는 일반 함수를 사용해 비트 연산을 수행한다. 커스텀 타입도 그와 비슷한 함수를 정의해 사용할 수 있다.

- `shl`: 왼쪽시프트 (`<<`)
- `shr`: 오른쪽시프트 (`>>`)
- `ushr`: 0으로 부호비트설정 오른쪽 시프트 (`>>>`)
- `and`: 비트곱 (`&`)
- `or`: 비트합 (`|`)
- `xor`: 비트 배타 합 (`^`)
- `inv`: 비트 반전 (`~`)

### 복합 대입 연산자 오버로딩

`plus`와 같은 연산자를 오버로딩하면 코틀린은 `+` 연산자뿐 아니라 그와 관련있는 연산자인 `+=`도 자동으로 함께 지원한다

```kotlin
>>> var point = Point(1, 2)
>>> point += Point(3, 4)
>>> println(point)
Point(x=4, y=6)
```

경우에 따라 변경가능한 컬렉션에 원소를 추가하는 내부상태 변경을 의도하고 싶을 때가 있다.

```kotlin
>>> val numbers = ArrayList<Int>()
>>> numbers += 42
>>> println(numbers[0])
42
```

반환타입이 `Unit`인 `plusAssign` 함수를 정의하면 코틀린은 `+=` 연산자에 그 함수를 사용한다.
마찬가지로 `minusAssign`, `timesAssign`등에 적용할 수 있다.

코틀린 표준 라이브러리는 변경가능한 컬렉션에 대해 `plusAssign`을 정의한다.
일반적으로 `plus`, `plusAssign`을 동시에 정의하면 오류발생 가능성이 있으므로 이는 피해야 한다.

코틀린 표준 라이브러리는 컬렉션에 대해 두 가지 접근방법을 함께 제공한다.
`+`와 `-`는 항상 새로운 컬렉션을 반환하며, `+=`과 `-=` 연산자는 항상 변경 가능한 컬렉션에 작용해 객체 상태를 변화시킨다.
또한 읽기 전용 컬렉션에서 `+=`와 `-=`는 변경을 적용한 복사본을 반환한다. (`var`로 선언한 변수가 가리키는 읽기 전용 컬렉션에서만 사용가능)

### 단항 연산자 오버로딩

단항 연산자를 오버로딩하는 절차도 이항연산자와 동일하다.

```kotlin
operator fun Point.unaryMinus(): Point {
  return Point(-x, -y)
}

>>> val p = Point(10, 20)
>>> println(-p)
Point(x=-10, y=-20)
```

오버로딩할 수 있는 단항 산술 연산자

| 식 | 함수 이름 |
| --- | --- |
| `+a` | `unaryPlus` |
| `-a` | `unaryMinus` |
| `!a` | `not` |
| `++a`, `a++` | `inc` |
| `--a`, `a--` | `dec` |

## 비교연산자 오버로딩

### 동등성 연산자: `equals`

`!=` 연산자를 사용하는 식도 `equals` 호출로 컴파일 된다.
`==`와 `!=`는 내부에서 인자가 널인지 검사하므로 다른 연산과 달리 널이 될 수 잇는 값에도 적용할 수 있다.

`identity equals` 연산자 (`===`)를 사용해 `equals`의 파라미터가 수신객체와 같은지 살펴본다.
식별자 비교 연산자는 자바 `==` 연산자와 같다. `equals`를 구현할 때는 `===`를 사용해 자기 자신과의 비교를 최적화하는 경우가 많다.

`===` 연산자는 오버로딩 할 수 없다.

`equals` 함수에는 `operator`가 아닌 `override`가 붙어있다. `Any`의 `equals`에는 `operator`가 붙어있기 때문이다.
또한 `Any`에서 상속받은 `equals`가 확장 함수보다 우선순위가 높기 때문에 `equals`를 확장함수로 정의할 수 없다.

### 순서 연산자: `compareTo`

코틀린도 자바와 같은 `Comparable` 인터페이스를 지원한다. 게다가 코틀린은 `Comparable` 인터페이스의 `compareTo` 메소드를 호출하는 관례를 제공한다.
따라서 비교연산자 (`<`, `>`, `<=`, `>=`)는 `compareTo` 호출로 컴파일된다.

`p1 < p2`는 `p1.compareTo(p2) < 0`와 같다.

`equals`와 마찬가지로 `Comparable`의 `compareTo`에도 `operator` 변경자가 붙어있으므로 하위 클래스의 오버라이딩 함수에 `operator`를 붙일필요가 없다.

`Comparable` 인터페이스를 구현하는 모든 자바클래스는 코틀린에서 비교연산자 구문을 사용할 수 있다.

```kotlin
>>> println("abc" < "bac")
true
```

## 컬렉션과 범위에 쓸 수 있는 관례

### 인덱스로 원소에 접근: `get`, `set`

코틀린에서는 맵의 원소에 접근할때 `[]` 괄호를 사용한다.
같은 연산자를 사용해 변경 가능 맵에 키/값 쌍을 넣거나, 이미 있는 값을 변경할 수 있다.

`mutableMap[key] = newValue`, 이 코드의 동작 방식을 살펴보자

`get` 관례 구현

```kotlin
operator fun Point.get(index: Int): Int {
  // index 0은 x좌표, index 1은 y좌표
  return when (index) {
    0 -> x
    1 -> y
    else -> throw IndexOutOfBoundsException()
  }
}

>>> val p = Point(10, 20)
>>> println(p[1])
20
```

`set` 관례 구현

```kotlin
data class MutablePoint(var x: Int, var y: Int)

operator fun MutablePoint.set(index: Int, value: Int) {
  when (index) {
    0 -> x = value
    1 -> y = value
    else -> throw IndexOutOfBoundsException()
  }
}

>>> val p = MutablePoint(10, 20)
>>> p[1] = 42
>>> println(p)
MutablePoint(x=10, y=42)
```

### `in` 관례

`in`은 객체가 컬렉션에 들어있는지 검사하는 membership test를 한다.
`in` 연산자와 대응하는 함수는 `contains`이다.

```kotlin
data class Rectangle(val upperLeft: Point, val lowerRight: Point)

operator fun Rectangle.contains(p: Point): Boolean {
  return p.x in upperLeft.x until lowerRight.x &&
      p.y in upperLeft.y until lowerRight.y
}

>>> val rect = Rectangle(Point(10, 20), Point(50, 30))
>>> println(Point(20, 30) in rect)
true
```

`in`의 우항에 있는 객체는 `contains` 메소드의 수신객체가 되고, `in`의 좌항에 있는 객체는 `contains` 메소드의 인자가 된다.

### `rangeTo` 관례

범위를 만드려면 `..` 구문을 사용해야 한다. 예를 들어 `1..10`은 1부터 10까지 모든 수가 들어있는 범위를 가리킨다.
`..` 연산자는 `rangeTo` 함수를 간략하게 표현하는 방법이다.

`..` 연산자는 아무클래스에 정의할 수 있지만, 클래스가 `Comparable` 인터페이스를 구현하면 `rangeTo`를 정의할 필요가 없다.
코틀린의 표준 라이브러리에는 모든 `Comparable` 객체에 대해 적용 가능한 `rangeTo` 함수가 들어있다.

예를 들어 `LocalDate` 클래스를 사용해 날짜의 범위를 만들어보자

```kotlin
>>> val now = LocalDate.now()
>>> val vacation = now..now.plusDay(10) // 오늘부터 10일짜리 범위 생성
>>> println(now.plusWeeks(1) in vacation)
true
```

`rangeTo` 연산자는 다른 산술연산자보다 우선순위가 낮지만 혼동을 피하기 위해 괄호로 인자를 감싸주면 좋다.

```kotlin
>>> val n = 0
>>> println(0..(n+1))
0..10
```

또한 범위 연산자는 우선순위가 낮아서 `0..n.forEach { }`와 같은 식을 컴파일 하려면 `(0..n).forEach { }`처럼 사용해야 한다.

### `for` 루프를 위한 `iterator` 관례

`for` 루프는 범위 검사와 똑같이 `in` 연산자를 사용하지만 이 경우 `in`의 의미는 다르다.
`for (x in list) { ... }`와 같은 문장은 `list.iterator()`를 호출해서 이터레이터에 대해 `hasNext`와 `next` 호출을 반복하는 식으로 변환된다.

코틀린에서는 이 또한 관례이므로 `iterator` 메소드를 확장 함수로 정의할 수 있다.
이런 성질로 인해 일반 자바 문자열에 대한 `for` 루프가 가능하다.
코틀린 표준 라이브러리는 `String`의 상위 클래스인 `CharSequence`에 대한 `iterator` 확장함수를 제공한다.

클래스 안에 직접 `iterator` 메소드를 구현할 수도 있다.

```kotlin
operator fun ClosedRange<LocalDate>.iterator(): Iterator<LocalDate> =
    object : Iterator<LocalDate> {
      var current = start

      override fun hasNext() = current <= endInclusive

      override fun next() = current.apply {
        current = plusDays(1)
      }
    }

>>> val newYear = LocalDate.ofYearDay(2017, 1)
>>> val daysOff = newYear.minusDays(1)..newYear
>>> for (dayOff in daysOff) { println(dayOff) }
2016-12-31
2017-01-01
```

`rangeTo` 라이브러리 함수는 `ClosedRange`의 인스턴스를 반환한다.
코드에서 `ClosedRange<LocalDate>`에 대한 확장함수 `iterator`를 정의했기 때문에 `LocalDate`의 범위객체를 `for` 루프에 사용할 수 있다.

## 구조 분해 선언과 `component` 함수

구조 분해 선언(destucturing declaration)을 사용하면 복합적인 값을 분해해서 여러 다른 변수를 한번에 초기화 할 수 있다.

```kotlin
>>> val p = Point(10, 20)
>>> val (x, y) = p
>>> println(x)
10
>>> println(y)
20
```

내부에서 구조분해 선언은 관례를 사용한다. 구조 분해 선언의 각 변수를 초기화하기 위해 `componentN` 이라는 함수를 호춣한다.
`N`은 구조 분해 선언에 있는 변수 위치에 따라 붙는 번호이다.

`data` 클래스의 주 생성자에 들어있는 프로퍼티에 대해서는 컴파일러가 자동으로 `componentN` 함수를 만들어 준다.

앞의 코드는 다음과 같이 컴파일 된다.

```kotlin
val a = p.component1()
val b = p.component2()
```

구조 분해 선언은 함수에서 여러 값을 반환할 때 유용하다.
여러 값을 한번에 반환해야 하는 함수가 있다면 반환해야 하는 모든 값이 들어갈 데이터 클래스를 정의하고 함수의 반환 타입을 그 데이터 클래스로 바꾼다.

배열이나 컬렉션에도 `componentN` 함수가 있다.

```kotlin
data class NameComponents(val name: String, val extension: String)

fun splitFilename(fullName: String): NameComponents {
  val (name, extension) = fullName.split('.', limit = 2)
  return NameComponents(name, extension)
}
```

코틀린 표준 라이브러리는 `N <= 5`에 대한 `componentN`을 제공한다.
표준 라이브러리의 `Pair`나 `Triple` 클래스를 사용하면 함수에서 여러 값을 더 간단하게 반환할 수 있다.

### 구조 분해 선언과 루프

함수 본문 내의 선언문뿐 아니라 변수 선언이 들어갈 수 있는 장소라면 어디든 구조분해 선언을 사용할 수 있다.

```kotlin
fun printEntries(map: Map<String, String>) {
  for ((key, value) in map) { // 루프변수에 구조분해 선언을 사용함
    println("${key} -> ${value}")
  }
}
```

## 위임 프로퍼티

위임 프로퍼티(delegated property)를 사용하면 값을 뒷받침하는 필드에 단순히 저장하는 것 보다 더 복잡한 방식으로 작동하는 프로퍼티를 쉽게 만들 수 있다.
이러한 특성의 기반에는 위임이 있다. 위임은 객체가 직접 작업을 수행하지 않고 위임 객체가 작업을 처리하게 하는 디자인 패턴을 말한다.

### 위임 프로퍼티 문법

```kotlin
class Delegate {
  operator fun getValue(...) { ... }
  operator fun setValue(..., value: Type) { ... }
}

class Foo {
  var p: Type by Delegate() // 프로퍼티와 위임객체 연결
}
```

이런경우 foo.p의 게터와 세터는 `Delegate` 타입의 위임 프로퍼티 객체에 있는 메소드를 호출한다.

### `by lazy()`를 사용한 프로퍼티 초기화 지연

lazy initialization은 객체의 일부분을 초기화하지 않고 남겨뒀다, 실제로 그 값이 필요한 경우 초기화 할때 쓰이는 패턴이다.
초기화 과정에 자원을 많이 사용하거나 객체를 사용할 때 마다 꼭 초기화하지 않아도 되는 프로퍼티에 대해 지연 초기화를 사용할 수 있다.

위임 프로퍼티가 아닌 지연초기화를 backing property를 사용한 코드는 다음과 같다

```kotlin
class Person(val name: String) {
  private var _emails: List<Email>? = null
  val emails: List<Email>
    get() {
      if (_emails == null) {
        _emails = loadEmails(this)
      }
      return _emails!!
    }
}
```

이런 코드는 작성이 성가시다. 위임 프로퍼티를 사용하면 코드를 더 간단하게 작성할 수 있다.
이와 같은 경우를 위해 위임객체를 반환하는 표준 라이브러리 함수인 `lazy`를 활용하면 된다.

```kotlin
class Person(val name: String) {
  val emails by lazy { loadEmails(this) }
}
```

`lazy` 함수는 코틀린 관례에 맞는 시그니처의 `getValue` 메소드가 들어있는 객체를 반환한다.
따라서 `by lazy` 키워드를 통해 위임 프로퍼티를 만들 수 있다. `lazy` 함수의 인자는 값을 초기화할 때 호출할 람다이다.
기본적으로 `lazy` 함수는 thread safe 하지만, 필요에 따라 동기화하거나 동기화 하지 못하게 설정할 수 있다.

### 위임 프로퍼티 컴파일 규칙

다음과 같은 위임 프로퍼티가 있는 클래스가 있다고 가정하자

```kotlin
class C {
  var prop: Type by MyDelegate()
}
val c = C()
```

컴파일러는 `MyDelegate` 클래스의 인스턴스를 감춰진 프로퍼티에 저장하며 그 감춰진 프로퍼티를 `<delegate>`라는 이름으로 부른다.
또한 컴파일러는 프로퍼티를 표현하기 위해 KProperty 타입의 객체를 사용한다. 이 객체를 `<property>`라고 부른다.

컴파일러는 다음과 같은 코드를 생성한다

```kotlin
class C {
  private val <delegate> = MyDelegate()
  var prop: Type
    get() = <delegate>.getValue(this, <property>)
    set(value: Type) = <delegate>.setValue(this, <property>, value)
}
```

이런 특성을 활용할 여러 방법이 있다. 프로퍼티 값이 저장될 장소를 바꿀 수도 있고(맵, 데이터베이스 테이블, 사용자 세션쿠키 등)
프로퍼티를 읽거나 쓸 때 벌어질 일을 변경할 수도 있다(값 검증, 변경 통지 등)

### 프로퍼티 값을 맵에 저장

자신의 프로퍼티를 동적으로 정의할 수 있는 객체를 만들 때 위임 프로퍼티를 활용하는 경우가 있다.
그런 객체를 확장 가능 객체(expando object)라고 부른다.

그런 상황을 구현하는 방법 중에는 정보를 모두 맵에 저장하고, 맵을 통해 처리하는 프로퍼티로 필수정보를 제공하는 방법이 있다.

```kotlin
class Person {
  private val _attributes = hashMapOf<String, String>()

  fun setAttribute(attrName: String, value: String) {
    _attributes[attrName] = value
  }

  val name: String by _attributes // 맵에게 위임한다
}
```

이와 같은 코드가 작동하는 이유는 표준 라이브러리 `Map`과 `MutableMap` 인터페이스에 대해 `getValue`와 `setValue` 확장함수를 제공하기 때문이다.
`getValue`와 `setValue`에서 자동으로 프로퍼티 이름을 키로 활용한다.

### 프레임워크에서 위임 프로퍼티 활용

객체 프로퍼티를 저장하거나 변경하는 방법을 바꿀 수 있으면 프레임워크 개발시 유용하다.

위임 프로퍼티를 사용해 데이터베이스 칼럼에 접근하는 예제이다

```kotlin
// 데이터베이스 테이블에 대응하는 객체 (싱글톤)
object Users : IdTable() {
  val name = varchar("name", length = 50).index()
  val age = integer("age")
}

class User(id: EntityID) : Entity(id) {
  var name: String by Users.name
  var age: Int by Users.age
}
```
