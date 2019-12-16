# CLI env for Server

환경 초기화 스크립트

- `bash-it`
- `java`
- `fnm`
- `nodejs`
- `pm2`
- `.vimrc`

```bash
#!/usr/bin/env bash

sudo yum update -y

# timezone
sudo ln -sf /usr/share/zoneinfo/Asia/Seoul /etc/localtime

# install git
sudo yum install -y git

# install java
sudo yum install -y java-1.8.0-openjdk-devel.x86_64

# install bash-it
git clone --depth=1 https://github.com/Bash-it/bash-it.git ~/.bash_it && ~/.bash_it/install.sh
source .bashrc

sed -i "s/'bobby'/\"candy\"/" .bashrc

bash-it enable plugin dirs docker git z

cat > .bashrc.alias <<- EOM
alias l="ls"
alias la="ls -A"
alias ll="ls -AFlh"

alias lg="\ls -al | grep"
alias pg="\pgrep -fl"
alias cntf="\ls -1A | wc -l"
alias ssa="\ss -natp"
alias ssl="\ss -nltp"
alias tf="\tail -f"
EOM

echo "[ -f ~/.bashrc.alias ] && source ~/.bashrc.alias" >> .bashrc
source .bashrc

# install fnm
curl -fsSL https://github.com/Schniz/fnm/raw/master/.ci/install.sh | bash

export PATH=~/.fnm:$PATH
eval "`fnm env --multi`"

fnm use latest-erbium
fnm default $(node -v | sed 's/v//')
source .bashrc

# install pm2
npm i -g pm2

# vimrc
cat > .vimrc <<- EOM
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

map <F12> mzgg=G\`z
EOM
```
