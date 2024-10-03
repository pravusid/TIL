# Node.js MySQL Drivers

선택 가능한 두 가지 패키지가 있다

- <https://github.com/mysqljs/mysql>
- <https://github.com/sidorares/node-mysql2>

## timezone

> [[mysql#TIMEZONE]] 참고

`mysqljs`, `node-mysql2` 모두 기본 설정값은 `'local'`

그러나 timezone은 database session timezone을 설정하는 것이 아니라 JavaScript Date 객체의 처리 방식을 설정하는 것이다.

JavaScript Date 객체는 내부적으로 Milliseconds Unix Time을 사용하지만 항상 LocalTime으로 표시된다.
그것을, 주어진 timezone 옵션에 맞춰서 컨버팅 하는 것이며 Database에서 가져올때는 그 반대를 적용한다.

이 과정은 드라이버의 `typeCast` 옵션을 사용해서 추가 구현할 수도 있다.

`typeCast` 옵션은 기본적으로 컬럼 값을 Native JavaScript type으로 변환할지 선택하는 옵션으로 두 드라이버 모두 기본값은 `true`이다.

옵션값으로 함수를 설정하면 사용자 정의 타입변환을 추가 적용할 수 있다.

- `mysqljs`: `mysql/lib/protocol/packets/RowDataPacket.js`
- `node-mysql2`: `node-mysql2/lib/parsers/text_parser.js`

```js
const config = {
  //...
  typeCast: function (field, next) {
    if (field.type === 'DATETIME') {
      return new Date(`${field.string()}Z`); // can be 'Z' for UTC or an offset in the form '+HH:MM' or '-HH:MM'
    }
    return next();
  },
};
```

데이터를 UTC로 관리하려면 다음 설정을 사용하면 된다 (타임존 처리는 Application 위임)

- set node-mysql driver's timezone option to `'+0'` or `'+00:00'` or `'Z'`
- run the DBMS with the timezone environment variable `TZ=UTC`

### timezone 테스트

> 실행시각 2023-03-01 10:35:40+09

driver option: `timezone: 'local'`

- 입력 system_tz: UTC

  - datetime (조회 session_tz)
    - 2023-03-01 10:35:40 (UTC)
    - 2023-03-01 10:35:40 (KST) *

  - timestamp (조회 session_tz)
    - 2023-03-01 10:35:40 (UTC)
    - 2023-03-01 19:35:40 (KST)

- 입력 system_tz: KST

  - datetime (조회 session_tz)
    - 2023-03-01 10:35:40 (UTC)
    - 2023-03-01 10:35:40 (KST) *

  - timestamp (조회 session_tz)
    - 2023-03-01 01:35:40 (UTC) *
    - 2023-03-01 10:35:40 (KST) *

driver option: `timezone: 'z'`

- 입력 system_tz: UTC

  - datetime (조회 session_tz)
    - 2023-03-01 01:35:40 (UTC) *
    - 2023-03-01 01:35:40 (KST)

  - timestamp (조회 session_tz)
    - 2023-03-01 01:35:40 (UTC) *
    - 2023-03-01 10:35:40 (KST) *

- 입력 system_tz: KST

  - datetime (조회 session_tz)
    - 2023-03-01 01:35:40 (UTC) *
    - 2023-03-01 01:35:40 (KST)

  - timestamp (조회 session_tz)
    - 2023-02-28 16:35:40 (UTC)
    - 2023-03-01 01:35:40 (KST)

> **datetime** 타입은 타임존에 맞춰 변환하지 않기 때문에 값을 입력할 때 DBMS의 타임존 설정을 신경 쓸 필요가 없다.
> 그러나 여러 지역에 서비스가 분산되어 있다면, DB에 datetime 값을 입력할 때 app에서 보낸 값을 그대로 사용할 것이므로
> 저장된 값의 타임존이 각자 달라서 같은 시간에 다른 값이 저장되지만 타임존 정보가 없으므로 정보유실이 발생하게 된다.
> 따라서, driver(client) 설정은 항상 UTC로 하는 것이 좋을 것 같다 (이렇게 하면 DB에 저장된 datetime 값은 항상 UTC임)
>
> **timestamp** 타입은 입력할 때 DBMS의 타임존과 driver(client) 타임존을 맞춰야 다른 타임존으로 조회했을때 정상값이 출력된다.
> 실수를 방지하기 위해서 datetime과 마찬가지로 모든 설정을 UTC기준으로만 사용하는 것이 좋을 것 같다.

## 에러처리

- 에러코드: <https://github.com/sidorares/node-mysql2/blob/master/lib/constants/errors.js>
- 에러패킷처리 (저수준): <https://github.com/sidorares/node-mysql2/blob/092dc7f103f622b0c9d013b286cea48ed01883a9/lib/packets/packet.js#L718>
- 에러처리 (고수준): <https://github.com/sidorares/node-mysql2/blob/092dc7f103f622b0c9d013b286cea48ed01883a9/promise.js#L7>

## mysqljs vs node-mysql2

<https://github.com/sidorares/node-mysql2/tree/master/documentation#known-incompatibilities-with-node-mysql>

### zeroFill

`zeroFill` flag is ignored in type conversion.
You need to check corresponding field's zeroFill flag and convert to string manually if this is of importance to you.

### DECIMAL and NEWDECIMAL

mysql의 집계함수 결과값은 DECIMAL/DOUBLE로 출력되는데, 두 드라이버의 처리방식이 다르다.

> DECIMAL and NEWDECIMAL types always returned as string unless you pass this config option

```js
{
  decimalNumbers: true;
}
```

### BigNumbers: BigInt, Decimal

#### `supportBigNumbers`

When dealing with big numbers (BIGINT and DECIMAL columns) in the database, you should enable this option

> default: `false`

#### `bigNumberStrings`

Enabling both supportBigNumbers and bigNumberStrings forces big numbers (BIGINT and DECIMAL columns) to be always returned as JavaScript String objects.
Enabling supportBigNumbers but leaving bigNumberStrings disabled will return big numbers as String objects
only when they cannot be accurately represented with JavaScript Number objects (which happens when they exceed the [-2^53, +2^53] range),
otherwise they will be returned as Number objects. This option is ignored if supportBigNumbers is disabled.

> default: `false`

## issues in node-mysql2 with Jest

### `Encoding not recognized`

<https://github.com/sidorares/node-mysql2/issues/489>

> 다음 인코딩 불러오기 코드를 database connection 생성 코드 상단에 입력한다

JavaScript

```js
require('iconv-lite').encodingExists('cesu8');
```

TypeScript

```ts
import * as iconv from 'iconv-lite';
iconv.encodingExists('cesu8');
```

## 쿼리결과 스트림 처리

<https://github.com/sidorares/node-mysql2/issues/677>
