Brian Goff and Luke Marsden paired on an MVP of the plugin handshake interface.

# pre-MVP design

We ended up with the following design:

```
$ docker run --plugin \
    -v /var/run/docker.sock:/var/run/docker.sock \
    clusterhq/flocker-plugin
```

Internally this translates into:

```
$ docker run -d -e $ARGS \
	-v /var/lib/docker/containers/<container_id>/plugin/:/var/run/docker/ \
	-v /var/run/docker.sock:/var/run/docker.sock \
	clusterhq/flocker-plugin $@
```

Docker blocks on waiting for a successful handshake with the plugin on `/var/run/docker/plugin.sock` with an HTTP `POST /v1/handshake` with an empty body.

# Plugin API

## `POST /v1/handshake`

**Request:** POST with empty body

**Response:** `application/json` as follows:

```
{
 InterestedIn: ["volume"],
 PluginName: "flocker",
 PluginAuthor: "Luke Marsden <luke@clusterhq.com>",
 PluginOrg: "ClusterHQ, Inc.",
 PluginWebsite: "https://clusterhq.com/",
}
```

PluginName is a human readable short string identifying the plugin.

InterestedIn is a list of extension points.

Extension points currently supported are:

* volume - called when a volume is bind-mounted into a container at create or start time

Planned extension points include:

* api - pre-hooks and post-hooks on the Docker remote API (powerstrip-style)
* network - called when ???

## `POST /v1/volume/volumes`

The simplest of plugin types, `volume` provides a single request-response type on `POST /v1/volume/volumes` (`POST ~= create`):

**Request**

```
{HostPath: "/path",
 ContainerID: "abcdef123"}
```

**Response**

```
{HostPath: "/newpath",
 ContainerID: "abcdef123"}
```

# Caveats

NB: plugin loading order is undefined, so plugins should not depend on eachother.

---

TODO

* sort out plugins from normal containers and load plugins first so that containers that want to use services (e.g. portable volumes) from containers get to consume those services from other plugins.
