# Date Object in JavaScript

- Date 생성자는 특정 시점을 나타내는 `Date` 객체를 생성한다
- JavaScript `Date`객체는 기본적으로 OS에 설정된 현재지역 타임존(Local timezone)을 따른다.
- Date 객체는 내부 데이터로 1970년 1월 1일 (UTC) 00:00으로부터 지난 시간을 miliseconds 값을 가지고 있다
- Date 객체는 유닉스 시간으로부터 약 ±100,000,000일 (1억일)을 기록할 수 있으므로, 293,742년이 오면 문제가 발생할 수 있다

> Unix 시간은 UTC 시작부터 경과시간을 **seconds**로 환산하여 나타낸 것이다. JavaScript Date 객체는 **miliseconds** 단위이다.

## 생성

```js
// 현재시각 Date 객체
new Date();
// unix timestamp in miliseconds
new Date(value);
// IETF 호환 RFC 2822 타임스탬프 또는 ISO-8601 형식으로 작성되어야 함
new Date(dateString);
// 개별 날짜 및 시각
new Date(year, monthIndex[, day[, hour[, minutes[, seconds[, milliseconds]]]]]);

// UTC 기준으로 1970년 1월 1일 0시 0분 0초부터 현재까지 경과된 밀리 초를 반환
Date.now();
// 쉼표로 구분된 날짜 매개변수를 받아 사용자가 지정한 시간과 1970년 1월 1일 00:00:00 사이의 밀리초 수를 반환 (local time 대신 UTC 사용)
Date.UTC(year, month[, day[, hour[, minute[, second[, millisecond]]]]])

// Unix 시간 얻기
Math.floor(Date.now() / 1000);
```

- 입력 및 수정에 사용되는 함수(`생성자`, `parse()`, `getHour()`, `setHour()`...)들은 모두 실행환경의 local time 기준으로 작동한다.
- 연도 입력시 0~99 범위는 1900~1999로 처리됨
- 월 입력시 0부터 시작 (1월 == 0, 2월 == 1)
- `Date.parse()`는 구현방식이 상이하여 사용하지 않기를 권장함 (브라우저 기준, ES6 이전의 경우)
- `dateString`의 경우 구현에 따라 마지막 `Z`가 붙어있지 않으면 local time으로 처리되므로 유의해야 함: `2019-07-28T21:44:20`

## 변환(formatting)

다음 문자열 변환은 UTC+0(`Z`) 기준으로 출력된다.

- ISO-8601 포맷: `toISOString()`, `toJSON()`
- RFC-1123 포맷: `toGMTString()`, `toUTCString()`

## offset 적용

`getTimeZoneOffset()` 메소드를 사용하면 현재 timezone offset을 분 단위로 출력한다

> i.e. +09:00(Asia/Seoul) => -540 출력 (offset과 부호가 반대이므로 유의)

그러나 offset만 가지고는 time zone 적용이 불가능하다 (time zone은 해당 지역과 offset 데이터의 집합임)

## Date Library

편리한 날짜 계산과 time zone 적용을 위해서 사용할 수 있는 라이브러리 목록이다

- moment: <https://github.com/moment/moment>
- moment-timezone: <https://github.com/moment/moment-timezone/>
- dayjs: <https://github.com/iamkun/dayjs>
- date-fns: <https://github.com/date-fns/date-fns>
