# Agentic Coding

- <https://en.wikipedia.org/wiki/Vibe_coding>
- [[agents]]
- [[mcp]]
- [[skills]]

## 실천방법

> 제작사별 모델마다 프롬프팅 방식에 차이가 있고, 신규모델이 나올 때마다 달라지므로 Learn & Unlearn 중요

- [커스텀 룰](#agentsmd) 사용
  - 커스텀 룰은 지속적으로 개선
  - 하위 디렉토리에서 추가 규칙이 필요하다면 디렉토리별 커스텀 룰 사용
  - 그러나 코딩 에이전트들은 커스텀 룰을 제대로 따르지 않는 경우도 많기 때문에 중요한 지시사항이라면 프롬프트에 포함
- 한 세션을 불필요하게 길게 유지하지 않음
- 명확한 요구사항 전달, 활용할 기존 코드베이스 참조 제공
  - > Claude Code의 성공률은 특히 첫 시도에서 더 구체적인 지침을 제공할 때 크게 향상됩니다
- 작업을 시작하기 전 모델에게 아이디어를 구상하거나 계획을 세우도록 요청 (plan mode)
  - 아이디어, 계획을 피드백 하면서 방향을 맞춤
  - 계획을 세우기에 context 불충분 한 경우 재질문 하도록 프롬프트 (대부분의 agent plan mode 내장기능임)
  - 계획의 적용범위가 부족하다고 생각될 때는 "추가로 고려해야 할 사항"이 있는지 역질문
  - 계획은 대형 모델(claude opus ...)로 진행하고, 코드 작성은 일반 모델(claude sonnet ...)으로 진행하는 것도 방법
  - 계획이 어느정도 정리되면 파일로 저장하도록 요청하고 본격적인 실행 전 다듬는 작업 진행
  - 계획으로부터 Task (todo list) 생성하고 해결 여부 표시하도록 요청 (최신 모델은 별도의 지시 없이도 단계별 Task 생성함)
- 계획을 세웠다면 프롬프트에서 지시를 정확히 수행하도록 요구 (PLAN-TASK file/context 제공)
- 실행 중 원하는 코드 방향이 아니면 계획단계로 돌아 가는 것이 나을 수도 있음 (계획-실행-반복수정실행 사이의 컨텍스트 불일치로 전체 품질 저하)
- 분리할 수 있는 작업이면서 결과만 사용해도 품질 저하가 적은 경우 subagent 사용해서 작업 분리 (context 관리)
- feedback loop (검증 조건, 빌드/테스트가 작동할 때 까지 코드를 수정해줘)
- 각 단계가 끝날 때 커밋 요청, git-hooks를 섞어 쓰면 타입오류, lint, 테스트 ... 적용할 수 있음
- 여러 모델을 섞어 구현-리뷰 교차실행
- 모델별로 차이는 있으나 context window 사용량이 40~50% 초과할 때부터 품질 저하 발생할 수 있음 (opus-4.5, gpt-5.2 이후 모델은 품질 저하 개선됨)
- AI에게는 오히려 코드 중간중간의 주석이 도움이 될 때가 있음 (why, what, how 모두)

### Prompt Engineering

#### claude

최근의 모델은 복잡한 프롬프팅이 필요하지 않음. 필수 지침만 최소로 관리.

- 명시적이고 명확하게 지시: 모델이 추론할 것이라고 가정하지 말고, 원하는 바를 직접적이고 명확하게 진술
  - 서문을 건너뛰고 바로 요청
  - 작업할 내용뿐만 아니라 결과물에 포함되기를 원하는 것을 명시
  - 품질 및 깊이 기대치에 대해 구체적으로 명시
- 맥락과 동기 부여 제공: 왜 이 작업이 중요한지 설명하면 모델이 목표를 더 잘 이해하고 더 목적에 맞는 응답을 제공하는 데 도움이 됨
- 구체적으로 작성
  - 명확한 제약 조건 (단어 수, 형식, 일정)
  - 관련 컨텍스트 (대상은 누구인지, 목표는 무엇인지)
  - 원하는 결과물 구조 (표, 목록, 문단)
  - 모든 요구 사항 또는 제한 사항 (식단 요구 사항, 예산 제한, 기술적 제약)
- 예시는 항상 필요한 것은 아니지만, 개념을 설명하거나 특정 형식을 시연할 때 효과적, 다음은 예시를 사용해야 할 때의 조건임
  - 원하는 형식을 설명하는 것보다 보여주는 것이 더 쉬울 때
  - 특정한 어조나 스타일이 필요할 때
  - 작업에 미묘한 패턴이나 관례가 포함될 때
  - 간단한 지침으로 일관된 결과가 나오지 않았을 때
- 불확실성을 표현할 권한을 부여
  - 추측 대신 불확실성을 표현할 수 있는 명시적인 권한을 AI에게 부여하세요. 이는 환각(hallucination)을 줄이고 신뢰도를 높입니다.
  - 예시: "이 재무 데이터를 분석하고 추세를 파악해 줘. 결론을 도출하기에 데이터가 불충분하다면, 추측하는 대신 그렇게 말해 줘."

이전 모델에서 인기 있었던 일부 기법은 최신 모델에서 덜 필요함

- 최신 모델은 XML 태그 없이도 구조를 잘 이해함
  - 여러 콘텐츠가 혼합된 복잡한 프롬프트, 콘텐츠 경계 확신이 필요한 경우 에는 사용
  - 명확한 제목, 공백, 명시적 언어 (아래의 정보를 사용하여 ...)로 대체
- 역할 프롬프팅
  - 최신 모델은 과도한 역할 프롬프팅이 불필요 (출력에 일관된 톤 유지, 복잡한 도메인 전문지식의 틀 유지 가 필요한 경우 사용)
  - 역할을 할당하는 대신 원하는 관점이 무엇인지 명시적으로 밝히는 것으로 대체

[추가내용](./agentic-coding/claude-best-practices-for-prompt-engineering.md)

refs

- <https://docs.anthropic.com/ko/docs/build-with-claude/prompt-engineering/overview>
- <https://platform.claude.com/docs/ko/build-with-claude/prompt-engineering/claude-4-best-practices>

#### codex

- (codex-max) rollout(긴 에이전트 실행) 중에 "upfront plan, preambles, status updates"(계획 설명, 진행상황 보고) 를 말하게 하는 프롬프트는 제거

refs

- <https://cookbook.openai.com/examples/gpt-5-codex_prompting_guide>
- <https://cookbook.openai.com/examples/gpt-5/gpt-5-1-codex-max_prompting_guide>

### AGENTS.md

- <https://agents.md/>
- [[claude-code-best-practices#CLAUDE.md]]
- <https://www.claude.com/blog/using-claude-md-files>
- <https://www.humanlayer.dev/blog/writing-a-good-claude-md>

> 커스텀 룰, 프로젝트 메모리

목적

- 지속적인 컨텍스트 제공: 코드베이스 구조, 코딩 표준, 선호 워크플로우를 모델에게 알려줌
- 반복 작업 감소: 모든 대화의 시작마다 프로젝트 기본 정보를 설명하지 않아도 됨

best practices

- 최소한의 내용으로 유지 (100줄 정도)
- 명확한 지침 (`DO`, `DON'T`), 그러나 지침은 최소화
- 주제별 세부사항은 적당한 개수로 나열 (3~4개 이내)
- 세부 지침이 필요한 상황이라면 별도의 파일로 분리하고, 파일위치와 설명 첨부
- 주요 구성
  - 프로젝트 요약, 아키텍처 패턴, 주요 라이브러리
  - 프로젝트 구성 요소가 위치한 핵심 디렉토리 구조
  - 표준 및 규칙 (linter 규칙, 코딩스타일 등은 제외)
  - 개발 서버 실행, 테스트 실행 등 자주 사용하는 명령어
  - 배포, 테스트 등 팀의 사용자 정의 도구 및 스크립트 사용 방법, MCP 등
  - 표준 워크플로우: 코드 변경 전 따라야 할 계획 수립, 테스트, 커밋 형식 등의 단계
  - 오답노트 (반복해서 잘못 작동하는 행동에 대한 지침)

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

### [Claude Code: Best practices for agentic coding - anthropic](./agentic-coding/claude-code-best-practices.md)

- 설정 사용자 정의 (Customize your setup)
- Claude에게 더 많은 도구 제공 (Give Claude more tools)
- 일반적인 워크플로우 시도 (Try common workflows)
- 워크플로우 최적화 (Optimize your workflow)
- 헤드리스 모드를 사용해 인프라 자동화 (Use headless mode to automate your infra)
- 다중 Claude 워크플로우로 수준 향상 (Uplevel with multi-Claude workflows)

### [Here’s how I use LLMs to help me write code](https://simonwillison.net/2025/Mar/11/using-llms-for-code/)

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

### [Vibe 코딩은 저품질 작업에 대한 변명이 될 수 없어요](./agentic-coding/vibe-coding-is-not-an-excuse-for-low-quality-work.md)

- Rule 1: Always Review AI-Generated Code / AI 코드 반드시 리뷰하기
- Rule 2: Establish Coding Standards and Follow Them / 코딩 스타일과 기준을 설정하고 준수할 것
- Rule 3: Use AI for Acceleration, Not Autopilot / AI는 가속기이지 자동 조종 장치가 아님
- Rule 4: Test, Test, Test / 테스트는 무조건 해야 한다
- Rule 5: Iterate and Refine / 반복하고 다듬기
- Rule 6: Know When to Say No / 거절할 줄 알아야 함
- Rule 7: Document and Share Knowledge / 문서화하고 지식을 공유하라

### [Vibe Coding 매뉴얼: AI 지원 개발을 위한 템플릿](./agentic-coding/vibe-coding-manual.md)

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

## Context Engineering

- <https://github.com/davidkimai/Context-Engineering>
- [바이브 코딩에는 컨텍스트 엔지니어링이 필요하다](https://blogbyash.com/translation/vibe-coding-needs-context-engineering/)
- <https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents>
- [복잡한 코드베이스에서 AI를 제대로 작동하게 만드는 법; "research, plan, implement" workflow](https://news.hada.io/topic?id=23257)
  - [No Vibes Allowed: Solving Hard Problems in Complex Codebases – Dex Horthy, HumanLayer](https://www.youtube.com/watch?v=rmvDxxNubIg)
  - <https://github.com/ai-that-works/ai-that-works/tree/main/2025-08-05-advanced-context-engineering-for-coding-agents>
- <https://github.com/muratcankoylan/Agent-Skills-for-Context-Engineering>

## Spec-Driven-Development

- <https://github.com/github/spec-kit>
- <https://martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html>
  - <https://news.hada.io/topic?id=23776>
- [스펙 주도 개발(SDD): 워터폴의 귀환](https://news.hada.io/topic?id=24400)
- <https://developers.googleblog.com/conductor-introducing-context-driven-development-for-gemini-cli/>

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
- [프롬프팅 101](https://www.youtube.com/watch?v=ysPbXH0LpIE)
- [Improve your AI code output with AGENTS.md](https://www.builder.io/blog/agents-md)
- [에이전트 루프 설계하기 (simonwillison.net)](https://news.hada.io/topic?id=23470)
- [코딩을 위한 효과적인 하위 에이전트를 활성화하는 방법](https://www.youtube.com/watch?v=jxDy33IqtjI)
- [AI가 이 코드의 작동 방식을 깊이 이해하고 있다](https://news.hada.io/topic?id=24637)
- [Claude Code의 모든 기능 활용법](https://news.hada.io/topic?id=24099)
- [AI를 활용한 프로그래밍 역량을 높이는 방법](https://news.hada.io/topic?id=25060)
- [Claude Code를 활용한 예측 가능한 바이브 코딩 전략](https://helloworld.kurly.com/blog/vibe-coding-with-claude-code/)
- [My Current AI Dev Workflow | Peter Steinberger](https://steipete.me/posts/2025/optimal-ai-development-workflow)
- [Just Talk To It - the no-bs Way of Agentic Engineering | Peter Steinberger](https://steipete.me/posts/just-talk-to-it)

## Tools

- <https://github.com/steveyegge/beads> A memory upgrade for your coding agent
- <https://github.com/yamadashy/repomix> packs your entire repository into a single, AI-friendly file
- <https://github.com/Ryandonofrio3/osgrep> Open Source Semantic Search for your AI Agent

### git-worktree-runner

<https://github.com/coderabbitai/git-worktree-runner>

```bash
git clone https://github.com/coderabbitai/git-worktree-runner.git ~/.git-worktree-runner
cd ~/.git-worktree-runner
ln -sfn "$(pwd)/bin/git-gtr" ~/.local/bin/git-gtr
mkdir -p ~/.oh-my-zsh/completions && ln -sfn "$(pwd)/completions/_git-gtr" ~/.oh-my-zsh/completions/_git-gtr
```
