# 클래스, 객체, 인터페이스

코틀린의 클래스와 인터페이스는 자바와 약간 다르다.
인터페이스에 프로퍼티 선언이 들어갈 수 있으며, 자바와 달리 기본적으로 final이며 public이다.
또한 중첩 클래스는 기본적으로 내부클래스가 아니다. 즉 코틀린 중첩 클래스는 외부 클래스에 대한 참조가 없다.

## 클래스 계층 정의

### 인터페이스

코틀린 인터페이스는 자바8 인터페이스와 비슷하다. 즉 내부에 추상 메소드뿐 아니라 구현이 있는 메소드를 정의할 수 있다.
다만 인터페이스에는 아무런 상태(필드)도 들어갈 수 없다.

인터페이스 구현을 위해서 자바에서는 implements 키워드를 사용하지만,
코틀린에서는 클래스 이름 뒤에 콜론(`:`)을 붙이는 것으로 클래스 확장(extends)과 인터페이스 구현(implements)을 모두 처리한다.
자바와 마찬가지로 인터페이스는 제한 없이 구현할 수 있지만, 클래스는 하나만 확장할 수 있다.

자바의 `@Override` 애노테이션과 비슷한 `override` 변경자는 상위 클래스나 상위 인터페이스에 있는 프로퍼티나 메소드를 오버라이드 한다는 표시다.
자바와 다르게 코틀린에서는 `override`변경자를 반드시 사용해야 한다.

인터페이스 메소드도 디폴트 구현을 제공한다. 그런 경우 메소드 앞에 default를 붙여야 하는 자바8과 달리 아무런 키워드를 요구하지 않는다.

동일한 이름의 메소드를 가진 두 인터페이스를 한 클래스에서 동시에 구현하는 경우 해당 클래스에서 override 하지 않는 한 컴파일 되지 않는다.

### 인터페이스에 선언된 프로퍼티

코틀린에서는 인터페이스에 추상 프로퍼티 선언을 넣을 수 있다.

```kt
interface User {
  val nickname: String
}
```

이는 User 인터페이스를 구현하는 클래스가 nickname의 값을 얻을 수 있는 방법을 제공해야 한다는 뜻이다.
해당 인터페이스를 구현해 보자

```kt
class PrivateUser(override val nickname: String): User

class SubscribingUser(val email: String): User {
  override val nickname: String
    get() = email.substringBefore('@')
}

class FacebookUser(val accountId: Int): User {
  override val nickname = getFacebookName(accountId) // 다른곳에 정의되어 있는 함수
}

>>> println(PrivateUser("test@kotlinlang.org").nickname)
test@kotlinlang.org
>>> println(SubscribingUser("test@kotlinlang.org".nickname))
test
```

SubscribingUser의 nickname은 매번 호출될 때마다 계산하는 커스텀 게터를 활용하고,
FacebookUser의 nickname은 객체 초기화 시 계산한 데이터를 뒷받침하는 필더에 저장했다가 불러오는 방식을 활용한다.

```kt
interface User {
  val email: String
  val nickname: String
    get() = email.substringBefore('@')
}
```

위와 같은 경우 하위클래스는 추상 프로퍼티인 email은 반드시 오버라이드 해아하지만 nickname은 오버라이드 하지 않아도 된다.

인터페이스에는 추상 프로퍼티 뿐만 아니라 게터와 세터가 있는 프로퍼티를 선언할 수 있다.
그러나 인터페이스는 상태를 저장할 수 없기 때문에 게터와 세터를 뒷받침 하는 필드를 참조할 수 없다.

### open, final, abstract 변경자

fragile base class는 어떤 클래스가 자신을 상속하는 방법에 대해 정확한 규칙(어떤 메소드를 어떻게 오버라이드 해야 하는지...)을 제공하지 않는다면,
그 클래스의 클라이언트는 기반 클래스를 작성한 사람의 의도와 다른 방식으로 메소드를 오버라이드 할 위험이 있다.
따라서 기반 클래스를 변경하는 경우 하위 클래스스의 동작이 예기치 않게 바뀔 수도 있다는 면에서 기반 클래스는 취약하다.

이 문제를 해결하기 위해 '이펙티브 자바'에서는 상속을 위한 설계와 문서를 갖추거나, 그럴 수 없다면 상속을 금지하라는 조언을 한다.
코틀린도 마찬가지로 모든 클래스와 메소드는 기본적으로 `final`이다.

어떤 클래스의 상속을 허용하려면 클래스 앞에 `open` 변경자를 붙여야 하며, 오버라이드를 허용하고 싶은 메소드나 프로퍼티의 앞에도 `open` 변경자를 붙여야 한다.

기반 클래스나 인터페이스의 멤버를 오버라이드 하는 경우 그 메소드는 기본적으로 열려있다.
(열려있어야 현 클래스에서 오버라이드가 가능한 것이고, 열려있는 것을 오버라이드 했으므로 여전히 열려있다.)
오버라이드 하는 메소드를 하위 클래스에서 오버라이드 하지 못하게 하려면 `final`을 명시해야 한다.

> 클래스의 기본 상속 기능 상태를 final로 하면서 스마트 캐스트가 가능하다. 스마트 캐스트는 타입 검사 뒤에 변경될 수 없는 변수에만 적용가능하기 때문이다.

코틀린에서도 클래스를 abstract로 선언할 수 있다. abstract로 선언한 클래스는 인스턴스화 할 수 없다.
추상 멤버는 항상 열려 있으므로 앞에 `open`변경자를 명시할 필요가 없다.

인터페이스의 멤버는 항상열려있으며 final, open, abstract를 사용하지 않는다.

### 접근 변경자

기본적으로 코틀린의 접근 변경자는 자바와 비슷하다. 하지만 자바와 다르게 아무런 변경자가 없는 경우 공개(`public`)상태이다.
자바의 기본 접근 변경자인 `package-private`은 코틀린에 없다. 코틀린의 패키지는 가시성 제어에 사용되지 않고 코드의 네임스페이스 관리에만 사용된다.

패키지 전용 접근 변경자로 `internal`이 도입되었다. internal은 모듈 내부에서만 접근 가능하다. 모듈 단위의 접근으로 캡슐화를 지킬 수 있다.

자바와의 또 다른 차이는 최상위 선언(클래스, 함수, 프로퍼티)에 대해 `private`을 적용한다는 점이다.
비공개 최상위 선언은 그 선언이 들어있는 파일 내부에서만 사용할 수 있다.

자바에서 `protected` 접근 변경자는 같은 패키지안의 멤버에 접근할 수 있었지만, 코틀린에서는 오로지 해당 클래스와 그 클래스를 상속한 클래스 내에서만 보인다.
따라서 확장 함수내에서는 `private`이나 `protected` 멤버에 접근할 수 없다.

### 내부 클래스와 중첩된 클래스

코틀린의 중첩 클래스(nested class)는 명시적으로 요청하지 않는 한 바깥쪽 클래스 인스턴스에 대한 접근권한이 없다.
코틀린 충첩클래스에 아무런 변경자가 붙지 않으면 자바 static 중첩 클래스와 같다.
이를 내부 클래스로 변경해서 바깥쪽 클래스에 대한 참조를 포함하게 만들고 싶다면 inner 변경자를 붙여야 한다.

| 클래스 B 안에 정의된 클래스 A | 자바 | 코틀린 |
| ----- | ----- | ----- |
| 중첩 클래스(바깥쪽 클래스 참조 X) | static class A | class A |
| 내부 클래스(바깥쪽 클래스 참조 O) | class A | inner class A |

코틀린에서는 바깥쪽 클래스의 인스턴스를 가리키는 참조를 표기하는 방법도 다르다.
내부클래스 Inner 에서 바깥쪽 클래스 Outer를 참조하려면 `this@Outer` 라고 써야 한다.

```kt
class Outer {
  inner class Inner {
    fun getOuterReference(): Outer = this@Outer
  }
}
```

### 봉인된 클래스 (`sealed`)

상위 클래스에 `sealed` 변경자를 붙이면 그 상위 클래스를 상속한 하위클래스 정의를 제한할 수 있다.
sealed 클래스의 하위 클래스를 정의할 때는 반드시 상위 클래스 안에 중첩시켜야 한다.
sealed class는 기본적으로 `open`이므로 접근 변경자 `open`을 붙일 필요가 없다.

```kt
sealed class Expr {
  class Num(val value: Int): Expr()
  class Sum(val left: Expr, val right: Expr): Expr()
}

fun eval(e: Expr): Int =
  when (e) {
    is Expr.Num -> e.value
    is Expr.Sum -> eval(e.right) + eval(e.left)
  }
```

when 식이 모든 하위클래스를 검사하므로 별도의 else 분기가 필요 없다
즉 Expr을 상속하는 경우의 수를 모두 파악한 상태로 컴파일하므로 else 분기가 필요 없는 것

## 클래스 선언

코틀린은 primary constructor와 secondary constructor를 구분한다.
또한 코틀린에서는 initializer block을 통해 초기화 로직을 추가할 수 있다.

### 주 생성자와 초기화 블록

클래스 이름뒤에 오는 괄호로 둘러싸인 코드를 주 생성자라고 한다.
주 생성자는 생성자 파라미터를 지정하고, 그 생성자 파라미터에 의해 초기화되는 프로퍼티를 정의한다.

주 생성자를 명시적 선언으로 풀어서 살펴보자

```kt
class User constructor(_nickname: String) {
  val nickname: String
  init {
    nickname = _nickname
  }
}
```

`constructor` 키워드는 주 생성자나 부 생성자 정의를 시작할 때 사용한다.
주 생성자 앞에 별다른 애노테이션이나 가시성 변경자가 없다면 `constructor`를 생략해도 된다.

`init` 키워드는 초기화 블록을 시작한다.
초기화 블록에는 클래스의 인스턴스가 생성될 때 실행될 초기화 코드가 들어간다.
초기화 블록은 주 생성자와 함께 사용된다. 주 생성자는 제한적이어서 별도의 코드를 포함할 수 없으므로 초기화 블록이 필요하다.
필요하다면 클래스 안에 여러 초기화 블록을 선언할 수 있다.

함수 파라미터와 마찬가지로 생성자 파라미터도 디폴트 값을 정의할 수 있다.
클래스 인스턴스를 만드려면 new 키워드 없이 생성자를 직접 호출하면 된다.
모든 생성자 파라미터에 디폴트 값을 지정하면 컴파일러가 자동으로 파라미터가 없는 생성자를 만들어준다.

기반 클래스를 초기화 하려면 기반 클래스 이름 뒤에 괄호 내부로 생성자 인자를 넘겨면 된다.

```kt
open class User(val nickname: String) { ... }
class TwitterUser(nickname: String): User(nickname) { ... }
```

클래스를 정의할 때 별도로 생성자를 정의하지 않으면 컴파일러가 자동으로 인자가 없는 디폴트 생성자를 만들어 준다.
생성자 인자가 없는 클래스를 상속하는 하위 클래스는 반드시 Button 클래스의 기본 생성자를 호출해야 한다.
이 규칙으로 인해 기반 클래스의 이름 뒤에는 꼭 괄호가 들어간다.
반면 인터페이스는 생성자가 ㅇ벗기 때문에 어떤 클래스가 인터페이스를 구현하는 경우 인터페이스 이름뒤에는 괄호가 없다.

```kt
open class Button
class RadioButton: Button()
```

어떤 클래스를 클래스 외부에서 인스턴스화하지 못하게 막고 싶다면 주 생성자에 private 변경자를 붙일 수 있다.

`class Secretive private constructor() { ... }`

비공개 생성자는 companion object와 함께 사용하여 활용할 수 있다.

### 부 생성자

코틀린에서는 디폴트 파라미터와 이름을 붙인 인자를 사용해서 생성자를 줄일 수 있다.
그래도 생성자가 여럿 필요한 경우가 있을 것이다. 일반적으로 프레임워크 클래스 확장을 위해 다양한 생성자를 지원해야 하는 경우이다.

```kt
open class View {
  constructor(ctx: Context) {
    ...
  }

  constructor(ctx: Context, attr: AttributeSet) {
    ...
  }
}
```

위의 클래스는 주 생성자를 선언하지 않고(클래스 헤더 클래스 이름 뒤 괄호가 없음) 부 생성자만 2개 선언하였다.

마찬가지로 클래스를 확장하면서 부 생성자를 정의 할 수 있다. 이 때 `super()` 키워드를 통해 상위클래스 생성자를 호출한다.

```kt
class MyButton: View {
  constructor(ctx: Context): super(ctx) {
    ...
  }

  constructor(ctx: Context, attr: AttributeSet): super(ctx, attr) {
    ...
  }
}
```

자바와 마찬가지로 `this()`를 통해 클래스 자신의 다른 생성자를 호출 할 수 있다.

```kt
class MyButton: View {
  constructor(ctx: Context): this(ctx, MY_STYLE) { // 아래의 생성자에게 위임한다
    ...
  }

  constructor(ctx: Context, attr: AttributeSet): super(ctx, attr) {
    ...
  }
}
```

만약 클래스에 주 생성자가 없다면 모든 부 생성자는 반드시 상위 클래스를 초기화 하거나 다른 생성자에게 생성을 위임해야 한다.

### 게터와 세터에서 뒷받침하는 필드에 접근

```kt
class User(val name: String) {
  var address: String = "unspecified"
    set(value: String) {
      field = value
    }
}
```

접근자의 본문에서는 field라는 특별한 식별자를 통해 뒷받침하는 필드에 접근할 수 있다.

접근자의 가시성을 바꿀 필요가 있을때 어떻게 하는지 살펴보자

```kt
class LengthCounter {
  var counter: Int = 0
    private set
  fun addWord(word: String) {
    counter += word.length
  }
}
```

접근자의 가시성은 기본적으로 프로퍼티의 가시성과 같다.
하지만 원한다면 get이나 set앞에 접근 변경자를 추가해서 접근을 제어할 수 있다.

### 동등성 연산

자바에서는 `==`를 원시타입과 참조타입 비교시 사용한다.
원시타입에서는 equality, 참조타입의 경우 reference comparision으로 작동한다.
따라서 두 객체의 동등성 확인을 위해서는 `equals`를 호출해야 한다.

코틀린에서는 `==`연산자가 두 객체를 비교하는 방법이다. `==`는 내부적으로 `equals`를 호출해서 객체를 비교한다.
따라서 클래스가 `equals`를 호버라이드 하면 `==`를 통해 안전하게 클래스의 인스턴스를 비교할 수 있다.

참조 비교를 위해서는 `===` 연산자를 사용할 수 있다.

### 데이터 클래스

코틀린 컴파일러는 기계적으로 생성하는 메소드를 보이지 않는곳에서 처리해준다.

코틀린 컴파일러는 모든클래스가 정의해야 하는 메소드 `toString`, `equals`, `hashCode`등을 자동으로 생성해준다.

`data`라는 변경자를 클래스 앞에 붙이면 필요한 메소드를 컴파일러가 자동으로 만들어준다. 해당 클래스를 데이터 클래스라고 부른다.
데이터 클래스의 프로퍼티가 반드시 `val`일 필요는 없지만 불변객체를 생성하기를 권장한다.

```kt
data class Client(val name: String, val postalCode: Int)
```

Client 데이터 클래스는 아래의 내용을 포함한다.

- `equals()`
- `hashCode()`
- `toString()`
- `copy()`: 불변 객체의 데이터 변경을 위해서 인스턴스를 복사본을 만들어서 관리한다

### 클래스 위임: by 키워드 사용

대규모 객체지향 시스템을 설계할 때 시스템을 취약하게 만드는 문제는 보통 구현 상속(implementation inheritance)에 의해 발생한다.
하위 클래스가 상위 클래스의 메소드 중 일부를 오버라이드 하면 하위 클래스는 상위 클래스의 세부 구현사항에 의존하게 된다.
시스템이 변함에 따라 상위 클래스가 변화하면 하위클래스에서 상위클래스에 대해 갖고있던 가정이 깨져 클래스가 정상적으로 작동하지 못할 수 잇다.

따라서 코틀인에서는 기본적으로 클래스를 `final`로 취급하여 클래스 변경 시 `open` 변경자를 보고 하위 클래스를 깨지 않기 위해 조심할 수 있다.
하지만 종종 상속을 허용하지 않는 클래스에 새로운 동작을 추가해야 할 때가 있다.

이럴 때 사용하는 방법이 데코레이터(Decorator) 패턴이다.
상속을 허용하지 않는 기존 클래스 대신 사용할 수 있는 새로운 클래스(데코레이터)를 만들되
기존 클래스와 같은 인터페이스를 데코레이터가 제공하게 하고, 기존 클래스를 데코레이터 내부에 필드로 유지하는 것이다.

데코레이터 패턴의 문제점은 준비 코드가 상당히 많이 필요하다는 점이다.

```kt
class DelegatingCollection<T>: Collection<T> {
  private val innerList = arrayListOf<T>()

  override val size: Int get() = innerList.size
  override fun isEmpty(): Boolean = innerList.isEmpty()
  override fun contains(element: T): Boolean = innerList.contains(element)
  override fun iterator(): Iterator<T> = innerList.iterator()
  override fun containsAll(elements: Collection<T>): Boolean = innerList.containsAll(elements)
}
```

코틀린은 이런 위임을 언어가 제공하는 일급 시민 기능으로 지원한다.
인터페이스를 구현할 때 `by` 키워드를 통해 그 인터페이스에 대한 구현을 다른 객체에 위임 중이라는 사실을 명시할 수 있다.

위의 예제를 위임을 사용해 재작성하면 다음과 같다.

```kt
class DelegatingCollection<T>(
  innerList: Collection<T> = ArrayList<T>()
): Collection<T> by innerList {}
```

코틀린은 `by` 키워드로 클래스안에 있던 전달 메소드를 자동 생성한다.
메소드 중 일부의 동작을 변경하고 싶은 경우 메소드를 오버라이드하면 컴파일러가 생성한 메소드 대신 오버라이드한 메소드가 쓰인다.

클래스 위임을 사용한 예제를 살펴보자

```kt
class CountingSet<T>(
  val innerSet: MutableCollection<T> = HashSet<T>
): MutableCollection<T> by innerSet { // MutableCollection의 구현을 innerSet에게 위임함
  var objectAdded = 0

  override fun add(element: T): Boolean {
    objectsAdded++
    return innerSet.add(element)
  }

  override fun addAll(c: Collection<T>): Boolean {
    objectAdded += c.size
    return innerSet.addAll(c)
  }
}
```

`add`와 `addAll`을 오버라이드 해서 카운터를 증가시키고, MutableCollection 인터페이스의 나머지 메소드는 내부 컨테이너(innerSet)에게 위임한다.

### object 키워드

코틀린에서는 `object` 키워드는 클래스를 정의하면서 동시에 인스턴스를 생성한다.

#### 쉽게 싱글톤을 생성하자

object declaration은 singleton을 정의하는 방법 중 하나이다.
자바에서는 보통 클래스의 생성자를 `private`으로 제한하고 정적인 필드에 그 클래스의 유일한 객체를 저장하는 방식을 사용한다.

코틀린은 객체 선언 기능을 통해 singleton을 언어에서 기본 지원한다.
객체 선언은 클래스 선언과 그 클래스에 속한 단일 인스턴스의 선언을 합친 선언이다.

```kt
object Payroll {
  val allEmployees = arrayListOf<Person>()

  fun calculateSalary() {
    for (person in allEmployees) {
      ...
    }
  }
}
```

객체선언시 `object` 키워드를 사용하면 된다.
클래스와 마찬가지로 객체 선언 안에도 프로퍼티, 메소드, 초기화 블록등이 들어갈 수 있다.
하지만 생성자는 객체선언에 쓸 수 없다. 싱글턴 객체는 객체 선언문이 있는 위치에서 생성자 호출 없이 즉시 만들어진다.

객체선언도 클래스나 인스턴스를 상속할 수 있다. 프레임워크를 사용하기 위해 특정 인터페이스를 구현해야 하는데,
구현 내부에 다른 상태가 필요하지 않은 경우에 이런 기능을 사용할 수 있다.

`java.util.Comparator` 인터페이스는 두 객체를 인자로 받아 어느객체가 더 큰지 알려주는 정수를 반환한다.
`Comparator` 안에는 데이터를 저장할 필요가 없다. 따라서 `Comparator` 인스턴스를 만드는 방법으로 객체선언이 가장 좋은 방법이다.

```kt
object CaseInsensitiveFileComparator: Comparator<File> {
  override fun compare(file1: File, file2: File): Int {
    return file1.path.compareTo(file2.path, ignoreCase = true)
  }
}

>>> println(CaseInsensitiveFileComparator.compare(File("/User"), File("/user")))
0
```

> 의존관계가 별로 많지 않은 소규모 소프트웨어에서는 싱글턴이나 객체 선언이 유용하지만 시스템을 구현하는 다양한 구성요소와 상호작용하는 대규모 컴포넌트에는 싱글턴이 적합하지 않다. 객체 생성을 제어할 방법이 없고 생성자 파라미터를 지정할 수 없기 때문이다. 생성자를 제어할 수 없고 생성자 파라미터를 지정할 수 없으므로 단위 테스트를 하거나 소프트웨어 시스템 설정이 달라질 때 객체를 대체하거나 객체의 의존관계를 바꿀 수 없다.

클래스 안에서 객체를 선언할 수도 있다. 그런 객체도 인스턴스는 하나 뿐이다.

```kt
data class Person(val name: String) {
  object NameComparator: Comparator<Person> {
    override fun compare(p1: Person, p2 Person): Int = p1.name.compareTo(p2.name)
  }
}

>>> val persons = listOf(Person("Bob"), Person("Alice"))
>>> println(persons.sortedWith(Person.NameComparator))
[Person(name=Alice), Person(name=Bob)]
```

> 코틀린 객체 선언은 유일한 인스턴스에 대한 정적인 필드가 있는 자바 클래스로 컴파일 된다. 이때 인스턴스 필드의 이름은 항상 INSTANCE이다.

#### Companion Object

코틀린 클래스안에는 정적인 멤버가 없다. 즉 코틀린은 자바 `static` 키워드를 지원하지 않는다.
대신 패키지 수준의 최상위 함수와 객체 선언을 지원하고, 대부분의 경우 최상위 함수를 더 권장한다.

하지만 클래스 내부 정보에 접근해야 하는 함수가 필요할 때는 클래스에 중첩된 객체선언의 멤버함수로 정의해야한다.
이럴 때 companion object는 인스턴스 메소드는 아니지만 어떤 클래스와 관련 있는 메소드와 팩토리 메소드를 담을 때 쓰인다.
동반 객체 메소드에 접근할 때는 동반 객체가 포함된 클래스의 이름을 사용할 수 있다.

결과적으로 동반 객체의 멤버를 사용하는 구문은 자바의 정적 메소드 호출이나 정적 필드 사용 구문과 같아진다

```kt
class A {
  companion object {
    fun bar() {
      println("Companion object called")
    }
  }
}

>>> A.bar()
Companion object called
```

companion object에서 private 생성자를 호출할 수 있다. 이를 활용한 팩토리 패턴등을 구현할 수 있다.

```kt
class User private constructor(val nickname: String) {
  companion object {
    fun newSubscribingUser(email: String) = User(email.substringBefore('@'))
    fun newFacebookUser(accountId: Int) = User(getFacebookName(accountId))
  }
}

>>> val subscribingUser = User.newSubscribingUser("bob@gmail.com")
>>> val facebookUser = User.newFacebookUser(4)
>>> println(subscribingUser.nickname)
bob
```

동반 객체에 이름을 붙이거나, 인터페이스를 상속 하거나, 확장함수와 프로퍼티를 정의하는 등, 일반객체처럼 사용할 수도 있다.

##### 동반객체에 이름 붙이기

```kt
class Person(val name: String) {
  companion object Loader {
    fun fromJSON(jsonText: String): Pserson = ...
  }
}

// 두 가지 방법 모두 같은 결과이다.
>>> person = Person.Loader.fromJSON("{name: 'Dmitry'}")
>>> person = Person.fromJSON("{name: 'Dmitry'}")
```

##### 동반 객체에서 인터페이스 구현

```kt
interface JSONFactory<T> {
  fun fromJSON(jsonText: String): T
}

class Person(val name: String) {
  companion object: JSONFactory<Person> {
    overrider fun fromJSON(jsonText: String): Person = ...
  }
}
```

동반객체가 구현한 인터페이스 인스턴스를 넘길때 동반객체 외부 클래스형을 사용할 수 있다.

```kt
fun loadFromJSON<T>(factory: JSONFactory<T>): T {
  ...
}

loadFromJSON(Person) // JSONFactory 타입에 Person 타입을 넘겼다
```

> 동반객체에 이름을 붙이지 않았다면 자바에서 Companion이라는 이름으로 참조에 접근할 수 있다

##### 동반 객체 확장

앞에서 본 Person의 관심사를 명확히 분리하고 싶다고 하자. Person 클래스는 핵심 비즈니스 로직의 일부다.
비즈니스 모듈이 특정 데이터 타입에 의존하기 원하지 않는다. 따라서 역직렬화 함수를 클라이언트/서버 통신을 담당하는 모듈에 포함시키고 싶다.
확장함수를 사용하면 이런 구조를 잡을 수 있다.

```kt
class Person(val firstName: String, val lastName: String) {
  companion object { // 비어있는 동반객체 선언
  }
}

fun Person.Companion.fromJSON(json: String): Person { // 확장 함수 선언
  ...
}

val p = Person.fromJSON(json)
```

이렇게 하면 마치 동반 객체 안에서 fromJSON 함수를 정의한 것 처럼 호출 가능하다.

#### 객체 식: 익명 내부 클래스

객체 식은 자바의 익명 내부 클래스(annoymous inner class) 대신 쓰인다.
자바에서 흔히 익명 내부 클래스로 구현하는 이벤트 리스너를 코틀린에서 구현해 보자.

```kt
window.addMouseListener(
  object: MouseAdapter() { // 이름이 없고 MouseAdapter를 확장하는 객체선언
    override fun mouseClicked(e: MouseEvent) {
      ...
    }
    ...
  }
)
```

사용한 구문은 객체 선언과 같지만, 객체 이름이 빠졌다는 점만 다르다.

만약 이름을 붙여야 한다면 변수에 익명 객체를 대입하면 된다.

```kt
val listener = object: MouseAdapter() {
  override fun mouseClicked(e: MouseEvent) { ... }
  ...
}
```

한 인터페이스만 구현하거나 한 클래스만 확장할 수 있는 자바의 익명 내부 클래스와 달리,
코틀린에서는 여러 인터페이스를 구현하거나 클래스를 확장하면서 인터페이스를 구현할 수 있다.

자바의 익명 클래스와 같이 객체 식 안의 코드는 그 식이 포함된 함수의 변수에 접근할 수 있다.
하지만 자바와 달리 final이 아닌 지역변수도 객체 식 안에서 사용할 수 있다.

```kt
fun countClicks(window: Window) {
  var clickCount = 0

  window.addMouseListener(object: MouseAdapter() {
    override fun mouseClicked(e: MouseEvent) {
      clickCount++
    }
  })
}
```
