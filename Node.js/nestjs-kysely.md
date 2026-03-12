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

const TRX_ALS = Symbol('TRX_ALS');

@Injectable()
export class TrxSession {
  constructor(
    @Inject(TRX_ALS) private readonly storage: AsyncLocalStorage<Transaction<Tables>>,
    private readonly db: Database,
  ) {}

  /** Nested Transaction Error: https://github.com/kysely-org/kysely/blob/2ceae395058da071c4948a632a3576fd4e0390f4/src/kysely.ts#L584 */
  current(): Omit<Transaction<Tables> | DB, 'transaction'> {
    return this.storage.getStore() ?? this.db;
  }

  /** Transactional Propagation = REQUIRED */
  create<R>(callback: (trx: Transaction<Tables>) => Promise<R>): Promise<R> {
    const sessTrx = this.storage.getStore();
    if (sessTrx) {
      return callback(sessTrx);
    }

    return this.db.transaction().execute((trx) => this.storage.run(trx, () => callback(trx)));
  }
}

@Global()
@Module({
  providers: [
    {
      provide: TRX_ALS,
      useValue: new AsyncLocalStorage(),
    },
    TrxSession,
  ],
  exports: [TrxSession],
})
export class TrxModule {}
```
