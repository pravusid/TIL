# Let's Encrypt

<https://letsencrypt.org/docs/>

Let’s Encrypt is a free, automated, and open Certificate Authority

## 발급: NGINX + certbot

<https://certbot.eff.org/>

Let's Encrypt 인증서는 3개월의 유효기간을 갖고 있으며, 자동 발급/갱신을 도와주는 certbot을 이용함

```shell
sudo apt-get install software-properties-common
sudo add-apt-repository universe
sudo add-apt-repository ppa:certbot/certbot
sudo apt-get install python-certbot-nginx
```

- `/etc/nginx/sites-available/default`
- `/etc/nginx/sites-available/example.com`

두 파일 중 하나의 `server_name`을 도메인 이름으로 변경

```text
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

```text
server {
  port_in_redirect off;
}
```

적용 (defalut 파일은 기본적으로 sites-enabled에 심볼릭 링크가 있음)

```shell
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
