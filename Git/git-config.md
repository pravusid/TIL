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
