package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestStruct2Router struct {
}

// InitTestStruct2Router 初始化 TestStruct2 路由信息
func (s *TestStruct2Router) InitTestStruct2Router(Router *gin.RouterGroup) {
	testStruct2Router := Router.Group("testStruct2").Use(middleware.OperationRecord())
	testStruct2RouterWithoutRecord := Router.Group("testStruct2")
	var testStruct2Api = v1.ApiGroupApp.PkgTest2ApiGroup.TestStruct2Api
	{
		testStruct2Router.POST("createTestStruct2", testStruct2Api.CreateTestStruct2)   // 新建TestStruct2
		testStruct2Router.DELETE("deleteTestStruct2", testStruct2Api.DeleteTestStruct2) // 删除TestStruct2
		testStruct2Router.DELETE("deleteTestStruct2ByIds", testStruct2Api.DeleteTestStruct2ByIds) // 批量删除TestStruct2
		testStruct2Router.PUT("updateTestStruct2", testStruct2Api.UpdateTestStruct2)    // 更新TestStruct2
	}
	{
		testStruct2RouterWithoutRecord.GET("findTestStruct2", testStruct2Api.FindTestStruct2)        // 根据ID获取TestStruct2
		testStruct2RouterWithoutRecord.GET("getTestStruct2List", testStruct2Api.GetTestStruct2List)  // 获取TestStruct2列表
	}
}
