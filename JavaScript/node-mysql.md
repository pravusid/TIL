# Node Mysql

두 가지 선택가능한 패키지가 있다

<https://github.com/mysqljs/mysql>

<https://github.com/sidorares/mysql2>

## mysqljs vs mysql2

<https://github.com/sidorares/mysql2/tree/master/documentation#known-incompatibilities-with-node-mysql>

`mysql2`에서는 `mysqljs`과 두 가지 큰 차이점이 있다

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

`typeorm`의 경우 다음 두 옵션과 관계있지만 기본적으로 `mysqljs`에서도 `string`을 반환한다

- `supportBigNumbers`: When dealing with big numbers (BIGINT and DECIMAL columns) in the database, you should enable this option **(Default: true)**

- `bigNumberStrings`: Enabling both supportBigNumbers and bigNumberStrings forces big numbers (BIGINT and DECIMAL columns) to be always returned as JavaScript String objects **(Default: true)**. Enabling supportBigNumbers but leaving bigNumberStrings disabled will return big numbers as String objects only when they cannot be accurately represented with JavaScript Number objects (which happens when they exceed the [-2^53, +2^53] range), otherwise they will be returned as Number objects. This option is ignored if supportBigNumbers is disabled.

> timezone connection option is not supported by mysql2

- mysqljs: set timezone option to +0
- mysql2: run the node process with the timezone environment variable `TZ=UTC`
