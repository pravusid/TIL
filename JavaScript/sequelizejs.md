# Sequelize.js

<http://docs.sequelizejs.com/>

sequelize-typescript: <https://github.com/RobinBuschmann/sequelize-typescript>

## Options

- supportBigNumbers: true (driver default: false)
- bigNumberStrings: false (driver default: false)

`node-mysql2` 드라이버의 경우 집계함수 결과값이 `string`으로 처리되는데,
이를 `number`로 변경하려면 `sequelize`의 경우 다음처럼 적용한다

```js
const sequelize = new Sequelize("database", "username", "password", {
  dialect: "mysql",
  dialectOptions: { decimalNumbers: true }
});
```

## Composite foreign keys

<https://github.com/sequelize/sequelize/issues/311>
