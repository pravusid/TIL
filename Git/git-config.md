# Git Config

## overall

<https://github.com/pravusid/sys-config/blob/main/.gitconfig>

> ref: <https://johngrib.github.io/wiki/git-alias/>

## git diff with delta

<https://github.com/dandavison/delta>

```conf
[core]
    pager = delta

[delta]
    plus-color = "#012800"
    minus-color = "#340001"
    syntax-theme = Nord

[interactive]
    diffFilter = delta --color-only
```

## include / includeIf

```sh
[include]
    path = ~/.gitconfig_common

[includeIf "gitdir:~/Documents/private/**"]
    path = ~/.gitconfig_private

[includeIf "gitdir:~/Documents/work/**"]
    path = ~/.gitconfig_work
```

## multiple ssh-keys

`~/.ssh/config`

```conf
Host github-private
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_rsa_private

Host github-work
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_rsa_work
```

`.git/config`

```conf
[remote "orgin"]
    url = git@github-work:{GithubID}/{RepositoryName}.git
```
