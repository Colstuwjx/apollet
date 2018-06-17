package client

import (
	"fmt"

	"github.com/philchia/agollo"
)

type Client struct {
	// we assumes that ONLY one appId.cluster choice on each node
	clients map[string]*agollo.Client
	server  string
}

func NewClient(serverIp string) *Client {
	return &Client{
		clients: map[string]*agollo.Client{},
		server:  serverIp,
	}
}

func (this *Client) path(appId, cluster string) string {
	return fmt.Sprintf("%s.%s", appId, cluster)
}

func (this *Client) exists(path string) bool {
	if _, exists := this.clients[path]; exists {
		return true
	} else {
		return false
	}
}

func (this *Client) makeAgolloClient(appId, cluster string, namespaces []string) *agollo.Client {
	conf := &agollo.Conf{
		AppID:          appId,
		Cluster:        cluster,
		NameSpaceNames: namespaces,
		IP:             this.server,
	}
	agolloClient := agollo.NewClient(conf)
	return agolloClient
}

func (this *Client) appendToWatch(path string, agolloClient *agollo.Client) error {
	if agolloClient == nil {
		return fmt.Errorf("invalid agollo client instance for path %s", path)
	}

	// FIXME: should change to manual poll, and add rate limit.
	agolloClient.Start()
	this.clients[path] = agolloClient
	return nil
}

func (this *Client) GetString(appId, cluster, namespace, key, defaultValue string) string {
	path := this.path(appId, cluster)
	if this.exists(path) {
		// FIXME: add GetInt, GetYAML, ... supports.
		value := this.clients[path].GetStringValueWithNameSapce(namespace, key, defaultValue)
		return value
	}

	// FIXME: agollo cache path should be customized.
	namespaces := []string{namespace}
	agolloClient := this.makeAgolloClient(appId, cluster, namespaces)
	if agolloClient == nil {
		return defaultValue
	}

	err := this.appendToWatch(path, agolloClient)
	if err != nil {
		return defaultValue
	}

	value := agolloClient.GetStringValueWithNameSapce(namespace, key, defaultValue)
	return value
}
