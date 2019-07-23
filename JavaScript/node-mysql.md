# Node.js MySQL Drivers

두 가지 선택가능한 패키지가 있다

<https://github.com/mysqljs/mysql>

<https://github.com/sidorares/node-mysql2>

## timezone

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
      return new Date(`${field.string()}Z`) // can be 'Z' for UTC or an offset in the form '+HH:MM' or '-HH:MM'
    }
    return next();
  }
}
```

Database 입력 값에 UTC를 적용하려면 다음 두 방식을 사용할 수 있다 (검증 필요)

- set timezone option to `'+0'` or `'Z'`
- run the node process with the timezone environment variable `TZ=UTC`

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

`sequelize`의 경우 다음처럼 적용한다

```js
const sequelize = new Sequelize("database", "username", "password", {
  dialect: "mysql",
  dialectOptions: { decimalNumbers: true }
});
```
