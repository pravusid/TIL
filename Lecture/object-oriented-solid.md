# SOLID

Robert C. Martin (Uncle Bob)이 정리한 객체지향을 위한 원칙을, Michael Feathers가 머릿글자를 따서 부름.

## Single Responsibility Principle (단일 책임 원칙)

하나의 클래스는 하나의 책임만 가지며, 클래스가 제공하는 모든 기능은 그 하나의 책임을 수행하는 데 집중되어야 한다.

책임이란 ‘변경을 위한 이유’이다.
만약 하나의 클래스에 변경을 위한 두 가지 이상의 이유가 있다면 그 클래스는 단일 책임을 위반한 것이다.

이는 관심사의 분리(Separation of Concern) 관점에서도 옳다.
하나의 기능의 작성/변경하고 이를 수행하는 것이 다른 기능을 오염시키지 않아야 하기 때문이다.

응집(Cohension)과 결합(Coupling)도 SRP와 깊은 관련이 있다.
하나의 책임만 가진 클래스의 경우 다른 모듈과의 결합도는 낮고, 동일한 책임을 수행하는 기능들의 응집도는 높은 상태이다.

```java
class Animal {
    Animal(String name){ ... }
    String getAnimalName() { ... }
    void saveAnimal(Animal animal) { ... }
}
```

위의 클래스는 Animal 객체의 프로퍼티 관리와 Animal 데이터베이스 관리가 결합되어 있다.
이는 다음과 같이 분리할 수 있다.

```java
class Animal {
    Animal(String name){ ... }
    String getAnimalName() { ... }
}

class AnimalDB {
    Animal getAnimal(String name) { ... }
    void saveAnimal(animal: Animal) { ... }
}
```

## Open-Closed Principle (개방-폐쇄 원칙)

소프트웨어의 구성요소(컴포넌트, 클래스, 모듈, 함수)는 확장에는 열려있고, 변경에는 닫혀있어야 한다는 원리이다.

만약 소프트웨어에 대한 요구사항이 변경된다면 기존 구성요소는 수정하지 말아야 하고,
기존 구성요소를 확장해서 재사용하는 것은 쉽게 가능해야 한다는 의미이다.

고객 유형에 따른 할인 로직 예제를 살펴보자.

```java
class Discount {
    DiscountPrice giveDiscount() {
        return this.price * 0.2
    }
}
```

요구사항 변경으로 VIP 고객에게는 20%를 추가 할인해 준다면,

```ts
class Discount {
    DiscountPrice giveDiscount() {
        if(this.customer.equals("fav")) {
            return this.price * 0.2;
        }
        if(this.customer.equals("vip")) {
            return this.price * 0.4;
        }
    }
}
```

제어문으로 논리 분기를 만들어 적용할 수도 있을 것이다.
그러나 기존에 작성된 코드를 수정하게 되므로 OCP원칙을 지켰다고 볼 수는 없다.

대신 Discount 클래스를 확장하여 사용할 수 있을 것이다.

```ts
class VIPDiscount extends Discount {
    DiscountPrice getDiscount() {
        return super.getDiscount() * 2;
    }
}
```

## Liskov Substitution Principle (리스코프 치환원칙)

하위 타입은 반드시 상위타입을 대체할 수 있어야 한다.
즉, 하위클래스는 상위클래스의 행위를 모두 수행할 수 있어야 하며, 상위클래스의 가정을 어기면 안된다.

행동의 하위형은 일반적인 함수의 하위형화(형 이론에서 인수형의 반공변성과 반환형의 공변성에 의존하여 정의한)보다 강한 개념이며 일반적으로 결정 불가능하다.

리스코프 치환원칙은 객체 지향 프로그래밍 언어에서 채용된 몇 가지 표준적인 요구사항을 강제한다.

- 하위형에서 메서드 인수의 반공변성
- 하위형에서 반환형의 공변성
- 하위형에서 메서드는 상위형 메서드에서 던져진 예외의 하위형을 제외하고 새로운 예외를 던지면 안된다

여기에 더하여 하위형이 만족해야하는 행동 조건 몇 가지가 있다.

- 하위형에서 선행조건은 강화될 수 없다.
- 하위형에서 후행조건은 약화될 수 없다.
- 하위형에서 상위형의 불변조건은 반드시 유지되어야 한다.
- 이력 제약 조건 (History constraint): 객체는 그 자신의 메서드를 통해서만 수정(캡슐화)할 수 있는 것으로 간주된다
  - 하위형은 상위형에 없는 메서드를 추가할 수 있기 때문에, 추가된 메서드를 통해 상위형에서 허용하지 않는 하위형 상태의 변경을 일으킬 수 있다.
  - 이 제약조건의 위반을 확인하기 위해 변경 가능 지점을 변경 불가 지점의 하위형으로 정의해 볼 수 있다.
  - 변경 불가 지점의 이력은 생성한 이후 상태가 항상 동일해야 하기 때문에, 앞에서 가정한 정의는 이력 제약조건의 위반이다
  - 따라서 일반적으로 변경 가능 위치를 이력에 포함할 수 없다.
  - 반면 하위형에 추가된 필드는 상위형의 메소드로 감시할 대상이 아니기 때문에 안정적으로 수정할 수 있다.

## Interface Segregation Principle (인터페이스 분리 원칙)

클라이언트는 사용하지 않는 인터페이스 메소드에 의존하지 않아야 한다.
(스스로 구현/사용하지 않을 기능의 인터페이스의 메소드까지 갖지 않아야 한다)

```java
interface Shape {
    void drawCircle();
    void drawSquare();
    void drawRectangle();
}
```

위의 인터페이스를 구현하는 클래스는 세 메소드 모두 구현해야 한다.

```java
class Circle implements Shape {
    void drawCircle(){
        //...
    }
    void drawSquare(){
        //...
    }
    void drawRectangle(){
        //...
    }
}
class Square implements Shape {
    void drawCircle(){
        //...
    }
    void drawSquare(){
        //...
    }
    void drawRectangle(){
        //...
    }
}
class Rectangle implements Shape {
    void drawCircle(){
        //...
    }
    void drawSquare(){
        //...
    }
    void drawRectangle(){
        //...
    }
}
```

각각의 도형은 불필요한 메소드를 구현하게 된다. (원이 정사각형과 사각형을 그리는 등...)

ISP 원칙에 따라 필요한 행위(action)를 구분하여 다른 인터페이스로 분리하여야 한다.

```java
interface Circle {
    void drawCircle();
}
interface Square {
    void drawSquare();
}
interface Rectangle {
    void drawRectangle();
}

class Circle implements Circle {
    void drawCircle() {
        //...
    }
}
class Square implements Square {
    void drawSquare() {
        //...
    }
}
class Rectangle implements Rectangle {
    void drawRectangle() {
        //...
    }
}
```

## Dependency Inversion Principle (의존성 역전 원칙)

Dependency Inversion Principle는 IoC의 한 형태이다.

의존(종속)은 구체가 아닌 추상과 이뤄져야 한다. 동시에 의존성 역전 원칙은 리스코프 치환 원칙을 위반하지 않도록 유의해야 한다.

- 고수준(High-Level)의 모듈은 저수준(Low-Level)의 모듈에 의존하면 안된다. 둘다 추상(인터페이스)에 의존해야한다.
- 추상(Abstractions)은 세부사항(Details)에 의존해서는 안된다. 세부사항(Details)은 추상(Abstractions)에 의존해야 한다

로그를 저장하는 고수준의 모듈을 예로 들어보자. 일반적으로 로그는 콘솔 혹은 파일로 출력한다.
이를 위해서 `PrintStream` / `FileOutputStream` 등으로 바이트를 출력하는 저수준의 모듈을 작성하게 된다.

고수준 모듈에서 로그를 저장하기 위해서는 저수준 모듈에 의존해야 한다.
이 때 저수준 모듈에 직접 의존하는게 아니라, 저수준 모듈들이 공통적으로 구현하고 있는 `OutputStream`이라는 인터페이스에 의존한다.

고수준 모듈과 저수준 모듈은 서로 추상계층인 인터페이스와 의존관계를 맺으므로 구체적인 의존관계는 런타임에 결정되며 결합도를 낮출 수 있다.

### Inversion of Control (IoC)

IoC는 일반적으로 Hollywood Principle라 불리기도 하는데 "나에게 연락하지마, 내가 필요하면 연락할께" 라고 표현된다(GoF).

IoC는 명령 흐름이 전통적인 방식에서 역전되는 것을 말하는 (너무나도)일반적인 용어이다.
일반적으로 프레임워크에서 내가 규칙에 따라 작성한 코드를 호출하는 것이 이에 해당한다.

IoC 존재여부가 프레임워크와 라이브러리를 구분하기도 한다.

라이브러리는 호출한 기능을 수행한 후 제어를 반환한다.
그러나 프레임워크를 사용할 때는 약속된 규칙에 따라 프레임워크가 호출할 코드를 작성하게 된다.

### IoC Container

코드에서 객체의 생성과 주입을 명령하는 일은 반복적으로 이루어진다.
이러한 의존성의 생성과 주입 제어를 넘겨받은 컨테이너를 IoC Container이다(컨테이너가 스스로 구현을 찾는다).

IoC Container 역시 너무나 일반적인 용어이므로 혼란을 불러오게 된다.

따라서 Martin Fowler는 Container가 수행하는 IoC(의존성 생성/주입을 제어하는 것)를 구체화하여 DI라고 명칭하였다.

### Dependency Injection (DI)

DIP를 적용한다고 해도 반드시 DI를 할 필요는 없다. (역으로, DI를 적용하였다고 하더라도 DIP를 적용하여야 하는 것은 아니다.)

다만 DIP는 DI를 동반하는 경우에 의존의 추상화를 최대한 유지하여, 런타임에 `type`과 인스턴스가 결정될 것이다.

위와 같은 관점은 다음의 Martin Fowler의 정의와 일치한다.

> DI에서 dependency를 의존객체라고 보는 것은 적절치 않다. 대신 코드 레벨에서 제거했지만, 런타임시에 가지게 해주는 의존관계 또는 의존성으로 볼 수 있다.

해당 객체가 의존은 추상에 하더라도 구체적인 인스턴스를 생성하거나(추상 타입에 new 연산자로 생성한 object를 할당), Service Locator를 활용할 수도 있다. (둘다 안티패턴이다)

- Service Locator는 클래스의 의존성을 숨긴다(인터페이스에 드러나지 않음).
- Service Locator에서 의존성이 호출되기 전 등록되었는지 컴파일러가 확인해주지 않는다(런타임에 오류가 발생할 수 있다).
- 마찬가지로 테스트 작성시에도 테스트상 호출될 의존성을 사전에 Service Locator에 등록해야 한다. 즉, 테스트 작성시 테스트가 실행될 코드 내부를 파악하고 있어야 한다는 것이다.
- Service Locator는 static 영역에서 구성되기 마련이므로 추가적인 문제점을 불러온다(사용 및 테스트에서 유연하지 않다: 병렬 테스팅이 어려움...)
- Service Locator의 문제점은 리팩토링에서도 동일한 어려움을 불러온다.
- 역으로 Service Locator의 문제점은 DI의 장점이기도 하다.

객체에서 추상화를 유지한 채 의존성을 표현하기 위해 여러 방법을 사용할 수 있으나, 일반적으로 DI가 많이 사용된다.

DI를 사용하는 경우 객체내부에서 의존관계를 전혀 표현하지 않고 주입받은 의존성(인터페이스)을 사용하기만 하면된다.
