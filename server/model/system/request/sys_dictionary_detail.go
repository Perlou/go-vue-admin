package request

import (
	"github.com/perlou/go-vue-admin/server/model/common/request"
	"github.com/perlou/go-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
