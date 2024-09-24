// 自动生成模板TestStruct2
package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// TestStruct2 结构体
type TestStruct2 struct {
      global.GVA_MODEL
      TestF  string `json:"testF" form:"testF" gorm:"column:test_f;comment:;"`
}


// TableName TestStruct2 表名
func (TestStruct2) TableName() string {
  return "test_Struct2"
}

