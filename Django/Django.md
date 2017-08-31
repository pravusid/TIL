# Django

## 개요

파일명 | 기능
---|---
wsgi.py | gateway
url.py | mapping
views.py | controller
models.py | model
manager | MyBatis
form.py | 폼 관리
template | 화면출력 양식

## 시작하기

1. 프로젝트 생성 `django-admin startproject (프로젝트명)`
1. app 생성 `./manage.py startapp (app이름)`

### settings.py

항목 | 내용
---|---
DEBUG | 디버그 모드설정
INSTALLED_APPS | third-party app 추가
MIDDLEWARE_CLASSES | request와 response 사이의 주요 기능 레이어
TEMPATES | django template 관련설정, 실제 뷰
DATABASES | 데이터베이스 엔진의 연결 설정
STATIC_URL | 정적 파일의 URL(css, javascript, images ...)

### manage.py

프로젝트 관리 명령어 모음

명령어 | 기능
---|---
startapp | 앱 생성
runserver | 서버 실행
createsuperuser | 관리자 생성
makemigrations app | app의 모델 변경사항 체크
migrate | 변경사항을 db에 반영
shell | 쉘을 통해 데이터 확인
collectstatic static | 파일을 한곳에 모음

외부접속 : `./manage.py runserver 0.0.0.0:8080`

## Work Flow

1. settings.py에 생성한 app 등록
1. models.py 작성
1. model을 db에 생성
  ```sh
  ./manage.py makemigrations (app이름)
  ./manage.py migrate
  ```
1. urls.py에 mapping
1. views.py에 (url에서 연결된 함수)작성
1. views와 연결된 templates 작성 (noname.html)
1. forms.py로 모델에서 선언한 폼 생성
1. views.py에서 생성한 폼을 템플릿으로 연결
1. views.py에서 받은 요청이 post면 form값 저장