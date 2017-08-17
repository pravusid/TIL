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

## 