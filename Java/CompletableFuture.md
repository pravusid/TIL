# CompletableFuture

자바 1.8에서 도입된 CompletableFuture는 자바 1.5 Future API의 확장이다

## Future API의 한계점

1. 임의로 종료할 수 없다
2. Future는 자신의 상태를 알리지 않으며, 유효한 결과값을 반환하는 `get()` 메소드를 호출할 시 blocking이 발생한다
3. 여러개의 Future를 Chain하여 함께 사용할 수 없다
4. 여러개의 Future를 합성할 수 없다 (병렬로 실행하고 결과값을 취합하는 등의...)
5. 예외 처리기능이 없다

## CompletableFuture 실행

### `runAsync()` 메소드로 비동기 실행

반환값을 받을 필요가 없는 `Runnable` 인터페이스를 구현한 작업을 비동기로 실행한다.

```java
CompletableFuture<Void> future = CompletableFuture.runAsync(() -> {
    // Simulate a long-running Job
    try {
        TimeUnit.SECONDS.sleep(1);
    } catch (InterruptedException e) {
        throw new IllegalStateException(e);
    }
    System.out.println("작업종료");
});

// future가 완료될 때 까지 block 상태로 대기한다
System.out.println(future.get());
```

### `supplyAsync()`

반환값이 없는 경우 `runAsync()` 메소드를 사용했다.

만약 반환되는 값이 있다면 `CompletableFuture.supplyAsync()` 메소드를 사용하며,
함수형 인터페이스 `Supplier<T>`를 파라미터로 받아 `CompletableFuture<T>`를 반환한다(체이닝 가능).

```java
CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
    try {
        TimeUnit.SECONDS.sleep(1);
    } catch (InterruptedException e) {
        throw new IllegalStateException(e);
    }
    return "작업종료";
});
```

### Executor

`CompletableFuture`는 작업들을 글로벌 `ForkJoinPool.commonPool()`에서 얻은 쓰레드에서 실행한다.
만약 parallelism을 지원하지 않는다면 `ForkJoinPool` 대신 `ThreadPerTaskExecutor`를 사용한다.

`Executor`를 지정하는 메소드를 통해 사용자 정의 Thread pool을 활용할 수 있다.

```java
static CompletableFuture<Void>  runAsync(Runnable runnable)
static CompletableFuture<Void>  runAsync(Runnable runnable, Executor executor)
static <U> CompletableFuture<U> supplyAsync(Supplier<U> supplier)
static <U> CompletableFuture<U> supplyAsync(Supplier<U> supplier, Executor executor)

Executor executor = Executors.newFixedThreadPool(10);
CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
    try {
        TimeUnit.SECONDS.sleep(1);
    } catch (InterruptedException e) {
        throw new IllegalStateException(e);
    }
    return "작업종료";
}, executor);
```

## CompletableFuture 반환값 처리

반환값을 감싼 `CompletableFuture`를 처리하는 방법은 크게 다음과 같다

- `Apply`: 인자O, 반환O
- `Run`: 인자X, 반환X
- `Accept`: 인자O, 반환X

API에서는 여기에 추가적인 옵션을 붙인 메소드를 제공한다.

- `Either`:다른 두 CompletableFuture 중 먼저 하나가 완료되면 처리됨
- `Both`: 다른 CompletableFuture와 체인에서 반환된 CompletableFuture를 모두 기다림 (다른 타입의 반환값 수용)
- `Async` : 추가적인 쓰레드에서 비동기로 실행 (executor 지정가능)

### `thenApply()`

함수형 인터페이스 `Function<T,R>`를 인자로 받는 `thenApply()` 메소드를 이용해서 반환된 CompletableFuture 결과값을 처리할 수 있다.

```java
CompletableFuture<String> whatsYourNameFuture = CompletableFuture.supplyAsync(() -> {
   try {
       TimeUnit.SECONDS.sleep(1);
   } catch (InterruptedException e) {
       throw new IllegalStateException(e);
   }
   return "world";
});

// thenApply()를 통해 Future에 callback을 정의한다
CompletableFuture<String> greetingFuture = whatsYourNameFuture.thenApply(str -> {
   return "hello " + str;
});
```

`thenApply()` 콜백 메소드를 연속으로 사용해서 연속한 값 처리를 할 수 있다

```java
CompletableFuture<String> welcomeText = CompletableFuture.supplyAsync(() -> {
    try {
        TimeUnit.SECONDS.sleep(1);
    } catch (InterruptedException e) {
       throw new IllegalStateException(e);
    }
    return "world";
}).thenApply(str -> {
    return "hello " + str;
}).thenApply(str -> {
    return str + "!";
});
```

### `thenRun()` & `thenAccept()` => 실행 후 반환값이 없음

콜백 메소드에서 값 처리후 반환값이 없거나(`thenRun()`),
반환값을 받아서(`thenAccept()`) 연속적인 작업을 실행하려 할 때

```java
CompletableFuture.supplyAsync(() -> {
    // 작업실행
}).thenRun(() -> {
    // 작업 후 처리
});

CompletableFuture.supplyAsync(() -> {
    return "jobs done";
}).thenAccept(str -> {
    System.out.println("status:" + str)
});
```

## CompletableFutures 결합

### `thenCompose()`를 사용한 의존관계의 Future 결합

전에 Async 프로세스로 응답 받은 값을 다음 Async 프로세스의 인자로 사용하는 경우 (합성)

회원 정보를 얻는 `getUsersDetail` 서비스와 회원의 신용점수를 얻는 별도의 서비스인 `getCreditRating`을 가정해보자

```java
CompletableFuture<User> getUsersDetail(String userId) {
  return CompletableFuture.supplyAsync(() -> UserService.getUserDetails(userId));
}

CompletableFuture<Double> getCreditRating(User user) {
  return CompletableFuture.supplyAsync(() -> CreditRatingService.getCreditRating(user));
}
```

`thenApply()`를 사용한다면 결과를 얻기 위해서 다음과 같이 구성해야 할 것이다

```java
CompletableFuture<CompletableFuture<Double>> result = getUserDetail(userId)
        .thenApply(user -> getCreditRating(user));
```

하지만 `thenApply()`의 반환값은 `CompletableFuture`로 wrapping 되어 있다.
이런경우 `thenCompose()` 사용하여 연속한 `CompletableFuture`를 사용할 수 있다.

```java
CompletableFuture<Double> result = getUserDetail(userId)
        .thenCompose(user -> getCreditRating(user));
```

### `thenCombine()`을 사용한 독립적인 Future 결합

`thenCompose()`는 하나의 Future가 다른 Future에 종속적일 때 사용했다. (합성)
`thenCombine()`은 독립적인 두 future 종료 후(병렬실행) 실행할 무언가가 있을 때 사용한다.

```java
System.out.println("Retrieving weight.");
CompletableFuture<Double> weightInKgFuture =
        CompletableFuture.supplyAsync(() -> {
            try {
                TimeUnit.SECONDS.sleep(1);
            } catch (InterruptedException e) {
                throw new IllegalStateException(e);
            }
            return 65;
        });

System.out.println("Retrieving height.");
CompletableFuture<Double> heightInCmFuture =
        CompletableFuture.supplyAsync(() -> {
            try {
                TimeUnit.SECONDS.sleep(1);
            } catch (InterruptedException e) {
                throw new IllegalStateException(e);
            }
            return 175;
        });

System.out.println("Calculating BMI.");
CompletableFuture<Double> combinedFuture =
        weightInKgFuture.thenCombine(heightInCmFuture, (weightInKg, heightInCm) -> {
                    Double heightInMeter = heightInCm/100;
                    return weightInKg/(heightInMeter*heightInMeter);
                });
```

`thenCombine()`으로 전해진 콜백 메소드는 두 Future가 모두 끝나고 호출될 것이다.

## 다수의 CompletableFutures 결합

여러개의 `CompletableFuture`를 결합하기 위해서 다음의 메소드를 사용할 수 있다.

```java
static CompletableFuture<Void> allOf(CompletableFuture<?>... cfs)
static CompletableFuture<Object> anyOf(CompletableFuture<?>... cfs)
```

### CompletableFuture.allOf()

`CompletableFuture.allOf()`는 여러개의 독립적인 Future를 결합하여 병렬로 실행하고,
모든 Future의 수행이 끝나면 공통작업을 하기 위해 사용한다.

```java
CompletableFuture<String> downloadWebPage(String pageLink) {
    return CompletableFuture.supplyAsync(() -> {
        // Code to download and return the web page's content
    });
}
```

웹페이지를 다운로드 할때 'CompletableFuture` 키워드가 포함된 페이지 수를 알고싶다고 하자.

```java
List<String> webPageLinks = Arrays.asList(...) // A list of 100 web page links

// 비동기로 웹페이지 콘텐츠를 다운로드 함
List<CompletableFuture<String>> pageContentFutures =
        webPageLinks.stream()
                .map(webPageLink -> downloadWebPage(webPageLink))
                .collect(Collectors.toList());

// allOf()를 사용하여 Future를 결합함
CompletableFuture<Void> allFutures = CompletableFuture.allOf(
        pageContentFutures.toArray(new CompletableFuture[pageContentFutures.size()])
);
```

`CompletableFuture.allOf()` 메소드의 문제점은 `CompletableFuture<Void>`를 반환한다는 것이다.
하지만 CompletableFuture로 감싼 결과값을 `future.join()` 메소드를 통해 얻을 수 있다.

```java
CompletableFuture<List<String>> allPageContentsFuture = allFutures.thenApply(v -> {
   return pageContentFutures.stream()
           .map(pageContentFuture -> pageContentFuture.join())
           .collect(Collectors.toList());
});
```

`future.join()` 메소드는 모든 future가 완료된 후 호출 되므로 blocking이 발생하지 않는다.
`join()` 메소드는 `get()`메소드와 비슷하지만 unchecked exception을 throw 한다는 점이 다르다.

```java
// 키워드를 갖고있는 웹페이지 수를 센다
CompletableFuture<Long> countFuture = allPageContentsFuture.thenApply(pageContents -> {
    return pageContents.stream()
            .filter(pageContent -> pageContent.contains("CompletableFuture"))
            .count();
});
```

### CompletableFuture.anyOf()

`CompletableFuture.anyOf()` 메소드는 주어진 CompletableFuture 중 하나라도 끝나면, 새로운 CompletableFuture로 감싼 결과를 반환한다.

```java
CompletableFuture<String> future1 = CompletableFuture.supplyAsync(() -> {
    try {
        TimeUnit.SECONDS.sleep(2);
    } catch (InterruptedException e) {
        throw new IllegalStateException(e);
    }
    return "Result of Future 1";
});

CompletableFuture<String> future2 = CompletableFuture.supplyAsync(() -> {
    try {
        TimeUnit.SECONDS.sleep(1);
    } catch (InterruptedException e) {
        throw new IllegalStateException(e);
    }
    return "Result of Future 2";
});

CompletableFuture<String> future3 = CompletableFuture.supplyAsync(() -> {
    try {
        TimeUnit.SECONDS.sleep(3);
    } catch (InterruptedException e) {
        throw new IllegalStateException(e);
    }
    return "Result of Future 3";
});

CompletableFuture<Object> anyOfFuture = CompletableFuture.anyOf(future1, future2, future3);

System.out.println(anyOfFuture.get()); // Result of Future 2
```

`CompletableFuture.anyOf()` 메소드는 Future의 가변배열을 인자로 받는다. 해당 메소드의 문제점은 결과로 반환되는 값의 타입을 모른다는 것이다.

## CompletableFuture 예외 처리

```java
CompletableFuture.supplyAsync(() -> {
    // 예외가 발생할수 있는 코드
    return "Some result";
}).thenApply(result -> {
    return "processed result";
}).thenApply(result -> {
    return "result after further processing";
}).thenAccept(result -> {
    // do something with the final result
});
```

첫 `supplyAsync()` 에서 예외가 발생한다면 이어지는 `thenApply()` 콜백은 하나도 실행되지 않는다.

### exceptionally() 콜백을 활용한 예외 처리

`exceptionally()` 콜백은 주어진 Future에서 예외가 발생했을때 처리할 수 있는 기회를 준다.

```java
Integer age = -1;

CompletableFuture<String> maturityFuture = CompletableFuture.supplyAsync(() -> {
    if(age < 0) {
        throw new IllegalArgumentException("Age can not be negative");
    }
    if(age > 18) {
        return "Adult";
    } else {
        return "Child";
    }
}).exceptionally(ex -> {
    System.out.println("Oops! We have an exception - " + ex.getMessage());
    return "Unknown!";
});

System.out.println("Maturity : " + maturityFuture.get());
```

예외는 처리되고 나면 이어지는 콜백에 전파되지 않는다.

### generic handle() 메소드를 활용한 예외 처리

API에서는 보다 일반적으로 예외를 처리할 수 있는 `handle()` 메소드를 제공한다.
해당 메소드에서는 예외 발생여부를 확인할 수 있는 파라미터를 제공한다.

만약 예외가 발생한다면 `handle()` 메소드의 첫 번째 파라미터인 `res`의 값은 `null`이고, `ex`의 값은 `null`이 아니다.

```java
Integer age = -1;

CompletableFuture<String> maturityFuture = CompletableFuture.supplyAsync(() -> {
    if(age < 0) {
        throw new IllegalArgumentException("Age can not be negative");
    }
    if(age > 18) {
        return "Adult";
    } else {
        return "Child";
    }
}).handle((res, ex) -> {
    if(ex != null) {
        System.out.println("Oops! We have an exception - " + ex.getMessage());
        return "Unknown!";
    }
    return res;
});

System.out.println("Maturity : " + maturityFuture.get());
```
