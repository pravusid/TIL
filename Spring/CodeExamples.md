# Spring MVC Codes

## 게시판

### 페이징

```java
private int rowSize; // 한 화면에 표시할 행
private int total; // 전체 게시물 수
private int start; // 출력 화면의 시작 행
private int end; // 출력 화면의 종료 행

private int blockSize;
private int totalPage; // 전체 페이지 수
private int startBlock; // block의 시작
private int endBlock; // block의 마지막
private int prevBtn; // 이전 block 버튼
private int nextBtn; // 다음 block 버튼

private int page;

{
  rowSize = 15;
  blockSize = 10;
}

public Map<String, Integer> calcPage(int total) {
  this.total = total;
  if (page == 0) { page = 1; }

  end = rowSize * page;
  start = end - rowSize + 1;
  if (end>total) { end=total; }

  Map<String, Integer> map = new HashMap();
  map.put("end", end);
  map.put("start", start);

  calcBlock(map);

  return map;
}

private void calcBlock(Map<String,Integer> map) {
  totalPage = (int) (Math.ceil((float)total/rowSize));
  startBlock = page - (page - 1) % blockSize;
  endBlock = startBlock + blockSize - 1;
  if (endBlock > totalPage) { endBlock = totalPage; }
  prevBtn = (startBlock==1)? 1: startBlock-1;
  nextBtn = (endBlock==totalPage)? totalPage: endBlock+1;
}
```

### 게시판 파일 업로드

```java
@RequestMapping("main/board_insert_ok.do")
public String board_insert_ok(DataBoardVO vo) {
  List<MultipartFile> list = vo.getUpload();
  if (list.isEmpty()) {
    vo.setFilename("");
    vo.setFilesize("");
    vo.setFilecount(0);
  } else {
    StringBuffer strName = new StringBuffer();
    StringBuffer strSize = new StringBuffer();
    for(MultipartFile mf : list) {
      try {
        String fileName = mf.getOriginalFilename();
        Long fileSize = mf.getSize();
        mf.transferTo(new File("c:\\upload\\"+fileName));
        strName.append(fileName+",");
        strSize.append(fileSize+",");

      } catch (IllegalStateException e) {
        e.printStackTrace();
      } catch (IOException e) {
        e.printStackTrace();
      }
    }
    vo.setFilename(strName.substring(0, strName.length()-1));
    vo.setFilesize(strSize.substring(0, strSize.length()-1));
    vo.setFilecount(list.size());
  }
  service.dataBoardInsert(vo);
  return "redirect:board_list.do";
}
```

### 파일 다운로드

```java
@RequestMapping("main/board_download")
public void board_download(String fn, HttpServletResponse resp) {
  try {
    File file = new File("c:\\upload\\"+fn);
    resp.setHeader("Content-Disposition", "attatchment;filename="+URLEncoder.encode(fn, "utf-8"));
    resp.setContentLength((int)file.length());

    BufferedInputStream bis = new BufferedInputStream(new FileInputStream(file));
    BufferedOutputStream bos = new BufferedOutputStream(resp.getOutputStream());
    byte[] b = new byte[1024];
    while (true) {
      if (bis.read(b)==-1) {break;}
      bos.write(b);
    }
    bis.close();
    bos.close();

  } catch (Exception e) {
    e.printStackTrace();
  }
}
```

## @ResponseBody 이용

### ResponseBody 일반 활용

```java
@RequestMapping("main/board_update_ok.do")
@ResponseBody
public String board_update_ok(DataBoardVO vo, int page) {
  boolean bChk = false;
  String send = null;
  if (bChk == true) {
    send = "<script>"
        +"location.href=\"board_content.do?no="+vo.getNo()+"&page="+page+"\";"
        +"</script>";
  } else {
    send = "<script>"
        +"alert(\"비밀번호가 일치하지 않습니다\");"
        +"history.back();"
        +"</script>";
  }
  return send;
}
```

### @ResponseBody에서 Object를 JSON으로 변환해서 반환

pom.xml

```xml
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.7.3</version>
</dependency>
```

jackson example

```java
@RequestMapping("/login")
public @ResponseBody UsersVO login(UsersVO vo) {
  vo = dao.selectUserData(email);
  return vo;
}
```

## eclipse-maven-plugin 배포 : pom.xml에 plugin 추가

- run as config에 goals : `tomcat7:redeploy`
  ```xml
  <plugin>
      <groupId>org.apache.tomcat.maven</groupId>
      <artifactId>tomcat7-maven-plugin</artifactId>
      <version>2.2</version>
      <configuration>
          <path>/</path>
          <url>http://211.238.142.123:80/manager/text</url>
          <username>itmenu</username>
          <password>unemti</password>
      </configuration>
  </plugin>
  ```

- tomcat-users.xml
  ```xml
  <role rolename="admin-gui"/>
  <role rolename="manager-gui"/>
  <role rolename="manager-script"/>
  <user username="tomcat" password="tomcat"  roles="admin-gui,manager-gui,manager-script" />
  ```

- 배포를 한 적이 있다면 다음 배포 이전에 mvn clean을 실행