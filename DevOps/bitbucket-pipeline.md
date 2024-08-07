# Bitbucket Pipeline

## Docs

- <https://support.atlassian.com/bitbucket-cloud/docs/build-test-and-deploy-with-pipelines/>
- <https://support.atlassian.com/bitbucket-cloud/docs/javascript-nodejs-with-bitbucket-pipelines/>

### SSH 사용

- <https://support.atlassian.com/bitbucket-cloud/docs/variables-and-secrets/>
- <https://support.atlassian.com/bitbucket-cloud/docs/deploy-using-scp/>

## pipeline 무시

커밋 메시지에 `[skip ci]` 또는 `[ci skip]` 포함

## Cache

<https://support.atlassian.com/bitbucket-cloud/docs/cache-dependencies/>

### npm cache

- `npm ci` 명령어의 경우 실행 직후 `node_modules` 디렉토리를 삭제하므로 캐시를 사용하기 위해서는 커스텀 캐시를 선언해야 한다
- <https://community.developer.atlassian.com/t/caching-and-installing-node-dependencies-in-pipeline/35659>

```yml
pipelines:
  default:
    - step:
        name: Install deps
        caches:
          - npm
        script:
          - npm ci

definitions:
  caches:
    npm:
      key:
        files:
          - package-lock.json
          - '**/package-lock.json'
      path: $HOME/.npm
```

## Example

### atlassian pipes

<https://bitbucket.org/atlassian/workspace/projects/BPP>

### deployment

<https://support.atlassian.com/bitbucket-cloud/docs/set-up-and-monitor-deployments/>

> We support deploying to test, staging, and production type environments, and they must be listed in this order in each pipeline.

### PR 빌드 & 테스트

```yml
image: node:22

pipelines:
  pull-requests:
    '**':
      - step:
          name: Install deps
          script:
            - npm ci
          artifacts:
            - 'node_modules/**'
      - parallel:
          - step:
              name: Typecheck
              script:
                - npm run typecheck
          - step:
              name: Lint
              script:
                - npm run lint:quiet
          - step:
              name: Test
              script:
                - npm run test
```

### S3 배포 후 cloudfront invalidation

```yml
image: node:22

pipelines:
  branches:
    master:
      - step:
          name: Deploy
          deployment: production
          caches:
            - npm
          script:
            - npm install
            - npm run build
            - pipe: atlassian/aws-s3-deploy:1.1.0
              variables:
                AWS_ACCESS_KEY_ID: $AWS_ACCESS_KEY_ID
                AWS_SECRET_ACCESS_KEY: $AWS_SECRET_ACCESS_KEY
                AWS_DEFAULT_REGION: $AWS_DEFAULT_REGION
                S3_BUCKET: $S3_BUCKET
                LOCAL_PATH: './dist'
                DELETE_FLAG: 'true'
                # ACL: 'public-read'
                # CONTENT_ENCODING: '<string>' # Optional.
                # STORAGE_CLASS: '<string>' # Optional.
                # CACHE_CONTROL: '<string>' # Optional.
                # EXPIRES: '<timestamp>' # Optional.
                # EXTRA_ARGS: '<string>' # Optional.
                # DEBUG: '<boolean>' # Optional.
            - pipe: atlassian/aws-cloudfront-invalidate:0.10.0
              variables:
                AWS_ACCESS_KEY_ID: $AWS_ACCESS_KEY_ID
                AWS_SECRET_ACCESS_KEY: $AWS_SECRET_ACCESS_KEY
                AWS_DEFAULT_REGION: $AWS_DEFAULT_REGION
                DISTRIBUTION_ID: $DISTRIBUTION_ID
                PATHS: '/*'
                # DEBUG: '<boolean>' # Optional
            - pipe: atlassian/slack-notify:2.3.0
              variables:
                WEBHOOK_URL: $WEBHOOK_URL
                MESSAGE: 'SLACK MESSAGE!'
```

### codedeploy to ec2

<https://support.atlassian.com/bitbucket-cloud/docs/deploy-to-aws-with-codedeploy/>

```yml
image: node:22

pipelines:
  branches:
    master:
      - step:
          name: Build & Compress
          caches:
            - npm
          script:
            - apt-get update && apt-get install -y zip
            - npm ci --only=prod
            - npm run build
            - zip -q -y -r ${APPLICATION_NAME}.dist.zip ./ -x '.git/*' -x '.vscode/*'
          artifacts:
            - '*.dist.zip'

      - step:
          name: Upload & Deploy with CodeDeploy
          deployment: production
          script:
            - pipe: atlassian/aws-code-deploy:1.5.0
              variables:
                AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
                AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}
                AWS_DEFAULT_REGION: ${AWS_DEFAULT_REGION}
                COMMAND: 'upload'
                APPLICATION_NAME: ${APPLICATION_NAME}
                S3_BUCKET: '${AWS_DEPLOY_BUCKET}'
                ZIP_FILE: '${APPLICATION_NAME}.dist.zip'
            - pipe: atlassian/aws-code-deploy:1.5.0
              variables:
                AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
                AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}
                AWS_DEFAULT_REGION: ${AWS_DEFAULT_REGION}
                COMMAND: 'deploy'
                APPLICATION_NAME: ${APPLICATION_NAME}
                DEPLOYMENT_GROUP: ${DEPLOYMENT_GROUP}
                S3_BUCKET: '${AWS_DEPLOY_BUCKET}'
                IGNORE_APPLICATION_STOP_FAILURES: 'false'
                # FILE_EXISTS_BEHAVIOR: 'OVERWRITE'
                WAIT: 'true'
```
