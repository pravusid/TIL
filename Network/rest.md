# REST (Representational State Transfer)

분산 하이퍼미디어 시스템(웹)을 위한 아키텍처 스타일

[MS API 디자인 가이드](https://docs.microsoft.com/ko-kr/azure/architecture/best-practices/api-design)를 참고하여 서술함

## REST 기본 디자인 원칙

2000년 Roy Fielding은 웹 서비스 디자인 아키텍처 접근방식으로 REST를 제안하였다.
REST는 프로토콜 독립적이지만 대부분의 REST 구현에서 HTTP 프로토콜을 사용하므로 HTTP REST API 디자인에 중점을 두어 설명함.

다음은 RESTful API의 기본 디자인 원칙이다

- Client-Server:
  - 서로의 관심사를 분리하여 독립적으로 기능을 발전시키고 추가할 수 있어야 함
  - REST API는 리소스를 중심으로 디자인 되며, 리소스에는 클라이언트에서 액세스할 수 있는 모든 개체, 데이터, 서비스가 포함된다.
  - 클라이언트가 리소스의 표현을 교환하여 서비스와 상호작용한다. 많은 HTTP API가 교환형식으로 JSON을 사용한다.

- Stateless:
  - REST API는 상태 비저장 요청 모델을 사용한다.
  - HTTP 요청은 독립적이고 임의순서로 발생하므로 요청 사이 상태정보를 유지할 수 없다.
  - 정보는 리소스 자체에만 저장된다.

- Cache:
  - 모든 서버 응답은 캐시 가능여부를 알 수 있어야 하며 캐시를 고려한 설계가 필요하다.

- Uniform Interface:
  - Identification of resource: 리소스를 고유하게 식별하는 식별자인 URI가 있다.
  - manipulation of resources through representation:
    - 리소스의 표현계층을 리소스의 식별자(URI)로 부터 분리한다.
    - HTTP 기반 REST API의 경우 리소스에 표준 HTTP 수행작업을 사용하는 균일한 인터페이스를 사용한다. (GET, POST, PUT, PATCH, DELETE...)
  - self-descriptive message: 각각의 요청에 필요한 정보가 모두 담겨야 한다.
  - hypermedia as the engine of application state(HATEOAS): REST API는 표현에 포함된 하이퍼미디어 링크에 따라 구동된다.

      ```json
      {
        "orderID": 3,
        "productID": 2,
        "quantity": 4,
        "orderValue": 16.6,
        "links": [
          {
            "rel": "product",
            "href": "https://adventure-works.com/customers/3",
            "action": "GET"
          },
          {
            "rel": "product",
            "href": "https://adventure-works.com/customers/3",
            "action": "PUT"
          }
        ]
      }
      ```

  - 이를 통해 내부 구현과 관계없이 어느 클라이언트에서든 플랫폼 독립적으로 API를 호출할 수 있어야 함 (표준 프로토콜 및 교환 가능한 데이터 형식 사용)

- Layered System:
  - 클라이언트 / 서버에 미들웨어 구성요소를 추가할 수 있는 구조이다
  - 서버와 클라이언트 간 상호작용을 일관성 있게 유지해야 한다.

- Code-On-Demand (Optional):
  - 서버가 네트워크를 통해 클라이언트에 프로그램을 전달하면 클라이언트에서 그 프로그램이 실행될 수 있어야 한다.
  - i.e. Java applet, JavaScript...

## 성숙도 모델

2008년에 Leonard Richardson은 Web API에 대한 다음과 같은 [성숙도 모델](https://martinfowler.com/articles/richardsonMaturityModel.html)을 제안하였다.

- 수준 0: 한 URI를 정의하고, 모든 작업은 이 URI에 대한 POST 요청이다
- 수준 1: 개별 리소스에 대한 별도 URI를 만든다
- 수준 2: HTTP 메소드를 사용하여 리소스에 대한 작업을 정의한다
- 수준 3: 하이퍼미디어(HATEOAS)를 사용한다
- 수준 3이어야 Roy Fielding이 정의한 RESTful API이지만 실제로 사용되는 다수의 Web API는 수준 2의 어디인가에 해당한다.

## API 디자인 가이드

### 리소스 중심 API 구성

비즈니스 entity에 집중해야 한다.

예를 들어, 전자 상거래 시스템의 기본 entity는 고객과 주문이다.
주문 정보가 포함된 HTTP POST 요청을 전송하여 주문을 생성한다.

`https://adventure-works.com/orders`

> 가능하다면 리소스 URI는 동사(리소스에 대한 작업)가 아닌 명사(리소스)를 기반으로 해야 한다

REST의 목적은 entity에서 애플리케이션이 수행할 수 있는 작업을 모델링 하는 것이므로 단순히 내부 구조를 반영하는 API를 만들면 안된다.

URI에 일관성있는 명명 규칙을 적용한다. 일반적으로 컬렉션 및 항목에 대한 URI를 계층 구조로 구성하는 것이 좋다.
예를 들어, `/customers`는 고객 컬렉션의 경로이고 `/customers/5`는 ID가 5인 고객의 경로이다.

서로 다른 리소스와의 관계도 고려해야 한다.
예를 들어 `/customer/5/orders`는 고객 5에 대한 모든 주문을 나타낼 수 있다.
반대로 `/orders/99/customer`는 주문 99의 고객을 나타낼 수 있다.

HTTP 응답 메시지 본문에 연결된 리소스에 대한 탐색 가능한 링크(HATEOAS)를 제공하는 방법이 좋다.

복잡한 시스템에서는 `/customer/1/orders/99/products`처럼 여러 관계 수준을 탐색할 수 있는 URI를 제공하고 싶을 수 있다.
그러나 복잡성은 유지보수성을 낮추고 리소스간 관계가 변하면 유연성이 떨어진다.

> 리소스 URI를 `/컬렉션/항목/컬렉션` 수준보다 복잡한 수준으로 요구하지 않는 것이 좋다

모든 웹 요청은 서버의 부하를 높이므로 다수의 작은 리소스를 표시하는 번잡한 Web API를 피하도록 노력해야 한다.
클라이언트 애플리케이션이 요구하는 모든 데이터를 찾기 위해 여러 요청을 보내는 대신,
데이터를 비정규화하고 단일 요청을 통해 관련 정보를 표시하는 더 큰 리소스로 결합하는 것이 좋다.
단, 이 접근 방식과 클라이언트에 필요 없는 데이터를 가져오는 오버헤드 사이의 균형을 맞춰야 한다.

Web API와 데이터 원본 사이에 종속성이 발생하지 않도록 해야 한다.
예를 들어, 데이터가 관계형 데이터베이스에 저장되는 경우 API는 각 테이블을 리소스 컬렉션으로 표시할 필요가 없다.
필요하다면 데이터베이스와 API 사이 매핑 계층을 도입하여 인터페이스가 데이터베이스 스키마 변경과 독립적이어야 한다.

### HTTP 메소드를 기준으로 작업 정의

HTTP 프로토콜은 요청에 체계의미를 할당하는 다양한 메소드를 정의한다.

- GET: 지정된 URI에서 리소스 표현을 검색한다. 응답 메시지의 본문은 요청한 리소스의 세부 정보를 포함한다.
- POST: 지정된 URI에 새 리소스를 만든다. 요청 메시지의 본문은 새 리소스의 세부정보를 제공한다. POST를 사용하여 실제로 리소스를 만들지 않는 작업을 트리거할 수도 있다.
- PUT: 지정된 URI에 리소스를 만들거나 대체한다. 요청 메시지의 본문은 만드럭나 업데이트할 리소스를 지정한다.
- PATCH: 리소스의 부분 업데이트를 수행한다. 요청 본문은 리소스에 적용할 변경 내용을 지정한다.
- DELETE: 지정된 URI의 리소스를 제거한다.

특정 요청의 효과는 리소스가 컬렉션인지 개별 항목인지에 따라 달라진다.
다음 표는 RESTful 구현에서 채택하는 일반적인 규칙을 보여준다.

| 리소스 | POST | GET | PUT | DELETE |
| --- | --- | --- | --- | --- |
| `/customers` | 새 고객 생성 | 모든 고객 검색 | 고객 일괄 업데이트 | 모든 고객 제거 |
| `/customers/1` | **오류** | 고객 1 세부 정보 | 고객 1의 세부 정보 업데이트 | 고객 1 제거 |
| `/customers/1/orders` | 고객 1의 새 주문 생성 | 고객 1의 모든 주문 검색 | 고객 1의 주문 일괄 업데이트 | 고객 1의 모든 주문 제거 |

POST, PUT, PATCH 차이점을 구분하기 어려울 수 있다

- POST
  - 리소스를 생성한다
  - 일반적으로 컬렉션에 POST 요청을 자주 하고 새 리소스는 컬렉션에 추가된다
  - 새 리소스를 만들지 않고 기존 리소스에 처리할 데이터를 보낼 목적으로 사용할 수도 있다
- PUT
  - 리소스를 생성하거나 기존 리소스를 업데이트 한다
  - 컬렉션 보다는 개별 리소스에 자주 적용된다
  - 리소스의 URI를 지정하고, 요청 본문에는 리소스의 완전한 표현이 포함된다
  - URI를 사용하는 리소스가 있으면 리소스가 대체되고 아직 없다면 새 리소스를 생성한다(서버에서 지원한다면)
- PATCH
  - 기존 리소스에 **부분 업데이트**를 수행한다
  - 리소스의 URI를 지정하고, 요청 본문에는 리소스에 적용할 변경내용을 지정한다
  - 리소스의 전체 표현이 아닌 변경 내용만 보내므로 PUT 보다 효율적일 수 있다
  - 서버에서 리소스 생성을 지원하는 경우 새 리소스를 만들수도 있다(null 리소스를 업데이트)

클라이언트가 동일한 PUT 요청을 여러번 제출하여도 결과는 항상 같아야한다.
POST, PATCH 요청은 여러번 제출하여도 반드시 결과가 같을 필요는 없다.

### HTTP 의미 체계 준수

#### 미디어 유형

HTTP 프로토콜에서 형식은 MIME 유형이라고도 하는 미디어 유형을 사용하여 지정된다.
binary가 아닌 데이터의 경우 대부분 JSON, XML을 지원한다.

요청 또는 응답의 Content-Type 헤더는 표현 형식을 지정한다.
다음은 JSON 데이터를 포함하는 POST 요청 예제이다.

```http
POST https://adventure-works.com/orders HTTP/1.1
Content-Type: application/json; charset=utf-8
Content-Length: 57

{"Id":1,"Name":"Gizmo","Category":"Widgets","Price":1.99}
```

서버에서 미디어 유형을 지원하지 않으면 HTTP 상태코드 415(Unsupported Media Type)를 반환해야 한다.

클라이언트 요청에는 클라이언트가 응답 메시지에서 서버로부터 받을 미디어 유형 목록을 포함하는 Accept 헤더가 포함될 수 있다.

서버가 나열된 미디어 유형 중 어떤 것도 일치시킬 수 없는 경우 HTTP 상태코드 406(Not Acceptable)을 반환해야 한다.

#### GET 메소드

성공적인 GET 메소드는 일반적으로 HTTP 상태코드 200(OK)을 반환한다.
리소스를 찾을 수 없는 경우 상태코드 404(Not Found)를 반환해야 한다.

#### POST 메소드

POST 메소드는 새 리소스를 만드는 경우 HTTP 상태코드 201(Created)을 반환한다.
새 리소스의 URI는 응답의 Location 헤더에 포함된다. 응답 본문은 리소스의 표현을 포함한다.

일부 처리를 수행하지만 새 리소스를 만들지 않는 경우 메소드는 HTTP 상태코드 200(OK)을 반환하고 작업 결과를 응답 본문에 포함할 수 있다.
혹은 반환할 결과가 없다면 메소드는 응답 본문 없이 상태코드 204(No Content)를 반환할 수 있다.

클라이언트가 잘못된 데이터를 요청에 포함하면 서버에서 HTTP 상태코드 400(Bad Request)을 반환한다.
응답 본문에는 오류에 대한 정보 또는 정보를 제공하는 URI 링크가 포함될 수 있다.

#### PUT 메소드

PUT 메소드는 POST 메소드와 마찬가지로 새 리소스를 만드는 경우 HTTP 상태코드 201(Created)를 반환한다.
기존 리소스를 업데이트하는 경우 200(OK) 또는 204(No Content)를 반환한다.

경우에 따라 기존 리소스를 업데이트 할 수 없는 경우도 있는데 이 경우 HTTP 상태코드 409(Conflict)를 반환할 수도 있다.

#### PATCH 메소드

클라이언트는 PATCH 요청시 리소스 전체가 아니라 적용할 변경내용만 포함하여 보낸다.
PATCH 메소드의 사양에는 패치 문서의 형식을 특정하지 않았으므로 요청의 미디어 형식에서 추론해야 한다.

Web API에 대한 일반적인 데이터 형식은 JSON이고 두 가지 주요 JSON 기반 패치 형식으로 JSON 패치 및 JSON 병합 배치가 있다.

[JSON 병합패치](https://tools.ietf.org/html/rfc7396)는 원래 리소스와 동일한 구조를 가진 JSON의 변경/추가할 필드의 하위 집합만 포함한다.
또한 필드값에 `null`을 지정하여 필드를 삭제할 수 있다. (리소스가 명시적 null 값을 가질 수 있으면 병합 패치가 적합하지 않음)

JSON 병합 패치의 미디어 유형은 `application/merge-patch+json`이다.

[JSON 패치](https://tools.ietf.org/html/rfc6902)는 보다 유연하다.
작업의 결과로 적용할 변경 내용을 지정하는대, 작업에는 추가, 제거, 바꾸기, 복사 및 테스트(값의 유효성 검사)가 포함된다.

JSON 패치의 미디어 유형은 `application/json-patch+json`이다.

다음은 PATCH 요청을 처리할 때 발생할 수 있는 몇 가지 일반적 오류 조건이다

- 지원되지 않는 패치 문서 형식: 415(Unsupported Media Type)
- 패치 문서의 형식이 오류: 400(Bad Request)
- 패치 문서가 유효하지만 현재 상태에서는 변경 내용을 리소스에 적용할 수 없음: 409(Conflict)

#### DELETE 메소드

삭제 작업이 성공하면 웹 서버는 성공적 처리를 응답 본문에 추가 정보가 포함되지 않는 HTTP 상태코드 204(No Content)로 응답해야 한다.

리소스가 없는 경우 웹서버는 상태코드 404(Not Found)를 반환할 수 있다.

#### 비동기 작업

경우에 따라 POST, PUT, PATCH, DELETE 작업을 완료 하는데 시간이 소요되어
작업이 완료될 때 까지 기다렸다가 클라이언트에 응답을 보내는 경우 허용되지 않는 수준의 대기시간이 발생할 수 있다.

이 경우 비동기 작업을 수행하는 방안을 고려해보아야 한다.
요청 처리가 수락되었지만 아직 완료되지 않았음을 나타내는 HTTP 상태코드 202(Accepted)를 반환한다.

클라이언트가 상태 엔드포인트를 폴링하여 상태를 모니터링할 수 있도록 비동기 요청 상태를 반환하는 엔드포인트를 표시해야 한다.
202 응답의 Location 헤더에 상태 엔드포인트 URI를 포함한다.

```http
HTTP/1.1 202 Accepted
Location: /api/status/12345
```

클라이언트가 GET 요청을 보내는 경우 응답에 요청의 현재 상태가 포함되어야 한다.
필요에 따라 예상 완료시간 또는 작업 취소 링크를 포함할 수 있다.

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "status":"In progress",
  "link": { "rel":"cancel", "method":"delete", "href":"/api/status/12345" }
}
```

비동기 작업에서 새 리소스를 생성하는 경우 작업 완료 후 상태 엔드포인트에서 상태코드 303(See Other)를 반환해야 한다.
303 응답에 새 리소스의 URI를 제공하는 Location 헤더를 포함한다.

```http
HTTP/1.1 303 See Other
Location: /api/orders/12345
```

### 데이터 필터링 및 페이지 매기기

단일 URI를 통해 리소스 컬렉션을 표시하면 정보의 하위 집합만 필요할 때 이에 맞춰 대량의 데이터를 가져올 수 있다.
예를 들어, `/orders?min_cost=n` 처럼 API가 URI의 query-string에서 조건을 전달할 수 있다.

컬렉션 리소스에 대한 GET 요청은 단일 요청에서 반환되는 데이터 양이 제한되도록 Web API를 디자인해야 한다.

```http
/orders?limit=25&offset=50
```

클라이언트 애플리케이션을 돕기위해, 페이지가 매겨진 데이터를 반환하는 GET 요청은 컬렉션의 총 리소스 수를 나타내는 메타데이터를 포함해야 한다.

`/orders?sort=product_id` 같이 정렬 매개 변수를 제공하여 데이터를 가져올 때 정렬하는 전략을 사용할 수 있다.
그러나 쿼리 문자열 매개변수는 여러 캐시 구현에서 캐시된 데이터의 키로 사용되는 리소스 식별자의 일부를 구성하므로 캐싱에 나쁜 영향을 미칠 수 있다.

각 항목에 대량의 데이터가 포함된 경우 각 항목에 대해 반환되는 필드를 제한하도록 할 수 있다.
예를 들어, 쉽표로 구분된 필드목록을 받는 `/orders?fields=product_id,quantity` 같은 쿼리 문자열 매개변수를 사용할 수 있다.

쿼리 문자열의 모든 선택적 매개변수에 의미 있는 기본 값을 제공한다.
예를 들어, 페이지를 구현하는 경우 `limit` 매개변수를 10으로, `offset` 매개변수를 0으로 설정하고,
주문을 구현하는 경우 정렬 매개변수를 리소스의 키로 설정하고,
프로젝션을 지원하는 경우 `fields` 매개 변수를 리소스의 모든 필드로 설정한다.

### 대용량 이진 리소스에 대한 부분 응답 지원

리소스에 파일 또는 이미 같은 대용량 이진 필드가 포함될 수 있다.
신뢰할 수 없는 간헐적 연결에 의한 문제를 해결하고 응답시간을 개선하려면 이러한 리소스를 chunk로 검색할 수 있는 방안을 고려해야 한다.

이렇게 하려면 Web API가 대용량 리소스의 GET 요청에 대해 Accept-Ranges 헤더를 지원해야 한다.
이 헤더는 GET 작업이 부분 요청을 지원한다는 것을 나타낸다.

또한 이런 리소스에 대해 HTTP HEAD 요청을 구현하는 방안을 고려해야 한다.
HEAD 요청은 리소스에 대해 설명하는 HTTP 헤더만 반환하고 메시지 본문이 비어있다는 점만 빼면 GET 요청과 비슷하다.

클라이언트 애플리케이션은 부분적인 GET 요청을 사용하여 리소스를 가져올지 여부를 결정하는 HEAD 요청을 사용할 수 있다.

```http
HEAD https://adventure-works.com/products/10?fields=productImage HTTP/1.1
```

다음은 응답메시지 예제이다

```http
HTTP/1.1 200 OK

Accept-Ranges: bytes
Content-Type: image/jpeg
Content-Length: 4580
```

Content-Length 헤더는 총 리소스 크기를 제공하고, Accept-Ranges 헤더는 해당 GET 작업이 일부결과를 지원함을 나타낸다.
클라이언트 애플리케이션은 이 정보를 사용하여 더 작은 chunk에서 이미지를 검색할 수 있다.

첫 번째 요청은 범위 헤더를 사용하여 처음 2500 바이트를 가져온다.

```http
GET https://adventure-works.com/products/10?fields=productImage HTTP/1.1
Range: bytes=0-2499
```

응답 메시지는 HTTP 상태코드 206(Partial Content)을 반환하여 이 응답이 부분 응답임을 나타낸다.
Content-Length 헤더는 메시지 본문에 반환된 실제 바이트 수를 지정하며, Content-Range 헤더는 해당 바이트가 리소스의 어느 부분 인지를 나타낸다.

```http
HTTP/1.1 206 Partial Content

Accept-Ranges: bytes
Content-Type: image/jpeg
Content-Length: 2500
Content-Range: bytes 0-2499/4580

...
```

### HATEOAS를 사용한 관련 리소스 탐색
