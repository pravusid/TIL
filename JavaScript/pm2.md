# PM2

Process Manager 2

<https://pm2.io/doc/en/runtime/overview/>

## 시작

### Install

```sh
sudo npm install pm2@latest -g
pm2 kill
```

### Update

```sh
npm install pm2@latest -g
pm2 update
```

### 확장 모듈

```sh
# pm2 tab-completion
pm2 completion install
# log rotate https://github.com/keymetrics/pm2-logrotate
pm2 install pm2-logrotate
```

## CLI

<https://pm2.io/doc/en/runtime/reference/pm2-cli/>

```sh
pm2 start app.js
pm2 ls
pm2 restart app
pm2 stop app
pm2 delete app

# save your list in hard disk memory
pm2 save
# resurrect your list previously saved
pm2 resurrect

# all apps logs
pm2 logs
# only app logs
pm2 logs app
# empty all application logs
pm2 flush
```

## Echosystem File

<https://pm2.io/doc/en/runtime/guide/ecosystem-file/>

`pm2 init`: `ecosystem.config.js` 파일이 생성됨

```js
module.exports = {
  apps: [
    {
      name: "app",
      script: "./app.js",
      instances: "max", // n (n개), -n (core 수 - n개)
      output: "./logs/out.log",
      error: "./logs/error.log",
      log_type: "json",
      merge_logs: true,
      env: {
        NODE_ENV: "production"
      }
    }
  ]
};
```

`pm2 start [/path/to/ecosystem.config.js] [--only app] [--env production]`

설정된 환경변수는 재시작해도 변하지 않으므로 변경을 위해서는 update 명령을 해야함

`pm2 restart ecosystem.config.js --update-env [--env production]`

## Cluster

`pm2 start app.js -i <number-instances>`

`pm2 reload <app_name>`: 한번에 종료하지 않고 클러스터 내에서 순차적으로 재시작
