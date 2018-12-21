# AWS CodeDeploy

## 과정

- 외부 트리거(CI...)에서 사용할 S3 접근권한과 CodeDeploy 권한이 있는 IAM 사용자를 추가
- S3 버킷을 생성함
- EC2에서 작업을 수행하기 위한 CodeDeploy 역할을 생성 후 EC2에 역할 할당
- CodeDeploy가 배포를 진행하기 위한 역할을 생성
- EC2의 aws-cli에서 생성한 사용자로 인증함
- CodeDeploy CLI 설치
  - `aws s3 cp s3://aws-codedeploy-ap-northeast-2/latest/install . --region ap-northeast-2`
  - `sudo ./install auto`
- CodeDeploy Agent 자동실행: `systemctl enable codedeploy-agent`
- `appspec.yml` 파일 작성
- CodeDeploy 웹콘솔에서 배포 애플리케이션을 생성함(ARN에서 기존에 생성한 CodeDeploy 역할 선택)
