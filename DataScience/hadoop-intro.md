# Hadoop 기초

## Hadoop 실행

- 하둡 초기화 `hadoop namenode -format`

- 하둡 서버 실행 / 중지

  ```sh
  start-all.sh
  stop-all.sh
  ```

- jps : 하둡 작동 프로세서 확인

## Hadoop 명령어

- 하둡에 파일 전송 : `hadoop fs -appendToFile /home/hostname/{local} /{hadoop}`
- 하둡에서 다운로드 : `hadoop -copyToLocal`
- 하둡 파일 보기 : `hadoop fs -cat /path`
- 하둡 파일삭제 : `hadoop fs -rmr /path`
- 하둡 폴더생성 : `hadoop fs -mkdir /path`
- 하둡에 파일 삭제 : `hadoop fs rm -r /test_ns1`

## HADOOP 환경설정

### openssh, java 설치

  ```sh
  sudo apt install openssh-server
  ssh-keygen -t rsa
  sudo apt install openjdk-7-jdk`
  ```

### java path 설정

/etc/environment : 아래의 코드 추가

  ```sh
  PATH="/usr/lib/jvm/java-7-openjdk-amd64"
  JAVA_HOME="/usr/lib/jvm/java-7-openjdk-amd64/
  ```

### /etc/hosts

  ```text
  127.0.0.1 localhost localhost.localdomain localhost4 localhost4.localdomain4
  ::1 localhost localhost.localdomain localhost6 localhost6.localdomain6

  {ip_address} ns1 zookeeper1 NameNode
  {ip_address} ns2 zookeeper2 SecondaryNode
  {ip_address} ns3 zookeeper3 DataNode01
  {ip_address} ns4 zookeeper4 DataNode02
  {ip_address} ns5 zookeeper5 DataNode03
  {ip_address} ns6 zookeeper6 DataNode04
  {ip_address} ns7 zookeeper7 DataNode05
  {ip_address} ns8 zookeeper8 DataNode06

  # The following lines are desirable for IPv6 capable hosts
  ::1     ip6-localhost ip6-loopback
  fe00::0 ip6-localnet
  ff00::0 ip6-mcastprefix
  ff02::1 ip6-allnodes
  ff02::2 ip6-allrouters
  ```

### Hadoop 관련설정 수정

- /etc/environment

경로에 맞춰 아래 내용 추가

  ```sh
  HADOOP_HOME="/usr/local/hadoop-2.7.3"
  HADOOP_PREFIX="/usr/local/hadoop-2.7.3"
  HADOOP_MAPRED_HOME="/usr/local/hadoop-2.7.3"
  HADOOP_COMMON_HOME="/usr/local/hadoop-2.7.3"
  HADOOP_HDFS_HOME="/usr/local/hadoop-2.7.3"
  HADOOP_COMMON_LIB_NATIVE_DIR="/usr/local/hadoop-2.7.3/lib/native"
  HADOOP_OPTS="-Djava.library.path=/usr/local/hadoop-2.7.3/lib"
  HADOOP_CONF_DIR="/usr/local/hadoop-2.7.3/etc/hadoop"
  YARN_HOME="/usr/local/hadoop-2.7.3"
  YARN_CONF_DIR="/usr/local/hadoop-2.7.3/etc/hadoop"
  ```

- ~/.profile

`export PATH="$PATH:$JAVA_HOME/bin:$HADOOP_HOME/bin:$HADOOP_HOME/sbin"`

- core-site.xml

  ```xml
  <configuration>
    <property>
      <name>hadoop.tmp.dir</name>
      <value>/home/hadoop/hdfs/temp</value>
    </property>
    <property>
      <name>fs.default.name</name>
      <value>hdfs://NameNode:9000</value>
    </property>
    <!--
    hive 적용시
    <property>
        <name>hadoop.proxyuser.{server.namenode}.hosts</name> 
        <value>*</value> 
    </property> 
    <property>
        <name>hadoop.proxyuser.{server.namenode}.groups</name>
        <value>*</value>
    </property>
    -->
  </configuration>
  ```

- hdfs-site.xml

  ```xml
  <configuration>
    <property>
      <name>dfs.replication</name>
      <value>3{노드 수}</value>
    </property>
    <property>
      <name>dfs.namenode.name.dir</name>
      <value>/home/hadoop/hdfs/name</value>
    </property>
    <property>
      <name>dfs.datanode.data.dir</name>
      <value>/home/hadoop/hdfs/data</value>
    </property>
  </configuration>
  ```

- mapred-site.xml

  ```xml
  <configuration>
    <!--<property>
    <name>mapred.job.tracker</name>
    <value>NameNode:9001</value>
    </property>-->
    <property>
    <name>mapreduce.framework.name</name>
    <value>yarn</value>
    </property>
    <property>
    <name>mapred.local.dir</name>
    <value>/home/hadoop/hdfs/mapred</value>
    </property>
    <property>
    <name>mapred.system.dir</name>
    <value>/home/hadoop/hdfs/mapred</value>
    </property>
  </configuration>
  ```

- yarn-site.xml

  ```xml
  <configuration>
  <!-- Site specific YARN configuration properties -->
    <property>
      <name>yarn.nodemanager.aux-services</name>
      <value>mapreduce_shuffle</value>
    </property>
    <property>
      <name>yarn.nodemanager.aux-services.mapreduce_suffle.class</name>
      <value>org.apache.hadoop.mapred.ShuffleHandler</value>
    </property>
    <property>
      <name>yarn.resourcemanager.hostname</name>
      <value>NameNode</value>
    </property>
    <property>
      <name>yarn.resourcemanager.resource-tracker.address</name>
      <value>NameNode:8025</value>
    </property>
    <property>
      <name>yarn.resourcemanager.scheduler.address</name>
      <value>NameNode:8030</value>
    </property>
    <property>
      <name>yarn.resourcemanager.address</name>
      <value>NameNode:8040</value>
    </property>
    <property>
      <name>yarn.resourcemanager.webapp.address</name>
      <value>NameNode:8088</value>
    </property>
  </configuration>
  ```

- hadoop-env.sh

  ```sh
  export JAVA_HOME=${JAVA_HOME}
  export HADOOP_PREFIX=${HADOOP_HOME}
  export HADOOP_MAPRED_HOME=${HADOOP_HOME}
  export HADOOP_COMMON_HOME=${HADOOP_HOME}
  export HADOOP_HDFS_HOME=${HADOOP_HOME}
  export YARN_HOME=${HADOOP_HOME}
  export YARN_CONF_DIR="$HADOOP_HOME/etc/hadoop"
  export HADOOP_CONF_DIR="$HADOOP_HOME/etc/hadoop"
  export HADOOP_OPTS="$HADOOP_OPTS -Djava.library.path=$HADOOP_PREFIX/lib/native"
  export HADOOP_HOME_WARN_SUPPRESS="TRUE"
  ```

- yarn-env.sh

  ```sh
  export JAVA_HOME=${JAVA_HOME}
  export HADOOP_PREFIX=${HADOOP_HOME}
  export HADOOP_MAPRED_HOME=${HADOOP_HOME}
  export HADOOP_COMMON_HOME=${HADOOP_HOME}
  export HADOOP_HDFS_HOME=${HADOOP_HOME}
  export YARN_HOME=${HADOOP_HOME}
  export YARN_CONF_DIR="$HADOOP_HOME/etc/hadoop"
  export HADOOP_CONF_DIR="$HADOOP_HOME/etc/hadoop"
  ```

### /etc/hostname을 각자 pc에 맞춰 변경

`sudo vi /etc/hostname`

### 하둡이 사용할 폴더 생성

  ```sh
  hostname@ns1:~$ mkdir -p hadoop/hdfs/name
  hostname@ns1:~$ mkdir -p hadoop/hdfs/data
  hostname@ns1:~$ mkdir -p hadoop/hdfs/mapred
  hostname@ns1:~$ mkdir -p hadoop/hdfs/temp
  ```

- 폴더 소유권한 변경 : `sudo sudo chown hostname hadoop`

- ssh key 일괄 생성(master)

  ```sh
  cd ..sh
  cat id_rsa.pub >> authorized_keys
  more authorized_keys
  ```

- master가 slave의 인증키 등록

  ```sh
  ssh-copy-id user@123.45.56.78
  OR
  ssh hostname@ns1 cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
  ```

- slave에게 생성한 rsa키 전송 : `scp -rp authorized_keys hostname@ns1:~/.ssh/authorized_keys`

- ssh ns1~8 date 로 비밀번호 확인 여부 확인
  known_hosts에 SecondaryNode, DataNode01~06이 등록되어 있지 않으므로 한 번씩 접속해서 known_hosts에 등록한다

### ssh 인증 오류시

`~/.ssh` 내부 삭제 (hosts 재적용)

  ```sh
  rm -rf .ssh
  ssh-keygen -t rsa
  ```
