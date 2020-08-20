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

설정파일 위치: `/etc/nginx/nginx.conf`

환경설정 구문 체크: `nginx -t`

설정 생성후 심볼릭 링크, 적용

```sh
sudo rm /etc/nginx/sites-enabled/default
sudo ln -s /etc/nginx/sites-available/<conf_name>.conf /etc/nginx/sites-enabled/<conf_name>.conf

sudo nginx -t
sudo systemctl stop nginx
sudo systemctl start nginx
```

### 설정 파일 예시

`site-available/<conf_name>.conf`

동일 sub-domain에서 443(https default), 8080 포트를 각각 listen

> full chain -> 도메인 인증서 + 체인 인증서 + 루트 인증서

```conf
# Default server configuration

server {
    root /var/www/html;

    # Add index.php to the list if you are using PHP
    index index.html index.htm index.nginx-debian.html;

    server_name aws.pravusid.kr;

    location / {
        # First attempt to serve request as file, then
        # as directory, then fall back to displaying a 404.
        # try_files $uri $uri/ =404;
        try_files $uri $uri/ /index.html;
    }

    port_in_redirect off;

    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/aws.pravusid.kr/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/aws.pravusid.kr/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
}

server {
    server_name aws.pravusid.kr;

    port_in_redirect off;

    location / {
        proxy_set_header X-Forwarded-Host $host:$server_port;
        proxy_set_header X-Forwarded-Server $host;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://localhost:8000;
        proxy_redirect http://localhost:8000 https://$host:$server_port;
    }

    listen [::]:8080 ssl ipv6only=on; # managed by Certbot
    listen 8080 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/aws.pravusid.kr/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/aws.pravusid.kr/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
}

server {
    if ($host = aws.pravusid.kr) {
        return 301 https://$host$request_uri;
    }

    listen 80 default_server;
    listen [::]:80 default_server;

    server_name aws.pravusid.kr;
    return 404;
}
```

### gzip compression

<http://nginx.org/en/docs/http/ngx_http_gzip_module.html>

```conf
http {
  gzip on;
  gzip_min_length 1000;
  gzip_proxied expired no-cache no-store private auth;
  gzip_types text/plain application/json application/xml;
  gzip_disable "msie6";
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

> 설정에 사용할 수 있는 변수 참고: <http://nginx.org/en/docs/http/ngx_http_core_module.html#var_arg_>

## URL matching

```conf
server {
    listen 80;
    server_name *.example.com;

    location optional_modifier location_match {
        # ...
    }
}
```

### location modifier

- `none`: prefix match
  - `/site` 경로에 대해 `/site/page1/index.html` 혹은 `/site/index.html` 등이 대응될 수 있다

- `=`: 정확히 일치하는 경로

- `~`: 대소문자를 구분하는 regular expression match

- `~*`: 대소문자를 구분하지 않는 regular expression match

- `^~`: prefix match에서 가장 일치하는 결과인 경우 regular expression match를 하지 않고 연결함
  - `^~ /costumes`인 경우 `/costumes/ninja.html` 경로는 정규식 매칭을 하지 않고 바로 연결함

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

- SSL을 사용하는 경우
  - 헤더설정: <https://www.nginx.com/resources/wiki/start/topics/examples/likeapache/>
  - 프록시 리다이렉트: <http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_redirect>

```text
location / {
    proxy_set_header X-Forwarded-Host $host:$server_port;
    proxy_set_header X-Forwarded-Server $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://localhost:8000;
    proxy_redirect http://localhost:8000 https://$host:$server_port;
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

### 정규표현식 매칭

URI를 정규표현식 매칭으로 분류했으면, 리버스 프록시에서도 capturing 그룹을 넘겨줄 수 있다.

`location ~ ^/api/(admin|setting)/(.*)`

위와 같은 정규표현식 매칭이 있다면 리버스프록시 호출에서 캡쳐 그룹을 `$순서`로 사용할 수 있다.
또한, `$is_args`, `$args`로 query-string을 받을 수 있다.

`proxy_pass http://localhost:8080/$1/$2$is_args$args;`

위와 같은 방법 대신 `rewrite`를 사용할 수도 있다. 이 경우 query-string은 nginx가 처리한다.

```conf
location /api/ {
    rewrite ^/api/(admin|setting)/(.*) /$1/$2 break;
    proxy_pass http://localhost:8080;
}
```

## IP 제한

위 -> 아래 순서대로 작동함

```conf
location ^~ /admin/ {
    deny 1.1.1.1; # Deny a single IP
    deny 2.2.2.0/24; #Deny a IP range
    allow 3.3.3.3;
    allow 1.1.1.0/24;
    deny all; # Deny everyone else
}
```

### Load Balancer (ELB...)등으로 인해 nginx에 ip가 전달되지 않는 경우

```conf
http {
    # ...
    real_ip_header X-Forwarded-For;
    set_real_ip_from 10.0.0.0/8; # <- subnet IPs or Elastic Load Balance IP
    # ...
}
```
