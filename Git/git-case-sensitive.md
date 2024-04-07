# Git Case Sensitive Troubleshooting

> 파일이름(대소문자) 변경 후 오류 발생

```sh
Type error: Cannot find module '../foo/Bar' or its corresponding type declarations.
```

> <https://stackoverflow.com/questions/67420394/vercel-deploy-build-fail-failed-to-compile-type-error-cannot-find-module>

```sh
git rm -r --cached .
git add --all .
git commit -a -m "fix: untracked files"
git push
```
