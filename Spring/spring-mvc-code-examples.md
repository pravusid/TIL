# Spring MVC Codes

## 게시판

### 페이징(Pagination)

```java
package kr.pravusid.dto;

import org.springframework.data.domain.Page;

public class Pagination {

    private int currPage; // range: 0 ~
    private int totalPages; // 게시물 없을 때 0, 게시물 있을 때 range: 1 ~

    private int firstBlock; // block 시작, range: 0 ~
    private int lastBlock; // block 마지막, range: 0 ~

    private int prev; // 좌측 화살표 (전 블록 마지막으로)
    private int next; // 우측 화살표 (다음 블록 처음으로)

    /* 검색 관련 parameter */
    private String filter;
    private String keyword;

    public enum FilterType {
        TITLE, CONTENT, USER, COMMENTS, ALL;
    }

    public Pagination calcPage(Page page, int blockSize) {
        this.currPage = page.getNumber();
        this.totalPages= (page.getTotalPages() == 0) ? 0 : page.getTotalPages() - 1;

        firstBlock = currPage - (currPage % blockSize);
        lastBlock = (firstBlock + (blockSize - 1) > totalPages) ? totalPages : firstBlock + (blockSize - 1);
        prev = (firstBlock == 0) ? 0 : firstBlock - 1;
        next = (lastBlock == totalPages) ? totalPages : lastBlock + 1;

        return this;
    }

    public int getCurrPage() {
        return currPage;
    }

    public int getTotalPages() {
        return totalPages;
    }

    public int getFirstBlock() {
        return firstBlock;
    }

    public int getLastBlock() {
        return lastBlock;
    }

    public int getPrev() {
        return prev;
    }

    public int getNext() {
        return next;
    }

    public String getFilter() {
        return filter;
    }

    public void setFilter(String filter) {
        this.filter = filter;
    }

    public boolean filterMatcher(FilterType type) {
        return Pagination.FilterType.valueOf(this.filter.toUpperCase()).equals(type);
    }

    public String getKeyword() {
        return keyword;
    }

    public void setKeyword(String keyword) {
        this.keyword = keyword;
    }

    public String getSearchQuery() {
        return (filter == null || keyword == null) ? "" : "&filter=" + filter + "&keyword=" + keyword;
    }

}
```

### 파일 업로드

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
