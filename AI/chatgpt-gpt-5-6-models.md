# GPT-5.6 패밀리 벤치마크·비용 효율·Driver 추천 보고서

#chatgpt

## 1. 요약 결론

### 1.1 작업 유형별 추천

| 작업 유형                                | 데일리 드라이버 | 어려운 작업용  | 핵심 근거                                                                   |
| ---------------------------------------- | --------------- | -------------- | --------------------------------------------------------------------------- |
| 일반 코딩·리팩터링·디버깅                | **Sol Medium**  | **Sol High**   | CursorBench와 DeepSWE에서 Medium이 비용 변곡점, High가 현실적인 고난도 상한 |
| GitHub PR 리뷰·코드 리뷰 봇              | **Terra High**  | **Terra High** | 실사용 사례에서 약 40% 빠르고 품질 저하가 거의 없음; XHigh는 리뷰 개선이 미미함 |
| 대규모 저장소·장기 에이전트 코딩         | **Sol Medium**  | **Sol High**   | DeepSWE에서 High 이후 XHigh·Max의 추가 비용 대비 성공률 상승이 작음         |
| 간단한 코드 수정·테스트·대량 반복 코딩   | **Luna High**   | **Sol Medium** | 저비용 반복 실행은 Luna가 유리하고, 실패 항목만 Sol로 승격 가능             |
| 문서·리서치·보고서·슬라이드·스프레드시트 | **Luna Max**    | **Sol XHigh**  | GDPval에서 Luna Max가 가장 강한 실무 가성비, 최고 품질 구간은 Sol XHigh     |
| 제품 기획·아키텍처·복합 분석             | **Sol Medium**  | **Sol XHigh**  | Intelligence Index에서 Sol이 범용 추론 우위, Max는 한계효율 낮음            |
| 요약·분류·변환·대량 콘텐츠 처리          | **Luna High**   | **Luna Max**   | 절대 비용이 낮고 High까지 효율이 좋으며 Max도 저렴함                        |
| 최고 품질의 최종 산출물                  | **Sol High**    | **Sol XHigh**  | XHigh가 실질적인 프리미엄 상한, Max는 비용 증가 대비 점수 상승이 작음       |
| 실패 비용이 매우 큰 1회성 작업           | —               | **Sol Max**    | 가성비보다는 절대 최고점이 중요한 예외적 상황                               |

### 1.2 가장 단순한 운영안

| 역할              | 권장 설정      |
| ----------------- | -------------- |
| 범용 Daily driver | **Sol Medium** |
| 범용 Hard driver  | **Sol XHigh**  |

### 1.3 가장 비용 효율적인 운영안

| 역할                      | 권장 설정     |
| ------------------------- | ------------- |
| 대량·간단 작업            | **Luna High** |
| 일반 지식작업             | **Luna Max**  |
| 어려운 코딩               | **Sol High**  |
| 어려운 지식작업·복합 추론 | **Sol XHigh** |
| 최후 에스컬레이션         | **Sol Max**   |

---

## 2. 분석 방법

인접한 추론 수준으로 한 단계 올렸을 때의 **한계 가성비**를 계산했다.

```text
한계 가성비 = 점수 증가량(ΔScore) / 비용 증가량(ΔCost)
```

- CursorBench·DeepSWE: 추가 비용 1달러당 점수 또는 Pass@1 상승폭
- GDPval-AA v2: 추가 비용 1달러당 Elo 상승폭
- Intelligence Index: 추가 비용 1달러당 지수 상승폭

한계 가성비가 최대인 단계는 수학적 효율이 가장 높은 지점이다. 그러나 실제 운영에서는 절대 성능, 실패 비용, 검증 가능성, 재시도 전략을 함께 고려하여 **실사용 권장 수준**을 별도로 정했다.

---

# 3. CursorBench 3.2

코드 편집·저장소 작업 중심의 코딩 벤치마크로 해석한다.

## 3.1 전체 점수와 비용

| 모델  | 수준   |  점수 | 평균 비용/작업 |
| ----- | ------ | ----: | -------------: |
| Sol   | Low    | 52.6% |          $0.95 |
| Sol   | Medium | 60.0% |          $1.83 |
| Sol   | High   | 63.5% |          $2.62 |
| Sol   | XHigh  | 64.5% |          $3.67 |
| Sol   | Max    | 67.2% |          $5.22 |
| Terra | Low    | 46.9% |          $0.50 |
| Terra | Medium | 50.3% |          $0.58 |
| Terra | High   | 54.2% |          $0.84 |
| Terra | XHigh  | 59.2% |          $1.36 |
| Terra | Max    | 64.9% |          $2.73 |
| Luna  | Low    | 37.6% |          $0.15 |
| Luna  | Medium | 47.7% |          $0.37 |
| Luna  | High   | 56.8% |          $0.77 |
| Luna  | XHigh  | 57.7% |          $1.07 |
| Luna  | Max    | 61.1% |          $1.85 |

## 3.2 한계 가성비

| 모델  | Low→Medium | Medium→High | High→XHigh | XHigh→Max |
| ----- | ---------: | ----------: | ---------: | --------: |
| Sol   |   **8.41** |        4.43 |       0.95 |      1.74 |
| Terra |  **42.50** |       15.00 |       9.62 |      4.16 |
| Luna  |  **45.91** |       22.75 |       3.00 |      4.36 |

단위: 점수 %p / 추가 $1

## 3.3 해석

- **Sol:** Medium에서 가장 효율적이며, High 이후 한계효율이 급격히 감소한다.
- **Terra:** 수학적 최고는 Medium이지만 High~XHigh까지도 상대적으로 완만하게 효율이 감소한다.
- **Luna:** Medium이 최고 효율이지만 High까지 효율이 매우 높고, XHigh에서 급락한다.

## 3.4 CursorBench 추천

| 모델  | 수학적 최고 가성비 | 실사용 추천         |
| ----- | ------------------ | ------------------- |
| Sol   | Medium             | **Medium**          |
| Terra | Medium             | **High 또는 XHigh** |
| Luna  | Medium             | **High**            |

---

# 4. DeepSWE v1.1

장기 소프트웨어 엔지니어링 과제를 수행하는 자율 코딩 에이전트 성능으로 해석한다.

## 4.1 전체 점수와 비용

| 모델  | 수준   | Pass@1 | 평균 비용/작업 |
| ----- | ------ | -----: | -------------: |
| Sol   | Low    |    45% |          $1.07 |
| Sol   | Medium |    61% |          $1.86 |
| Sol   | High   |    69% |          $3.47 |
| Sol   | XHigh  |    71% |          $4.70 |
| Sol   | Max    |    73% |          $8.39 |
| Terra | Low    |    24% |          $0.43 |
| Terra | Medium |    35% |          $0.58 |
| Terra | High   |    54% |          $1.13 |
| Terra | XHigh  |    60% |          $2.13 |
| Terra | Max    |    70% |          $4.95 |
| Luna  | Low    |     2% |          $0.07 |
| Luna  | Medium |    11% |          $0.22 |
| Luna  | High   |    44% |          $0.78 |
| Luna  | XHigh  |    57% |          $1.54 |
| Luna  | Max    |    67% |          $3.03 |

## 4.2 한계 가성비

| 모델  | Low→Medium | Medium→High | High→XHigh | XHigh→Max |
| ----- | ---------: | ----------: | ---------: | --------: |
| Sol   |  **20.25** |        4.97 |       1.63 |      0.54 |
| Terra |  **73.33** |       34.55 |       6.00 |      3.55 |
| Luna  |  **60.00** |       58.93 |      17.11 |      6.71 |

단위: Pass@1 %p / 추가 $1

## 4.3 해석

### Sol

- Low→Medium이 가장 효율적이다.
- Medium→High에서 성공률이 8%p 상승하지만 비용은 $1.61 증가한다.
- High→XHigh, XHigh→Max는 각각 2%p만 상승한다.
- **일반 코딩은 Medium, 어려운 코딩은 High**가 가장 현실적이다.

### Terra

- Medium이 수학적 최고 효율이다.
- Medium→High도 34.55%p/$로 충분히 효율적이다.
- High 이후 한계효율이 크게 떨어진다.
- Terra를 사용해야 한다면 **High**가 적절하다.

### Luna

- Medium과 High의 한계효율 차이가 거의 없다.
- Medium 11%에서 High 44%로 상승하므로 실사용에서는 High까지 올리는 편이 합리적이다.
- XHigh도 중간 품질 구간에서 활용 가능하지만 자율 코딩 Daily driver로는 성공률이 낮다.

## 4.4 DeepSWE 추천

| 모델  | 수학적 최고 가성비 | 실사용 추천                        |
| ----- | ------------------ | ---------------------------------- |
| Sol   | Medium             | **Medium**, 어려운 작업은 **High** |
| Terra | Medium             | **High**                           |
| Luna  | Medium             | **High**                           |

## 4.5 목표 성공률별 선택

| 목표 Pass@1 | 권장 설정  |  비용 |
| ----------- | ---------- | ----: |
| 약 40%      | Luna High  | $0.78 |
| 약 50%      | Terra High | $1.13 |
| 약 55%      | Luna XHigh | $1.54 |
| 약 60%      | Sol Medium | $1.86 |
| 약 65%      | Luna Max   | $3.03 |
| 약 69%      | Sol High   | $3.47 |
| 약 71%      | Sol XHigh  | $4.70 |
| 최고 점수   | Sol Max    | $8.39 |

---

# 5. GDPval-AA v2

문서·슬라이드·스프레드시트·다이어그램 등 실제 지식작업 결과물을 평가한다. 인간 전문가 결과물은 Elo 1,000에 고정되어 있다.

## 5.1 전체 점수와 비용

| 모델  | 수준          |  Elo | 평균 비용/작업 |
| ----- | ------------- | ---: | -------------: |
| Sol   | Non-reasoning | 1384 |          $0.32 |
| Sol   | Low           | 1445 |          $0.33 |
| Sol   | Medium        | 1562 |          $0.74 |
| Sol   | High          | 1630 |          $1.37 |
| Sol   | XHigh         | 1702 |          $2.11 |
| Sol   | Max           | 1748 |          $3.49 |
| Terra | Non-reasoning | 1239 |          $0.16 |
| Terra | Low           | 1249 |          $0.11 |
| Terra | Medium        | 1404 |          $0.23 |
| Terra | High          | 1513 |          $0.61 |
| Terra | XHigh         | 1572 |          $0.92 |
| Terra | Max           | 1593 |          $1.44 |
| Luna  | Non-reasoning | 1068 |          $0.04 |
| Luna  | Low           | 1146 |          $0.04 |
| Luna  | Medium        | 1273 |          $0.07 |
| Luna  | High          | 1472 |          $0.21 |
| Luna  | XHigh         | 1539 |          $0.33 |
| Luna  | Max           | 1592 |          $0.51 |

## 5.2 한계 가성비

| 모델  |  Low→Medium | Medium→High | High→XHigh | XHigh→Max |
| ----- | ----------: | ----------: | ---------: | --------: |
| Sol   |  **285.37** |      107.94 |      97.30 |     33.33 |
| Terra | **1291.67** |      286.84 |     190.32 |     40.38 |
| Luna  | **4233.33** |     1421.43 |     558.33 |    294.44 |

단위: Elo / 추가 $1

## 5.3 모델별 해석

### Sol

- Medium이 첫 번째 가성비 변곡점이다.
- High와 XHigh의 한계효율은 비슷하며, 높은 품질이 필요하면 XHigh까지 합리적이다.
- Max는 XHigh보다 $1.38 비싸지만 Elo는 46만 증가한다.

### Terra

- Medium이 수학적 최고지만 High까지는 충분히 효율적이다.
- 모델 내부 추천은 High다.
- 전체 패밀리 비교에서는 거의 모든 수준이 Luna 또는 Sol에 의해 비용·점수 양쪽에서 지배된다.

### Luna

- Medium이 수학적 최고 효율이다.
- Medium→High도 1421.43 Elo/$로 매우 높다.
- High 이후 효율은 떨어지지만 절대 비용이 낮다.
- **Luna Max는 Sol Medium보다 저렴하면서 Elo가 높다.**

## 5.4 전체 Pareto frontier

| 설정          |      Elo |      비용 |
| ------------- | -------: | --------: |
| Luna Low      |     1146 |     $0.04 |
| Luna Medium   |     1273 |     $0.07 |
| Luna High     |     1472 |     $0.21 |
| Luna XHigh    |     1539 |     $0.33 |
| **Luna Max**  | **1592** | **$0.51** |
| **Sol High**  | **1630** | **$1.37** |
| **Sol XHigh** | **1702** | **$2.11** |
| Sol Max       |     1748 |     $3.49 |

## 5.5 GDPval 추천

| 운영 목적            | 권장 설정     |
| -------------------- | ------------- |
| 저비용 대량 지식작업 | Luna High     |
| 일반 지식작업 기본값 | **Luna Max**  |
| 1600 Elo 이상        | **Sol High**  |
| 1700 Elo 이상        | **Sol XHigh** |
| 절대 최고 품질       | Sol Max       |

---

# 6. Artificial Analysis Intelligence Index

지식·추론·수학·코딩·에이전트 작업을 합성한 범용 종합 지수다. 아래 비용은 제공된 그래프를 기준으로 읽은 근사값이며, 센트 단위 반올림 오차가 있을 수 있다.

## 6.1 전체 점수와 비용

| 모델  | 수준   | Intelligence Index | 평균 비용/작업 |
| ----- | ------ | -----------------: | -------------: |
| Sol   | Low    |                 49 |       약 $0.20 |
| Sol   | Medium |                 54 |       약 $0.31 |
| Sol   | High   |                 56 |       약 $0.45 |
| Sol   | XHigh  |                 58 |       약 $0.68 |
| Sol   | Max    |                 59 |       약 $1.04 |
| Terra | Low    |                 41 |       약 $0.10 |
| Terra | Medium |                 46 |       약 $0.13 |
| Terra | High   |                 49 |       약 $0.24 |
| Terra | XHigh  |                 52 |       약 $0.33 |
| Terra | Max    |                 55 |       약 $0.55 |
| Luna  | Low    |                 33 |       약 $0.04 |
| Luna  | Medium |                 38 |       약 $0.05 |
| Luna  | High   |                 46 |       약 $0.09 |
| Luna  | XHigh  |                 49 |       약 $0.14 |
| Luna  | Max    |                 51 |       약 $0.21 |

## 6.2 한계 가성비

| 모델  | Low→Medium | Medium→High | High→XHigh | XHigh→Max |
| ----- | ---------: | ----------: | ---------: | --------: |
| Sol   |   **45.5** |        14.3 |        8.7 |       2.8 |
| Terra |  **166.7** |        27.3 |       33.3 |      13.6 |
| Luna  |  **500.0** |       200.0 |       60.0 |      28.6 |

단위: Intelligence Index / 추가 $1

## 6.3 해석

- **Luna:** High가 효율 변곡점이며 Max도 절대 비용이 매우 낮다.
- **Terra:** 모델 내부에서는 XHigh까지 고려 가능하지만 전체 패밀리에서는 다른 설정에 의해 지배된다.
- **Sol:** Medium이 명확한 첫 번째 변곡점이며, XHigh가 프리미엄 품질의 실질적 상한이다.

## 6.4 전체 Pareto frontier

| 설정           |   지수 |      비용 |
| -------------- | -----: | --------: |
| Luna Low       |     33 |     $0.04 |
| Luna Medium    |     38 |     $0.05 |
| Luna High      |     46 |     $0.09 |
| Luna XHigh     |     49 |     $0.14 |
| **Luna Max**   | **51** | **$0.21** |
| **Sol Medium** | **54** | **$0.31** |
| **Sol High**   | **56** | **$0.45** |
| **Sol XHigh**  | **58** | **$0.68** |
| Sol Max        |     59 |     $1.04 |

결과적으로 범용 지수에서는 다음 구조가 가장 단순하다.

```text
저비용·중간 성능: Luna
54점 이상 고성능: Sol
Terra: Pareto frontier에서 제외
```

---

# 7. 벤치마크 통합 해석

## 7.1 코딩 작업

### Daily driver: Sol Medium

- CursorBench에서 Medium이 Sol의 가장 효율적인 상승 지점이다.
- DeepSWE에서 61% Pass@1을 기록한다.
- High 대비 비용 절감폭이 크면서 실사용 성공률이 충분하다.

### Hard driver: Sol High

- DeepSWE에서 69% Pass@1이다.
- XHigh는 71%로 2%p만 높지만 비용은 $1.23 증가한다.
- Max는 73%지만 비용이 $8.39로 크게 증가한다.

### 저비용 코딩: Luna High

적합한 작업:

- 작은 범위의 CRUD 수정
- 테스트·타입·주석 생성
- 단순 리팩터링
- 반복적인 코드 변환
- 검증 가능한 대량 작업

DeepSWE Pass@1 44%이므로 장기 자율 코딩의 단일 기본 모델로는 부족하다. 실패 항목을 Sol로 승격하는 구조가 적절하다.

### GitHub PR 리뷰·코드 리뷰 봇: Terra High

벤치마크와 별개로 Peter Steinberger(@steipete)는 GitHub 리뷰 봇 `clawsweeper`를 5.6 Terra High로 전환한 실사용 결과를 공유했다.

- 전체적으로 약 **40% 빠르며**, 품질 저하는 거의 없었다.
- 5.5보다 모든 측면에서 나았고, 비용도 크게 낮았다.
- XHigh도 시험했지만 속도 이점이 사라졌고, 리뷰 평가에서 눈에 띄는 개선은 없었다.

따라서 지연시간·처리량·비용이 중요한 GitHub PR 리뷰 자동화에서는 **Terra High를 우선 기본값으로 검증**하고, 리뷰 품질 개선이 확인되지 않는 한 XHigh로 자동 승격하지 않는 편이 합리적이다. 이 사례를 모든 코딩 작업에 일반화하지는 않는다.

## 7.2 지식작업

### Daily driver: Luna Max

GDPval-AA v2:

| 설정       |      Elo |      비용 |
| ---------- | -------: | --------: |
| Luna Max   | **1592** | **$0.51** |
| Sol Medium |     1562 |     $0.74 |

지식작업에서는 Luna Max가 Sol Medium보다 저렴하고 점수도 높다.

적합한 작업:

- 리서치·자료 종합
- 문서·보고서 작성
- 슬라이드 초안
- 스프레드시트 분석
- 정책·계약 비교
- 구조화된 결과물 생성

### Hard driver: Sol XHigh

- GDPval Elo 1702
- 높은 품질을 확보하면서 Max의 급격한 한계효율 저하를 피한다.
- 복잡한 자료 간 모순 해결, 경영진용 최종 문서, 고위험 의사결정 자료에 적합하다.

## 7.3 복합 추론·기획·아키텍처

### Daily driver: Sol Medium

Intelligence Index에서 Sol Medium은 54로 Luna Max의 51보다 높다. 수학·코딩·추론·지식작업이 혼합된 문제에서는 Sol Medium이 안정적인 기본값이다.

### Hard driver: Sol XHigh

- Intelligence Index 58
- Max는 59로 1점만 높지만 비용은 약 53% 증가한다.
- 시스템 아키텍처, 복합 장애 분석, 여러 제약을 포함한 전략 설계에 적합하다.

## 7.4 대량 반복 작업

### Daily driver: Luna High

- 비용이 낮고 High까지 효율이 좋다.
- 작업 결과를 자동 검증할 수 있는 경우 특히 유리하다.

### Hard driver: Luna Max

- 일부 어려운 항목만 Luna Max로 재시도한다.
- 여전히 실패하면 Sol Medium 또는 Sol High로 승격한다.

---

# 8. 권장 라우팅 정책

## 8.1 코딩 라우팅

```text
작은 수정·반복 작업
  → Luna High

GitHub PR 리뷰·코드 리뷰 봇
  → Terra High

일반 개발·리팩터링·디버깅
  → Sol Medium

저장소 전체 변경·복잡한 버그·다단계 테스트
  → Sol High

High 실패 + 실패 비용 큼
  → Sol XHigh

절대 최고 성공률이 필요함
  → Sol Max
```

## 8.2 지식작업 라우팅

```text
대량 요약·분류·초안
  → Luna High

일반 보고서·리서치·슬라이드·스프레드시트
  → Luna Max

복잡한 모순 해결·고품질 최종 산출물
  → Sol XHigh

최고 품질이 비용보다 중요함
  → Sol Max
```

## 8.3 자동 승격 조건 예시

다음 중 하나에 해당하면 상위 driver로 승격한다.

- 테스트 실패 또는 빌드 실패
- 요구사항 누락
- 인용·근거 부족
- 출력 형식 검증 실패
- 도구 호출 실패
- 상호 모순된 결론
- 낮은 자체 확신도
- 사람이 재작업해야 하는 수준의 품질 저하

---

# 9. Terra의 위치

현재 확인한 네 벤치마크만 보면 Terra는 전체 모델군의 기본 driver로 추천하기 어렵다.

- GDPval에서 거의 전 구간이 Luna 또는 Sol에 의해 지배된다.
- Intelligence Index에서도 Terra 전체가 Pareto frontier 밖이다.
- 코딩 벤치마크만으로는 Terra High의 실사용 가치를 충분히 포착하지 못한다.

### 9.1 GitHub 리뷰 작업에 대한 보정

Peter Steinberger의 실사용 보고에 따르면 `clawsweeper`를 5.6 Terra High로 바꾼 뒤 전체 처리 시간이 약 **40% 단축**됐고, 품질 저하는 거의 없었다. 5.5보다 전반적으로 나았으며 비용도 크게 낮았다. 반면 XHigh는 속도 이점을 상쇄했고 리뷰 평가에서 눈에 띄는 차이를 만들지 못했다.

따라서 Terra를 일괄적으로 제외해서는 안 된다.

- **GitHub PR 리뷰·코드 리뷰 봇:** Terra High를 우선 검증
- **일반 코딩·장기 자율 코딩:** 기존 벤치마크 기반의 Sol Medium/High 유지
- **리뷰 작업의 Terra XHigh:** 품질 개선이 실제로 확인될 때만 사용

즉, 전체 모델 라우팅은 여전히 Luna + Sol 중심이 단순하지만, **코드 리뷰 자동화라는 좁고 반복적인 작업에서는 Terra High가 비용·속도·품질의 실용적인 Pareto 선택지**가 될 수 있다.

---

# 10. 최종 권장안

## 10.1 작업 유형별 최종 선택

| 작업 유형           | Daily driver   | Hard driver    |
| ------------------- | -------------- | -------------- |
| 일반 코딩           | **Sol Medium** | **Sol High**   |
| GitHub PR 리뷰·코드 리뷰 봇 | **Terra High** | **Terra High** |
| 장기 자율 코딩      | **Sol Medium** | **Sol High**   |
| 간단·대량 코딩      | **Luna High**  | **Sol Medium** |
| 일반 지식작업       | **Luna Max**   | **Sol XHigh**  |
| 복합 추론·아키텍처  | **Sol Medium** | **Sol XHigh**  |
| 요약·분류·변환      | **Luna High**  | **Luna Max**   |
| 최고 품질 최종물    | **Sol High**   | **Sol XHigh**  |
| 마지막 에스컬레이션 | —              | **Sol Max**    |

## 10.2 운영 원칙

1. **코딩 기본값은 Sol Medium**
2. **어려운 코딩은 Sol High**
3. **GitHub PR 리뷰·코드 리뷰 봇은 Terra High를 우선 검증**
4. **지식작업 기본값은 Luna Max**
5. **복합 추론과 어려운 지식작업은 Sol XHigh**
6. **저비용 대량 처리는 Luna High**
7. **Sol Max는 최후의 에스컬레이션 전용**
8. **Terra XHigh는 리뷰 품질 개선이 확인될 때만 사용**

---

# 11. 주의사항

- 벤치마크별 작업 유형과 평가 방식이 다르므로 점수를 직접 합산해서는 안 된다.
- GDPval과 Intelligence Index는 일부 평가 요소가 중복될 수 있다.
- 오차막대가 겹치는 작은 점수 차이는 확정적인 우열로 해석하기 어렵다.
- Intelligence Index 비용은 제공된 산점도에서 읽은 근사값이므로 소수점 단위 오차가 있을 수 있다.
- 실제 비용은 프롬프트 길이, 출력 길이, 캐시 적중률, 도구 호출, 재시도 횟수에 따라 달라진다.
- Terra High의 약 40% 속도 향상과 품질 저하가 거의 없다는 평가는 단일 GitHub 리뷰 봇의 실사용 사례이므로 모든 작업에 일반화해서는 안 된다.
- 최종 모델 라우팅은 사내 작업 50~200개로 별도 검증하는 것이 바람직하다.

---

# 12. 출처

- Cursor Evals: <https://cursor.com/ko/evals>
- DeepSWE: <https://deepswe.datacurve.ai/>
- Artificial Analysis GDPval-AA: <https://artificialanalysis.ai/evaluations/gdpval-aa>
- Artificial Analysis 모델·Intelligence Index: <https://artificialanalysis.ai/>
- Peter Steinberger의 5.6 Terra High 실사용 의견: <https://x.com/steipete/status/2078236791329657017>
