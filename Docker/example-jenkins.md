# Docker Example: Jenkins

```sh
docker run \
  -u root \
  -p 8080:8080 \
  -v $(pwd)/jenkins:/home \
  -v $(pwd)/jenkins/jenkins-data:/var/jenkins_home \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --name jenkins \
  jenkins/jenkins
```
