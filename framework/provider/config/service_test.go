package config

import (
	"github.com/gohade/hade/framework"
	"path/filepath"
	"testing"

	"github.com/gohade/hade/framework/contract"
	tests "github.com/gohade/hade/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeConfig_GetInt(t *testing.T) {
	Convey("test hade env normal case", t, func() {
		c := framework.NewHadeContainer()
		basePath := tests.BasePath
		folder := filepath.Join(basePath, "config")
		serv, err := NewHadeConfig(folder, map[string]string{}, contract.EnvDevelopment, c)
		So(err, ShouldBeNil)
		conf := serv.(*HadeConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
