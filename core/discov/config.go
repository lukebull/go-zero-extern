package discov

import "errors"

// EtcdConf is the config item with the given key on etcd.
type EtcdConf struct {
	Hosts    []string
	Key      string
	Tls      bool   `json:",default=false,options=true|false"`
	Cafile   string `json:",optional"`
	Certfile string `json:",optional"`
	Keyfile  string `json:",optional"`
}

// Validate validates c.
func (c EtcdConf) Validate() error {
	if len(c.Hosts) == 0 {
		return errors.New("empty etcd hosts")
	} else if len(c.Key) == 0 {
		return errors.New("empty etcd key")
	} else {
		return nil
	}
}
