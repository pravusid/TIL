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

JPA는 JPQL에서 `new` 키워드를 사용하여 임의의 객체를 동적으로 생성할 수 있는 기능을 제공하고 있다.

### 하이버네이트 `@Subselect` 사용

`@Subselect`는 쿼리결과를 `@Entity`로 매핑할 수 있는 기능이다.

```java
@Entity
@Immutable
@Subselect("select o.order_number as number, " +
          "o.orderer_id, o.orderer_name, o.total_amounts, " +
          "o.receiver_name, o.state, o.order_date, " +
          "p.product_id, p.name as product_name " +
          "from purchase_order o inner join order_line ol " +
          "  on o.order_number = ol.order_number " +
          "  cross join product p " +
          " where ol.line_idx = 0 and ol.product_id = p.product_id")
@Synchronize({"purchage_order", "order_line", "product"})
public class OrderSummary {
  @Id
  private String number;
  private String ordererId;
  private STring ordererName;
  private int totalAmounts;
  private String receiverName;
  private String state;
  @Temporal(TemporalType.TimeStamp)
  @Column(name = "orderDate")
  private Date orderDate;
  private String productId;
  private String productName;
  ...
}
```

`@Immutable`, `@Subselect`, `@Synchronize`는 하이버네이트 전용 애노테이션이다.

`@Subselect`는 조회 쿼리를 값으로 갖는다. 하이버네이트는 `select` 쿼리 결과를 매핑할 테이블처럼 사용한다.
뷰를 수정할 수 없듯이 `@Subselect`로 조회한 Entity역시 수정할 수 없다.

트랜잭션 내에서 우선 반영된 사항을 함께 조회하기 위해서 `@Synchronize` 애노테이션을 사용한다.

## 응용서비스와 표현영역

### 응용서비스의 역할

- 도메인 객체간의 흐름 제어: 요청을 처리하기 위해 도메인 객체를 구하고 사용한다
- 트랜잭션 처리
- 접근제어
- 이벤트 처리

### 도메인 로직

도메인 로직을 도메인 영역과 응용서비스에 분산해서 구현하면 코드품질에 문제가 발생한다.
따라서 도메인 로직은 도메인 영역에 모아 응집도를 높여야 한다.

### 응용서비스의 구현

#### 응용서비스의 크기

- 한 응용 서비스 클래스의 회원 도메인 모든 기능 구현: 중복제거 용이
- 구분되는 기능별로 응용서비스 클래스를 따로 구현: 관련없는 코드가 공존할 가능성

#### 응용서비스의 인터페이스과 클래스

인터페이스와 - 인터페이스 구현 클래스가 1:1 관계인 경우 인터페이스를 분리해야할 필요가 있을지?

#### 메소드 파라미터와 값 리턴

파라미터 값이 여러개 이면 별도 클래스를 사용하는 것이 좋다. (DTO, 커맨드 오브젝트)

응용서비스에서 애그리거트 객체를 그대로 리턴하여 표현영역 코드까지 보낼 수도 있다.
이런경우 도메인 로직 실행을 응용서비스와 표현 영역 두곳에서 하여 코드 응집도가 낮아질 수 있다.

응용서비스는 표현영역에 필요한 데이터만 리턴하는 것이 기능 실행 로직의 응집도를 높이는 방법이 된다.

#### 표현영역에 의존하지 않을 것

응용서비스의 파라미터 타입에 표현영역과 관련된 타입을 사용하면 안된다.

예를 들어 `HttpServletRequest`나 `HttpSession`을 응용서비스에 파라미터로 전달하면 안된다.

응용서비스에서 표현영역에 대한 의존이 발생하면 테스트가 어려워 진다.
게다가 표현영역 구현이 변경되면 응용서비스 구현도 함께 변경해야 하는 문제도 발생한다.

#### 트랜잭션 처리

트랜잭션을 관리하는 것은 용용서비스의 중요한 역할이다.

#### 도메인 이벤트 처리

응용 서비스의 역할 중 하나는 도메인 영역에서 발생시킨 이벤트를 처리하는 것이다.
이벤트는 도메인에서 발생한 상태변경을 의미하며 '암호 변경됨', '주문 취소함'과 같은 것이 될 수 있다.

도메인 영역은 상태가 변경되면 이를 외부에 알리기 위해 이벤트를 발생시킬 수 있다.
예를 들어, 암호 초기화 기능은 암호 변경 후에 '암호 변경됨' 이벤트를 발생시킬 수 있다.

```java
public class Member {
  private Password password;

  public void initializePassword() {
    String newPassword = generateRandomPassword();
    this.password = new Password(newPassword);
    Events.raise(new PasswordChangedEvent(this.id, password));
  }
}
```

도메인에서 이벤트를 발생시키면 이벤트를 받아서 처리할 코드가 필요한데, 그 역할을 하는 것이 응용 서비스이다.
암호 초기화됨 이벤트가 발생하면 변경한 암호를 이메일로 발송하는 이벤트 핸들러를 등록할 수 있을 것이다.

```java
public class InitPasswordService {
  @Transactional
  public void initializePassword(String memberId) {
    Event.handle((PasswordChangedEvent evt) -> {
      ...
    });
    Member member = memberRepository.findById(memberId);
    checkMemberExist(member);
    member.initializePassword();
  }
}
```

`member.initializePassword()` 메소드를 실행하면 `PasswordChangedEvent`가 발생하고
`Event.handle()`에 등록한 이벤트 핸들러가 이벤트를 받아서 메일 발송을 할 것이라 예상할 수 있다.

### 표현영역

표현영역의 책임은 크게 다음과 같다.

- 사용자가 시스템을 사용할 수 있는 화면흐름을 제공하고 제어한다.
- 사용자의 요청을 알맞은 응용 서비스에 전달하고 결과를 사용자에게 제공한다.
- 사용자의 세션을 관리한다.

### 값 검증

Validation은 표현 영역과 응용 서비스 두 곳에서 모두 수행할 수 있다.

하지만 표현 영역은 잘못된 값이 존재하면 이를 사용자에게 알려주고 다시 값을 입력받아야 한다.
스프링 MVC의 경우 폼의 값이 잘못된 경우 `Errors`나 `BindingResult`를 사용하여 에러를 알려주는데,
응용서비스에서 값 검증을 하게 되면 코드가 복잡해 질 수 있다.

또한 `Exception`을 사용하면 사용자에게 좋지않은 경험을 제공한다.
값이 올바른지 한번에 확인하고 싶으나 예외가 발생하면 그 이후의 값은 검증하지 않기 때문이다.

이를 피하기 위해 응용서비스에 값을 전달하기 전에 표현영역에서 값을 검사할 수 있다.
하이버네이트등에서는 `@Valid`와 같은 값 검사 기능을 제공하므로 이 기능을 활용하여 표현영역에서 값 검사가 가능하다.

- 표현영역: 필수 값, 값의 형식, 범위 등을 검증한다
- 응용 서비스: 데이터의 존재 유무와 같은 논리적 오류를 검증한다

응용 서비스를 실행하는 주체가 표현 영역이면 응용 서비스는 논리적 오류 위주로 값을 검증해도 문제 없지만,
응용 서비스를 실행하는 주체가 다양하면 응용 서비스에서 반드시 파라미터로 전달받은 값이 올바른지 검사해야 한다.

### 권한 검사

표현 영역에서 할 수 있는 가장 기본적인 검사는 인증(Authentication)이다.
회원 관련 URL은 인증된 사용자만 접근해야 한다. 이런 접근제어를 하기 좋은 위치가 서블릿 필터이다.
인증여부 뿐만아니라 권한(Authorization)에 대해서도 동일한 방식으로 필터를 사용해서 URL별 검사를 할 수 있다.
스프링 시큐리티는 이와 유사한 방식으로 필터를 이용해서 인증 정보를 생성하고 웹 접근을 제어한다.

URL만으로 접근 제어를 할 수 없는 경우 응용서비스의 메소드 단위로 권한검사를 수행해야 한다.
반드시 응용서비스의 코드에서 직접 권한 검사를 수행할 필요는 없다.
스프링 시큐리티의 경우 AOP를 활용한 권한검사를 할 수 있다. `@PreAuthorize("hasRole('admin')")`

개별 도메인 단위로 권한 검사를 하는 경우는 복잡해 질 수 있다.
도메인 객체 수준의 권한 검사 로직은 도메인별로 다르므로 도메인에 맞게 보안 프레임워크를 확장하려면 프레임워크에 대한 이해도가 필요하다.

### 조회 전용 기능과 응용 서비스

서비스에서 조회 전용 기능을 사용하게 되면 서비스 코드가 다음과 같이 단순히 조회 전용 기능을 호출하는 것으로 끝나는 경우가 많다.

```java
public class OrderListService {
  public List<OrderView> getOrderList(String ordererId) {
    return orderViewDao.selectByOrderer(ordererId);
  }
}
```

서비스에서 수행하는 추가적인 로직이 없을뿐더러 조회 전용 기능이어서 트랜잭션이 필요하지도 않다.
이런 경우라면 굳이 서비스를 만들 필요 없이 표현 영역에서 바로 조회전용 기능을 사용할 수도 있다.

## 도메인 서비스

### 여러 애그리거트가 필요한 기능

도메인 영역의 코드를 작성하다 보면 한 애그리거트로 기능을 구현할 수 없을 때가 있다.

그 예로 결제 금액 계산 로직을 들 수 있다.

- 상품 애그리거트: 구매하는 상품 가격, 배송비
- 주문 애그리거트: 상품별 구매 개수
- 할인 쿠폰 애그리거트: 쿠폰별로 금액, 비율 할인, 제약조건
- 회원 애그리거트: 회원 등급에 따라 추가할인 가능

이런 경우 어떤 애그리거트가 결제 금액 계산의 책임을 갖고 있을까?
생각해 볼 수 있는 방법은 주문 애그리거트가,
필요 애그리거트와 필요 데이터를 모두 가지도록 한 뒤 할인 금액 계산 책임을 주문 애그리거트에 할당하는 것이다.

```java
public class Order {
  ...
  private Orderer orderer;
  private List<OrderLine> orderLines;
  private List<Coupon> usedCoupons;

  private Money calculatePayAmounts() {
    ...
  }

  private Money calculateDiscount(Coupon coupon) {
    ...
  }

  private Money calculateDiscount(MemberGrade grade) {
    ...
  }
}
```

이렇게 한 애그리거트에 넣기 애매한 도메인 기능을 특정 애그리거트에 억지로 구현하면 안된다.
이 경우 애그리거트는 책임 범위를 넘어서는 기능을 구현하기 때문에 코드가 길어지고 외부 의존이 높아지게 된다.
또한 애그리거트의 범위를 넘어서는 도메인 개념이 애그리거트에 숨어들어 명시적으로 드러나지 않게된다.

이런 문제를 해결하기 위해 도메인 서비스를 별도로 구현할 수 있다.

### 도메인 서비스 구현

한 애그리거트에 넣기 애매한 도메인 개념을 구현하려면 도메인 서비스를 이용해서 도메인 개념을 명시적으로 드러내면 된다.
응용 영역의 서비스가 응용 로직을 다룬다면 도메인 서비스는 도메인 로직을 다룬다.

도메인 서비스가 도메인 영역의 애그리거트나 밸류 같은 구성요소와 다른점은 상태 없이 로직만 구현한다는 점이다.
도메인 서비스를 구현하는데 필요한 상태는 애그리거트나 다른 방법으로 전달받는다.

할인 금액 계산 로직을 위한 도메인 서비스는 다음과 같이 도메인의 의미가 드러나는 용어를 타입과 메서드 이름으로 갖는다.

```java
public class DiscountCalculationService {
  public Money calculateDiscountAmoutns(List<OrderLine> orderLines, List<Coupon> coupons, MemberGrade grade) {
    Money couponDiscount = coupons.stream()
            .map(coupon -> calculateDiscount(coupon))
            .reduce(Money(0), (v1, v2) -> v1.add(v2));
    Money membershipDiscount = calculateDiscount(orderer.getMember().getGrade());
    return couponDiscount.add(membershipDiscount);
  }

  private Money calculateDiscount(Coupon coupon) {
    ...
  }

  private Money calculateDiscount(MemberGrade grade) {
    ...
  }
}
```

할인 계산 서비스를 사용하는 주체는 애그리거트가 될 수도 있고 응용 서비스가 될 수도 있다.
`DiscountCalculationService`를 다음과 같이 애그리거트의 결제 금액 계산 기능에 전달하면 사용주체는 애그리거트가 된다.

```java
public class Order {
  public void calculateAmoutns(DiscountCalculationService disCalSvc, MemberGrade grade) {
    Money totalAmounts = getTotalAmounts();
    Money discountAmounts = disCalSvc.calculateDiscountAmoutns(this.orderLines, this.coupons, grade);
    this.paymentAmounts = totalAmounts.minus(discountAmounts);
  }
}
```

애그리거트 객체에 도메인 서비스를 전달하는 것은 응용 서비스 책임이다.

```java
public class OrderService {
  private DiscountCalculationService discountCalculationService;

  @Transactional
  public OrderNo placeOrder(OrderRequest orderRequest) {
    OrderNo orderNo = orderRepository.nextId();
    Order order = createOrder(orderNo, orderRequest);
    orderRepository.save(order);
    ...
    return orderNo;
  }

  private Order createOrder(OrderNo orderNo, OrderRequest orderReq) {
    Member member = findMemeber(orderReq.getOrdererId());
    Order order = new Order(orderNo, orderReq.getOrderLines(), orderReq.getCoupons(),
            createOrderer(member), orderReq.getShippingInfo());
    order.calculateAmounts(this.discountCalculationService, member.getGrade());
    return order;
  }
  ...
}
```

도메인 서비스의 객체를 애그리거트에 주입하지 않을 것

애그리거트에 도메인 서비스 객체를 파라미터로 전달한다는 것은 애그리거트가 도메인 서비스에 의존한다는 것이다.
애그리거트가 의존하는 도메인 서비스를 DI로 처리하고 싶을지도 모르겠으나 이는 좋은 방법이 아니다.
애그리거트에서 도메인서비스는 일부 기능에서만 사용하고 데이터 자체와 관련이 없다.

애그리거트 메소드를 실행할 때 도메인 서비스를 인자로 전달하지 않고 반대로 도메인 서비스의 기능을 실행할 때 애그리거트를 전달할 수 있다.
계좌이체의 경우 두 계좌 애그리거트가 관여하는데 한 애그리거트는 금액을 출금하고 한 애그리거트는 금액을 입금한다.

```java
public class TransferService {
  public void transfer(Account fromAcc, Account toAcc, Money Amounts) {
    fromAcc.withdraw(amounts);
    toAcc.credit(amounts);
  }
  ...
}
```

응용 서비스는 두 Account 애그리거트를 구한 뒤에 해당 도메인 영역의 `TransferService`를 이용해서 계좌이체 도메인 기능을 실행할 것이다.

도메인 서비스는 도메인 로직을 수행하지 응용 로직을 수행하지는 않으므로, 트랜잭션 처리와 같은 로직은 응용서비스에서 처리해야 한다.

> 특정기능이 도메인서비스인지 알기 위해서는 해당 로직이 애그리거트의 상태를 변경하거나 애그리거트의 상태 값을 계산하는지 보면된다. (해당사항 없으면 응용서비스)

### 도메인 서비스의 패키지 위치

도메인 서비스는 도메인 로직을 실행하므로 도메인 서비스의 위치는 다른 도메인 구성요소와 동일한 패키지에 위치한다.

## 애그리거트 트랜잭션 관리

한 애그리거트를 두 사용자가 거의 동시에 변경할 때 트랜잭션이 필요하다.
운영자와 고객이 동시에 한 주문 애그리거트를 수정하는 상황을 가정해보자.
운영자 스레드와 고객 스레드는 개념적으로 동일한 애그리거트이지만 물리적으로 서로 다른 애그리거트 객체를 사용한다.
따라서 운영자 스레드가 주문 애그리거트 객체 상태를 변경해도 고객 스레드가 사용하는 주문 애그리거트 객체에 영향을 주지 않는다.

이런 문제가 발생하지 않도록 하려면 다음 두가지 중 하나를 해야한다.

- 운영자가 배송정보를 조회하고 상태를 변경하는 동안 고객이 애그리거트를 수정하지 못하게 막는다
- 운영자가 배송지 정보를 조회한 이후에 고객이 정보를 변경하면 운영자가 애그리거트를 다시 조회한 뒤 수정하게 한다

위의 두 가지는 애그리거트 자체의 트랜잭션과 관련있다.
대표적인 트랜잭션 처리방법에는 선점 잠금과 비선점 잠금 두가지가 있다.

### 선점 잠금

선점 잠금은 먼저 애그리거트를 구한 스레드가 애그리거트 사용이 끝날 때까지 다른 스레드가 해당 애그리거트를 수정하는 것을 막는 방식이다.
이런경우 한 스레드는 다른 스레드가 애그리거트에 대한 잠금을 해제할 때 까지 블로킹 된다.

선점잠금은 보통 DBMS가 제공하는 행 단위 잠금을 사용해서 구현한다.
오라클을 비롯한 다수 DBMS가 `for update`와 같은 쿼리를 사용해서 특정레코드에 한 사용자만 접근할 수 있는 잠금을 제공한다.

JPA의 `EntityManager`는 `LockModeType`을 인자로 받는 `find()` 메소드를 제공하는데,
`LockModeType.PESSIMISTIC_WRITE`를 값으로 전달하면 선점 잠금 방식을 적용할 수 있다.
`Order order = entityManager.find(Order.class, orderNo, LockModeType.PESSIMISTIC_WRITE)`

#### 선점잠금과 교착상태

선점 잠금 기능을 사용할 때는 잠금 순서에 따른 deadlock이 발생하지 않도록 주의해야 한다.
선점 잠금에 의한 교착상태는 상대적으로 사용자 수가 많을 때 발생가능성이 높아진다.
이런 문제가 발생하지 않도록 하려면 잠금을 구할 때 최대 대기시간을 지정하여야 한다.

```java
Map<String, Object> hints = new HashMap<>();
hints.put("javax.persistence.lock.timeout", 2000); // ms 단위
Order order = entityManager.find(Order.class, orderNo, LockModeType.PESSIMISTIC_WRITE, hints);
```

지정 시간내에 잠금을 구하지 못하면 예외를 발생시킨다. DBMS에 따라 힌트가 적용되지 않을 수도 있다.

### 비선점 잠금

stateless 인경우 데이터 접근 순서에 따라 마지막 행위가 최신 정보를 반영하지 못할 수도 있다.
이런 경우 선점 잠금 방식으로 해결할 수 없는데, 이때 비선점 잠금을 사용할 수 있다.

비선점 잠금은 변경한 데이터를 실제 DBMS에 반영하는 시점에 변경 가능여부를 확인하는 방식이다.
비선점 잠금을 구현하려면 애그리거트에 버전으로 사용할 숫자 타입의 프로퍼티를 추가해야 한다.
애그리거트를 수정할 때마다 버전으로 프로퍼티의 값이 1씩 증가한다. 이때 다음과 같은 쿼리가 적용된다.

```sql
UPDATE aggtable SET version = version + 1, colx = ?, coly = ?
WHRER aggid = ? and version = 현재버전
```

JPA는 버전을 이용한 비선점 잠금기능을 지원한다.
버전으로 사용할 필드에 `@Version` 애노테이션을 붙이고 매핑되는 테이블에 버전을 저장할 칼럼을 추가하면 된다.

```java
@Entity
public class Order {
  @EmbeddedId
  private OrderNo number;
  @Version
  private long version;
}
```

JPA는 엔티티가 변경되어 `UPDATE` 쿼리를 실행할 때 `@Version`에 명시한 필드를 이용해 비전점 잠금쿼리를 실행한다.

응용서비스는 버전에 대해 알 필요가 없이 기능만 실행하면 된다.

```java
public class ChangeShippingService {
  @Transactional
  public void changeShipping(ChangeShippingRequest changeReq) {
    Order order = orderRepository.findById(new OrderNo(changeReq.getNumber()));
    checkNoOrder(order);
    order.changeShippingInfo(changeReq.getShippingInfo());
  }
  ...
}
```

비선점 잠금 쿼리를 실행할 때 실행결과로 수정된 행 개수가 0이면 트랜잭션이 충돌한 것이다.
트랜잭션 충돌이 발생하면 `OptimisticLockingFailureException`이 발생한다.

#### 강제 버전 증가

애그리거트에 애그리거트 루트 외 다른 엔티티가 존재하는데 기능 실행 도중 루트가 아닌 다른 엔티티의 값만 변경된다고 하자.
연관된 엔티티의 값이 변경된다고 해도 루트 엔티티의 값은 바뀌는 것이 없으므로 루트 엔티티의 버전이 갱신되지 않는다.

JPA는 이런 문제를 처리할 수 있도록 `EntityManager.find()` 메소드로 엔티티를 구할 때 강제로 버전 값을 증가시키는 잠금모드를 지원한다.

```java
@Repository
public class JpaOrderRepository implements OrderRepository {
  @PersistenceContext
  private EntityManager entityManager;

  @Override
  public Order findByIdOptimisticLockMode(OrderNo id) {
    return entityManager.find(Order.class, id, LockModeType.OPTIMISTIC_FORCE_INCREMENT);
  }
  ...
}
```

`LockModeType.OPTIMISTIC_FORCE_INCREMENT`를 사용하면 해당 엔티티의 상태 변경여부와 관계없이 트랜잭션 종료시점에 버전 값 증가처리를 한다.

### 오프라인 선점 잠금

오프라인 선점 잠금방식은 여러 트랜잭션에 걸쳐 동시 변경을 막는다.
누군가 수정화면을 보고 있을 때 수정 화면 자체를 실행 못하도록 만드는 방식을 예로 들 수 있다.

#### `LockManager` 인터페이스와 관련 클래스

오프라인 선점 잠금은 잠금 선점 시도, 잠금 확인, 잠금 해제, 락 유효시간 연장의 네가지 기능을 제공해야 한다.

오프라인 선점잠금이 필요한 코드는 `LockManager.tryLock()`을 이용해서 잠금선점을 시도한다.
잠금 선점에 성공하면 `LockId`를 리턴한다. `LockId`는 잠금 해제시 사용된다.

잠금 유효시간이 지났으면 다른 사용자가 잠금 선점을 할 수 있다.

잠금을 선점하지 않은 사용자가 기능을 실행한다면 기능실행을 막아야 한다.

#### DB를 이용한 `LockManager` 구현

잠금 정보 테이블

```sql
CREATE TABLE locks (
  `type` varchar(255),
  id varchar(255),
  lockid varchar(255),
  expiration_time datetime,
  primary key (`type`, id)
) character set utf8;

CREATE UNIQUE INDEX locks_idx ON locks (lockid);
```

## 도메인 모델과 Bounded Context

처음 도메인 모델을 만들 때 빠지기 쉬운 함정이 도메인을 완벽하게 표현하는 단일 모델을 만드는 시도를 하는 것이다.
한 도메인은 여러 하위도메인으로 구분되기 때문에 한 개의 모델로 완벽히 표현할 수 없다.

예를 들어 상품이라는 모델을 생각해보자.
카탈로그에서 상품, 재고관리에서 상품, 주문에서 상품, 배송에서 상품은 이름만 같고 의미하는 바가 다르다.
카탈로그에서 상품은 상품 이미지, 상품명, 상품가격, 옵션목록, 상세 설명과 같은 상품 정보 위주이다.
재고 관리에서 상품은 개별 상품 객체를 추적하기 위한 목적으로 상품을 사용한다.
즉 카탈로그에서 물리적으로 하나인 상품이 재고 관리에서는 여러개 존재할 수 있다.

하위 도메인마다 사용하는 용어가 다르므로 올바른 도메인 모델을 개발하려면 하위 도메인마다 모델을 만들어야 한다.
각 모델은 명시적으로 구분되는 경계를 가져서 섞이지 않도록 해야 한다.
이렇게 구분되는 경계를 DDD에서 Bounded Context라고 부른다.

### BOUNDED CONTEXT

Bounded Context는 모델의 경계를 결정하며 하나의 컨텍스트에 논리적으로 하나의 모델을 갖는다.
여러 하위 도메인을 하나의 Bounded Context에서 개발할 때 하위 도메인의 모델이 섞이지 않도록 해야한다.

#### BOUNDED CONTEXT의 구현

컨텍스트가 도메인 모델만 포함하는 것은아니다.
도메인 모델 뿐만 아니라 사용자에게 기능을 제공하기 위한 표현영역, 응용서비스, 인프라 영역 모두를 포함한다.

### BOUNDED CONTEXT간 관계

Bounded Context는 연결되기 때문에 두 컨텍스트는 다양한 방식으로 관계를 맺는다.
가장 흔한 관계는 한쪽에서 API를 제공하고 다른쪽에서 API를 호출하는 관계이다.

upstream conponent는 서비스 공급자 역할을하고 downstream conponent는 서비스를 사용하는 고객 역할을 한다.
하류 컴포넌트는 상류 서비스의 모델이 자신의 도메인 모델에 영향을 주지 않기 위해 보호용의 완충지대를 만들어야 한다.
이 계층을 Anticorruption Layer라 부른다.

## 이벤트

### 시스템 간 강결합 문제

쇼핑몰에서 구매 취소시 환불을 처리해야 한다.
보통 결제 시스템은 외부에 존재하므로 외부 환불시스템 서비스 호출 시 몇 가지 문제가 발생한다.

첫째, 외부 서비스가 정상이 아닐 경우 트랜잭션 처리 문제이다.
외부의 환불 서비스를 실행하는 과정에서 예외가 발생하면 환불에 실패했으므로 주문 취소 트랜잭션을 롤백해야한다.
하지만 반드시 트랜잭션을 롤백해야 하는 것은 아니다. 주문은 취소상태로 변경하고 환불만 나중에 다시 시도할 수도 있다.

두번째, 환불을 처리하는 외부시스템 응답시간이 길어지면 대기시간도 길어지는 성능문제이다.

셋째 도메인 객체에 서비스를 전달하면 설계상 문제가 나타날 수 있다.

지금까지의 문제는 주문과 결제 두 Bounded Context가 강결합이기 때문에 발생한다.
강결합을 약화시키기 위해서 이벤트를 사용할 수 있다.
특히 비동기 이벤트를 사용하면 두 시스템간의 결합도를 크게 낮출 수 있다.

### 이벤트 개요

이벤트가 발생한다는 것은 상태가 변경되었다는 것을 의미한다.
이벤트는 발생에서 끝나지 않고, 이벤트에 반응하여 원하는 동작을 수행하는 기능을 구현할 수 있다.

#### 이벤트 관련 구성요소

이벤트를 위해서는 다음의 네 구성요소를 구현해야 한다.

- 이벤트 생성 주체
- 이벤트 디스패처 (이벤트 퍼블리셔)
- 이벤트 핸들러 (이벤트 구독자)

도메인 모델에서 도메인 객체는 로직을 실행해서 상태가 바뀌면 관련 이벤트를 발생시킨다.
이벤트 핸들러는 이벤트 생성 주체가 발생한 이벤트에 반응한다.
이벤트 생성 주체와 이벤트 핸들러를 연결해 주는 것이 이벤트 디스패처이다.
이벤트 생성주체는 이벤트를 생성해서 디스페처에 이벤트를 전달하고 디스페처는 이벤트를 핸들러에 전파한다.

#### 이벤트 구성

- 이벤트 종류: 클래스 이름으로 표현
- 이벤트 발생 시각
- 추가 데이터: 주문번호, 신규 배송지등 이벤트 관련 정보

배송지 변경 이벤트를 생각해보자

```java
public class ShippingInfoChangedEvent {
  private String orderNumber;
  private long timestamp;
  private ShippingInfo newShippingInfo;
  ...
}
```

이벤트 발생 주체는 `Order` 애그리거트이다.

```java
public class Order {
  public void changeShippingInfo(ShippingInfo newShippingInfo) {
    verifyNotYetShipped();
    seShippingInfo(newShippingInfo);
    Events.raise(new ShippingInfoChangedEvent(number, newShippingInfo));
  }
  ...
}
```

이벤트를 처리하는 핸들러는 다음과 같이 구성할 수 있다.

```java
public class ShippingInfoChangedHandler implements EventHandler<ShippingInfoChangedEvent> {
  @Override
  public void handle(ShippingInfoChangedEvent evt) {
    // 이벤트에 데이터가 없다면 직접 데이터를 조회해야 한다.
    Order order = orderRepository.findById(evt.getOrderNo());

    shippingInfoSynchronizer.sync(evt.getOrderNumber(), evt.getNewShippingInfo());
  }
  ...
}
```

#### 이벤트 용도

이벤트는 크게 두가지 용도로 쓰인다.

첫 번째 용도는 트리거이다. 도메인의 상태가 바뀔 때 다른 후처리를 해야할 경우 이벤트를 사용할 수 있다.
주문취소의 경우 환불 처리를 예로 들 수 있다. 환불 처리를 위한 트리거로 주문 취소 이벤트를 사용할 수 있다.

이벤트의 두 번째 용도는 다른 시스템간의 데이터 동기화이다.
배송지를 변경하면 외부 배송 서비스에 바뀐 배송지 정보를 전송해야 한다.

#### 이벤트의 장점

이벤트를 사용하면 서로 다른 도메인 로직이 섞이는 것을 방지할 수 있다. (도메인에 서비스가 주입되는 경우)

```java
public class Order {
  public void cancel(RefundService refundSvc) {
    verifyNotYetShipped();
    this.state = OrderState.CANCELED;
    this.refundStatus = State.REFUND_STARTED;
    try {
      refundSvc.refund(getPaymentId());
      this.refundStatus = State.REFUND_COMPLETED;
    } catch (Exception ex) {
      ...
    }
  }
}
```

위와 같은 코드를 아래로 변경할 수 있다.

```java
public class Order {
  public void cancel() {
    verifyNotYetShipped();
    this.state = OrderState.CANCELED;
    this.refundStatus = State.REFUND_STARTED;
    Events.raise(new OrderCanceledEvent(number.getNumber()));
  }
}
```

이벤트 핸들러를 사용하면 기능확장도 용이하다.
다른 기능을 처리하는 핸들러를 구현하고 디스패처에 등록하면된다. 도메인 로직은 수정할 필요가 없다.

### 이벤트, 핸들러, 디스패처 구현

#### 이벤트 클래스

이벤트를 위한 상위타입은 존재 하지 않으므로 원하는 클래스를 이벤트로 사용한다.
이벤트 클래스는 이벤트를 처리하는 데 필요한 최소한의 데이터를 포함해야 한다.

```java
public class OrderCanceledEvent {
  // 이벤트는 핸들러에서 이벤트를 처리하는 데 필요한 데이터를 포함한다.
  private String orderNumber;
  public OrderCanceledEvent(String number) {
    this.orderNumber = number;
  }
  ...
}
```

#### `EventHandler` 인터페이스

`EventHandler` 인터페이스는 이벤트 핸들러를 위한 상위 인터페이스이다.

```java
import net.jodah.typetools.TypeResolver;

public interface EventHandler<T> {
  void handle(T event);

  default boolean canHandle(Object event) {
    Class<?>[] typeArgs = TypeResolver.resolveRawArguments(EventHandler.class, this.getClass());
    return typeArgs[0].isAssignableFrom(event.getClass());
  }
}
```

#### 이벤트 디스패처인 `Events` 구현

도메인을 사용하는 응용 서비스는 이벤트를 받아 처리할 핸들러를 Events.handle()로 등록하고, 도메인 기능을 실행한다.

```java
public class CancelOrderService {
  private OrderRepository orderRepository;
  private RefundService refundService;

  @Transactional
  public void cancel(OrderNo orderNo) {
    Events.handle((OrderCanceldedEvent evt) -> {
      refundService.refund(evt.getOrderNumber())
    });
    Order order = findOrder(orderNo);
    order.cancel();
    Events.reset();
  }
}
```

`Events`는 내부적으로 핸들러 목록 유지를 위해 `ThreadLocal`을 사용한다.
`Events.handle()` 메소드는 인자로 전달받은 `EventHandler`를 `List`에 보관한다.
이벤트를 발생시킬 때는 `Events.raise()` 메소드를 사용한다.

```java
public class Order {
  public void cancel() {
    verifyNotYetShipped();
    this.state = OrderState.CANCELED;
    Events.raise(new OrderCanceldedEvent(number.getNumber()));
  }
}
```

`Events` 클래스의 구현 코드는 다음과 같다.

```java
public class Events {
  private static ThreadLocal<List<EventHandler<?>>> handlers = new ThreadLocal<>();
  private static ThreadLocal<Boolean> publishing = new ThreadLocal<Boolean> {
    @Override
    protected Boolean initalValue() {
      return Boolean.FALSE;
    }
  };

  public static void raise(Object event) {
    if (publishing.get()) return;
    try {
      publishing.set(Boolean.TRUE);
      List<EventHandler<?>> eventHandlers = handlers.get();
      if (eventHandlers == null) return;
      for (EventHandler handler : eventHandlers) {
        if (handler.canHandle(event)) {
          handler.handle(event);
        }
      }
    } finally {
      publishing.set(Boolean.FALSE);
    }
  }

  public static void handle(EventHandler<?> handler) {
    if (publishing.get()) return;
    List<EventHandler<?>> eventHandlers = handlers.get();
    if (eventHandlers == null) {
      eventHandlers = new ArrayList<>();
      handlers.set(eventHandlers);
    }
    eventHandlers.add(handler);
  }

  public static void reset() {
    if (!publishing.get()) {
      handlers.remove();
    }
  }
}
```

`Events`는 핸들러 목록을 유지하기 위해 `ThreadLocal` 변수를 사용한다.
톰캣과 같은 WAS는 스레드를 재사용하므로 `ThreadLocal`에 보관한 값을 제거해야 한다. (`reset()` 메소드 호출)

#### 이벤트 처리 흐름

1. 이벤트 처리에 필요한 이벤트 핸들러를 생성한다
2. 이벤트 발생 전에 이벤트 핸들러를 `Events.handle()` 메소드로 등록한다
3. 이벤트를 발생하는 도메인 기능을 실행한다
4. 도메인은 `Events.raise()`를 이용하여 이벤트를 발생한다
5. `Events.raise()`는 등록된 핸들러의 `canHandle()`을 이용해서 이벤트를 처리할 수 있는지 확인한다
6. 핸들러가 이벤트를 처리할 수 있다면 `handle()` 메소드를 이용해서 이벤트를 처리한다
7. `Events.raise()` 실행을 끝내고 리턴한다
8. `Events.reset()`을 이용해서 `ThreadLocal`을 초기화한다

### 동기 이벤트 처리 문제

외부 서비스가 느려지만 내부 서비스의 성능저하가 발생하는 문제를 해결해야 한다.
성는 뿐만아니라 트랜잭션도 문제가 된다.

외부 시스템과 연동을 동기로 처리할 때 발생하는 성능과 트랜잭션 범위 문제를 해소하는 방법중 하나는 이벤트를 비동기로 처리하는 것이다.

### 비동기 이벤트 처리

회원 가입 검증시 이메일을 보내는 서비스나, 주문취소시 결제 취소가 이루어지는 서비스는 즉시 발생하지 않아도 문제가 없다.

이를 비동기로 구현하려면 A 이벤트가 발생하면 별도 스레드로 B를 수행하는 핸들러를 실행할 수 있따.

#### 로컬 핸들러의 비동기 실행

별도 스레드를 이용해서 이벤트 핸들러를 실행하면 이벤트 발생코드와 같은 트랜잭션 범위에 묶을 수 없기 때문에,
한 트랜잭션으로 실행해야 하는 이벤트 핸들러는 비동기로 처리하면 안된다.

> 스프링의 트랜잭션 관리자는 일반적으로 스레드를 이용하여 트랜잭션을 전파한다.

#### 메시징 시스템을 이용한 비동기 구현

비동기로 이벤트를 처리해야 할 때 사용하는 또 다른 방법은 메시징 큐를 사용하는 것이다.

- 이벤트가 발생하면 이벤트 디스패처는 이벤트를 메시지 큐에 보낸다
- 메시지 큐는 이벤트를 메시지 리스너에 전달한다
- 메시지 리스너는 알맞은 이벤트 핸들러를 이용해서 이벤트를 처리한다

이때 이벤트를 메시지 큐에 저장하는 과정과 메시지 큐에서 이벤트를 읽어와 처리하는 과정은 별도 스레드나 프로세스로 처리된다.

필요하다면 이벤트를 발생하는 도메인 기능과 메시지 큐에 이벤트를 저장하는 절차를 한 트랜잭션에 묶어야 한다.
이를 위해서는 글로벌 트랜잭션이 필요한데, 안전하게 이벤트를 메시지큐에 전달하겠지만 성능이 떨어지는 단점이 있다.

많은경우 이벤트 발생 주체와 이벤트 핸들러가 별도 프로세스에 동작한다.
자바의 경우 이벤트 발생 JVM과 이벤트 처리 JVM 이 다르다는 것을 의미한다.

메시지 전달을 위해 많이 사용되는 것중에 Kafka도 있다.
Kafka는 글로벌 트랜잭션을 지원하지는 않지만 다른 메시징 시스템에 비해 높은 성능을 보여준다.

#### 이벤트 저장소를 이용한 비동기 처리

이벤트를 처리하는 방법중 하나는 이벤트를 DB에 저장한 뒤 이벤트 핸들러에 전달하는 것이다.
이 방식은 도메인의 상태와 이벤트 저장소로 동일한 DB를 사용하므로 도메인의 상태변화와 이벤트 저장이 로컬 트랜잭션으로 처리된다.

다른 방법은 이벤트를 외부에 제공하는 API를 사용하는 것이다.

## CQRS

조회화면의 특성상 속도가 빠를수록 좋은데 여러 애그리거트에서 데이터를 가져와야 할 경우 구현방법을 고민해야 한다.
이런 구현 복잡도를 낮추는 방법이 있는데 상태변경을 위한 모델과 조회를 위한 모델을 분리하는 것이다.

### CQRS 정의

도메인 모델 관점에서 상태 변경 기능은 주로 한 애그리거트의 상태를 변경한다.
상태를 변경하는 범위와 상태를 조회하는 범위가 정확하게 일치하지 않기 때문에 단일 모델로 두 기능을 구현하면 모델이 복잡해지게 된다.

CQRS는 Command Query Responsibility Segregation의 약자로 상태를 변경하는 Command 모델과 상태를 조회하는 Query 모델을 분리하는 것이다.

CQRS를 사용하면 각 모델에 맞는 구현 기술을 선택할 수 있다.
명령모델은 객체지향 기반의 JPA를 사용하고, 조회모델은 MyBatis를 사용하여 구현할 수 있다.

서로 다른 데이터 저장소를 사용할 수도 있는데 명령모델은 트랜잭션을 지원하는 RDMBS를 사용하고,
조회모델은 조회성능이 좋은 메모리기반 NoSQL을 사용할 수도 있다.

명령 모델과 조회 모델이 같은 기술을 사용할 수도 있다.
JPQL을 이용한 동적 인스턴트 생성과 하이버네이트의 `Subselect`를 이용할 수 있다.

### CQRS의 장단점

장점

- 명령 모델을 구현할 때 도메인 자체에 집중할 수 있다
- 명령 모델에서 조회 관련 로직이 사라져 복잡도가 낮아진다
- 조회 성능을 향상시키는데 유리하다

단점

- 구현해야 할 코드가 더 많아진다
- 더 많은 구현기술이 필요하다
