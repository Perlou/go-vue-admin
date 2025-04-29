package request

import (
	"github.com/perlou/go-vue-admin/server/model/common/request"
	"github.com/perlou/go-vue-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
