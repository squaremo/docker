# Docker Plugins MVP

Plugins provide a mechanism for hooking into the Docker engine to extend its behavior.

HTTP over a local UNIX socket is used to communicate between Docker and its plugins.

Plugins are distributed as containers.

## Usage

Plugins can be loaded using a special `docker plugin-load` command, as follows:

```
$ docker plugin-load volume:flocker clusterhq/flocker-plugin
```

This is really just syntactic sugar for the following:

```
$ docker run -d -e PLUGIN_TYPE=volume -e PLUGIN_NAME=flocker \
	-v /var/run/docker-plugins/flocker.sock:/var/run/plugin.sock \
	clusterhq/flocker-plugin
```

It also marks the plugin container as hidden from `docker ps`.

Docker then waits for the plugin to start listening on the socket, and according to the type of plugin, sends it HTTP requests on certain events.

## Types of plugin

Each plugin type defines a straightforward (MVP) protocol.

### Volume

The simplest of plugin types, `volume` provides a single request-response type:

**Request**

```
{DockerVolumesExtensionVersion: 1,
 Action: "create",
 HostPath: "/path",
 ContainerID: "abcdef123",
}
```

**Response**

```
{DockerVolumesExtensionVersion: 1,
 ModifiedHostPath: "/newpath"} 
```

In the initial version, if the plugin responds with an empty string (`""`) as the `ModifiedHostPath`, the response is ignored.

See reference implementation:
* https://github.com/cpuguy83/docker/compare/ddb366ee9a07e3feab766cc712c9683ad0c3c309
* https://github.com/ClusterHQ/powerstrip-flocker/compare/docker-volume-extension

**Demo**

```
$ git clone https://github.com/clusterhq/powerstrip-flocker
$ cd powerstrip-flocker/vagrant-aws
$ git checkout docker-volume-extension
$ vagrant up --provider=aws
$ vagrant ssh node1
node1$ docker run -v /flocker/test:/data ubuntu sh -c "echo fish > /data/file.txt"
node1$ exit
$ vagrant ssh node2
node2$ docker run -v /flocker/test:/data ubuntu cat /data/file.txt
fish
```

### API



### Network

???

## Possible future functionality

* While waiting for the plugin to initialize, Docker should not allow any API requests to succeed.
  This is to avoid startup race where e.g. `create` requests might not get passed through the plugins.
* Plugins could negotiate with Docker via an initial handshake HTTP request.
* Plugins could use a protocol other than HTTP.