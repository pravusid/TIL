# 코틀린 제네릭스

## 제네릭 타입 파라미터

제네릭스를 사용하면 type parameter를 받는 타입을 정의할 수 있다.

코틀린 컴파일러는 보통 타입과 마찬가지로 타입 인자도 추론할 수 있다.
`val authors = listOf("Dmitry", "Svetlana")`

빈 리스트를 만든다면 타입인자를 추론할 근거가 없으므로 직접 타입인자를 명시해야 한다.
`val readers: MutableList<String> = mutableListOf()` 또는 `val readers = mutableListOf<String>()`

### 제네릭 함수와 프로퍼티
