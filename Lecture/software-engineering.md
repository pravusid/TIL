# 소프트웨어 공학

## 소프트웨어 공학 개요

### 소프트웨어 정의

> 소프트웨어는 컴퓨터 프로그램, 프로시져 규칙, 관련된 문서와 데이터의 묶음이다.

### 소프트웨어 분류

기능측면 분류

- 시스템 소프트웨어: 컴퓨터를 운영하기 위한 목적의 소프트웨어. 하드웨어를 동작/제어하는 기본기능을 수행하고, 응용소프트웨어의 실행을 위해 여러 컴퓨터 부품들의 협력적 작업을 관리하는 저수준의 프로그램들을 말함.
- 응용소프트웨어: 사용자의 실제 업무를 수행하는 프로그램들로 특정 용도에 사용되도록 만들어진 것.

고객에 따른 분류

- 일반소프트웨어(generic S/W): 요구사항이 일반적/안정적인 불특정 다수를 대상으로 설계된 소프트웨어.
- 맞춤형소프트웨어(customized S/W): 응용도메인, 사용환경 및 요구사항이 특별한 고객을 위해 개발되는 맞춤 소프트웨어

### 소프트웨어의 성질

- 무형의 인공물이며 물질적인 성질과 형태를 가지는 하드웨어와는 다르다
- 소프트웨어는 요구사항에 맞추어 새롭게 만들어진다.
- 소프트웨어 최종 제품의 결과는 추상적이다. 소프트웨어 비용은 설계 과정에 집중된다.
- 소프트웨어 개발비용의 대부분이 노동력에 투입된다.
- 소프트웨어는 상대적으로 변경이 용이하다. 그러나 변경을 위해 필요한 테스트 작업과 요구사항 검증 작업은 쉽지 않다.
- 소프트웨어는 마모되지 않는다. 그러나 환경의 변화나 새로운 요구사항의 등장 또는 기대수준의 향상으로 폐기될 수 있다.
- 소프트웨어 유지보수는 하드웨어와 달리 많은 경우 설계의 변경이 요구된다. 소프트웨어의 변경이 요구될 때 효율적인 형상관리가 필요하다.

### 소프트웨어 공학 정의

신뢰성 있고 요구기능을 효율적으로 수행하는 소프트웨어를 경제적으로 생산하기 위해 건전한 공학적 원리와 방법을 만들고 사용하는 것.

### 좋은 소프트웨어의 기준

- 신뢰도: 오랜시간 작동되며 치명적 오류가 없으며 오류 발생후에도 복구될 수 있다
- 정확성: 요구사항에 비추어 차이가 없음을 나타냄
- 성능: 지정된 시간 안에 시스템에서 처리할 수 있는 작업량
- 사용성: 본래의 설계 목적에 따라 효율성 있게 사용할 수 있는 정도
- 상호운용성: 다른 시스템과 공존하며 협력할 수 있는 능력
- 유지보수성: 새로운 기능추가, 기존기능 개선, 기존기능 수정, 오류수정이 상대적으로 용이함
- 이식성: 소프트웨어가 탑재되는 하드웨어, 운영체제 또는 상호작용하는 다른시스템에서 동작하기 쉬운 정도
- 검사성: 여러 품질요소들을 포함하여 소프트웨어의 속성을 쉽게 검사할 수 있는 경우
- 추적성: 요구사항들, 시스템 설계문서들, 소스코드간에는 구체적인 관계가 존재한다. 그러한 관계정보를 쉽게 추적할 수 있어야 한다.

## 소프트웨어 프로세스

### 주요 소프트웨어 프로세스 활동

1. 소프트웨어 명세: 소프트웨어의 기능과 운영상의 제약 조건을 정한다
2. 소프트웨어 개발: 요구사항 명세를 만족하는 소프트웨어를 설계하고 프로그래밍 한다.
3. 소프트웨어 검증: 소프트웨어가 고객이 원하는 것을 수행하는지 검증한다.
4. 소프트웨어 진화: 고객과 시장의 요구사항 변화를 수용하며 소프트웨어를 수정한다.

소프트웨어 유형에 따라 프로세스가 달라지고, 주요 프로세스 활동들의 서술 내용과 활동들의 구성방법이 달라지며, 서술 내용의 구체성 정도도 다르다.
또한 개별 활동의 결과물이 다르다면 활동들의 순서도 바뀔 것이다.

### 폭포수 모델

폭포수 모델은 1970년대에 유행했던 잘 알려진 전통적 소프트웨어 개발 프로세스 모델이다.
각 단계가 긑날 때 마다 결과를 확인한 후에 다음 단계로 진행해야 한다.
그러나 개발 과정중 문제점이 발생되거나 요구사항이 변경되는 일을 피하기 힘들기 때문에 수정을 위한 재작업을 위해 앞 단계로의 피드백이 필요하다.

#### 폭포수 모델의 단계

- 타당성 조사: 제안된 시스템 개발에 투입되는 비용과 그로 인한 이득을 평가.
  - 타당성 조사
    - 문제정의
    - 기술적인 면과 경제적인 면을 고려한 결정 사항
    - 몇 가지 해결 방안들과 각각의 기대효과
    - 제안된 방안들 각각에 대한 필요한 자원과 비용 및 인도날짜
  - 타당성 평가
    - 조직 측면의 타당성: 제안된 시스템이 조직 목적달성을 지원할 수 있는가
    - 경제적 타당성: 예상되는 비용절약 효과, 기대수익 등이 제안된 시스템을 개발하고 운영하는 데 드는 비용을 상쇄하는가
    - 기술적 타당성: 요구조건을 충족하고 신뢰성있는 개발을 현재의 기술수준으로 할 수 있는가
    - 운영의 타당성: 제안된 시스템을 운영하고 사용할 사람들의 능력과 시스템에 대한 호감도 판단
- 요구 분석과 명세
  - 요구 분석 단계는 소프트웨어가 갖추어야 하는 품질 요소들을 기능, 성능, 사용편의성 및 이식성등의 관점에서 파악하는 것
  - 요구사항 명세서(SRS)
    - 문제에 관한 구체적인 기술
    - 문제 해결을 위한 가능한 대안
    - 소프트웨어 시스템의 기능적 요구사항들
    - 소프트웨어 시스템의 제약사항들
- 설계와 명세
  - 설계단계의 본질은 SRS에 표현된 what을 how로 변환하는 것, 기술된 요구사항들을 구현 작업에 적합하게 명확하고 조직화된 구조의 형태로 바구는 것
  - 전통적 설계: 요구사항을 명세화하기 위한 구조적 분석 방법을 먼저 수행하고 구조적 설계 방법을 적용, 구조적 설계는 아키텍처 설계와 상세설계로 구분된다.
  - 객체지향 설계: 시스템에 필요한 다양한 객체를 파악한 후, 그들간의 관계를 설계한다.
- 코딩과 단위 테스트: 테스트 작업은 각 모듈이 해당 명세서를 만족하는지 확인하는 것이다.
- 통합과 시스템 테스트: 모듈들을 통합하고 테스트하여 전체 시스템이 요구사항을 만족하는지 확인한다.
  - 알파테스트: 소프트웨어 개발 현장에서 수행한다.
  - 베타테스트: 고객의 실제 사용환경에서 수행한다.
- 인도와 유지보수
  - 유지보수 분류
    - 시스템에 남아 있는 오류를 고치기 위한 작업을 수정 유지보수
    - 환경의 변화에 적응하기 위해 수정작업을 하는 것을 적응 유지보수
    - 기능을 개선하거나 성능을 향상시키기 위한 것을 완전 유지보수
    - 잠재적 결함을 찾아 수정함으로써 실제 오류로 연결되지 못하도록 하는 것을 예방 유지보수

#### 폭포수 모델의 장단점

장점

- 순차적인 선형모델이므로 이해하기 쉽다
- 단계가 분리되어 있어 정형화된 접근 방법과 체계적인 문서화 작업이 가능하다
- 각 단계별로 산출물을 체크함으로써 진행상황을 명확하게 알 수 있다

단점

- 프로젝트 초기에 모든 요구사항을 완벽하게 작성하는 것을 불가능하다
- 변경을 수용하기 적합한 형태가 아니다
- 동작이 되는 시스템 버전을 프로젝트 후반부에나 볼 수 있다
- 대형 프로젝트에 적용하기 위해 확장되기 어렵다
- 문서화를 위해 지나친 노력을 소모하는 경향이 있다
- 기본적으로 생명주기를 거슬러 갈 수 없다
- 고객의 요구를 명확히 확인해 줄 방법을 갖지 못한다
- 위험 분석을 수행하지 않는다
- 실수나 오류가 생기면 일정이 지연되거나 소요 자원이 증가한다
- 각 단계의 종료를 위해 정형화된 문서를 요구하므로 모든 절차가 문서 위주로 수행된다.

### 반복 진화형 모델

진화형 모델은 초기버전을 만들고 사용자의 의견을 수렴한 후,
요구사항을 정제하여 새로운 시스템 버전을 개발하는 작업을 반복함으로써 시스템을 완성해가는 방식이다.

#### 반복 진화형 모델의 장단점

장점

요구사항이 부분적으로 완성되지 못한 상태에서도 먼저 시스템의 초기버전을 만들고, 이를 통해 불명확한 요구사항을 명확히 도출해낼 수 있다.

단점

- 요구사항과 개발범위가 분명히 정해지지 않은 상황에서는 개발비용을 예상하기 힘들고, 개발 종료시점이 늦어질 가능성이 크다.
- 공학적관점에서 계속되는 수정작업은 소프트웨어 구조에 악영향을 주어 통합이나 유지보수성에 문제를 야기할 수 있다.

따라서 진화형 모델은 50만 라인 이하의 중소형 시스템에 적합하다.

#### 프로토타이핑 방법

1. 빠른 계획
2. 빠른 설계
3. 프로토타입 만들기
4. 프로토 타입 실행과 피드백
5. 다시 1번 실행 반복

### 점증적 모델

점증적 모델에서 소프트웨어는 여러개의 모듈들로 분해되고 각각은 점증적으로 개발되어 인도된다.
이 모델은 선형 순차 모델을 여러번 적용하여 결과들을 조합하는 것이다.

#### 점증적 모델 구조

1. 요구사항 개요 정의
2. 요구사항들을 점증에 배정
3. 시스템 구조설계
4. 점증을 개발
5. 점증을 확인
6. 점증을 통합
7. 시스템 확인
8. (시스템 미완성) -> 4번으로
9. 최종 시스템

#### 점증적 모델 장단점

장점

- 첫 번째 점증은 가장 중요한 요구사항들을 구현한 것이므로, 사용자는 전체 시스템이 개발될 때 까지 기다리지 않아도 된다
- 소프트웨어 전체가 한꺼번에 릴리즈되는 것보다 시간차를 두고 점증을 개발하여 릴리즈 하는 방식이 사용자 요구사항의 변화에 쉽게 대응할 수 있게 한다.
- 점증들은 점차적으로 규모나 기능이 축소되므로 프로젝트 관리가 어렵지 않다.
- 중요한 부분이 먼저 개발되므로 차후에 개발되는 점증들을 통합하면서 반복적으로 테스트를 수행하게 되어 중요한 부분의 오류가 줄어들게 된다.

단점

- 기능적으로 분해하기 어려운 문제들이 잇으며, 또 중요 기능들이 시스템의 여러 부분으로 나뉘어 배치되는 경우가 있다.
- 요구사항들을 적당한 크기의 점증들로 나누기 쉽지 않다.
- 점증을 구현하기 전에 명확하게 요구사항을 정의하지 못하면 개발 후에도 점증을 수정해야 하는 경우가 있다. 모듈 공통기능을 미리 파악하지 못하는 경우 그러하다.

### 나선형 모델

반복 진화형 모델을 확장한 형태로 전체 생명주기에 걸쳐 일련의 프로토 타이핑과 위험분석을 계획적으로 사용하여 프로젝트 수행 시 발생하는 위험을 관리하고 최소화 하려는 목적을 가진다.

나선형 모델은 새로운 많은 기능들이 실험적으로 처음 시도되는 복잡한 대형 시스템을 개발할 때 적합한 모델로 알려져 있다.

#### 나선형 모델의 과정

타당성 조사, 요구사항 정의, 시스템설계를 순차적으로 진행하고 각 단계를 작업하는 과정에서 대안의 타당성을 평가할 때마다 프로토타이핑과 위험분석을 수행한다.

#### 나선형 모델의 장단점

장점: 큰 위험을 안고 시작해야 하는 대형 프로젝트에서 위험 관리를 통해 위험을 줄임으로써 프로젝트의 성공 가능성을 높이고, 프로젝트나 개발조직에 맞게 변형될 수 있는 융통성을 갖고 있다.

단점: 비교적 새로운 방법으로 아직까지 많이 사용한 경험이 없기 때문에 충분한 검증을 거치지 못했다는 점과 모델 자체가 복잡하나 소프트웨어 개발을 위한 분명한 지침이나 엄격한 표준을 제시하지 못하고 위험 관리 기술에만 의존한다는 점. 또한 각 단계의 시작과 끝이 불분명하여 프로젝트 관리가 어렵다.

### V모델

V 모델은 폭포수 모델이 확장된 형태로 이 모델에서는 아래방향으로 내려가다가 코딩단계를 거친후 위로 향하면서 V자 모향을 형성한다.

- V 모델은 각 개발단계의 작업을 확인하기 위해 테스트 작업을 수행할 것을 제안한다. 이러한 과정을 통해 시스템이 요구사항을 만족하는지를 신뢰할 수 있게 한다.
- 생명주기 초반부터 테스트 작업을 고려함으로써 요구사항이나 기능 명세와 같은 문서들의 품질보증을 지원한다.
- V 모델은 폭포수 모델과 달리 반복과 재처리 과정을 명백히 한다.
- 폭포수 모델이 문서와 결과물에 초점을 둔 반면, V 모형은 활동과 활동의 정확성에 초점을 둔다.
- 각 단계별로 테스트 작업을 구분함으로써 책임을 명확히 한다. 사용자는 인수테스트, 시스템 테스터는 시스템 테스트, 팀 리더는 통합테스트, 그리고 프로그래머는 단위 테스트를 책임진다.

테스트 작업을 생명주기 초반까지 확장하여 테스트 작업이 전체 생명주기상 항상 수행되는 작업이어야 하고, 코드뿐만 아니라 소프트웨어 요구사항과 설계 결과도 테스트가 가능해야 한다.

#### V 모델 과정

1. 사용자 요구 명세서
2. 시스템 요구 명세서
3. 시스템 설계
4. 단위 설계
5. 코드작성
6. 단위 테스트(4)
7. 통합 테스트(3)
8. 시스템 테스트(2)
9. 인수 테스트(1)

### 애자일 방법

애자일 방법은 품질의 저하 없이 변화를 수용하고, 협업을 강조하며, 제품의 빠른 인도를 강조하는 반복적 방법이다.
이 방법에서는 기존의 방법들과는 다르게 계획, 모델의 작성 및 문서화 작업이 중요하기는 하나, 이러한 것들은 작동되는 코드를 개발하기 위한 과정들일 뿐이며, 진정한 개발 산출물은 코드임을 강조한다.

요구사항의 잦은 변화에 대처하려면 설계 단계를 지연시킬 수 있는 초기 문서화 작업에 매달리지 말아야 한다는 것이다. 또, 애자일 방법은 프로세스 중심이 아닌 사람 중심의 방법으로 개발을 즐기기 위해 요구사항이나 설계 문서를 최소하할 필요가 있다는 것이다.

#### 익스트림 프로그래밍(XP)

소프트웨어 엔지니어인 Kent Beck의 저서에서 처음등장한 용어로 반복적 개발이나 고객의 개발 참여와 같은 기존의 좋은 실천 기술들을
'extreme' 수준으로 끌어올려라라는 의미이다.

켄트 벡은 테스트주도개발을 퍼뜨린 장본이기도 한데 그것은 익스트림 프로그래밍에서의 테스트 선행 프로그래밍과 관계 있다.

XP에서 모든 요구사항들은 시나리오(사용자 스토리)들로 표현되고 각 스토리카드의 내용은 여러 구현단위(태스크) 카드로 나뉜다.
프로그래머들은 짝(pair)를 이루어 작업하며 태스크마다 테스트를 수행한다. 모든 테스트가 성공하면 새로운 코드가 시스템에 통합된다.

소프트웨어의 새로운 버전은 하루에도 여러 번 만들어질 수 있고, 보통 2주에 한 번꼴로 점증이 전달된다.

XP에서는 변경을 고려한 설계가 쓸모없는 노력이라 본다. 예상했던 변경은 자주 일어나지 않으며, 오히려 예상하지 못했던 변경으로 소프트웨어 구조가 나빠지고 구현이 어려워질 수 있다는 것이다.
따라서 XP에서는 리팩토링을 통한 지속적인 코드개선을 제안한다.

##### pair programming

두 사람이 하나의 작업대에서 짝이 되어 함께 작업하는 기술이다. 한 사람은 코딩을 하고 다른 사람은 코딩 과정을 검사한다.

- 시스템 코드에 대한 소유권과 지식을 공유하고 문제 해결에 대한 책임을 팀이진다
- 적어도 두 사람이 코드를 살피기 때문에 비형식적 검토가 가능하다
- 코드 개선을 위한 리팩토링이 장려된다

##### 테스트 선행 개발

테스트 선행개발은 테스트 케이스를 먼저 작성한 후, 이것을 통과하는 코드를 만들고 다듬어서 기능을 구현하는 것이다.
이것을 위해서는 먼저 개발하고자 하는 기능에 관한 인터페이스와 요구사항들이 정의되어야 한다.

XP에서는 스토리 카드가 태스크들로 분해되어 구현되기 때문에 요구사항과 코드는 명확한 관계를 가진다.
코드가 구현된 후 즉시 테스트를 수행할 수 있도록 코드 작성전에 각 태스크별로 하나 이상의 테스트케이스를 만든다.

이후 과정은 자동화된 테스트 도구를 사용하여 진행한다.