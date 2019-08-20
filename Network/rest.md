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
