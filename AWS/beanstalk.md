# AWS Beanstalk

<https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/Welcome.html>

## 개념

- Application

  - SourceCode -> Application Version (N)
  - AWS resources -> Environment (M)

> 동시에 M개의 배포 가능 (배포 총 경우의 수 == N * M)

## 리눅스 플랫폼 확장

<https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/platforms-linux-extend.html>

### `.platform/nginx`

> [[nginx]] 설정 참고

[기본설정 값 확인](https://stackoverflow.com/questions/66074722/what-is-the-port-number-of-the-web-application-to-which-default-proxy-config-on) (AMZ linux 2 기준)

`cat /etc/nginx/conf.d/elasticbeanstalk/00_application.conf`

```conf
location / {
    proxy_pass          http://127.0.0.1:8080;
    proxy_http_version  1.1;

    proxy_set_header    Connection          $connection_upgrade;
    proxy_set_header    Upgrade             $http_upgrade;
    proxy_set_header    Host                $host;
    proxy_set_header    X-Real-IP           $remote_addr;
    proxy_set_header    X-Forwarded-For     $proxy_add_x_forwarded_for;
}
```
