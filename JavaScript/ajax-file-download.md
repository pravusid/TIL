# Download files via AJAX

## client

axios 사용을 가정한 예제

```js
const response = await this.axios.get(link, { responseType: "blob" });
const url = window.URL.createObjectURL(
  new Blob([response.data], { type: response.headers["content-type"] })
);

const contentDisposition = response.headers["content-disposition"];
const fileName = contentDisposition
  ? contentDisposition.match(/filename\*=UTF-8''(.+)/)[1]
  : String(new Date().getTime());

const tag = document.createElement("a");
tag.href = url;
tag.setAttribute("download", fileName);
document.body.appendChild(tag);
tag.click();

tag.remove();
window.URL.revokeObjectURL(url);
```

## server

CORS 설정: `Access-Control-Expose-Headers: Content-Disposition`

express.js 예제

```js
app.use(
  cors({
    origin: "http://localhost:8080",
    credentials: true,
    exposedHeaders: ["Content-Disposition"]
  })
);
```

response

```js
resp.setHeader("Content-Type", contentType);
resp.setHeader(
  "Content-Disposition",
  `attachment;filename*=UTF-8''${encodeURIComponent(fileName)}`
);

stream.pipe(resp);
stream.on("error", err => {
  resp.status(404).json({ message: "요청하신 파일이 존재하지 않습니다" });
});
```
