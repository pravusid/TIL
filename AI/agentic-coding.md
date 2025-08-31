# Agentic Coding

- <https://en.wikipedia.org/wiki/Vibe_coding>
- [[agents]]

## 실천방법

- 커스텀 룰 사용
  - 결과 확인하면서 커스텀 룰을 지속적으로 개선
  - 프로젝트 root 규칙 (참고: [[claude-code-best-practices#CLAUDE.md]])
    - 최소한의 내용으로 유지 (100줄 정도)
    - 프로젝트의 목적, 프로젝트 구조, 주요 명령어 (npm script, makefile ...), 코드 스타일, 핵심 지침 (너무 일반적인 내용이 아니라 구체적인 내용)
  - 하위 디렉토리에서 추가 규칙이 필요하다면 디렉토리별 커스텀 룰 사용
  - 그러나 코딩 에이전트들은 커스텀 룰을 제대로 따르지 않는 경우도 많기 때문에 중요한 지시사항이라면 프롬프트에 포함
- 한 세션을 불필요하게 길게 유지하지 않음
- 명확한 요구사항 전달, 활용할 기존 코드베이스 참조 제공
  - > Claude Code의 성공률은 특히 첫 시도에서 더 구체적인 지침을 제공할 때 크게 향상됩니다
- 작업을 시작하기 전 모델에게 아이디어를 구상하거나 계획을 세우도록 요청 (plan mode)
  - 아이디어, 계획을 피드백 하면서 방향을 맞춤
  - 계획을 세우기에 context 내용이 불충분 한 경우 재질문 하도록 프롬프트
  - 계획은 대형/추론 모델로 진행하고 (claude opus ...), 코드 작성은 일반모델(claude sonnet ...)으로 진행하는 것도 방법
  - 계획이 어느정도 정리되면 파일로 저장하도록 요청하고 본격적인 실행 전 다듬는 작업 진행
  - 계획에 TODO list 생성하고 해결여부 표시하도록 요청
    - 각 단계가 끝날 때마다, 현재 파일의 step 체크 처리하고, 커밋해줘
    - 단계별 TODO list에 타입체크, 테스트 등을 확인하도록 해서 커밋 전 작동여부 추가 수정하도록 가이드
- 계획을 세웠다면 프롬프트에서 지시를 정확히 수행하도록 요구 (`지시사항을 빠짐없이 정확히 준수해서 작업을 진행해줘`)
- feedback loop (빌드/테스트가 작동할 때 까지 코드를 수정해줘)
- agent에게 각 단계가 끝날 때 커밋 요청
  - husky 등의 git-hooks를 섞어 쓰면 타입오류, lint, 테스트 ... 적용할 수 있음
- 여러 모델을 섞어 구현-리뷰 교차실행
- AI에게는 오히려 코드 중간중간의 주석이 도움이 됨 (why, what, how 모두)

### 계획 단계에서 context 보강

```txt
이 작업을 진행하기 위해 다음 정보들을 먼저 확인해주세요
- 불명확하거나 부족한 정보가 있다면 반드시 먼저 질문해주세요
- 가정을 하기보다는 명확히 확인 후 진행해주세요

정보가 충분하지 않다고 판단되면 "다음 정보가 필요합니다:"로 시작하는 질문 목록을 제시해주세요.
```

### 탐색

- 질문을 작성하기 위해 모델과 질의응답
- 내용이 어느정도 확보되면 질의응답을 다시 질문으로 요약
- 추론모델에게 정리된 질문을 다시 입력

### [Augmented Coding: Beyond the Vibes - kent beck](./agentic-coding/augmented-coding-beyond-the-vibes.md)

AI가 잘못하고 있다는 세 가지 신호

> 1. 비슷한 행동을 반복한다 (무한루프 등)
> 2. 내가 요청하지 않은 기능 구현. 그게 논리적인 다음 단계가 맞을지라도.
> 3. 테스트를 삭제하거나 비활성화는 등, AI가 치팅하는 걸로 느껴지는 그 외 모든 신호.

[[augmented-coding-beyond-the-vibes#TDD를 돕는 시스템 프롬프트|TDD를 돕는 시스템 프롬프트]]

### [에이전트 기반 AI는 왜 좋은 페어 프로그래머가 아닌가](./agentic-coding/why-agents-are-bad-pair-programmers.md)

> 병목지점은 인간이고, 인간의 속도 향상 정확도 유지를 위한 실천방법이 중요할 듯

### [Cloudflare가 Claude와 함께 OAuth를 빌드하고 모든 프롬프트를 공개함](https://news.hada.io/topic?id=21256)

- <https://github.com/cloudflare/workers-oauth-provider/commits/main/>
- [Cloudflare의 AI가 작성한 OAuth 라이브러리 살펴보기](https://news.hada.io/topic?id=21354)

### [Cursor 사용 방법 (+ 최고의 팁)](./agentic-coding/cursor-tips.md)

### [⭐️Claude Code: Best practices for agentic coding - anthropic](./agentic-coding/claude-code-best-practices.md)

- 설정 사용자 정의 (Customize your setup)
- Claude에게 더 많은 도구 제공 (Give Claude more tools)
- 일반적인 워크플로우 시도 (Try common workflows)
- 워크플로우 최적화 (Optimize your workflow)
- 헤드리스 모드를 사용해 인프라 자동화 (Use headless mode to automate your infra)
- 다중 Claude 워크플로우로 수준 향상 (Uplevel with multi-Claude workflows)

### [⭐️Here’s how I use LLMs to help me write code](https://simonwillison.net/2025/Mar/11/using-llms-for-code/)

### [Accelerating Large-Scale Test Migration with LLMs](https://medium.com/airbnb-engineering/accelerating-large-scale-test-migration-with-llms-9565c208023b)

> Enzyme to React Testing Library

1. File Validation and Refactor Steps
2. Retry Loops & Dynamic Prompting
3. Increasing the Context
4. From 75% to 97%: Systematic Improvement

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

### [바이브 코딩 바이블: AI 에이전트 시대의 새로운 코딩 패러다임](./agentic-coding/vibe-coding-bible.md)

- 맥락과 목표 충분 제공
- 단계별 접근(체인 프롬프팅)
- 에이전트와 도구 활용
- 프롬프트 템플릿과 예시 활용
- 반복적 실험과 개선

### [AI 시대의 시니어 개발자 역량 : 더 나은 결과를 위한 경험 활용](./agentic-coding/how-seasoned-developers-can-achieve-great-results-with-ai-coding-agents.md)

- 실제 AI 코딩 프로젝트에서는 구현 전 미리 빈 껍데기 파일(stub files) 을 생성해 둠
  - 예: API 엔드포인트, API 클라이언트, 컨트롤러 클래스, Twig 템플릿 등
- 이러한 키프레임 파일은 AI에게 다음과 같은 중요한 문맥 정보를 제공함
  - 프로젝트의 파일 구성 방식, 네임스페이스 구조, 명명 규칙, 일관된 코드 패턴

### [Vibe 코딩은 저품질 작업에 대한 변명이 될 수 없어요](agentic-coding/vibe-coding-is-not-an-excuse-for-low-quality-work.md)

- Rule 1: Always Review AI-Generated Code / AI 코드 반드시 리뷰하기
- Rule 2: Establish Coding Standards and Follow Them / 코딩 스타일과 기준을 설정하고 준수할 것
- Rule 3: Use AI for Acceleration, Not Autopilot / AI는 가속기이지 자동 조종 장치가 아님
- Rule 4: Test, Test, Test / 테스트는 무조건 해야 한다
- Rule 5: Iterate and Refine / 반복하고 다듬기
- Rule 6: Know When to Say No / 거절할 줄 알아야 함
- Rule 7: Document and Share Knowledge / 문서화하고 지식을 공유하라

### [Vibe Coding 매뉴얼: AI 지원 개발을 위한 템플릿](agentic-coding/vibe-coding-manual.md)

- [[vibe-coding-manual#Global rules]]
- [[vibe-coding-manual#Workspace rules]]

## Series of Articles

### <https://simonwillison.net/tags/ai-assisted-programming/>

### <https://github.com/humanlayer/12-factor-agents>

### [AI 지원 프로그래밍](https://wiki.g15e.com/pages/AI-aided%20programming)

- [AI 지원 프로그래밍을 위한 새 실천법](https://wiki.g15e.com/pages/New%20practices%20for%20AI-aided%20programming)
- [AI 시대의 소스코드 품질](https://wiki.g15e.com/pages/Source%20code%20quality%20in%20the%20AI%20era)
- [에이전트 기반 코딩](https://wiki.g15e.com/pages/Agentic%20coding)

### [Agentic Engineering - Zed](https://zed.dev/agentic-engineering)

## Prompting

- <https://cookbook.openai.com/examples/gpt-5/gpt-5_prompting_guide>
- <https://cookbook.openai.com/examples/prompt_migration_guide>
- <https://docs.anthropic.com/ko/docs/build-with-claude/prompt-engineering/overview>
- [Gemini CLI Plan Mode](https://gist.github.com/philschmid/379cf06d9d18a1ed67ff360118a575e5)
- [Gemini CLI: Explain Mode](https://gist.github.com/philschmid/64ed5dd32ce741b0f97f00e9abfa2a30)

## Context Engineering

- <https://github.com/davidkimai/Context-Engineering>
- [바이브 코딩에는 컨텍스트 엔지니어링이 필요하다](https://blogbyash.com/translation/vibe-coding-needs-context-engineering/)

## 참고, 의견

- [무신사 X GitHub Copilot은 정말로 우리의 생산성을 높였을까?](https://medium.com/musinsa-tech/de149ad7b7f6)
- [나만의 Visual Studio Code Copilot 지침 만들고 활용하기](https://d2.naver.com/helloworld/6615449)
- [LLM이 실제로 프로그래머의 생산성을 얼마나 향상시키고 있을까?](https://news.hada.io/topic?id=19672)
- [AI Blindspots – AI 코딩 중에 발견한 LLM의 맹점들](https://news.hada.io/topic?id=19859)
- [에이전틱 코딩에서 개발자 역량의 역할](https://news.hada.io/topic?id=20006)
- [AI 시대의 새로운 개발자 패턴들](https://news.hada.io/topic?id=21115)
- [에이전트와 함께 프로그래밍하는 방법](https://news.hada.io/topic?id=21418)
  - > LLM 기반 코드 생성 도구에 대한 흔한 비판 중 하나는 코드 생성 자체는 전체 소프트웨어 비용에서 극히 일부에 불과하다는 것임 (...) 그러나 대규모 유지보수 경험만을 전체 산업의 본질로 확대 해석하면 안 됨
  - > LLM 기반 에이전트는 주석과 설명을 코드 작성에 적극적으로 반영함
- [에이전틱 코딩 추천사항](https://news.hada.io/topic?id=21435)
- [TDD, AI Agents and Coding - 켄트백](https://news.hada.io/topic?id=21446)
- [생성형 AI 코딩 툴과 에이전트가 나에게 효과 없는 이유](https://news.hada.io/topic?id=21514)
- [AI가 오픈소스 개발자를 느리게 만든다. Peter Naur가 그 이유를 알려줄 수 있다](https://news.hada.io/topic?id=21996)
- [코딩 에이전트 체크리스트: Claude Code ver](https://speakerdeck.com/nacyot/koding-eijeonteu-cekeuriseuteu-claude-code-ver)
- [AI로 개발을 어떻게 가속화하는가](https://drive.google.com/file/d/1SJ7-1YXo4r4pkHDuMdKLR9NtgbUsSRoZ/view)
- [AI 시대의 지식 노동](https://wiki.g15e.com/pages/Knowledge%20work%20in%20the%20AI%20era)
- [나날이 발전하고픈 개발를 위한 AI 활용법](https://drive.google.com/file/d/1h99VB5Ra5nn78ZpcXzvN8HyJbSmcX-Qn/view?ref=stdy.blog)
- [Claude Code 6주 사용기](https://news.hada.io/topic?id=22375)
  - <https://blog.puzzmo.com/posts/2025/07/30/six-weeks-of-claude-code/>
  - > 과거에는 수주~수개월이 걸렸던 마이그레이션, 리팩토링, 기술 부채 해소 작업을 Claude Code 도입 이후 6주 만에 모두 병행·완료
  - > 기존의 "기술 부채는 일정 확보→대규모 투입" 공식이 깨지고, 즉석에서 시작→진행→완료까지 ‘즉각성’ 이 구현됨
  - > 실패 비용이 극적으로 낮아져, 실험-학습-결정의 사이클이 대폭 가속
- [Claude Code로 좋은 결과 얻기](https://news.hada.io/topic?id=22425)
- [AI에 대한 열풍 속에서 소프트웨어 엔지니어들은 어떻게 생각하고 있나요?](https://news.hada.io/topic?id=22489)
- [GenAI를 통한 소프트웨어 엔지니어링](https://news.hada.io/topic?id=22583)
