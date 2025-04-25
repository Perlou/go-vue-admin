package initialize

import "github.com/perlou/go-vue-admin/server/global"

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
