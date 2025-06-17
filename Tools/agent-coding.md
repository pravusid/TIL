# Agent Coding

- <https://en.wikipedia.org/wiki/Vibe_coding>
- [[cursor#Prompt]]
- [[vs-code#Copilot Prompt]]

## 실천방법

- 커스텀 룰 사용, 결과 확인하면서 커스텀 룰을 지속적으로 개선
- 한 세션을 짧게 유지
- 명확한 요구사항 전달 및 활용할 기존 코드베이스 참조
- feedback loop (빌드/테스트가 작동할 때 까지 코드를 수정해줘)
- agent에게 중간 커밋 요청
- 여러 모델을 섞어 구현-리뷰 교차실행

### [Vibe Coding 매뉴얼: AI 지원 개발을 위한 템플릿](https://roboco.io/posts/vibe-coding-manual/)

1. 명세(Specification): 목표를 정의합니다(예: “로그인 기능이 있는 Twitter 클론 구축”).
1. 규칙(Rules): 명시적인 제약 조건을 설정합니다(예: “Python 사용, 복잡성 피하기”).
1. 감독(Oversight): 프로세스를 모니터링하고 조정하여 일관성을 보장합니다.

`.cursor/rules` 디렉토리(또는 `.windsurfrules`)에 각각 고유한 목적을 가진 네 개의 파일(또는 섹션)로 구성됩니다

1. 코딩 선호도 – 코드 스타일 및 품질 표준을 정의합니다.
1. 기술 스택 – 도구 및 기술을 명시합니다.
1. 워크플로우 선호도 – AI의 프로세스 및 실행을 관리합니다.
1. 커뮤니케이션 선호도 – AI-인간 상호작용에 대한 기대치를 설정합니다.

#### 예시

Global rules

```md
1️⃣ 구현 작업 원칙

- SOLID 원칙을 사용해서 구현하세요:
  - 단일 책임 원칙 (Single Responsibility Principle)
  - 개방-폐쇄 원칙 (Open-Closed Principle)
  - 리스코프 치환 원칙 (Liskov Substitution Principle)
  - 인터페이스 분리 원칙 (Interface Segregation Principle)
  - 의존성 역전 원칙 (Dependency Inversion Principle)
- TDD로 구현하세요: 테스트 주도 개발 방식으로 먼저 테스트를 작성하고 구현하세요.
- Clean Architecture를 사용해서 구현하세요: 책임과 관심사를 명확히 분리하여 구현하세요.

2️⃣ 코드 품질 원칙

- 단순성: 언제나 복잡한 솔루션보다 가장 단순한 솔루션을 우선시하세요.
- 중복 방지: 코드 중복을 피하고, 가능한 기존 기능을 재사용하세요 (DRY 원칙).
- 가드레일: 테스트 외에는 개발이나 프로덕션 환경에서 모의 데이터를 사용하지 마세요.
- 효율성: 명확성을 희생하지 않으면서 토큰 사용을 최소화하도록 출력을 최적화하세요.

3️⃣ 리팩토링

- 리팩토링이 필요한 경우 계획을 설명하고 허락을 받은 다음 진행하세요.
- 코드 구조를 개선하는 것이 목표이며, 기능 변경은 아닙니다.
- 리팩토링 후에는 모든 테스트가 통과하는지 확인하세요.

4️⃣ 디버깅

- 디버깅 시에는 원인 및 해결책을 설명하고 허락을 받은 다음 진행하세요.
- 에러 해결이 중요한 것이 아니라 제대로 동작하는 것이 중요합니다.
- 원인이 불분명할 경우 분석을 위해 상세 로그를 추가하세요.

5️⃣ 언어

- 한국어로 소통하세요.
- 문서와 주석도 한국어로 작성하세요.
- 기술적인 용어나 라이브러리 이름 등은 원문을 유지해도 됩니다.

6️⃣ Git 커밋

- --no-verify를 절대 사용하지 마세요.
- 명확하고 일관된 커밋 메시지를 작성하세요.
- 적절한 크기로 커밋을 유지하세요.

7️⃣ 문서화

- 주요 컴포넌트 개발 후에는 /docs/[component].md에 간략한 요약을 작성하세요.
- 문서는 코드와 함께 업데이트하세요.
- 복잡한 로직이나 알고리즘은 주석으로 설명하세요.
```

Workspace rules

```txt
1️⃣ 기술 스택 - “이 도구들을 사용하세요”
- 개발 도구
  - 백엔드: Python 사용
  - 인프라: Pulumi for TypeScript, CloudFormation
  - 데이터 저장: MySQL 호환 Aurora Serverless
  - 테스트: pytest, Jest
- 추가 정보
  - 추가 도구가 명시적으로 요청되면 여기에 포함될 수 있습니다.
  - 명시적인 승인 없이는 스택을 변경하지 마세요.
  - Terraform이나 CDK의 리소스 Description은 영어로 작성하세요.

2️⃣ 워크플로우 선호도 - “이런 방식으로 작업하세요”
- 기본 과정
  - 초점: 지정된 코드만 수정하고, 다른 부분은 그대로 두세요.
  - 단계: 큰 작업을 단계로 나누고, 각 단계 후에는 승인을 기다리세요.
  - 계획: 큰 변경 전에는 설계 및 작업개요 문서 [이슈명]_design.md 와 구현 계획 문서 [이슈명]_plan.md를 작성하고 확인을 기다리세요.
  - 추적: 완료된 작업은 progress.md에 기록하고, 다음 단계는 TODO.txt에 기록하세요.
- 고급 워크플로우
  - 테스팅: 주요 기능에 대한 포괄적인 테스트를 포함하고, 엣지 케이스 테스트를 제안하세요.
  - 컨텍스트 관리: 컨텍스트가 100k 토큰을 초과하면 context-summary.md로 요약하고 세션을 재시작하세요.
  - 적응성: 피드백에 따라 체크포인트 빈도를 조정하세요(더 많거나 적은 세분화).

3️⃣ 커뮤니케이션 선호도 - “이렇게 소통하세요”
- 기본 소통
  - 요약: 각 컴포넌트 후에 완료된 작업을 요약하세요.
  - 변경 규모: 변경을 작은, 중간, 큰 규모로 분류하세요.
  - 명확화: 요청이 불명확하면 진행 전에 질문하세요.
- 정밀 소통
  - 계획: 큰 변경의 경우 구현 계획을 제공하고 승인을 기다리세요.
  - 추적: 항상 완료된 작업과 대기 중인 작업을 명시하세요.
  - 감정적 신호: 긴급성이 표시되면(예: “이것은 중요합니다—집중해주세요!”) 주의와 정확성을 우선시하세요.

4️⃣ 프로젝트 구조
- 디렉토리 구조
  - docs/: 모든 문서 파일
    - architecture/: 아키텍처 문서
    - guides/: 개발자 가이드
    - runbooks/: 운영 매뉴얼
  - src/: 소스 코드
    - core/: 핵심 비즈니스 로직
    - infrastructure/: 인프라 관련 코드
    - api/: API 엔드포인트
  - tests/: 테스트 파일
- 명명 규칙
  - 파일명: 스네이크 케이스(snake_case) 사용 (예: user_service.py)
  - 클래스명: 파스칼 케이스(PascalCase) 사용 (예: UserService)
  - 함수와 변수명: 스네이크 케이스(snake_case) 사용 (예: get_user())
  - 상수: 대문자 스네이크 케이스(UPPER_SNAKE_CASE) 사용 (예: MAX_USERS)

5️⃣ 활용 방법
- 이 규칙 세트는 AI 지원 개발을 위한 템플릿입니다. 다음과 같이 사용하세요:
  - 프로젝트 시작 시 이 규칙을 참조하세요.
  - 필요에 따라 규칙을 조정하세요.
  - AI 모델에게 이 파일의 내용을 따르도록 지시하세요.
  - 프로젝트를 진행하면서 이 규칙이 어떻게 도움이 되는지 평가하세요.
- 이 규칙 세트를 통해 AI와의 협업이 더 효율적이고 예측 가능해질 것입니다.
```

### [Vibe 코딩은 저품질 작업에 대한 변명이 될 수 없어요](https://news.hada.io/topic?id=20449)

Rules for high-quality vibe coding (practical guidelines)

- Rule 1: Always Review AI-Generated Code / AI 코드 반드시 리뷰하기
  - 예외 없음. AI가 작성한 모든 코드는 주니어 개발자가 작성한 코드처럼 리뷰해야 함
  - 개별 리뷰이든 동료 리뷰이든 반드시 수행
  - Copilot, ChatGPT, Cursor 등 어떤 AI라도 마찬가지
  - 리뷰할 시간이 없다면, 그 코드를 쓸 시간도 없는 것
  - 리뷰 없이 AI 코드를 머지하는 건 리스크를 그대로 끌어안는 것과 같음
- Rule 2: Establish Coding Standards and Follow Them / 코딩 스타일과 기준을 설정하고 준수할 것
  - AI는 학습한 코드 스타일을 그대로 반영하므로, 일관된 팀 기준이 없다면 품질이 들쭉날쭉해짐
  - 팀의 스타일 가이드, 아키텍처 패턴, 코딩 규칙을 명확히 정의해야 함
  - 예: “모든 함수에는 JSDoc과 유닛 테스트가 있어야 한다” → AI가 생성한 코드도 마찬가지로 적용
  - 계층 구조나 레이어드 아키텍처를 사용하는 프로젝트에서,
  - AI가 UI 코드 안에 DB 호출을 넣지 않도록 리팩터링 필수
  - 흔한 AI 실수(ex: 복잡한 함수, deprecated API 사용 등)를 잡는 lint 또는 정적 분석 룰 도입 추천
- Rule 3: Use AI for Acceleration, Not Autopilot / AI는 가속기이지 자동 조종 장치가 아님
  - vibe coding은 잘 알고 있는 작업을 빠르게 처리하는 용도로 사용해야 함
  - 좋은 활용 예:
    - 보일러플레이트 생성
    - 컴포넌트 스캐폴딩
    - 언어 변환
    - 간단한 알고리즘 뼈대 작성
  - 위험한 사용 예:
    - 애매한 설명으로 모듈 전체를 설계하게 하기
    - 잘 모르는 도메인에 코드 생성 시도
  - 코드가 영구적으로 남을 예정이라면, 반드시 vibe 모드에서 engineering 모드로 전환 필요
- Rule 4: Test, Test, Test / 테스트는 무조건 해야 한다
  - AI가 코드를 생성했다고 해서 자동으로 정답이 되는 건 아님
  - 모든 주요 경로에 대해 테스트 작성 필수
  - AI가 테스트도 만들어줬다면, 그 테스트가 실제로 유효한지도 검토 필요
  - 특히 UI 기능이나 유저 입력이 많은 부분은 직접 클릭, 비정상 입력 테스트 필수
  - vibe-coded 앱은 해피 패스만 잘 작동하고, 예외 입력에 취약한 경우가 많음
- Rule 5: Iterate and Refine / 반복하고 다듬기
  - AI가 처음 준 결과물이 만족스럽지 않다면, 그냥 넘어가지 말고 다시 시도하거나 리팩터링
  - vibe coding은 대화 기반의 반복적 프로세스임
  - 예:
    - “이 코드 더 간결하게 해줘”
    - “작은 함수들로 나눠줘” 등 프롬프트 재조정
  - 또는 직접 리팩터링 → 수정 포인트 → 다시 프롬프트 → 반복
  - AI와의 사이클 사용 전략이 효과적
- Rule 6: Know When to Say No / 거절할 줄 알아야 함
  - vibe coding이 항상 최선은 아님
  - 중요한 설계나 보안이 필요한 상황에선 직접 작성하는 것이 낫다
  - 예:
    - 보안 관련 모듈은 직접 설계하고, 일부만 AI 활용
    - 단순한 문제에 대해 AI가 복잡하게 답할 경우, 직접 짜는 게 더 빠름
  - AI가 문제를 제대로 해결하지 못할 때는 고집하지 말고 수동 모드로 전환할 것
  - "AI가 해줬으니까"는 내가 내 코드를 이해하지 못해도 된다는 핑계가 아님
- Rule 7: Document and Share Knowledge / 문서화하고 지식을 공유하라
  - AI가 생성한 코드도 직접 쓴 코드만큼 문서화가 되어야 함 (때로는 더 많이)
  - 비직관적인 결정이나 특이한 구현이 있다면 주석을 남겨야 함
  - 어떤 부분이 AI 생성인지 팀원에게 명확히 공유
  - 일부 팀은 주요 AI 코드에 사용한 프롬프트를 그대로 저장함 → 디버깅에 유용

### [AI 시대의 시니어 개발자 역량 : 더 나은 결과를 위한 경험 활용](https://news.hada.io/topic?id=20152)

따라서 AI 코딩 세션 시작 전에 다음과 같은 품질 보장 도구들을 반드시 세팅함

- black, isort: 코드 포맷팅
- ruff: 린팅
- mypy: 타입 검사
- bandit: 보안 분석
- 테스트 스위트 전반

파일 기반의 키프레임 기법

- AI는 창의적인 구현에는 강하지만, 코드 구조나 파일 구성에 대해서는 방향성이 부족할 수 있음
  - 이를 보완하기 위해 사용하는 전략이 파일 기반 키프레임(file-based keyframing) 임
- 이 기법은 애니메이션 제작에서의 키프레임 방식에서 영감을 얻음:
  - 숙련된 애니메이터가 중요 장면(키프레임) 을 먼저 만들고, 나머지는 보조 인력이 채워 넣는 방식
  - 품질을 유지하면서 작업 효율을 높일 수 있음
- 실제 AI 코딩 프로젝트에서는 구현 전 미리 빈 껍데기 파일(stub files) 을 생성해 둠
  - 예: API 엔드포인트, API 클라이언트, 컨트롤러 클래스, Twig 템플릿 등
- 이러한 키프레임 파일은 AI에게 다음과 같은 중요한 문맥 정보를 제공함
  - 프로젝트의 파일 구성 방식
  - 네임스페이스 구조
  - 명명 규칙
  - 일관된 코드 패턴
- 프롬프트로 모든 구조를 설명하기보다, 코드베이스 자체에 힌트를 제공함으로써 AI의 추론 정확도를 높일 수 있음
- 이 접근 방식은 AI 시대에도 여전히 중요한 "이름 짓기" 의 원칙을 강조함
  - AI는 언어를 기반으로 동작하기 때문에, 의도와 의미가 담긴 텍스트는 더 나은 결과를 이끌어냄

### [바이브 코딩 바이블: AI 에이전트 시대의 새로운 코딩 패러다임](https://tech.kakao.com/posts/696)

#### 실전 중심 코드 프롬프팅 4가지 유형

- 코드 작성 (Code Generation)
  - 요구사항은 구체적으로 열거하세요. 단순히 “웹 크롤러 만들어줘”보다는 “파이썬으로 requests와 BeautifulSoup을 사용해 특정 URL의 HTML을 가져오고 <a> 태그의 링크를 추출하는 함수”처럼 상세히 적으면 더 정확한 결과를 얻습니다. 프롬프트에 언어, 목적, 제약조건(예: 시간복잡도, 사용 라이브러리)까지 포함하면 좋습니다.
  - 생성된 코드는 바로 실행하여 테스트하고, 의도와 다르면 다시 프롬프트를 수정하거나 추가 지시를 내리는 반복 과정을 거치세요. 처음 한 번에 완벽한 코드를 얻기보다 AI와 인터랙티브하게 개선해간다는 생각으로 접근하면 효율적입니다.
- 코드 설명 (Code Explanation)
  - 코드 블럭 전체를 프롬프트에 넣어야 AI가 정확한 설명을 할 수 있습니다. 부분만 주면 문맥이 부족해 잘못된 설명이 나올 수 있으니, 필요한 경우 코드 길이에 맞춰 프롬프트를 나누어 전달하세요.
  - 설명이 너무 복잡하거나 전문적이면, 이해하기 쉬운 비유나 단계별 풀이를 요청해보세요. 예를 들어 “한글로 차근차근 설명해줘” 또는 “간단한 예를 들어 설명해줘”처럼 프롬프트를 조정하면 더 쉬운 설명을 얻을 수 있습니다.
  - 보안이나 민감한 코드라면, 직접 실행하기 전에 AI의 설명이 맞는지 검증이 필요합니다. LLM이 가끔 헷갈려 할 수 있으므로, 중요한 부분은 다시 질문하거나 다른 방법으로 이중 체크하세요.
- 프로그래밍 언어 변환 (Programming Language Translation)
- 디버깅 및 코드 리뷰 (Debugging & Code Review)
  - 오류 메시지가 있다면 프롬프트에 함께 포함하세요. AI는 스택 트레이스나 에러 메시지를 단서로 문제를 더 정확히 파악할 수 있습니다. 예: “이 코드를 실행하니 IndexError가 발생하는데, 원인을 찾아 고쳐줘” 같이 맥락을 주세요.
  - AI의 수정안을 적용하기 전에 반드시 테스트 케이스를 실행해보세요. AI가 논리적으로 그럴듯한 답을 내놔도 실제로 문제를 완벽히 해결하지 못했을 수 있으므로, 다양한 입력으로 검증하는 단계가 필요합니다.
  - 단순 버그뿐 아니라 코드 스타일, 성능, 보안 측면도 함께 검토해달라고 요청해보세요. “이 코드의 개선점을 리뷰해줘” 처럼 요구하면 AI가 종합적인 리뷰를 수행하여 더 나은 코드를 제안해줍니다. 다만 최종 판단은 개발자가 해야 한다는 점을 잊지 마세요.

#### 바이브 코딩 관점에서의 프롬프트 설계 모범 사례

- 맥락과 목표 충분 제공
  - 프롬프트에는 필요한 맥락(context)을 최대한 담도록 합니다. 예를 들어 기존에 작성된 코드 조각이나 환경 정보, 사용하고자 하는 라이브러리 버전, 목표 플랫폼 등을 알려주면 AI가 그에 맞춰 답변할 확률이 높아집니다.
  - 또한 달성하려는 목표를 명확히 기술하세요. 모호한 요구보다 "주어진 JSON 문자열을 파싱하여 특정 필드를 추출"처럼 구체적 목표를 제시하면 결과가 더 정확합니다.
- 단계별 접근(체인 프롬프팅)
  - 복잡한 문제는 한 번의 프롬프트로 모두 해결하려 하지 말고, 여러 단계로 나눠서 해결하세요. 이를 체인 프롬프팅(chain prompting) 또는 연쇄 프롬프트 기법이라고 합니다. 예를 들어 먼저 “필요한 모듈 설계”를 묻고, 다음 프롬프트에서 “각 모듈의 상세 구현”을 요청하는 식입니다.
  - 또는 AI에게 “우선 해결 방안을 계획하고 차례로 실행하자”라고 해서 사고 과정을 문장으로 풀게 한 뒤, 그 계획에 따라 코드를 작성하도록 유도할 수도 있습니다. 한꺼번에 안 되는 경우 프롬프트를 재구성하거나 문제를 작은 단위로 쪼개어 제시하면 돌파구를 얻을 수 있으며, 이런 단계별 접근은 복잡한 프로젝트에서 특히 효과적입니다.
- 에이전트와 도구 활용
  - 최신 AI 모델과 프레임워크들은 프롬프트를 통해 툴 사용이나 자동 실행까지 수행하는 에이전트 기능을 제공합니다. 예를 들어 OpenAI의 플러그인을 사용하면 모델이 인터넷 검색이나 코드 실행 같은 액션을 프롬프트의 일부로 수행할 수 있고, GitHub 코파일럿의 에이전트 모드나 Replit의 Ghostwriter 에이전트는 프롬프트만으로 전체 애플리케이션을 만들어 배포하는 것을 목표로 하고 있습니다.
  - 가능하다면 이러한 도구 연계형 프롬프트를 활용해보세요. 예를 들어 "테스트를 실행하고 실패하면 출력 오류 내용을 알려줘"라고 하면, 에이전트가 코드를 실행해보고 결과를 해석한 뒤 후속 조치를 취해줄 수도 있습니다.
  - 다만 현재 시점의 에이전트 기술은 완전하지 않으므로, 자동화된 결과도 사람이 모니터링하면서 조정하는 human-in-the-loop 절차가 필요합니다.
- 프롬프트 템플릿과 예시 활용
  - 자주 쓰는 프롬프트 패턴이 있다면 템플릿으로 저장해 두세요. 예를 들어 “버그 찾기”나 “리팩토링 제안” 등은 일정한 형식이 효과적일 수 있으므로, 잘 된 프롬프트와 그 출력 예시를 모아서 필요할 때 참고하면 편리합니다.
  - 또한 Few-shot 학습 방식으로, 프롬프트에 원하는 출력의 예시를 간단히 포함시키는 것도 유용합니다. 예를 들어 JSON 출력이 필요하면 프롬프트 끝에 작은 예시 JSON을 보여주어 AI가 그 형식을 따르게 하는 것이죠. 이런 기법들은 모델이 사용자 의도를 더 정확히 파악하도록 도와줍니다.
- 반복적 실험과 개선
  - 프롬프트 설계는 한 번에 완벽해지지 않습니다. AI가 의도대로 답하지 않으면 프롬프트를 수정하고 재시도하세요. 때로는 같은 질문을 약간 다른 표현으로 여러 번 시도해보고 최적의 대답을 선택하는 실험적 접근도 필요합니다. 예를 들어 첫 응답이 부족하면, 요구사항을 다시 강조하거나 문제를 다른 각도에서 묻는 등 프롬프트를 조절해 볼 수 있습니다.
  - 중요한 것은 AI를 협력자로 여기고 대화하듯이 계속 방향을 잡아나가는 것입니다. 이러한 iterative한 프롬프트 개선 과정 자체가 바이브 코딩의 일부이며, 개발자의 의도에 AI를 맞춰가는 훈련이 되어 장기적으로 프롬프트 실력을 향상시켜줍니다.

### [The Prompt Engineering Playbook for Programmers](https://addyo.substack.com/p/the-prompt-engineering-playbook-for)

| Technique                 | Prompt template                                                                                                                            | Purpose                                                                |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------- |
| 1. Role Prompting         | “You are a senior {language} developer. Review this function for {goal}.”                                                                  | Simulate expert-level code review, debugging, or refactoring           |
| 2. Explicit Context Setup | “Here’s the problem: {summary}. The code is below. It should do {expected behavior}, but instead it’s doing {actual behavior}. Why?”       | Frame the problem clearly to avoid generic, surface-level responses    |
| 3. Input/Output Examples  | “This function should return {expected output} when given {input}. Can you write or fix the code?”                                         | Guide the assistant by showing intent through examples                 |
| 4. Iterative Chaining     | “First, generate a skeleton of the component. Next, we’ll add state. Then handle API calls.”                                               | Break larger tasks into steps to avoid overwhelming or vague prompts   |
| 5. Debug with Simulation  | “Walk through the function line by line. What are the variable values? Where might it break?”                                              | Get the assistant to simulate runtime behavior and surface hidden bugs |
| 6. Feature Blueprinting   | “I’m building {feature}. Requirements: {bullets}. Using: {tech stack}. Please scaffold the initial component and explain your choices.”    | Kick off feature development with AI-led planning and scaffolding      |
| 7. Code Refactor Guidance | “Refactor this code to improve {goal}, such as {e.g., readability, performance, idiomatic style}. Use comments to explain changes.”        | Make AI refactors align with your goals, not arbitrary changes         |
| 8. Ask for Alternatives   | “Can you rewrite this using a functional style? What would a recursive version look like?”                                                 | Explore multiple implementation paths and expand your toolbox          |
| 9. Rubber Ducking         | “Here’s what I think this function does: {your explanation}. Am I missing anything? Does this reveal any bugs?”                            | Let the AI challenge your understanding and spot inconsistencies       |
| 10. Constraint Anchoring  | “Please avoid {e.g., recursion} and stick to {e.g., ES6 syntax, no external libraries}. Optimize for {e.g., memory}. Here’s the function:” | Prevent the AI from overreaching or introducing incompatible patterns  |

### [Accelerating Large-Scale Test Migration with LLMs](https://medium.com/airbnb-engineering/accelerating-large-scale-test-migration-with-llms-9565c208023b)

### [Here’s how I use LLMs to help me write code](https://simonwillison.net/2025/Mar/11/using-llms-for-code/)

### [Claude Code: Best practices for agentic coding](https://www.anthropic.com/engineering/claude-code-best-practices)

#### `CLAUDE.md`

<https://docs.anthropic.com/ko/docs/build-with-claude/prompt-engineering/overview>

- Common bash commands
- Core files and utility functions
- Code style guidelines
- Testing instructions
- Repository etiquette (e.g., branch naming, merge vs. rebase, etc.)
- Developer environment setup (e.g., pyenv use, which compilers work)
- Any unexpected behaviors or warnings particular to the project
- Other information you want Claude to remember

예시

```md
# Bash commands

- npm run build: Build the project
- npm run typecheck: Run the typechecker

# Code style

- Use ES modules (import/export) syntax, not CommonJS (require)
- Destructure imports when possible (eg. import { foo } from 'bar')

# Workflow

- Be sure to typecheck when you’re done making a series of code changes
- Prefer running single tests, and not the whole test suite, for performance
```

CLAUDE.md 파일은 Claude에게 주는 지시사항과 같기 때문에, 자주 사용하는 프롬프트처럼 계속 다듬고 개선해야 합니다.
많은 사람들이 실수를 하는 부분은 일단 많은 내용을 추가하고 나서 그 내용이 제대로 작동하는지 확인하는 과정을 생략하는 것입니다.
하지만 Claude가 지시를 더 잘 따르도록 하려면, 시간을 들여 다양한 방법으로 실험해보고 어떤 내용이 가장 효과적인지 파악해야 합니다.

#### Try common workflows

주요 내용은 Claude Code를 사용하여 코딩 작업을 수행하는 세 가지 효과적인 워크플로우입니다. Claude는 유연하게 사용할 수 있지만, 다음 방법들이 특히 유용하다는 것이 밝혀졌습니다.

##### 탐색, 계획, 코딩, 커밋 (Explore, plan, code, commit)

이 워크플로우는 대부분의 코딩 문제에 적합합니다.

- 탐색/조사: Claude에게 관련 파일(예: logging.py), 이미지 또는 URL을 읽도록 요청합니다. 이때는 아직 코드를 작성하지 않도록 명확하게 지시합니다. 복잡한 문제의 경우 서브 에이전트를 사용하여 세부 정보를 확인하거나 질문을 조사하도록 하면 컨텍스트를 유지하는 데 도움이 됩니다.
  - 예시: "Claude, user_authentication.py 파일과 database_schema.sql 파일을 읽어줘. 아직 코드는 작성하지 마."
- 계획: Claude에게 문제 해결을 위한 계획을 세우도록 요청합니다. "think", "think hard", "ultrathink"와 같은 단어를 사용하여 Claude가 더 많은 생각 시간을 갖도록 지시합니다. 계획이 만족스러우면 Claude에게 해당 계획을 문서나 GitHub 이슈로 만들도록 요청할 수 있습니다.
  - 예시: "Claude, 사용자 로그인 기능 구현에 대한 계획을 세워줘. think hard."
- 코딩: Claude에게 계획에 따라 솔루션을 코드로 구현하도록 요청합니다. 코드를 구현하면서 합리적인지 확인하도록 요청할 수 있습니다.
  - 예시: "이제 계획대로 사용자 로그인 기능을 코드로 구현해줘."
- 커밋: 결과물을 커밋하고 Pull Request를 생성하도록 요청합니다. 관련이 있다면 README나 변경 로그를 업데이트하도록 요청할 수도 있습니다.
  - 예시: "코드 변경 사항을 커밋하고 Pull Request를 생성해줘. 그리고 README.md도 업데이트해줘."

##### 테스트 작성, 커밋; 코딩, 반복, 커밋 (Write tests, commit; code, iterate, commit)

이 워크플로우는 테스트를 통해 쉽게 확인할 수 있는 변경 사항에 특히 강력합니다. **테스트 주도 개발(TDD)**과 유사합니다.

- 테스트 작성: 예상 입력/출력 쌍을 기반으로 테스트를 작성하도록 요청합니다. TDD를 하고 있다는 것을 명확히 하여 Claude가 가짜 구현(mock implementations)을 만들지 않도록 합니다.
  - 예시: "Claude, 사용자 이름과 비밀번호를 입력했을 때 올바른지 확인하는 로그인 기능의 테스트 케이스를 작성해줘. 아직 구현 코드는 작성하지 마."
- 테스트 실행 및 실패 확인: 테스트를 실행하고 실패하는지 확인하도록 지시합니다.
  - 예시: "작성된 테스트를 실행하고 실패하는지 확인해줘."
- 테스트 커밋: 테스트가 만족스러우면 커밋하도록 요청합니다.
  - 예시: "이제 테스트 코드를 커밋해줘."
- 코드 작성 및 반복: Claude에게 테스트를 통과하는 코드를 작성하도록 요청하고, 테스트를 수정하지 않도록 지시합니다. Claude는 코드를 작성하고, 테스트를 실행하고, 코드를 조정하고, 다시 테스트를 실행하는 과정을 반복합니다. 이때 독립적인 서브 에이전트를 사용하여 구현이 테스트에 과적합(overfitting)되지 않는지 확인하도록 요청할 수 있습니다.
  - 예시: "이제 작성된 테스트를 통과하는 코드를 작성해줘. 테스트 코드는 수정하지 마."
- 코드 커밋: 변경 사항이 만족스러우면 코드를 커밋하도록 요청합니다.
  - 예시: "모든 테스트가 통과되었으니 코드를 커밋해줘."

##### 코드 작성, 결과 스크린샷, 반복 (Write code, screenshot result, iterate)

이 워크플로우는 시각적인 결과가 중요한 경우에 유용합니다.

- 스크린샷 도구 제공: Claude가 브라우저 스크린샷을 찍을 수 있는 방법(예: Puppeteer, iOS 시뮬레이터, 수동 복사/붙여넣기)을 제공합니다.
  - 예시: "Claude, 이 Puppeteer 서버를 사용해서 웹페이지 스크린샷을 찍을 수 있어."
- 시각적 Mock 제공: 이미지 파일을 제공하거나, 이미지를 복사/붙여넣기하여 Claude에게 시각적 Mock을 제공합니다.
  - 예시: "이 이미지는 우리가 만들고자 하는 로그인 페이지의 디자인이야."
- 코드 구현 및 반복: Claude에게 코드를 구현하고, 결과 스크린샷을 찍고, 결과가 Mock과 일치할 때까지 반복하도록 요청합니다.
  - 예시: "이 디자인을 HTML/CSS로 구현하고, 매번 스크린샷을 찍어서 디자인과 일치하는지 확인해줘."
- 커밋: 결과가 만족스러우면 커밋하도록 요청합니다.
  - 예시: "디자인과 정확히 일치하니 코드를 커밋해줘."

#### Optimize workflow

Claude Code’s success rate improves significantly with more specific instructions, especially on first attempts.
Giving clear directions upfront reduces the need for course corrections later.

- add tests for foo.py
  - Good: write a new test case for foo.py, covering the edge case where the user is logged out. avoid mocks
- why does ExecutionFactory have such a weird api?
  - Good: look through ExecutionFactory's git history and summarize how its api came to be
- add a calendar widget
  - look at how existing widgets are implemented on the home page to understand the patterns and specifically how code and interfaces are separated out.
    HotDogWidget.php is a good example to start with. then, follow the pattern to implement a new calendar widget that lets the user select a month and paginate forwards/backwards to pick a year.
    Build from scratch without libraries other than the ones already used in the rest of the codebase.

### [Cursor 사용 방법 (+ 최고의 팁)](https://siosio3103.medium.com/e931ced0429f)

YOLO Mode

```md
any kind of tests are always allowed like vitest, npm test, nr test, etc. also basic build commands like build, tsc, etc.
creating files and making directories (like touch, mkdir, etc) is always ok too
```

복잡한 작업

```md
Create a function that converts a markdown string to an HTML string.
Write tests first, then the code, then run the tests and update the code until tests pass.
```

오류 수정

```md
I've got some build errors. Run nr build to see the errors, then fix them, and then run build until build passes.
```

로그 기반 디버깅과 반복적 수정

복잡하고 까다로운 문제를 다룰 때는, Cursor에 로그를 출력하도록 지시한 뒤 그 로그를 활용해 문제를 해결하는 방식이 꽤 유용합니다. 제가 자주 사용하는 방법은 이렇습니다.
기존 이슈를 처리하기 위해 여러 방법을 시도했지만, 원하는 결과가 안 나오는 경우 전 Cursor에 이렇게 말합니다.

- “코드에 로그를 추가해서 내부 동작을 좀 더 잘 파악할 수 있게 해 줘. 내가 코드를 실행한 후에 로그 결과를 너한테 다시 알려줄게.”
- Cursor는 주요 지점에 로그 코드를 추가할 겁니다.
- 코드를 실행하고 로그를 수집합니다.
- Cursor로 돌아가 다음과 같이 말합니다 “여기에 로그 결과물이 있어. 지금 보니까 어떤 문제가 있는 것 같아? 어떻게 고치면 될까?
- Cursor가 분석할 로그 결과물을 붙여 넣습니다.

### [Cloudflare가 Claude와 함께 OAuth를 빌드하고 모든 프롬프트를 공개함](https://news.hada.io/topic?id=21256)

- <https://github.com/cloudflare/workers-oauth-provider/commits/main/>
- [Cloudflare의 AI가 작성한 OAuth 라이브러리 살펴보기](https://news.hada.io/topic?id=21354)

### [AI 지원 프로그래밍을 위한 새 실천법](https://wiki.g15e.com/pages/New%20practices%20for%20AI-aided%20programming)

## 참고, 의견

- [무신사 X GitHub Copilot은 정말로 우리의 생산성을 높였을까?](https://medium.com/musinsa-tech/de149ad7b7f6)
- [나만의 Visual Studio Code Copilot 지침 만들고 활용하기](https://d2.naver.com/helloworld/6615449)
- [LLM이 실제로 프로그래머의 생산성을 얼마나 향상시키고 있을까?](https://news.hada.io/topic?id=19672)
- [AI Blindspots – AI 코딩 중에 발견한 LLM의 맹점들](https://news.hada.io/topic?id=19859)
- [에이전틱 코딩에서 개발자 역량의 역할](https://news.hada.io/topic?id=20006)
- [AI 시대의 새로운 개발자 패턴들](https://news.hada.io/topic?id=21115)
