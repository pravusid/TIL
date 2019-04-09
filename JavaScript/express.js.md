# Express.js

## 배포

배포 후 환경변수를 production으로 변경해야 함: `NODE_ENV=production node app.js`

앱 내부에서 `process.env.NODE_ENV` 값에 할당되어 express 배포시 최적화 처리됨
