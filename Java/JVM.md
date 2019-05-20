# Java Virtual Machine

## JVM

자바 가상 머신은 자바 바이트 코드를 실행하는 환경이다. 자바 바이트코드는 자바와 기계어 사이의 중간 언어이다.
자바 바이트코드는 JVM이 해석하므로 플랫폼 의존적이지 않다.

## JVM 특징

- 스택 기반의 가상 머신: 대표적인 컴퓨터 아키텍처인 인텔 x86 아키텍처나 ARM 아키텍처와 같은 하드웨어가 레지스터 기반으로 동작하는 데 비해 JVM은 스택 기반으로 동작한다.
- 심볼릭 레퍼런스: 기본 자료형(primitive data type)을 제외한 모든 타입(클래스와 인터페이스)을 명시적인 메모리 주소 기반의 레퍼런스가 아니라 심볼릭 레퍼런스를 통해 참조한다.
- 가비지 컬렉션(garbage collection): 클래스 인스턴스는 사용자 코드에 의해 명시적으로 생성되고 가비지 컬렉션에 의해 자동으로 파괴된다.
- 기본 자료형을 명확하게 정의하여 플랫폼 독립성 보장: C/C++ 등의 전통적인 언어는 플랫폼에 따라 int 형의 크기가 변한다. JVM은 기본 자료형을 명확하게 정의하여 호환성을 유지하고 플랫폼 독립성을 보장한다.
- 네트워크 바이트 오더(network byte order): 자바 클래스 파일은 네트워크 바이트 오더를 사용한다. 인텔 x86 아키텍처가 사용하는 리틀 엔디안이나, RISC 계열 아키텍처가 주로 사용하는 빅 엔디안 사이에서 플랫폼 독립성을 유지하려면 고정된 바이트 오더를 유지해야 하므로 네트워크 전송 시에 사용하는 바이트 오더인 네트워크 바이트 오더를 사용한다. 네트워크 바이트 오더는 빅 엔디안이다.

## 바이트코드

<https://d2.naver.com/helloworld/1230>

자바 소스코드를 자바 바이트 코드로 컴파일 하면 자바 클래스 파일이 생성된다.
바이트코드는 바이너리 파일이므로 이를 읽기 위해서 disassembler를 사용할 수 있다.

### 메소드 바이트코드

```java
public void add(java.lang.String);
Code:
0: aload_0
1: getfield #15; //Field admin:Lcom/nhn/user/UserAdmin;
4: aload_1
5: invokevirtual #23; //Method com/nhn/user/UserAdmin.addUser:(Ljava/lang/String;)V
8: return
```

바이트코드는 `위치: 바이트코드 명령어(OpCode) #피연산자`로 구성된다.

바이트 코드 명령어는 1 바이트의 바이트 번호로 표현되므로 OpCode는 최대 256개가 된다.

OpCode들은 `aload_0 = 0x2a, getfield = 0xb4, invokevirtual = 0xb6`로 표현되며
위의 코드를 실제로 보면 `2a b4 00 0f 2b b6 00 17 57 b1` 와 같이 표기된다

피연산자가 필요없는 OpCode는 바로 다음 바이트가 다음 명령어의 OpCode가 되지만, 피연산자가 필요한 경우 다음 바이트는 피연산자가 위치한다.
위의 경우 `getfield`와 `invokevirtual`은 각각 2바이트의 피연산자를 가지며(피연산 명령어의 위치), 따라서 다음 OpCode와의 간격이 2바이트가 된다.

### 클래스 바이트코드

```java
ClassFile {
u4 magic; // 클래스 파일 구분자, 항상 0xCAFEBABE임
u2 minor_version; u2 major_version; // 버전정보; 하위호환성 유지가 중요함
u2 constant_pool_count; cp_info constant_pool[constant_pool_count-1]; // 상수 풀 정보
u2 access_flags; // 클래스 modifier 정보
u2 this_class; u2 super_class; // this, super에 해당하는 클래스의 상수풀 내의 인덱스
u2 interfaces_count; u2 interfaces[interfaces_count]; // 클래스가 구현한 인터페이스 개수와, 인터페이스에 대한 상수풀 내의 인덱스
u2 fields_count; field_info fields[fields_count]; // 클래스 필드 개수와 정보
u2 methods_count; method_info methods[methods_count]; // 클래스의 메소드 개수와 정보
u2 attributes_count; attribute_info attributes[attributes_count]; // 클래스내의 속성 개수와 정보
}

Compiled from "UserService.java"
public class com.nhn.service.UserService extends java.lang.Object
SourceFile: "UserService.java"
minor version: 0
major version: 50
Constant pool:
const #1 = class #2; // com/nhn/service/UserService
const #2 = Asciz com/nhn/service/UserService;
const #3 = class #4; // java/lang/Object
const #4 = Asciz java/lang/Object;
const #5 = Asciz admin;
// ... 후략
```

## JVM 구조

![JVM](https://raw.githubusercontent.com/pravusid/TIL/master/Java/img/jvm-struct.png)

Class Loader가 컴파일된 자바 바이트코드를 Runtime Data Area에 불러오고 Execution Engine이 바이트코드를 실행한다.

### Class Loader(클래스 로더)

자바는 런타임에 클래스를 처음으로 참조할 때 해당 클래스를 로드하고 링크하는데, JVM의 클래스 로더가 동적로드를 담당한다.

- 계층 구조
  - 클래스 로더끼리 부모-자식 관계를 이루어 계층 구조로 생성된다
  - 최상위 클래스 로더는 부트스트랩 클래스 로더(Bootstrap Class Loader)

- 위임 모델
  - 계층 구조를 바탕으로 클래스 로더끼리 로드를 위임하는 구조로 동작한다

- 가시성(visibility) 제한
  - 하위 클래스 로더는 상위 클래스 로더의 클래스를 찾을 수 있지만, 상위 클래스 로더는 하위 클래스 로더의 클래스를 찾을 수 없다

- 언로드 불가
  - 클래스 로더는 클래스를 로드할 수는 있지만 언로드할 수는 없다
  - 언로드 대신, 현재 클래스 로더를 삭제하고 새로운 클래스 로더를 생성할 수 있다

각 클래스 로더는 로드된 클래스들을 보관하는 네임스페이스(namespace)를 갖는다.
클래스를 로드할 때 이미 로드된 클래스인지 확인하기 위해 네임스페이스에 보관된 FQCN(Fully Qualified Class Name)을 기준으로 클래스를 찾는다.
비록 FQCN이 같더라도 네임스페이스가 다르면(다른 클래스 로더가 로드한 클래스) 다른 클래스로 간주한다.

로더가 클래스 로드를 요청받으면, 이전에 로드된 클래스인지 클래스 로더 캐시를 확인하고, 없으면 상위 클래스 로더를 거슬러 올라가며 확인한다.
부트스트랩 클래스 로더까지 확인해도 없으면 요청받은 클래스 로더가 파일 시스템에서 해당 클래스를 찾는다.

클래스 로더 위임모델의 계층은 다음과 같다

- 부트스트랩 클래스 로더
  - JVM을 기동할 때 생성되며, Object 클래스들을 비롯하여 자바 API들을 로드한다
  - 다른 클래스 로더와 달리 자바가 아니라 네이티브 코드로 구현되어 있다

- 익스텐션 클래스 로더(Extension Class Loader)
  - 기본 자바 API를 제외한 확장 클래스들을 로드한다. 다양한 보안 확장 기능 등을 여기에서 로드하게 된다.
  - 부트스트랩 클래스 로더와 익스텐션 클래스 로더는 JVM 자체의 구성 요소들을 로드하는 것이라 할 수 있다

- 시스템 클래스 로더(System Class Loader)
  - 시스템 클래스 로더는 애플리케이션의 클래스들을 로드한다
  - 사용자가 지정한 $CLASSPATH 내의 클래스들을 로드한다.

- 사용자 정의 클래스 로더(User-Defined Class Loader)
  - 애플리케이션 사용자가 직접 코드 상에서 생성해서 사용하는 클래스 로더이다.

웹 애플리케이션 서버(WAS)와 같은 프레임워크는 웹 애플리케이션들, 엔터프라이즈 애플리케이션들이 서로 독립적으로 동작하게 하기 위해 사용자 정의 클래스 로더를 사용한다.

클래스 로더가 아직 로드되지 않은 클래스를 찾으면, 다음과 같은 과정을 거쳐 클래스를 로드하고 링크하고 초기화한다.

- 로드
  - 클래스를 파일에서 가져와서 JVM의 메모리에 로드한다

- 검증(Verifying)
  - 읽어 들인 클래스가 자바 언어 명세(Java Language Specification) 및 JVM 명세에 명시된 대로 잘 구성되어 있는지 검사한다
  - 클래스 로드의 전 과정 중에서 가장 복잡하고 시간이 많이 걸린다

- 준비(Preparing)
  - 클래스가 필요로 하는 메모리를 할당하고, 클래스에서 정의된 필드, 메소드, 인터페이스들을 나타내는 데이터 구조를 준비한다

- 분석(Resolving)
  - 클래스의 상수 풀 내 모든 심볼릭 레퍼런스를 다이렉트 레퍼런스로 변경한다

- 초기화
  - 클래스 변수들을 적절한 값으로 초기화한다. (static initializer들을 수행하고, static 필드들을 설정된 값으로 초기화)

### Runtime Data Area

런타임 데이터 영역은 JVM이라는 프로그램이 운영체제 위에서 실행되면서 할당받는 메모리 영역이다.

- 런타임 데이터 영역은 6개의 영역
- 쓰레드별 생성: PC 레지스터(PC Register), JVM 스택(JVM Stack), 네이티브 메소드 스택(Native Method Stack)
- 모든 쓰레드 공유: 힙(Heap), 메소드 영역(Method Area), 런타임 상수 풀(Runtime Constant Pool)

#### PC Register

각 Thread마다 하나씩 존재하며 현재 수행중인 Java Virtual Machine Instruction의 주소를 갖는다.

PC Register는 Native Pointer 혹은 Method Bytecode 값을 가지며, Native Method를 수행할 경우 JVM을 거치지 않고 API를 통해 바로 수행한다.

#### JVM Stack

JVM 스택은 각 Thread마다 하나씩 존재하며 Thread가 시작될 때 생성된다. Stack Frame이라는 구조체를 저장하는 스택이다.
예외 발생 시 printStackTrace() 등의 메소드로 보여주는 Stack Trace의 각 라인은 하나의 스택 프레임을 표현한다.

##### 스택 프레임 (Stack Frame)

JVM 내에서 메소드가 수행될 때마다 하나의 스택 프레임이 생성되어 해당 스레드의 JVM 스택에 추가되고 메소드가 종료되면 스택 프레임이 제거된다.
각 스택 프레임은 Local Variable Array, Operand Stack, 현재 실행 중인 메소드가 속한 클래스의 런타임 상수 풀에 대한 레퍼런스를 갖는다.
지역 변수 배열, 피연산자 스택의 크기는 컴파일 시에 결정되기 때문에 스택 프레임의 크기도 메소드에 따라 크기가 고정된다.

- 지역 변수 배열 (Local Variable Array)
  - index 0은 메소드가 속한 클래스 인스턴스의 this 레퍼런스
  - index 1부터는 메소드에 전달된 파라미터들이 저장됨
  - 메소드 파라미터 이후에는 메소드의 지역 변수들이 저장됨

- 피연산자 스택 (Operand Stack)
  - 각 메소드는 피연산자 스택과 지역 변수 배열 사이에서 데이터를 교환하고, 다른 메소드 호출 결과를 추가하거나(push) 꺼낸다(pop)
  - 피연산자 스택 공간이 얼마나 필요한지는 컴파일할 때 결정할 수 있으므로, 피연산자 스택의 크기도 컴파일 시에 결정된다

- 메소드가 속한 클래스 런타임 상수풀 레퍼런스

#### Native method stack

코드가 아닌 실제 실행할 수 있는 기계어로 작성된 프로그램을 실행시키는 영역이다.
JAVA가 아닌 다른 언어로 작성된 코드를 위한 공간이다. JAVA Native Interface를 통해 바이트 코드로 전환하여 저장하게 된다.
일반 프로그램처럼 커널이 스택을 잡아 독자적으로 프로그램을 실행시키는 영역이다. 이 부분을 통해 C code를 실행시켜 Kernel에 접근할 수 있다.

#### Heap (힙 영역)

생성된 객체와 배열을 저장하는 가상 메모리 공간이다.

힙은 세 부분으로 나눌 수 있다

- Permanent Generation
  - 생성된 객체 정보의 주소값이 저장된 공간이다
  - Class loader에 의해 load되는 Class, Method 등에 대한 Meta 정보가 저장되는 영역이고 JVM에 의해 사용된다
  - Reflection을 사용하여 동적으로 클래스가 로딩되는 경우에 사용된다

- New/Young 영역
  - Eden : 객체들이 최초로 생성되는 공간
  - Survivor 0 / 1 : Eden에서 참조되는 객체들이 저장되는 공간

- Old 영역
  - New area에서 일정 시간 참조된, 살아남은 객체들이 저장되는 공간
  - Eden영역에 객체가 가득차게 되면 첫번째 GC(minor GC)가 발생한다
  - Eden영역에 있는 값들을 Survivor 1 영역에 복사하고 이 영역을 제외한 나머지 영역의 객체를 삭제한다

#### Method Area (Class area / Static area)

클래스 정보를 처음 메모리 공간에 올릴 때 초기화되는 대상을 저장하기 위한 메모리 공간.
 JVM이 읽어 들인 각각의 클래스와 인터페이스에 대한 런타임 상수 풀, 필드와 메소드 정보, Static 변수, 메소드의 바이트코드 등을 보관한다.

이 공간에는 **Runtime Constant Pool**이라는 별도의 관리 영역도 함께 존재한다.
클래스 파일 포맷에서 constant_pool 테이블에 해당하는 영역이다.

각 클래스와 인터페이스의 상수뿐만 아니라, 메소드와 필드에 대한 모든 레퍼런스까지 담고 있는 테이블이다.
즉, 어떤 메소드나 필드를 참조할 때 JVM은 런타임 상수 풀을 통해 해당 메소드나 필드의 실제 메모리상 주소를 찾아서 참조한다.

- Field Information: 멤버변수의 이름, 데이터 타입, 접근 제어자에 대한 정보

- Method Information: 메소드의 이름, 리턴타입, 매개변수, 접근제어자에 대한 정보

- Type Information
  - Class, Interface 구분
  - Type의 속성, 전체 이름
  - Super Class의 전체 이름 (Interface 이거나 Object인 경우 제외)

### Execution Engine(실행 엔진)

클래스 로더를 통해 JVM 내의 런타임 데이터 영역에 배치된 바이트코드는 실행 엔진에 의해 실행된다.
실행 엔진은 자바 바이트코드를 명령어 단위로 읽어서 기계가 실행할 수 있는 형태로 변경하여 실행한다.

이 때 인터프리터, JIT 두 가지 방식을 사용하게 된다.

#### Interpreter(인터프리터)

인터프리터 방식은 자바 바이트 코드를 명령어 단위로 하나씩 읽어서 실행한다.

하나씩 해석하고 실행하기 때문에 바이트코드 하나하나의 해석은 빠른 대신 인터프리팅 결과의 실행은 느리다는 단점을 가지고 있다.
흔히 얘기하는 인터프리터 언어의 단점을 그대로 가진다고 볼 수 있다.

바이트코드라는 '언어'는 기본적으로 인터프리터 방식으로 동작한다.

#### JIT(Just-In-Time)

인터프리터 방식의 단점을 보완하기 위해 도입된 JIT 컴파일러이다.
인터프리터 방식으로 실행하다가 적절한 시점에 바이트코드 전체를 컴파일하여 네이티브 코드로 변경하고,
이후에는 해당 더 이상 인터프리팅 하지 않고 네이티브 코드로 직접 실행하는 방식이다.
네이티브 코드는 캐시에 보관하기 때문에 한 번 컴파일된 코드는 빠르게 수행하게 된다.

JIT 컴파일러를 통하면 바이트코드를 인터프리팅하는 것보다 오래걸리므로, 한 번만 실행되는 코드라면 컴파일하지 않고 인터프리팅하는 것이 유리하다.
따라서 JIT 컴파일러를 사용하는 JVM들은 내부적으로 해당 메소드가 얼마나 자주 수행되는지 체크하고, 일정 정도를 넘을 때에만 컴파일을 수행한다.

#### Garbage collector

Java에서는 개발자가 프로그램 코드로 메모리를 명시적으로 해제하지 않기 때문에
가비지 컬렉터(Garbage Collector)가 더 이상 필요 없는 (쓰레기) 객체를 찾아 지우는 작업을 한다.
