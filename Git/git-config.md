# Git Config

## overall

```sh
[user]
    name = Sangdon Park
    email = sandpark@pravusid.kr
    signingkey = SOME_SIGN_KEY

[core]
    autocrlf = input
    excludesfile = ~/.gitignore_global

[commit]
    gpgsign = true

[pull]
    rebase = true

[alias]
    a = "!git add $(git status -s | fzf -m | awk '{print $2}')"
    b = "!git checkout $(\
        _height=$(stty size | awk '{print $1}');\
        git branch | egrep -v '^\\*' | fzf --preview \"git log {1} | head -n $_height\";\
    )"
```

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
