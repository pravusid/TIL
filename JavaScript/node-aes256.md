# AES256 in nodejs

```ts
// ENC_KEY => AES256 -> 256bit (32bytes)
// INITIALIZATION_VECTOR => AES -> 128bit (16bytes) == block size

import { createCipheriv, createDecipheriv, randomBytes } from 'crypto';

export function encrypt(val: string, encKey: string) {
  const iv = randomBytes(16);
  const cipher = createCipheriv('aes-256-cbc', encKey, iv);

  let encrypted = iv.toString('base64');
  encrypted += cipher.update(val, 'utf8', 'base64');
  encrypted += cipher.final('base64');
  return encrypted;
}

export function decrypt(encrypted: string, encKey: string) {
  const iv = Buffer.from(encrypted.substring(0, 24), 'base64');
  const encStr = encrypted.substring(24);

  const decipher = createDecipheriv('aes-256-cbc', encKey, iv);

  let decrypted = decipher.update(encStr, 'base64', 'utf8');
  decrypted += decipher.final('utf8');
  return decrypted;
}
```
