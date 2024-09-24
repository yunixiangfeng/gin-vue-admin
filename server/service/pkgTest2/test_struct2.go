package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTest2Req "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2/request"
)

type TestStruct2Service struct {
}

// CreateTestStruct2 创建TestStruct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct2Service *TestStruct2Service) CreateTestStruct2(testStruct2 *pkgTest2.TestStruct2) (err error) {
	err = global.GVA_DB.Create(testStruct2).Error
	return err
}

// DeleteTestStruct2 删除TestStruct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct2Service *TestStruct2Service)DeleteTestStruct2(testStruct2 pkgTest2.TestStruct2) (err error) {
	err = global.GVA_DB.Delete(&testStruct2).Error
	return err
}

// DeleteTestStruct2ByIds 批量删除TestStruct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct2Service *TestStruct2Service)DeleteTestStruct2ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest2.TestStruct2{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTestStruct2 更新TestStruct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct2Service *TestStruct2Service)UpdateTestStruct2(testStruct2 pkgTest2.TestStruct2) (err error) {
	err = global.GVA_DB.Save(&testStruct2).Error
	return err
}

// GetTestStruct2 根据id获取TestStruct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct2Service *TestStruct2Service)GetTestStruct2(id uint) (testStruct2 pkgTest2.TestStruct2, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&testStruct2).Error
	return
}

// GetTestStruct2InfoList 分页获取TestStruct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (testStruct2Service *TestStruct2Service)GetTestStruct2InfoList(info pkgTest2Req.TestStruct2Search) (list []pkgTest2.TestStruct2, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest2.TestStruct2{})
    var testStruct2s []pkgTest2.TestStruct2
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&testStruct2s).Error
	return  testStruct2s, total, err
}
