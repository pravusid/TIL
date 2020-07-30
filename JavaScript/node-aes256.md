# AES256 in nodejs

```ts
import { createCipheriv, createDecipheriv } from 'crypto';

const ENC_KEY = 'SOME_ENC_KEY'; // AES256 -> 256bit (32bytes)
const IV = 'SOME_INITIALIZATION_VECTOR'; // AES -> 128bit (16bytes) == block size

export function encrypt(val: string) {
  const cipher = createCipheriv('aes-256-cbc', ENC_KEY, IV);
  let encrypted = cipher.update(val, 'utf8', 'base64');
  encrypted += cipher.final('base64');
  return encrypted;
}

export function decrypt(encrypted: string) {
  const decipher = createDecipheriv('aes-256-cbc', ENC_KEY, IV);
  const decrypted = decipher.update(encrypted, 'base64', 'utf8');
  return decrypted + decipher.final('utf8');
}
```
