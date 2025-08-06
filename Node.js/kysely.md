---
tags: nodejs/db querybuilder
---

# Kysely

<https://github.com/kysely-org/kysely>

> Kysely (pronounce “Key-Seh-Lee”) is a type-safe and autocompletion-friendly typescript SQL query builder. Inspired by knex.

## troubleshooting

### `select - limit 1`

> querybuilder 구현에 따라 첫 행만 반환하는 함수에 `limit 1`이 포함되어 있지 않을 수 있음

- kysely

  - `executeTakeFirst`에서 집계 쿼리가 아니라 여러 행이 반환되는 경우 `limit 1`을 호출해야 함
  - <https://github.com/kysely-org/kysely/blob/0982d61a59d6dccaaf2a026cad82a77b843be7ac/src/query-builder/select-query-builder.ts#L2114>

- typeorm querybuilder

  - `getOne`에서 `limit 1` 호출 필요
  - <https://github.com/typeorm/typeorm/blob/d184d8598c057ce8fa54815e669b567238f3a86e/src/query-builder/SelectQueryBuilder.ts#L1696>

- knex

  - `first` 구현에 `limit 1` 포함되어 있음
  - <https://github.com/knex/knex/blob/3ba9550346a4b0220566c32c94751e4c1fc85771/lib/query/querybuilder.js#L1144>

## examples

### Upsert

- <https://kysely-org.github.io/kysely-apidoc/classes/InsertQueryBuilder.html#onConflict>
- <https://kysely-org.github.io/kysely-apidoc/classes/InsertQueryBuilder.html#onDuplicateKeyUpdate>
- [MySQL 8 row aliases](https://github.com/kysely-org/kysely/issues/664)

```ts
db.insertInto('foo')
  .values(list)
  .onDuplicateKeyUpdate(($) => ({
    ...Array.of('name', 'type')
      .map((k) => ({ [k]: sql`values(${$.ref(k)})` }))
      .reduce((a, c) => ({ ...a, ...c }), {}),
  }))
  .execute();
```

upsert util function

> 참고: <https://orm.drizzle.team/docs/guides/upsert>

```ts
import { ExpressionBuilder, StringReference, sql } from 'kysely';

export const buildDuplicateKeyUpdateCols = <DB, TB extends keyof DB>(
  expr: ExpressionBuilder<DB, TB>,
  cols: StringReference<DB, TB>[]
) => {
  return cols.map((k) => ({ [k]: sql`values(${expr.ref(k)})` })).reduce((a, c) => ({ ...a, ...c }), {});
};
```

### JSON value

- <https://github.com/kysely-org/kysely/issues/209>
- <https://kysely.dev/docs/recipes/extending-kysely#expression>

```ts
import { RawBuilder, sql } from 'kysely';

export function jsonValue<T>(obj: T): RawBuilder<T> {
  return sql`${JSON.stringify(obj)}`;
}
```

### Optimizer Hints

`modifyFront`, `modifyEnd` 사용

```ts
db.modifyFront(sql`/*+ MAX_EXECUTION_TIME(5000) */`);
```

### Type inference failure

> 발생한 오류: `'isSelectQueryBuilder' 속성이 'RawBuilder<unknown>' 형식에 없지만 'SelectQueryBuilderExpression<Record<string, number>>' 형식에서 필수입니다.ts(2345)`

오류가 발생한 코드 (실제 오류가 있는 것은 아니고 language server 변경사항이 있으면 추론 과부하가 걸리는 것으로 보임)

```ts
.values([/* ... */])
.onDuplicateKeyUpdate($ => ({
  foo: sql`${$.ref('foo')} + values(${$.ref('foo')})`,
}))
```

타입 캐스팅으로 타입추론 오류 방지

```ts
.values([/* ... */])
.onDuplicateKeyUpdate($ => ({
  foo: sql<number>`${$.ref('foo')} + values(${$.ref('foo')})`,
}))
```
