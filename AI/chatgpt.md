# ChatGPT

#chatgpt

## Work와 Codex 사용량

ChatGPT Plus 기준으로 웹의 Work와 Codex는 같은 ‘agentic usage’ 사용량을 공유합니다.

- 웹 Work를 많이 사용하면 Codex의 남은 사용량이 줄어듭니다.
- Codex를 많이 사용하면 Work에서 사용할 수 있는 양도 줄어듭니다.
- 추가 구매한 크레딧도 계정에 지원되는 경우 두 기능에서 공동으로 사용됩니다.
- 일반적인 빠른 대화인 **Chat/Quick Chat 사용량은 이 풀과 별도**입니다.

남은 사용량은 **Codex 설정 → Usage Dashboard**에서 확인할 수 있습니다. Work 작업의 소모량은 작업 복잡도, 입력 파일·컨텍스트 크기, 출력 길이에 따라 달라집니다.

- [ChatGPT Work와 Codex](https://help.openai.com/en/articles/20001275-chatgpt-work-and-codex)
- [Plus·Pro 크레딧 사용 방식](https://help.openai.com/en/articles/12642688-using-credits-for-flexible-usage-in-chatgpt-freegopluspro-sora)
- [Codex 요금·사용량 기준](https://help.openai.com/en/articles/20001106-codex-rate-card)

## GPT-5.6

<https://openai.com/index/gpt-5-6/>

### 모델 등급: Sol / Terra / Luna

[[chatgpt-gpt-5-6-models]]

| 모델      | 포지션             | API 가격 입력/출력 | 적합한 작업                       |
| --------- | ------------------ | ------------------ | --------------------------------- |
| **Sol**   | 최고 성능형        | $5 / $30           | 고난도 코딩·연구·장기 에이전트    |
| **Terra** | 성능·비용 균형형   | $2.5 / $15         | 일반적인 프로덕션 에이전트        |
| **Luna**  | 속도·비용 최적화형 | $1 / $6            | 대량 처리·반복 작업·간단한 자동화 |

가격 단위는 100만 토큰이다. Terra는 GPT-5.5급 성능을 더 낮은 비용으로 제공하며, Luna는 제품군 중 가장 빠르고 저렴하다. `Sol`, `Terra`, `Luna`는 독립적으로 발전하는 성능 등급이다.

### 추론 수준: `max`와 `ultra`

#### `max`

기존 `xhigh`보다 더 많은 시간과 연산을 사용해 다음을 수행한다.

- 여러 해결책 탐색
- 가정과 결과 검증
- 접근법 수정
- 최종 답변 재검토

즉, **단일 모델이 하나의 문제를 깊게 추론하는 모드**다. 복잡한 버그 분석, 아키텍처 설계, 수학·과학 문제에 적합하다.

#### `ultra`

기본적으로 **4개의 에이전트를 병렬 실행한 뒤 결과를 통합**한다. 토큰 사용량은 늘지만, 병렬화 가능한 작업의 품질을 높이고 완료 시간을 줄이는 데 초점을 둔다.

적합한 작업은 다음과 같다.

- 여러 출처를 동시에 조사하는 리서치
- 구현·테스트·보안 검토의 병렬 수행
- 여러 해결책의 독립 탐색과 비교
- 대규모 저장소의 영역별 분석

> **`max`** = 단일 모델의 심층 추론
> **`ultra`** = 다중 에이전트의 병렬 작업과 통합

`ultra`는 ChatGPT Work의 Pro·Enterprise와 Codex의 Plus 이상에서 제공된다.

### Programmatic Tool Calling

Programmatic Tool Calling은 모델이 **메모리 안에서 경량 프로그램을 실행해 여러 도구를 조정하는 기능**이다. Responses API에서 제공된다.

주요 기능은 다음과 같다.

- 여러 API와 도구의 호출 순서 제어
- 반복 호출과 조건 분기
- 중간 결과 필터링·집계
- 진행 상태에 따른 다음 행동 선택
- 필요한 데이터만 모델 문맥에 전달

#### 기존 방식과의 차이

```text
기존 방식
모델 → 도구 호출 → 전체 결과 반환
     → 모델 판단 → 다음 도구 호출
```

```text
Programmatic Tool Calling
모델 → 프로그램 작성
     → 여러 도구 호출·필터링·집계
     → 정제된 결과만 반환
```

이를 통해 다음 효과를 기대할 수 있다.

- 토큰 사용량 감소
- 모델과 도구 사이의 왕복 횟수 감소
- 복잡한 도구 흐름의 안정성 향상
- 장기 작업에서 사용자 개입 감소
