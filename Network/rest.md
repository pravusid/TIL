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
