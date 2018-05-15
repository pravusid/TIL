# Domain Driven Design

`DDD START!`를 읽으며 정리한 내용임

## 도메인 모델

### Entity

각 엔티티 객체는 서로다른 식별자를 갖는다 또한 엔티티의 식별자는 바뀌지 않는다.
만약 두 엔티티의 식별자가 같다면 같은 객체라고 판단할 수 있다.

식별자 생성

- 규칙에 따라 생성
- UUID (universally unique identifier)
- 일렬번호 사용 (시퀀스, auto increment ...)

도메인 모델의 엔티티는 데이터 뿐만 아니라 도메인 기능을 제공한다.
즉 주문을 표현하는 엔티티는 주문과 관련된 데이터와 배송지 주소 변경을 위한 기능도 제공한다.

### Value

밸류 타입은 개념적으로 완전한 하나를 표현할 때 사용한다.

받는 사람과 주소에 대한 데이터를 갖고있는 `ShippingInfo` 클래스가 있다고 하자.

```java
public class ShippingInfo {
  private String receiverName;
  private String receiverPhoneNumber;

  private String shippingAddress1;
  private String shippingAddress2;
  private String shippingZipcode;

  // constructor & getters
}
```

위의 예제에서 받는사람을 위한 밸류 타입인 `Receiver`를 다음과 같이 표현할 수 있다.

```java
public class Receiver {
  private String name;
  private String phoneNumber;

  // constructor & getters
}
```

마찬가지로 주소 관련 데이터도 `Address` 밸류타입을 생성하여 표현할 수 있다.

```java
public class Address {
  private String address1;
  private String address2;
  private String zipcode;

  // constructor & getters
}
```

밸류타입을 이용해서 처음에 정의한 `ShippingInfo` 클래스를 다시 쓰면 다음과 같다.

```java
public class ShippingInfo {
  private Receiver receiver;
  private Address address;

  // constructor & getters
}
```

밸류타입이 반드시 두 개 이상의 데이터를 가질필요는 없다. 의미를 명확하게 정의하기 위한 목적이면 충분하다.

가격 (`int price`)는 int 타입이지만 의미를 명확히 하기 위해서 `Money` 타입을 만들어 사용할 수 있다.
의미만 명확해지는 것이 아니라, `Money` 타입과 관련된 연산을 처리하기에도 적합하다.

```java
public class Money {
  private int value;

  // constructor & getter다

  public Money add(Money money) {
    return new Money(this.value + money.value);
  }

  public Money multiply(int multiplier) {
    return new Money(value * multiplier);
  }
}
```

### Aggregate

도메인이 키질수록 도메인 모델도 커지면서 많은 엔티티와 밸류가 늘어나고 모델은 복잡해진다.
애그리거트는 관련된 엔티티와 밸류 객체를 개념적으로 묶은 것이다.
애그리거트를 사용하여 관련 객체를 묶어 애그리거트간의 관계로 도메인 모델을 이해하고 구현/관리할 수 있게 된다.

애그리거트에는 루트 엔티티가 있다.
루트 엔티티는 애그리거트에 속해있는 엔티티와 밸류 객체를 이용해서 애그리거트가 구현해야 할 기능을 제공한다.
애그리거트를 사용하는 코드는 애그리거트 루트가 제공하는 기능을 실행하고 루트를 통해서 간접적으로 애그리거트 내의 다른 엔티티나 밸류 객체에 접근하게 된다.
이는 애그리거트의 내부 구현을 숨겨서 애그리거트 단위로 구현을 캡슐화할 수 있도록 돕는다.

### Repository

리포지토리는 도메인 모델의 영속성을 처리한다.
리포지토리는 애그리거트 단위로 도메인 객체를 저장하고 조회하는 기능을 정의한다.

응용서비스에서 CRUD 작업을 하거나 트랜잭션을 관리할 때 리포지토리 구현에 영향을 받기때문에 서로 밀접한 연관이 있다.

### Domain Service

특정 엔티티에 속하지 않은 도메인 로직을 제공한다. 도메인 로직이 여러 엔티티와 밸류를 필요로 하는 경우 도메인 서비스에서 로직을 구현한다.

## 아키텍처

### 아키텍처 영역

아키텍처의 전형적인 네 개의 영역은 표현, 응용, 도메인, 인프라스트럭처이다.

표현 영역은 사용자의 요청을 받아 응용 영역에 전달하고 응용 영역의 처리 결과를 다시 사용자에게 보여주는 역할을 한다.

응용영역은 시스템이 사용자에게 제공해야 할 기능을 구현한다.
응용영역은 기능을 구현하기 위해 도메인영역의 도메인 모델을 사용한다.
응용서비스는 로직을 직접 수행하기보다는 도메인 모델에 로직 수행을 위임한다.

도메인 영역은 도메인 모델을 구현한다. 도메인 모델은 도메인의 핵심 로직을 구현한다.

인프라스트럭처 영역은 구현기술에 대한것을 다룬다. DBMS 연동처리, 메시징 큐에 메시지 전송, 메일발송기능 등이 있다.

### DIP (Dependency Inversion Principle)

일반적으로 아키텍처를 적용할 때는 `표현 -> 응용 -> 도메인 -> 인프라스트럭처`와 같이 계층적으로 구성한다.
각 계층은 상위계층에 의존하지 않고 하위계층에만 의존성이 있다.
이런 경우 대부분이 인프라스트럭처 영역에 의존하게 된다.
인프라스트럭처에 의존하게 되면 '테스트'와 '기능확장'에 어려움이 발생하게 된다. 이를 해결하기 위해 DIP를 적용한다.

저수준 모듈이 고수준 모듈에 의존하도록 하려면 추상화한 인터페이스를 사용한다.

우선 가격계산을 위해 인프라스트럭처 영역의 `DroolsRuleEngine`을 사용한 예제를 보자

```java
public class CalculateDiscountService {
  private DroolsRuleEngine ruleEngine;

  public CalculateDiscountService() {
    ruleEngine = new DroolsRuleEngine();
  }

  public Money calculateDiscount(OrderLine orderLines, String customerId) {
    Customer customer = findCustomer(customerId);
    MutableMoney money = new MutableMoney(0);
    List<?> facts = Arrays.asList(customer, money);
    facts.addAll(orderLines);
    ruleEngine.evalute("discountCalculation", facts);
    return money.toImmutableMoney();
  }

  private Customer findCustomer(String customerId) {
    ...
  }
}
```

`CalculateDiscountService` 입장에서 룰 적용을 `Drools`를 사용하든 `java`를 사용하든 상관없다.
따라서 룰 적용을 추상화한 인터페이스를 다음과 같이 정의할 수 있다.

```java
public interface RuleDiscounter {
  public Money applyRules(Customer customer, List<OrderLine> orderLines);
}
```

인터페이스를 이용해서 `CalculateDiscountService`를 바꾸면 다음과 같다.

```java
public class CalculateDiscountService {
  private CustomerRepository customerRepository;
  private RuleDiscounter ruleDiscounter;

  public CalculateDiscountService(CustomerRepository customerRepository, RuleDiscounter ruleDiscounter) {
    this.customerRepository = customerRepository;
    this.ruleDiscounter = ruleDiscounter;
  }조

  public Money calculateDiscount(OrderLine orderLines, String customerId) {
    Customer customer = findCustomer(customerId);
    return ruleDiscounter.applyRules(customer, orderLines);
  }

  private Customer findCustomer(String customerId) {
    Customer customer = customerRepository.findById(customerId);
    if (customer == null) throw new NoCustomerException();
    return customer;
  }
}
```

이와 같은 의존역전 원칙을 활용하면 구현체 교체 및 테스트에 용이하다.
다음은 위 코드를 테스트 하는 경우이다.

```java
public class CalculateDiscou조ntServiceTest {
  @Test(expect = NoCustomerException.class)
  public void noCustomer_thenExceptionShouldBeThrown() {
    CustomerRepository stubRepo = mock(CustomerRepository.class);
    when(stubRepo.findById("noCustId")).thenReturn(null);
    RuleDiscounter stubRule = (cust, lines) -> null;

    CalculateDiscountService discountSvc = new CalculateDiscountService(stubRepo, stubRule);
    discountSvc.calculateDiscount(someLines, "noCustId");
  }
}
```

위 테스트 코드를 실행하면 `NoCustomerException`이 발생한다.

## 애그리거트 (Aggregate)

애그리거트는 모델을 이해하는 데 도움을 줄 뿐만 아니라 일관성을 관리하는 기준이 된다.

애그리거트는 관련된 모델을 하나로 모은 것이기 때문에 한 애그리거트에 속한 객체는 유사하거나 동일한 라이프사이클을 갖는다.
애그리거트에 속한 구성요소는 대부분 함께 생성하고 함께 제거한다.

애그리거트는 경계를 갖는다. 한 애그리거트에 속한 객체는 다른 애그리거트에 속하지 않는다.
애그리거트는 독립된 객체 군이며, 각 애그리거트는 자기자신을 관리할 뿐 다른 애그리거트를 관리하지 않는다.
경계를 설정할 때 기본이 되는 것은 도메인 규칙과 요구사항이다.
도메인 규칙에 따라 함께 생성되는 구성요소는 한 애그리거트에 속할 가능성이 높다.

그러나 'A'가 'B'를 갖는다고 볼 수 있는 요구사항이 있더라도, A와 B가 한 애그리거트에 속하지 않을 수 있다.
상품과 리뷰를 예로 들 수 있다. 상품 상세페이지에서 리뷰를 보여줘야 하는 경우에 둘을 같이 묶어 생각할 수 있다.
하지만 상품과 리뷰는 함께 생성되지도 않고 함께 변경되지도 않는다.
또한 리뷰의 변경이 상품에 영향을 주지도 않고, 상품의 변경이 리뷰에 영향을 주지 않기 때문에 둘은 다른 애그리거트에 속한다.

일반적으로 다수의 애그리거트는 한 개의 엔티티 객체를 갖는 경우가 많고, 두 개이 이상의 엔티티로 구성되는 애그리거트는 드물게 존재한다.

### 애그리거트 루트

애그리거트 루트는 속한 모든객체의 일관성을 관리한다.
이를 위해 애그리거트가 제공해야 할 도메인 기능을 구현한다.

불필요한 중복을 피하고 애그리거트 루트를 통해서만 도메인 로직을 구현하게 만드려면 다음의 두가지를 일관성있게 적용해야 한다.

- 단순히 필드를 변경하는 `set` method를 public 범위로 만들지 않는다
- 밸류타입은 불변으로 구현한다 (즉 변경시 값이 적용된 새 객체를 반환한다 - equality는 유지한 채로)

### 트랜잭션 범위

성능문제가 있기 때문에 트랜잭션 범위는 작을수록 좋다.
한 트랜잭션에서는 한 개의 애그리거트만 수정해야 한다.
한 트랜잭션에서 한 애그리거트만 수정한다는 것은 애그리거트에서 다른 애그리거트를 변경하지 않음을 의미한다.

만약 부득이하게 한 트랜잭션으로 두 개 이상의 애그리거트를 수정해야 한다면,
애그리거트에서 다른 애그리거트를 직접 수정하지 말고 응용서비스에서 두 애그리거트를 수정하도록 구현해야 한다.

### 리포지토리와 애그리거트

애그리거트는 개념상 완전한 한 개의 도메인 모델을 표현하므로 객체의 영속성을 처리하는 리포지토리는 애그리거트 단위로 존재한다.

애그리거트의 상태가 변경되면 모든 변경을 원자적으로 저장소에 반영해야 한다.
RDBMS를 이용해서 리포지토리를 구현하면 트랜잭션을 이용하여 애그리거트 변경의 반영을 보장할 수 있다.
NO-SQL을 이용하면 한 개 애그리거트를 한 개 문서에 저장함으로써 변경을 손실없이 반영할 수 있다.

### ID를 이용한 애그리거트 참조

ORM 기술을 사용하면 필드를 사용해서 다른 애그리거트를 쉽게 참조할 수 있다.
하지만 애그리거트를 직접 참조할 때 여러 문제가 발생할 수 있다.

우선 다른 애그리거트에 쉽게 접근할 수 있으면, 다른 애그리거트의 상태를 쉽게 변경할 수 있게 된다.

또한 애그리거트를 직접 참조하면 성능과 관련된 여러 고민을 해야한다.
JPA의 경우 지연로딩과 즉시로딩을 지원한다.
단순히 연관된 객체의 데이터를 함께 화면에 보여주어야 하면 eager loading이 유리하지만,
애그리거트의 상태를 변경하는 기능을 위해서는 불필요한 객체를 함께 로딩할 필요가 없으므로 lazy loading이 유리하다.

직접 참조는 확장에 불리하다. 시스템이 확장되면서 하위 도메인별로 다른종류의 DBMS를 사용할 수 있다.

위의 문제들을 완화히기 위해서 ID를 이용한 참조를 사용할 수 있다.
ID를 이용한 참조는 DB 테이블에서 외래키를 참조하는 것과 비슷하다.

#### ID를 이용한 참조와 조회성능

다른 애그리거트를 ID로 참조하면 여러 애그리거트를 읽어야 할 때 조회속도문제가 발생할 수 있다.

조회대상이 N개일 때 N개를 읽어오는 한번의 쿼리와 연관된 데이터를 읽어오는 쿼리를 N번실행하는 N+1 조회문제가 생길 수 있다.
이 문제가 발생하지 않으려면 JOIN을 사용해야 한다.
JOIN을 사용하는 가장 쉬운방법은 객체참조 방식에서 즉시로딩을 사용하는 것이다. 하지만 ID 참조를 사용하지 못하게 되므로 해결책은 아니다.

ID 참조 방식을 사용하면서 N+1 조회문제를 발생시키지 않으려면 전용 조회쿼리를 사용해야 한다.
예를들어 데이터 조회를 위한 별도 DAO를 만들고 DAO의 조회 메서드에서 세타조인을 이용하면 된다.

```java
@Repository
public class JpaOrderViewDao implements OrderViewDao {
  @PersistenceContext
  private EntityManager em;

  @Override
  public List<OrderView> selectByOrderer(String ordererID) {
    String selectQuery =
        "SELECT new OrderView(o, m, p) " +
        "FROM Order o join o.orderLinews o1, Member m, Product p " +
        "WHERE o.orderer.memberId.id = :ordererID " +
        "AND o.orderer.memberId = m.id " +
        "AND index(o1) = 0 " +
        "AND o1.productId = p.id " +
        "ORDER BY o.number.number DESC";
    TypedQuery<OrderView> query = em.createQuery(selectQuery, OverView.class);
    query.setParameter("orderId", orderId);
    return query.getResultList();
  }
}
```

애그리거트마다 서로 다른 저장소를 사용하는 경우 한번의 쿼리로 조회할 수 없다.
이런 경우 조회 성능을 높이기 위해 캐시를 적용하거나 조회 전용 저장소를 따로 구성한다.

### 애그리거트간 집합 연관

개념적으로 애그리거트 간에 1:N 연관이 있더라도 성능상의 문제점으로 1:N 연관을 직접 반영하는 것은 드물다.
반대로 N:1 연관을 활용하여 성능상의 문제점을 해결할 수 있다.

M:N 연관은 양쪽 애그리거트에 컬렉션으로 연관을 만든다.
상품이 여러 카테고리에 속할 수 있다고 가정하면 상품과 카테고리는 M:N 연관을 맺는다.
1:N 연관처럼 M:N 연관도 실제 요구사항을 고려해서 구현에 포함시킬지 고려해야 한다.
요구사항을 고려하면 상품에서 카테고리로의 단방향 M:N 연관만 적용하면 된다.

RDBMS를 이용하여 M:N 연관을 구하려면 조인 테이블을 사용한다.
JPA를 사용하면 다음과 같은 매핑설정을 사용해서 ID 참조를 이용한 M:N 단방향 연관을 구현할 수 있다.

```java
@Entity
public class Product {
  @EmbeddedId
  private ProductId id;

  @ElementCollection
  @CollectionTable(name = "product_category", joinColumns = @JoinColumn(name = "product_id"))
  private Set<CategoryId> categoryIds;
  ...
}
```

## 리포지토리

### 엔티티와 밸류 기본 매핑

- 애그리거트 루트는 `@Entity`로 매핑
- Value는 `@Embeddable`로 매핑 (Value type을 명시한 class)
- Value type property는 `@Embedded`로 매핑 (Entity 및 Embeddable 내부의 property)
- Value 저장 전략
  - Table의 column과 Entity의 value가 일치하지 않는다면 (주소 : 주소1, 주소2) `AttributeConverter` 사용
  - Value 값이 컬렉션으로 저장될 때 별도의 테이블로 정규화 하는 경우 `@ElementCollection`과 `@CollectionTable`을 사용한다.

### 별도 테이블에 저장하는 밸류

애그리거트에서 루트 엔티티를 뺀 나머지 구성요소는 대부분 밸류이다.
루트 엔티티 외에 또 다른 엔티티가 있다면 진짜 엔티티인지 의심해봐야 한다.
밸류가 아니라 엔티티가 확실하다면 다른 애그리거트는 아닌지 확인해야 한다. 특히 자신만의 라이프사이클이 있다면 다른 애그리거트일 가능성이 높다.

애그리거트에 속한 객체가 밸류인지 엔티티인지 구분하는 방법은 고유 식별자를 갖는지 여부를 확인하는 것이다.
하지만 식별자를 찾을 때 매핑되는 테이블의 식별자를 애그리거트 구성요소의 식별자와 동일한 것으로 착각하면 안된다.

### 애그리거트 로딩전략

애그리거트는 개념적으로 하나여야 하지만 루트 엔티티를 로딩할 때 애그리거트에 속한 객체를 모두 로딩해야하는 것은 아니다.
JPA는 트랜잭션 범위내에서 지연로딩을 허용하기 때문에 실제로 상태를 변경하는 시점에서 필요한 구성요소만 로딩해도 문제가 되지 않는다.

### 애그리거트의 영속성 전파

- 저장 메소드는 애그리거트 루트만 저장하면 안 되고 애그리거트에 속한 모든객체를 저장해야 한다
- 삭제 메소드는 애그리거트 루트뿐만 아니라 애그리거트에 속한 모든 객체를 삭제해야 한다

`@Embeddable` 매핑타입의 경우 함께 저장되고 삭제되므로 `cascade`속성을 추가하지 않아도 된다.
반면 애그리거트에 속한 `@Entity` 타입에 대한 매핑은 `cascade` 속성을 사용해서 저장/삭제시 동시 처리되도록 해야한다.

## 리포지토리의 조회 기능

Specification은 애그리거트가 특정 조건을 충족하는지 여부를 검사한다.

```java
public interface Specification<T> {
  public boolean isSatisfiedBy(T agg);
}
```

리포지토리는 스펙을 전달받아 애그리거트를 걸러내는 용도로 사용한다.
스펙은 `AND`연산자나 `OR`연산자로 조합해서 새로운 조건을 만들 수 있다.

### JPA를 위한 스펙 구현

JPA에서 다양한 검색조건을 조합하기 위해 `CriteriaBuilder`와 `Predicate`를 사용한다.
JPA 리포지토리를 위한 `Specification`의 인터페이스는 다음과 같이 정의할 수 있다.

```java
public interface Specification<T> {
  Predicate toPredicate(Root<T> root, CriteriaBuilder cb);
}
```

위 스펙을 구현하면 다음과 같다.

```java
public class OrdererSpec implements Specification<Order> {
  private String ordererId;

  public OrdererSpec(String ordererId) {
    this.ordererId = ordererId
  }

  @Override
  public Predicate toPredicate(Root<Order> root, CriteriaBuilder cb) {
    return cb.equal(
      root.get(Order_.orderer).get(Orderer_.memberId).get(MemberId_.id),
      ordererId
    );
  }
}
```

서비스는 원하는 스펙을 생성하고 리포지토리에 전달하면 된다.

`List<Order> orders = orderRepository.findAll(new OrdererSpec("user"))`

### 조회전용 기능 구현

리포지토리는 애그리거트의 저장소를 표현하는 것으로 다음용도로는 적합하지 않다.

- 여러 애그리거트를 조합해서 한 화면에 보여주는 데이터 제공
- 각종 통계 데이터 제공

이런 기능은 조회 전용 쿼리로 처리해야 한다.
JPA와 하이버네이트에서는 동적 인스턴스 생성, 하이버네이트 `@Subselect` 확장기능, 네이티브 쿼리를 사용할 수 있다.

#### 동적 인스턴스 생성
