# CLI env for Server

## sshd 설정

## sudoer 설정

```sh
export EDITOR=vim
visudo

# User privilege specification
<USER>    ALL=(ALL:ALL) NOPASSWD:ALL
```

## 환경 초기화 스크립트

- `git`
- `java`
- `bash-it`
- `fnm`
- `nodejs`
- `pm2`
- `.vimrc`

```bash
#!/usr/bin/env bash

# timezone
sudo ln -sf /usr/share/zoneinfo/Asia/Seoul /etc/localtime

if [[ $(grep '^ID=' /etc/os-release) == @(*centos*) ]]; then
  sudo yum update -y
  sudo yum install -y git
  # sudo yum install -y java-1.8.0-openjdk-devel.x86_64
elif [[ $(grep '^ID_LIKE=' /etc/os-release) == @(*debian*) ]]; then
  sudo apt update && sudo apt upgrade -y
  sudo apt install -y git
  # sudo apt install -y openjdk-8-jdk
fi

# install bash-it
git clone --depth=1 https://github.com/Bash-it/bash-it.git ~/.bash_it && ~/.bash_it/install.sh --silent
source .bashrc

sed -i "s/\"bobby\"/\"candy\"/" .bashrc

export BASH_IT="/home/$(whoami)/.bash_it"
source "$BASH_IT"/bash_it.sh
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
curl -fsSL https://fnm.vercel.app/install | bash

PATH=~/.fnm:$PATH
eval "$(fnm env --shell=bash)"

# set node version
NODE_VERSION=v12.20.0

fnm install $NODE_VERSION
fnm use $NODE_VERSION
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

## CodeDeploy agent in AmazonLinux

```sh
sudo yum install -y ruby
curl -O https://aws-codedeploy-ap-northeast-2.s3.amazonaws.com/latest/install
chmod +x install
sudo ./install auto

sudo sed -i 's/""/"ec2-user"/g' /etc/init.d/codedeploy-agent

# 중요: Amazon Linux 2 AMI의 경우 다음 추가 명령을 실행합니다.
sudo sed -i 's/#User=codedeploy/User=ec2-user/g' /usr/lib/systemd/system/codedeploy-agent.service

sudo systemctl daemon-reload
sudo chown ec2-user:ec2-user -R /opt/codedeploy-agent/
sudo chown ec2-user:ec2-user -R /var/log/aws/

sudo service codedeploy-agent start
sudo service codedeploy-agent status

ps aux | grep codedeploy-agent
```
