# Husky

- <https://www.npmjs.com/package/husky>
- <https://typicode.github.io/husky/>

## alternatives

- <https://www.npmjs.com/package/lefthook>

## 설치

`npx husky init && npm i`

## 설정

`package.json`

```sh
{
  "scripts": {
    "prepare": "if [ \"$CI\" != \"true\" ]; then husky; fi"
  }
}
```

## hooks

### node package lock 변경감지

<https://jshakespeare.com/use-git-hooks-and-husky-to-tell-your-teammates-when-to-run-npm-install/>

`.husky/post-merge`

```sh
function changed {
  git diff --name-only HEAD@{1} HEAD | grep "^$1" > /dev/null 2>&1
}

if changed 'package-lock.json'; then
  echo "📦 package-lock.json changed. Run npm install to bring your dependencies up to date."
  npm ci
fi
```

### hooks 실행권한

```sh
chmod +x ./husky/*
```
