# Cursor 사용 방법 (+ 최고의 팁)

<https://siosio3103.medium.com/e931ced0429f>

YOLO Mode

```md
any kind of tests are always allowed like vitest, npm test, nr test, etc. also basic build commands like build, tsc, etc.
creating files and making directories (like touch, mkdir, etc) is always ok too
```

복잡한 작업

```md
Create a function that converts a markdown string to an HTML string.
Write tests first, then the code, then run the tests and update the code until tests pass.
```

오류 수정

```md
I've got some build errors. Run nr build to see the errors, then fix them, and then run build until build passes.
```

로그 기반 디버깅과 반복적 수정

복잡하고 까다로운 문제를 다룰 때는, Cursor에 로그를 출력하도록 지시한 뒤 그 로그를 활용해 문제를 해결하는 방식이 꽤 유용합니다. 제가 자주 사용하는 방법은 이렇습니다.
기존 이슈를 처리하기 위해 여러 방법을 시도했지만, 원하는 결과가 안 나오는 경우 전 Cursor에 이렇게 말합니다.

- “코드에 로그를 추가해서 내부 동작을 좀 더 잘 파악할 수 있게 해 줘. 내가 코드를 실행한 후에 로그 결과를 너한테 다시 알려줄게.”
- Cursor는 주요 지점에 로그 코드를 추가할 겁니다.
- 코드를 실행하고 로그를 수집합니다.
- Cursor로 돌아가 다음과 같이 말합니다 “여기에 로그 결과물이 있어. 지금 보니까 어떤 문제가 있는 것 같아? 어떻게 고치면 될까?
- Cursor가 분석할 로그 결과물을 붙여 넣습니다.
