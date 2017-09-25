# Spring 설정 Java편

<https://docs.spring.io/spring/docs/current/spring-framework-reference/htmlsingle/#beans-java>

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

```java
public class WebApplicationInitializer extends AbstractAnnotationConfigDispatcherServletInitializer {
  // ContextLoaderListener (contextConfigLocation 인식)
  @Override
  protected Class<?>[] getRootConfigClasses() {
    return new Class<?>[]{RootContextConfig.class};
  }

  @Override
  protected Class<?>[] getServletConfigClasses() {
    return new Class[] { ServletContextConfig.class };
  }

  // URL pattern
  @Override
  protected String[] getServletMappings() {
    return new String[] { "/" };
  }

  /* Encoding Filter(POST) */
  @Override
  protected Filter[] getServletFilters() {
    CharacterEncodingFilter characterEncodingFilter = new CharacterEncodingFilter();
    characterEncodingFilter.setEncoding("UTF-8");
    return new Filter[]{characterEncodingFilter, new HiddenHttpMethodFilter()};
  }
}
```

### RootContextConfig(root-context.xml)

`@Configuration`를 통해 설정 클래스임을 확인

```java
@Configuration
@PropertySource("classpath:application.properties")
@ComponentScan(basePackages = { "kr.pravusid" })
@MapperScan(basePackages = "kr.pravusid.persistence", annotationClass = Mapper.class)
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

  // Spring MyBatis 설정
  @Bean
  public SqlSessionFactory sqlSessionFactory() throws Exception {
    SqlSessionFactoryBean sqlSessionFactory = new SqlSessionFactoryBean();
    sqlSessionFactory.setDataSource(dataSource());
    return (SqlSessionFactory) sqlSessionFactory.getObject();
  }

  /* @MapperScan을 사용한다면 bean 등록필요 X
  @Bean
  public UserMapper userMapper() throws Exception {
    SqlSessionTemplate sessionTemplate = new SqlSessionTemplate(sqlSessionFactory());
    return sessionTemplate.getMapper(UserMapper.class);
  }
  */

  // 트랜잭션 설정
  @Bean
  public PlatformTransactionManager transactionManager() {
    return new DataSourceTransactionManager(dataSource());
  }
}
```

### ServletContextConfig(servlet-context.xml)

```java
@Configuration
@EnableWebMvc
public class ServletContextConfig extends WebMvcConfigurerAdapter {
  // 리소스 처리
  @Override
  public void addResourceHandlers(ResourceHandlerRegistry registry) {
    registry.addResourceHandler("/resources/**").addResourceLocations("/resources/");
  }

  // Resonponse Body (jackson bean, 인코딩처리:org.springframework.http.converter.StringHttpMessageConverter)
  @Override
  public void configureMessageConverters(List<HttpMessageConverter<?>> converters) {
    Jackson2ObjectMapperBuilder builder = new Jackson2ObjectMapperBuilder();
    builder.dateFormat(new SimpleDateFormat("yyyy-MM-dd HH:mm"));
    builder.featuresToDisable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
    converters.add(new MappingJackson2HttpMessageConverter(builder.build()));
  }

  @Bean
  public HttpMessageConverter<String> responseBodyConverter() {
    return new StringHttpMessageConverter(Charset.forName("UTF-8"));
  }

  // ViewResolver
  @Bean
  public InternalResourceViewResolver jstlViewResolver(){
    InternalResourceViewResolver viewResolver = new InternalResourceViewResolver();
    viewResolver.setViewClass(JstlView.class);
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
}
```
