# GitHub Packages

## NPM

<https://help.github.com/en/github/managing-packages-with-github-packages/configuring-npm-for-use-with-github-packages>

## with Yarn

<https://github.com/yarnpkg/yarn/issues/4451>

`yarn`에서 private package 사용시 `.npmrc` 설정

```conf
//npm.pkg.github.com/:_authToken=<TOKEN>
<@ORG_NAME>:registry=https://npm.pkg.github.com/
```
