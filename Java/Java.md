# JAVA

<!-- TOC -->

- [JAVA](#java)
  - [자바 정의](#자바-정의)
  - [환경설정](#환경설정)
    - [개발에 필요한 환경설정](#개발에-필요한-환경설정)
    - [환경변수](#환경변수)
      - [윈도우즈 기준](#윈도우즈-기준)
  - [자료형](#자료형)
    - [기본자료형 (Primitive Type)](#기본자료형-primitive-type)
    - [참조 자료형 (Reference Type)](#참조-자료형-reference-type)
    - [Call by Reference, Call by Value](#call-by-reference-call-by-value)
    - [(자료)형변환](#자료형변환)
      - [자동형변환](#자동형변환)
      - [강제형변환](#강제형변환)
  - [수정자](#수정자)
  - [JVM 메모리구조](#jvm-메모리구조)
    - [Garbage Collection(가비지 컬렉션)](#garbage-collection가비지-컬렉션)
  - [클래스간의 관계](#클래스간의-관계)
    - [상속](#상속)
  - [Overloading vs Overriding](#overloading-vs-overriding)
  - [추상클래스와 인터페이스](#추상클래스와-인터페이스)
    - [추상메소드](#추상메소드)
    - [추상클래스](#추상클래스)
      - [추상클래스 선언](#추상클래스-선언)
      - [추상클래스의 목적](#추상클래스의-목적)
    - [인터페이스](#인터페이스)
      - [인터페이스 용도](#인터페이스-용도)
      - [인터페이스의 디폴트 메소드](#인터페이스의-디폴트-메소드)
  - [String Class](#string-class)
    - [String 클래스 메소드](#string-클래스-메소드)
  - [Wrapper Class](#wrapper-class)
    - [Wrapper Class 목록](#wrapper-class-목록)
  - [클래스와 인터페이스 중첩](#클래스와-인터페이스-중첩)
    - [중첩클래스(Nested Class)](#중첩클래스nested-class)
    - [중첩 인터페이스](#중첩-인터페이스)
    - [익명객체](#익명객체)
  - [Collection Framework](#collection-framework)
  - [스트림](#스트림)
  - [데이터 처리방법에 따른 스트림 유형](#데이터-처리방법에-따른-스트림-유형)
  - [File](#file)
  - [예외처리](#예외처리)
    - [예외 유형](#예외-유형)
    - [예외처리 방식](#예외처리-방식)
    - [try-with-resources 예제](#try-with-resources-예제)
  - [Thread](#thread)
    - [Thread의 생명주기](#thread의-생명주기)
    - [Thread 클래스 정의방법 3가지](#thread-클래스-정의방법-3가지)
  - [Annotation](#annotation)
    - [어노테이션의 용도](#어노테이션의-용도)
    - [유지 정책(@Retention)](#유지-정책retention)
    - [대상(@Target)](#대상target)
    - [Annotation 정의](#annotation-정의)
  - [JDBC](#jdbc)
    - [JDBC 주요객체](#jdbc-주요객체)
      - [ResultSet](#resultset)
    - [JDBC드라이버](#jdbc드라이버)
    - [바인드변수](#바인드변수)
    - [Connection Pooling](#connection-pooling)
    - [JNDI (Java Naming Directory Interface)](#jndi-java-naming-directory-interface)

<!-- /TOC -->

## 자바 정의

- sun micro >> oracle

> __자바 슬로건 : Write once run anywhere__(즉, 운영체제에 독립적)

- JAVA 특징
  - 기본 자료형과 객체자료형
  - 객체 지향 개념의 특징인 캡슐화, 상속, 다형성이 잘 적용된 언어
  - Garbage Collector를 통한 메모리 관리
  - 멀티쓰레드(Multi-thread) 사용의 간편함

## 환경설정

### 개발에 필요한 환경설정

cmd에서 자바 버전확인 java -version

자바 컴파일러 연결 확인 javac -version

### 환경변수

OS 사용할 동안 계속 참조할 수 있는 변수 : java_home

cmd에서 echo를 통해 환경변수 출력가능
  > echo %path%

%___%를 사용시 해당 환경 변수로 생각한다

#### 윈도우즈 기준

시스템 >> 고급시스템설정 >> 환경변수 에서 등록

JAVA_HOME : jdk(java development kit)경로

classpath : java.exe가 기본적으로 참조하는 경로

두 개를 등록하고 시스템 환경변수 PATH에는 JAVA_HOME\bin 을 추가한다.

## 자료형

### 기본자료형 (Primitive Type)

1. 문자 Char(2byte) : (2^16) 유니코드까지 표현(음수는 사용하지 않는다)

    사실상 자바를 포함한 응용프로그램에서 문자란 존재하지 않고 숫자를 변환해서 표현한다

    > char c = ‘A’; char c = 65;

1. 숫자

    정수 : Byte(1byte) Short(2byte) Int(4byte) Long(8byte)

    실수 : Float(4byte) Double(8byte)

1. 논리값 Boolean : 자바에서 논리값은 숫자로 대체될 수 없다.

> 문자열은 기본 자료형이 아닌 참조자료형이다.

### 참조 자료형 (Reference Type)

레퍼런스 타입 변수에 객체를 대입하면 객체가 변수에 저장 되는 것이 아니라 메모리상에 객체가 있는 위치를 가리키는 정보만 변수에 저장

### Call by Reference, Call by Value

- Call by Value - 변수에 할당된 메모리에 기본 자료형 값이 들어가 있음, 이 변수를 호출하면 값 자체를 직접 사용하게 됨

- Call by Reference - 변수에 할당된 메모리에 인스턴스의 주소값을 저장(객체를 참조한다), 호출시 원본 주소를 참조하므로 원본을 간접 조작하게 된다

### (자료)형변환

#### 자동형변환

작은 타입에서 큰 타입으로는 자동 형변환이 일어난다.

```java
byte byteVal = 5;
int intVal = byteVal; //byteVal이 int로 자동 형변환 됨
```

자식객체를 부모타입의 변수에 대입하면 자동형변환이 일어난다.
부모타입으로 변환된 이후에는 super클래스에 선언된 필드와 메소드만 접근가능하다.
다만 오버라이딩된 메소드가 있다면 자식의 메소드가 호출된다.

> cpu가 32비트에 최적화 되어있기 때문에, int형보다 작은 자료형의 연산은 내부적으로 int로 형변화시켜서 연산을 수행한다

#### 강제형변환

(자료형) 부모클래스타입 자료형 변수 : cast연산자를 사용하여 형변환 한다.

자동형변환이 발생하지 않는 관계에서 연산대상(피연산자)의 자료형이 같지 않으면 연산을 수행할 수 없기 때문에 형 변환이 발생한다

```java
byte b =3;
int a =7;
a=b;
b=(byte)a; // (byte) : cast연산자
```

강제형변환이 가능한지 확인하기 위해서 instanceof 연산자로 객체타입을 확인한다

[객체 레퍼런스 변수] instanceof [타입] : 반환형 boolean

형변환이 불가능한데 casting을 하면 ClassCastException 이 발생할 수 있다.

## 수정자

- 접근제어자
  - public : 접근 제한이 없다
  - protected : 같은 패키지 + 상속받은 자식 클래스에서 접근가능
  - default : 같은 패키지 내에서만 접근
  - private : 같은 클래스 내에서만 접근
- static : 클래스 변수, 클래스메서드를 지정하는 수정자
- final : 클래스 >> 상속불가, 메서드 >> 오버라이딩 불가, 변수 >> 그 값이 변경될 수 없음
- abstract : 클래스 >> 추상클래스, 메서드 >> 추상메서드(몸체{}없는), 변수>>X

## JVM 메모리구조

- Class(Method)영역 - ClassLoader가 읽은 데이터(클래스 정보, 멤버변수, 접근제한자...), Constant Pool, Static으로 선언된 클래스들 의 데이터가 저장된다
- Stack영역 - 호출된 method와 관련데이터(인수, 지역변수...)들이 stack 구조로 쌓인다
- Heap영역 - 인스턴스가 생성이후 위치하는 영역

### Garbage Collection(가비지 컬렉션)

- 자바는 메모리를 직접 관리하지 않는다
- 시스템에서 알아서 사용하지 않는 메모리 영역/인스턴스를 찾아 다시 사용 가능한 자원으로 회수
- System.gc()로 직접호출할 수 있으나 부른다고 바로 실행되는 것이 아니라 시스템이 적절할 때를 찾아 수행한다
- [가비지 컬렉션](http://d2.naver.com/helloworld/1329)

## 클래스간의 관계

클래스들 간에는 관계가 있다. = 이 세상의 모든 객체는 서로 관련성을 맺는다.

Has a : 객체가 합체된 상태, 즉 부품관계

Is a : 상속관계, 상 하위 관계

> Duck extends(is a) Bird (오리는 새의 기능을 온전히 수행할 수 있어야 한다)

상속관계에서 부모클래스가 생성자의 인수를 둔 경우, 모든 클래스 생성자의 첫줄에는 눈에 보이지는 않지만 super()가 생략되어 있다.
특정 객체의 인스턴스를 메모리에 올리기 위해서는 해당 객체의 부모 인스턴스가 필요하기 때문이다.

### 상속

어떤 클래스를 부모로 둘 것인지 정하고 extends 뒤에 클래스 이름을 써서 상속받는다
자바는 다중상속 (여럿의 부모클래스를 두는 행위)을 지원하지 않는다.

상속관계에서 부모클래스가 생성자의 인수를 둔 경우, 모든 클래스 생성자의 첫 줄에는 눈에 보이지는 않지만 super()가 생략되어 있다.
특정 객체의 인스턴스를 메모리에 올리기 위해서는 해당 객체의 부모 인스턴스가 필요하기 때문이다.

## Overloading vs Overriding

- Overloading(오버로딩)
  - 같은 이름의 메소드를 여러개 정의하는 것
  - 매개변수의 타입이 다르거나 개수가 달라야 한다.
  - return type과 접근 제어자는 영향을 주지 않음.
  - 부모의 메소드와 동일한 형태(리턴타입, 메소드명, 매개변수)이어야 한다.
  - 더 강한 접근제한자를 사용할 수 없다.
  - 새로운 Exception을 throws 할 수 없다.
- Overriding(오버라이딩)
  - 상속에서 나온 개념
  - 상위 클래스(부모 클래스)의 메소드를 하위 클래스(자식 클래스)에서 재정의

## 추상클래스와 인터페이스

### 추상메소드

내용이 정의되지 않은 메소드이다. 메소드에 접근제한자, 반환형, 메소드이름, 매개변수만 명시되어있다.

### 추상클래스

#### 추상클래스 선언

실질적으로는 불완전한 클래스를 의미한다. 추상메소드가 없어도 선언가능하고, 추상메서드가 단 하나라도 있으면 반드시 추상클래스로 선언해야 한다.

선언을 위해서는 abstract 키워드를 사용한다.

추상클래스는 new 연산자로 인스턴스를 생성하지 못한다. 따라서 사용을 위해서는 상속을 통한 구현만 가능하다.
  > 불완전한 부분을 overriding 해야하기 때문에

#### 추상클래스의 목적

공통되는 특성을 정의하기 위해서 사용한다. 자식클래스의 필드와 기능을 통일하고 구현을 강제하기 위해서 사용한다.

### 인터페이스

객체의 사용방법을 정의한 타입이다. 즉 기능만을 보유한 객체이다.

인터페이스는 상수와 추상메소드만을 구성멤버로 가질 수 있다. 따라서 객체로 생성할 수 없고 사용을 위해서는 클래스에서 implements 키워드로 불러와 구현해야한다.

인터페이스는 인터페이스간 다중상속을 받을 수 있다. 따라서 다중구현을 활용한 유연한 타입변환을 할 수 있다.

#### 인터페이스 용도

인터페이스는 개발한 코드에서 여건에 따라 변경될 수 있는 부분을 분리하여 사용할 객체에서 구현하게 하고, 본래의 코드는 수정하지 않기 위해 사용한다.

추가적으로 인터페이스는 여러사람이 분업 할 때 서로의 기능연결을 위한 최소한의 약속도구로 사용할 수도 있고,
한 가지 기능을 여러곳에서 별도로 구현할 때 가이드라인으로 사용할 수 있다.

#### 인터페이스의 디폴트 메소드

인터페이스를 구현하는 클래스에서는 인터페이스에서 선언된 추상메서드를 모두 오버라이드해서 구현해야한다.
이 상황에서 인터페이스에 변경이 발생한다면 인터페이스를 구현하는 모든 클래스에서 새롭게 추가된 메소드를 구현해야하는 상황이 발생한다.
이 때 인터페이스에서 디폴트 메소드를 추가하면 기존에 인터페이스를 구현하던 클래스에서 추가로 메소드를 구현해야 할 의무가 사라진다.

## String Class

```java
String str = “”; //implicit (암묵적) >> Constant Pool에 올라감
String str = new String(); //explicit (명시적)
```

String 클래스는 java.lang 패키지에 포함되어 import없이 사용가능하다.

### String 클래스 메소드

- `charAt(n)` : n번째에 있는 문자를 반환한다 (첫 글자는 0번째로 간주)

- `equals(String str)`
  - identical : 두 변수가 동일한지 비교하기 위해서는 비교연산자인 ==을 사용한다. 비교연산자 == 는 기본자료형인경우 그 값을 비교하고, 참조자료형인경우 변수가 갖고있는 주소값을 비교하게 된다.
  - equality : String 클래스는 new연산자를 사용하여 인스턴스를 생성하기도 하지만, 암묵적 생성의 경우 Constant Pool에 위치한다. String 객체의 동등성을 알아내기 위해서 `equals()`메소드를 사용한다. (즉, 값이 같은가)

- `getBytes()` : 문자열을 바이트 배열로 변환한다.

- `indexOf("x")` : 문자열에서 x가 포함되어있는지를 찾고(포함되어 있지 않다면 -1 리턴) x를 찾았다면 그 위치를 반환한다.(첫 글자는 0번째로 간주)

- `length()` : 문자열의 길이를 반환한다.

- `replace("A", "B")` : 문자열에서 A를 모두 찾아 B로 바꾼다.

- `substring(n, m)` : n번째 부터 m-1번째 사이의 문자열을 반환한다. (첫 글자는 0번째로 간주)

- `trim()` : 문자열의 앞, 뒤 공백을 제거한다.

- `valueOf(기본자료형)` : 기본자료형의 변수값을 문자열로 변환한다.

## Wrapper Class

기본자료형에 1:1로 대응되는 객체자료형

기본자료형과 객체자료형간의 형변환을 지원
  > “3” >> 3 , 3 >> “3”

- 기본자료형을 객체로 만드는 것을 Boxing이라한다.

    Boxing을 위해서는 해당WrapperClass(기본자료형 변수)를 사용하면 된다. (생성자 매개변수에 기본자료형 입력)

    만약 동일타입의 WrapperClass 자리에 기본자료형이 대입되면 자동 형변환이 일어난다.

- 객체에서 기본자료형을 얻는 것을 Unboxing이라 한다.

    Unboxing을 위해서는 WrapperClass.Value() 메소드를 이용하면 된다.

### Wrapper Class 목록

- byte >> Byte
- short >> Short
- int >> Integer
- long >> Long
- float >> Float
- double >> Double
- boolean >> Boolean
- char >> Character

## 클래스와 인터페이스 중첩

### 중첩클래스(Nested Class)

클래스 내부에 선언한 클래스이며 멤버클래스와 로컬클래스로 구분된다.
충첩클래스를 사용하면 두 클래스간 멤버들을 서로 쉽게 접근할 수 있는 장점이 있다.

내부 중첩클래스에서 외부 클래스를 참조하려면 `외부클래스명.this`

### 중첩 인터페이스

중첩 인터페이스는 클래스 멤버로 선언된 인터페이스를 말한다.

### 익명객체

클래스가 다른 클래스를 상속하거나, 인터페이스를 구현할 때 내부의 초기값, 매개변수의 매개값으로 대입된다.

## Collection Framework

java.util 패키지에서 지원한다.

- Set, List, Map 이 있다.
- List : 중복을 허용하고 순서를 가진다.
- Set : 중복을 허용하지 않고 순서도 가지지 않는다.
- Map : key 와 value 의 형태로 저장한다

## 스트림

스트림(Stream)이란? 현실에서는 물의 흐름을 말하고, 전산에서는 데이터의 흐름을 의미

- 방향에 따른 분류
  - 실행중인 프로그램으로 데이터가 흘러들어가는 스트림을 “입력” Input
  - 실행중인 프로그램에서 데이터가 흘러 나가는 스트림을 “출력” output

## 데이터 처리방법에 따른 스트림 유형

- 바이트 기반 스트림 : 입출력시 (1btye씩 처리)
- 문자 기반 스트림 : (한 문자씩- 2byte 처리)
- 버퍼 기반 스트림 : 한 라인씩 처리

## File

자바에서는 디렉토리도 일종의 파일로 간주 : `isDirectory()`, `isFile()`

## 예외처리

개발자는 코드 작성마다 모든 코드에 대하여 예외를 처리해야 할지 고민X

예외를 처리해야 할 코드인지에 대한 결정은 sun에서 이미 각 에러에 대해 case by case로 결정지어 놓았고, 예외를 처리하도록 이미 컴파일러가 강요하고 있다.

예외처리의 목적 : 비 정상 종료의 방지 == 프로그램의 불안정한 실행 방지

### 예외 유형

- Checked Exception
- unChecked Exception
  - ArrayIndexOutOfBoundsException
  - NumberFormatException
  - NullpointerException
  - ClassCastException
  - ArithmeticException

### 예외처리 방식

1. 직접처리 : try-catch-finally
1. 간접처리 : throws
1. 임의발생 : throw exception
1. 사용자정의 : extends exception

### try-with-resources 예제

```java
try(
    FileInputStream fis = new FileInputStream("input.txt");
    FileOutputStream fos = new FileOutputStream("ouput.txt");
) {
    int data = -1;
    while(true) {
        data = fis.read();
        if (fis.read()==-1) {break;}
        fos.write(data);
    }
} catch (IOException e) {
  // 예외처리
}
```

## Thread

하나의 프로세스내에서 독립적으로 실행될 수 있는 세부실행단위

하나의 프로그램에서 동시에 여러작업을 진행하면 동기화가 안되므로 불안정함

독립적으로 실행하려는 로직을 run메서드에 넣어 오버라이딩

### Thread의 생명주기

|언제 만들어져서|언제 돌아가고|언제 죽는가|
|---|---|---|
|new 할 때|JVM에게 맡길 때(start)|run()메서드 종료|

### Thread 클래스 정의방법 3가지

1. extends Thread
1. implements Runnable
1. 익명 클래스

## Annotation

### 어노테이션의 용도

- 컴파일러에게 코드 문법에러를 체크하도록 정보를 제공
- 소프트웨어 개발 툴이 빌드나 배치 시 코드를 자동으로 생성할 수 있도록 정보를 제공
- 실행시(런타임시) 특정 기능을 실행하도록 정보를 제공

### 유지 정책(@Retention)

RetentionPolicy | 설명
--- | ---
SOURCE | 소스상에서만 유지
CLASS | 바이트 코드까지 유지, 리플렉션에서 사용 불가
RUNTIME | 런타임까지 유지, 리플렉션에서 사용 가능

### 대상(@Target)

ElementType | 적용 대상
--- | ---
TYPE | 클래스, 인터페이스, ENUM
ANNOTATION_TYPE | 어노테이션
FIELD | 필드
CONSTRUCTOR | 생성자
METHOD | 메소드
LOCAL_VARIABBLE | 지역 변수
PACKAGE | 패키지

### Annotation 정의

- Controller : 클래스를 메모리에 올리는지 구분

  ```java
  @Retention(RetentionPolicy.RUNTIME)
  @Target(ElementType.TYPE)
  public @interface Controller {
  }
  ```

- RequestMapping : model을 method로 mapping하기 위해

  ```java
  @Retention(RetentionPolicy.RUNTIME)
  @Target(ElementType.METHOD)
  public @interface RequestMapping {
    public String value(); // 기본값 == value()
  }
  ```

## JDBC

자바의 데이터베이스 연동기술 java.sql 패키지에서 지원함

### JDBC 주요객체

- DriverManager : 접속시도객체
- Connection : 접속성공이후 그 정보를 담는객체 따라서 끊을때도 사용할 수 있다.
- PreparedStatement : 쿼리 수행 객체
- ResultSet : select문의 수행결과를 담아놓는 객체

#### ResultSet

아무런 옵션을 부여하지 않으면, 기본이 1칸씩 앞으로(forward)이동할 수 있는 커서를 지원한다. scroll 가능한 커서로 변환해서 쓸 수 있다.

```java
con.prepareStatement(쿼리문, 상수, 상수2);
상수1 - ResultSet.TYPE_SCROLL_INSENSITIVE
상수2 - ResultSet.CONCUR_READ_ONLY
```

레코드 총 개수 가져오는 법

```java
rs.last(); //커서를 제일 마지막으로 이동
rs.getRow(); //현재 커서의 위치반환
```

### JDBC드라이버

sun에서 명시한 스펙에따라 개발된 컴포넌트, 라이브러리

따라서 개발자는 어떠한 종류의 DBMS를 사용한다 하더라도, 동일한 인터페이스로 사용 가능하다.

### 바인드변수

DBMS도 소프트웨어이기 때문에 쿼리문장 파싱과 컴파일과정을 거치게 된다.하지만 매번 DB에서 컴파일이나 문법검사등을 일으킨다면 성능에 상당한 영향을 주게된다.

> `select * from member where id=입력값 and pass=입력값` 입력값이 변하더라도, 전체문장의 변경으로 간주하지 않도록 한다. 이때 바인드 변수 지원됨

### Connection Pooling

웹분야처럼 __서버와 연결이 유지되지 않는 stateless 특징__을 갖는 경우,
클라이언트의 요청마다 db와의 접속을 시도하게 되면 너무 많은 자원을 낭비하게 되므로 (속도저하, 접속시도에 따르는 시간지연등)
클라이언트의 접속이 없더라도 메모리에 미리 여유분의 접속객체를 확보해놓고,
요청이 있을 때마다 새로운 접속을 일으키는 것이 아니라 이미 생성된 접속객체를 할당하여 데이터베이스 업무를 처리할 수 있게 하는 것

> stateful(실시간 연결된 상태 ex-socket)

### JNDI (Java Naming Directory Interface)

server.xml

```xml
<Host name="127.0.0.1"  appBase="webapps" unpackWARs="true" autoDeploy="true">
  <Context path="" docBase="경로\WebContent" reloadable="true">
  <!-- 이하 context.xml에 작성하는 방식으로도 적용 가능 -->
    <Resource name="jdbc/myoracle" auth="Container"
              type="javax.sql.DataSource" driverClassName="oracle.jdbc.driver.OracleDriver"
              url="jdbc:oracle:thin:@127.0.0.1:1521:mysid"
              username="scott" password="tiger" maxTotal="20" maxIdle="10"
              maxWaitMillis="-1"/>
  </Context>
</Host>
```
