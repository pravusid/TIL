# Hadoop 기초

## Hadoop 실행

- 하둡 초기화 `hadoop namenode -format`

- 하둡 서버 실행 / 중지

  ```sh
  start-all.sh
  stop-all.sh
  ```

- jps : 하둡 작동 프로세서 확인

## Hadoop 명령어

- 하둡에 파일 전송 : `hadoop fs -appendToFile /home/hostname/{local} /{hadoop}`
- 하둡에서 다운로드 : `hadoop -copyToLocal`
- 하둡 파일 보기 : `hadoop fs -cat /path`
- 하둡 파일삭제 : `hadoop fs -rmr /path`
- 하둡 폴더생성 : `hadoop fs -mkdir /path`
- 하둡에 파일 삭제 : `hadoop fs rm -r /test_ns1`
