# Bitbucket API

<https://developer.atlassian.com/cloud/bitbucket/rest/intro/>

## repository 경로 파일다운로드

`https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/{username}/{repo_slug}/downloads/{filename}`

```bash
curl -s -S --user USERNAME:APPPASSWORD -L -O \
  https://api.bitbucket.org/2.0/repositories/<ORGANISATION>/<REPO>/src/main/<FOLDER>/<FILE>
```

```bash
curl -s -S -H "Authorization: Bearer BITBUCKET_ACCESS_TOKEN" -L -O \
  https://api.bitbucket.org/2.0/repositories/<ORGANISATION>/<REPO>/src/main/<FOLDER>/<FILE>
```
