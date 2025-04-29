package response

import "github.com/perlou/go-vue-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
