# Spring Boot REST

## 기본설정

## Controller Mapping for CRUD

- CRUD를 위한 URL mapping
  1. Create : HTTP Method: POST, URL: /user/article
  1. Read : HTTP Method: GET, URL: /user/articles & /user/article/{id}
  1. Update : HTTP Method: PUT, URL: /user/article/{id}
  1. Delete : HTTP Method: DELETE, URL: /user/article/{id}

- hidden사용
  - name="_method" / value="put" / value="delete"

## controller에서 JSON 데이터 처리

클라이언트 request 예제

  ```javascript
  $ajax({
    type: "post",
    contentType: "application/json; charset=UTF-8",
    url: "",
    data: JSON.stringify(json),
    success: function(data, status) {
      /* 결과값으로 기능구현 */
    },
    error: function(xhr, status) {
      /* 결과값으로 기능구현 */
    }
  });
  ```

### 데이터를 String으로 받는경우

컨트롤러 인수로 `@RequestBody String data`를 받는다.

JSON이 문자열로 변환된 형태로 JSON parsing을 하거나 직접 분리해서 사용.

### 데이터를 Map으로 받는경우

컨트롤러 인수로 `@RequestBody Map<String, Object> data`를 받는다.

`data.get("key")`로 value를 꺼내면 `Object`가 나오는데
단일 값이면 `String` 배열이면 `List<LinkedHashMap<String, T>>` 형태의 데이터이다.

### 데이터를 Entity class로 받는경우

컨트롤러 인수로 `@RequestBody EntityClass data`를 받는다.

데이터와 매칭되는 class를 인수로 요청하면 적절한 결과물이 나온다. 가장 편리한 방법이다.