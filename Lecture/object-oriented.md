# 객체지향 프로그래밍

## SOLID

Robert C. Martin (Uncle Bob)이 정리한 객체지향을 위한 원칙을, Michael Feathers가 머릿글자를 따서 부름.

### Single Responsibility Principle (단일 책임 원칙)

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

### Open-Closed Principle (개방-폐쇄 원칙)

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

### Liskov Substitution Principle (리스코프 치환원칙)

하위 타입은 반드시 상위타입을 대체할 수 있어야 한다.
즉, 하위클래스는 상위클래스의 행위를 모두 수행할 수 있어야 하며, 상위클래스의 가정을 어기면 안된다.

행동의 하위형은 형 이론에서 인수형의 반공변성과 반환형의 공변성에 의존하여 정의한 일반적 기능의 하위형화보다 강한 개념이며 일반적으로 결정 불가능하다.
그러나 일반적 기능의 하위형화에 추가적인 요구사항을 강제하여 리스코프 치환 원칙을 정의하였고, 이는 클래스 계층구조의 설계에 유용하다.

일반적 기능의 하위형이 만족하는 요구사항에 덧붙인, 리스코프 치환 원칙을 만족하는 하위형의 추가적인 요구사항은 다음과 같다

- 하위형에서 선행조건은 강화될 수 없다.
- 하위형에서 후행조건은 약화될 수 없다.
- 하위형에서 상위형의 불변조건은 반드시 유지되어야 한다.

### Interface Segregation Principle (인터페이스 분리 원칙)

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

### Dependency Inversion Principle (의존성 역전 원칙)

의존(종속)은 구체가 아닌 추상과 이뤄져야 한다.
동시에 의존성 역전 원칙은 리스코프 치환 원칙을 위반하지 않도록 유의해야 한다.

- 고수준(High-Level)의 모듈은 저수준(Low-Level)의 모듈에 의존하면 안된다. 둘다 추상화에 의존해야한다.
- 추상은 세부사항(Details)에 의존해서는 안된다. 세부사항 대신 추상에 의존해야 한다

로그를 저장하는 고수준의 모듈을 예로 들어보자. 일반적으로 로그는 콘솔 혹은 파일로 출력한다.
이를 위해서 `PrintStream` / `FileOutputStream` 등으로 바이트를 출력하는 저수준의 모듈을 작성하게 된다.

고수준 모듈에서 로그를 저장하기 위해서는 저수준 모듈에 의존해야 한다.
이 때 저수준 모듈에 직접 의존하는게 아니라, 저수준 모듈들이 공통적으로 구현하고 있는 `OutputStream`이라는 인터페이스에 의존한다.

고수준 모듈과 저수준 모듈은 서로 추상계층인 인터페이스와 의존관계를 맺고 결합도를 낮출 수 있다.

## Object Oriented Programming (객체지향프로그래밍)

- 객체와 객체의 상호작용을 통해 프로그램이 동작하는 것을 말한다.
- 캡슐화 : 클래스의 내부 정의에 대해 외부에서 볼 수 없도록 하기 위한 특징
- 상속성 : 상위 클래스에서 정의한 속성을 재사용, 확장하여 개발속도를 개선시키려고 만든 개념.
- 다형성 : 같은 이름의 함수 호출에 대하여 각 객체에 따라 다른 동작을 할 수 있도록 구성하는 방법.

## Class와 Instance (추상화와 인스턴스화)

- Class : 객체를 표현하는 변수와 메서드를 담고 있는 표현식.
- Instance : class를 통해서 표현하는 객체를 실체화한 상태.

## Domain

소프트웨어가 다루고 해결해야 할 문제영역
