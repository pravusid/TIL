# Dynamo DB

종합 관리형 NoSQL 데이터베이스 서비스

테이블 / 항목 / 속성으로 구성됨

## 구조

### 기본키

테이블에는 기본키가 필요하고(고유식별자) 두 가지를 지원함

- 파티션 키(단순 기본키)
- 파티션 키 및 정렬 키(복합 기본키)
- Partition Key는 고유 값이 많고(cardinality) 균일 비율로 무작위 요청 되는 속성이 좋다
- Sort Key는 관계 모델링, 선택적 조회, 범위 조회를 위해 사용한다

단기적인 워크로드 불균형 문제를 완화하기 위해 적응형 용량이 5~30분 내에 활성화됨.
하지만 각 파티션에는 1,000개의 쓰기 용량 유닛 및 3,000개의 읽기 용량 유닛 제한이 계속 적용되므로 테이블 또는 파티션 설계에서 발생하는 더 큰 용량 문제는 적응형 용량으로 해결할 수 없음.

AWS SDK에서는 스로틀된 요청은 기본적으로 10번 재시도 함

- <https://github.com/aws/aws-sdk-js/blob/master/lib/services/dynamodb.js>
- <https://github.com/aws/aws-sdk-js/blob/master/test/services/dynamodb.spec.js>

### 보조 인덱스

테이블에서 종류별 최대 5개의 보조 인덱스를 생성할 수 있다

- Global secondary index: 파티션 키 및 정렬 키가 테이블의 파티션 키 및 정렬 키와 다를 수 있는 인덱스

  - 인덱스 크기 제약 없음
  - 테이블 생성 이후에도 생성 및 삭제가 가능하다
  - Eventual consistent read만 가능함
  - 별도의 읽기&쓰기 용량이 할당된다 (쓰기 용량이 부족하면 병목 발생)
  - <https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/GSI.html>

- Local secondary index: 테이블과 파티션 키는 동일하지만 정렬 키는 다른 인덱스
  - 인덱스는 10GB 단위인 파티션 데이터와 함께 저장함
  - 테이블 생성 이후에는 생성 및 삭제 불가
  - 테이블에 할당된 읽기&쓰기 용량을 사용함
  - Eventual consistent read || Strong consistent read 선택가능
  - <https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/LSI.html>

### 용량 (Capacity Unit)

- 1RCU

  - Eventual consistent read: 8KB/sec 처리
  - Strong consistent read: 4KB/sec 처리
  - 4KB 단위로 올림계산 (4KB block)

- 1WCU
  - 1KB/sec 처리
  - 1KB 단위로 올림계산 (1KB block)

트랜잭션 처리의 경우 2배의 용량이 필요(읽기의 경우 Strong consistent read 기준)

### 예약어

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/ReservedWords.html>

## 입출력

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/WorkingWithItems.html>

1MB 단위로 반환함 (1MB 보다 크면 잘라서 처리함)

### 비교연산자 / 조건식

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/Expressions.OperatorsAndFunctions.html>

## 모델링

> 1 애플리케이션 당 1테이블 생성

### GSI 오버로딩

다양한 패턴으로 데이터를 조회

- PK: 일반적인 조건으로 둔다 (cardinality 및 분포비율 고려)
- SK: 기본 데이터는 `master` 값으로 두고 조회하고 나머지 조건은 `key:value`로 저장한다
  - 기본 SK를 GSI의 PK로 설정하고 필요에 따라 GSI의 SK와 프로젝션 설정

### 결합 정렬키

- 계층 구조를 정의함
- 특정 조건에 맞는 데이터를 빠르게 조회하고 조회 복잡도를 줄임

> i.e. 국가, 주, 도시, 위치 별로 건물 검색

- PK: country: KOREA
- SK: SEOUL#MAPO#HAPJUNG

검색시 SK에서 `begin_with`를 사용한다

### DynamoDB 스트림

DynamoDB 테이블의 데이터 수정 이벤트가 발생한 순서대로 거의 실시간으로 스트림에 표시됨

AWS Lambda와 연결하여 거의 실시간으로 관심있는 정보를 처리할 수 있다

### 쓰기 샤딩

PK 분포도가 고르지 않을 경우 샤딩크기를 관리하는 테이블을 별도로 관리하여 파티션키를 분할하여 저장함

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

jar 실행

`java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb`

docker 컨테이너 실행

`docker run -d -p 8000:8000 --name dynamo amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb`

- `-sharedDb` 옵션을 사용하면 DynamoDB에서 이름이 shared-local-instance.db인 단일 데이터베이스 파일이 생성함
- `-sharedDb`를 누락하면 데이터베이스 파일의 이름은 myaccesskeyid_region.db로 설정되고, AWS 액세스 키 ID 및 리전은 애플리케이션 구성에 표시된 대로 지정됨
- `-inMemory` 옵션을 사용하면 DynamoDB는 데이터베이스 파일을 기록하지 않고 메모리를 사용하며 종료시 데이터는 사라짐
- `-port <port>` 옵션으로 실행시 포트를 변경할 수 있음

shell 접속: <http://localhost:8000/shell>

shell 접속 후 option에서 Access Key Id 설정(옵션사항)

## AWS SDK for DynamoDB

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/CurrentAPI.html>

### Node에서 DynamoDB endpoint 설정

```js
export const awsDynamo = new AWS.DynamoDB.DocumentClient({
  convertEmptyValues: true,
});
```

### Node에서 DynamoDB endpoint 설정 (Local DynamoDB)

```js
const AWS = require('aws-sdk');

AWS.config.update({
  credentials,
  region: 'ap-northeast-2',
});

const dynamoClient = new AWS.DynamoDB({ endpoint: 'http://localhost:8000' });
const dynamoDocClient = new AWS.DynamoDB.DocumentClient({
  endpoint: 'http://localhost:8000',
});
```

또는 AWS config에서 설정할 수도 있다

```js
AWS.config.update({
  credentials,
  region: 'ap-northeast-2',
  dynamodb: {
    endpoint: 'http://localhost:8000',
  },
});
```

### 테이블 생성

```js
dynamodb.createTable(
  {
    TableName: 'cba-test',
    KeySchema: [
      { AttributeName: 'uid', KeyType: 'HASH' },
      { AttributeName: 'createdAt', KeyType: 'RANGE' },
    ],
    AttributeDefinitions: [
      { AttributeName: 'uid', AttributeType: 'S' },
      { AttributeName: 'createdAt', AttributeType: 'N' },
    ],
    ProvisionedThroughput: {
      ReadCapacityUnits: 5,
      WriteCapacityUnits: 5,
    },
  },
  (err, data) => {
    if (err) ppJson(err);
    // an error occurred
    else ppJson(data); // successful response
  }
);
```

### Scan

전체 데이터 순회

```ts
type EvalKey = { uid: string };

async function readFromDynamo() {
  let lastEvalKey: EvalKey;
  do {
    const { Count, ScannedCount, Items, LastEvaluatedKey } = await dynamoDbClient
      .scan({
        TableName: 'foo-bar',
        ...(lastEvalKey && { ExclusiveStartKey: lastEvalKey }),
      })
      .promise();

    console.log('Count', Count, 'ScannedCount', ScannedCount, 'LastEvaluatedKey', LastEvaluatedKey);
    lastEvalKey = LastEvaluatedKey as EvalKey;
  } while (lastEvalKey);
}
```

### Query

```js
const response = dynamodb
  .query({
    TableName: 'cba-test',
    KeyConditionExpression: '#uid = :uid',
    FilterExpression: '#createdAt >= :createdAt',
    ExpressionAttributeNames: {
      '#uid': 'uid',
      '#createdAt': 'createdAt',
    },
    ExpressionAttributeValues: {
      ':uid': query.uid,
      ':createdAt': query.time,
    },
    ScanIndexForward: false, // ASC || DESC
    Limit: query.limit,
  })
  .promise();
```

### Put

```js
const result = {
  /* content */
};
const response = dynamodb
  .put({
    TableName: 'cba-test',
    Item: result,
    ConditionExpression: 'attribute_not_exists(#uid)',
    ExpressionAttributeNames: {
      '#uid': 'uid',
    },
  })
  .promise();
```

### Update

```js
const response = dynamodb
  .update({
    TableName: 'cba-test',
    Key: {
      uid: result.uid,
      createdAt: result.createdAt,
    },
    UpdateExpression: 'SET #comment = :comment, #updatedAt = :updatedAt',
    ExpressionAttributeNames: {
      '#comment': 'comment',
      '#updatedAt': 'updatedAt',
    },
    ExpressionAttributeValues: {
      ':comment': result.comment,
      ':updatedAt': new Date().getTime(),
    },
  })
  .promise();
```
