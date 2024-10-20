# NEXT.js image optimization

- <https://github.com/vercel/next.js/blob/main/packages/next/src/server/image-optimizer.ts>
- <https://nextjs.org/docs/pages/building-your-application/optimizing/images>
- <https://fe-developers.kakaoent.com/2022/220714-next-image/>

## `next.config.js`

<https://nextjs.org/docs/app/api-reference/next-config-js/images>

```js
/** @type {import('next').NextConfig} */
module.exports = {
  images: {
    remotePatterns: [{ hostname: 'localhost' }, { hostname: 'foo.remote-host.com' }],
    minimumCacheTTL: 31536000,
    formats: ['image/avif', 'image/webp'],
  },
  async headers() {
    return [
      {
        source: '/(images|videos)/(.*)',
        headers: [
          {
            key: 'Cache-Control',
            value: 'public, max-age=86400, s-maxage=31536000, stale-while-revalidate=604800',
          },
        ],
      },
    ];
  },
};
```

## Custom Loader, Middlewares

static(소스코드의 public 경로에 포함된) image가 아닌 경우 최적화된 이미지의 응답헤더 cache-control 값은 `public, max-age=${maxAge}, must-revalidate`임

> <https://github.com/vercel/next.js/blob/9a1cd356dbafbfcf23d1b9ec05f772f766d05580/packages/next/src/server/image-optimizer.ts#L701>

이 경우 [Custom Loader](https://nextjs.org/docs/app/api-reference/next-config-js/images#example-loader-configuration)를 사용해서 CDN 캐시설정을 그대로 사용하도록 수정할 수 있음.
그러나 CDN에서 edge function 등을 이용한 이미지 최적화도 작성해야 하므로 번거롭다면 [Middleware](https://nextjs.org/docs/app/building-your-application/routing/middleware)를 사용한 꼼수로 임시처리 할 수 있음.

```js
import { NextResponse } from 'next/server';

export function middleware(request) {
  const response = NextResponse.next();

  if (request.nextUrl.pathname.startsWith('/some-path')) {
    response.headers.set('Cache-Control', 'public, max-age=315360000, immutable');
  }

  return response;
}
```
