# Claude Code: Best practices for agentic coding

<https://www.anthropic.com/engineering/claude-code-best-practices>

## CLAUDE.md

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

## Try common workflows

주요 내용은 Claude Code를 사용하여 코딩 작업을 수행하는 세 가지 효과적인 워크플로우입니다. Claude는 유연하게 사용할 수 있지만, 다음 방법들이 특히 유용하다는 것이 밝혀졌습니다.

### 탐색, 계획, 코딩, 커밋 (Explore, plan, code, commit)

이 워크플로우는 대부분의 코딩 문제에 적합합니다.

- 탐색/조사: Claude에게 관련 파일(예: logging.py), 이미지 또는 URL을 읽도록 요청합니다. 이때는 아직 코드를 작성하지 않도록 명확하게 지시합니다. 복잡한 문제의 경우 서브 에이전트를 사용하여 세부 정보를 확인하거나 질문을 조사하도록 하면 컨텍스트를 유지하는 데 도움이 됩니다.
  - 예시: "Claude, user_authentication.py 파일과 database_schema.sql 파일을 읽어줘. 아직 코드는 작성하지 마."
- 계획: Claude에게 문제 해결을 위한 계획을 세우도록 요청합니다. "think", "think hard", "ultrathink"와 같은 단어를 사용하여 Claude가 더 많은 생각 시간을 갖도록 지시합니다. 계획이 만족스러우면 Claude에게 해당 계획을 문서나 GitHub 이슈로 만들도록 요청할 수 있습니다.
  - 예시: "Claude, 사용자 로그인 기능 구현에 대한 계획을 세워줘. think hard."
- 코딩: Claude에게 계획에 따라 솔루션을 코드로 구현하도록 요청합니다. 코드를 구현하면서 합리적인지 확인하도록 요청할 수 있습니다.
  - 예시: "이제 계획대로 사용자 로그인 기능을 코드로 구현해줘."
- 커밋: 결과물을 커밋하고 Pull Request를 생성하도록 요청합니다. 관련이 있다면 README나 변경 로그를 업데이트하도록 요청할 수도 있습니다.
  - 예시: "코드 변경 사항을 커밋하고 Pull Request를 생성해줘. 그리고 README.md도 업데이트해줘."

### 테스트 작성, 커밋; 코딩, 반복, 커밋 (Write tests, commit; code, iterate, commit)

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

### 코드 작성, 결과 스크린샷, 반복 (Write code, screenshot result, iterate)

이 워크플로우는 시각적인 결과가 중요한 경우에 유용합니다.

- 스크린샷 도구 제공: Claude가 브라우저 스크린샷을 찍을 수 있는 방법(예: Puppeteer, iOS 시뮬레이터, 수동 복사/붙여넣기)을 제공합니다.
  - 예시: "Claude, 이 Puppeteer 서버를 사용해서 웹페이지 스크린샷을 찍을 수 있어."
- 시각적 Mock 제공: 이미지 파일을 제공하거나, 이미지를 복사/붙여넣기하여 Claude에게 시각적 Mock을 제공합니다.
  - 예시: "이 이미지는 우리가 만들고자 하는 로그인 페이지의 디자인이야."
- 코드 구현 및 반복: Claude에게 코드를 구현하고, 결과 스크린샷을 찍고, 결과가 Mock과 일치할 때까지 반복하도록 요청합니다.
  - 예시: "이 디자인을 HTML/CSS로 구현하고, 매번 스크린샷을 찍어서 디자인과 일치하는지 확인해줘."
- 커밋: 결과가 만족스러우면 커밋하도록 요청합니다.
  - 예시: "디자인과 정확히 일치하니 코드를 커밋해줘."

## 워크플로우 최적화

Claude Code의 성공률은 특히 첫 시도에서 더 구체적인 지침을 제공할 때 크게 향상됩니다.
명확한 방향을 미리 제시하면 나중에 수정할 필요성이 줄어듭니다.

- foo.py에 대한 테스트 추가
  - 좋은 예: 사용자가 로그아웃된 엣지 케이스를 다루는 새로운 테스트 케이스를 foo.py에 작성하세요. 모의(mock) 사용은 피하세요.
- ExecutionFactory의 API는 왜 이렇게 이상한가요?
  - 좋은 예: ExecutionFactory의 git 히스토리를 살펴보고 API가 어떻게 현재의 모습을 갖추게 되었는지 요약해 주세요.
- 캘린더 위젯 추가
  - 홈페이지에 기존 위젯이 어떻게 구현되어 있는지 살펴보고 패턴, 특히 코드와 인터페이스가 어떻게 분리되어 있는지 파악하세요.
    HotDogWidget.php가 시작하기 좋은 예입니다. 그런 다음 그 패턴에 따라 사용자가 월을 선택하고 앞/뒤로 페이지를 넘겨 연도를 선택할 수 있는 새로운 캘린더 위젯을 구현하세요.
    코드베이스의 나머지 부분에서 이미 사용 중인 라이브러리 외에 다른 라이브러리 없이 처음부터 직접 만드세요.
