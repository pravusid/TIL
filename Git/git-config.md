# Git Config

<https://github.com/pravusid/dotfiles/blob/main/.gitconfig>

## git utils

- <https://github.com/jesseduffield/lazygit>
- <https://github.com/extrawurst/gitui>
- <https://github.com/bigH/git-fuzzy>
- <https://github.com/arxanas/git-branchless>
- <https://github.com/gitext-rs/git-stack>

## refs

- <https://git-scm.com/docs/git-config>
- <https://johngrib.github.io/wiki/git-alias/>
- <https://gist.github.com/junegunn/f4fca918e937e6bf5bad>

## git diff with delta

- <https://github.com/dandavison/delta>
- <https://dandavison.github.io/delta/>

## 키 생성

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
chmod 400 ~/.ssh/id_rsa

# macOS 키체인 등록
ssh-add --apple-use-keychain ~/.ssh/id_rsa
```

이후 서버(github.com 등)에 공개키 등록

## ssh config for github

```conf
Host *
    IdentitiesOnly yes
    AddKeysToAgent yes
    UseKeychain yes

Host github.com
    IdentityFile ~/.ssh/id_rsa
```

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
