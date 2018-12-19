# NGINX

Web Server (Apache와 2강 구도를 형성하고 있다)

- Apache : 요청을 Multi Processing Module 방식으로 처리
  - Prefork: 실행중인 프로세스를 복제하여 처리
  - Worker: 여러 프로세스가 여러 쓰레드를 사용한다

- Nginx : Event Driven 방식으로 요청을 비동기 처리함

## 설치 및 실행

- 설치: `sudo apt install nginx`
- 실행: `sudo service nginx start`

## 명령

- `nginx -s stop`: 바로 종료 (TERM)
- `nginx -s quit`: 정상 종료 (QUIT)
- `nginx -s reload`: 환경설정을 다시 읽음
- `nginx -s reopen`: 로그 파일을 다시 연다

## 설정

- 설정파일 위치: `/etc/nginx/nginx.conf`
- 환경설정 구문 체크: `nginx -t`

### 기본 설정파일 내용

```conf
user www-data;
worker_processes auto;
pid /run/nginx.pid;
include /etc/nginx/modules-enabled/*.conf;

events {
    worker_connections 768;
    # multi_accept on;
}

http {

    ##
    # Basic Settings
    ##

    sendfile on;
    tcp_nopush on;
    tcp_nodelay on;
    keepalive_timeout 65;
    types_hash_max_size 2048;
    # server_tokens off;

    # server_names_hash_bucket_size 64;
    # server_name_in_redirect off;

    include /etc/nginx/mime.types;
    default_type application/octet-stream;

    ##
    # SSL Settings
    ##

    ssl_protocols TLSv1 TLSv1.1 TLSv1.2; # Dropping SSLv3, ref: POODLE
    ssl_prefer_server_ciphers on;

    ##
    # Logging Settings
    ##

    access_log /var/log/nginx/access.log;
    error_log /var/log/nginx/error.log;

    ##
    # Gzip Settings
    ##

    gzip on;

    # gzip_vary on;
    # gzip_proxied any;
    # gzip_comp_level 6;
    # gzip_buffers 16 8k;
    # gzip_http_version 1.1;
    # gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;

    ##
    # Virtual Host Configs
    ##

    include /etc/nginx/conf.d/*.conf;
    include /etc/nginx/sites-enabled/*;
}
```

### HTTP 모듈

HTTP 모듈 설정은 세 가지 계층 블럭을 제공한다

- http (프로토콜 수준)
  - HTTP와 관련된 모든 모듈의 지시어를 정의함

- server (서버 수준)
  - 하나의 웹사이트를 선언하는 데 사용
  - http 블럭 안에서만 사용 가능

- location (요청 URI 수준)
  - 웹사이트의 특정 위치에 적용할 설정 그룹 정의
  - server 블럭이나 다른 location 블럭 안에 삽입할 수 있다

## 리버스 프록시

`/etc/nginx/sites-available/default`

### 요청 흐름제어

server 아래의 location을 다음과 같이 수정

- `proxy_pass`
  - 요청이 오면 해당 위치로 전달
  - <http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass>

- `proxy_set_header <헤더> <헤더값>`
  - 요청의 헤더에 정의한 값 할당
  - <http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_set_header>

```text
location / {
    proxy_pass       http://localhost:8000;
    proxy_set_header Host      $host;
    proxy_set_header X-Real-IP $remote_addr;
}
```

해당 위치를 변수로 설정하고 싶으면

```text
include /etc/nginx/conf.d/service-url.inc;

location / {
    proxy_pass $service_url;
}
```

`/etc/nginx/conf.d/service-url.inc` 파일 내용은 다음으로 한다

```text
set $service_url http://localhost:8000;
```

### 로드밸런싱

```text
upstream test_proxy {
    <LOAD_BALANCING_METHOD>
    server web-01;
    server web-02;
}
server {
    listen 80 default_server;
    listen [::]:80 default_server;

    root /var/www/html;

    index index.html index.htm index.nginx-debian.html;

    server_name _;

    location / {
        proxy_pass http://test_proxy;
    }
}
```

로드밸런싱 메소드(`<LOAD_BALANCING_METHOD>` 위치)

- default: round-robin
- least_conn: 연결이 적은곳으로
- ip_hash: 클라이언트 IP주소를 기준으로 요청 분배
- hash: 유저정의 변수 조합을 해싱하여 기준으로 사용
- least_time: 평균 레이턴시와 연결을 기준으로
