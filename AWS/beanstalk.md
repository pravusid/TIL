# AWS Beanstalk

<https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/Welcome.html>

## 개념

- Application

  - SourceCode -> Application Version (N)
  - AWS resources -> Environment (M)

> 동시에 M개의 배포 가능 (배포 총 경우의 수 == N * M)

## 환경구성

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

`.platform/nginx/conf.d/timeout.conf`: 타임아웃 설정

```conf
send_timeout 90s;
proxy_connect_timeout 90s;
proxy_send_timeout 90s;
proxy_read_timeout 90s;
```

### `.ebextensions`

- <https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/ebextensions.html>
- <https://docs.aws.amazon.com/ko_kr/elasticbeanstalk/latest/dg/command-options-general.html>

### ALBv2

생성 환경의 로드밸런서 설정 변경

`.ebextensions/alb.config`

```yml
option_settings:
  aws:elbv2:loadbalancer:
    IdleTimeout: 90
    ManagedSecurityGroup: <SG_ID>
    SecurityGroups: <SG_ID>,<OTHER_SG_ID>
```

## 참고

### timeout

- <https://stackoverflow.com/questions/72130038/how-to-solve-aws-elastic-beanstalk-504-timeout-error>
- <https://stackoverflow.com/questions/68974919/elasticbeanstalk-returns-504-gateway-timeout>

### security group (ssh)

> Remove default security group from EC2-Instance
>
> -- <https://github.com/aws/elastic-beanstalk-roadmap/issues/44>
