# TypeORM

<http://typeorm.io>

## 설치

```sh
npm install typeorm --save
npm install reflect-metadata --save
npm install @types/node --save
npm install mysql --save
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

프로젝트 루트에 `ormconfig.json`을 작성하고 db연결정보를 기록한다

json 뿐만 아니라 js, yml, .env 등의 다양한 방식으로 작성 가능하다

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
    "entities": [
      "src/entity/**/*.ts"
    ],
    // DDL을 실행하는 MigrationInterface를 구현한 파일을 불러올 위치
    "migrations": [
      "src/migration/**/*.ts"
    ],
    // @After___, @Before___ 같은 entity listeners and subscribers를 불러올 위치
    "subscribers": [
      "src/subscriber/**/*.ts"
    ]
  },
  {
    "name": "another",
    "type": "mysql",
    "host": "localhost",
    "port": 3306,
    "username": "test",
    "password": "test",
    "database": "test",
  }
]
```

mysql/mariadb 관련 연결옵션은 다음을 참조: <http://typeorm.io/#/connection-options/mysql--mariadb-connection-options>

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
const connectionOptions = await (process.env.NODE_ENV === 'production' ? getConnectionOptions()
    : new ConnectionOptionsReader({ configName: 'ormconfig.dev' }).get('default'));
await createConnection(Object.assign(connectionOptions, {
  namingStrategy: new CustomNamingStrategy(),
}));
```

### Custom Naming Stragegy

```ts
import { DefaultNamingStrategy } from 'typeorm';
import { snakeCase } from 'typeorm/util/StringUtils';

export class CustomNamingStrategy extends DefaultNamingStrategy {
  columnName(propertyName: string, customName: string, embeddedPrefixes: string[]): string {
    return snakeCase(embeddedPrefixes.join('_')) + (customName ? customName : snakeCase(propertyName));
  }

  relationName(propertyName: string): string {
    return snakeCase(propertyName);
  }

  joinColumnName(relationName: string, referencedColumnName: string): string {
    return snakeCase(`${relationName}_${referencedColumnName}`);
  }

  joinTableColumnName(tableName: string, propertyName: string, columnName?: string): string {
    return snakeCase(`${tableName}_${columnName ? columnName : propertyName}`);
  }
}
```

### Custom Logger

```ts
import { Logger } from "typeorm";

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

## Transaction 예시

```ts
export class PostController {

    @Transaction("mysql") // "mysql" is a connection name. you can not pass it if you are using default connection.
    async save(post: Post, category: Category, @TransactionManager() entityManager: EntityManager) {
        await entityManager.save(post);
        await entityManager.save(category);
    }

    // this save is not wrapped into the transaction
    async nonSafeSave(entityManager: EntityManager, post: Post, category: Category) {
        await entityManager.save(post);
        await entityManager.save(category);
    }

    @Transaction("mysql") // "mysql" is a connection name. you can not pass it if you are using default connection.
    async saveWithRepository(
        post: Post,
        category: Category,
        @TransactionRepository(Post) postRepository: Repository<Post>,
        @TransactionRepository() categoryRepository: CategoryRepository,
    ) {
        await postRepository.save(post);
        await categoryRepository.save(category);

        return categoryRepository.findByName(category.name);
    }

    @Transaction({ connectionName: "mysql", isolation: "SERIALIZABLE" }) // "mysql" is a connection name. you can not pass it if you are using default connection.
    async saveWithNonDefaultIsolation(post: Post, category: Category, @TransactionManager() entityManager: EntityManager) {
        await entityManager.save(post);
        await entityManager.save(category);
    }

}
```
