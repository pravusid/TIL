# DSL

DSL (Domain Specific Language) : 영역 특화 언어

## API에서 DSL로

API를 깔끔하게 작성했다는 말은 두 가지 뜻이다.

- 코드를 읽는 사람들이 명확하게 이해할 수 있어야 한다
- 코드가 간결해야 한다.

깔끔한 코드 작성을 위해서 코틀린에서는 확장 함수, 중위 함수 호출, 람다 구문의 `it` 등 여러 문법적 편의와 연산자 오버로딩을 지원한다.

| 일반 구문 | 간결한 구문 | 사용한 언어 특성
| ----- | ----- | ----- |
| `StringUtil.capitalize(s)` | `s.capitalize()` | 확장함수 |
| `1.to("one")` | `1 to "one"` | 중위 호출 |
| `set.add(2)` | `set += 2` | 연산자 오버로딩 |
| `map.get("key")` | `map["key"]` | `get` 메소드에 대한 관례 |
| `file.use({ f -> f.read() })` | `file.use { it.read() }` | 람다의 관례
| `sb.append("yes"); sb.append("no");` | `with (sb) { append("yes"); append("no") }` | 수신객체 지정 람다 |

### 영역 특화 언어의 개념

DSL은 범용 언어를 사용하는 경우보다 특정 영역에 대한 연산을 더 간결하게 기술할 수 있다.
DLSL은 범용 프로그래밍 언어와 달리 더 선언적이다. 범용 프로그래밍 언어는 보통 명령적이다.

DSL은 자체 문법이 있기 때문에 다른 언어의 프로그램안에 포함하기 어렵다.
이런 문제를 해결하면서 DSL의 이점을 살리는 방법으로 internal DSL이라는 개념이 도입되고 있다.

### Internal DSL

내부 DSL은 범용 언어로 작성된 프로그램의 일부이며 범용 언어와 동일한 문법을 사용한다.

SQL 문법을 코틀린에서 제공하는 `Exposed` 프레임워크를 살펴보자

```kotlin
(Country join Customer)
    .slice(Country.name, Count(Customer.id))
    .selectAll()
    .groupBy(Country.name)
    .orderBy(Count(CUstomer.id), isAsc = false)
    .limit(1)
```

### DSL의 구조

다른 API에는 존재하지 않지만 DSL에만 존재하는 특징은 구조 또는 문법이다.
전형적인 라이브러리는 여러 메소드로 이뤄지면 클라이언트는 그런 메소드를 한번에 하나씩 호출해서 라이브러리를 사용한다.
함수 호출 시퀸스에는 아무런 구조가 없고, 한 호출과 다른 호출 사이에는 아무 맥락도 존재하지 않는다.
그런 API를 command-query API라고 부른다.

반대로 DSL의 메소드 호출은 DSL 문법에 의해 정해지는 구조에 속한다.

## DSL에서 수신 객체 지정 DSL 사용

수신 객체 지정 람다는 구조화된 API를 만들 때 도움이 되는 기능이다.

### 수신 객체 지정 람다와 확장 함수 타입

`buildString` 함수를 통해 코틀린이 수신객체지정 람다를 어떻게 구현하는지 살펴보자.

```kotlin
fun buildString(builderAction: (StringBuilder) -> Unit): String {
  val sb = StringBuilder()
  builderAction(sb)
  return sb.toString()
}

>>> val s = buildString {
  it.append("hello, ")
  it.append("World!")
}
>>> println(s)
Hello, World!
```

매번 메소드 앞에 `it`을 넣지 않으려면 수신객체지정 람다로 바꿔야 한다.

```kotlin
fun buildString(builderAction: StringBuilder.() -> Unit): String {
  ...
}

>>> val s = buildString {
  this.append("Hello, ")
  append("World")
}
>>> println(s)
Hello, World!
```

수신객체지정 람다를 변수에 저장

```kotlin
val appendExc1: StringBuilder.() -> Unit = { this.append("!") }

>>> val stringBuilder = StringBuilder("Hi")
>>> stringBuilder.appendExc1()
>>> println(stringBuilder)
Hi!
```

표준 라이브러리의 `buildString` 구현은 `builderAction`을 명시적으로 호출하는 대신 `apply` 함수에 인자로 넘긴다.

```kotlin
fun buildString(builderAction: StringBuilder.() -> Unit): String =
    StringBuilder().apply(builderAction).toString()
```

`apply` 함수와 `with` 함수의 구현을 살펴보자

```kotlin
inline fun <T> T.apply(block: T.() -> Unit): T {
  block()
  return this
}

inline fun <T, R> with(receiver: T, block: T.() -> R): R =
    receiver.block()
```

### 수신 객체 지정 람다를 HTML 빌더에서 사용

type-safe builder 를 코틀린으로 작성할 수 있다.

```kotlin
fun createSimpleTable() = createHTML().
    table {
      tr{
        td { +"cell" }
      }
    }
```

각 블록의 이름 결정 규칙은 각 람다의 수신객체에 의해 결정된다.
`table` 에 전달된 수신객체는 `TABLE` 이라는 타입이며 그 안에 `tr` 메소드 정의가 있다.

```kotlin
open class Tag

class TABLE : tag {
  fun tr(init: TR.() -> Unit)
}

class TR : tag {
  fun td(init: TD.() -> Unit)
}

class TD : tag
```

앞의 빌더예제에서 수신객체를 명시하여 다시 작성하면 다음과 같다

```kotlin
fun createSimpleTable() = createHTML.
    table {
      (this@table).tr {
        (this@tr).td {
          +"cell"
        }
      }
    }
```

내부 람다에서는 외부 람다에 정의된 수신 객체를 사용할 수 있다.
코틀린 1.1 부터는 `@DslMarker` 애노테이션을 사용해서 중첩된 람다에서 외부 람다의 수신객체를 접근하지 못하게 할 수 있다.

## 간단한 태그 빌더

```kotlin
open class Tag(val name: String) {
  private val children = mutableListOf<Tag>()
  protected fun <T : Tag> doInit(child: T, init: T.() -> Unit) {
    child.init()
    children.add(child)
  }

  override fun toString() = "<$name>${children.joinTostring("")}</$name>"
}

fun table(init: TABLE.() -> Unit) = TABLE().apply(init)

class TABLE : Tag("table") {
  fun tr(init: TR.() -> Unit) = doInit(TR(), init)
}

class TR : Tag("tr") {
  fun td(init: TD.() -> Unit) = doInit(TD(), init)
}

class TD : Tag("td")
```

## `invoke` 관례를 사용한 유연한 블록중첩

`invoke` 관례를 사용하면 객체를 함수처럼 호출할 수 있다.
일상적으로 사용하면 이해하기 어려운 코드를 만들 수 있지만, DSL에서는 유용할 때가 있다.

### 함수처럼 호출할 수 있는 객체

`operator` 변경자가 붙은 `invoke` 메소드 정의가 들어있는 클래스의 객체를 함수처럼 호출할 수 있다.

```kotlin
class Greeter(val greeting: String) {
  operator fun invoke(name: String) {
    println("$greeting, $name!")
  }
}

>>> val bavarianGreeter = Greeter("Servus")
>>> bavarianGreeter("Dmitry")
Servus, Dmitry!
```

`Greeter` 안에 `invoke` 메소드를 정의하면 `Greeter` 인스턴스를 함수처럼 호출할 수 있다.
`bavarianGreeter("Dmitry")` 는 내부적으로 `bavarianGreeter.invoke("Dmitry")`로 컴파일 된다.

### invoke 관례와 함수형 타입

`invoke` 관례를 봤으므로 람다뒤에 괄호를 붙이는 방식이 실제로는 `invoke` 관례를 적용한 것에 지나지 않음을 알 수 있다.
인라인 람다를 제외한 모든 함수는 함수형 인터페이스를 구현하는 클래스로 컴파일된다.

함수타입을 확장하면서 `invoke()`를 오버라이드

```kotlin
data class Issue(
  val id: String, val project: String, val type: String,
  val priority: String, val description: String
)

class ImportantIssuesPredicate(val project: String) : (Issue) -> Boolean {
  override fun invoke(issue: Issue): Boolean {
    return issue.project == project && issue.isImportant()
  }

  private fun Issue.isImportant(): Boolean {
    return type == "Bug" && (priority == "Major" || priority == "Critical")
  }
}

>>> val i1 = Issue("IDEA-154446", "IDEA", "Bug", "Major", "Save settings failed")
>>> val i2 = Issue("KT-12183", "Kotlin", "Feature", "Normal",
        "Intention: convert several calls on the same receiver to with/apply")
>>> val predicate = ImportantIssuesPredicate("IDEA")
>>> for (issue in listOf(i1, i2).filter(predicate)) {
  println(issue.id)
}
IDEA-154446
```

람다를 여러 메소드로 나누고 각 메소드에 뜻을 명확히 알 수 있는 이름을 붙인다.
람다를 함수 타입 인터페이스를 구현하는 클래스로 변환하고 그 클래스의 `invoke` 메소드를 오버라이드하면 리팩토링이 가능하다.

### DSL의 invoke 관례

```groovy
dependencies.compile("junit:junit:4.11")
dependencies {
  compile("junit:junit:4.11")
}
```

`gradle` 설정처럼 사용자가 설정해야 할 항목이 많으면 중첩된 블록 구조를 사용하고 설정할 항목이 하나 뿐이면 간단한 함수호출 구조를 사용하게 해보자.

```kotlin
class DependencyHandler {
  fun compile(coordinate: String) {
    println("Added dependency on $coordinate")
  }

  operator fun inoke(body: DependencyHandler.() -> Unit) {
    body()
  }
}
```
