# Kotlin Type System

## Nullability

널 가능성은 `NullPointerException`(NPE)를 피하기 위한 코틀린 타입 시스템의 특성이다.

코틀린을 비롯한 최신 언어에서 `null`에 대한 접근 방법은 되도록이면 컴파일 시점에서 확인하는 것이다.
`nuul`이 될 수 있는지 여부를 타입 시스템에 추가함으로써 컴파일러가 오류를 미리 감지 할 수 있다.

### 널이 될 수 있는 타입

코틀린과 자바의 가장 중요한 차이는 코틀린 타입 시스템이 `nuul`이 될 수 있는 타입을 명시적으로 지원한다는 것이다.

어떤 변수가 널이 될 수 있다면 그 변수를 수신객체로 메소드를 호출 할 시 NPE가 발생할 수 있으므로 안전하지 않다.
코틀린은 그런 메소드 호출을 금지함으로써 많은 오류를 방지한다.

```java
int strLen(String s) {
  return s.length();
}
```

만약 위의 함수에서 s가 `null`이면 NPE가 발생한다.
이 함수를 코틀린으로 다시 작성하면 다음과 같다.

```kt
fun strLen(s: String) = s.length
```

코틀린에서 strLen에 `null`이거나 널이 될 수 있는 인자를 넘기는 것은 금지되어 있으며,
그런 값을 넘긴다면 컴파일 시 오류가 발생한다.

이 함수가 널과 문자열 모두를 인자로 받으려면 타입 이름 뒤에 물음표 `?`를 명시해야 한다.

```kt
fun strLenSafe(s: String?) = ...
```

널이 될 수 있는 타입의 값이 있다면 그 값에 대해 수행할 수 있는 연산의 종류는 제한적이다.
이렇게 제약이 많다면 널이 될 수 있는 타입의 값으로 무엇을 할 수 있을까.
가장 중요한 것은 결과 값을 `null`과 비교하는 것이다.

```kt
fun strLenSafe(s: String?): Int = if (s != null) s.length else 0
```

널 가능성을 `if`로 확인할 수 밖에 없다면 코드가 복잡하지겠지만, 코틀링네서는 널이 될 수 있는 값을 다룰 때 여러 도구를 제공한다.

### 안전한 호출연산자 (`?.`)

코틀린에서 실행시점에 널이 될 수 있는 타입이나 널이 될 수 없는 타입의 객체는 같다.
널이 될 수 있는 타입은 널이 될 수 없는 타입을 감싼 래퍼 타입이 아니다.
널 검사는 컴파일 시점에서 수행되므로 널이 될 수 있는 타입을 처리하는데 별도의 부가 비용이 들지 않는다.

안전한 호출 연산자 `?.`는 `null` 검사와 메소드 호출을 한 번의 연산으로 수행한다.
다시 말해서 호출하려는 값이 `null`이 아니라면 `?.`은 일반 메소드 호출처럼 작동한다.

안전한 호출의 결과 타입도 `null`이 될 수 있는 타입이라는 점을 유의해야 한다.
만약 호출하려는 값이 `null`이면 이 호출은 무시되고 `null` 최종 결과가 된다.

메소드 호출 뿐만 아니라 프로퍼티를 읽거나 쓸 때도 안전한 호출을 사용할 ㅜㅅ 있다.

```kt
class Employee(val name: String, val manager: Employee?)

fun managerName(employee: Employee): String? = employee.manager?.name

>>> println(managerName(ceo))
null
```

자바에서는 `if`문으로 `null` 확인을 해야 NPE가 발생하지 않는 method chaining 의 경우를 보자.
코틀린에서는 널이 될 수 있는 중간 객체가 여럿 있다면 식 하나에서 안전한 호출을 연쇄사용하면 편리하다.

```kt
class Address(val streetAddress: String, val zipCode: Int,
    val city: String, val country: String)

class Company(val name: String, val address: Address?)

class Person(val name: String, val company: Company?)

fun Person.countryName(): String {
  val country = this.company?.address?.country
  return if (country != null) country else "Unknown"
}
```

### elvis 연산자 (`?:`)

코틀린은 `null` 대신 사용할 디폴트 값을 지정할 때 사용할 수 있는 연산자를 제공하는데, 엘비스 연산자라고 한다.

엘비스 연산자는 이항연산자로 좌항이 널인지 검사하고, 좌항값이 널이 아니면 좌항을, 널이면 우항값을 결과로 한다.

```kt
val t: String = s ?: ""
```

이를 이용하면 앞의 예제를 고쳐쓸 수 있다.

```kt
class Address(val streetAddress: String, val zipCode: Int,
    val city: String, val country: String)

class Company(val name: String, val address: Address?)

class Person(val name: String, val company: Company?)

fun Person.countryName(): String {
  val country = this.company?.address?.country
  return country ?: throw IllegalArgumentException("No address")
}
```

### 안전한 cast ('as?`)

앞에서 코틀린 타입 캐스트 연산자인 `as`를 살펴보았다.
자바와 마찬가지로 대상을 지정한 타입으로 바꿀 수 없으면 `ClassCastException`이 발생한다.
물론 `is`를 사용해서 변환가능한 타입인지 검사할 수도 있지만, 보다 간결한 해법이 있다.

`as?` 연산자는 값을 지정한 타입으로 캐스트한다. 대상 타입으로 변환할 수 없으면 `null`을 반환한다.
따라서 안전한 캐스트를 사용할 때는 엘비스 연산자와 함께 사용하는 것이 좋다.

```kt
class Person(val firstName: String, val lastName: String) {
  override fun equals(o: Any?): Boolean {
    val otherPerson = o as? Person ?: return false
    // 안전한 cast를 사용했으므로 otherPerson이 Person으로 스마트 캐스트 됨
    return otherPerson.firstName == firstName && otherPerson.lastName == lastName
  }

  override fun hashCode(): Int = firstName.hashCode() * 37 + lastName.hashCode()
}

>>> val p1 = Person("Dmitry", "Jemerov")
>>> val p2 = Person("Dmitry", "Jemerov")
>>> println(p1 === p2) // == 연산자는 equals 메소드를 호출함
```

### `null`이 아님 (`!!`)

not-null assertion은 어떤 값이든 강제로 널이 될 수 없는 타입으로 바꿀 수 있다.
`null`에 대해 `!!`를 사용하면 NPE가 발생한다.

```kt
fun ignoreNulls(s: String?) {
  val sNotNull: String = s!! // 예외 발생지점
  println(sNotNull.length)
}

>>> ignoreNulls(null)
kotlin.KotlinNullPointerException
```

`!!`은 컴파일러가 검증할 수 없는 단언이기 때문에 일반적으로 더 나은 방법을 찾는 것이 좋다.

하지만 not-null-assertion이 더 나은 해법인 경우도 있다.
어떤 함수가 값이 널인지 검사한 다음 다른함수를 호출한다고 해도 컴파일러는 호출된 함수 안에서 안전하게 그 값을 사용할 수 있음을 인식할 수 없다.
이런 경우 호출된 함수가 언제나 다른 함수에서 널이 아닌 값을 전달받는다는 사실이 분명하다면 굳이 널 검사를 다시 수행할 필요가 없다.
이런경우 `!!` 단언문을 쓸 수 있다.

```kt
class CopyRowAction(val list: JList<String>) : AbstractAction() {
  override fun isEnabled(): Boolean = list.selectedValue != null
  override fun actionPerformed(e: ActionEvent) { // actionPerformed는 isEnabled가 true인 경우에만 호출되므로
    val value = list.selectedValue!!
  }
}
```

위의 경우 `!!`를 사용하지 않으려면 `val value = list.selectedValue ?: return` 처럼 널이 될 수 없는 타입의 값을 얻어야 한다.

`!!`를 널에 대해 사용해서 발생하는 예외의 stack trace에는 어떤 위치에서 예외가 발생했는지 정보가 들어있지만,
어떤 식에서 예외가 발생했는지에 대한 정보가 들어있지 않으므로, `!!` 단언문을 한 줄에 함께 쓰는 것은 피해야 한다.

`person.company!!.address!!.country` 와 같이 작성하지 말아야 한다.

### let함수

`let` 함수를 사용하면 널이 될 수 있는 식을 더 쉽게 다룰 수 있다.
`let` 함수를 안전한 호출 연산자와 함께 사용하면 원하는 식의 결과가 널인지 검사한 다음 결과를 변수에 넣는 작업을 한번에 처리할 수 있다.

```kt
fun sendEmailTo(email: String) {
  println("Sending email to ${email})
}

getTheBestPersonInTheWorld()?.let { sendEmailTo(it.email) }
```

### 나중에 초기화할 프로퍼티

객체 인스턴스를 일단 생성한 다음에 나중에 초기화하는 프레임워크가 많다.
하지만 코틀린에서 클래스 안의 널이 될 수 없는 프로퍼티를 생성자 안에서 초기화 하지 않고 특별한 메소드 안에서 초기화할 수는 없다.
생성자에서 그런 초기화 값을 제공할 수 없으면 널이 될 수 잇는 타입을 사용할 수 밖에 없다.

이를 해결하기 위해 프로퍼티를 late-initialized 할 수 있다.
`lateinit` 변경자를 붙이면 프로퍼티를 나중에 초기화 할 수 있다.

```kt
class MyService {
  fun performAction(): String = "foo"
}

class MyTest {
  private lateinit var myService: MyService // myService 프로퍼티는 아직 초기화 되지 않았다.

  @Before fun setUp() {
    myService = MyService() // 여기에서 프로퍼티 초기화가 이루어진다
  }

  @Test fun testAction() {
    Assert.assertEquals("foo", myService.performAction())
  }
}
```

나중에 초기화하는 프로퍼티는 항상 `var`여야 한다. val 프로퍼티는 `final`필드로 컴파일되므로 생성자안에서 반드시 초기화 해야한다.

나중에 초기화하는 프로퍼티를 초기화 하기 이전에 프로퍼티에 접근하면 "lateinit property ooo has not been initialized" 예외가 발생한다.

### 널이 될 수 있는 타입 확장

어떤 메소드를 호출하기 전에 수신 객체 역할을 하는 변수가 널이 될 수 없다고 보장하는 대신,
직접 변수에 대해 메소드를 호출해도 확장 함수인 메소드가 알아서 널을 처리해준다. 이런 처리는 확장함수에서만 가능하다.

```kt
fun verifyUserInput(input: String?) {
  if (input.isNullOrBlank()) { // 안전한 호출을 하지 않아도 된다
    println("please fill in the required fields")
  }
}
```

일반 멤버 호출은 객체 인스턴스를 통해 dispatch되므로 그 인스턴스가 널인지 여부를 검사하지 않는다.
즉, 널이 될 수 있는 타입의 확장 함수는 안전한 호출 없이도 호출 가능하다.
그렇기 때문에 해당 확장함수 내부에서는 `null` 여부를 검사해야 한다.

```kt
fun String?.isNullOrBlank(): Boolean = this == null || this.isBlank() // 두번째 this는 스마트 캐스트가 적용된다.
```

### 타입 파라미터의 널 가능성

코틀린에서는 함수나 클래스의 모든 타입 파라미터는 기본적으로 널이 될 수 있다.
널이 될 수 있는 타입을 포함하는 어떤 타입이라도 타입 파라미터를 대신할 수 있다.
따라서 타입 파라미터 `T`를 클래스나 함수 안에서 타입 이름으로 사용하면 이름 끝에 `?`가 없어도 `T`는 널이 될 수 있는 타입이다.

```kt
fun <T> printHashCode(t: T) {
  println(t?.hashCode()) // t는 null이 될 수 있으므로 안전한 호출을 써야 한다
}
```

타입 파라미터가 널이 아님을 확실히 하려면 널이 될 수 없는 upper bound를 지정해야 한다.

```kt
fun <T: Any> printHashCode(t: T) { // T는 null이 될 수 없는 타입이다
  println(t.hashCode())
}
```

### 널 가능성과 자바

자바코드에 애노테이션으로 표시된 널 가능성 정보가 있으면 코틀린도 그 정보를 활용한다.
자바의 `@Nullable String`은 코틀린의 `String?`과 같고 자바의 `@NotNull String`은 코틀린의 `String`과 같다.

#### 플랫폼 타입

플랫폼 타입은 자바코드에서 코틀린이 널 관련 정보를 알 수 없는 타입을 말한다.
그 타입은 널이 될 수 있는 타입으로 처리해도 되고 널이 될 수 없는 타입으로 처리해도 된다.
즉, 사용자에게 온전히 널 가능성 확인 여부를 위임하는 것이다.

코틀린에서 자바타입을 플랫폼 타입이 아닌 널이 될 수 있는 타입으로 다룰수도 있었겠지만, 불필요한 널 검사를 없애기 위해서 플랫폼 타입을 도입하였다.

코틀린에서는 플랫폼 타입을 선언할 수 없지만, 오류 메시지에서는 플랫폼 타입을 확인할 수 있다.

`ERROR: Type mismatch: inferred type is String! but Int was expected`

여기에서 `String!` 이라는 느낌표가 붙은 타입은, 해당 타입의 널 가능성에 대한 아무런 정보가 없다는 뜻이다.

### 상속

코틀린에서 자바 메소드를 오버라이드 할 때 그 메소드의 파라미터와 반환 타입을 널이 될 수 있는 타입으로 선언할지,
널이 될 수 없는 타입으로 선언할지 결정해야 한다.

```java
interface StringProcessor {
  void process(String value);
}
```

코틀린 컴파일러는 다음의 두 가지를 모두 받아들인다

```kt
class StringPrinter : StringProcessor {
  override fun process(value: String) {
    println(value)
  }
}

class NullableStringPrinter : StringProcessor {
  override fun process(value: String?) {
    if (value != null) {
      println(value)
    }
  }
}
```

자바 클래스나 인터페이스를 코틀린에서 구현할 경우 널 가능성을 제대로 처리하는 일이 중요하다.
구현 메소드를 다른 코틀린 코드가 호출할 수 있으므로
코틀린 컴파일러는 널이 될 수 없는 타입으로 선언한 모든 파라미터에 대해 널이 아님을 검사하는 단언문을 만들어주고
자바 코드가 그 메소드에 널 값을 넘기면 단언문이 발동되어 예외가 발생한다.

## 코틀린의 원시 타입

### 원시타입 Int, Boolean 등

자바는 primary type과 reference type을 구분한다.
원시타입의 변수에는 값이 직접 들어가지만 (call by value), 참조타입의 변수에는 메모리상 객체 위치가 들어간다 (call by reference).

자바는 원시타입에 대한 참조가 필요한 경우 래퍼 타입으로 감싸서 사용한다.
코틀린은 원시 타입과 래퍼타입을 구분하지 않으므로 항상 같은 타입을 사용한다.

코틀린에서 원시타입과 참조타입을 같게 표현하지만 항상 객체로 표현하는 것은 아니다.
항상 객체로 표현 하면 비효율적이기 때문이다.

실행 시점에 `Int` 타입은 가능한 가장 효율적인 방식으로 표현된다. 대부분의 경우 코틀린 `Int`타입은 자바 `int` 타입으로 컴파일 된다.
이런 컴파일이 불가능 한 경우에는 컬렉션과 같은 제네릭 클래스를 사용하는 경우 뿐이다.

다음은 자바 원시 타입에 해당하는 코틀린 타입이다

- 정수타입: Byte, Short, Int, Long
- 부동소수점 수 타입: Float, Double
- 문자타입: Char
- 불리언타입: Boolean

`Int`와 같은 코틀린 타입에는 널 참조가 들어갈 수 없기 때문에 쉽게 그에 상응하는 자바 원시타입으로 컴파일 할 수 있다.
마찬가지로 자바 원시타입 값은 널이 될 수 없으므로 해당 값을 코틀린에서 사용할 때도 플랫폼 타입이 아니라 널이 될 수 없는 타입으로 취급할 수 있다.

### 널이 될 수 있는 원시 타입: Int?, Boolean? 등

`null` 참조를 자바의 참조 타입의 변수에만 대입할 수 있기 때문에 널이 될 수 있는 코틀린 타입은 자바 원시타입으로 표현할 수 없다.
따라서 코틀린에서 널이 될 수 있는 원시타입을 사용하면 그 타입은 자바의 래퍼타입으로 컴파일 된다.

원시타입이라도 널이 될 수 있는 타입으로 선언하였다면 널이 아닌지 검사를 해야 일반적인 값으로 다룰 수 있다.

### 숫자 변환
