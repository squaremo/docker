# Experimental: Docker network driver plugins

Docker supports network driver plugins via 
[LibNetwork](https://github.com/docker/libnetwork). Network driver plugins are 
implemented as "remote drivers" for LibNetwork, which shares plugin 
infrastructure with Docker. In effect this means that network driver plugins 
are activated in the same way as other plugins, and use the same kind of 
protocol.

## Using network driver plugins

The means of installing and running a network driver plugin will depend on the
particular plugin.

Once running however, network driver plugins are used just like the built-in
network drivers: by being mentioned as a driver in network-oriented Docker
commands. For example,

    docker network create -d weave mynet

(assuming the [Weave](https://github.com/weaveworks/docker-plugin) plugin is
currently running). Other plugins are listed in [plugins.md](plugins.md)

The network thus created is owned by the plugin, so subsequent commands
referring to that network will also be run through the plugin.

Plugins might provide their own tools e.g. for checking the status of the
plugin or for working with network policy. For example 
[Calico](https://github.com/metaswitch/calico-docker) provides a calicoctl
command for this purpose.

## Network driver plugin protocol

The network driver protocol, additional to the plugin activation call, is
documented as part of LibNetwork:
[https://github.com/docker/libnetwork/blob/master/docs/remote.md](https://github.com/docker/libnetwork/blob/master/docs/remote.md).

# Related GitHub PRs and issues

 - [#13441](https://github.com/docker/docker/pull/13441) Networks API & UI

Send us feedback and comments on the usual Google Groups, or the IRC channel
#docker-network.
