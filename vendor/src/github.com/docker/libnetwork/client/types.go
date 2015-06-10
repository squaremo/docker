package client

import "github.com/docker/libnetwork/types"

/***********
 Resources
************/

// networkResource is the body of the "get network" http response message
type networkResource struct {
	Name      string              `json:"name"`
	ID        string              `json:"id"`
	Type      string              `json:"type"`
	Endpoints []*endpointResource `json:"endpoints"`
}

// endpointResource is the body of the "get endpoint" http response message
type endpointResource struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Network string `json:"network"`
}

/***********
  Body types
  ************/

// networkCreate is the expected body of the "create network" http request message
type networkCreate struct {
	Name        string                 `json:"name"`
	NetworkType string                 `json:"network_type"`
	Options     map[string]interface{} `json:"options"`
}

// endpointCreate represents the body of the "create endpoint" http request message
type endpointCreate struct {
	Name         string                `json:"name"`
	NetworkID    string                `json:"network_id"`
	ExposedPorts []types.TransportPort `json:"exposed_ports"`
	PortMapping  []types.PortBinding   `json:"port_mapping"`
}

// endpointJoin represents the expected body of the "join endpoint" or "leave endpoint" http request messages
type endpointJoin struct {
	ContainerID       string                 `json:"container_id"`
	HostName          string                 `json:"host_name"`
	DomainName        string                 `json:"domain_name"`
	HostsPath         string                 `json:"hosts_path"`
	ResolvConfPath    string                 `json:"resolv_conf_path"`
	DNS               []string               `json:"dns"`
	ExtraHosts        []endpointExtraHost    `json:"extra_hosts"`
	ParentUpdates     []endpointParentUpdate `json:"parent_updates"`
	UseDefaultSandbox bool                   `json:"use_default_sandbox"`
}

// EndpointExtraHost represents the extra host object
type endpointExtraHost struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// EndpointParentUpdate is the object carrying the information about the
// endpoint parent that needs to be updated
type endpointParentUpdate struct {
	EndpointID string `json:"endpoint_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
}
