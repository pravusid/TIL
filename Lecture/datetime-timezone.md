# DateTime & TimeZone

- <https://en.wikipedia.org/wiki/Coordinated_Universal_Time>
- <https://en.wikipedia.org/wiki/ISO_8601>

## GMT(Greenwich Mean Time)

- 영국의 그리니치 천문대를 기점으로 하는 시간대이다.
- 1925년 2월 5일 시행

## UTC(Coordinated Universal Time)

- 국제원자시와 윤초 보정을 적용하고 GMT에 기반한 국제 표준시이다
- 1970년 1월 1일을 기점으로 1972년 1월 1일 시행

## ISO-8601

- 날짜와 시간 관련 데이터 교환을 다루는 국제표준이며 1988년 처음 공개되었다
- 일반적으로 그레고리력에서 날짜와 24시간제에 기반한 시간, 시간간격 및 그들의 조합과 표현/형식에 적용된다
- 날짜와 시간을 함께 표현하기 위해 구분문자 `T`를 사용한다, 상호 동의하에 `T`를 생략가능하다: `<date>T<time>`
- UTC 표기에서 offset `0`을 표현하기 위해서는 시간 뒤에 빈칸없이 `Z`나 `+00:00`를 추가한다

> i.e. UTC 표기 예시: `2016-10-27T17:13:40+00:00` || `2016-10-27T17:13:40Z` || `20161027T171340Z`

## Time offsets

- offset `+09:00`은 기준인 UTC보다 9시간 **빠르다**는 뜻이다
- <https://en.wikipedia.org/wiki/List_of_UTC_time_offsets>

## Time zone

- 국가/지역에 따라 사용하는 시간대역(offset)에 이름을 부여하여 타임존이라고 한다. (한국: KST)

- 타임존은 DST(Daylight Saving Time == summer time)에 따라 변화하기도한다.
  - Pacific Standard Time: UTC-08:00
  - Pacific Daylight Time: UTC-07:00

- 또한 타임존은 시점에 따라 변화하기도 한다
  - 조선말기 타임존이 지정되지 않았던 때: +08:28
    - (+08:27:52): ISO-8601 표준은 offset에 분단위까지만 있지만 표준이 없던 시점을 불러오면 초가 젹용되는 경우(경도에 맞춰 적용됨)가 있음
  - 1908년 이후 +08:30
  - 1912년 이후 +09:00
  - 이후로도 몇 차례 변화가 있음 (+08:30, +09:00 및 일광절약시간제)

## IANA time zone database

- <https://www.iana.org/time-zones>
- 타임존은 규칙이라기 보다는 해당 지역의 offset 데이터 집합이라고 볼 수 있다
- IANA tz database는 이러한 타임존 표준시와 변경내역을 가장 충실하게 반영한 곳이다
- `Area / Location` 형태로 구분된다 (Asia / Seoul)
- 유닉스 기반 OS, 대중적 프로그래밍 언어의 Date 시스템에서 사용하고 있다
