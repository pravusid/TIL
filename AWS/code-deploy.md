# AWS CodeDeploy

## CodeDeploy Agent

- CodeDeploy CLI 설치
  - `sudo yum install ruby`
  - `curl -O https://aws-codedeploy-ap-northeast-2.s3.amazonaws.com/latest/install`
  - `chmod +x install`
  - `sudo ./install auto`

- CodeDeploy Agent 실행
  - `sudo service codedeploy-agent start`
  - `sudo service codedeploy-agent status`

- CodeDeploy Agent 사용자 변경: <https://aws.amazon.com/ko/premiumsupport/knowledge-center/codedeploy-agent-non-root-profile/>

### agent 실행 환경변수

agent 실행시 환경변수는 다음과 같음

1. LIFECYCLE_EVENT : This variable contains the name of the lifecycle event associated with the script.
2. DEPLOYMENT_ID :  This variables contains the deployment ID of the current deployment.
3. APPLICATION_NAME :  This variable contains the name of the application being deployed. This is the name the user sets in the console or AWS CLI.
4. DEPLOYMENT_GROUP_NAME :  This variable contains the name of the deployment group. A deployment group is a set of instances associated with an application that you target for a deployment.
5. DEPLOYMENT_GROUP_ID : This variable contains the ID of the deployment group in AWS CodeDeploy that corresponds to the current deployment

## 준비

- 배포할 데이터를 업로드할 **S3 bucket 생성**
- 외부 트리거(CI...)에서 사용할 **[IAM 사용자를 추가](https://docs.aws.amazon.com/ko_kr/codedeploy/latest/userguide/getting-started-provision-user.html)**
  - S3 접근권한 (`action:s3:PutObject`)
  - CodeDeploy 권한 (`POLICY:AWSCodeDeployDeployerAccess`)
- CodeDeploy가 배포를 진행하기 위한 **[CodeDeploy Role 생성](https://docs.aws.amazon.com/ko_kr/codedeploy/latest/userguide/getting-started-create-service-role.html)**
- EC2에서 작업을 수행하기 위해 **[S3 접근 권한이 있는 Role 생성](https://docs.aws.amazon.com/ko_kr/codedeploy/latest/userguide/getting-started-create-iam-instance-profile.html)** 후 EC2에 Role 할당
- CodeDeploy 웹 콘솔에서 **[배포 그룹 생성(ARN에서 기존에 생성한 CodeDeploy Role 선택)](https://docs.aws.amazon.com/ko_kr/codedeploy/latest/userguide/deployment-groups-create.html)** 후 배포 애플리케이션 생성

## `appspec.yml`

<https://docs.aws.amazon.com/ko_kr/codedeploy/latest/userguide/reference-appspec-file.html>

최소 기능 스크립트 (`runas` 사용시 사용자 암호 입력이 필요하다면 오류발생함)

```yml
version: 0.0
os: linux
files:
  - source: ./
    destination: /home/ec2-user/deployment/
hooks:
  BeforeInstall:
    - location: scripts/deploy-prepare.sh
      runas: ec2-user
 ApplicationStart:
   - location: scripts/deploy-run.sh
     runas: ec2-user
```

- 스크립트가 실행되는 경로는 Code Deploy Agent가 설치된 경로임 `/opt/codedeploy-agent`
- 실제 파일 경로는: `/opt/codedeploy-agent/deployment-root/{deployment-group-ID}/{deployment-ID}`
