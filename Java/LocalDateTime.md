# JSR-310

## LocalDate

날짜 정보만 필요할 때 사용한다

### LocalDate 필드

```java
LocalDate.MIN // -999999999-01-01 - 과거 표현시
LocalDate.MAX // +999999999-12-31 - 미래 표현시
```

### LocalDate 생성

```java
LocalDate.now([ZoneId zone]); // Timezone의 현재날짜, 지정되어 있지 않으면 현재 Timezone의 날짜를 반환한다
LocalDate.of(int year, int month, int dayOfMonth); // 인자로 받은 날짜를 반환한다
LocalDate.ofEpochDay(long epochDay) // UNIX 타임을 받은 경우
LocalDate.parse(CharSequence text, [DatetimeFormatter formatter]) // 날짜 문자열을 파싱함
```

### LocalDate 메소드: 조회

- `int getYear()`: 연도를 얻는다
- `int getMonthValue()`: 1 to 12
- `Month getMonth()`: Month enum 값을 얻는다
- `int getDayOfMonth()`: 1 to 31
- `int getDayOfYear()`: 1 to 365 (366 윤년)
- `DayOfWeek getDayOfWeek()`: DayOfWeek enum 값을 얻는다
- `boolean isLeafYear()`: 윤년 여부
- `int lengthOfMonth()`: 28 to 31
- `int lengthOfYear()`: 365 / 366
- `LocalDate --With--(*)`: 해당 값으로 설정한 LocalDate를 반환하여 Chaining 할 수 있다
- `String format(DateTimeFormatter formatter)`: 포맷을 지정하여 문자열로 반환한다
- `LocalDateTime atTime(LocalTime time)`: `LocalTime`을 인자로 받아 합성한 후 `LocalDateTime`을 반환한다
- `LocalDateTime atTime(int hour, int minute, [int second], [int nanoOfSecond])`: 직접 입력받은 시각으로 합성한 후 `LocalDateTime`을 반환한다
- `LocalDateTime atStartOfDay([ZoneId zone])`: 해당날짜의 `LocalDateTime` 00:00분을 반환한다

### LocalDate 메소드: 연산

- `LocalDate plus/minus(long amount, TemporalUnit unit)`: 날짜를 더한다/뺀다, `ChronoUnit`의 enum 값으로 단위를 설정한다
- `LocalDate plus/minusYears(long years)`: 연도를 더한다/뺀다
- `LocalDate plus/minusMonths(long months)`: 달을 더한다/뺀다 (마지막 날짜를 넘을때는 3/31 -> 4/31 마지막 날짜가 반환된다 4/30)
- `LocalDate plus/minusWeeks(long weeks)`: 주를 더한다/뺀다
- `LocalDate plus/minusDays(long days)`: 날짜를 더한다/뺀다
- `long until(Temporal endExclusive, TemporalUnit unit)`: 주어진 날짜와 인자의 날짜간의 차이를 구한다
- `Period until(ChronoLocalDate endDateExclusive)`: 주어진 날짜와 인자의 날짜간의 차이를 구한다

### LocalDate 메소드: 비교

- `int compareTo(ChronoLocalDate other)`: 날짜를 비교하여 결과를 반환한다 (음수, 0, 양수)
- `boolean isAfter(ChronoLocalDate other)`: 인자보다 이후 날짜인지 비교
- `boolean isBefore(ChronoLocalDate other)`: 인자보다 이전 날자인지 비교
- `boolean isEqual(ChronoLocalDate other)`: 인자와 같은 날짜인지 비교

날짜 비교는 `Period` 클래스를 통해서도 할 수 있다

```java
Period period = currentDate.until(targetDate);
period.getDays();
```

## LocalTime

시간 정보만 필요할 때 사용한다

### LocalTime 필드

```java
LocalTime.MAX // 23:59:59.999999999
LocalTime.MIN // 00:00
LocalTime.MIDNIGHT // 00:00, the start of the day
LocalTime.NOON // 12:00, the middle of the day
```

### LocalTime 생성

```java
LocalTime.now([ZoneId zone]); // Timezone의 현재시각, 지정되어 있지 않으면 현재 Timezone의 시각을 반환한다
LocalTime.of(int, hour, int minute, [int second, int nanoOfSecond]) // 인자로 받은 시각을 반환한다
LocalTime.ofSecondOfDay(long secondOfDay) // the second-of-day, from 0 to 24 * 60 * 60 - 1
LocalTime.ofNanoOfDay(long nanoOfDay) // the nano of day, from 0 to 24 * 60 * 60 * 1,000,000,000 - 1
LocalTime.parse(CharSequence text, [DatetimeFormatter formatter]) // 시각 문자열을 파싱함
```

### LocalTime 메소드: 조회

- `int getHour()`: 0 to 23
- `int getMinute()`: 0 to 59
- `int getSecond()`: 0 to 59
- `int getNano()`: 0 to 999,999,999
- `LocalTime --With--(*)`: 해당 값으로 설정한 LocalTime을 반환하여 Chaining 할 수 있다
- `int toSecondOfDay()`: 현재 시각이 하루 시작으로 부터 몇 초가 지났는지 반환
- `long toNanoOfDay()`: 현재 시각이 하루 시작으로 부터 몇 나노초가 지났는지 반환

### LocalTime 메소드: 연산

- `LocalTime plus/minus(long amount, TemporalUnit unit)`: 시간을 더한다/뺀다, `ChronoUnit`의 enum 값으로 단위를 설정한다
- `LocalTime plus/minusHours(long hours)`: 시간을 더한다/뺀다
- `LocalTime plus/minusMinutes(long minutes)`: 분을 더한다/뺀다
- `LocalTime plus/minusSeconds(long seconds)`: 초를 더한다/뺀다
- `LocalTime plus/minusNanos(long nanos)`: 나노초를 더한다/뺀다
- `long until(Temporal endExclusive, TemporalUnit unit)`: 주어진 시각과 인자의 시각간의 차이를 구한다
- `LocalDateTime atDate(LocalDate date)`: `LocalDate`를 입력받아 합성한 후 `LocalDateTime`을 반환한다

### LocalTime 메소드: 비교

- `int compareTo(LocalTime other)`: 시각을 비교하여 결과를 반환한다 (음수, 0, 양수)
- `boolean isAfter(LocalTime other)`: 인자보다 이후 시각 여부
- `boolean isBefore(LocalTime other)`: 인자보다 이전 시각 여부

시간 비교는 `Duration` 클래스를 통해서도 할 수 있다

```java
Duration duration = Duration.between(startTime, endTime);
duration.getSeconds();
```

## LocalDateTime

날짜, 시간 모두 필요할 때 사용한다

### LocalDateTime 필드

```java
LocalDateTime.MIN // -999999999-01-01T00:00:00
LocalDateTime.MAX // +999999999-12-31T23:59:59.99999999
```

### LocalDateTime 생성

```java
LocalDateTime.now([ZoneId zone]); // Timezone의 현재날짜시각, 지정되어 있지 않으면 현재 Timezone의 날짜시각을 반환한다
LocalDateTime.of(int year, int month, int dayOfMonth, int, hour, int minute, [int second, int nanoOfSecond]) // 인자로 받은 날짜시각을 반환한다
LocalDateTime.of(LocalDate date, LocalTime time) // 인자로 받은 날짜시각을 반환한다
LocalDateTime.ofInstant(Instant instant, ZoneId zone) // 인자로 받은 날짜시각을 반환한다
LocalDateTime.ofEpochSecond(long epochSecond, int nanoOfSecond, ZoneOffset offset) // UNIX 시각으로 현재 날짜시각을 구한다
LocalDateTime.parse(CharSequence text, [DatetimeFormatter formatter]) // 시각 문자열을 파싱함
```

### LocalDateTime 메소드

조회 및 연산 메소드는 `LocalDate`와 `LocalTime`에 해당하는 기능을 대부분 사용할 수 있다.

날짜와 시각이 동시에 있는 경우 비교는 `ChronoUnit`을 통해 할 수 있다

```java
ChronoUnit.YEARS.between(start, end);
ChronoUnit.MONTHS.between(start, end);
ChronoUnit.DAYS.between(start, end);
ChronoUnit.HOURS.between(start, end);
ChronoUnit.SECONDS.between(start, end);
```

## TemporalAdjusters

`--with--` 메소드와 `TemporalAdjuster` 메소드를 사용하면 날짜와 시각을 상대적으로 변경할 수 있다.

```java
LocalDateTime targetDateTime = currentDateTime
    .with(TemporalAdjusters.firstDayOfYear()) // 이번 년도의 첫 번째 일(1월 1일)
    .with(TemporalAdjusters.lastDayOfYear()) // 이번 년도의 마지막 일(12월 31일)
    .with(TemporalAdjusters.firstDayOfNextYear()) // 다음 년도의 첫 번째 일(1월 1일)
    .with(TemporalAdjusters.firstDayOfMonth()) // 이번 달의 첫 번째 일(1일)
    .with(TemporalAdjusters.lastDayOfMonth()) // 이번 달의 마지막 일
    .with(TemporalAdjusters.firstDayOfNextMonth()) // 다음 달의 첫 번째 일(1일)
    .with(TemporalAdjusters.firstInMonth(DayOfWeek.MONDAY)) // 이번 달의 첫 번째 요일(여기서는 월요일)
    .with(TemporalAdjusters.lastInMonth(DayOfWeek.FRIDAY)) // 이번 달의 마지막 요일(여기서는 마지막 금요일)
    .with(TemporalAdjusters.next(DayOfWeek.FRIDAY)) // 다음주 금요일
    .with(TemporalAdjusters.nextOrSame(DayOfWeek.FRIDAY)) // 다음주 금요일(오늘 포함. 오늘이 금요일이라면 오늘 날짜가 표시 된다.)
    .with(TemporalAdjusters.previous(DayOfWeek.FRIDAY)) // 지난주 금요일
    .with(TemporalAdjusters.previousOrSame(DayOfWeek.FRIDAY));// 지난주 금요일(오늘 포함)
```

## DateTimeFormatter

DateTime을 생성하거나 파싱할 때 사용

### format()

```java
dt.format(DateTimeFormatter.ISO_DATE_TIME);              // 2017-06-12T14:28:59. 147
date.format(DateTimeFormatter.ISO_LOCAL_DATE);           // 2017-06-12
time.format(DateTimeFormatter.ISO_LOCAL_TIME);           // 14:28:59. 147
dt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME);        // 2017-06-12T14:28:59. 147
odt.format(DateTimeFormatter.ISO_OFFSET_DATE);           // 2017-06-12+09:00
odt.format(DateTimeFormatter.ISO_OFFSET_TIME);           // 14:28:59. 147+09:00
odt.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME);      // 2017-06-12T14:28:59. 147+09:00
zdt.format(DateTimeFormatter.ISO_ZONED_DATE_TIME);       // 2017-06-12T14:28:59. 147+09:00[Asia/Seoul]
zdt.format(DateTimeFormatter.ISO_INSTANT);               // 2017-06-12T05:28:59. 147Z
date.format(DateTimeFormatter.BASIC_ISO_DATE);           // 20170612
date.format(DateTimeFormatter.ISO_DATE);                 // 2017-06-12
time.format(DateTimeFormatter.ISO_TIME);                 // 14:28:59. 147
date.format(DateTimeFormatter.ISO_ORDINAL_DATE);         // 2017-163
date.format(DateTimeFormatter.ISO_WEEK_DATE);            // 2017-W24-1
odt.format(DateTimeFormatter.RFC_1123_DATE_TIME);        // Mon. 12 Jun 2017 14:28:59 +09:00
```

### ofPattern()

```java
DateTimeFormatter dtf = DateTimeFormatter.ofPattern("YYYY-MM-dd");
```

| Symbol | Meaning                    | Presentation     | Examples                                         |
| ------ | -------------------------- | ---------------- | ------------------------------------------------ |
| `G`    | era                        | text             | AD; Anno Domini; A                               |
| `u`    | year                       | year             | 2004; 04                                         |
| `y`    | year-of-era                | year             | 2004; 04                                         |
| `D`    | day-of-year                | number           | 189                                              |
| `M/L`  | month-of-year              | number/text      | 7; 07; Jul; July; J                              |
| `d`    | day-of-month               | number           | 10                                               |
|                                                                                                           |
| `Q/q`  | quarter-of-year            | number/text      | 3; 03; Q3; 3rd quarter                           |
| `Y`    | week-based-year            | year             | 1996; 96                                         |
| `w`    | week-of-week-based-year    | number           | 27                                               |
| `W`    | week-of-month              | number           | 4                                                |
| `E`    | day-of-week                | text             | Tue; Tuesday; T                                  |
| `e/c`  | localized day-of-week      | number/text      | 2; 02; Tue; Tuesday; T                           |
| `F`    | week-of-month              | number           | 3                                                |
|                                                                                                           |
| `a`    | am-pm-of-day               | text             | PM                                               |
| `h`    | clock-hour-of-am-pm (1-12) | number           | 12                                               |
| `K`    | hour-of-am-pm (0-11)       | number           | 0                                                |
| `k`    | clock-hour-of-am-pm (1-24) | number           | 0                                                |
|                                                                                                           |
| `H`    | hour-of-day (0-23)         | number           | 0                                                |
| `m`    | minute-of-hour             | number           | 30                                               |
| `s`    | second-of-minute           | number           | 55                                               |
| `S`    | fraction-of-second         | fraction         | 978                                              |
| `A`    | milli-of-day               | number           | 1234                                             |
| `n`    | nano-of-second             | number           | 987654321                                        |
| `N`    | nano-of-day                | number           | 1234000000                                       |
|                                                                                                           |
| `V`    | time-zone ID               | zone-id          | America/Los_Angeles; Z; -08:30                   |
| `z`    | time-zone name             | zone-name        | Pacific Standard Time; PST                       |
| `O`    | localized zone-offset      | offset-O         | GMT+8; GMT+08:00; UTC-08:00;                     |
| `X`    | zone-offset 'Z' for zero   | offset-X         | Z; -08; -0830; -08:30; -083015; -08:30:15;       |
| `x`    | zone-offset                | offset-x         | +0000; -08; -0830; -08:30; -083015; -08:30:15;   |
| `Z`    | zone-offset                | offset-Z         | +0000; -0800; -08:00;                            |
|                                                                                                           |
| `p`    | pad next                   | pad modifier     | 1                                                |
|                                                                                                           |
| `'`    | escape for text            | delimiter        |                                                  |
| `''`   | single quote               | literal          | '                                                |
| `[`    | optional section start     |                  |                                                  |
| `]`    | optional section end       |                  |                                                  |
| `#`    | reserved for future use    |                  |                                                  |
| `{`    | reserved for future use    |                  |                                                  |
| `}`    | reserved for future use    |                  |                                                  |

## Date(Time) / LocalDate(Time) 사이의 변환

```java
public class DateHelper {

    public static Date asDate(LocalDate localDate) {
        return Date.from(localDate.atStartOfDay().atZone(ZoneId.systemDefault()).toInstant());
    }

    public static Date asDate(LocalDateTime localDateTime) {
        return Date.from(localDateTime.atZone(ZoneId.systemDefault()).toInstant());
    }

    public static LocalDate asLocalDate(Date date) {
        // Java 8
        return date.toInstant().atZone(ZoneId.systemDefault()).toLocalDate();
        // Java 9 이상
        return LocalDate.ofInstant(date.toInstant(), ZoneId.systemDefault())
    }

    public static LocalDateTime asLocalDateTime(Date date) {
        // Java 8
        return date.toInstant().atZone(ZoneId.systemDefault()).toLocalDateTime();
        // Java 9 이상
        return LocalDateTime.ofInstant(date.toInstant(), ZoneId.systemDefault());
    }
}
```
