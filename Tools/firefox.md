# Firefox

## AI Chatbot 추가 설정

<https://connect.mozilla.org/t5/discussions/advanced-configuration-of-the-ai-chatbot-in-firefox/td-p/85031>

`about:config`

- `browser.ml.chat.prompts.8`
  - `{"label": "번역","id": "translate","value": "<instruction>내용을 한국어로 번역해줘. 절대! 요약하지 말고 원문을 그대로 사용해줘.</instruction>"}`
- `browser.ml.chat.prompts.9`
  - `{"label": "취합","id": "translate","value": "<instruction>본문과 댓글을 사용자별로 요약해서 한국어로 출력해줘</instruction>"}`

## 세션 복구

- 주소창에 about:support 입력
- 아래쪽의 프로필 폴더 경로 확인 후, 해당 폴더로 이동
- 그 안의 sessionstore-backups 폴더 확인
- `recovery.jsonlz4`, `previous.jsonlz4`, `upgrade.jsonlz4-XXXX` 등 파일 중에서 가장 최근의 파일을 `sessionstore.jsonlz4`로 복사/이름변경
- Firefox 완전히 종료 후 프로필 폴더에 `sessionstore.jsonlz4` 파일 붙여넣기
- 다시 실행하면 복원되어 있음
