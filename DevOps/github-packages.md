# GitHub Packages

## NPM

<https://docs.github.com/en/free-pro-team@latest/packages/guides/configuring-npm-for-use-with-github-packages>

## `.npmrc`

<https://docs.npmjs.com/cli/v6/configuring-npm/npmrc>

<https://github.com/yarnpkg/yarn/issues/4451>

`yarn`에서 private package 사용시 `.npmrc` 설정

```conf
//npm.pkg.github.com/:_authToken=<TOKEN>
<@ORG_NAME>:registry=https://npm.pkg.github.com/
```
