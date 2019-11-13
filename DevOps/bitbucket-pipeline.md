# Bitbucket Pipeline

## Docs

- <https://confluence.atlassian.com/bitbucket/configuring-your-pipeline-872013574.html>
- <https://confluence.atlassian.com/x/14UWN>

## Example

```yml
image: node:10.15.3

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
                # DEBUG: "<boolean>" # Optional
            - pipe: atlassian/slack-notify:0.3.2
              variables:
                WEBHOOK_URL: $WEBHOOK_URL
                MESSAGE: 'SLACK MESSAGE!'
```
