Brian Goff and Luke Marsden paired on an MVP of the plugin handshake interface.

We ended up with the following design:

```
$ docker run --plugin \
    -v /var/run/docker.sock:/var/run/docker.sock \
    clusterhq/flocker-plugin
```

Internally this translates into:

```
$ docker run -d -e $ARGS \
	-v /var/lib/docker/containers/<container_id>/plugin/plugin.sock:/var/run/docker/plugin.sock \
	-v /var/run/docker.sock:/var/run/docker.sock \
	clusterhq/flocker-plugin $@
```

Docker blocks on waiting for a successful handshake with the plugin.

