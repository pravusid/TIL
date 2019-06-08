# CLI env for Server

## bash-it

### 설치

`git clone --depth=1 https://github.com/Bash-it/bash-it.git ~/.bash_it && ~/.bash_it/install.sh`

### 플러그인 설정

플러그인 확인: `bash-it show plugins`

플러그인 활성화: `bash-it enable plugin dirs docker git z`

### Theme

`.bashrc`

```sh
export BASH_IT_THEME='candy'
```

## `.vimrc`

```sh
set number
set cursorline
set hlsearch
set smartindent
set expandtab
set tabstop=4
set softtabstop=4
set shiftwidth=4
set showmatch

set list
set listchars=tab:\|\ ,trail:·

ca ㅈ w
ca ㅈㅂ wq

map <F12> mzgg=G`z
```

### `.bashrc`

`.bashrc.alias`

```sh
alias lsg="\ls -al | grep "
alias pg="\pgrep -fl"
alias cntf="\ls -1A | wc -l"
alias ssa="\ss -natp"
alias ssl="\ss -nltp"
alias tf="\tail -f"
```

`.bashrc`

```sh
[ -f ~/.bashrc.alias ] && source ~/.bashrc.alias
```
