# Docker Build

<https://docs.docker.com/build/>

## Multi Stage

<https://docs.docker.com/build/building/multi-stage/>

## Secrets

- <https://docs.docker.com/build/building/secrets/>
- <https://docs.npmjs.com/docker-and-private-modules>

## Buildx

docker-buildx is a Docker plugin. For Docker to find the plugin, add "cliPluginsExtraDirs" to `~/.docker/config.json`:

```json
{
  "cliPluginsExtraDirs": ["/opt/homebrew/lib/docker/cli-plugins"]
}
```
