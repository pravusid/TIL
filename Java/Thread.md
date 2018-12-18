# Thread in Java

## Java 쓰레드 구현

Java의 일반 쓰레드 구현 방식은 다음과 같다.

1. Thread 클래스를 상속한 클래스의 객체를 생성
2. Runnable 인터페이스를 구현한 클래스의 객체를 생성
3. Runnable 인터페이스를 inline or anonymous로 구현

### Thread 객체 생성

1. Thread class를 상속
2. public void run()을 구현
3. start()로 run() 메소드를 호출

### Runnable 인터페이스 구현

1. Runnable 인터페이스 구현
2. public void run()을 구현
3. Thread(Runnable) 객체를 만들어 start()로 run() 메소드를 호출

### 동기화 메소드와 동기화 블록

쓰레드가 사용중인 객체를 다른 쓰레드가 변경할 수 없도록 하려면 객체에 잠금을 걸어야 한다.

멀티 쓰레드 프로그램에서 하나의 쓰레드가 실행할 수 있는 코드 영역을 임계영역(critical section)이라고 한다.

자바는 임계 영역을 지정하기 위해 동기화 메소드와 동기화 블록을 제공한다.

```java
public synchronized void method() {
    // 단 하나의 쓰레드만 동시에 접근 가능
}
```

동기화 메소드는 전체가 임계영역이고 쓰레드가 동기화 메소드 영역을 실행하는 즉시 객체에 잠금이 일어나고,
동기화 메소드가 종료되면 잠금이 풀린다.

일부 내용만 임계영역으로 만드려면 동기화 블록을 만들면 된다.

```java
public void method() {
    // ...
    synchronized(공유객체) {
        // 임계영역
    }
    // ...
}
```

### 쓰레드 상태 제어

- `interrupt()`: 일시 정지 상태의 쓰레드에서 `InterruptedException` 예외를 발생시킨다
- `notify()`, `notifyAll()`: 동기화 블록 내에서 `wait()` 메소드에 의해 일시정지 상태에 있는 쓰레드를 실행대기상태로 만든다
- `sleep()`: 주어진 시간동안 쓰레드를 일시 정지 상태로 만든다
- `join()`: `obj.join()` obj 쓰레드가 종료되거나 인자로 주어진 시간이 지나면 실행 대기상태가 된다
- `wait()`: 동기화 블록내에서 쓰레드를 일시 정지 상태로 만든다. 주어진 시간이 지나거나 `notify()` 메소드에 의해 실행대기상태가 된다.
- `yield()`: 실행 중 우선순위가 동일한 다른 쓰레드에게 실행을 양보하고 실행 대기상태가 된다

### 쓰레드 그룹

쓰레드를 묶어서 사용하기 위해 사용된다

```java
ThreadGroup tg = new ThreadGroup([ThreadGroup parent], String name);
// 쓰레드를 생성할 때 쓰레드 그룹을 지정할 수 있다
Thread t = new Thread(ThreadGroup, Runnable target);
```

## Java Multi Threading

JDK 1.5 부터 포함된 Concurrent 패키지에서 Executor, Callable, Future 지원

### 개요

쓰레드 구동을 위해 다음 단계를 수행하면 된다

1. Task를 정의한 클래스의 생성
2. Executor Service에 Task 객체를 제공

### 특징

- 쓰레드 풀을 사용
- 무거운 쓰레드는 미리 할당 가능
- Task 와 쓰레드를 생성하고 관리하는 것을 분리
- 쓰레드 풀안의 쓰레드는 한번해 하나씩 여러 Task를 실행
- Task 큐를 이용해 Task를 관리
- Executor Service가 더이상 필요 없으면 중지
- Executor Service가 멈추면 모든 쓰레드도 중지

### 주요 클래스와 인터페이스

- Executor 인터페이스: Task와 쓰레드를 분리하고 실행을 담당
- ExecutorService 인터페이스: Executor 인터페이스를 확장하며 라이프 사이클을 제어
- Executors 클래스: 다양한 executor서비스의 인스턴스를 생성하는 Factory 클래스
- Future 인터페이스: Task가 중지되었는지 아닌지를 확인하거나 Task로부터 응답 획득
- Callable 인터페이스

### Executor를 이용한 쓰레드 구현

#### FixedThreadPool 사용

지정한 수의 Thread Pool을 사용한다

CPU코어 수만큼 최대 쓰레드를 지정하려면 인자로 `Runtime.getRuntime().availableProcessors()`를 사용한다

```java
ExecutorService execService = Executors.newFixedThreadPool(2);

execService.execute(new MyThreadTask());
execService.execute(new MyThreadTask());

execService.shutdown();
```

#### CachedThreadPool 사용

CachedThreadPool은 FixedThreadPool과 달리 Task의 숫자에 따라 쓰레드 숫자가 변한다

```java
ExecutorService execService = Executors.newCachedThreadPool();

execService.execute(new MyThreadTask());
execService.execute(new MyThreadTask());

execService.shutdown();
```

#### SingleThreadExecutor

쓰레드가 하나로 구성되어 있다. Task간 Thread safe 하다.

```java
ExecutorService execService = Executors.newSingleThreadExecutor();

execService.execute(new MyThreadTask());
execService.execute(new MyThreadTask());

execService.shutdown();
```

### Executors 종료

Executors의 쓰레드는 기본적으로 daemon thread가 아니므로 main 스레드가 종료되더라도 계속 실행 상태로 남아있다.

따라서 ExecutorService는 종료와 관련해서 세 가지의 메소드를 제공한다

- `void shutdown()`: 현재 작업과 작업큐의 대기중인 작업을 모두 처리하고 쓰레드풀 종료
- `List<Runnable> shutdownNow()`: 현재 처리중인 쓰레드에 interrupt를 걸고 미처리된 작업목록을 반환하며 쓰레드풀을 종료시킨다
- `boolean awaitTermination(long timeout, TimeUnit unit)`: timeout 시간내에 작업을 완료하면 true, 작업이 남으면 interrupt 걸고 쓰레드 풀 종료후 false 반환

### 쓰레드 이름 설정

1. 실행 시점에 이름 부여: run() 메소드안에서 Thread.currentThread().setName()을 이용한다
2. 생성 시점에 부여하기: Thread() 생성자의 인자로 이름을 입력한다.

### Executor의 쓰레드 이름 설정

Executor에서 쓰레드에 이름을 부여하려면 ThreadFactory 인터페이스를 구현해야 한다.
ThreadFactory 인터페이스는 다음의 메소드를 구현해야 한다.

`public Thread newThread(Runnable r);`

구현부의 반환값인 Thread constructor에 이름을 추가하여 반환하면 Executor가 생성한 쓰레드에 이름이 할당된다.

### 쓰레드에서 값 반환

일반 쓰레드의 경우와 Executor에서 값을 반환반는 방법을 각각 살펴보자.

#### 일반 쓰레드의 경우

일반 쓰레드에서 리턴 값을 얻는 방법은 두 가지 이다

##### 블록킹

synchronized를 이용해 변경 시점까지 락을 거는 방법으로 데이터를 읽는 메소드를 락을 걸고 쓰레드가 실행되는 마지막에 락을 푼다.
this.wait()로 기다리고 this.notifyAll()로 해제한다.

done 변수를 이용해 synchronized로 무조건 들어가지 않도록 했으며, done은 volatile로 선언해서 어떤 스레드가 값을 변경하든,
항상 최신값을 읽어갈 수 있게 해준다.

참고로 volatile은 항상 최신 값을 읽게는 해주지만 operation을 atomic하게는 만들지 않으며, synchronized는 operation을 atomic하게 해준다.
즉, volatile은 동기화를 할 뿐이지 lock은 아니다.

```java
public class ReturningValueFirstWay {
    public static void main(String args[]) throws InterruptedException {
        System.out.println("Main thread starts here...");

        MyThreadTask task1 = new MyThreadTask();
        MyThreadTask task2 = new MyThreadTask();

        new Thread(task1,"firstThread").start();
        new Thread(task2,"secondThread").start();

        System.out.println("task1 result:" + task1.getRandomSum());
        System.out.println("task2 result:" + task2.getRandomSum());

        System.out.println("Main thread ends here...");
    }
}

class MyThreadTask implements Runnable {
    private static int count = 0;
    private int id;
    private volatile boolean done = false;
    private int randomSum = 0;

    @Override
    public void run(){
        for(int i = 0; i<5; i++) {
            System.out.println("<" + id + ">TICK TICK " + i);
            randomSum += Math.random()*1000;
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        /********** Unlock ************/
        done = true;
        synchronized (this) {
            System.out.println(Thread.currentThread().getName() + ": Notifying the result");
            this.notifyAll();
        }
    }

    public int getRandomSum(){
        /*********** Lock ************/
        if(!done) {
            synchronized (this) {
                try {
                    System.out.println(Thread.currentThread().getName() + ": Waiting for the result");
                    this.wait();
                } catch (InterruptedException e){
                    e.printStackTrace();
                }
            }
            System.out.println(Thread.currentThread().getName() + ": Woken up");
        }
        return randomSum;
    }

    public MyThreadTask() {
        this.id = ++count;
    }
}
```

main 스레드가 getRandomSum()에서 wait()를 하고 있다가 스레드에서 notifyAll()을 한 시점에 풀려나서 값을 읽는다.
타이밍 상 main 스레드는 한번만 락이 걸렸다가 풀렸음에도 두 스레드 모두 done이 true가 되어 있었다.

##### 논블록킹

Observer pattern을 이용하는 것으로 스레드가 끝났을 때 등록된 Listener의 메소드를 호출하는 방식이다.

```java
public class ReturningValueSecondWay {
    public static void main(String args[]) throws InterruptedException {
        System.out.println("Main thread starts here...");

        MyThreadTask task1 = new MyThreadTask(new RandomSumObserver("task1"));    //Register the listeners
        MyThreadTask task2 = new MyThreadTask(new RandomSumObserver("task2"));    //Register the listeners

        new Thread(task1,"firstThread").start();
        new Thread(task2,"secondThread").start();

        System.out.println("Main thread ends here...");
    }
}

/** Listener interface **/
interface ResultListener<T> {
    public void notifyResult(T t);
}

/** Listener **/
class RandomSumObserver implements ResultListener<Integer> {

    private String taskId;

    public RandomSumObserver(String taskId) {
        this.taskId = taskId;
    }

    @Override
    public void notifyResult(Integer result) {
        System.out.println(taskId + " result:" + result);
    }
}

class MyThreadTask implements Runnable {
    private static int count = 0;
    private int id;
    private int randomSum = 0;
    private ResultListener<Integer> listener; //listener

    @Override
    public void run(){
        for(int i = 0; i<5; i++) {
            System.out.println("<" + id + ">TICK TICK " + i);
            randomSum += Math.random()*1000;
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        this.listener.notifyResult(randomSum); //Invoke the listener
    }

    public MyThreadTask(ResultListener<Integer> listener) {    //Register listener
        this.id = ++count;
        this.listener = listener;
    }
}
```

블록킹 없이 쓰레드의 연산이 수행되면 자동으로 등록된 Listener를 실행시킨다.

#### Executor의 경우

##### Callable 인터페이스의 사용

Executor는 Runnable 인터페이스 대신 Callable 인터페이스를 이용해 Task를 만들어서 결과를 반환한다.
callable 인터페이스에서 오버라이드해야 하는 메소드는 call()이다.

`public T call() throws Exception`

ExecutorService.submit()을 하고 Future를 반환 받는다.

`Future<T> ExecutorService.submit(Callable c)`

ExecutorService.shutdown() 이후에 Task의 결과는 Future의 get() 메소드를 이용해 반환 받는다.
값을 반환하지 않는 Runnable 객체도 동일하게 할 수 있는데 이때 반환 값은 null이다.

Runnable 객체를 submit할때 두번째 인자 값은 Future로 반환 시 동일한 값이 반환되게 된다.

```java
public class ReturningValuesUsingExecutorsFirstWay {
    public static void main(String args[]) throws InterruptedException {
        System.out.println("Main thread starts here...");

        ExecutorService execService = Executors.newCachedThreadPool();

        //Callable
        Future<Integer> result1 = execService.submit(new MyCallableTask());
        Future<Integer> result2 = execService.submit(new MyCallableTask());

        //Runnable
        Future<?> result3 = execService.submit(new MyRunnableTask());
        Future<?> result4 = execService.submit(new MyRunnableTask(), 110.1);

        execService.shutdown();

        try {
            System.out.println("result1:" + result1.get());
            System.out.println("result2:" + result2.get());
            System.out.println("result3:" + result3.get());
            System.out.println("result4:" + result4.get());
        } catch(Exception e) {
            e.printStackTrace();
        }

        System.out.println("Main thread ends here...");
    }
}

class MyCallableTask implements Callable<Integer> {
    private static int count = 0;
    private int id;
    private static String taskName = "CallableTaks";
    private int randomSum = 0;

    @Override
    public Integer call() throws Exception {
        for(int i = 0; i<5; i++) {
            System.out.println("<" + taskName +"-"+id + ">TICK TICK " + i);
            randomSum += Math.random()*1000;
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        return randomSum; //값을 반환
    }
    public MyCallableTask() {
        this.id = ++count;
    }
}

class MyRunnableTask implements Runnable {
    private static int count = 0;
    private int id;
    private static String taskName = "RunnableTaks";

    @Override
    public void run() {
        for(int i = 0; i<5; i++) {
            System.out.println("<" + taskName +"-"+id + ">TICK TICK " + i);
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
    public MyRunnableTask() {
        this.id = ++count;
    }
}
```

Runnable과 Callable이 동일하게 실행되며,
결과도 반환받는데 Runnable은 null이고, 인자를 추가한 경우 그 인자값이 동일하게 반환된다.

##### CompletionService의 사용

CompletionService 인터페이스를 이용해 Task의 리턴을 획득할 수 있다.

- Future submit(Callable task): Callable 등록
- Future submit(Runnable task, V result); Runnable 등록
- Future take(): 종료된 Task의 인스턴스. 종료가 될때까지 블록되고 종료가 된 Task가 있으면 그 Task의 인스턴스를 반환
- Future poll(): take() 메소드와 동일하지만 블록되지않음. 만약 종료된 Task가 없으면 null을 반환

`Future submit(Runnable task)`가 없어서 결과값을 반환하지 않는 Task는 처리하지 않는다

```java
public class ReturingValuesUsingExecutorsSecondWay {
    public static void main(String args[]) throws InterruptedException {
        System.out.println("Main thread starts here...");

        ExecutorService execService = Executors.newCachedThreadPool();

        CompletionService<Integer> tasks = new ExecutorCompletionService<>(execService);

        //Callable
        tasks.submit(new MyCallableTask());
        tasks.submit(new MyCallableTask());

        //Runnable
        tasks.submit(new MyRunnableTask(), 110);

        execService.shutdown();

        for(int i = 0; i< 3; i++) {
            try {
                System.out.println("Result: " + tasks.take().get()); //block
            } catch (ExecutionException e) {
                e.printStackTrace();
            }
        }

        System.out.println("Main thread ends here...");
    }
}

class MyCallableTask implements Callable<Integer> {
    private static int count = 0;
    private int id;
    private static String taskName = "CallableTaks";
    private int randomSum = 0;

    @Override
    public Integer call() throws Exception {
        for(int i = 0; i<5; i++) {
            System.out.println("<" + taskName +"-"+id + ">TICK TICK " + i);
            randomSum += Math.random()*1000;
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        return randomSum; //값을 반환
    }
    public MyCallableTask() {
        this.id = ++count;
    }
}

class MyRunnableTask implements Runnable {
    private static int count = 0;
    private int id;
    private static String taskName = "RunnableTaks";

    @Override
    public void run() {
        for(int i = 0; i<5; i++) {
            System.out.println("<" + taskName +"-"+id + ">TICK TICK " + i);
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
    public MyRunnableTask() {
        this.id = ++count;
    }
}
```

Generic을 이용해 다음과 같은 반환 값을 위한 클래스를 정의한 경우의 예제이다

```java
//TaskResult.java
public class TaskResult<S, R> {
    private S taskId;
    private R result;
    public TaskResult(S taskId, R result) {
        this.taskId = taskId;
        this.result = result;
    }
    public S getTaskId() {
        return taskId;
    }
    public R getResult() {
        return result;
    }
    @Override
    public String toString() {
        return "TaskResult [taskId=" + taskId + ", result=" + result + "]";
    }
}

public class ReturingValuesUsingExecutorsThirdWay {
    public static void main(String args[]) throws InterruptedException {
        System.out.println("Main thread starts here...");

        ExecutorService execService = Executors.newCachedThreadPool();

        CompletionService<TaskResult<String,Integer>> tasks = new ExecutorCompletionService<>(execService); //type change

        tasks.submit(new MyCallableTask());
        tasks.submit(new MyCallableTask());
        tasks.submit(new MyRunnableTask(), new TaskResult<String, Integer>("RunnableTask", 101));//type change

        execService.shutdown();

        for(int i = 0; i< 3; i++) {
            try {
                System.out.println(tasks.take().get());
            } catch (ExecutionException e) {
                e.printStackTrace();
            }
        }

        System.out.println("Main thread ends here...");
    }
}

class MyCallableTask implements Callable<TaskResult<String, Integer>> {//type change
    private static int count = 0;
    private int id;
    private static String taskName = "CallableTaks";
    private int randomSum = 0;

    @Override
    public TaskResult<String, Integer> call() throws Exception {//type change
        for(int i = 0; i<5; i++) {
            System.out.println("<" + taskName +"-"+id + ">TICK TICK " + i);
            randomSum += Math.random()*1000;
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        return new TaskResult<String, Integer>(taskName + id, randomSum); // type change
    }
    public MyCallableTask() {
        this.id = ++count;
    }
}

class MyRunnableTask implements Runnable {
    private static int count = 0;
    private int id;
    private static String taskName = "RunnableTaks";

    @Override
    public void run() {
        for(int i = 0; i<5; i++) {
            System.out.println("<" + taskName +"-"+id + ">TICK TICK " + i);
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
    public MyRunnableTask() {
        this.id = ++count;
    }
}
```

## Daemon Thread

데몬 쓰레드
데몬쓰레드는 다른 User 스레드의 작업을 돕는 보조적인 역할을 수행하도록 되어 있다.
다른 User 쓰레드가 모두 종료되면 데몬 쓰레드는 강제적으로 종료된다.

### 일반 쓰레드와 Daemon Thread

setDaemon(true)후에 start()를 시키면 된다.

### Executor와 Daemon Thread

Executor의 경우도 ThreadFactory 인터페이스를 통해 구현시 setDaemon()을 동일하게 사용한다.

```java
public class DaemonThreadsUsingExecutors {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        ExecutorService execService = Executors.newCachedThreadPool( new DaemoneThreadsFactory());

        execService.execute(new MyRunnableTask(100));
        execService.execute(new MyRunnableTask(400));

        execService.shutdown();

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class NamedThreadsFactory implements ThreadFactory {

    private static int count = 0;
    private static String Name = "MyThread-";

    @Override
    public Thread newThread(Runnable r) {
        return new Thread(r, Name + ++count);
    }
}

class DaemoneThreadsFactory extends NamedThreadsFactory {

    private static int count = 0;

    @Override
    public Thread newThread(Runnable r) {
        Thread t = super.newThread(r);
        count++;
        if(count%2 == 0){
            t.setDaemon(true); //setDaemon()
        }
        return t;
    }
}
```

## 쓰레드 동작의 완료를 확인하기

- 일반 쓰레드에서 쓰레드가 살아있는지 확인하기
- Executor에서 Task가 종료되었는지 확인하기

### 일반 쓰레드에서 쓰레드가 살아있는지 확인하기

쓰레드가 살이있는지 확인하는 방법은 Thread.isAlive() 메소드를 사용하는 것이다

### Executor에서 Task가 종료되었는지 확인하기

예외가 발생한 것도 포함해서 종료되었는지 확인하는 것은 Future.isDone() 메소드이다.

## 쓰레드 중지하기

- 일반 쓰레드: 쓰레드를 종료한다.
- Executor: Task를 종료한다. 모든 Task가 종료되면 쓰레드는 자동으로 종료한다.

### 일반 쓰레드에서 쓰레드 중지 시키기

- 플래그 사용하여 쓰레드 중지
- Non-blocking Task의 경우 인터럽트 발생 여부 확인하여 쓰레드 중지
- Blocking Task의 경우 인터럽트를 catch하여 쓰레드 중지

#### 플래그 사용하여 쓰레드 중지

Loop을 돌면서 종료 플래그를 확인하여 true이면 종료하는 방법이다.
쓰레드간의 동기화를 위해 volatile 키워드를 이용해 플래그 변수를 선언하고 synchronized를 이용해 atomic하게 만든다.

```java
public class TerminatingNormalThreadWithFlag {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        LoopTask task1 = new LoopTask();
        LoopTask task2 = new LoopTask();

        new Thread(task1, "Thread-1").start();
        new Thread(task2, "Thread-2").start();

        TimeUnit.MILLISECONDS.sleep(1000);
        task1.cancel();
        task2.cancel();

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class LoopTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;
    private volatile boolean shutdown = false;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            try {
                TimeUnit.MILLISECONDS.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            synchronized(this) {
                if(shutdown) {
                    break;
                }
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public void cancel() {
        System.out.println("... <" + Thread.currentThread().getName() +"," + taskId + "> shutting down...");
        synchronized (this) {
            this.shutdown = true;
        }
    }

    public LoopTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
}
```

#### Non-blocking Task의 경우 인터럽트 발생 여부 확인하여 쓰레드 중지

블록되지 않고 동작하는 쓰레드를 중지시키는 좋은 방법은 인터럽트를 사용하여 종료하는 것이다.

- void Thread.interrupt(): 쓰레드에 인터럽트를 호출
- static boolean Thread.interrupted() : 쓰레드 내부에서 쓰레드가 인터럽트 되었는지 확인하기 위해 사용
- boolean Thread.isInterrupted(): 다른 쓰레드에서 호출하여 쓰레드가 인터럽트 되었는지 확인

주의할 점은 blocking 메소드(타이머 등)와 함께 사용하면 Exception이 발생한다는 것이다.

```java
public class TerminatingNormalNonBlockingThreadWithInterrupt {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        LoopTask task1 = new LoopTask();
        LoopTask task2 = new LoopTask();

        Thread thread1 = new Thread(task1, "Thread-1");
        Thread thread2 = new Thread(task2, "Thread-2");
        thread1.start();
        thread2.start();

        TimeUnit.MILLISECONDS.sleep(500);

        thread1.interrupt();
        thread2.interrupt();

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class LoopTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;
    private final int DATA_SIZE = 100000;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            doSomeWork(); //Interrupt can not work with timer because the timer is a blocking function
            if(Thread.interrupted()){ //Thread.interrupted()
                break;
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public LoopTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
    private void doSomeWork(){
        for(int i = 0; i<2 ; i++) {
            Collections.sort(generateDataSet());
        }
    }
    private List<Integer> generateDataSet(){
        List<Integer> intList = new ArrayList<>();
        Random randomGenerator = new Random();
        for(int i = 0; i<DATA_SIZE; i++) {
            intList.add(randomGenerator.nextInt(DATA_SIZE));
        }
        return intList;
    }
}
```

#### Blocking Task의 경우 인터럽트를 catch하여 쓰레드 중지

인터럽트를 외부 쓰레드에서 발생시키면 Sleep과 같은 Blocking 메소드에서 Interrupted Exception이 발생하는데
이때 그 Exception catch 블록에서 loop을 중지시키는 것이다.

```java
public class TerminatingNormalBlockingThreadWithInterrupt {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        LoopTask task1 = new LoopTask();
        LoopTask task2 = new LoopTask();

        Thread thread1 = new Thread(task1, "Thread-1");
        Thread thread2 = new Thread(task2, "Thread-2");
        thread1.start();
        thread2.start();

        TimeUnit.MILLISECONDS.sleep(500);

        thread1.interrupt();
        thread2.interrupt();

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class LoopTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            try {
                TimeUnit.MILLISECONDS.sleep(100);
            } catch (InterruptedException e) {
                System.out.println("<" + currentThreadName + "," + taskId + "> Sleep Interrupted. Cancelling...");
                break;
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public LoopTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
}
```

### Excutor에서 쓰레드 중지 시키기

Executor에서는 쓰레드를 중지하는 개념이 아니라 Task를 중지시킨다.

- 플래그 사용하여 Task 중지
- Non-blocking Task의 경우 인터럽트 발생 여부 확인하여 중지
- Blocking Task의 경우 인터럽트를 catch하여 중지

#### 플래그 사용하여 Task 중지

일반적인 쓰레드에서 loop을 도는 Task를 종료하는 대표적인 방법은 플래그를 사용하는 것이다.
Loop을 돌면서 종료 플래그를 확인하여 true이면 종료하는 방법이다.
쓰레드간의 동기화를 위해 volatile 키워드를 이용해 플래그 변수를 선언하고 synchronized를 이용해 atomic하게 만든다.

```java
public class TerminatingExecutorWithFlag {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        ExecutorService execService = Executors.newCachedThreadPool();

        LoopTask task1 = new LoopTask();
        execService.execute(task1);

        execService.shutdown();

        TimeUnit.MILLISECONDS.sleep(1000);

        task1.cancel();

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class LoopTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;
    private volatile boolean shutdown = false;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            try {
                TimeUnit.MILLISECONDS.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            synchronized(this) {
                if(shutdown) {
                    break;
                }
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public void cancel() {
        System.out.println("... <" + Thread.currentThread().getName() +"," + taskId + "> shutting down...");
        synchronized (this) {
            this.shutdown = true;
        }
    }

    public LoopTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
}
```

#### Non-blocking Task의 경우 인터럽트 발생 여부 확인하여 중지

Executor에서 블록되지 않고 동작하는 Task를 중지시키는 좋은 방법은 인터럽트를 사용하여 종료하는 것이다.

- boolean Future.cancel(boolean mayInterruptIfRunning): Executor의 Task 멈추기. 만약 Task가 이미 다른 이유로 cancel되었다면 false를 반환. 일반 쓰레드에서의 Thread.interrupt()와 유사한 역할
- static boolean Thread.interrupted() : Task가 인터럽트 되었는지 확인하기 위해 사용. 일반 쓰레드와 동일
- boolean Future.isCanceled(): 다른 쓰레드에서 호출. boolean Thread.isInterrupted()와 유사

중요한 것은 위의 API는 쓰레드가 아닌 Task를 중지시키는 것이라는 것이다.
쓰레드를 점유한 첫번째 Task를 종료시키면 두번째 Task가 동작하고 두번째 Task를 종료시키면 더이상 Task가 없으므로 쓰레드가 멈추게 된다.

```java
public class TerminatingNonBlockingExecutorWithInterrupt {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        ExecutorService execService = Executors.newSingleThreadExecutor();

        LoopTask task1 = new LoopTask();
        LoopTask task2 = new LoopTask();

        Future<?> future1 = execService.submit(task1);
        Future<?> future2 = execService.submit(task2);

        execService.shutdown();

        TimeUnit.MILLISECONDS.sleep(200);
        future1.cancel(true);
        TimeUnit.MILLISECONDS.sleep(100);
        future2.cancel(true);

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class LoopTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;
    private final int DATA_SIZE = 100000;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            doSomeWork();
            if(Thread.interrupted()){
                System.out.println("<" + currentThreadName + "," + taskId + "> got an interrupt! ..canceling...");
                break;
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public LoopTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
    private void doSomeWork(){
        for(int i = 0; i<2 ; i++) {
            Collections.sort(generateDataSet());
        }
    }
    private List<Integer> generateDataSet(){
        List<Integer> intList = new ArrayList<>();
        Random randomGenerator = new Random();
        for(int i = 0; i<DATA_SIZE; i++) {
            intList.add(randomGenerator.nextInt(DATA_SIZE));
        }
        return intList;
    }
}
```

결과는 아래와 같이 첫번째 Task가 멈춘 후 같은 쓰레드에서 두번째 쓰레드가 동작한다.

#### Blocking Task의 경우 인터럽트를 catch하여 중지

외부 쓰레드에서 cancel()을 호출하면 Sleep과 같은 Blocking함수에서 Interrupted Exception이 발생하는데
그 Exception catch 블록에서 loop을 중지시킨다.

```java
public class TerminatingBlockingExecutorWithInterrupt {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        ExecutorService execService = Executors.newSingleThreadExecutor();

        LoopTask task1 = new LoopTask();
        LoopTask task2 = new LoopTask();

        Future<?> future1 = execService.submit(task1);
        Future<?> future2 = execService.submit(task2);

        execService.shutdown();

        TimeUnit.MILLISECONDS.sleep(200);
        future1.cancel(true);
        TimeUnit.MILLISECONDS.sleep(100);
        future2.cancel(true);

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class LoopTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            try {
                TimeUnit.MILLISECONDS.sleep(100);
            } catch (InterruptedException e) {
                System.out.println("<" + currentThreadName + "," + taskId + "> Sleep Interrupted. Cancelling...");
                break;
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public LoopTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
}
```

#### 모든 Executor를 한번에 중지시키기

ExecutorService는 모든 Task를 한번에 중지시키는 shutdownNow() 메소드를 제공하고 있다.

`List ExecutorService.shutdownNow()`

- 내부적으로는 위에서 살펴본 인터럽트를 사용하는 방식이다.
- 반환되는 List는 아직 종료되지 않은 Task들의 목록이다.
- 만약 모든 Task들이 종료가 되었다면 ExecutorService도 종료한다.
- Non-blocking Task는 물론이고 Blocking Task 또한 종료 시키는데 위에서 살펴본 인터럽트를 사용하는 방식이므로 Exception을 정확히 다루어야 한다.
- 만약 아직 실행이 시작되지 않은 Task가 있다면 종료된 것으로 처리한다.

shutdownNow()를 호출후에 Task가 완전히 종료 될 때까지 기다리기 위해서는 awaitTermiation() 메소드를 사용한다.

`boolean ExecService.awaitTermination(long timeout, TimeUnit unit)`

- shutdown 이후 일정시간동안 block된다. timeout과 unit을 이용해 blocking되는 시간을 지정한다.
- shutdown 이후 모든 Task가 종료될때 까지 block된다.
- 외부에서 실행되는 쓰레드에 인터럽트를 걸때까지 block된다.
- 위의 3가지 경우 중 가장 빨리 조건에 해제되는 시점까지 block되며 이때 모든 Task가 종료되었으면 true를 반환한다.

Blocking, Non-Blocking, Callable Task 모두를 한번에 종료하는 예제이다.
non-Blocking은 인터럽트 여부를 확인해야 하고, Blocking은 InterruptedException을 catch해야 한다.

```java
public class TerminatingAllExecutors {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        ExecutorService execService = Executors.newCachedThreadPool();

        BlockingTask blockingTask = new BlockingTask();
        NonBlockingTask nonBlockingTask = new NonBlockingTask();
        CallableTask callableTask = new CallableTask();

        //Runnable
        execService.execute(blockingTask);
        execService.execute(nonBlockingTask);
        //Callable
        execService.submit(callableTask);

        TimeUnit.MILLISECONDS.sleep(1000);

        execService.shutdownNow(); //shutDownNow
        System.out.println("["+ currentThreadName + "]" + " shutdownNow() invoked ");

        System.out.println("["+ currentThreadName + "]" + " All threads terminated: " +
                execService.awaitTermination(500, TimeUnit.MILLISECONDS)); //awaitTermination

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class BlockingTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            try {
                TimeUnit.MILLISECONDS.sleep(100);
            } catch (InterruptedException e) {
                System.out.println("<" + currentThreadName + "," + taskId + "> Sleep Interrupted. Cancelling...");
                break;
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public BlockingTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
}

class NonBlockingTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;
    private final int DATA_SIZE = 100000;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            doSomeWork();
            if(Thread.interrupted()){
                System.out.println("<" + currentThreadName + "," + taskId + "> Thread.interrupted() is true: Cancelling...");
                break;
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public NonBlockingTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
    private void doSomeWork(){
        for(int i = 0; i<2 ; i++) {
            Collections.sort(generateDataSet());
        }
    }
    private List<Integer> generateDataSet(){
        List<Integer> intList = new ArrayList<>();
        Random randomGenerator = new Random();
        for(int i = 0; i<DATA_SIZE; i++) {
            intList.add(randomGenerator.nextInt(DATA_SIZE));
        }
        return intList;
    }
}

class CallableTask implements Callable<Integer> {
    private static int count = 0;
    private int id;
    private String taskId;
    private int randomSum = 0;

    @Override
    public Integer call() throws Exception {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        while(true) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            randomSum += Math.random()*1000;
            try {
                TimeUnit.MILLISECONDS.sleep(100);
            } catch (InterruptedException e) {
                System.out.println("<" + currentThreadName + "," + taskId + "> Sleep Interrupted. Cancelling...");
                break;
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
        return randomSum;
    }
    public CallableTask() {
        this.id = ++count;
        this.taskId = "Task-" + id;
    }
}
```

## 쓰레드 Exception 처리하기

### 일반 쓰레드에서 Exception 처리

#### try/catch 블록을 이용한 직접 Exception을 처리하기

try/catch 블록을 이용해서 Task에서 던진 RuntimeException을 처리해 보자

이를 통해 쓰레드가 던진 Exception을 쓰레드 바깥에서 try/catch로 바로 받을 수 없다는 것을 알 수 있다.

```java
public class HandlingUncaughtExceptionsForThreads {
    public static void main(String args[]) {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        try {
            new Thread(new ExceptionLeakingTask(), "Mythread-1").start();
        } catch(RuntimeException re) {
            System.out.println("!!!!!!["+ currentThreadName + "]" + " Caught Exception!!!");
        }

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class ExceptionLeakingTask implements Runnable {

    @Override
    public void run() {
        throw new RuntimeException();
    }

}
```

#### UncaughtExceptionHandler 인터페이스로 처리하기

UncaughtException을 처리하기 위해 UncaughtExceptionHandler 인터페이스를 구현하자.
UncaughtExceptionHandler 인터페이스는 다음의 메소드 하나만을 가진다.

`void uncaughtException(Thread thread, Throwable e)`

UncaughtExceptionHandler를 이용하는 방법은 3가지가 있다.

##### 시스템 내의 모든 쓰레드를 위한 디폴트 handler 지정하기

시스템 내의 모든 쓰레드를 위한 디폴트 handler로 지정하여 UncaughtException을 처리한다.
쓰레드 하나가 아닌 모든 쓰레드의 Exception을 처리한다.

이를 위해 아래 메소드를 이용해 쓰레드를 시작하기 이전에 디폴트로 지정한다. 이 경우 try/catch는 필요없다.

`static void Thread.setDefaultUncaughtExceptionHandler(UncaughtExceptionHandler eh)`

```java
public class HandlingUncaughtExceptionsForThreads {
    public static void main(String args[]) {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        Thread.setDefaultUncaughtExceptionHandler(new ThreadExceptionHandler("DefaultHandler"));

        Thread thread1 = new Thread(new ExceptionLeakingTask(), "Mythread-1");
        Thread thread2 = new Thread(new ExceptionLeakingTask(), "Mythread-2");

        thread1.start();
        thread2.start();

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class ExceptionLeakingTask implements Runnable {

    @Override
    public void run() {
        throw new RuntimeException();
    }

}

class ThreadExceptionHandler implements UncaughtExceptionHandler{
    private String handlerName;

    public ThreadExceptionHandler(String handlerName) {
        this.handlerName = handlerName;
    }

    @Override
    public void uncaughtException(Thread thread, Throwable e) {
        System.out.println(handlerName + " caught Exception in Thread - "
            + thread.getName()
            + " => " + e);
    }

}
```

##### 쓰레드 별로 handler를 각기 지정하기

각 쓰레드 별로 각기 다른 Exception Handler를 지정하는 방법은 아래 메소드를 이용하는 것이다.

`void Thread.setUncaughtExceptionHandler(UncaughtExceptionHandler eh)`

각기 다른 쓰레드마다 별도의 handler를 등록하여 처리할 수 있다

```java
public class HandlingUncaughtExceptionsForThreads {
    public static void main(String args[]) {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        //Thread.setDefaultUncaughtExceptionHandler(new ThreadExceptionHandler("DefaultHandler"));

        Thread thread1 = new Thread(new ExceptionLeakingTask(), "Mythread-1");
        thread1.setUncaughtExceptionHandler(new ThreadExceptionHandler("Handler-1"));

        Thread thread2 = new Thread(new ExceptionLeakingTask(), "Mythread-2");
        thread2.setUncaughtExceptionHandler(new ThreadExceptionHandler("Handler-2"));

        thread1.start();
        thread2.start();

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class ExceptionLeakingTask implements Runnable {

    @Override
    public void run() {
        throw new RuntimeException();
    }

}

class ThreadExceptionHandler implements UncaughtExceptionHandler{
    private String handlerName;

    public ThreadExceptionHandler(String handlerName) {
        this.handlerName = handlerName;
    }

    @Override
    public void uncaughtException(Thread thread, Throwable e) {
        System.out.println(handlerName + " caught Exception in Thread - "
            + thread.getName()
            + " => " + e);
    }

}
```

##### 디폴트 handler와 쓰레드 별 handler를 조합해서 사용하기

위에서 살펴본 디폴트 handler와 쓰레드 별로 각기 다른 handler를 등록하는 것을 조합해서 사용해 보자. 이것은 다음과 같다.

- 디폴트 UncaughtException handler를 등록
- 디폴트 handler를 오버라이드해서 쓰레드별 handler를 등록하기

아래 예제에서는 thread1은 DefaultHandler를 이용하고, thread2는 Handler-2를 이용한다.

```java
public class HandlingUncaughtExceptionsForThreads {
    public static void main(String args[]) {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        Thread.setDefaultUncaughtExceptionHandler(new ThreadExceptionHandler("DefaultHandler"));

        Thread thread1 = new Thread(new ExceptionLeakingTask(), "Mythread-1");
        //thread1.setUncaughtExceptionHandler(new ThreadExceptionHandler("Handler-1"));

        Thread thread2 = new Thread(new ExceptionLeakingTask(), "Mythread-2");
        thread2.setUncaughtExceptionHandler(new ThreadExceptionHandler("Handler-2"));

        thread1.start();
        thread2.start();

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class ExceptionLeakingTask implements Runnable {

    @Override
    public void run() {
        throw new RuntimeException();
    }

}

class ThreadExceptionHandler implements UncaughtExceptionHandler{
    private String handlerName;

    public ThreadExceptionHandler(String handlerName) {
        this.handlerName = handlerName;
    }

    @Override
    public void uncaughtException(Thread thread, Throwable e) {
        System.out.println(handlerName + " caught Exception in Thread - "
            + thread.getName()
            + " => " + e);
    }

}
```

### Executor에서의 UncaughtException 처리하기

일반 쓰레드의 경우와 동일하게 UncaughtExceptionHandler 인터페이스를 구현한다.
UncaughtExceptionHandler 인터페이스는 다음의 메소드 하나만을 가진다.

`void uncaughtException(Thread thread, Throwable e)`

UncaughtExceptionHandler를 이용하는 방법은 3가지가 있다.

- 시스템 내의 모든 쓰레드를 위한 디폴트 handler 지정하기
- 쓰레드 별로 handler를 각기 지정하기
- 디폴트 handler와 쓰레드 별 handler를 조합해서 사용하기

## 쓰레드 Join하기

### 일반 쓰레드에서 Join

#### 일반 쓰레드가 다른 쓰레드가 끝날때까지 대기하게 하기

쓰레드가 다른 쓰레드가 종료될 때까지 대기하는 방법

- 쓰레드 동작의 완료를 확인하기에서 다룬 Thread.isAlive()를 사용하여 계속 확인
- 쓰레드에서 값 반환에서 다룬 synchronized에서 wait()와 notify()를 사용하는 블록킹 방식

첫번째 폴링하는 방법은 CPU의 낭비가 심하고 두번째 블록킹 방법은 제대로 구현하기 어렵다.

#### join()을 사용해서 다른 쓰레드 종료를 기다리기

Thread.join() 메소드를 사용하여 다른 쓰레드 종료를 기다릴 수 있다.

`void Thread.join()`

thread1, 2, 3가 실행이 되는데 thread3이 가장 먼저 끝나고, thread1이 다음, 마지막으로 thread2가 종료되도록 되어 있다.
하지만, join()을 호출하는 순서가 thread1, 2, 3이므로 main thread는 thread1을 join한 뒤에야 thread2를 join하고 마지막으로 thread3를 join한다.

```java
public class JoiningThreads {
    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        Thread thread1 = new Thread(new LoopTask(200), "Thread-1");
        Thread thread2 = new Thread(new LoopTask(500), "Thread-2");
        Thread thread3 = new Thread(new LoopTask(100), "Thread-3");

        thread1.start();
        thread2.start();
        thread3.start();

        thread1.join();
        System.out.println("["+ currentThreadName + "]" + " joined " + thread1.getName());

        thread2.join();
        System.out.println("["+ currentThreadName + "]" + " joined " + thread2.getName());

        thread3.join();
        System.out.println("["+ currentThreadName + "]" + " joined " + thread3.getName());

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class LoopTask implements Runnable {
    private static int count = 0;
    private int id;
    private long sleepTime;
    private String taskId;

    @Override
    public void run() {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        for(int i = 0; i<4; i++) {
            System.out.println("<" + currentThreadName + "," + taskId + "> TICK TICK");
            try {
                TimeUnit.MILLISECONDS.sleep(sleepTime);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        System.out.println("#### <" + currentThreadName + "," + taskId + "> done...####");
    }

    public LoopTask(long sleepTime) {
        this.id = ++count;
        this.taskId = "Task-" + id;
        this.sleepTime = sleepTime;
    }
}
```

#### join()을 사용해서 일반 쓰레드에서 값 반환하기

join()은 마치 wait()와 notify()처럼 동작한다. 이런 join()의 동작을 이용하여 일반 쓰레드에서 값을 반환하는 것을 구현해 보자.

쓰레드에서 값 반환에서는 boolean 변수인 done과 synchronized를 사용했는데 join을 사용하면 그럴 필요가 없다.

쓰레드는 종료되었지만, Task 객체는 남아있기 때문에 join()을 한 후에 get() 메소드를 이용해 결과값을 읽어오면 된다.

```java
public class ReturningValueFromThreads {
    public static void main(String args[]) throws InterruptedException {
        System.out.println("Main thread starts here...");

        MyThreadTask task1 = new MyThreadTask();
        MyThreadTask task2 = new MyThreadTask();
        MyThreadTask task3 = new MyThreadTask();

        Thread thread1 = new Thread(task1,"firstThread");
        Thread thread2 = new Thread(task2,"secondThread");
        Thread thread3 = new Thread(task3,"thirdThread");
        thread1.start();
        thread2.start();
        thread3.start();

        thread1.join();
        System.out.println("task1 result:" + task1.getRandomSum());
        thread2.join();
        System.out.println("task2 result:" + task2.getRandomSum());
        thread3.join();
        System.out.println("task3 result:" + task3.getRandomSum());

        System.out.println("Main thread ends here...");
    }
}

class MyThreadTask implements Runnable {
    private static int count = 0;
    private int id;
    private int randomSum = 0;

    @Override
    public void run(){
        for(int i = 0; i<5; i++) {
            System.out.println("<" + id + ">TICK TICK " + i);
            randomSum += Math.random()*1000;
            try {
                TimeUnit.MICROSECONDS.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    public int getRandomSum(){
        return randomSum;
    }

    public MyThreadTask() {
        this.id = ++count;
    }
}
```

### Executor Task가 종료 될때까지 대기하기

다른 Executor Task가 끝날때까지 한 Task를 정지해 놓는 방법을 알아보자.
Java는 이를 위해 java.util.concurrent.CountDownLatch 클래스를 제공한다.
일반 쓰레드는 Thread.join()하나를 사용할 뿐이지만, Executor는 클래스를 사용해야 한다.
CountDownLatch object는 종료할 Task들과 그 Task들이 종료하길 기다리는 Task들이 모두 공유한다.
몇 개의 Task들이 종료해야 하는 지를 0보다 큰 숫자로 정해놓고 하나씩 줄여나가면서 기다리는 방식이다.

- 대기하는 Task: void await() 메소드를 호출하여 count가 0이 될때까지 블록되어 기다린다.
- 종료하는 Task: void coutndown() 메소드를 실행이 끝난후에 호출한다. 이러면 count가 1 감소한다.

count가 0이 되면 대기하던 Task들은 모두 블록킹 상태에서 풀려나서 실행을 지속하게 된다.

```java
public class JoiningExecutors {
    public static void main(String args[]) {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("["+ currentThreadName + "]" + " thread starts here...");

        ExecutorService execService1 = Executors.newCachedThreadPool(new NamedThreadsFactory());
        CountDownLatch doneSignal = new CountDownLatch(4); //set the initial count

        execService1.execute(new LoopTask(doneSignal));
        execService1.execute(new LoopTask(doneSignal));
        execService1.execute(new LoopTask(doneSignal));
        execService1.execute(new LoopTask(doneSignal));

        execService1.shutdown();

        try {
            doneSignal.await(); //wait for the count = 0
            System.out.println("["+ currentThreadName + "]" + " got the signal to continue...");
        } catch(InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println("["+ currentThreadName + "]" + " thread ends here...");
    }
}

class NamedThreadsFactory implements ThreadFactory {

    private static int count = 0;
    private static String Name = "MyThread-";

    @Override
    public Thread newThread(Runnable r) {
        return new Thread(r, Name + ++count);
    }
}

class LoopTask implements Runnable {
    private static int count = 0;
    private int id;
    private String taskId;
    private CountDownLatch doneCountLatch;

    @Override
    public void run(){
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId + "> starting...####");
        for(int i = 0; i<5; i++) {
            System.out.println("<" + currentThreadName +"," + taskId + "> TICK TICK" + i);
            try {
                TimeUnit.MILLISECONDS.sleep((long)Math.random()*1000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
        System.out.println("#### <" + currentThreadName +"," + taskId + "> done...####");
        if(doneCountLatch != null) {
            doneCountLatch.countDown();    //count--
            System.out.println("#### <" + currentThreadName +"," + taskId + "> LATCH COUNT =" + doneCountLatch.getCount());
        }
    }

    public LoopTask(CountDownLatch doneCountLatch) {
        this.id = ++count;
        this.taskId = "LoopTask" + id;
        this.doneCountLatch = doneCountLatch;
    }
}
```

## 쓰레드 스케쥴링

쓰레드의 개수가 코어의 수보다 많을 경우, 쓰레드를 어떤 순서에 의해 동시성으로 실행할 것인가를 결정해야 하는데,
이것을 쓰레드 스케줄링이라고 한다.

쓰레드 스케줄링에 의해 쓰레드 들은 아주 짧은 시간 번갈아가면서 `run()` 메소드를 조금씩 실행한다.

### 일반 쓰레드의 스케쥴링

쓰레드 스케줄링을 위해 Java Thread는 두가지 API를 제공한다.

- java.util.Timer
- java.util.TimerTask

특징

- Timer 객체당 한번에 하나의 Task만 실행된다.
- Timer Task는 재빨리 종료되어야 한다.
- Application이 종료될 때 반드시 Timer.cancel()이 호출되어야 한다. 그렇지 않으면 leakage가 발생한다.
- Timer가 cancel되면 더이상 Task는 스케줄되지 않는다.
- Timer 클래스는 thread-safe이다.
- Timer 클래스는 실시간성을 보장하지 않는다.

#### java.util.Timer

하나의 쓰레드를 생성하며, 그 하나의 쓰레드에서 모든 Task를 실행한다.

- void schedule(TimerTask task, Date time) : 정해진 time에 한번 실행
- void schedule(TimerTask task, long delay) : 현재부터 delay millisec 이후 한번 실행
- void schedule(TimerTask task, Date firstTime, long period): firstTime에 처음 실행하고 period 간격으로 계속 실행
- void schedule(TimerTask task, long delay, long period): 현재부터 delay millisec 이후 처음 실행하고 period 간격으로 계속 실행
- void scheduleAtFixedRate(TimerTask task, Date firstTime, long period): firstTime에 처음 실행하고 period 간격으로 계속 고정 Rate으로 실행
- void scheduleAtFixedRate(TimerTask task, long delay, long period): 현재부터 delay millisec 이후 처음 실행하고 period 간격으로 계속 계속 고정 Rate으로 실행
- void cancel(): Timer를 terminate. **실행되고 있던 Task는 계속 실행** 되며, 더이상 스케줄은 되지 않음

파라미터

- TimerTask task: 실행될 Task
- Date time: 실행 시작할 절대 시간
- long delay: 현재 시간으로부터의 delay offset. Milliseconds
- long period: Task 실행시간 간의 고정 delay

#### java.util.TimerTask

스케줄링 대상이 되는 Task API이다. Runnable 인터페이스를 구현한 것이다. 반복된 실행을 위해 짧게 수행(short-lived)되어져야 한다.

- long scheduleExecutionTime(): 동작하는 시간으로 가장 최신에 실행된 시간을 반환한다.
- boolean cancel(): Task를 취소한다. 만약 Task가 구동중이면 구동은 지속되지만 다시 실행되지 않는다

#### 한번 실행되는 Task 스케줄링하기

다음 두 메소드를 이용해서 Task를 한번 스케줄링 할 수 있다.

- void schedule(TimerTask task, Date time) : 정해진 time에 한번 실행
- void schedule(TimerTask task, long delay) : 현재부터 delay millisec 이후 한번 실행

```java
public class SchedulingTasksForOneTimeExecution {
    private static SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("Main thread starts here...");

        Timer timer = new Timer("Timer-thread", false); //false: user thread, true: daemon
        Date currentTime = new Date();
        System.out.println("["+currentThreadName+"] Current time" + dateFormatter.format(currentTime));


        Date scheduledTime = TimeUtils.getFutureTime(currentTime, 5000);
        ScheduledTask task0 = new ScheduledTask(100);
        timer.schedule(task0, scheduledTime); //schedule(TimerTask task, Date time)
        System.out.println("["+currentThreadName+"] Task-0 is scheduled for running at " + dateFormatter.format(currentTime));

        long delayMillis = 10000;
        ScheduledTask task1 = new ScheduledTask(100);
        timer.schedule(task1, delayMillis); //schedule(TimerTask task, long delay)
        System.out.println("["+currentThreadName+"] Task-1 is scheduled for running "+ delayMillis/1000 + " at " + dateFormatter.format(currentTime));

        ScheduledTask task2 = new ScheduledTask(100);
        timer.schedule(task2, delayMillis); //schedule(TimerTask task, long delay)
        System.out.println("["+currentThreadName+"] Task-2 is scheduled for running "+ delayMillis/1000 + " at " + dateFormatter.format(currentTime));

        task1.cancel();    //task1 canceled

        TimeUnit.MILLISECONDS.sleep(12000);
        timer.cancel();
        System.out.println("Main thread ends here...");
    }
}

class ScheduledTask extends TimerTask {
    private static int count = 0;
    private long sleepTime;
    private String taskId;
    private SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    @Override
    public void run(){
        Date startTime = new Date();
        Date SchdulingForRunningTime = new Date(super.scheduledExecutionTime());
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId
                + " > Scheduled to run at "
                + dateFormatter.format(SchdulingForRunningTime)
                + ", Actually started at "
                + dateFormatter.format(startTime)
                + "####");

        for(int i = 0; i<5; i++) {
            try {
                TimeUnit.MICROSECONDS.sleep(sleepTime);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("#### <" + currentThreadName + "," + taskId
                + " >Finished at "
                + dateFormatter.format(new Date())
                + "####");
    }

    public ScheduledTask(long sleepTime) {
        this.sleepTime = sleepTime;
        this.taskId = "ScheduledTask-" + count++;
    }
}

/* Utility class to get the future time*/
class TimeUtils {
    private TimeUtils() {

    }
    public static Date getFutureTime(Date initialTime, long millisToAdd) {
        Calendar cal = GregorianCalendar.getInstance();
        cal.setTimeInMillis(initialTime.getTime() + millisToAdd);
        return cal.getTime();
    }
}
```

#### 반복되는 Task 스케줄링하기

아래 메소드를 이용하여 반복되는 Task를 스케줄링한다.

- void schedule(TimerTask task, Date firstTime, long period): firstTime에 처음 실행하고 period 간격으로 계속 실행
- void schedule(TimerTask task, long delay, long period): 현재부터 delay millisec 이후 처음 실행하고 period 간격으로 계속 실행

간격은 이전 Task 시작부터 다음 Task 시작까지이다. 만약 정한 간격보다 길어지면 시작 시간이 밀려서 시작하게 된다.

```java
public class SchedulingTasksFixedDelayRepeatedExecution {
    private static SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("Main thread starts here...");

        Timer timer = new Timer("Timer-thread", false); //false: user thread, true: daemon
        Date currentTime = new Date();
        System.out.println("["+currentThreadName+"] Current time" + dateFormatter.format(currentTime));

        Date scheduledTime = TimeUtils.getFutureTime(currentTime, 1000);
        ScheduledTask task0 = new ScheduledTask(100);
        long periodMillis = 1000;
        timer.schedule(task0, scheduledTime, periodMillis); //schedule(TimerTask task, Date firstTime, long period)
        System.out.println("["+currentThreadName+"] Task-0 is scheduled for running at " + dateFormatter.format(currentTime));

        long delayMillis = 5000;
        periodMillis = 5000;
        ScheduledTask task1 = new ScheduledTask(100);
        timer.schedule(task1, delayMillis, periodMillis); //schedule(TimerTask task, long delay, long period)
        System.out.println("["+currentThreadName+"] Task-1 is scheduled for running "+ delayMillis/1000 + " at " + dateFormatter.format(currentTime));

        TimeUnit.MILLISECONDS.sleep(11000);
        timer.cancel();
        System.out.println("Main thread ends here...");
    }
}

class ScheduledTask extends TimerTask {
    private static int count = 0;
    private long sleepTime;
    private String taskId;
    private SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    @Override
    public void run(){
        Date startTime = new Date();
        Date SchdulingForRunningTime = new Date(super.scheduledExecutionTime());
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId
                + " > Scheduled to run at "
                + dateFormatter.format(SchdulingForRunningTime)
                + ", Actually started at "
                + dateFormatter.format(startTime)
                + "####");

        for(int i = 0; i<5; i++) {
            try {
                TimeUnit.MICROSECONDS.sleep(sleepTime);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("#### <" + currentThreadName + "," + taskId
                + " >Finished at "
                + dateFormatter.format(new Date())
                + "####");
    }

    public ScheduledTask(long sleepTime) {
        this.sleepTime = sleepTime;
        this.taskId = "ScheduledTask-" + count++;
    }
}

/* Utility class to get the future time*/
class TimeUtils {
    private TimeUtils() {

    }
    public static Date getFutureTime(Date initialTime, long millisToAdd) {
        Calendar cal = GregorianCalendar.getInstance();
        cal.setTimeInMillis(initialTime.getTime() + millisToAdd);
        return cal.getTime();
    }
}
```

#### FixedRate로 반복되는 Task 스케줄링하기

아래 메소드를 이용하여 반복되는 Task를 스케줄링한다.

- void scheduleAtFixedRate(TimerTask task, Date firstTime, long period)
- void scheduleAtFixedRate(TimerTask task, long delay, long period)

시작하는 시점은 맨 처음 Task 실행 시간을 기준으로 고정되어 계산 된다.
FixedRate가 아닌 경우는 바로 이전의 시작을 기준으로 계산한다.
만약 정한 간격보다 길어지면 시작 시간이 밀려서 시작하는 것이 아니라 바로 그 시간에 시작한다.
그러므로 절대 시간에 관한 요구사항을 가진 경우 FixedRate을 사용한다.
컨셉을 이해한다면 이전 예와 동일하므로 에제를 가지고 실행하지 않는다.

### Executor Task 스케줄링

Task 스케줄링을 위해 Executors는 다음 두개의 쓰레드 풀을 별도로 제공한다.

- Single-thread-scheduled-executor: 하나의 쓰레드로 여러 Task의 스케줄링
- Scheduled-thread-pool: 여러 쓰레드로 여러 Task의 스케줄링

Executors 클래스를 사용해 생성하는 것은 다음과 같이 한다.

- static ScheduledExecutorService newSingleThreadScheduledExecutor()
- static ScheduledExecutorService newScheduledThreadPool(int corePoolSize)

ScheduledExecutorService 인터페이스는 ExecutorService 인터페이스를 확장한 것으로 다음과 같은 메소드를 정의하고 있다.

- `ScheduledFuture<?> schedule(Runnable command, long delay, TimeUnit unit)`
- `ScheduledFuture<?> schedule(Callable<V> callable, long delay, TimeUnit unit)`
- `ScheduledFuture<?> scheduleWithFixedDelay(Runnable, long initialDelay, long delay, TimeUnit unit)`
- `ScheduledFuture<?> scheduleWithFixedRate(Runnable, long initialDelay, long delay, TimeUnit unit)`

일반 쓰레드와 다른 점은 정확한 시간을 지정하여 구동하는 것이 없다는 것이다.

ScheduledFuture 인터페이스는 Future 인터페이스와 Delayed 인터페이스를 상속한다.
Delayed 인터페이스는 다음의 메소드만을 정의한다. 이것은 다음 실행 스케줄까지 얼마나 남았는지 반환한다.

`long getDelay(TimeUnit unit)`

#### Executor에서 한번 실행되는 Task 스케줄링하기

다음 두 메소드를 이용해서 Task를 한번 스케줄링 할 수 있다.

- ScheduledFuture<?> schedule(Runnable command, long delay, TimeUnit unit)
- ScheduledFuture<?> schedule(Callable callable, long delay, TimeUnit unit)

실행은 Single-thread-scheduled-executor 또는 Scheduled-thread-pool 중 하나를 사용하여 스케줄될 수 있다.
TimerTask는 필요없으며 Runnable 또는 Callable이면 된다.

```java
public class SchedulingTasksFixedDelayRepeatedExecution {
    private static SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("Main thread starts here...");

        ScheduledExecutorService execService = Executors.newSingleThreadScheduledExecutor(new NamedThreadsFactory());

        System.out.println("["+currentThreadName+"] Current time" + dateFormatter.format(new Date()));

        ScheduledFuture<?> schedFuture1 = execService.schedule(new ScheduledRunnableTask(0), 2, TimeUnit.SECONDS);
        ScheduledFuture<Integer> schedFuture2 = execService.schedule(new ScheduledCallableTask(0), 4, TimeUnit.SECONDS);

        execService.shutdown();
        try {
            System.out.println("Task1 result = " + schedFuture1.get());
            System.out.println("Task2 result = " + schedFuture2.get());
        } catch (ExecutionException e) {
            e.printStackTrace();
        }

        System.out.println("Main thread ends here...");
    }
}

class NamedThreadsFactory implements ThreadFactory {

    private static int count = 0;
    private static String Name = "MyThread-";

    @Override
    public Thread newThread(Runnable r) {
        return new Thread(r, Name + ++count);
    }
}

class ScheduledRunnableTask implements Runnable {
    private static int count = 0;
    private long sleepTime;
    private String taskId;
    private SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    @Override
    public void run(){
        Date startTime = new Date();
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId
                + " > Actually started at "
                + dateFormatter.format(startTime)
                + "####");

        for(int i = 0; i<5; i++) {
            try {
                TimeUnit.MICROSECONDS.sleep(sleepTime);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("#### <" + currentThreadName + "," + taskId
                + " >Finished at "
                + dateFormatter.format(new Date())
                + "####");
    }


    public ScheduledRunnableTask(long sleepTime) {
        this.sleepTime = sleepTime;
        this.taskId = "ScheduledRunnableTask-" + count++;
    }
}

class ScheduledCallableTask implements Callable<Integer> {
    private static int count = 0;
    private long sleepTime;
    private String taskId;
    private SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    @Override
    public Integer call() throws Exception {
        Date startTime = new Date();
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId
                + " > Actually started at "
                + dateFormatter.format(startTime)
                + "####");

        for(int i = 0; i<5; i++) {
            try {
                TimeUnit.MICROSECONDS.sleep(sleepTime);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("#### <" + currentThreadName + "," + taskId
                + " >Finished at "
                + dateFormatter.format(new Date())
                + "####");
        return 10; //whatever
    }


    public ScheduledCallableTask(long sleepTime) {
        this.sleepTime = sleepTime;
        this.taskId = "ScheduledCallableTask-" + count++;
    }
}
```

#### Executor에서 반복 실행되는 Task 스케줄링하기

반복 실행되는 Task는 다음의 메소드를 이용한다.

- ScheduledFuture<?> scheduleWithFixedDelay(Runnable, long initialDelay, long delay, TimeUnit unit)

주의할 점은 다음과 같다.

- 인터벌은 이전 Task 실행이 언제 끝났느냐에 따라 달라진다. 일반 쓰레드가 시작을 기준으로 정하는 것과 구별된다.
- Callable Task의 반복은 지원하지 않는다. Runnable만 지원된다.

```java
public class SchedulingTasksFixedDelayRepeatedExecution {
    private static SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("Main thread starts here...");

        ScheduledExecutorService execService = Executors.newSingleThreadScheduledExecutor(new NamedThreadsFactory());

        System.out.println("["+currentThreadName+"] Current time" + dateFormatter.format(new Date()));

        //ScheduledFuture<?> scheduleWithFixedDelay(Runnable, long initialDelay, long delay, TimeUnit unit)
        ScheduledFuture<?> schedFuture = execService.scheduleWithFixedDelay(new ScheduledRunnableTask(0), 4, 2, TimeUnit.SECONDS);

        TimeUnit.MILLISECONDS.sleep(10000);
        execService.shutdown();

        System.out.println("Main thread ends here...");
    }
}

class NamedThreadsFactory implements ThreadFactory {

    private static int count = 0;
    private static String Name = "MyThread-";

    @Override
    public Thread newThread(Runnable r) {
        return new Thread(r, Name + ++count);
    }
}

class ScheduledRunnableTask implements Runnable {
    private static int count = 0;
    private long sleepTime;
    private String taskId;
    private SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    @Override
    public void run(){
        Date startTime = new Date();
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("#### <" + currentThreadName +"," + taskId
                + " > Actually started at "
                + dateFormatter.format(startTime)
                + "####");

        for(int i = 0; i<5; i++) {
            try {
                TimeUnit.MICROSECONDS.sleep(sleepTime);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        System.out.println("#### <" + currentThreadName + "," + taskId
                + " >Finished at "
                + dateFormatter.format(new Date())
                + "####");
    }

    public ScheduledRunnableTask(long sleepTime) {
        this.sleepTime = sleepTime;
        this.taskId = "ScheduledRunnableTask-" + count++;
    }
}
```

#### Executor에서 FixedRate로 반복 실행되는 Task 스케줄링하기

반복 실행되는 Task는 다음의 메소드를 이용한다.

- ScheduledFuture<?> scheduleAtFixedRate(Runnable, long initialDelay, long delay, TimeUnit unit)

주의할 점은 다음과 같다.

- 맨처음 Task의 시작 시간을 기준으로 각각 실행이 스케줄 된다. 일반 쓰레드의 FixedRate와 동일한 점이다.
- Callable Task의 반복은 지원하지 않는다. Runnable만 지원된다.

```java
public class SchedulingTasksFixedDelayRepeatedExecution {
    private static SimpleDateFormat dateFormatter = new SimpleDateFormat("dd-MM-yyyy HH:mm:ss.SSS");

    public static void main(String args[]) throws InterruptedException {
        String currentThreadName = Thread.currentThread().getName();
        System.out.println("Main thread starts here...");

        ScheduledExecutorService execService = Executors.newSingleThreadScheduledExecutor(new NamedThreadsFactory());

        System.out.println("["+currentThreadName+"] Current time" + dateFormatter.format(new Date()));

        //ScheduledFuture<?> schedFuture = execService.scheduleWithFixedDelay(new ScheduledRunnableTask(0), 4, 2, TimeUnit.SECONDS);
        ScheduledFuture<?> schedFuture = execService.scheduleAtFixedRate(new ScheduledRunnableTask(0), 4, 2, TimeUnit.SECONDS);

        TimeUnit.MILLISECONDS.sleep(10000);
        execService.shutdown();

        System.out.println("Main thread ends here...");
    }
}
```
