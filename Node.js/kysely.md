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

### Optimizer Hints

`modifyFront`, `modifyEnd` 사용

```ts
db.modifyFront(sql`/*+ MAX_EXECUTION_TIME(5000) */`);
```
