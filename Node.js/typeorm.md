# TypeORM

<http://typeorm.io>

## 설치

```sh
npm install typeorm --save
npm install reflect-metadata --save
npm install @types/node --save
```

드라이버를 설치한다 (택1)

```sh
npm install mysql --save
npm install mysql2 --save
```

앱 실행 시작점에서 reflect-metadata를 불러온다

```ts
import 'reflect-metadata';
```

`tsconfig.json`에 다음을 추가한다

```json
"emitDecoratorMetadata": true,
"experimentalDecorators": true,
```

## 설정

프로젝트 루트에 `ormconfig.json`을 작성하고 db연결정보를 기록한다

json 뿐만 아니라 `js`, `yml`, `.env` 등의 다양한 방식으로 작성 가능하다

```json
[
  {
    // 여러 연결을 구분하기 위해서 사용 (default는 생략가능)
    "name": "default",
    // 사용할 db 종류를 정의
    "type": "mysql",
    "host": "localhost",
    "port": 3306,
    "username": "test",
    "password": "test",
    "database": "test",
    // 앱 실행마다 데이터베이스 스키마를 동기화 할 것인지 설정
    "synchronize": true,
    // 앱 실행마다 스키마를 삭제할 것인지 설정(개발용)
    "dropSchema": false,
    // 로그 출력: true/false 대신 로그 레벨을 쓸 수 있다
    // ["query", "error", "schema", "warn", "info", "log", "all"]
    "logging": true,
    // 기록된 시간 이상이 소요되는 쿼리의 로그를 남긴다
    "maxQueryExecutionTime": 1000,
    // 로그 출력방식을 선택한다
    // simple-console(기본에서 색이 빠짐), file(ormlogs.log 파일로 출력)
    "logger": "advanced-console",
    // @Entity()로 선언된 파일을 불러올 위치
    "entities": ["src/entity/**/*.ts"],
    // DDL을 실행하는 MigrationInterface를 구현한 파일을 불러올 위치
    "migrations": ["src/migration/**/*.ts"],
    // @After___, @Before___ 같은 entity listeners and subscribers를 불러올 위치
    "subscribers": ["src/subscriber/**/*.ts"]
  },
  {
    "name": "another",
    "type": "mysql",
    "host": "localhost",
    "port": 3306,
    "username": "test",
    "password": "test",
    "database": "test"
  }
]
```

> mysql/mariadb 관련 연결옵션은 다음을 참조: <http://typeorm.io/#/connection-options/mysql--mariadb-connection-options>

### nodejs-mysql 드라이버(mysqljs, node-mysql2)와 기본 옵션이 다른 경우

#### `supportBigNumbers`

When dealing with big numbers (BIGINT and DECIMAL columns) in the database, you should enable this option

- mysqljs: `false`
- node-mysql2: `false`
- typeorm: `true`

#### `bigNumberStrings`

Enabling both supportBigNumbers and bigNumberStrings forces big numbers (BIGINT and DECIMAL columns) to be always returned as JavaScript String objects.
Enabling supportBigNumbers but leaving bigNumberStrings disabled will return big numbers as String objects
only when they cannot be accurately represented with JavaScript Number objects (which happens when they exceed the [-2^53, +2^53] range),
otherwise they will be returned as Number objects. This option is ignored if supportBigNumbers is disabled.

- mysqljs: `false`
- node-mysql2: `false`
- typeorm: `true`

## 연결

db 연결을 시도할 때 기본 옵션을 동적으로 override 할 수 있다

```ts
const connectionOptions = await getConnectionOptions();
Object.assign(connectionOptions, {
  namingStrategy: new CustomNamingStrategy(),
  logger: new MyCustomLogger(),
});
const connection = await createConnection(connectionOptions);
```

db 연결 설정 파일을 변경할 수 있다

```ts
import { createConnection, ConnectionOptionsReader } from 'typeorm';
import { CustomNamingStrategy } from './custom.naming.strategy';

export const connectToDatabase = async (env?: string) => {
  const configName = env ? `ormconfig.${env}` : 'ormconfig';
  const connectionOptions = new ConnectionOptionsReader({ configName }).all();
  const [connectionOption] = await connectionOptions;

  return createConnection(
    Object.assign(connectionOption, {
      namingStrategy: new CustomNamingStrategy(),
    })
  );
};
```

직접 연결설정을 입력할 수도 있다

```ts
import { join } from 'path';
import { createConnection, Connection } from 'typeorm';

const connection: Connection = await createConnection({
  type: 'mysql',
  host: 'localhost',
  port: 3306,
  username: 'test',
  password: 'test',
  database: 'test',
  entities: [`${join(__dirname, '..')}/domain/**/*.{js,ts}`],
  namingStrategy: new CustomNamingStrategy(),
});
```

### Custom Naming Stragegy

```ts
import { DefaultNamingStrategy } from 'typeorm';
import { snakeCase } from 'typeorm/util/StringUtils';

export class CustomNamingStrategy extends DefaultNamingStrategy {
  columnName(propertyName: string, customName: string, embeddedPrefixes: string[]): string {
    return snakeCase(embeddedPrefixes.join('_')) + (customName || snakeCase(propertyName));
  }

  relationName(propertyName: string): string {
    return snakeCase(propertyName);
  }

  joinColumnName(relationName: string, referencedColumnName: string): string {
    return snakeCase(`${relationName}_${referencedColumnName}`);
  }

  joinTableColumnName(tableName: string, propertyName: string, columnName?: string): string {
    return snakeCase(`${tableName}_${columnName || propertyName}`);
  }
}
```

### Custom Logger

```ts
import { Logger } from 'typeorm';

export class MyCustomLogger implements Logger {
  // implement all methods from logger class
}
```

### 연결 객체

- `createConnection()` / `createConnections()`: 연결

- `getConnectionManager()`: 생성된 db연결을 보유하고 있는 객체 / get()으로 연결을 얻음

- `getConnection()`: db 연결을 얻음

  - `getEntityManager({connection-name?})`: 연결로부터 entity manager를 얻음
  - `getRepository({Type}, {connection-name?})`: 연결로부터 repository를 얻음
  - `getTreeRepository({Type}, {connection-name?})`: 연결로부터 tree repository를 얻음
  - `getCustomRepository({Type}, {connection-name?})`: `AbstractRepository<T>`를 상속하거나 클래스에 `@EntityRepository()`를 사용하여 정의한 사용자 정의 repository를 얻는다

## 작동방식

- `EntityManager` & `Repository` Interface

  - `Repository` 내부에서 실제로는 `EntityManager` 호출
  - <https://github.com/typeorm/typeorm/blob/master/src/entity-manager/EntityManager.ts>

- 메소드 `Executor`

  - `EntityManager`는 단순 작업은 직접 `QueryBuilder`를 호출하고 복잡한 작업은 `Executor`를 통해 `QueryBuilder`를 호출함
  - <https://github.com/typeorm/typeorm/blob/master/src/persistence/SubjectExecutor.ts>

## 대용량 처리

TypeOrm에서 대용량 작업을 `chunk` 단위로 나누어 처리할 수 있다

- <https://typeorm.io/#/repository-api/additional-options>
- 해당옵션은 `SaveOptions` & `RemoveOptions`에 포함되어 있다

> 실제 처리되는 코드는 다음에서 확인: <https://github.com/typeorm/typeorm/blob/master/src/persistence/EntityPersistExecutor.ts>

데이터 처리시 `insert`, `remove` 메소드는 단순한 작업만을 수행하므로(쿼리빌더 호출),
성능상 이점이 있지만 조건이 많은 요구사항을 처리하기는 부적합하다.

대용량 작업이 필요하다면 직접 대상 배열을 쪼갠뒤 `insert`, `delete` 메소드를 순차적으로 호출해도 된다

## 활용

### raw query with parameters

<https://github.com/typeorm/typeorm/blob/master/src/driver/mysql/MysqlDriver.ts#L396>

```ts
// 테스트 위해 id = 1 고정
const [query, parameters] = getConnection().driver.escapeQueryWithParameters(
  `
    INSERT INTO post(id, title, content, author, hit)
    VALUES(1, :title, :content, :author, 0)
    ON DUPLICATE KEY UPDATE hit = hit + 1
  `,
  { title: '제목', content: '본문', author: '작성자' }, // parameters
  {} // native parameters
);

const result = await getManager().query(query, parameters);
```

### Transaction 예시

```ts
export class PostController {
  @Transaction('mysql') // "mysql" is a connection name. you can not pass it if you are using default connection.
  async save(post: Post, category: Category, @TransactionManager() entityManager: EntityManager) {
    await entityManager.save(post);
    await entityManager.save(category);
  }

  // this save is not wrapped into the transaction
  async nonSafeSave(entityManager: EntityManager, post: Post, category: Category) {
    await entityManager.save(post);
    await entityManager.save(category);
  }

  @Transaction('mysql') // "mysql" is a connection name. you can not pass it if you are using default connection.
  async saveWithRepository(
    post: Post,
    category: Category,
    @TransactionRepository(Post) postRepository: Repository<Post>,
    @TransactionRepository() categoryRepository: CategoryRepository
  ) {
    await postRepository.save(post);
    await categoryRepository.save(category);

    return categoryRepository.findByName(category.name);
  }

  @Transaction({ connectionName: 'mysql', isolation: 'SERIALIZABLE' }) // "mysql" is a connection name. you can not pass it if you are using default connection.
  async saveWithNonDefaultIsolation(
    post: Post,
    category: Category,
    @TransactionManager() entityManager: EntityManager
  ) {
    await entityManager.save(post);
    await entityManager.save(category);
  }
}
```

### columns > value-transformer

<https://github.com/typeorm/typeorm/blob/master/test/functional/columns/value-transformer/value-transformer.ts>

```ts
const lowercase: ValueTransformer = {
  to: (entityValue: string) => {
    return entityValue.toLocaleLowerCase();
  },
  from: (databaseValue: string) => {
    return databaseValue;
  },
};

const encode: ValueTransformer = {
  to: (entityValue: string) => {
    return encodeURI(entityValue);
  },
  from: (databaseValue: string) => {
    return decodeURI(databaseValue);
  },
};

const encrypt: ValueTransformer = {
  to: (entityValue: string) => {
    return Buffer.from(entityValue).toString('base64');
  },
  from: (databaseValue: string) => {
    return Buffer.from(databaseValue, 'base64').toString();
  },
};

@Entity()
export class User {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column({ transformer: [lowercase, encode, encrypt] })
  email: string;
}
```

## Troubleshooting

### 테스트 실행 오류

테스트코드에서 db 커넥션을 생성하고 테스트를 실행할 때 다음 오류가 발생하는 경우가 있다

```txt
Test suite failed to run
Cannot add a test after tests have started running. Tests must be defined synchronously.
```

- <https://stackoverflow.com/a/66058192>
- <https://github.com/typeorm/typeorm/issues/1654#issuecomment-368618650>

이는 orm 설정의 entities 경로에 테스트파일이 포함되어 있는 경우 발생할 수 있으며, 설정을 다음과 같이 변경한다

```json
{
  entities: ['src/entities/**/!(*.spec.ts)']
}
```
