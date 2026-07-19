<https://www.npmjs.com/package/lefthook>

## [[agents#Claude Code]] Sandbox에서 Lefthook `stage_fixed` 실패 원인과 해결

### 1. 문제 요약

Claude Code 샌드박스 환경에서 Lefthook의 `stage_fixed: true`가 동작하지 않고 다음 계열의 오류가 발생했다.

```text
operation not permitted
```

처음에는 `stage_fixed` 내부의 `git add` 문제처럼 보였지만, 검증 결과 실제 원인은 `git add`가 아니라 **Lefthook이 command 실행 전에 PTY를 할당하려다 실패한 것**이었다.

### 2. Lefthook 옵션 동작 정리

#### 2.1 `stage_fixed: true`

`stage_fixed: true`는 `pre-commit` hook에서 formatter/linter가 파일을 수정한 뒤, 수정된 파일을 다시 stage하기 위한 옵션이다.

개념적으로는 다음 흐름이다.

```text
1. staged files 확인
2. command 실행
3. command가 파일 수정
4. Lefthook이 수정 파일을 git add
5. commit 계속 진행
```

이번 이슈에서는 `stage_fixed` 단계까지 도달하지 못했다.
따라서 `stage_fixed` 자체나 `git add`는 근본 원인이 아니었다.

#### 2.2 `no_tty: true`

`no_tty: true`는 Lefthook의 TTY 기반 UI/interactive 동작을 줄이는 옵션이다.

```yml
no_tty: true
```

주요 효과는 다음과 같다.

- spinner/progress UI 비활성화
- Lefthook의 interactive 처리 억제
- `interactive: true` 무시
- command 실행 방식 자체를 항상 pipe 기반으로 바꾸는 옵션은 아님
- child command 실행에 필요한 PTY 할당 문제를 우회하지 못할 수 있음

이번 Claude Code sandbox 이슈에서는 `no_tty: true` 또는 `--no-tty`를 적용해도 실패했다.

이유는 문제의 핵심이 “화면 UI 출력”이 아니라, Lefthook이 command 실행 과정에서 기본적으로 PTY를 만들려고 하는 동작이었기 때문이다.

#### 2.3 `interactive: true`

`interactive: true`는 command가 `/dev/tty`를 사용할 수 있도록 interactive 실행을 허용하는 옵션이다.

```yml
pre-commit:
  jobs:
    - run: some-command
      interactive: true
```

Claude Code sandbox에서는 `/dev/ptmx` 접근이 차단되어 있으므로, 이 옵션은 해결책이 아니다.
오히려 PTY/TTY 의존성을 강화하는 방향이라 실패를 해결하지 못한다.

#### 2.4 `use_stdin: true`

`use_stdin: true`는 Lefthook이 받은 stdin을 command/script에 전달하도록 하는 옵션이다.

```yml
pre-commit:
  jobs:
    - run: pnpm prettier --write {staged_files}
      use_stdin: true
      stage_fixed: true
```

일반적으로는 `pre-push`처럼 Git hook이 stdin으로 ref 정보를 넘기는 경우에 사용한다.

하지만 이번 환경에서는 부수적으로 중요한 효과가 있었다.

```text
use_stdin: true
→ Lefthook이 child stdin을 PTY 대신 pipe로 연결
→ /dev/ptmx 접근 없음
→ Claude Code sandbox의 PTY 차단을 우회
```

즉, 이번 문제에서는 `use_stdin: true`가 stdin 전달 옵션이면서 동시에 **PTY 할당 우회 수단**으로 작동했다.

### 3. 근본 원인

검증된 근본 원인은 다음과 같다.

```text
Claude Code sandbox
→ macOS seatbelt sandbox
→ /dev/ptmx 접근 차단
→ PTY master 생성 실패
→ Lefthook command 실행 실패
→ formatter/linter 실행 전 종료
→ stage_fixed 단계 미도달
```

Lefthook은 기본적으로 command를 PTY 안에서 실행한다.
이 과정에서 `creack/pty`를 통해 `/dev/ptmx`에 접근한다.

Claude Code sandbox에서는 이 접근이 `EPERM`으로 차단된다.

따라서 실제 실패 지점은 다음이다.

```text
prettier/eslint 실행
X stage_fixed
X git add
```

`prettier`나 `eslint`가 잘못된 것도 아니고, `git add`가 실패한 것도 아니다.
Lefthook이 command를 실행하기 위한 PTY를 만들지 못해 command 자체가 시작되지 못한 것이다.

### 4. 핵심 증거

#### 4.1 샌드박스 OFF 시 정상 동작

Claude Code sandbox를 끄면 다음이 모두 정상 동작했다.

```text
prettier/eslint 실행
stage_fixed
git add --force
commit 진행
```

따라서 Git write 권한이나 `.git/objects` 접근 문제는 아니었다.

#### 4.2 `/dev/ptmx` 직접 접근 실패

샌드박스 내부에서 `/dev/ptmx`를 직접 열면 다음 오류가 발생했다.

```text
Operation not permitted
```

이는 Lefthook에서 발생한 오류와 동일했다.

#### 4.3 seatbelt profile 실험

seatbelt profile에서 다음 장치를 허용하면 PTY 생성이 성공했다.

```text
/dev/ptmx
/dev/ttys*
```

따라서 실패 원인은 sandbox의 PTY 관련 device 접근 차단으로 확정할 수 있다.

### 5. 적용한 해결책

최종 해결책은 `format`과 `lint` job에 `use_stdin: true`를 추가하는 것이다.

```yml
pre-commit:
  jobs:
    - name: format
      run: pnpm prettier --write {staged_files}
      use_stdin: true
      stage_fixed: true

    - name: lint
      run: pnpm eslint --fix {staged_files}
      use_stdin: true
      stage_fixed: true
```

이 설정을 적용하면 Lefthook이 child process의 stdin을 PTY 대신 pipe로 연결한다.
그 결과 `/dev/ptmx`에 접근하지 않으므로 Claude Code sandbox에서도 command 실행이 가능해진다.

### 6. 실패한 해결 시도

다음 시도들은 효과가 없었다.

```text
no_tty: true
--no-tty
interactive: true
--colors off
```

이 옵션들은 Lefthook 출력 UI나 interactive 동작에는 영향을 줄 수 있지만, 이번 문제의 핵심인 **child command 실행을 위한 PTY 생성 실패**를 해결하지 못했다.

특히 `no_tty: true`는 “Lefthook UI를 조용하게 만드는 옵션”에 가깝고, 이번 환경에서 필요한 “PTY 할당 자체를 피하는 옵션”으로는 충분하지 않았다.

### 7. 회귀 검증 결과

`use_stdin: true` 적용 전후를 동일 입력으로 비교했다.

검증 결과는 다음과 같다.

```text
exit code 동일
stage_fixed 동작 동일
formatter 출력 동일
lint 실패 전파 정상
다중 파일 처리 정상
hang 없음
```

즉, `use_stdin: true`를 추가해도 `prettier`, `eslint`, `stage_fixed`의 정상 동작에는 영향이 없었다.

### 8. `use_stdin`의 “stdin은 하나만 받는다” 주의사항

Lefthook 문서상 `use_stdin: true`를 여러 command/script에 붙이면 stdin 스트림은 하나의 command만 소비할 수 있다.

하지만 이번 `pre-commit`의 `prettier`/`eslint` 사례에서는 문제가 되지 않는다.

이유는 다음과 같다.

```text
prettier/eslint는 hook stdin을 읽는 command가 아님
pre-commit은 pre-push처럼 ref 정보를 stdin으로 받는 hook이 아님
이번 use_stdin은 stdin 데이터 소비 목적이 아니라 PTY 우회 목적으로 사용됨
```

따라서 이번 설정에서는 “only one receives stdin” 문제가 실질적으로 발생하지 않았다.

### 9. 부수 경고: Lefthook hook 재동기화

설정 변경 후 Lefthook이 `.git/hooks/pre-commit` 재동기화를 시도하면서 경고를 낼 수 있다.

하지만 이 경고는 커밋을 막지 않는다.

```text
exit code 0
commit 진행 가능
```

경고를 없애려면 Claude Code sandbox 밖에서 다음 명령을 한 번 실행하면 된다.

```bash
pnpm exec lefthook install
```

### 10. 대안

#### 10.1 sandbox device 허용

`~/.claude/settings.json`의 `sandbox.filesystem.allowWrite`에 다음 경로를 추가한다.

```text
/dev/ptmx
/dev/ttys*
```

이 방식은 PTY 생성을 허용하므로 근본 원인을 직접 우회할 수 있다.
다만 sandbox 허용 범위를 넓히는 방식이므로 최소 변경 원칙에는 맞지 않는다.

#### 10.2 `git commit*` sandbox 제외

`excludedCommands`에 다음을 추가한다.

```text
git commit*
```

이 방식은 commit 자체를 sandbox 밖에서 실행하게 만든다.
하지만 Lefthook뿐 아니라 commit 과정 전체를 sandbox 밖으로 빼는 방식이므로 영향 범위가 크다.

### 11. 최종 결론

이번 문제의 원인은 `stage_fixed`나 `git add`가 아니라 Claude Code sandbox에서 PTY 생성이 차단된 것이다.

정확한 실패 흐름은 다음과 같다.

```text
Lefthook command 실행
→ PTY 생성 시도
→ /dev/ptmx 접근
→ sandbox가 EPERM으로 차단
→ prettier/eslint 실행 전 실패
→ stage_fixed 미도달
```

최소 변경으로 해결하려면 `stage_fixed`를 쓰는 `format`/`lint` job에 다음 옵션을 추가한다.

```yml
use_stdin: true
```

이 설정은 Lefthook의 child stdin 연결을 PTY 대신 pipe 기반으로 바꾸어 `/dev/ptmx` 접근을 피하게 만든다.

따라서 Claude Code sandbox 환경에서는 다음 구성이 가장 안전한 해결책이다.

```yml
pre-commit:
  jobs:
    - name: format
      run: pnpm prettier --write {staged_files}
      use_stdin: true
      stage_fixed: true

    - name: lint
      run: pnpm eslint --fix {staged_files}
      use_stdin: true
      stage_fixed: true
```

## pre-push hook

1. **`--changed HEAD`를 그대로 옮기면 게이트가 무력화된다** — push 시점엔 변경이 이미 HEAD에 포함돼 잡히는 게 0개 → `--passWithNoTests`로 항상 통과. pre-push의 base는 "원격에 아직 없는 커밋 범위"여야 한다 → `@{push}`.
2. **`@{push}` 단독은 새 브랜치에서 push를 막는다** — upstream 없으면 `git rev-parse @{push}`가 **exit 128**. 첫 push(`-u`)는 upstream 생기기 전에 훅이 도므로 차단됨(닭-달걀). → **폴백 필수**.
3. **git은 브랜치의 base(부모)를 저장하지 않는다** — base 자동검출은 reflog 휴리스틱뿐이라 신뢰 불가(CI·새 클론서 깨짐). 고정 폴백이 현실적. lefthook도 동일 문제(#785)를 **아직 미해결**.
4. **`{push_files}`보다 `--changed`를 택한 이유**
   - `{push_files}`+glob은 `package.json`만 바뀐 push에서 test를 **skip** → vitest `forceRerunTriggers`(의존성 변경 시 전체 실행) 무력화
   - `{push_files}`는 폴백 base를 **제어 불가**(origin/HEAD 고정, #785)
   - → `--changed`는 vitest가 변경을 직접 판단하고, 폴백 base를 우리가 명시 가능
5. **glob 제거** — lefthook이 미리 거르지 않게 해서 vitest `--changed`가 변경 판단을 전담.
6. **남는 한계(공통)** — 테스트 파일만 수정 시 누락 가능(소스 기준) → CI 전체 실행으로 보완. `origin/develop` 로컬 최신 유지 필요.

```yaml
pre-push:
  commands:
    test:
      run: pnpm run test -- --changed "$(git rev-parse @{push} 2>/dev/null || echo origin/develop)" --passWithNoTests
```

### references

- <https://github.com/evilmartians/lefthook/issues/568> (upstream 없을 때 exit 128)
- <https://github.com/evilmartians/lefthook/issues/785> (push_files base 지정 미지원, OPEN)
- <https://github.com/evilmartians/lefthook/issues/302> (새 브랜치 첫 push 전체 실행 + 폴백 논의)
- <https://github.com/vitest-dev/vitest/issues/1113> (`--changed`가 테스트 파일 단독 변경 무시)
