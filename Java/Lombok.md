# Lombok

## Annotations

- `@Data`
- `@Getter`
- `@Setter`
- `@NoArgsConstructor`
- `@AllArgsConstructor`

## 설치

빌드툴에서 Lombok 의존성 추가: <https://mvnrepository.com/artifact/org.projectlombok/lombok>

## IDE별 설정

### IntelliJ IDEA

플러그인 설치

- Settings > Pluglins > Browse Repositories
- `Lombok Plugin` 검색하여 설치

빌드시 어노테이션 적용 설정

- Settings > Build, Execution, Deployment > Complier > Annotation Processors
- `Enable annotation processing` 체크

### Eclipse

Eclipse를 종료한 상태에서 아래의 과정을 실행한다.

`lombok.jar`가 다운로드 된 경로에 가서 `lombok.jar` 실행 (이름에 버전이 붙어있음)

- 경로: `~/.m2/repository/org/projectlombok/lombok`
- 실행: 더블클릭 or `java -jar lombok.jar`

Lombok 인스톨러가 실행되고 `Specify location`을 눌러서 Eclipse 실행파일을 선택한 후 `install/update` 클릭
