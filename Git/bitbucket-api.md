# Bitbucket API

## repository 경로 파일다운로드

<https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/{username}/{repo_slug}/downloads/{filename}>

```sh
curl -s -S --user username:apppassword -L -O \
  https://api.bitbucket.org/2.0/repositories/<ORGANISATION>/<REPO>/src/master/<FOLDER>/<FILE>
```
