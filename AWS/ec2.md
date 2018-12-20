# Elastic Compute Cloud (EC2)

## aws-cli로 인스턴스 public DNS 조회

`aws configure` 설정한 이후 사용

`aws ec2 describe-instances --instance-ids <인스턴스id> --query 'Reservations[].Instances[].PublicDnsName'`
