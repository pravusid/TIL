# AWS Beanstalk

<https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/Welcome.html>

## 개념

- Application

  - SourceCode -> Application Version (N)
  - AWS resources -> Environment (M)

> 동시에 M개의 배포 가능 (배포 총 경우의 수 == N x M)

## 환경 구성

### 리눅스 플랫폼 확장

<https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/platforms-linux-extend.html>

#### `.platform/hooks/prebuild`

> 인스턴스에 스왑 할당
>
> -- <https://stackoverflow.com/questions/36119306/can-i-configure-linux-swap-space-on-aws-elastic-beanstalk>

```bash
#!/usr/bin/env bash

set -o xtrace
set -e

if grep -E 'SwapTotal:\s+0+\s+kB' /proc/meminfo; then
  echo "Positively identified no swap space, creating some."
  sudo dd if=/dev/zero of=/var/swapfile bs=1M count=512
  sudo /sbin/mkswap /var/swapfile
  sudo chmod 000 /var/swapfile
  sudo /sbin/swapon /var/swapfile
else
  echo "Did not confirm zero swap space, doing nothing."
fi
```

#### `.platform/nginx`

> [[nginx]] 설정 참고

[Beanstalk Nginx 기본설정 값 확인](https://stackoverflow.com/questions/66074722/what-is-the-port-number-of-the-web-application-to-which-default-proxy-config-on) (AMZ linux 2 기준)

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

`.platform/nginx/conf.d/timeout.conf`: [[nginx#Timeout|타임아웃 설정]]

```conf
send_timeout 90s;
proxy_connect_timeout 90s;
proxy_send_timeout 90s;
proxy_read_timeout 90s;
```

인스턴스 nginx 기본 설정은 다음과 같다 (상단의 설정은 Include를 통해 적용된다)

`/etc/nginx/nginx.conf`

```conf
user                    nginx;
error_log               /var/log/nginx/error.log warn;
pid                     /var/run/nginx.pid;
worker_processes        auto;
worker_rlimit_nofile    131591;

events {
    worker_connections  1024;
}

http {
    server_tokens off;

    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    include       conf.d/*.conf;

    map $http_upgrade $connection_upgrade {
        default     "upgrade";
    }

    server {
        listen        80 default_server;
        access_log    /var/log/nginx/access.log main;

        client_header_timeout 60;
        client_body_timeout   60;
        keepalive_timeout     60;
        gzip                  off;
        gzip_comp_level       4;
        gzip_types text/plain text/css application/json application/javascript application/x-javascript text/xml application/xml application/xml+rss text/javascript;

        # Include the Elastic Beanstalk generated locations
        include conf.d/elasticbeanstalk/*.conf;
    }
}
```

### `.ebextensions`

- <https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/ebextensions.html>
- <https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/command-options-general.html>

#### timezone

`.ebextensions/00-set-timezone.config`

```yml
commands:
  set_time_zone:
    command: ln -f -s /usr/share/zoneinfo/Asia/Seoul /etc/localtime
```

#### ALBv2

생성 환경의 로드밸런서 설정 변경

`.ebextensions/alb.config`

```yml
option_settings:
  aws:elbv2:loadbalancer:
    IdleTimeout: 90
    ManagedSecurityGroup: <SG_ID>
    SecurityGroups: <SG_ID>,<OTHER_SG_ID>
```

### 구성 옵션

#### 환경 속성

<https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/custom-platforms-scripts.html>

### 저장된 구성

<https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/environment-configuration-savedconfig.html>

`save configuration` 값은 S3에 저장된다: `s3://elasticbeanstalk-{region}-{account}/resources/templates/{application}/`

## 참고

### timeout

- <https://stackoverflow.com/questions/72130038/how-to-solve-aws-elastic-beanstalk-504-timeout-error>
- <https://stackoverflow.com/questions/68974919/elasticbeanstalk-returns-504-gateway-timeout>

### security group (ssh)

> Remove default security group from EC2-Instance
>
> -- <https://github.com/aws/elastic-beanstalk-roadmap/issues/44>

EC2 key pair 등록한 경우 보안그룹에 22번 포트 ingress 생성되는 것으로 보임

> 등록한 key pair 삭제 (웹 콘솔에는 기능 없음)
>
> `aws elasticbeanstalk update-environment --environment-name $ENV --options-to-remove 'Namespace=aws:autoscaling:launchconfiguration,OptionName=EC2KeyName'`
>
> -- <https://github.com/aws/elastic-beanstalk-roadmap/issues/78>

### EIP

[Remove public IP on single instance deployments](https://github.com/aws/elastic-beanstalk-roadmap/issues/47)
