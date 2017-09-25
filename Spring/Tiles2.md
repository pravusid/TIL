# Spring MVC Tiles2 설정

## ViewResolver로 Tiles2 활용

default.jsp

```jsp
<%@ taglib prefix="tiles" uri="http://tiles.apache.org/tags-tiles" %>
<tiles:insertAttribute name="header"/>
```

## Tiles2 레이아웃 구성예제

tiles.xml

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE tiles-definitions PUBLIC
"-//Apache Software Foundation//DTD Tiles Configuration 3.0//EN"
"http://tiles.apache.org/dtds/tiles-config_3_0.dtd">
<tiles-definitions>
  <definition name="default" template="/WEB-INF/views/default.jsp">
    <put-attribute name="header" value="/WEB-INF/views/header.jsp"/>
    <put-attribute name="navi" value="/WEB-INF/views/navi.jsp"/>
    <put-attribute name="body" value="/WEB-INF/views/body.jsp"/>
    <put-attribute name="footer" value="/WEB-INF/views/footer.jsp"/>
  </definition>
  <definition name="*" extends="default">
    <put-attribute name="body" value="/WEB-INF/views/{1}.jsp"/>
  </definition>
  <definition name="*/*" extends="default">
    <put-attribute name="body" value="/WEB-INF/views/{1}/{2}.jsp"/>
  </definition>
  <definition name="*/*/*" extends="default">
    <put-attribute name="body" value="/WEB-INF/views/{1}/{2}/{3}.jsp"/>
  </definition>
</tiles-definitions>
```
