# Sequelize.js

<http://docs.sequelizejs.com/>

## Options

- supportBigNumbers: true (driver default: false)
- bigNumberStrings: false (driver default: false)

`node-mysql2` 드라이버의 경우 집계함수 결과값이 `string`으로 처리되는데,
이를 `number`로 변경하려면 `sequelize`의 경우 다음처럼 적용한다

```js
const sequelize = new Sequelize("database", "username", "password", {
  dialect: "mysql",
  dialectOptions: { decimalNumbers: true },
});
```

## Composite foreign keys

<https://github.com/sequelize/sequelize/issues/311>

## sequelize-typescript

sequelize-typescript: <https://github.com/RobinBuschmann/sequelize-typescript>

### Model definition

```ts
import { Table, Column, Model, HasMany } from "sequelize-typescript";

@Table
class Person extends Model<Person> {
  @Column
  name: string;

  @Column
  birthday: Date;

  @HasMany(() => Hobby)
  hobbies: Hobby[];
}
```

### Model Attributes 추출

```ts
import { Model } from "sequelize-typescript";

export type Props<T> = Omit<T, keyof Model<T>>;
```

`sequelize` 모듈의 `Model` 추상클래스는 다음 프로퍼티를 포함한다.

- `createdAt?: Date | any`
- `updatedAt?: Date | any`
- `deletedAt?: Date | any`

Entity에서 해당 프로퍼티를 override 한다면 `Props<T>` 타입으로 정확한 타입을 추출할 수 없다.

## Troubleshooting

### connection deadlock

- <https://github.com/sequelize/sequelize/issues/11024>
- <https://github.com/sequelize/sequelize/issues/11571>

```ts
import PQueue from "p-queue";
import { Op, Transaction } from "sequelize";

const transactionQueue = new PQueue({ concurrency: connectionPoolSize - 1 });

export const dbTransaction = <T>(
  fn: (transaction: Transaction) => Promise<T>
) =>
  transactionQueue.add(() =>
    sequelize.transaction(async (transaction) => {
      try {
        const result = await fn(transaction);
        return result;
      } catch (e) {
        if (e.parent?.code === "ER_LOCK_DEADLOCK") {
          await (transaction as any).cleanup();
        }
        throw e;
      }
    })
  );
```
