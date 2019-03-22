# AWS sdk for JavaScript

aws credentials(`aws.credentials.json`)

```json
{
  "region": "ap-northeast-2",
  "accessKeyId": "<ACCESS_KEY_ID>",
  "secretAccessKey": "<SECRET_ACCESS_KEY>"
}
```

sdk 설정

```ts
import * as AWS from "aws-sdk";

const configs = require("./aws.credentials.json");

AWS.config.update({
  region: configs.region,
  accessKeyId: configs.accessKeyId,
  secretAccessKey: configs.secretAccessKey
});

export const awsDynamo = new AWS.DynamoDB.DocumentClient({
  convertEmptyValues: true
});

export const awsS3 = new AWS.S3({ apiVersion: "2006-03-01" });
```
