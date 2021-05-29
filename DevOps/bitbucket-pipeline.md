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

### PR 빌드 & 테스트

```yml
image: node:12.14.0

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
image: node:12.14.0

pipelines:
  branches:
    master:
      - step:
          deployment: production
          caches:
            - node
          script:
            - npm ci
            - npm run build
            - pipe: atlassian/aws-s3-deploy:0.3.7
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
            - pipe: atlassian/aws-cloudfront-invalidate:0.1.3
              variables:
                AWS_ACCESS_KEY_ID: $AWS_ACCESS_KEY_ID
                AWS_SECRET_ACCESS_KEY: $AWS_SECRET_ACCESS_KEY
                AWS_DEFAULT_REGION: $AWS_DEFAULT_REGION
                DISTRIBUTION_ID: $DISTRIBUTION_ID
                PATHS: '/*'
                # DEBUG: '<boolean>' # Optional
            - pipe: atlassian/slack-notify:0.3.2
              variables:
                WEBHOOK_URL: $WEBHOOK_URL
                MESSAGE: 'SLACK MESSAGE!'
```

### codedeploy to ec2

<https://support.atlassian.com/bitbucket-cloud/docs/deploy-to-aws-with-codedeploy/>

```yml
image: atlassian/default-image:2

pipelines:
  branches:
    master:
      - step:
          image: node:12.14.0
          name: Build
          caches:
            - node
          script:
            - npm ci --only=prod && npm run build
          artifacts:
            - '**/*'
            - '**/.*'

      - step:
          name: Zip
          script:
            - zip -r dist.zip ./
          artifacts:
            - dist.zip

      - step:
          name: Upload to S3
          services:
            - docker
          script:
            - pipe: atlassian/aws-code-deploy:0.2.10
              variables:
                AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
                AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}
                AWS_DEFAULT_REGION: ${AWS_DEFAULT_REGION}
                COMMAND: 'upload'
                APPLICATION_NAME: ${APPLICATION_NAME}
                S3_BUCKET: '${AWS_DEPLOY_BUCKET}'
                ZIP_FILE: 'dist.zip'

      - step:
          name: Deploy with CodeDeploy
          services:
            - docker
          script:
            - pipe: atlassian/aws-code-deploy:0.2.10
              variables:
                AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
                AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}
                AWS_DEFAULT_REGION: ${AWS_DEFAULT_REGION}
                COMMAND: 'deploy'
                APPLICATION_NAME: ${APPLICATION_NAME}
                DEPLOYMENT_GROUP: ${DEPLOYMENT_GROUP}
                S3_BUCKET: '${AWS_DEPLOY_BUCKET}'
                IGNORE_APPLICATION_STOP_FAILURES: 'true'
                # FILE_EXISTS_BEHAVIOR: 'OVERWRITE'
                WAIT: 'true'
```
