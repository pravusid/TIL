# Git Config

<https://github.com/pravusid/dotfiles/blob/main/.gitconfig>

## refs

- <https://git-scm.com/docs/git-config>
- <https://johngrib.github.io/wiki/git-alias/>
- <https://gist.github.com/junegunn/f4fca918e937e6bf5bad>

## git diff with delta

- <https://github.com/dandavison/delta>
- <https://dandavison.github.io/delta/>

## include / includeIf

```conf
[include]
    path = ~/.gitconfig_common

[includeIf "gitdir:~/Documents/private/**"]
    path = ~/.gitconfig_private

[includeIf "gitdir:~/Documents/work/**"]
    path = ~/.gitconfig_work
```

## multiple ssh-keys

### (방법1) git_config 사용

> gitconfig include / includeIf 옵션과 함께 사용

```conf
[core]
    sshCommand = ssh -o IdentitiesOnly=yes -i ~/.ssh/id_rsa_git -F /dev/null
```

### (방법2) ssh_config 사용

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
