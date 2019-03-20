# XML parsing

## 자바에서 XML 파싱방법

- JAXP
  - SAX -> Spring, MyBatis
  - DOM
- JAXB (class와 xml을 binding) (언마셜/마셜)

## XML DTD에 정의된 태그들

- 최상위 태그(a, b, c, d, e) : 사용 순서
- a? : 사용강제X
- b* : 사용강제X & 여러번 사용가능
- c+ : 반드시 한 번 이상 사용
- d : 반드시 한 번만 사용
- e|k : 둘 중 하나는 반드시 사용

## JAXB 사용예제

- xml 태그에 대응되는 클래스 생성
  - 최상위 태그 : `@XmlRootElement`
  - setter : `@XmlElement`, `@XmlAttribute`

```java
List<Item> list = new ArrayList<Item>();
try {
  URL url = new URL("XML경로"));
  JAXBContext jc = JAXBContext.newInstance(Rss.class);
  Unmarshaller unms = jc.createUnmarshaller();
  Rss rss = (Rss) unms.unmarshal(url);
  list = rss.getChannel().getItem();
} catch (Exception e) {
  /* 에러처리 */
}
```
