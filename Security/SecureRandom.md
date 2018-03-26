# Secure Random

Java Random 클래스는 불완전한 난수를 생성하기 때문에 보안을 위해 난수가 필요한 상황에서는 Secure Random 을 사용하는 것이 좋다.

자바 8에서는 `SecureRandom.getInstanceStrong()`이 추가되었는데,
리눅스 환경에서 `/dev/random` 를 호출하여 엄청난 퍼포먼스 저하를 불러일으킨다.

자바 8에서 기본 인스턴스인 `new SecureRandom()` 으로 `/dev/urandom`을 호출하여 충분한 랜덤문자를 생성할 수 있다.

또한 자바는 기본적으로 운영체제의 CSPRNG를 읽어오기 때문에

> CSPRNG : Cryptographically secure pseudorandom number generator

RNG seed를 임의로 지정할 필요가 없다.

> RNG seed : Random number generator seed

## 사용법

```java
import java.security.NoSuchAlgorithmException;
import java.security.SecureRandom;

public class KeyGen {
    public void generate() throws NoSuchAlgorithmException {
        SecureRandom random = SecureRandom.getInstanceStrong();
        byte[] values = new byte[32]; // 256 bit
        random.nextBytes(values);
        StringBuilder sb = new StringBuilder();
        for (byte b : values) {
          sb.append(String.format("%02x", b));
        }
        System.out.print("Key: ");
        System.out.println(sb.toString());
    }
}
```
