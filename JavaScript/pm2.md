# PM2

Process Manager 2

> <https://pm2.keymetrics.io/docs/usage/quick-start/>

## 시작

```sh
# Install
sudo npm install pm2@latest -g

# Update
pm2 update
```

## Cluster

```sh
// cluster: n (n개), -n (core 수 - n개)
pm2 start app.js -i <number-instances>

# 한번에 종료하지 않고 클러스터 내에서 순차적으로 재시작
pm2 reload <app_name>
```

## CLI

> <https://pm2.keymetrics.io/docs/usage/quick-start/#cheatsheet>

```sh
pm2 start app.js --name <app_name>

pm2 reload <app_id|app_name|all>
pm2 restart <app_id|app_name|all>

pm2 stop <app_id|app_name|all>
pm2 delete <app_id|app_name|all>

# Display all processes status
pm2 ls

# Monitor all processes
pm2 monit

# all apps logs / only app logs / only out or err
pm2 logs [--raw] [app_id|app_name] [--err | --out]

# empty all application logs
pm2 flush

# Reload all logs
pm2 reloadLogs

# Send system signal to script
pm2 sendSignal SIGUSR2 my-app
```

## Ecosystem File

> <https://pm2.keymetrics.io/docs/usage/application-declaration/#attributes-available>

`pm2 ecosystem`: `ecosystem.config.js` 파일이 생성됨

```js
module.exports = {
  apps: [
    {
      name: "app",
      script: "./app.js",
      instances: "max", // cluster
      exp_backoff_restart_delay: 100,
      max_memory_restart: "200M",
      output: "./logs/out.log",
      error: "./logs/error.log",
      log_type: "json",
      merge_logs: true,
      env: {
        NODE_ENV: "development"
      },
      env_production: {
        NODE_ENV: "production"
      }
    }
  ]
};
```

`pm2 <start|restart|stop|delete> [/path/to/ecosystem.config.js] [--only app]`

설정된 환경변수는 재시작해도 변하지 않으므로 변경을 위해 환경변수를 명시함

> <https://pm2.keymetrics.io/docs/usage/application-declaration/#switching-environments>

`pm2 <restart|reload> <app_name> --env production`

## Graceful Start/Stop

<https://pm2.keymetrics.io/docs/usage/signals-clean-restart/>

### Graceful Start

앱 실행까지 준비시간이 필요할 수 있다 (db연결, 데이터로드 ...)

pm2를 통해 실행된 어플리케이션에서는 `wait_ready: true` 옵션을 활성화하고
아래의 코드를 앱에서 실행하여 PM2에게 `ready` signal을 직접 보낼 수 있다

```js
process.send("ready");
```

### Graceful Stop

SIGINT signal을 가로채서 앱 종료준비를 한다 (`kill_timeout`내에 작업을 끝내야 한다)

```sh
process.on('SIGINT', function() {
   db.stop(function(err) {
     process.exit(err ? 1 : 0);
   });
});
```

### ecosystem file

```js
module.exports = {
  apps: [
    {
      name: "app",
      script: "./app.js",
      wait_ready: true, // Instead of reload waiting for listen event, wait for process.send(‘ready’)
      listen_timeout: 3000, // (default) time in ms before forcing a reload if app not listening
      kill_timeout: 1600 // (default) time in milliseconds before sending a final SIGKILL
    }
  ]
};
```

## log rotate

> <https://github.com/keymetrics/pm2-logrotate>

```sh
pm2 install pm2-logrotate
```
