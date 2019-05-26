# Apache

## 설정 예시

`httpd.conf` 파일에서 `Include /private/etc/apache2/extra/httpd-vhosts.conf` 주석 해제

다음 명령어로 ssl 모듈 활성화

```sh
sudo a2enmod ssl
sudo a2enmod rewrite
```

`extra/httpd-vhosts.conf`

```xml
<VirtualHost *:80>
    ServerName pravusid.kr
    Redirect / https://pravusid.kr
</VirtualHost>

<VirtualHost *:443>
    ServerName pravusid.kr

    SSLEngine on
    SSLProtocol all
    SSLCertificateKeyFile /etc/letsencrypt/live/pravusid/privkey.pem
    SSLCertificateFile /etc/letsencrypt/live/pravusid/cert.pem
    SSLCertificateChainFile /etc/letsencrypt/live/pravusid/chain.pem

    ErrorLog "logs/pravusid-error.log"
    ProxyPreserveHost On

    ProxyPass /.well-known/acme-challenge/ !

    ProxyPass /bitbucket-hook/ http://localhost:8080/bitbucket-hook/
    ProxyPassReverse /bitbucket-hook/ http://localhost:8080/bitbucket-hook/

    ProxyPass / http://localhost:3000/
    ProxyPassReverse / http://localhost:3000/

    Alias /.well-known/acme-challenge/ "/var/www/certbot/.well-known/acme-challenge/"
    <Directory "/var/www/certbot">
        Options None
        AllowOverride None
        ForceType text/plain
        RedirectMatch 404 "^(?!/\.well-known/acme-challenge/[\w-]{43}$)"
    </Directory>

    <Location "/">
        Order deny,allow
        Deny from all
        Allow from 0.0.0.0
    </Location>

    <Location "/.well-known/acme-challenge/">
        Allow from All
    </Location>
</VirtualHost>
```
