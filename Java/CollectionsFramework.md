# Collections Framework

초창기 자바에서는 Vector, Stack, Hashtable, Array만 구현되어 있었다
자바 1.2가 릴리즈 되면서 컬렉션 프레임워크들은 Colletions 인터페이스를 구현하고 해당 알고리즘도 구현되었다.

자바 1.5에서 도입된 제너릭은 컬렉션 프레임워크에도 적용되었고 제너릭을 통해 컬렉션내의 타입을 지정할수 있게 되었다.
이를 통해 컴파일 시점에서 컬렉션내에 타입이 일치하는지를 파악할 수 있게 되었고, 사용시 별도로 캐스팅을 하지 않아도 된다.

## hierarchy

### Collection hierarchy

![collection](https://raw.githubusercontent.com/pravusid/TIL/master/Java/img/collection_hierarchy.png)

### Map hierarchy

![map](https://raw.githubusercontent.com/pravusid/TIL/master/Java/img/map_hierarchy.png)

- java.util.SortedMap: key가 ascending order로 정렬된 map
- java.util.NavigableMap

## 종류

| Interface | Hash Table | Resizable Array | Balanced Tree | List | Hash Table + Linked List |
| --- | --- | --- | --- | --- | --- |
| Set | HashSet |  | TreeSet |  | LinkedHashSet|
| List |  | ArrayList |  | LinkedList |  |
| Deque |  | ArrayDeque |  | LinkedList |  |
| Map | HashMap |  | TreeMap |  | LinkedHashMap |

### Legacy Class

- Vector: 동기화 지원 List
- Stack: Vector를 이용한 LIFO 구현
- Dictionary: key/value 저장(Map과 유사)하는 abstract class
- Hashtable: Dictionary의 구현
- Properties: Hashtable의 subclass로 key/value 모두 String인 경우에 리스트를 관리
- BitSet: Bit 값을 저장하는 특별한 유형의 array. 필요시 사이즈 증가 가능

## List

정적 리스트 생성: `List<String> messages = Arrays.asList("Hello", "World!", "How", "Are", "You");`

List 자료구조는 순서가 있고 element의 중복을 허용한다.

- ArrayList: resizable array, synchronized 보장하지 않음
- Vector: resizable array, synchronized 보장
- LinkedList: linked list 자료구조 구현체
- Stack: stack 자료구조 구현체

### ArrayList

동기화를 보장하지 않으며 동기화가 필요할 때는 `Collections.synchronizeList()` 메소드를 통해 동기화가 보장되는 List를 반환받아 사용한다.

가변크기를 가진다. (element가 추가되면 컬렉션의 저장공간 크기가 변할 수도 있다)

### Vector

동기화를 보장되고 가변 크기의 자료구조이다.

ArrayList도 동기화 지원으로 변경하여 사용할수 있기 때문에 사실상 Legacy로 구분되어 쓰지 않는다.

### LinkedList

동기화를 보장하지 않으며 동기화가 필요할 때는 `Collections.synchronizeList()` 메소드를 통해 동기화가 보장되는 List를 반환받아 사용한다.

선입 선출인 Queue와 양쪽 끝에서의 처리를 하는 Deque의 속성과 메소드를 가지고 있다.

### Stack

Deque 인터페이스의 속성을 물려받아 메소드만 LIFO에 맞게 정의한 것이다.

Stack은 LIFO(후입 선출)을 지원

### ArrayList와 LinkedList의 차이점

ArrayList는 인덱스 기반의 Array로 구성되어 있어서 랜덤 엑세스를 할 경우 O(1)의 속도를 가진다.

LinkedList는 데이터들이 서로 연결된 node로 구성되어 있다.
인덱스 번호를 사용해서 엘리먼트에 접근 하더라도 내부적으로는 노드들을 순차적으로 순회하며 엘리먼트를 찾는다
LinkedList의 탐색 속도는 O(n)으로 ArrayList 보다 느리다.

엘리먼트의 추가 및 삭제는 LinkedList가 ArrayList보다 빠르다.
엘리먼트를 추가 및 삭제하는 중에 array를 리사이즈 하거나 인덱스를 업데이트를 할 일이 없기 때문이다.

LinkedList의 엘리먼트들은 이전, 다음 엘리먼트들에 대한 정보를 가지고 있기 때문에 LinkedList가 ArrayList보다 더 많은 메모리를 소비한다.

## Set

순서가 없고 중복을 허용하지 않는 자료구조이다. 순서가 없으므로 index를 통해 접근하지 않는다.

- HashSet: Java Collection의 대표 Set 자료구조
- TreeSet: Set 인터페이스를 상속한 SortedSet 인터페이스를 구현
- EnumSet: Enum 타입을 활용해서 Set을 구현

## Queue

### BlockingQueue

java.util.concurrent.BlockingQueue는 엘리먼트들을 검색하거나 삭제 할때 대기하고,
큐에 엘리먼트가 추가 될 때 저장공간이 충분해 질때까지 기다리는 기능을 제공하는 Queue 이다.
BlockingQueue는 자바 컬렉션 프레임워크에서 제공하는 인터페이스중에 하나로 주로 producer-consumer 문제에 주로 사용된다.
BlockingQueue를 사용하면 producser가 cosumer에게 Object를 전달할때 저장공간 부족에 따르는 여러 문제점을 걱정할 필요가 없다.
Java에서는 BlockingQueue를 구현한 ArrayBlockingQueue, LinkedBlockingQueue, PriorityBlockingQueue, SynchronousQueue등을 지원 한다.

## Map

Key를 이용해서 Value를 찾는 Dictionary 자료구조이다.

- HashMap: 동기화를 보장하지 않는 Java Collection Framework의 대표 Map
- Hashtable: 동기화를 보장 (Collection Framework이 아닌 Legacy class를 구현)
- TreeMap: Map 인터페이스를 상속한 SortedMap 인터페이스를 구현

### HashMap

![HashMap](https://raw.githubusercontent.com/pravusid/TIL/master/Java/img/java-hashmap.png)

HashMap은 키-값 쌍으로 사용하도록 구현되어 있다.
HashMap은 해싱 알고리즘을 사용하고 hashCode()와 equals()를 put() 과 get()을 쓸때 사용한다.

key-value를 저장하기 위해 put 메소드를 호출 하면 key의 hashCode()를 호출해서 맵에 저장되어 있는 값 중에 동일한 key가 있는지 찾는다.
이 Entry는 LinkedList에 저장되어 있고(요소가 증가하면 Tree) 만약 존재하는 entry면 equals()메서드를 사용해서 key가 이미 존재 하는지 확인 하고,
만약 존재 한다면 value값을 덮어 씌워서 새로운 키-값 으로 저장한다.

키를 가지고 get 메서드를 호출하면 hashCode()를 호출해서 array에서 값을 찾고 equals()메서드를 가지고 찾고자 하는 key와 동일한지 확인한다.

### Hashtable

HashMap은 보조 해시 함수(Additional Hash Function)를 사용하기 때문에
보조 해시 함수를 사용하지 않는 HashTable에 비하여 해시 충돌(hash collision)이 덜 발생할 수 있어 상대으로 성능상 이점이 있다.

HashMap은 키/값에 null을 허용하는 반면 Hashtable은 이를 허용하지 않는다.

Hashtable은 synchronized (synchronized) 되어 있지만 HashMap은 그렇지 않다.

LinkedHashMap 은 자바 1.4에서 HashMap의 서브클래스로 소개되었다.
그렇기 때문에 iteration 의 순서를 보장받고 싶다면, HashMap에서 LinkedHashMap으로쉽게 변경 가능하다.
그러나 Hashtable 에서는 그럴 수 없으므로 iteration 순서를 예측할 수 없다.

HashMap은 iterator 키 셋을 제공하므로 fail-fast 기능을 사용하나 Hashtable은 Enumeration 키를 사용하므로 이런 기능을 제공하지 못한다.

Hashtable은 legacy 클래스로 취급을 받기 때문에 만약 Map에서 iteration을 하는 도중에 수정가능한 Map을 사용하고 싶다면 ConcurrentHashMap을 사용하면 된다.

### TreeMap

일반적인 목적에서는 HashMap이 가장 좋은 선택이다.

하지만 만약 정렬되어 있는 key값에 따라 탐색을 하기 원한다면 TreeMap을 사용하는 것이 더 좋다.
컬렉션에 크기에 따라 다르지만 HashMap에 엘리먼트를 추가 하고 이를 TreeMap으로 변환하는게 키를 정렬해서 탐색하는 경우보다 더 빠르게 동작 한다.

## Iterator / Enumeration

Enumeration은 Iterator 보다 빠르고 더 작은 메모리를 사용한다.
Enumeration은 매우 간단하고 간단한 요구사항에 잘 동작되도록 최적화 되어 있다.
Iterator는 자바 컬렉션 프레임워크의 Enumeration에 포함된다.

Iterator는 사용될때 대상 컬렉션을 다른 쓰레드에서 접근해서 수정하는것을 막는다.
iterator는 작업을 수행하면서 해당 엘리먼트를 삭제할 수 있지만 Enumeration은 불가능 하다.

Iterator의 fail-fast 속성은 다음 엘리먼트에 접근 하려고 할 때 엘리먼트가 변한것이 있는지 확인하는 것이다.
만약 수정 사항이 발견된다면 ConcurrentModificationException를 발생시킨다.

모든 Iterator의 구현체는 ConcurrentHashMap이나 CopyOnWriteArrayList같은 동시성 관련 컬렉션을 제외하고
fail-fast를 사용하는 방법으로 디자인 되어 있다.

## Concurrent Collection

Java 1.5 Concurrent 패키지는 thread-safe 하고 이터레이팅 작업 중에 컬렉션을 수정할 수 있는 클래스들을 포함하고 있다.
Iterator는 fail-fast 하도록 디자인되어있고, ConcurrentModificationException을 발생 시킨다.
가장 잘 알려진 클래스로는 CopyOnWriteArrayList, ConcurrentHashMap, CopyOnWriteArraySet이 있다.

Vector, Hashtable, Properties, stack 은 synchronized 되어있는 클래스로 thread-safe기 때문에 multi-thread 환경에서도 정삭적으로 동작한다.
Java 1.5의 Concurrent API에 포함되어 있는 몇몇 컬렉션 클래스는 이터레이팅 작업을 수행하는 도중에 컬렉션을 수정할 수 있는데,
이는 컬렉션의 복사본을 통해 작업을 하고 있기 때문이고 이들 역시 multi-thread 환경에서 안전한다.

Concurrent programming을 지원하기 위한 Concurrent Collection 인터페이스

- BlockingQueue
- TransferQueue
- BlockingDeque
- ConcurrentMap
- ConcurrentNavigableMap

클래스

- LinkedBlockingQueue
- ArrayBlockingQueue
- PriorityBlockingQueue
- DelayQueue
- SynchronousQueue
- LinkedBlockingDeque
- LinkedTransferQueue
- CopyOnWriteArrayList
- CopyOnWriteArraySet
- ConcurrentSkipListSet
- ConcurrentHashMap
- ConcurrentSkipListMap

기존 컬렉션을 가지고 동기화된 컬렉션을 만들기

`Collections.synchronizedCollection(Collection c)`를 사용해서 동기화된(thread-safe)한 컬렉션을 만들 수 있다.

## Comparable / Comparator Interface

기본 정렬 기능을 구현하기 위해 `Comparable` 인터페이스를 제공한다.
이 인터페이스는 `compareTo(Object o)` 메소드를 통해 정렬한다.
이 메서드를 구현할때 리턴값으로 음수, 0, 양수를 통해 엘리먼트들을 정렬하는데 사용한다.

Comparator 인터페이스는 두개의 파라미터를 가지고 있는 `compare(Object o1, Object o2)` 메소드를 제공한다.
이 메소드를 통해 정렬 알고리즘을 직접 상세히 구현할 수 있다.

## 병렬처리를 위한 컬렉션

- `java.util.concurrent.ConcurrentHashMap`
- `java.util.concurrent.ConcurrentLinkedQueue`

병렬처리를 위한 컬렉션은 부분(segment)잠금을 사용한다.
따라서 처리하는 요소가 포함된 부분만 잠그고 다른 요소에는 다른 쓰레드가 접근 가능하다.
