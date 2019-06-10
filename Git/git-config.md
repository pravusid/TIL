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
```

## include / includeIf

```sh
[include]
    path = ~/.gitconfig_private
[includeIf "gitdir:~/Documents/dev/**"]
    path = ~/.gitconfig_work
```
