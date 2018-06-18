package apollctl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/Colstuwjx/apollet/pkg/apollet"
)

type Client struct {
	client *http.Client
	scheme string
	addr   string
}

const (
	timeout = time.Second * 10
)

func makeUnixDial(sock string) func(proto, addr string) (conn net.Conn, err error) {
	return func(proto, addr string) (conn net.Conn, err error) {
		return net.Dial("unix", sock)
	}
}

func NewClient(scheme, addr string) (*Client, error) {
	switch scheme {
	case "http":
		client := &http.Client{
			Timeout: timeout,
		}

		return &Client{
			client: client,
			scheme: scheme,
			addr:   addr,
		}, nil
	case "unix":
		tr := &http.Transport{Dial: makeUnixDial(addr)}
		client := &http.Client{
			Transport: tr,
			Timeout:   timeout,
		}

		return &Client{
			client: client,
			scheme: scheme,
			addr:   addr,
		}, nil
	default:
		return nil, fmt.Errorf("Unsupported scheme %s", scheme)
	}
}

func (this *Client) GetString(appId, cluster, namespace, key string) string {
	uri := fmt.Sprintf(
		"%s://%s/get_string?app_id=%s&cluster=%s&namespace=%s&key=%s",
		this.scheme,
		this.addr,
		appId,
		cluster,
		namespace,
		key,
	)

	resp, err := this.client.Get(uri)
	if err != nil {
		fmt.Println(err)
		return apollet.NotFoundDefaultValue
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		fmt.Println("Bad response")
		return apollet.NotFoundDefaultValue
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	respData := new(apollet.GetStringResponse)
	err = json.Unmarshal(data, respData)
	if err != nil {
		fmt.Println(err)
		return apollet.NotFoundDefaultValue
	}

	return respData.Data
}
