# NestJS + Kysely

[[nestjs]] + [[kysely]]

- <https://docs.nestjs.com/fundamentals/dynamic-modules#configurable-module-builder>
- <https://wanago.io/2023/08/07/api-nestjs-kysely-postgresql/>
- <https://github.com/kazu728/nestjs-kysely>

## 사용자정의 모듈 선언

```ts
import { ConfigurableModuleBuilder, Global, Module } from '@nestjs/common';
import { CamelCasePlugin, Kysely, MysqlDialect } from 'kysely';
import { PoolOptions, createPool } from 'mysql2';

type DatabaseModuleParams = PoolOptions;

export const { ConfigurableModuleClass, MODULE_OPTIONS_TOKEN } = new ConfigurableModuleBuilder<DatabaseModuleParams>()
  .setClassMethodName('forRoot')
  .build();

export type Tables = {};

export class Database extends Kysely<Tables> {}

@Global()
@Module({
  providers: [
    {
      provide: Database,
      inject: [MODULE_OPTIONS_TOKEN],
      useFactory: (opts: DatabaseModuleParams) => {
        const dialect = new MysqlDialect({
          pool: createPool(opts),
        });
        return new Kysely({
          dialect,
          plugins: [new CamelCasePlugin()],
        });
      },
    },
  ],
  exports: [Database],
})
export class DatabaseModule extends ConfigurableModuleClass {}
```

## 사용자정의 모듈 사용

```ts
import { DatabaseModule } from './database.module';
import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';

@Module({
  imports: [
    ConfigModule.forRoot({ isGlobal: true }),
    DatabaseModule.forRootAsync({
      inject: [ConfigService],
      useFactory: (config: ConfigService) => ({
        host: config.getOrThrow('DB_HOST'),
        port: config.getOrThrow<number>('DB_PORT'),
        user: config.getOrThrow('DB_USERNAME'),
        password: config.getOrThrow('DB_PASSWORD'),
        database: config.getOrThrow('DB_NAME'),
        connectionLimit: config.getOrThrow<number>('DB_CONN_SIZE'),
        timezone: config.getOrThrow('DB_TZ'),
      }),
    }),
  ],
})
export class AppModule {}
```

## TransactionSession 객체

- 트랜잭션 생성 및 관리
- [`AsyncLocalStorage`](https://nodejs.org/api/async_context.html#async_context_class_asynclocalstorage) 통합 가능
- <https://github.com/Papooch/nestjs-cls>

```ts
import { Database, Tables } from './database.module';
import { Global, Injectable, Module } from '@nestjs/common';
import { Transaction } from 'kysely';

@Injectable()
export class TrxSession {
  constructor(private readonly db: Database) {}

  create<R>(callback: (trx: Transaction<Tables>) => Promise<R>): Promise<R> {
    return this.db.transaction().execute(callback);
  }
}

@Global()
@Module({
  providers: [TrxSession],
  exports: [TrxSession],
})
export class TrxModule {}
```
