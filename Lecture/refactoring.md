# Refactoring

## 리팩토링 개론

### 리팩토링은 무엇인가

리팩토링: 겉으로 드러나는 기능은 그대로 둔 채, 알아보기 쉽고 수정하기 간편하게 소프트웨어 내부를 수정하는 작업

리팩토링의 목적은 소프트웨어를 더 이해하기 쉽고 수정하기 쉽게 만드는 것이다.
리팩토링은 성능 최적화와 상반된다. 성능 최적화를 수행하면 코드를 파악하기 더 어려워질 때가 많다.

소프트웨어 개발에 리팩토링을 적용할 때 기능 추가와 리팩토링이라는 별개의 두 작업에 시간을 분배해야 한다.

- 기능을 추가할 땐 코드를 수정하지 말고 기능만 추가해야 하며, 진행상태를 파악하려면 테스트를 추가하고 통과여부를 확인한다
- 리팩토링 할 때는 코드를 추가하지 말고 코드 구조 개선만 해야 한다
- 인터페이스가 변경되어 그에 맞게 코드를 고치는 것이 불가할 때는 테스트만 변경해야 한다

### 리팩토링 목적

- 소프트웨어 설계 개선
  - 지속적으로 코드를 수정하면 코드 구조가 복잡해진다
  - 코드의 처음 설계구조를 유지하려면 정기적으로 리팩토링을 실시해야 한다
  - 특히 설계가 깔끔하지 않으면 중복코드로 인해 쓸데없이 코드가 길어진다

- 소프트웨어 이해를 쉽게 한다
  - 코드에는 나중에 수정될 다른 개발자에 대한 고려가 반영되어야 한다
  - 코드를 파악하기 쉽게 수정하기 위한 과정이 리팩토링이다
  - 리팩토링을 하면서 낮선 코드를 쉽게 이해할 수도 있다

- 버그를 찾기 쉬워진다
  - 프로그램 구조를 명료하게 만들어 특정해놓은 전제들이 확실해지면 버그 발견하기가 쉬워진다

- 프로그래밍 속도가 빨라진다
  - 깔끔한 설계는 소프트웨어 개발속도를 높이기위한 핵심이다
  - 깔끔하지 않다면 버그를 찾고 중복코드를 찾느라 시간이 길어진다

### 리팩토링이 필요할 때

- 같은 작업의 삼진 아웃
  - 어떤 작업을 처음 하면 그냥 한다.
  - 비슷한 작업을 두 번째 하면 중복이라도 그냥 한다
  - 그러나 비슷한 작업을 세 번째 하게 되면 리팩토링을 실시해야 한다

- 기능을 추가할 때
  - 기능을 추가할 때 기존 코드를 이해하기 위해서, 혹은 이해하기 쉽게 만들기 위해 리팩토링이 필요하다
  - 리팩토링 여부는 이해가 쉽게 고칠 수 있는지 고민한 뒤 결정한다
  - 기능 추가 중 설계 때문에 작업의 어려움을 느낀다면 리팩토링을 실시해야 한다

- 버그 수정시
  - 버그를 수정하려면 코드의 기능을 파악하여야 하고 이를 위해 리팩토링을 실시한다
  - 버그 리포팅을 받았을 때도 리팩토링이 필요하다. 버그 존재를 파악하기 힘든 구조가 존재했기 때문이다

- 코드 검수시
  - 코드를 검수한다면 리팩토링을 통해 자신의 코드를 이해하기 쉽게 만들고 타인의 코드를 쉽게 이해하여야 한다

### 인다이렉션과 리팩토링

소프트웨어 공학자들은 인다이렉션을 광적으로 좋아한다는 점에서 대부분의 리팩토링에서 인다이렉션이 적용되는 것은 당연한 일이다.

그러나 인다이렉션은 양날의 검이다. 한 부분을 둘로 쪼개면 관리할 부분이 늘어난다.
또한 객체가 또 다른 객체에 작업을 위임하기 때문에 코드를 알아보기 힘들어질 수도 있다.

하지만 인다이렉션은 다음과 같은 장점이 있다

- 로직을 공유: 하나의 로직을 여러곳에서 공유할 수 있다
- 의도와 구현부를 분리: 클래스명과 메소드명으로 의도를 드러낼 수 있고, 클래스나 메소드의 내부 코드를 통해 의도를 어떻게 구현했는지 보여줄 수 있다
- 수정 부분을 분리: 한 객체를 두 위치에서 사용했는데 객체를 수정하면 두곳 다 변경될 위험이 있다. 이 때 하위클래스로 수정할 수 있다
- 조건문을 코드화: 조건문을 메시지로 바꾸면 중복코드가 줄어들어 명료해지며 유연성도 높아진다

### 리팩토링 관련 문제

#### 데이터베이스

수많은 비즈니스 애플리케이션은 데이터베이스 스키마와 강력히 결합되어 있다.
데이터베이스 스키마와 객체 모델의 상호 의존성을 최소화 하여 구성해도,
데이터베이스 스키마를 수정하면 데이터도 이전해야 하는데 시간도 오래 걸리며 위험성도 높다.

#### 인터페이스 변경

객체의 장점중 하나는 인터페이스를 건드리지 않고 내부 구현을 수정할 수 이:ㅆ다는 점이다.

하지만 인터페이스를 수정하면 어떤 문제가 발생할지 알 수 없다.
리팩토링에서 불안한 점은 상당수의 리팩토링이 인터페이스를 건드린다는 것이다.

만일 메소드를 사용하는 모든 코드에 접근할 수 있다면 메소드명을 변경하는 것은 문제가 되지 않을 수 이:ㅆ다.
public 메소드라 해도 메소드를 호출하는 부분을 수정할 수 있다면 변경가능하다.

그러나 인터페이스가 public 보다 더 적극적으로 공개된 published 인터페이스라면 문제가 생긴다.

published 인터페이스(공개한 라이브러리 인터페이스의 경우) 라면 해당 인터페이스를 사용하는 부분이 수정되기 전까지는
기존 인터페이스와 새 인터페이스를 모두 그대로 유지해야 한다.

이 경우 기존 인터페이스가 새 인터페이스를 호출하게 하면 된다. 더불어 deprecation을 통해 더 이상 사용하지 않는 인터페이스임을 알려야한다.

#### 리팩토링 하면 안 되는 상황

기존 코드를 리팩토링 할 수 있지만, 너무나 지저분해서 차라리 새로 작성하는 것이 나은 경우가 있다.

만일 코드가 돌아가지 않는다면 완전히 새로 작성해야 한다는 신호이다.
테스트만 하려다 코드가 안정화 될 수 없는 경우도 그런 경우이다.

종료 시점이 임박했을 때도 리팩토링은 피해야 한다.

### 리팩토링과 설계

리팩토링은 설계를 보완하는 특수한 역할을 한다.
설계가 핵심이고 프로그래밍은 구현 도구에 불과하다고 생각할 수 있으나, 설계에는 빈틈이 많다.

사전 설계 없이 리팩토링만 하는 방법을 택할수도 있지만, 결국 설계를 수정해야 하는 시점이 오면 큰 비용이 발생할 수도 있다.

리팩토링을 실시하더라도 여전히 설계는 해야 하지만, 사전 설계 과정에서 완벽함을 추구할 필요 없이 적당한 솔루션만 생각하면 된다.
솔루션을 구축해 나가면서 문제를 더 잘 이해하게 되며, 최상의 솔루션이 처음의 생각과는 다르다는 사실을 깨닫게 된다.

작업 중심 방식으로 이동하여 얻는 장점은 설계가 단순해진다는 점이다.

사전설계를 한다면 항상 발생하지도 않을 수정에 대한 비용을 걱정하느라 과도하게 유연한 솔루션을 갖추려하게 되고 이는 복잡성을 유발한다.

### 리팩토링과 성능

소프트웨어를 이해하기 쉽게 만드려면 수정할 일이 많은데, 그러한 수정으로 프로그램이 느려질수도 있다.

빠른 소프트웨어를 작성할 수 있는 일반적인 세 가지 방법이 있다

- 실시간 시스템에 주로 사용되는 시간분배
  - 설계를 분해하면서 각 구성요소에 시간이나 메모리 같은 자원별 예산을 할당하는 것
  - 컴포넌트간 할당된 시간의 상호교환은 혀용하나 예산을 초과하면 안된다
  - 이 방식은 철저히 성능 시간을 중시한다

- 성능에 관심을 가지는 것이다
  - 개발자에게 성능을 유지해야 함을 인식시키는 방법이다
  - 실제로 효과는 없는 편이다, 성능은 크게 오르지 않고 개발속도는 느려진다

- 성능에 관심을 가지는 방식의 문제점 개선
  - 개발 절차중 후기 단계에 있는 성능 최적화 전까지는 성능에 신경쓰지 않고 잘 쪼개진 방식의 프로그램을 제작하는 것
  - 성능 최적화 단계에서는 그 프로그램을 튜닝하는 절차를 따른다

## Code Smell

리팩토링이 필요하다고 판단되는 '의심되는 상황'이 code smell이다.
다음과 같은 상황에서 code smell이 존재한다고 볼 수 있다.

### 중복코드(Duplicated code)

가장 빈번하게 볼 수 있는 code smell이다

- 한 클래스의 두 메소드 안에 겹치는 코드가 들어있는 경우: **메소드 추출**
- 한 클래스의 두 하위클래스에 같은 코드가 들어있는 경우: **메소드 추출** & **메소드 상향**
- 한 클래스의 두 하위클래스에 비슷한 코드가 들어있는 경우: **메소드 추출** & **템플릿 메소드 형성**
- 두 메소드가 알고리즘만 다르고 기능이 같다면: **알고리즘 전환**
- 서로 상관없는 두 클래스 안에 중복코드 존재: **클래스 추출** or **모듈 추출**

### 장황한 메소드(Long method)

- 주석을 달아야 할 것 같은 부분에 주석을 넣는 대신 메소드를 작성한다
  - 메소드명은 기능 수행 방식이 아니라 목적을 나타내는 이름을 정한다
  - 메소드 호출이 원래 코드보다 길어지는 한이 있어도 메소드명은 그 코드의 의도를 잘 반영하는 것으로 정해야 한다

- 메소드를 줄이려면 대부분 **메소드 추출** 기법을 적용해야 한다
  - 메소드에서 하나로 묶으면 좋을 만한 부분을 찾아내어 메소드로 만드는 것
  - 메소드에 매개변수와 임시변수가 많으면 **메소드 추출**이 까다롭다
  - **임시변수를 메소드 호출로 전환**하거나 **임시변수를 메소드 체인으로 전환**하면 임시변수를 제거할 수 있다
  - 매개변수는 **매개변수 세트를 객체로 전환**하거나 **객체를 통째로 전달**하면 간결해진다
  - 리팩토링을 적용했어도 임시변수와 매개변수가 너무 많다면 **메소드를 메소드 객체로 전환**하면 된다

- 조건문과 루프문도 메소드로 빼야 한다
  - 조건문을 추출하려면 **조건문 쪼개기** 기법을 사용한다
  - **루프를 컬렉션 클로저 메소드로 전환**한 후, **클로저 메소드 호출과 클로저 자체에 메소드 추출**을 실시한다

### 거대한 클래스(Large class)

기능이 지나치게 많은 클래스에는 보통 엄청난 수의 인스턴스 변수가 들어있다.
클래스에 인스턴스 변수가 너무 많으면 중복 코드가 반드시 존재하기 마련이다.

- 인스턴스 변수가 너무 많은 클래스: **클래스 추출**
  - 서로 연관된 인스턴스 변수를 콜라서 클래스로 묶는다
  - 클래스 안의 일부 변수의 접두/접미사가 같다면 하나의 클래스로 추출할 것을 고려해봐야 한다
  - 하위클래스로 추출하는 것이 적합할 것 같다면 하위클래스 추출을 실시할 수도 있다
  - 추출할 클래스가 대리자로 부적절할 것 같다면 모듈 추출을 실시한다

- 코드 분량이 거대한 클래스
  - 중복코드가 존재할 가능성이 크며, **클래스 추출**, **모듈 추출**, **하위클래스 추출**, **인터페이스 추출**등을 실시한다

- 방대한 클래스가 GUI 클래스인경우
  - 데이터와 기능을 서로 다른 도메인 객체로 옮겨야 할 수 있다
  - 데이터와 기능을 분리하기 위해 일부 중복 데이터는 놔두고 싱크를 유지해야 할 수 있다
  - 이는 **관측 데이터 복제 기법**을 사용하면 된다

### 과다한 매개변수(Long parameter list)

매개변수 세트가 길면 서로 일관성이 없어지거나 사용이 불편해지고 지속적인 수정이 필요하다

- 객체지향이 아니라면
  - 전역 데이터를 사용하지 않기 위해서 루틴에 필요한 모든 것을 매개변수로 전달한다

- 객체를 사용할 때는
  - 메소드에 필요한 모든 데이터를 전달하는 것이 아니라 필요한 데이터를 가져올 수 있는 메소드만 전달하면 된다
  - 메소드가 필요로 하는 각종 데이터는 그 메소드가 속한 클래스에 들어있다

- **매개변수 세트를 메소드로 전환**
  - 이미 알고 있는 객체(인스턴스 변수/다른 매개변수)에 요청하여 한 매개변수에 들어 있는 데이터를 가져올 수 있을 때

- **객체를 통째로 전달**
  - 객체에 있는 데이터 세트를 가져온 후, 데이터 세트를 그 객체 자체로 전환할 때

- **매개변수 세트를 객체로 전환**
  - 여러 데이터 항목에 논리적 객체가 없다면

- 호출되는 객체가 호출 객체에 의존적일때
  - 데이터를 개별적으로 빼서 매개변수로 전달하는 것이 바람직함

### 수정의 산발(Divergent change)

- 한 클래스가 다양한 원인 때문에 다양한 방식으로 자주 수정될 때
  - 클래스를 여러 개의 변형 객체로 분리하는 것이 좋다
  - 분리를 통해 각 객체는 한 종류의 수정에 의해서만 변경된다

- 특정 원인으로 인해 변하는 모든 부분을 찾은 후 **클래스 추출**을 적용한다

### 기능의 산재(Shotgun surgery)

- 기능의 산재는 수정의 산발과 반대의 상황이다

- 수정할 때마다 여러 클래스에서 수많은 자잘한 부분을 고쳐야 할 때
  - **메소드 이동**과 **필드 이동**을 적용하여 수정할 부분을 합친다
  - 기존의 클래스 중 적절한 병합 대상이 없다면 새 클래스를 만들어야 한다
  - **클래스 내용 직접 삽입**을 적용하여 별도 클래스에 분산되어 있던 모든기능을 한 곳으로 가져와도 된다

### 잘못된 소속(Feature envy)

- 객체의 핵심은 데이터와 그 데이터에 사용되는 프로세스를 한 데 묶는 기술이란 점이다

- 어떤 메소드가 자신이 속하지 않은 클래스에 더 많이 접근한다면 잘못된 소속의 문제가 발생한 것이다
  - 다른 객체에 있는 여러 읽기 메소드를 호출한다면 **메소드 이동** 기법을 실시할 수 있다
  - 메소드의 일부만 소속이 잘못되었다면 **메소드 추출**을 적용한 뒤 **메소드 이동**을 적용할 수 있다

- 한 메소드가 여러 클래스에 들어있는 기능을 이용할 때
  - 문제의 메소드가 접근하는 데이터가 어느 클래스에 제일 많이 들어있는지 파악하여 해당 클래스로 옮긴다
  - 그 전에 **메소드 추출**을 통해 메소드를 다른 클래스로 옮길 여러 부분으로 쪼개면 더 작업이 쉬워진다

- strategy pattern, visitor pattern, self delegation pattern
  - 기본 규칙은 함께 수정되는 것은 하나로 뭉치는 것이다
  - 해당 패턴들은 재정의가 필요한 일부기능을 따로 빼내서 기능을 수정하기 쉽게하는 대신 인다이렉션이 늘어난다

### 데이터 뭉치(Data clumps)

- 몰려있는 데이터 뭉치는 객체로 만들어야 한다
  - 두 클래스에 들어 있는 인스턴스 변수나 여러 메소드 시그니처에 들어있는 매개변수
  - 데이터 뭉치를 객체로 전환하려면 해당 필드를 대상으로 **클래스 추출** 기법을 적용한다
  - 그리고 나서 메소드 시그니처를 대상으로 **매개변수 세트를 객체로 전환** 기법과 **객체를 통째로 전달** 기법을 적용한다

- 새로생긴 객체의 속성들 중 일부만 이용하는 데이터뭉치라 해도 효과가 있다
  - 둘 이상의 필드를 객체로 전환하면 코드 개선이 가능하다
  - 여러 데이터 값 중 하나를 삭제했을 때 나머지 데이터 값들이 제대로 돌아가야 한다
  - 객체로 전환하고 나면 전체적인 성능이 개선될 여지도 있다

### 강박적 기본 타입 사용(Primitive obsession)

프로그래밍 환경의 데이터는 기본타입과 레코드 타입이 있다. 레코드에는 항상 일정 양의 오버헤드가 따른다.

- 객체는 언어에 내장된 기본 타입과 구별하기 힘든 작은 클래스를 손쉽게 작성할 수 있다

- 데이터 값을 객체로 전환
  - 객체에 익숙하지 않다면 전화번호와 우편번호와 같은 특수 문자열 클래스 등의 사소한 작업에 작은 객체를 잘 사용하지 않으려는 경향이 있다
  - 데이터 값이 분류 부호일 때는 그 값이 기능에 영향을 주지 않는다면 **분류부호를 열거 클래스로 전환**하자
  - 조건문에 분류 부호가 사용될 때는 **분류부호를 하위클래스로 전환**하거나 **분류부호를 상태/전략패턴으로 전환** 기법을 적용하자

### Switch 문(Switch statement)

- 객체지향 코드에서는 switch-case 문이 비교적 적게 사용된다
  - switch 문의 단점은 반드시 중복이 생긴다는 점이다
  - 이를 해결하기 위한 방법은 객체지향의 다형성을 이용하는 것이다

- 대부분의 switch문은 재정의로 바꿔야 한다
  - switch문에 사용되는 분류부호 값이 들어있는 메소드나 클래스가 있어야 한다
  - **메소드 추출**을 실시해서 switch문을 메소드로 빼낸 후 **메소드 이동**을 실시해서 클래스를 옮긴다
  - **동시에 분류 부호를 하위클래스로 전환**하거나 **분류부호를 상태/전략 패턴으로 전환**한다
  - 상속구조를 만들었다면 **조건문을 재정의로 전환**한다

- 하나의 메소드에 영향을 미치는 case문 수가 적다면
  - 모든 case문을 수정할일이 없을 것 같으면 재정의로 전환하는것은 과하다
  - 그런경우 **매개변수를 메소드로 변환**하는 편이 낫다
  - 여러 case문 중 하나가 null인 경우 **Null 검사를 널 객체에 위임**해야 한다

### 평행 상속 계층(Parallel inheritance hierarchies)

- 평행 상속 계층은 기능의 산재의 특수한 상황임
  - 이 문제가 있다면 한 클래스의 하위 클래스를 만들 때마다 다른 클래스의 하위클래스도 만들어야 한다
  - 서로 다른 두 상속 계층의 클래스명 접두어가 같으면 이 문제를 의심할 수 있다

- 보통 한 상속계층의 인스턴스가 다른 상속 계층의 인스턴스를 참조하게 만들면 된다
  - **메소드 이동**과 **필드 이동**을 실시하면 참조하는 클래스에 있는 계층이 제거된다

### 직무유기 클래스(Lazy class)

- 하나의 클래스를 작성할 때 마다 유지관리 비용이 추가된다
  - 비용대비 효율이 떨어지는 하위클래스나 모듈은 **계층 병합**을 실시하면 된다
  - 거의 쓸모 없는 구성요소에는 **클래스 내용 직접 삽입**이나 **모듈 내용 직접 삽입** 기법을 적용해야 한다

### 막연한 범용 코드(Speculative generality)

- 조만간 필요할 것이란 생각에 아직은 필요없는 기능을 수행하고자 갖은 호출과 조건문을 넣으려는 경우
  - 모든 경우가 활용될 경우 문제가 없겠지만, 그렇지 않다면 제거해야 한다
  - 별다른 기능이 없는 클래스나 모듈이 있다면 **계층 병합**을 실시
  - 불필요한 위임을 제거하려면 **클래스 내용 직접 삽입**을 실시
  - 메소드에 사용되지 않는 매개변수가 있으면 **매개변수 제거**를 실시
  - 메소드명이 이상하다면 **메소드명 변경** 실시

- 메소드나 클래스가 오직 테스트케이스에서만 사용된다면 유력한 용의자로 막연한 범용코드를 지목할 수 있다
  - 해당 메소드나 클래스를 발견하면 그것과 그것을 실행하는 테스트케이스를 모두 삭제하자
  - 적절한 기능을 실행하는 테스트케이스용 헬퍼 메소드나 클래스는 삭제하지 않아도 된다

### 임시 필드(Temporary field)

- 어떤 객체 안에 인스턴스 변수가 특정 상황에서만 할당되는 경우가 간혹 있다
  - 임시 필드는 복잡한 알고리즘에 여러 변수를 사용해야 할때 자주 발생한다
  - 많은 매개변수를 전달하는 것을 꺼린 나머지 매개변수를 필드에 대입한다
  - 이럴 때는 인스턴스 변수와 그 변수를 사용하는 메소드 전부에 대해 **클래스 추출**을 적용한다

### 메시지 체인(Message chains)

- 메시지 체인은 클라이언트가 한 객체에 다른 객체를 요청하면 다른 객체가 또다른 객체를 요청하는 식의 연쇄 요청을 말한다
  - 메시지 체인은 수많은 코드행이 든 `getThis` 메소드나 임시변수 세트라고 봐도 된다
  - 호출 사이의 관계가 수정될 때마다 코드들도 모두 수정해야 한다

- 이런 경우 **대리객체 은폐**를 실시한다
  - 원칙적으로 체인을 구성하는 모든 객체에 적용할 수 있지만 과잉 중개 메소드 문제가 발생할 수 있다
  - 결과 객체가 어느 대상에 사용되는지 알아내서 객체가 사용되는 코드 부분을 **메소드 추출**한다
  - 추출한 메소드를 **메소드 이동**을 통해 체인 아래로 밀어낼 수 있는지 확인한다

### 과잉 중개 메소드(Middle Man)

- 캡슐화에는 대개 위임이 수반된다
  - 위임이 지나치면 문제가 된다. 어떤 클래스의 인터페이스의 절반이상이 다른 클래스에 위임을 하고 있다면 문제이다
  - **과잉 중개 메소드 제거**를 실시하여 원리가 구현된 객체에 직접 접근할 수 있다
  - 일부 메소드에 별 기능이 없다면 **메소드 내용 직접 삽입**을 한다
  - 부수적인 기능이 있다면 **위임을 상속으로 전환** 기법을 실시한다

### 지나친 관여(Inappropriate intimacy)

- 간혹 클래스끼리 관계가 지나치게 밀접한 경우
  - 밀접한 클래스는 서로의 private을 알아내는데 지나친 노력을 기울일 때가 있다
  - 서로 지나치게 관여하는 클래스는 **메소드 이동**과 **필드 이동**으로 분리해야 한다

- **클래스의 양방향 연결을 단방향으로 전환** 기법을 적용할 수 있으면
  - 해당 클래스들이 공통으로 필요로하는 부분이 있따면 **클래스 추출**을 실시하여 별도의 클래스로 빼낸다
  - 혹은 **대리 객체 은폐**를 실시하여 다른 클래스가 중개 메소드 역할을 하게 한다

- 상속으로 인해 지나친 관여가 발생하는 경우가 많다
  - 하위 클래스는 항상 상위 클래스가 공개하는 것보다 많은 데이터를 필요로 한다
  - 상위 클래스에서 하위 클래스를 빼내는 경우 **상속을 위임으로 전환** 기법을 사용한다

### 인터페이스가 다른 대용 클래스(Alternative classes with different interfaces)

- 기능은 같은데 시그니처가 다른 메소드
  - **메소드명 변경**을 실시해야 한다
  - 클래스에 충분한 기능이 구현되어 있지 않기 때문에 프로토콜이 같아질 때까지 **메소드 이동**을 실시해야 한다
  - 코드를 너무 여러 번 옮겨야 한다면 **상위클래스 추출**을 실시한다

### 미흡한 라이브러리 클래스(Incomplete library class)

- 재사용을 객체의 목적이라 생각하는 것은 재사용을 과대평가한 것

- 라이브러리 제작자라 해도 모든 설계를 파악한다는 것은 거의 불가능하다
  - 라이브러리 클래스에게 원하는 기능을 수정하도록 수정하는 것이 보통은 불가능하다
  - 이 경우 메소드 이동같은 방법이 무용지물이 된다
  - 라이브러리 클래스에 넣어야 할 메소드가 적다면 외래 클래스에 **메소드 추가** 기법을 사용한다
  - 추가할 기능이 많다면 **국소적 상속확장 클래스 사용**기법을 사용한다

### 데이터 클래스(Data class)

- 데이터 클래스는 필드와 필드 읽기/쓰기 메소드만 들어있는 클래스다
  - 해당 클래스는 데이터 보관만 담당한다
  - 구체적인 데이터 조작은 다른 클래스가 수행한다
  - 필드는 **필드 캡슐화** 기법을 실시한다
  - 컬렉션 필드는 **컬렉션 캡슐화** 기법을 적용한다
  - 변경되지 않아야 하는 필드에는 **쓰기 메소드 제거**를 적용한다

- 읽기/쓰기 메소드가 다른 클래스에 의해 사용되는 부분을 찾는다
  - 다른 클래스의 해당 부분을 데이터 클래스로 **메소드 이동**을 실시한다
  - 메소드 전체를 옮길 수 없다면 **메소드 추출**을 실시하고, 읽기/쓰기 메소드에 **메소드 은폐**를 적용한다

### 방치된 상속물(Refused Bequest)

- 하위클래스가 상속받은 부모 클래스의 메소드와 데이터중 더이상 쓰이지 않거나 필요없는 경우

- 잘못된 계층구조로 인하여 발생한 경우에는
  - 새로운 대등 클래스를 작성한다
  - **메소드 하향**과 **필드 하향**을 실시해서 사용되지 않는 메소드를 대등한 클래스에 몰아넣어야 한다

- 방치된 상속물을 반드시 처리해야 하는 것은 아니다
  - 일부기능을 재사용하고자 하위 클래스로 몰아넣는 작업은 효과적일 수 있다
  - 이 문제는 심각하지 않은 경우가 대부분이므로 리팩토링이 별로 필요하지 않을 수 있다

- 하위클래스가 기능은 재사용하지만 상위클래스의 인터페이스를 지원하지 않을 때
  - 상속 구현을 거부할 수는 있지만 인터페이스를 거부하는 것은 문제임
  - **상속을 위임으로 전환** 기법을 적용해서 계층 구조를 없애야 한다

### 불필요한 주석(Comments)

- 어떤 코드 구간의 기능을 설명할 주석이 필요한 경우 **메소드 추출** 실시
- 메소드가 추출된 상태임에도 기능을 설명할 주석이 필요하다면 **메소드명 변경**을 실시
- 시스템에 필수적인 상태에 관한 규칙을 설명해야 할때는 **Assertion 넣기**를 실시
- 주석은 무슨 작업을 해야 좋을지 모를때만 넣는것이 좋다

## 테스트작성

리팩토링을 실시하기 위한 필수 전제조건은 반드시 견고한 테스트를 해야 한다는 것이다.

### 자가 테스트 코드의 가치

자가 테스트 코드를 작성하지 않으면 작업을 파악하거나 설계하는 시간에 비해 디버깅 시간이 길었다.
일반적으로 테스트를 작성하면 훨씬 시간이 절약된다.

테스트를 작성하기 가장 적합한 시점 중 하나는 프로그래밍을 시작할 때다.

기능을 추가해야 할 때는 우선 테스트부터 작성하자.
테스트를 작성하면 그 기능을 추가하려고 해야 할 작업이 무엇인지 자문하게 된다.
그리고 테스트를 작성하면 구현부가 아니라 인터페이스에 집중하게 된다.

### 테스트 추가

모든 public 메소드를 테스트 할 수도 있지만, 위험 위주로 테스트를 작성할 수도 있다.

현재나 미래에 버그를 찾을 곳에서 테스트를 하며, 읽고 쓰는 기능뿐인 메소드는 테스트 하지 않을 수 있다.

너무 많은 테스트를 작성하려다 보면 오히려 질려서 테스트를 필요한 만크도 작성하지 못하게 되므로,
버그 가능성이 거의 없는 부분은 애시장초 테스트 작성 대상에서 제외시키는게 나을 수 있다.

테스트에서 가장 힘든 일은 경계 조건을 찾는 것이다.
경계찾기에는 테스트를 실패하게 할 가능성이 있는 특수조건을 찾는 작업도 포함이된다.

테스트를 실시할 때는 반드시 예상한 에러가 제대로 발생하는지 검사해야 한다.

## 메소드 정리

리팩토링의 주된 작업은 메소드를 적절히 정리하는 것이다. 거의 모든 문제점은 장황한 메소드로 인해 생긴다.

### 메소드 추출

> Extract Method

```java
void printOwing(double amount) {
  printBanner();

  System.out.println("name" + _name);
  System.out.println("amount" + amount);
}
```

메소드 추출을 적용하면

```java
void printOwing(double amount) {
  printBanner();
  printDetails(amount);
}

void printDetails(double amount) {
  System.out.println("name" + _name);
  System.out.println("amount" + amount);
}
```

#### 메소드 추출: 동기

- 메소드가 너무 길거나 코드에 주석을 달아야만 의도를 이해할 수 있을 때
  - 메소드가 적절히 잘게 쪼개져 있으면 다른 메소드에서 쉽게 사용할 수 있다
  - 상위계층의 메소드에서 메소드 이름으로 더 많은 정보를 읽어들일 수 있다
  - 재정의하기도 훨씬 수월하다

- 메소드 내용이 간결한 것도 중요하지만, 메소드 이름도 잘 지어야 한다
  - 메소드 추출로 코드의 명료성이 향상된다면 메소드명이 추출한 코드보다 길어도 추출을 실시해야 한다

#### 메소드 추출: 방법

- 기존 메소드에서 추출할 코드를 새로운 메소드로 이동시킨다
- 옮긴 코드에서 기존 메소드의 모든 지역변수 참조를 찾아, 새로 생성한 메소드의 지역변수나 매개변수로 만든다
- 추출 코드에 의해 변경되는 지역변수를 파악하여 처리한다
- 생성한 코드에서 사용하는 지역변수를 대상 메소드에 매개변수로 전달한다
- 둘 이상의 변수를 반환해야 할 때
  - 최선의 방법은 각기 다른 값을 하나씩 반환하는 여러개의 메소드를 만드는 방법이다.
  - 자신이 사용하는 언어가 출력 매개변수 기능이 있다면 출력 매개변수를 사용해도 된다.
- 테스트를 실시한다

### 메소드 내용 직접 삽입

> Inline Method

```java
int getRating() {
  return (moreThanFiveLateDeliveries()) ? 2 : 1;
}

boolean moreThanFiveLateDeliveries() {
  return _numberOfLateDeliveries > 5;
}
```

메소드 내용 직접 삽입을 적용하면

```java
int getRating() {
  return (_numberOfLateDeliveries > 5) ? 2 : 1;
}
```

#### 메소드 내용 직접 삽입: 동기

- 간혹 메소드명에 모든 기능이 반영될 정도로 지나치게 단순한 경우 해당 메소드를 없애야 한다
- 잘못 쪼개진 메소드의 내용을 하나의 큰 메소드에 직접 삽입 후, 합쳐진 메소드를 다시 각각 작은 메소드로 추출한다
- 과다한 인다이렉션과 동시에 모든 메소드가 다른 메소드에 단순 위임을 하고 있어 코드가 지나치게 복잡할 때 실시한다

#### 메소드 내용 직접 삽입: 방법

- 메소드가 재정의되어 있지 않은지 확인한다(재정의 되어 있다면 실시X)
- 해당 메소드를 호출하는 부분을 모두 찾는다
- 각 호출 부분을 메소드 내용으로 교체한다
- 테스트를 실시한다

### 임시변수 내용 직접 삽입

> Inline Temp

```java
double basePrice = anOrder.basePrice();
return (basePrice > 1000);
```

임시변수 내용 직접 삽입을 적용하면

```java
return (anOrder.basePrice() > 1000);
```

#### 임시변수 내용 직접 삽입: 동기

- 임시변수 내용 직접 삽입은 임시변수를 메소드 호출로 전환을 실시하는 도중에 병용하게 되는 경우가 많음
- 메소드 호출의 결과가 임시변수에 대입될 때, 임시변수가 다른 리팩토링에 방해가 되면 내용 직접삽입을 한다

#### 임시변수 내용 직접 삽입: 방법

- 대입문의 우변에 문제가 없는지 확인
- 문제가 없다면 임시변수를 `final`로 선언하고 확인해보자
- 임시변수를 참조하는 모든 부분을 찾아서 대입문 우변의 수식으로 바꾸자
- 하나씩 수정할 때마다 테스트를 실시한다

### 임시변수를 메소드 호출로 전환(Replace Temp with Query)

```java
double getPrice() {
  double basePrice = _quantity * _itemPrice;
  double discountFactor;
  if (basePrice > 1000) {
    discountFactor = 0.95;
  } else {
    discountFactor = 0.98;
  }
  return basePrice * discountFactor;
}
```

임시변수를 메소드 호출로 전환을 적용하면

```java
double getPrice() {
  return basePrice() * discountFactor();
}
// ...

double basePrice() {
  return _quantity * _itemPrice;
}

private double discountFactor() {
  if (basePrice() > 1000) return 0.95;
  else return 0.98;
}
```

#### 임시변수를 메소드 호출로 전환: 동기

- 임시변수는 일시적이고 국소 범위로 제한된다
- 임시변수를 메소드 호출로 수정하면 클래스 안 모든 메소드가 그 정보에 접근할 수 있다
- 임시변수를 메소드 호출로 전환은 대부분의 경우 메소드 추출을 적용하기 전에 반드시 적용해야 한다
- 지역변수가 많을 수록 메소드 추출이 힘들어진다

#### 임시변수를 메소드 호출로 전환: 방법

- 값이 한번만 대입되는 임시변수를 찾는다
- 값이 여러번 대입되는 임시변수는 임시변수 분리 기법을 실시한다
- 임시변수를 final로 선언한다
- 대입문 우변을 분리하여 private 메소드로 만든다
- 만약 메소드가 객체변경등을 한다면 상태 변경메소드와 값 반환 메소드를 분리한다
- 임시변수를 대상으로 임시변수 내용 직접 삽입을 실시한다
- 테스트를 실시한다

### 직관적 임시변수 사용

> Introduce Explaining Variable

```java
if ((platform.toUpeerCase().indexOf("MAC") > -1) &&
    (browser.toUpperCase().indexOf("IE") > -1) &&
    wasInitialized() && resize > 0) {
  //...
}
```

직관적 임시변수 사용 적용

```java
final boolean isMacOs = platform.toUpeerCase().indexOf("MAC") > -1);
final boolean isIEBrowser = browser.toUpperCase().indexOf("IE") > -1;
final boolean wasResized = resize > 0;

if (isMacOs && isIEBrowser && wasInitialized() && wasResized) {
  //...
}
```

#### 직관적 임시변수 사용: 동기

- 수식이 너무 복잡해져 이해하기 어려운 경우 임시변수를 사용하여 수식을 쪼갤 수 있다
- 직관적 임시변수 사용은 조건문에서 각 조건절을 가져와서 직관적 이름의 임시변수로 의미를 설명할 때 사용한다
- 직관적 임시변수 사용보다는 메소드추출을 우선 고려하는 것이 좋다
  - 임시변수는 하나의 지역내에서만 사용할 수 있기 때문이다
  - 하지만 지역변수로 인해 메소드추출을 적용하기 힘들다면 임시변수 사용 기법을 활용한다
- 나중에 코드나 로직의 복잡함이 덜해지면 임시변수를 메소드 호출로 전환 기법을 적용한다

#### 직관적 임시변수 사용: 방법

- 임시변수를 final로 선언하고 복잡한 수식에서 한 부분의 결과를 그 임시변수에 대입
- 수식에서 한 부분의 결과를 임시변수의 값으로 교체한다(여러개가 있어도 순차적으로 진행한다)
- 테스트를 실시한다

### 임시변수 분리

> Split Temporary Variable

```java
double temp = 2 * (_height + _width);
System.out.println(temp);
temp = _height * _width;
System.out.println(temp);
```

임시변수 분리 적용

```java
final double perimeter = 2 * (_height + _width);
System.out.println(perimeter);
final double area = _height * _width;
System.out.println(area);
```

#### 임시변수 분리: 동기

- 임시변수를 사용하다보면 임시변수에 값이 여러번 대입될 경우가 있다
- 많은 임시변수는 긴 코드의 계산 결과를 나중에 간편히 참조할 수 있게 하는 용도로 사용된다
- 이런 변수에는 값이 한 번만 대입되어야 한다
  - 값이 두 번 이상 대입된다는 것은 변수가 메소드 안에서 여러 용도로 사용된다는 반증이다
  - 임시변수 하나를 여러 용도로 사용하면 코드분석시 혼란을 줄 수 있다

#### 임시변수 분리: 방법

- 선언문과 첫 번째 대입문에 있는 임시변수 이름을 변경한다(값 누적용 임시변수는 분리하지 않는다)
- 이름을 바꾼 새 임시변수를 final로 선언한다
- 임시변수의 다음 대입문 전 범위의 임시변수 참조를 모두 수정한다
- 두 번째 대입문에 있는 임시변수를 선언한다
- 임시변수마다 차례대로 수정한다
- 테스트를 실시한다

### 매개변수로의 값 대입 제거

> Remove Assignments to Parameters

```java
int discount(int inputVal, int quantity, int yearToDate) {
  if (inputVal > 50) inputVal -= 2;
  // ...
}
```

매개변수로의 값 대입 제거 적용

```java
int discount(int inputVal, int quantity, int yearToDate) {
  int result = inputVal;
  if (result > 50) result -= 2;
  // ...
}
```

#### 매개변수로의 값 대입 제거: 동기

- 전달받은 매개변수에 다른 객체 참조를 대입할 때의 문제
  - 코드의 명료성이 떨어진다
  - 값을 통한 전달과 참조를 통한 전달을 혼동하게 된다 (자바는 값을 통한 전달만 사용한다)

- 매개변수로의 값 대입 제거 규칙은 출력 매개변수가 사용되는 다른언어에서는 사용하지 않아도 된다
  - 그러나 출력 매개변수는 가능하면 적게 사용하는 것이 좋다

#### 매개변수로의 값 대입 제거: 방법

- 매개변수 대신 사용할 임시변수를 선언한다
- 매개변수로 값을 대입하는 코드 뒤에 나오는 매개변수 참조를 모두 임시변수로 수정한다
- 매개변수로의 값 대입을 임시변수로의 값 대입으로 수정한다
- 테스트를 실시한다

### 메소드를 메소드 객체로 전환

> Replace Method with Method Object

```java
class Account {
  int gamma(int inputVal, int quantity, int yearToDate) {
    int importantValue1 = (inputVal * quantity) + delta();
    int importantValue2 = (inputVal * yearToDate) + 100;
    if ((yearToDate - importantValue1) > 100) {
      importantValue2 -= 20;
    }
    int importantValue3 = importantValue2 * 7;
    return importantValue3 - 2 * importantValue1;
  }
}
```

메소드를 메소드 객체로 전환 적용

```java
class Gamma {
  private final Account account;
  private int inputVal;
  private int quantity;
  private int yearToDate;
  private int importantValue1;
  private int importantValue2;
  private int importantValue3;

  Gamma(Account account, int inputVal, int quantity, int yearToDate) {
    this.account = account;
    this.inputVal = inputVal;
    this.quantity = quantity;
    this.yearToDate = yearToDate;
  }

  int compute() {
    importantValue1 = (inputVal * quantity) + account.delta();
    importantValue2 = (inputVal * yearToDate) + 100;
    // 아래의 계산에 메소드 추출을 적용할 수 있다!!
    if ((yearToDate - importantValue1) > 100) {
      importantValue2 -= 20;
    }
    importantValue3 = importantValue2 * 7;
    return importantValue3 - 2 * importantValue1;
  }
}

class Account {
  int gamma(int inputVal, int quantity, int yearToDate) {
    return new Gamma(this, inputVal, quantity, yearToDate).compute();
  }
}
```

#### 메소드를 메소드 객체로 전환: 동기

- 메소드 분해를 어렵게 많드는 것은 지역변수이다
  - 임시변수를 메소드 호출로 전환하면 어려움이 어느정도 해소된다

- 메소드를 메소드 객체로 전환을 적용하면
  - 모든 지역변수가 메소드 객체의 속성이된다
  - 그 객체에 메소드 추출을 적용해서 원래의 메소드를 쪼개어 여러개의 추가 메소드를 만든다

#### 메소드를 메소드 객체로 전환: 방법

- 전환할 메소드의 이름과 같은 이름으로 새 클래스를 생성한다
- 클래스에 원본 메소드가 들어있던 객체를 나타내는 final 필드를 추가한다
- final 필드는 원본 메소드안의 각 임시변수와 매개변수에 해당하는 속성이다
- 새 클래스에 원본 객체와 각 매개변수를 받는 생성자 메소드를 작성한다
- 새 클래스에 compute라는 이름의 메소드를 작성한다
- 원본 메소드 내용을 compute 메소드 안에 복사해넣는다
- 원본 객체에 있는 메소드 호출시 원본 객체를 나타내는 필드를 사용한다
- 테스트를 실시한다

### 알고리즘 전환

> Substitute Algorithm

#### 알고리즘 전환: 동기

- 기능을 수행하기 위한 비교적 간단한 방법이 있다면 복잡한 방법을 좀 더 간단한 방법으로 교체해야 한다
- 문제를 해결해나가다 더 간단한 방법이 있음을 발견하면 알고리즘을 교체해야 하는 상황이다
- 알고리즘 교체를 위해서는 메소드를 최대한 잘게 쪼개야 한다

#### 알고리즘 전환: 방법

- 교체할 간결한 알고리즘을 준비한다
- 새 알고리즘을 실행하면서 여러 번의 테스트를 실시한다

## 객체 간의 기능 이동

### 메소드 이동

> Move Method

```java
class Account {
  private AccountType type;
  private int daysOverdrawn;

  double overdraftCharge() {
    if (type.isPremium()) {
      double result = 10;
      if (daysOverdrawn > 7) {
        result += (daysOverdrawn - 7) * 0.85;
        return result;
      }
    } else {
      return daysOverdrawn * 1.75;
    }
  }

  double bankCharge() {
    double result = 4.5;
    if (daysOverdrawn > 0) {
      result += overdraftCharge();
      return result;
    }
  }
}
```

객체 간의 기능 이동 적용

```java
class AccountType {
  double overdraftCharge(int daysOverdrawn) {
    if (isPremium()) {
      double result = 10;
      if (daysOverdrawn > 7) {
        result += (daysOverdrawn - 7) * 0.85;
        return result;
      }
    } else {
      return daysOverdrawn * 1.75;
    }
  }
}

class Account {
  private AccountType type;
  private int daysOverdrawn;

  // 메소드는 위임으로 변경한거나 삭제한다
  double overdraftCharge() {
    return type.overdraftCharge(daysOverdrawn);
  }

  double bankCharge() {
    double result = 4.5;
    if (daysOverdrawn > 0) {
      result += overdraftCharge();
      return result;
    }
  }
}
```

#### 메소드 이동: 동기

- 클래스에 기능이 너무 많거나 클래스가 다른 클래스와 과하게 연동되어 의존성이 지나친 경우 메소드를 옮기는 것이 좋다
- 자신이 속한 객체보다 다른 객체를 더 많이 참조하는 것 같은 메소드를 찾는다
  - 그 메소드를 호출하는 메소드, 그 메소드가 호출하는 메소드, 상속계층에서 재정의 메소드를 확인한다
  - 옮기려는 메소드가 더 많이 참조한다고 보이는 메소드가 들어있는 객체를 기준으로 진행여부를 판단한다

#### 메소드 이동: 방법

- 원본 메소드에서 사용한 기능 중 원본 클래스에 정의되어 있는 기능을 확인한다
  - 원본 메소드에서만 사용하는 기능이라면 그 메소드와 함께 옮겨야 한다
  - 그 기능이 다른 메소드에서도 사용된다면 또 다른 메소드도 옮길지를 확인해보자
- 원본 클래스의 하위 클래스와 상위 클래스에서 그 메소드에 대한 다른 선언이 있는지를 검사한다
- 원본 메소드를 대상 클래스 안에 선언한다
- 원본 메소드의 코드를 대상 메소드에 복사한 후 대상 클래스에 맞춰 수정한다
- 원본 객체에서 대상 객체를 참조할 방법을 정한다
- 원본 메소드를 위임 메소드로 전환하거나 삭제한다
- 테스트를 실시한다

### 필드 이동

> Move Field

```java
class Account {
  private AccountType type;
  private double interestRate;

  double interestForAmountDays(double amount, int days) {
    return interestRate * amount * days / 365;
  }
  // ...
}
```

필드 이동 적용

```java
class AccountType {
  private double interestRate;

  void getInterestRate() {
    return interestRate;
  }

  void setInterestRate(double arg) {
    interestRate = arg;
  }
}

class Account {
  private AccountType type;

  double interestForAmountDays(double amount, int days) {
    return type.getInterestRate() * amount * days / 365;
  }
  // ...
}
```

#### 필드 이동: 동기

- 어떤 필드가 자신이 속한 클래스보다 다른 클래스에 있는 메소드를 더 많이 참조할 때
  - 필드를 해당 클래스로 옮기는 것을 고려한다
  - 인터페이스에 따라 메소드를 옮길수도 있지만, 메소드의 위치가 적절하다고 판단되면 필드를 옮긴다
- 클래스 추출을 실시하는 중에서 필드 이동이 수반된다. 이 경우 필드가 판단 기준에서 우선한다.

#### 필드 이동: 방법

- 필드가 public이면 필드 캡슐화 기법을 실시한다
- 대상 클래스 안에 읽기/쓰기 메소드와 함께 필드를 작성한다
- 원본 객체에서 대상 객체를 참조할 방법을 정한다
  - 기존 필드나 메소드에 대상 클래스를 참조하는 기능이 있을 수도 있다
  - 기존 기능이 없다면 얼마나 간편히 그런 기능의 메소드를 작성할 수 있는지 확인한다
- 원본 클래스에서 필드를 삭제한다
- 원본 필드를 참조하는 모든 부분을 대상 클래스에 있는 적절한 메소드를 참조하도록 수정한다
  - 변수 접근 참조는 대상객체의 getter 메소드로 대입 참조 부분은 setter 메소드 호출로 수정한다
- 테스트를 실시한다

### 클래스 추출

> Extract Class

```java
class Person {
  private String name;
  private String officeAreaCode;
  private String officeNumber;

  public String getName() {
    return name;
  }
  
  public String getTelephoneNumber() {
    return "(" + officeAreaCode + ")" + officeNumber;
  }

  String getOfficeAreaCode() {
    return officeAreaCode;
  }

  void setOfficeAreaCode(String arg) {
    officeAreaCode = arg;
  }

  String getOfficeNumber() {
    return officeNumber;
  }

  void setOfficeNumber(String arg) {
    officeNumber = arg;
  }
}
```

클래스 추출 적용

```java
class TelephoneNumber {
  private String number;
  private String areaCode;

  public String getTelephoneNumber() {
    return "(" + areaCode + ")" + number;
  }

  String getAreaCode() {
    return areaCode;
  }

  void setAreaCode(String arg) {
    areaCode = arg;
  }

  String getNumber() {
    return number;
  }

  void setNumber(String arg) {
    number = arg;
  }
}

class Person {
  private String name;
  private TelephoneNumber officeTelephone = new TelephoneNumber();

  public String getName() {
    return name;
  }
  
  public String getTelephoneNumber() {
    return officeTelephone.getTelephoneNumber();
  }

  TelephoneNumber getOfficeTelephone() {
    return officeTelephone;
  }
}
```

#### 클래스 추출: 동기

- 클래스는 확실히 추상화되어야 하며 명확한 기능을 담당해야 한다
- 개발이 진행되며 클래스에 많은 메소드와 데이터가 추가되어 방대해진다
  - 데이터와 메소드가 한 덩어리인 경우
  - 함께 변화하거나 서로 의존적인 데이터

#### 클래스 추출: 방법

- 클래스의 기능 분리 방법을 정한다
- 분리한 기능을 넣을 새 클래스를 작성한다
  - 원본 클래스의 기능이 변했다면 원본 클래스의 이름을 변경한다
- 원본 클래스에서 새 클래스의 링크를 만든다
  - 필요할때까지는 양방향 링크(역방향 링크 추가)를 만들지 않는다
- 옮길 필드마다 필드이동을 실시한다
- 필드 이동마다 테스트를 실시한다
- 메소드 이동을 실시하여 원본 클래스의 메소드를 새 클래스로 옮긴다
  - 피호출 메소드부터 시작해서 호출 메소드 순으로 적용한다
- 메소드 이동을 실시할 때마다 테스트를 실시한다
- 각 클래스를 다시 검사하여 인터페이스를 줄인다
  - 양방향 링크가 있다면 단방향으로 바꿀수 있는지 확인한다
- 여러 곳에서 클래스에 접근할 수 있도록 할지 결정한다

### 클래스 내용 직접 삽입

> Inline Class

클래스 추출의 역순으로 수행한다

#### 클래스 내용 직접 삽입: 동기

- 클래스 내용 직접 삽입은 클래스 추출과 반대이다
- 클래스의 기능 대부분을 다른곳으로 옮기는 리팩토링을 실시하여 남은 기능이 거의 없을 때
- 작은 클래스를 가장 많이 사용하는 다른 클래스에 병합한다

#### 클래스 내용 직접 삽입: 방법

- 원본 클래스의 public 프로토콜 메소드를 합칠 클래스에 선언하고, 이 메소드를 전부 원본 클래스에 위임한다
  - 원본 클래스의 메소드 대신 별도의 인터페이스가 필요하다고 판단되면 인터페이스 추출을 실시한다
- 원본 클래스의 모든 참조를 합칠 클래스 참조로 수정한다
  - 원본 클래스를 private로 선언하고 패키지 밖의 참조를 삭제한다
- 테스트를 실시한다
- 메소드 이동과 필드 이동을 실시해서 원본 클래스의 모든 기능을 합칠 클래스로 옮긴다
- 원본 클래스를 삭제한다

### 대리 객체 은폐

> Hide Delegate

```java
class Person {
  private Department department;

  public Department getDepartment() {
    return department;
  }

  public void setDepartment(Department arg) {
    department = arg;
  }
}

class Department {
  private String chargeCode;
  private Person manager;

  public Department(Person manager) {
    this.manager = manager;
  }

  public Person getManager() {
    return manager;
  }
  // ...
}
```

클라이언트 클래스는 매니저를 알아내려면 우선 부서를 알아야 한다: `manager = kim.getDepartment().getManager();`

이런 의존성을 줄이려면 Department 클래스를 클라이언트가 알 수 없게 감춰야 한다.
그러려면 Person 클래스에 위임 메소드를 작성하면 된다.

대리 객체 은폐 적용하면 다음과 같다

```java
class Person {
  private Department department;

  // ...

  public Person getManager() {
    return department.getManager();
  }
}
```

#### 대리 객체 은폐: 동기

- 객체의 핵심 개념중 하나가 캡슐화이다
- 클라이언트가 서버 객체의 필드 중 하나에 정의된 메소드를 호출할 때 클라이언트는 대리객체에 관해 알아야 한다
  - 대리객체가 변경될 때 클라이언트도 변경해야 할 가능성이 있다
  - 이런 의존성을 없애려면 대리객체를 감추는 위임 메소드를 서버에 두면 된다
  - 이 경우 변경은 서버에만 이루어지고 클라이언트에는 영향을 주지 않는다
- 서버의 일부 클라이언트나 모든 클라이언트에 대리 객체 은폐를 실시하는 것이 좋다
  - 모든 클라이언트를 대상으로 대리 객체를 감출 경우 서버의 인터페이스에서 대리객체에 관한 부분을 삭제해도 된다

#### 대리 객체 은폐: 방법

- 대리 객체에 들어 있는 각 메소드를 대상으로 서버에 간단한 위임 메소드를 작성한다
- 클라이언트를 수정해서 서버를 호출하게 만든다
- 각 메소드를 수정할 때마다 테스트를 실시한다
- 대리 객체를 읽고 써야 할 클라이언트가 하나도 남지 않게 되면 서버에서 대리 객체가 사용하는 읽기/쓰기 메소드를 삭제한다

### 과잉 중개 메소드 제거

> Remove Middle Man

대리 객체 은폐의 역순으로 수행한다

#### 과잉 중개 메소드 제거: 동기

- 대리 객체 은폐를 사용하면 장점을 얻는 대신 단점도 생긴다
  - 클라이언트가 대리 객체의 새 기능을 사용할 때마다 서버에 위임 메소드를 추가해야 한다
- 은폐의 적절한 정도를 알기란 쉽지 않다
  - 대리 객체 은폐와 과잉 중개 메소드 제거를 실시할 때는 은폐의 정도를 몰라도 된다
  - 시스템이 변경되면 은폐 정도의 기준도 변한다

#### 과잉 중개 메소드 제거: 방법

- 대리 객체에 접근하는 메소드를 작성한다
- 서버에서 위임 메소드를 제거하고 클라이언트에서 호출을 대리 메소드 호출로 교체한다
- 메소드를 수정할 때마다 테스트를 실시한다

### 외래 클래스에 메소드 추가

> Introduce Foreign Method

```java
Date newStart = new Date(prevEnd.getYear()), prevEnd.getMonth(), prevEnd.getDate() + 1);
```

외래 클래스에 메소드 추가 적용

```java
Date newStart = nextDay(prevEnd);

// ...

private static Date nextDay(Date arg) {
  // Date 클래스의 외래 메소드
  return new Date(prevEnd.getYear()), prevEnd.getMonth(), prevEnd.getDate() + 1);
}
```

#### 외래 클래스에 메소드 추가: 동기

- 현재 클래스에 없는 한가지 기능이 필요한데 원본 클래스를 수정할 수 없다면
  - 외래 클래스에 메소드로 추가 기법을 사용
- 서버 클래스에 수많은 외래 메소드를 작성해야 하거나 하나의 외래 메소드를 여러 클래스가 사용해야 할 때
  - 국소적 상속확장 클래스 사용 기법을 대신 사용한다
- 외래 메소드는 임시방편에 불과하다
  - 가능하다면 외래 메소드를 원래 있어야 할 위치로 옮겨야 한다

#### 외래 클래스에 메소드 추가: 방법

- 필요한 기능의 메소드를 우선 클라이언트 클래스안에 작성한다
  - 해당 메소드는 클라이언트 클래스의 어느 기능에도 접근하면 안된다
  - 값이 필요할 때는 매개변수로 전달해야 한다
- 서버 클래스의 인스턴스를 찾아 첫 번째 매개변수로 만든다
- 그 메소드에 '서버 클래스의 외래 메소드'와 같은 주석을 추가한다

### 국소적 상속확장 클래스 사용

> Introduce Local Extension

#### 국소적 상속확장 클래스 사용: 동기

- 원본 클래스를 수정하는 것이 불가능 할 경우
  - 필요 메소드가 적다면 외래 클래스에 메소드 추가 기법을 사용할 수 있다
  - 그러나 세 개이상이라면 필요 메소드를 적당한 곳에 모아두어야 한다
  - 이 경우 하위클래스화와 wrapper화를 적용한 국소적 상속확장 클래스를 사용한다
- 국소적 상속확장 클래스는 별도의 클래스지만 상속확장하는 클래스의 하위타입이다
- Wrapper 클래스와 하위 클래스 중 방식을 선택할 수 있다
  - 하위클래스의 작업량이 적다
  - 하위클래스의 문제점은 객체를 생성함과 동시에 하위클래스로 만들어야 한다는 점이다
  - 원본객체가 mutable이면 한 객체를 수정해도 다른객체가 수정되지 않으므로 이때는 wrapper 클래스를 사용해야 한다

#### 국소적 상속확장 클래스 사용: 방법

- 상속확장 클래스를 작성한 후 원본 클래스의 하위클래스나 래퍼클래스로 만든다
- 상속확장 클래스에 변환생성자 메소드를 작성한다
  - 생성자는 원본클래스를 인자로 받는다
- 상속확장 클래스에 새 기능을 추가한다
- 필요한 위치마다 원본 클래스를 상속확장 클래스로 수정한다
- 해당 클래스용으로 정의된 외래 메소드를 모두 상속확장 클래스로 옮긴다

#### 국소적 상속확장 클래스 사용: 예제

Date 클래스를 확장해서 사용하는 경우를 살펴보자

##### 하위 클래스 사용

```java
class MfDateSub extends Date {
  public MfDateSub(String dateString) {
    super(dateString);
  }

  public MfDateSub(Date arg) {
    super(arg.getTime());
  }

  Date nextDay(Date arg) {
    return new Date(prevEnd.getYear()), prevEnd.getMonth(), prevEnd.getDate() + 1);
  }
}
```

##### 래퍼 클래스 사용

```java
class MfDateWrap {
  private Date date;
  
  public MfDateWrap(String dateString) {
    date = new Date(dateString);
  }

  public MfDateWrap(Date arg) {
    date = arg;
  }

  // 원본 Date 클래스의 모든 메소드를 위임
  public int getYear() {
    return date.getYear();
  }

  // 중략 ...

  Date nextDay(Date arg) {
    return new Date(prevEnd.getYear()), prevEnd.getMonth(), prevEnd.getDate() + 1);
  }
}
```

Wrapper 클래스를 사용하는 경우 `equals`와 같은 시스템의 메소드를 재정의하는 것은 위험하다.
이런 경우 `public boolean equalsDate(Date arg)`와 같은 메소드명을 사용할 수 밖에 없다.

## 데이터 체계화

### 필드 자체 캡슐화

> Self Encapsulate Field

```java
class IntRange {
  private int low;

  IntRange(int low) {
    this.low = low;
  }

  boolean includes(int arg) {
    return arg >= low;
  }
}
```

필드 자체 캡슐화 적용

```java
class IntRange {
  private int low;

  IntRange(int low) {
    this.low = low;
  }

  boolean includes(int arg) {
    return arg >= getLow();
  }

  int getLow() {
    return low;
  }

  int setLow(int arg) {
    this.low = arg;
  }
}
```

#### 필드 자체 캡슐화: 동기

- 변수 간접 접근 방식
  - 하위클래스가 메소드에 해당 정보를 가져오는 방식을 재정의 할수 있다
  - 속성 초기화 시점을 바꾸는 등의 데이터 관리가 더 유연해진다
- 변수 직접 접근 방식
  - 코드를 알아보기 쉽다
- 필드 자체 캡슐화를 실시해야 할 시점
  - 상위클래스의 안의 필드에 접근하되 변수 접근을 하위클래스에서 계산된 값으로 재정의 해야 할 때

#### 필드 자체 캡슐화: 방법

- 필드 읽기 메소드와 쓰기 메소드를 작성한다
- 필드를 private으로 만든다
- 필드 참조 부분을 전부 찾아서 일긱 메소드와 쓰기 메소드로 연결한다
- 테스트를 실시한다

### 데이터 값을 객체로 전환

> Replace Data Value with Object

```java
class Order {
  private String customer;

  public Order(String customer) {
    this.customer = customer;
  }
  
  public String getCustomer() {
    return customer;
  }

  public void setCustomer(String arg) {
    this.customer = arg;
  }
}
```

데이터 값을 객체로 전환 적용

```java
class Customer {
  private final String name;

  public Customer(String name) {
    this.name = name;
  }

  public String getName() {
    return name;
  }
}

class Order {
  private Customer customer;

  public Order(String customerName) {
    this.customer = new Customer(customerName);
  }
  
  public String getCustomer() {
    return customer.getName();
  }

  public void setCustomer(String customerName) {
    this.customer = new Customer(customerName);
  }
}
```

#### 데이터 값을 객체로 전환: 동기

- 개발 초기 단계에서는 단순 정보를 간단한 데이터 항목으로 표현하는 사안이 개발 진행과 함께 복잡해진다
  - 한동안은 전화번호를 문자열로 표현해도 되지만 시간이 흐르면 형식화, 지역번호 추출등의 기능이 필요하다
  - 한두 항목은 객체 안에 메소드를 넣어도 되겠지만 금방 Code Smell이 발생하게된다

#### 데이터 값을 객체로 전환: 방법

- 데이터 값을 넣을 클래스를 작성한다
- 클래스에 원본 클래스 안의 값과 같은 타입의 필드를 추가하고 필드를 인자로 받는 생성자와 읽기 메소드를 추가한다
- 원본 클래스의 필드 타입을 새 클래스로 바꾼다
- 원본 클래스 안의 읽기 메소드를 새 클래스의 읽기 메소드를 호출하도록 수정한다
- 필드값 생성은 새 클래스의 생성자를 이용한다
- 테스트를 실시한다

### 값을 참조로 전환

> Change Value to Reference

```java
class Customer {
  private final String name;

  public Customer(String name) {
    this.name = name;
  }

  public String getName() {
    return name;
  }
}

class Order {
  private Customer customer;
  // ...

  public Order(string customerName) {
    cumstomer = new Customer(customerName);
  }

  // ...
}
```

값을 참조로 전환 적용

```java
class Customer {
  private final String name;
  private static Map<Customer> instances = new HashMap<>();

  public static Customer getNamed(String name) {
    return instances.get(name);
  }

  private Customer(String name) {
    this.name = name;
  }

  static void loadCustomers() {
    new Customer("고객1").store();
    new Customer("고객2").store();
  }

  private void store() {
    instances.put(this.getName(), this);
  }

  public String getName() {
    return name;
  }
}

class Order {
  private Customer customer;
  // ...

  public Order(string customerName) {
    cumstomer = Customer.getNamed(customerName)
  }

  // ...
}
```

#### 값을 참조로 전환: 동기

- 참조 객체와 값 객체중 어느 것을 사용할지 결정하기 애매할 때도 있다

#### 값을 참조로 전환: 방법

- 생성자를 팩토리 메소드로 전환한다
- 참조 객체로 접근을 담당할 객체를 정한다
  - 이 기능은 정적 딕셔너리나 레지스트리 객체가 담당할 수 있다
  - 참조 객체로의 접근을 둘 이상의 객체가 담당할 수도 있다
- 객체를 미리생성할지 사용하기 직전에 생성할지를 정한다
- 참조 객체를 반환하도록 팩토리 메소드를 수정한다
  - 객체를 미리생성하는 경우 존재하지 않는 객체 요청에 대한 예외처리 해야함
  - 팩토리 메소드가 원본 객체를 반환함을 알수있도록 메소드명 변경을 적용해야 할 수도 있다
- 테스트를 실시한다

### 참조를 값으로 전환

> Change Reference to Value

```java
class Currency {
  private String code;

  public String getCode() {
    return code;
  }

  private Currency(String code) {
    this.code = code;
  }
}

new Currency("KRW").equals(new Currency("KRW")) // false
```

참조를 값으로 전환

```java
class Currency {
  private String code;

  // ...

  public boolean equals(Object arg) {
    if (!(arg instanceof Currency)) return false;
    Currency other = (Currency) arg;
    return (code.equals(other.code));
  }

  public int hashCode() {
    return code.hashCode();
  }
}

new Currency("KRW").equals(new Currency("KRW")) // true
```

#### 참조를 값으로 전환: 동기

- 참조 객체를 사용한 작업이 복잡해지는 순간 참조를 값으로 바꿔야 할 시점이다
- 값 객체는 변경할 수 없어야 한다는 특성이 있다
  - 값 객체는 분산 시스템이나 병렬 시스템에 주로 사용된다
  - 값이 바뀌는 경우 기존의 객체를 새 객체로 교체해야 한다

#### 참조를 값으로 전환: 방법

- 전환할 객체가 변경불가인지 변경가능인지 확인한다
  - 전환할 객체가 변경불가가 아니면 변경불가가 될 때 까지 쓰기 메소드 제거 실시
- equals / hashCode 메소드를 작성한다
- 팩토리 메소드를 삭제하고 생성자를 public으로 만들어야 좋은지 생각해본다
- 테스트 실시
