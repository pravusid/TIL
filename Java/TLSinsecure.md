# TLS - insecure

TLS 프로토콜 사용시 인증서 검증을 하지 않음

```kt
import java.net.Socket
import java.security.cert.CertificateException
import java.security.cert.X509Certificate
import javax.net.ssl.*

object SslSocketFactory {
    val trustAllCerts = arrayOf<TrustManager>(object : X509ExtendedTrustManager() {
        @Throws(CertificateException::class)
        override fun checkClientTrusted(x509Certificates: Array<X509Certificate>, type: String) {
        }

        @Throws(CertificateException::class)
        override fun checkServerTrusted(x509Certificates: Array<X509Certificate>, type: String) {
        }

        override fun getAcceptedIssuers(): Array<X509Certificate>? {
            return arrayOf()
        }

        @Throws(CertificateException::class)
        override fun checkClientTrusted(x509Certificates: Array<X509Certificate>, type: String, socket: Socket) {
        }

        @Throws(CertificateException::class)
        override fun checkServerTrusted(x509Certificates: Array<X509Certificate>, type: String, socket: Socket) {
        }

        @Throws(CertificateException::class)
        override fun checkClientTrusted(x509Certificates: Array<X509Certificate>, type: String, sslEngine: SSLEngine) {
        }

        @Throws(CertificateException::class)
        override fun checkServerTrusted(x509Certificates: Array<X509Certificate>, type: String, sslEngine: SSLEngine) {
        }
    })

    fun getFactory(): SSLSocketFactory {
        val sslContext = SSLContext.getInstance("TLS")
        sslContext.init(null, trustAllCerts, java.security.SecureRandom())
        return sslContext.socketFactory
    }
}
```

클라이언트에서 `SSLSocketFactory`를 다음처럼 사용함

```kt
val client = OkHttpClient.Builder()
                .connectionSpecs(listOf(ConnectionSpec.MODERN_TLS, ConnectionSpec.COMPATIBLE_TLS))
                .sslSocketFactory(SslSocketFactory.getFactory(), SslSocketFactory.trustAllCerts.first() as X509TrustManager)
                .hostnameVerifier { _, _ -> true }.build()
val request = Request.Builder().url("https://localhost:8080").build()
```
