# Let's Encrypt

<https://letsencrypt.org/docs/>

Let’s Encrypt is a free, automated, and open Certificate Authority

## 발급: NGINX + certbot stand alone

`sudo certbot certonly --webroot -w /var/www/example -d example.com -d www.example.com`

인증서 생성 도중 대상 도메인에 대한 소유권 확인 과정을 거친다

`http://example.com/.well-known/acme-challenge/<RANDOM_STRING>` 경로로 접속할 때 값이 출력되어야 함

`/etc/nginx/conf.d/letsencrypt.conf` 파일에 다음을 임시로 추가한다

```conf
location ^~ /.well-known/acme-challenge/ {
    allow all;
    root /var/lib/letsencrypt/;
    default_type "text/plain";
    try_files $uri =404;
}
```

`/etc/nginx/servers-available/domain.conf`

```conf
server {
    server_name domain.tld
    # ..
    include conf.d/letsencrypt.conf;
}
```

임시파일을 생성할 디렉토리를 만든다

```sh
mkdir -p /var/lib/letsencrypt/.well-known
chgrp http /var/lib/letsencrypt
chmod g+s /var/lib/letsencrypt
```

random key 처리를 삭제하고 sites-available 에서 다음을 추가한다

```conf
server {
    server_name www.mysite.com;
    listen 443 ssl http2 default_server;
    listen [::]:443 ssl http2 default_server ipv6only=on;

    ssl_certificate /etc/letsencrypt/live/www.mysite.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/www.mysite.com/privkey.pem;
    ssl_trusted_certificate /etc/letsencrypt/live/www.mysite.com/fullchain.pem;
}
```

## 발급: NGINX + certbot auto

<https://certbot.eff.org/>

Let's Encrypt 인증서는 3개월의 유효기간을 갖고 있으며, 자동 발급/갱신을 도와주는 certbot을 이용함

```sh
sudo apt-get install software-properties-common
sudo add-apt-repository universe
sudo add-apt-repository ppa:certbot/certbot
sudo apt-get install python-certbot-nginx
```

- `/etc/nginx/sites-available/default`
- `/etc/nginx/sites-available/example.com`

두 파일 중 하나를 선택해(`example.com` 이름도 변경) `server_name`을 도메인 이름으로 변경

```conf
# http 요청을 https로 리다이렉트
server {
    server_name www.example.com example.com;

    return 301 https://example.com$request_uri;
}

# www 서브도메인을 리다이렉트
server {
    listen 443 ssl http2;

    server_name www.example.com;
    include snippets/ssl-params.conf;

    return 301 https://example.com$request_uri;
}

server {
    listen 443 ssl http2;

    server_name example.com;
    include snippets/ssl-params.conf;

    root /var/www/example.com;
    index index.html;

    location / {
        try_files $uri $uri/ =404;
    }
}
```

만약 다른 포트(:8080)가 리다이렉트로 인해 인증서 오류가 나면, 포트는 리다이렉트 제외 설정할 수 있음

```conf
server {
    port_in_redirect off;
}
```

적용 (defalut 파일은 기본적으로 sites-enabled에 심볼릭 링크가 있음)

```sh
sudo ln -s /etc/nginx/sites-available/example.com /etc/nginx/sites-enabled/
sudo nginx -t
sudo service nginx restart
```

인증서 발급: `sudo certbot --nginx -d example.com`

## 도메인 변경

`certbot certonly --cert-name <인증서 이름> -d <도메인1>,<도메인2>`

## 갱신

갱신 가능여부 확인: `sudo certbot renew --dry-run`

Ubuntu의 경우 `/etc/cron.d/`에 certbot이 생성되어있음

cron job을 등록한다

```shell
crontab -e
0 19 * * * certbot renew --post-hook "service nginx reload"
```

등록된 job을 확인한다: `crontab -l`

## 제거

인증제거: `sudo certbot delete --cert-name example.com`

인증서 revoke: <https://certbot.eff.org/docs/using.html#revoking-certificates>
