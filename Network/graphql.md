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

## Ecosystem

- <https://graphql.org/code/>
- <https://github.com/chentsulin/awesome-graphql>

### GraphQL Servers

- <https://github.com/graphql/express-graphql>
- <https://github.com/mercurius-js/mercurius>
- <https://github.com/apollographql/apollo-server>
- <https://github.com/prisma-labs/graphql-yoga>

#### DataLoader

- <https://github.com/graphql/dataloader>
- <https://shopify.engineering/solving-the-n-1-problem-for-graphql-through-batching>
- <https://www.apollographql.com/blog/backend/data-sources/a-deep-dive-on-apollo-data-sources/>

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
  - `schema.graphql` 파일을 대상으로 code generation 실행

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
