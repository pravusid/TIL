# Dynamo DB

종합 관리형 NoSQL 데이터베이스 서비스

테이블 / 항목 / 속성으로 구성됨

## 구조

### 기본키

테이블에는 기본키가 필요하고(고유식별자) 두 가지를 지원함

- 파티션 키(단순 기본키)
- 파티션 키 및 정렬 키(복합 기본키)

### 보조 인덱스

테이블에서 종류별 최대 5개의 보조 인덱스를 생성할 수 있다

- Global secondary index: 파티션 키 및 정렬 키가 테이블의 파티션 키 및 정렬 키와 다를 수 있는 인덱스
- Local secondary index: 테이블과 파티션 키는 동일하지만 정렬 키는 다른 인덱스

### DynamoDB 스트림

DynamoDB 테이블의 데이터 수정 이벤트가 발생한 순서대로 거의 실시간으로 스트림에 표시됨

AWS Lambda와 연결하여 거의 실시간으로 관심있는 정보를 처리할 수 있다

### 예약어

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/ReservedWords.html>

## 쿼리

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/WorkingWithItems.html>

### 비교연산자 / 조건식

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/Expressions.OperatorsAndFunctions.html>

## Index

### GSI

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/GSI.html>

### LSI

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/LSI.html>

## 에러처리

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/Programming.Errors.html>

JS error object 예시

```json
{
  "code": "ConditionalCheckFailedException",
  "time": "2019-03-21T14:40:32.171Z",
  "requestId": "94033QGK80J454DT5CR78O862NVV4KQNSO5AEMVJF66Q9ASUAAJG",
  "statusCode": 400,
  "retryable": false,
  "retryDelay": 13.794078870110315
}
```

## 로컬 실행

docker 기준

`docker run -d -p 8000:8000 --name dynamo amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb`

- `-sharedDb` 옵션을 사용하면 DynamoDB에서 이름이 shared-local-instance.db인 단일 데이터베이스 파일이 생성함
- `-sharedDb`를 누락하면 데이터베이스 파일의 이름은 myaccesskeyid_region.db로 설정되고, AWS 액세스 키 ID 및 리전은 애플리케이션 구성에 표시된 대로 지정됨
- `-inMemory` 옵션을 사용하면 DynamoDB는 데이터베이스 파일을 기록하지 않고 메모리를 사용하며 종료시 데이터는 사라짐
- `-port <port>` 옵션으로 실행시 포트를 변경할 수 있음

shell 접속: <http://localhost:8000/shell>

shell 접속 후 option에서 Access Key Id 설정

shell에서 테이블 생성

```js
dynamodb.createTable(
  {
    TableName: "cba-test",
    KeySchema: [
      { AttributeName: "uid", KeyType: "HASH" },
      { AttributeName: "createdAt", KeyType: "RANGE" }
    ],
    AttributeDefinitions: [
      { AttributeName: "uid", AttributeType: "S" },
      { AttributeName: "createdAt", AttributeType: "N" }
    ],
    ProvisionedThroughput: {
      ReadCapacityUnits: 5,
      WriteCapacityUnits: 5
    }
  },
  (err, data) => {
    if (err) ppJson(err);
    // an error occurred
    else ppJson(data); // successful response
  }
);
```

Node에서 DynamoDB endpoint 설정

```js
const AWS = require("aws-sdk");

AWS.config.update({
  credentials,
  region: "ap-northeast-2"
});

const dynamoClient = new AWS.DynamoDB({ endpoint: "http://localhost:8000" });
const dynamoDocClient = new AWS.DynamoDB.DocumentClient({ endpoint: "http://localhost:8000" });
```

또는 AWS config에서 설정할 수도 있다

```js
AWS.config.update({
  credentials,
  region: "ap-northeast-2",
  dynamodb: {
    endpoint: "http://localhost:8000"
  }
});
```

## AWS SDK for DynamoDB

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/CurrentAPI.html>

```js
export const awsDynamo = new AWS.DynamoDB.DocumentClient({
  convertEmptyValues: true
});
```

### Put

```js
const result = {
  /* content */
};
const response = dynamodb
  .put({
    TableName: "cba-test",
    Item: result,
    ConditionExpression: "attribute_not_exists(#uid)",
    ExpressionAttributeNames: {
      "#uid": "uid"
    }
  })
  .promise();
```

### Update

```js
const response = dynamodb
  .update({
    TableName: "cba-test",
    Key: {
      uid: result.uid,
      createdAt: result.createdAt
    },
    UpdateExpression: "SET #comment = :comment, #updatedAt = :updatedAt",
    ExpressionAttributeNames: {
      "#comment": "comment",
      "#updatedAt": "updatedAt"
    },
    ExpressionAttributeValues: {
      ":comment": result.comment,
      ":updatedAt": new Date().getTime()
    }
  })
  .promise();
```

### Query

```js
const response = dynamodb
  .query({
    TableName: "cba-test",
    KeyConditionExpression: "#uid = :uid",
    FilterExpression: "#createdAt >= :createdAt",
    ExpressionAttributeNames: {
      "#uid": "uid",
      "#createdAt": "createdAt"
    },
    ExpressionAttributeValues: {
      ":uid": query.uid,
      ":createdAt": query.time
    },
    ScanIndexForward: false, // ASC || DESC
    Limit: query.limit
  })
  .promise();
```
