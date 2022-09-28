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
