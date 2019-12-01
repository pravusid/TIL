# Google Chrome in Linux

## CentOS7

```sh
sudo vim /etc/yum.repos.d/google-chrome.repo
```

repository 연결

```conf
[google-chrome]
name=google-chrome
baseurl=http://dl.google.com/linux/chrome/rpm/stable/$basearch
enabled=1
gpgcheck=1
gpgkey=https://dl-ssl.google.com/linux/linux_signing_key.pub
```

설치: `yum install google-chrome-stable`

버전확인: `google-chrome --version`
