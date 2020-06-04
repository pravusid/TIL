# Git Config

## overall

```sh
[user]
    name = Sangdon Park
    email = pravusid@gmail.com
    signingkey = 4B0A009C0CC438F4
[core]
    autocrlf = input
    excludesfile = ~/.gitignore_global
[commit]
    gpgsign = true
[diff]
    tool = vimdiff
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

## include / includeIf

```sh
[include]
    path = ~/.gitconfig_private
[includeIf "gitdir:~/Documents/dev/**"]
    path = ~/.gitconfig_work
```

## multiple ssh-keys

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
