# Let's Encrypt & CertBot

<https://letsencrypt.org/docs/>

Let’s Encrypt is a free, automated, and open Certificate Authority

Let's Encrypt 인증서는 3개월의 유효기간을 갖고 있으며, 자동 발급/갱신을 도와주는 certbot을 이용함

```sh
sudo apt-get install software-properties-common
sudo add-apt-repository universe
sudo add-apt-repository ppa:certbot/certbot
sudo apt-get install python-certbot-nginx
```

## 인증서

아래의 발급 과정을 거치면 다음 인증서가 생성됨

`/etc/letsencrypt/archive` -> (symbolic link) `/etc/letsencrypt/live/<domain_name>`

- `fullchain.pem` : cert.pem + chain.pem
- `cert.pem` : 도메인 인증서
- `chain.pem` : Let’s Encrypt chain 인증서
- `privkey.pem` : 인증서 개인키

## 발급

## 발급: StandAlone

```sh
sudo certbot certonly --authenticator standalone \
    -d example.com -d www.example.com \
    --pre-hook "systemctl stop apache2" --post-hook "systemctl start apache2"
    --pre-hook "systemctl stop nginx" --post-hook "systemctl start nginx"
```

### 발급: Manual

DNS의 TXT record로 인증 후 발급

```sh
certbot certonly --manual --preferred-challenges dns -d pravusid.kr -d www.pravusid.kr

# Are you OK with your IP being logged?:
# -> Yes

# Please deploy a DNS TXT record under the name
# -> 출력되는 random string을 _acme-challenge.pravusid.kr TXT record에 등록함
# -> 등록후 $ nslookup -q=TXT _acme-challenge.pravusid.kr 입력하여 적용 확인
```

### 발급: NGINX + certbot webroot

Challenge Seed를 외부에서 접근 가능한 경로에(`/.well-known`)에 위치시켜 인증받는다

`sudo certbot certonly --cert-name <인증서이름> --webroot -w /var/www/certbot -d example.com -d www.example.com`

인증서 생성 도중 대상 도메인에 대한 소유권 확인 과정을 거친다

`http://example.com/.well-known/acme-challenge/<RANDOM_STRING>` 경로로 접속할 때 값이 출력되어야 함

WebRoot로 사용할 디렉토리를 생성한다

```sh
mkdir /var/www/certbot
chown nginx:nginx /var/www/certbot
chmod 700 /var/www/certbot
```

`/etc/nginx/conf.d/letsencrypt.conf` 파일에서 well-known 디렉토리와 WebRoot를 연결한다

```conf
location /.well-known/acme-challenge/ {
    root /var/www/certbot/;
}

# 또는

location ^~ /.well-known/acme-challenge/ {
    default_type "text/plain";
    root /home/www/certbot;
}
```

`/etc/nginx/servers-available/domain.conf`

```conf
server {
    server_name domain.tld
    # ...
    include conf.d/letsencrypt.conf;
}
```

`sites-available`에서 다음을 추가한다

```conf
server {
    server_name www.mysite.com;
    listen 443 ssl http2 default_server;
    listen [::]:443 ssl http2 default_server ipv6only=on;

    ssl_certificate /etc/letsencrypt/live/www.mysite.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/www.mysite.com/privkey.pem;
    ssl_trusted_certificate /etc/letsencrypt/live/www.mysite.com/chain.pem;
}
```

### 발급 Apache + certbot webroot

아파치 설정파일

```conf
Alias /.well-known/acme-challenge/ "/var/www/certbot/.well-known/acme-challenge/"
<Directory "/var/www/certbot/">
    Options None
    AllowOverride None
    ForceType text/plain
    RedirectMatch 404 "^(?!/\.well-known/acme-challenge/[\w-]{43}$)"
</Directory>
```

인증서 생성

`sudo certbot certonly --cert-name <인증서이름> --webroot -w /var/www/certbot -d example.com -d www.example.com`

리버스 프록시 사용시 다음 추가

`ProxyPass /.wellknown !`

### 발급: NGINX + certbot auto

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

인증서 발급: `sudo certbot --nginx -d example.com -d www.example.com`

## WildCard 사용

도메인에 WildCard 사용을 위해서는 TXT record를 통해서만 가능한데,
이를 자동으로 수행하기 위해 DNS 플러그인을 사용할 수 있다.

<https://certbot.eff.org/docs/using.html#dns-plugins>

## 도메인 변경

`certbot certonly --cert-name <인증서 이름> -d <도메인1>,<도메인2>`

## 갱신

갱신 가능여부 확인: `sudo certbot renew --dry-run`

certbot-auto의 경우: `--no-self-upgrade`를 추가하여 certbot의 업그레이드를 막음

Ubuntu의 경우 `/etc/cron.d/`에 certbot이 생성되어있음: <https://certbot.eff.org/docs/using.html#automated-renewals>

> Many Linux distributions provide automated renewal when you use the packages installed through their system package manager.

cron job을 등록한다

```shell
crontab -e
0 19 * * * certbot renew --quiet --post-hook "service nginx reload"
```

등록된 job을 확인한다: `crontab -l`

## 제거

인증제거: `sudo certbot delete --cert-name <인증서이름>`

인증제거(대화형): `sudo certbot delete`

인증서 revoke: <https://certbot.eff.org/docs/using.html#revoking-certificates>
