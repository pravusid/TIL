# 스프링 Validation

## @Valid

Controller에서 `@Valid` 어노테이션을 사용하면 대상 객체의 조건을 검사한다.

Controller routing method에서 `@Valid Entity entity` 형식으로 데이터를 받는다

`@Valid` 검증결과는 마찬가지로 Controller routing method에서 `BindingResult bindingResult`로 받을 수 있다.

받은 검증 결과는 다시 템플릿으로 돌려주고, 템플릿 엔진 설정에 따라 메시지를 보여주게 된다.

```java
if (bindingResult.hasErrors()) {
  return "template";
}
```

RestController 에서는 `BindingResult` 없이 `defaultMessage`와 `field`로 결과가 반환된다.

## 검증용 어노테이션

Entity에 명시

- @AssertFalse : false 값만
- @AssertTrue : true 값만
- @DecimalMax(value=) : 지정된 값 이하의 실수만
- @DecimalMin(value=) : 지정된 값 이상의 실수만
- @Digits(integer=,fraction=) : 대상 수가 지정된 정수와 소수 자리수보다 적을 경우
- @Future : 대상 날짜가 현재보다 미래일 경우만
- @Past : 대상 날짜가 현재보다 과거일 경우만
- @Max(value) : 지정된 값보다 아래일 경우만
- @Min(value) : 지정된 값보다 이상일 경우만
- @NotNull : null 값이 아닐 경우만
- @Null : null일 겨우만
- @Pattern(regex=, flag=) : 해당 정규식을 만족할 경우만
- @Size(min=, max=) : 문자열 또는 배열이 지정된 값 사이일 경우
- @Email : Email 형식
- @NotBlank : 빈칸 허용하지 않음

## Custom 어노테이션

커스텀 어노테이션 정의

```java
@Target({ ElementType.TYPE, ElementType.ANNOTATION_TYPE })
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = { FieldsMatchValidator.class })
public @interface FieldsMatcher {

  String message() default "다른 값이 입력되었습니다";

  Class<?>[] groups() default {};

  Class<? extends Payload>[] payload() default {};

  String baseField();

  String matchField();

  @Target({ElementType.TYPE, ElementType.ANNOTATION_TYPE})
  @Retention(RetentionPolicy.RUNTIME)
  @interface List {

    FieldsMatcher[] value();

  }

}
```

커스텀 validator 정의

```java
public class FieldsMatchValidator implements ConstraintValidator<FieldsMatcher, Object> {

  private String baseField;
  private String matchField;
  private String message;

  @Override
  public void initialize(FieldsMatcher constraint) {
    baseField = constraint.baseField();
    matchField = constraint.matchField();
    message = constraint.message();
  }

  @Override
  public boolean isValid(Object object, ConstraintValidatorContext context) {
    try {
      Object baseFieldValue = getFieldValue(object, baseField);
      Object matchFieldValue = getFieldValue(object, matchField);
      if (baseFieldValue != null && baseFieldValue.equals(matchFieldValue)) {
        return true;
      }
      context.buildConstraintViolationWithTemplate(message)
          .addPropertyNode(baseField)
          .addConstraintViolation()
          .disableDefaultConstraintViolation();
      return false;

    } catch (Exception e) {
      return false;
    }
  }

  private Object getFieldValue(Object object, String fieldName) throws Exception {
    Class<?> clazz = object.getClass();
    Field field = clazz.getDeclaredField(fieldName);
    field.setAccessible(true);
    return field.get(object);
  }

}
```

## Thymeleaf 적용

### Controller 예제: 해당페이지 관련 `GetMapping`, `PostMapping`에서 모두 커맨드 객체를 생성해야함

```java
@Controller
public class WebController implements WebMvcConfigurer {

  @Override
  public void addViewControllers(ViewControllerRegistry registry) {
    registry.addViewController("/results").setViewName("results");
  }

  @GetMapping("/")
  public String showForm(PersonForm personForm) {
    return "form";
  }

  @PostMapping("/")
  public String checkPersonInfo(@Valid PersonForm personForm, BindingResult bindingResult) {
    if (bindingResult.hasErrors()) {
      return "form";
    }
    return "redirect:/results";
  }
}
```

### Thymeleaf 예제: form 태그, input 태그, 에러 메세지를 위한 태그 모두 수정해야 함

```html
<html>
<body>
  <form action="#" th:action="@{/}" th:object="${personForm}" method="post">
    <table>
      <tr>
        <td>Name:</td>
        <td><input type="text" th:field="*{name}" /></td>
        <td th:if="${#fields.hasErrors('name')}" th:errors="*{name}">Name Error</td>
      </tr>
      <tr>
        <td>Age:</td>
        <td><input type="text" th:field="*{age}" /></td>
        <td th:if="${#fields.hasErrors('age')}" th:errors="*{age}">Age Error</td>
      </tr>
      <tr>
        <td><button type="submit">Submit</button></td>
      </tr>
    </table>
  </form>
</body>
</html>
```

Form에서 명시하지 않은 필드에 validation 어노테이션 사용시

```html
<ul th:if="${#fields.hasErrors('global')}">
  <li th:each="err : ${#fields.errors('global')}" th:text="${err}">Input is incorrect</li>
</ul>
```

```html
<p th:if="${#fields.hasErrors('global')}" th:errors="*{global}">Incorrect date</p>
```
