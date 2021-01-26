# AWS lambda

## Serverless Computing

사용자가 애플리케이션/운영체제/하드웨어를 모두 관리하는 전통적인 방식에서, 하드웨어는 관리하지 않는 가상서버로 넘어갔고,
추가로 Docker와 같은 컨테이너를 쓴다면 운영체제 영역도 관리하지 않아도 된다.

서버리스는 다른 모든 것을 신경쓸 필요 없이 CPU와 메모리만을 할당받아 코드를 실행한다.
이러한 접근법을 Fuction as a Service 라고 부르기도 한다.

## Serverless Framework

<https://serverless.com/framework/docs/>

## lambda

lambda + API gateway: <https://docs.aws.amazon.com/ko_kr/lambda/latest/dg/with-on-demand-https-example.html>

## AWS SAM

- <https://docs.aws.amazon.com/ko_kr/serverless-application-model/latest/developerguide/what-is-sam.html>
- <https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md>

### CLI

<https://docs.aws.amazon.com/ko_kr/serverless-application-model/latest/developerguide/serverless-sam-cli-command-reference.html>

`pip install --user aws-sam-cli`

- scaffolding: `sam init`
- build: `sam build`
- packaging: `sam package --template-file template.yaml --s3-bucket mybucket --output-template-file packaged.yaml`
- deploy
  - `sam deploy --guided`
  - `sam deploy --template-file ./packaged.yaml --stack-name mystack --capabilities CAPABILITY_IAM`
- run local
  - generate event: `sam local generate-event [SERVICE] --help`
  - run lambda endpoint: `sam local start-lambda` (AWS CLI 혹은 SDK로 호출)
  - run api gateway: `sam local start-api`
  - invoke function: `sam local invoke "HelloWorldFunction" -e event.json`
  - env var: `sam local <start-api|invoke> --env-vars env.json`
- cleanup: `aws cloudformation delete-stack --stack-name <name> --region <region>`

## Using AWS Lambda with Amazon SNS

CloudWatch -> SNS -> Lambda 메시지 처리

<https://docs.aws.amazon.com/ko_kr/lambda/latest/dg/with-sns.html>
