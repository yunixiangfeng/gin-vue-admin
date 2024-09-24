// 自动生成模板TestStruct
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// TestStruct 结构体
type TestStruct struct {
      global.GVA_MODEL
      TestF  string `json:"testF" form:"testF" gorm:"column:test_f;comment:;"`
}


// TableName TestStruct 表名
func (TestStruct) TableName() string {
  return "test_struct"
}

