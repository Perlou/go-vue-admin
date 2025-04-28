package system

import "github.com/perlou/go-vue-admin/server/config"

type System struct {
	Config config.Server `json:"config"`
}
