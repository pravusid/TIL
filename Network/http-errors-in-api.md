# HTTP 오류 처리

- <https://datatracker.ietf.org/doc/html/rfc9457>
- <https://datatracker.ietf.org/doc/html/rfc7807>

## 읽을거리

- [Google, Facebook 등 대형 서비스들의 REST API 에러 처리 비교](https://blog.ull.im/engineering/2019/03/14/service-providers-errors.html)
- [IllegalArgumentException은 400 Bad Request인가?](https://techblog.woowahan.com/21686/)

## 상태코드

### <https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/422>

본문 형식은 유효하지만 검증 규칙 위배 (일반적으로는 상태코드 400을 사용하는 경우가 많음)

- <https://stackoverflow.com/questions/16133923/400-vs-422-response-to-post-of-data>
- <https://fastapi.tiangolo.com/tutorial/handling-errors/?h=error#use-the-requestvalidationerror-body>

## 예시

## Stripe

<https://stripe.com/docs/api/errors>

```txt
200 - OK  Everything worked as expected.
400 - Bad Request  The request was unacceptable, often due to missing a required parameter.
401 - Unauthorized  No valid API key provided.
402 - Request Failed  The parameters were valid but the request failed.
403 - Forbidden  The API key doesn't have permissions to perform the request.
404 - Not Found  The requested resource doesn't exist.
409 - Conflict  The request conflicts with another request (perhaps due to using the same idempotent key).
429 - Too Many Requests  Too many requests hit the API too quickly. We recommend an exponential backoff of your requests.
500, 502, 503, 504 - Server Errors  Something went wrong on Stripe's end. (These are rare.)
```

- 전역 오류처리 문서가 있음
- statusCode: 400 ~ 500
- 본문은 `type`, `code`, `decline_code`, `message`... 등으로 구성됨
- errorType 값은 `api_error`, `card_error`, `idempotency_error`, `invalid_request_error`로 구성

## Notion

<https://developers.notion.com/reference/intro>

- 대부분 API 응답은 200, 400이고 400응답일 때 데이터 타입 제공하지 않음 (상태값과 빈 응답이 갈 수도...)

## Slack

<https://api.slack.com/quickstart>

- 성공: `{ "ok": true }`
- 오류: `{ "ok": false, "error": "error type" }`
