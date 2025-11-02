# Docker Network

<https://docs.docker.com/engine/network/>

## Drivers

- none: Completely isolate a container from the host and other containers.
- bridge: The default network driver.
- host: Remove network isolation between the container and the Docker host.
- overlay: Swarm Overlay networks connect multiple Docker daemons together.
- ipvlan: Connect containers to external VLANs.
- macvlan: Containers appear as devices on the host's network.

## Port publishing and mapping

<https://docs.docker.com/engine/network/port-publishing/>

- `-p 8080:80`: host(8080) to container(80)
- `-p 192.168.1.100:8080:80`: host(192.168.1.100:8080) to container(80)
- `-p 8080:80/tcp -p 8080:80/udp` host(tcp/8080) to container(tcp/80) & host(udp/8080) to container(udp/80)

> 컨테이너 포트를 게시하는 것은 기본적으로 안전하지 않습니다. 즉, 컨테이너의 포트를 게시하면 Docker 호스트뿐만 아니라 외부 세계에서도 접근할 수 있게 됩니다.
> 게시 플래그에 localhost IP 주소(127.0.0.1 또는 ::1)를 포함하면 Docker 호스트만 게시된 컨테이너 포트에 액세스할 수 있습니다.
>
> `docker run -p 127.0.0.1:8080:80 -p '[::1]:8080:80' nginx`
