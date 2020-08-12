# AES256 in Kotlin

- common codecs: <https://mvnrepository.com/artifact/commons-codec/commons-codec>

- crypto policy (`java.security.InvalidKeyException: Illegal key size`)
  - 파일교체(~ JDK1.8): <https://www.oracle.com/technetwork/java/javase/downloads/jce8-download-2133166.html>
  - 수정 (JDK1.8u151 ~)
    - `$JAVA_HOME/jre/lib/security`
    - crypto.policy: (limited -> unlimited)

```kt
import com.typesafe.config.ConfigFactory
import kotlinx.io.charsets.Charset
import org.apache.commons.codec.binary.Base64
import java.security.Key
import javax.crypto.Cipher
import javax.crypto.spec.IvParameterSpec
import javax.crypto.spec.SecretKeySpec

object AES256 {
    private val config = ConfigFactory.load()
    private val key = config.getString("ktor.secrets.logEncodingKey")

    private val iv: String
    private val keySpec: Key

    init {
        iv = key.substring(0, 16)
        val keyBytes = ByteArray(16)
        val b = key.toByteArray(charset("UTF-8"))
        var len = b.size
        if (len > keyBytes.size) {
            len = keyBytes.size
        }
        System.arraycopy(b, 0, keyBytes, 0, len)
        keySpec = SecretKeySpec(keyBytes, "AES")
    }

    fun encrypt(str: String): String {
        val c = Cipher.getInstance("AES/CBC/PKCS5Padding")
        c.init(Cipher.ENCRYPT_MODE, keySpec, IvParameterSpec(iv.toByteArray()))
        val encrypted = c.doFinal(str.toByteArray(charset("UTF-8")))
        return String(Base64.encodeBase64(encrypted))
    }


    fun decrypt(str: String): String {
        val c = Cipher.getInstance("AES/CBC/PKCS5Padding")
        c.init(Cipher.DECRYPT_MODE, keySpec, IvParameterSpec(iv.toByteArray()))
        val byteStr = Base64.decodeBase64(str.toByteArray())
        return String(c.doFinal(byteStr), Charset.defaultCharset())
    }
}
```
