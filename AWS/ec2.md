# Elastic Compute Cloud (EC2)

## 인스턴스 생성

인스턴스 생성시 유의점: 보안그룹 확인

- ssh 인바운드는 지정된 IP에서만
- 웹서비스용 80포트 443포트는 열어야함
- 인스턴스에 접근하기 위한 pem키를 잃어버리면 인스턴스를 다시 생성해야 함

ssh에서 EC2 인스턴스 접속을 위한 호스트설정을 한다: `~/.ssh/config`

```text
Host ec2-idpravus
    HostName <EIP||퍼블릭DNS>
    User <ec2-user||ubuntu>
    IdentityFile ~/.ssh/<pem파일>
```

기본 타임존은 UTC이므로 UTC+9로 변경한다

```shell
sudo rm /etc/localtime
sudo ln -s /usr/share/zoneinfo/Asia/Seoul /etc/localtime
```

## aws-cli로 인스턴스 public DNS 조회

`aws configure` 설정한 이후 사용

`aws ec2 describe-instances --instance-ids <인스턴스id> --query 'Reservations[].Instances[].PublicDnsName'`

## EBS 볼륨크기 변경

<https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/UserGuide/recognize-expanded-volume-linux.html>

## Windows 인스턴스 관리자 암호

<https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/WindowsGuide/ResettingAdminPassword_EC2Config.html>

윈도우 인스턴스내에서 암호를 변경 한 경우 `연결 > 암호 가져오기`가 작동하지 않는다

> 암호변경: `net user Administrator "new_password"`

### Windows 2016 이전 버전의 경우

이 경우 `C:\Program Files\Amazon\Ec2ConfigService\Settings\config.xml` 파일에서
`Ec2SetPassword` 항목을 `Enabled`로 변경하고 인스턴스를 종료후 다시 켠다.

부팅 후 4분 이상 대기한 뒤 `연결 > 암호 가져오기`를 다시 실행한다.

직접 암호를 변경했을 때와 마찬가지로 암호가 변경으로 인해 `Ec2SetPassword` 항목이 `Disabled`로 변경되어 있다.

<https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/WindowsGuide/ami-create-standard.html>

만약 해당 인스턴스의 이미지를 생성한다면 암호를 재생성 할 수 있도록
`Ec2SetPassword` 항목을 `Enabled`로 처리하는 방법도 좋을 듯 하다.

### Windows 2016 이상 버전

암호를 변경하려면 볼륨을 분리하여 다른 인스턴스에서 접근하는 방법으로 가능하다

<https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/WindowsGuide/ResettingAdminPassword_EC2Launch.html>

관리자 암호를 Random으로 지정하고, `EC2Launch > Shutdown with Sysprep` 실행후 이미지를 생성하면 암호를 재생성 할 수 있음

<https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/WindowsGuide/ec2launch.html#ec2launch-sysprep>

## 버스트 가능 인스턴스 (T)

<https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/UserGuide/burstable-performance-instances.html>

> 기존 Amazon EC2 인스턴스 유형은 고정된 CPU 리소스를 제공하는 반면,
> 성능 순간 확장 가능 인스턴스는 기본 수준의 CPU 사용률을 제공하면서 기본 수준 이상으로 CPU 사용률을 버스트하는 기능을 제공합니다.
> 이렇게 하면 기준 CPU와 추가 버스트 CPU 사용량에 대해서만 비용을 지불하면 되므로 컴퓨팅 비용이 절감됩니다.
> 기준 사용률과 버스트 기능은 CPU 크레딧에 의해 좌우됩니다.

### 핵심개념

<https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/UserGuide/burstable-credits-baseline-concepts.html>

- 시간당 적립되는 크레딧 수 = [기준 사용률(%)](https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/UserGuide/burstable-credits-baseline-concepts.html#burstable-performance-instances-credit-table)] x vCPU 수 x 60분
- 분당 소비되는 CPU 크레딧 = vCPU 수 x CPU 사용률 x 1분
- 적립 - 소비 크레딧은 누적 한도(일반적으로 24시간 동안 적립되는 최대 크레딧 수)까지 적립됨

### 무제한 모드 (unlimited)

- <https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode.html>
- <https://blog.wisen.co.kr/pages/blog/blog-detail.html?idx=2726>

> unlimited로 구성된 성능 순간 확장 가능 인스턴스의 CPU 크레딧 밸런스가 감소하면 잉여 크레딧을 사용하여 기준 이상으로 버스트할 수 있습니다.
> CPU 사용률이 기준 미만으로 떨어지면 획득한 CPU 크레딧을 사용하여 이전에 소비한 잉여 크레딧을 청산할 수 있습니다.
> CPU 크레딧을 획득하고 잉여 크레딧을 청산하는 기능을 통해 Amazon EC2은 24시간 동안 인스턴스의 CPU 사용률을 평균 수준으로 유지할 수 있습니다.
> 24시간 동안의 평균 CPU 사용량이 기준을 초과하는 경우 인스턴스에 추가 사용량에 대해 vCPU 시간당 고정 추가 요금이 청구됩니다.

- 무제한 모드에서는 잔여 크레딧이 0이 되면, 기준 사용률 이하로 제한되지 않고 잉여(surplus) 크레딧을 적립하게 된다
- 잉여 크레딧의 한도는 크레딧 누적 한도와 동일하다
- 잉여 크레딧 누적 한도를 초과하여 적립하게 되면 초과분에 대해 과금이 이루어진다

> [무제한 모드 vs 고정 CPU 손익분기](https://docs.aws.amazon.com/ko_kr/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#when-to-use-unlimited-mode)
