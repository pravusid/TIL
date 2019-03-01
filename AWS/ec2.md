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
