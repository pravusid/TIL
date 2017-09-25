# Spring 설정 Java편

## 설정 준비

Maven plugin 추가

```xml
<plugin>
  <groupId>org.apache.maven.plugins</groupId>
  <artifactId>maven-war-plugin</artifactId>
  <configuration>
    <failOnMissingWebXml>false</failOnMissingWebXml>
  </configuration>
</plugin>
```

## AbstractAnnotationConfigDispatcherServletInitializer <- web.xml

WAS가 인식하는 Deployment Descriptor (Servlet 3 이상부터 Java Config으로 지원)

### ContextLoaderListener (contextConfigLocation 인식)

루트 애플리케이션 컨텍스트를 만들어 초기화하고, 애플리케이션과 함께 컨텍스트를 종료시키는 이벤트를 처리하는 리스너로써 애플리케이션 컨텍스트를 구성합니다. 따라서, root-context.xml 말고도 다수의 애플리케이션 컨텍스트 파일을 구성하여 적용할 수도 있습니다. 만약, 애플리케이션의 규모가 커져서 관리해야할 빈이 많아진다면 이러한 애플리케이션 컨텍스트를 모듈별로 나누어 관리하도록 합시다.

@Override
protected Class<?>[] getRootConfigClasses() {
    // TODO Auto-generated method stub
    return new Class<?>[]{RootContextConfig.class};
}

### DispatcherServlet

스프링 웹 MVC에서 지원하는 프론트 컨트롤러 서블릿입니다. 만약, 서블릿 이름을 다르게 지정해주면 애플리케이션에 여러개의 DispatcherServlet을 등록할 수도 있습니다. 각 DispatcherServlet은 서블릿이 초기화될 때 루트 애플리케이션 컨텍스트를 찾아서 자신의 부모 컨텍스트로 사용합니다. 루트 애플리케이션 컨텍스트 처럼 별도의 모듈로 구성할 수도 있지만, 왠만해서는 그럴일이 없습니다.

```java
@Override
protected void registerDispatcherServlet(ServletContext servletContext) {
  WebApplicationContext servletAppContext = createServletApplicationContext();
  DispatcherServlet ds = new DispatcherServlet(servletAppContext);
  ServletRegistration.Dynamic appServlet = servletContext.addServlet("appServlet", ds);
  appServlet.setLoadOnStartup(1);
  appServlet.addMapping(getServletMappings());
}

// URL pattern
@Override
protected String[] getServletMappings() {
  return new String[] { "/" };
}

@Override
protected Class<?>[] getServletConfigClasses() {
  return new Class<?>[]{ServletContextConfig.class};
}
```

### Filter

```java
/* Encoding Filter(POST) */
@Override
protected Filter[] getServletFilters() {
  CharacterEncodingFilter characterEncodingFilter = new CharacterEncodingFilter();
  characterEncodingFilter.setEncoding("UTF-8");
  return new Filter[]{characterEncodingFilter, new HiddenHttpMethodFilter()};
}
```

### RootContextConfig(root-context.xml)

`@Configuration`를 통해 설정 클래스임을 확인

```java
@Import(value={FooConfig.class, BarConfig.class})
@Configuration
@EnableWebMvc
@ComponentScan(basePackages = { "com.idpravus" })
@PropertySource("classpath:application.properties")
@EnableTransactionManagement
public class RootContextConfig {
  @Resource
  private Environment env;

  @Bean
  public DataSource dataSource() {
    // commons DBCP
    BasicDataSource dataSource = new BasicDataSource();
    // tomcat DBCP
    DataSource dataSource = new org.apache.tomcat.jdbc.pool.DataSource();

    dataSource.setDriverClassName(env.getProperty("database.driver"));
    dataSource.setUrl(env.getProperty("database.url"));
    dataSource.setUsername(env.getProperty("database.username"));
    dataSource.setPassword(env.getProperty("database.password"));
    dataSource.setConnectionProperties("true");
    return dataSource;
  }

  // spring.SqlSessionFactoryBean 설정 (Spring MyBatis)

  // 트랜잭션 설정
  @Bean
  public PlatformTransactionManager transactionManager() {
    return new DataSourceTransactionManager(dataSource());
  }

  // JPA 설정
  @Bean
  public static PropertySourcesPlaceholderConfigurer propertySourcesPlaceholderConfigurer() {
      return new PropertySourcesPlaceholderConfigurer();
  }
}
```

### ServletContextConfig(servlet-context.xml)

```java
@Configuration
public class ServletContextConfig extends WebMvcConfigurerAdapter{
  // 컴포넌트 스캔

  // 리소스 처리
  @Override
  public void addResourceHandlers(ResourceHandlerRegistry registry) {
    registry.addResourceHandler("/resources/**").addResourceLocations("/resources/");
  }

  // Resonponse Body (jackson bean, 인코딩처리:org.springframework.http.converter.StringHttpMessageConverter)

  // ViewResolver
  @Bean
  public InternalResourceViewResolver jstlViewResolver(){
    InternalResourceViewResolver viewResolver = new InternalResourceViewResolver();
    viewResolver.setPrefix("/WEB-INF/views/");
    viewResolver.setSuffix(".jsp");
    return viewResolver;
  }

  // Multipart Resolver
  @Bean
  public MultipartResolver multipartResolver(){
    CommonsMultipartResolver cmr = new CommonsMultipartResolver();
    cmr.setMaxUploadSize(env.getProperty("file.maxUploadSize", Long.class));
    return cmr;
  }

  // Apache Tiles
}
```
