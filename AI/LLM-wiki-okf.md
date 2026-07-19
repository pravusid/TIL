# LLM wiki, OKF

Google Cloud 블로그의 핵심 주장은 **AI/LLM 에이전트의 성능 병목이 모델 자체보다 “조직 내부 지식의 맥락 부족”에 있다**는 것이다.
조직의 지식은 데이터 카탈로그, 위키, 문서, 코드, 노트북, 런북, 담당자의 암묵지 등으로 흩어져 있고, 이 때문에 LLM 에이전트가 정확한 답을 만들기 어렵다.

이를 해결하기 위해 Google Cloud는 **Open Knowledge Format, OKF**를 제안한다. OKF는 새로운 지식관리 플랫폼이 아니라, 조직 지식을 **사람도 읽고 LLM도 처리하기 쉬운 파일 포맷**으로 정리하기 위한 개방형 규약이다.

## Open Knowledge Format, OKF란?

OKF는 데이터셋, 테이블, 메트릭, API, 런북, 비즈니스 규칙 같은 조직 지식을 **Markdown 파일 + YAML frontmatter + 디렉터리 구조 + Markdown 링크**로 표현하는 포맷이다.

예를 들어 `weekly_active_users.md`, `payments.md`, `refund-policy.md`, `cloudfront-incident-runbook.md` 같은 파일을 만들고, 각 파일 상단에 다음과 같은 메타데이터를 둔다.

```md id="b6xk2p"
---
type: Metric
title: Weekly Active Users
description: 주간 활성 사용자를 계산하는 사내 표준 메트릭
tags: [analytics, product, metric]
---
```

본문에는 정의, 계산 방식, 관련 테이블, 주의사항, 예시 쿼리, 관련 문서 링크 등을 Markdown으로 작성한다.

OKF의 핵심은 **format, not platform**이다. 특정 클라우드, 데이터베이스, LLM, 벡터DB, 문서도구에 종속되지 않고, Git repo나 파일시스템처럼 단순한 구조 위에서 작동한다. 그래서 사람이 직접 관리할 수도 있고, 에이전트가 생성·수정할 수도 있으며, 검색 도구나 시각화 도구가 읽을 수도 있다.

## LLM Wiki란?

LLM Wiki는 LLM이 원본 자료를 읽고, 그 내용을 기반으로 **지속적으로 갱신되는 Markdown 기반 지식베이스를 만들어 가는 방식**이다.

기존 RAG는 보통 질문이 들어올 때마다 원본 문서에서 관련 chunk를 검색해 답변을 만든다. 반면 LLM Wiki는 원본 자료를 한 번 읽고 끝내는 것이 아니라, 요약 페이지, 개념 페이지, 엔티티 페이지, 비교 문서, 링크 구조를 계속 업데이트한다.

즉, LLM Wiki는 다음과 같은 구조를 가진다.

| 계층                 | 역할                                               |
| -------------------- | -------------------------------------------------- |
| Raw sources          | 원본 자료. 문서, 논문, 로그, 데이터, 코드 등       |
| Wiki                 | LLM이 생성·관리하는 Markdown 지식베이스            |
| Instruction / Schema | LLM에게 위키 작성 규칙과 관리 방식을 알려주는 파일 |

LLM Wiki의 중요한 작업은 세 가지다.

1. **Ingest**: 새 자료를 읽고 관련 위키 페이지를 생성·수정한다.
2. **Query**: 질문이 들어오면 위키를 기반으로 답변한다.
3. **Lint**: 오래된 정보, 모순, 누락된 링크, 고립된 문서를 점검한다.

## OKF와 LLM Wiki의 관계

LLM Wiki가 **“LLM이 관리하는 살아있는 지식베이스 패턴”**이라면, OKF는 그런 지식베이스를 여러 도구와 조직이 함께 사용할 수 있게 만드는 **공통 파일 포맷**이다.

| 구분      | LLM Wiki                      | OKF                          |
| --------- | ----------------------------- | ---------------------------- |
| 성격      | 지식관리 패턴                 | 지식 표현 포맷               |
| 목적      | LLM이 지식을 누적·갱신        | 지식을 교환 가능하게 표준화  |
| 저장 방식 | Markdown 위키                 | Markdown + YAML frontmatter  |
| 핵심 가치 | 지속적 학습, 연결, 갱신       | 상호운용성, 이식성, 단순성   |
| 비유      | LLM이 관리하는 Obsidian vault | 여러 vault를 호환시키는 규격 |

## 실무적 의미

OKF와 LLM Wiki를 함께 사용하면 사내 지식을 LLM 친화적으로 정리할 수 있다. 예를 들어 데이터팀의 메트릭 정의, 운영팀의 장애 런북, 개발팀의 API 문서, 비즈니스팀의 정책 문서를 하나의 Markdown 기반 knowledge bundle로 만들 수 있다.

이렇게 하면 LLM 에이전트는 단순히 문서를 검색하는 수준을 넘어, **메트릭 정의 → 관련 테이블 → 예시 쿼리 → 장애 대응 런북 → 담당 시스템**처럼 연결된 지식을 따라가며 더 정확한 답변과 작업을 수행할 수 있다.

결론적으로 **LLM Wiki는 지식을 살아 있게 만드는 운영 방식이고, OKF는 그 지식을 도구와 조직 사이에서 이동 가능하게 만드는 포맷**이다.

## References

[1] Google Cloud Blog, “How the Open Knowledge Format can improve data sharing”
<https://cloud.google.com/blog/products/data-analytics/how-the-open-knowledge-format-can-improve-data-sharing/?hl=en>

[2] GoogleCloudPlatform / knowledge-catalog, “Open Knowledge Format v0.1 Specification”
<https://github.com/GoogleCloudPlatform/knowledge-catalog/blob/main/okf/SPEC.md>

[3] Andrej Karpathy, “llm-wiki”, GitHub Gist
<https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f>
