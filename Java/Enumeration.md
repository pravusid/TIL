# Java Enumeration

Java 1.5 버전에서 추가된 열거형 상수

## 정의

enum은 열거형 상수이나 Java에서는 다른 특성들을 지니고 있다.

## enum의 실제 타입

기본적으로 enum은 추상 클래스이다
enum 내부에 선언된 열거형 상수는, 실제로는 enum의 타입을 상속 받은 하위 클래스이다.
이 하위 클래스는 public static final 타입이며, 직접 생성할 필요가 없이 각 하나의 인스턴스로 생성된다.

```java
enum Type { // abstract class
    ADD,    // public static final class ADD extends Type
    SUB;    // public static final class SUB extends Type
}
```

## enum에 인터페이스 구현

enum은 추상 클래스이기는 하나 다른 클래스로부터 상속을 받지는 못한다. 하지만 interface는 구현가능하다.

```java
interface Interface {
    void do();
}

enum Type implements Interface{
    ADD{
        public void do() {
          System.out.println("ADD");
        }
    },
    SUB{
        public void do() {
          System.out.println("SUB");
        }
    };
}
```

enum이 public static final 이라는 점은 Singleton과 유사한 특성을 지니고 있다.

이를 활용한 Singleton 패턴, Abstract Factory 패턴, Factory Method 패턴, State 패턴, Strategy 패턴, Visitor 패턴 등
enum을 활용한 다양한 패턴을 구현할 수 있다.

## enum에서 메소드 선언

추상 클래스이므로 내부에 메소드를 선언할 수도 있다. 추상클래스이므로 추상 메소드로 선언할 수도 있다.

```java
enum Type {
    ADD,
    SUB;

    public void do(){
      System.out.println("do()");
    }

    abstract public void destory();
}
```

## enum에서 static 메소드 선언

static 메소드 선언도 가능하다

```java
public static int size() {
    return values().length; // enum의 하위타입 개수 반환
}
```

## enum 출력값

enum을 호출하면 기본적으로 하위타입 이름과 같은 값이 출력된다. (순서값이 출력되지도 않음)

enum 타입이 제공하는 기본 함수로 enum의 순서를 알 수 있는 함수가 있다.
enum의 하위타입이 몇 번째로 선언된 것인지를 반환한다 (0부터 시작)

```java
public int ordinal();
```

## enum에서 생성자 선언하기

만약 enum의 하위타입이름을 값으로 사용하지 않으려면, 생성자를 선언하여 다양한 값을 할당할 수 있다.

enum Type{
    ADD(0),
    SUB(1);

    int value;

    private Type(int value) {
        this.value = value;
    }

    public int value() {
        return value;
    }
}

`enum` 생성자의 가시성은 `private`으로 하는 것이 좋다(상속받은 하위타입들이 `public static final` 이므로))
