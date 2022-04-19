# AES256 in nodejs

## aes-256-gcm

- <https://gist.github.com/rjz/15baffeab434b8125ca4d783f4116d81>
- <https://gist.github.com/AndiDittrich/4629e7db04819244e843>
- <https://stackoverflow.com/questions/53269132/aes-256-gcm-encryption-decryption-in-nodejs>

> In GCM mode, the `authTagLength` option is not required
> but can be used to set the length of the authentication tag that will be returned by `getAuthTag()` and defaults to 16 bytes
-- <https://nodejs.org/api/crypto.html#cryptocreatecipheralgorithm-password-options>

## aes-256-cbc (examples)

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
