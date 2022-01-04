# Bitbucket Pipeline

## Docs

- <https://support.atlassian.com/bitbucket-cloud/docs/build-test-and-deploy-with-pipelines/>
- <https://support.atlassian.com/bitbucket-cloud/docs/javascript-nodejs-with-bitbucket-pipelines/>

### SSH 사용

- <https://support.atlassian.com/bitbucket-cloud/docs/variables-and-secrets/>
- <https://support.atlassian.com/bitbucket-cloud/docs/deploy-using-scp/>

## pipeline 무시

커밋 메시지에 `[skip ci]` 또는 `[ci skip]` 포함

## Example

### atlassian pipes

<https://bitbucket.org/atlassian/workspace/projects/BPP>

### deployment

<https://support.atlassian.com/bitbucket-cloud/docs/set-up-and-monitor-deployments/>

> We support deploying to test, staging, and production type environments, and they must be listed in this order in each pipeline.

### PR 빌드 & 테스트

```yml
image: node:16.13.0

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
image: node:16.13.0

pipelines:
  branches:
    master:
      - step:
          name: Deploy
          deployment: production
          caches:
            - node
          script:
            - npm ci
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
            - pipe: atlassian/aws-cloudfront-invalidate:0.6.0
              variables:
                AWS_ACCESS_KEY_ID: $AWS_ACCESS_KEY_ID
                AWS_SECRET_ACCESS_KEY: $AWS_SECRET_ACCESS_KEY
                AWS_DEFAULT_REGION: $AWS_DEFAULT_REGION
                DISTRIBUTION_ID: $DISTRIBUTION_ID
                PATHS: '/*'
                # DEBUG: '<boolean>' # Optional
            - pipe: atlassian/slack-notify:2.0.0
              variables:
                WEBHOOK_URL: $WEBHOOK_URL
                MESSAGE: 'SLACK MESSAGE!'
```

### codedeploy to ec2

<https://support.atlassian.com/bitbucket-cloud/docs/deploy-to-aws-with-codedeploy/>

```yml
image: atlassian/default-image:3

pipelines:
  branches:
    master:
      - step:
          image: node:16.13.0
          name: Build & Compress
          caches:
            - node
          script:
            - apt-get update && apt-get install -y zip
            - npm ci --only=prod && npm run build
            - zip -r dist.zip ./
          artifacts:
            - dist.zip

      - step:
          name: Upload & Deploy with CodeDeploy
          deployment: production
          services:
            - docker
          script:
            - pipe: atlassian/aws-code-deploy:1.1.1
              variables:
                AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
                AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}
                AWS_DEFAULT_REGION: ${AWS_DEFAULT_REGION}
                COMMAND: 'upload'
                APPLICATION_NAME: ${APPLICATION_NAME}
                S3_BUCKET: '${AWS_DEPLOY_BUCKET}'
                ZIP_FILE: 'dist.zip'
            - pipe: atlassian/aws-code-deploy:1.1.1
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
