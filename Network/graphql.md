# GraphQL

- <https://graphql-kr.github.io/learn/>
- <https://graphql.org/learn/>
- <https://spec.graphql.org/>
- [GraphQL IDE Monorepo](https://github.com/graphql/graphiql)

## Refs

- <https://www.apollographql.com/blog/>
- <https://graphql.org/learn/best-practices/>
- <https://www.graphql.com/guides/>
- <https://fe-developers.kakaoent.com/2022/220113-designing-graphql-mutation/>
- [GraphQL API 까짓거 운영해보지 뭐](https://tv.naver.com/v/16969996)
- [GraphQL이 가져온 에어비앤비 프론트앤드 기술의 변천사](https://tv.naver.com/v/16970011)
- [Domain Graph Service를 활용한 광고 서비스의 GraphQL API 구현 사례](https://tv.naver.com/v/23652389)
- [GraphQL 잘 쓰고 계신가요? (Production-ready GraphQL)](https://youtu.be/9G2vT4C4sAY)

## Ecosystem

- <https://graphql.org/code/>
- <https://github.com/chentsulin/awesome-graphql>

### GraphQL Servers

<https://github.com/graphql/graphql-over-http>

> If you want a feature-full server with bleeding edge technologies, you're recommended to use one of the following.
>
> -- <https://github.com/graphql/graphql-http#servers>

- <https://github.com/mercurius-js/mercurius>
- <https://github.com/apollographql/apollo-server>
- <https://github.com/prisma-labs/graphql-yoga>

#### DataLoader

- <https://github.com/graphql/dataloader>
- <https://shopify.engineering/solving-the-n-1-problem-for-graphql-through-batching>
- <https://www.apollographql.com/blog/backend/data-sources/a-deep-dive-on-apollo-data-sources/>
- <https://github.com/kriasoft/relay-starter-kit>

데이터로더는 GraphQL에서 발생하는 N+1 문제를 애플리케이션 로직과 분리하여 해결하기 위한 라이브러리이다

GraphQL에서만 사용할 수 있는 것은 아니므로, 다른 곳에서도 동일한 문제를 해결할 수도 있다.
대표적인 사례로, Prisma에서 데이터로더를 통해 N+1 문제를 해결하고 있다.

<https://www.prisma.io/docs/guides/performance-and-optimization/query-optimization-performance#solving-the-n1-problem>

##### DataLoader 작동방식

<https://github.com/graphql/dataloader#batching>

> 데이터의 키 배열을 파라미터로 받는 일괄 처리함수를 선언하여 로더를 만든다

```js
const DataLoader = require('dataloader');
const userLoader = new DataLoader((keys) => myBatchGetUsers(keys));
```

> 코드 실행 중 로더를 통해 개별 값을 호출하면
> 데이터로더는 단일 실행프레임 (JS에서는 이벤트루프의 단일 틱) 내에서 발생하는 모든 개별 호출을 통합한 뒤
> 일괄처리 함수를 통해 요청한 모든 키에 대한 값을 반환한다

```js
const user = await userLoader.load(1);
const invitedBy = await userLoader.load(user.invitedByID);
console.log(`User 1 was invited by ${invitedBy}`);

// Elsewhere in your application
const user = await userLoader.load(2);
const lastInvited = await userLoader.load(user.lastInvitedID);
console.log(`User 2 last invited ${lastInvited}`);
```

##### DataLoader 일괄처리 함수

> 일괄처리함수는 데이터의 키 배열을 받아서 키에 대응하는 데이터 또는 오류객체 배열의 프로미스를 반환한다.
> 로더는 일괄처리함수에서 `this` context로 제공된다.

```js
async function batchFunction(keys) {
  const results = await db.fetchAllKeys(keys);
  return keys.map((key) => results[key] || new Error(`No result for ${key}`));
}

const loader = new DataLoader(batchFunction);
```

> 일괄처리함수는 다음과 같은 제약이 있다

- 반환 값의 배열 크기는 입력받은 키 배열의 크기와 같아야 한다
- 반환 값의 배열 순서는 입력받는 키의 순서에 대응해야 한다

##### DataLoader 캐싱

<https://github.com/graphql/dataloader#caching>

> 데이터로더는 단일요청(a single request)에서 발생하는 모든 로드(`load`)에 대한 메모이제이션을 제공한다

💡일반적으로 로더는 GraphQL의 컨텍스트에 주입하여 리졸버에서 호출하는 방식으로 사용한다 (개별 요청마다 생성한다, RequestScope)

```js
function createLoaders(authToken) {
  return {
    users: new DataLoader((ids) => genUsers(authToken, ids)),
  };
}

const app = express();

app.get('/', function (req, res) {
  const authToken = authenticateUser(req);
  const loaders = createLoaders(authToken);
  res.send(renderPage(req, loaders));
});

app.listen();
```

### GraphQL Tools

- <https://github.com/ardatan/graphql-tools>
- [schema loading & import expression](https://www.the-guild.dev/graphql/tools/docs/schema-loading#load-graphqlschema-by-using-different-loaders-from-different-sources)

### GraphQL Config

- <https://www.the-guild.dev/graphql/config/docs>
- [VSCode 확장](https://marketplace.visualstudio.com/items?itemName=GraphQL.vscode-graphql)기능 사용에 필요

## Schema-first, Code-first

<https://github.com/graphql/graphql-js>

- Schema-first

  - `schema.graphql` 파일을 문자열로 변환하여 `graphql.buildSchema` 함수에 전달
  - `schema.graphql` 파일을 대상으로 [code generation](https://the-guild.dev/graphql/codegen/docs/getting-started) 실행

- Code-first

  - <https://github.com/MichalLytek/type-graphql>
  - <https://github.com/graphql-nexus/nexus>

### Schema-first & GraphQL Code Generator

- <https://www.the-guild.dev/graphql/codegen/docs/getting-started>

### Code-first & TypeGraphQL

- <https://typegraphql.com/>

## GraphQLError

- <https://github.com/graphql/graphql-js/blob/main/src/error/GraphQLError.ts>
- <https://spec.graphql.org/draft/#sec-Errors>

클라이언트에서 확인 가능한 오류는 `GraphQLFormattedError`이다

<https://github.com/graphql/graphql-js/blob/main/src/error/GraphQLError.ts#L218>

- message: (짧고 이해할 수 있게 요약한) 오류메시지
- locations: 요청 GraphQL doucment(query, mutation)에서 오류가 발생한 위치 (line, column)
- path: 필드 처리중 오류가 발생한 경우 해당 필드 (null 응답을 클라이언트에서 구분하기 위함 == 런타임 오류 or 실제 null 응답)
- extensions: 추가 정보입력을 위한 값 (자유롭게 사용할 수 있고, 구현체에 따라 다름)

### Error in Apollo Server

> v4 부터 `ApolloError` 클래스 및 `toApolloError` 함수는 사라짐
>
> -- <https://www.apollographql.com/docs/apollo-server/migration/#apolloerror>

- <https://github.com/apollographql/apollo-server/blob/main/packages/server/src/errors/index.ts>
- <https://www.apollographql.com/docs/apollo-server/data/errors>
