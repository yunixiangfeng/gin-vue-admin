package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTest2Req "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type TestStruct2Api struct {
}

var testStruct2Service = service.ServiceGroupApp.PkgTest2ServiceGroup.TestStruct2Service


// CreateTestStruct2 创建TestStruct2
// @Tags TestStruct2
// @Summary 创建TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest2.TestStruct2 true "创建TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStruct2/createTestStruct2 [post]
func (testStruct2Api *TestStruct2Api) CreateTestStruct2(c *gin.Context) {
	var testStruct2 pkgTest2.TestStruct2
	err := c.ShouldBindJSON(&testStruct2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testStruct2Service.CreateTestStruct2(&testStruct2); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestStruct2 删除TestStruct2
// @Tags TestStruct2
// @Summary 删除TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest2.TestStruct2 true "删除TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStruct2/deleteTestStruct2 [delete]
func (testStruct2Api *TestStruct2Api) DeleteTestStruct2(c *gin.Context) {
	var testStruct2 pkgTest2.TestStruct2
	err := c.ShouldBindJSON(&testStruct2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testStruct2Service.DeleteTestStruct2(testStruct2); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestStruct2ByIds 批量删除TestStruct2
// @Tags TestStruct2
// @Summary 批量删除TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /testStruct2/deleteTestStruct2ByIds [delete]
func (testStruct2Api *TestStruct2Api) DeleteTestStruct2ByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testStruct2Service.DeleteTestStruct2ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestStruct2 更新TestStruct2
// @Tags TestStruct2
// @Summary 更新TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest2.TestStruct2 true "更新TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testStruct2/updateTestStruct2 [put]
func (testStruct2Api *TestStruct2Api) UpdateTestStruct2(c *gin.Context) {
	var testStruct2 pkgTest2.TestStruct2
	err := c.ShouldBindJSON(&testStruct2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testStruct2Service.UpdateTestStruct2(testStruct2); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestStruct2 用id查询TestStruct2
// @Tags TestStruct2
// @Summary 用id查询TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest2.TestStruct2 true "用id查询TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testStruct2/findTestStruct2 [get]
func (testStruct2Api *TestStruct2Api) FindTestStruct2(c *gin.Context) {
	var testStruct2 pkgTest2.TestStruct2
	err := c.ShouldBindQuery(&testStruct2)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retestStruct2, err := testStruct2Service.GetTestStruct2(testStruct2.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retestStruct2": retestStruct2}, c)
	}
}

// GetTestStruct2List 分页获取TestStruct2列表
// @Tags TestStruct2
// @Summary 分页获取TestStruct2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest2Req.TestStruct2Search true "分页获取TestStruct2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStruct2/getTestStruct2List [get]
func (testStruct2Api *TestStruct2Api) GetTestStruct2List(c *gin.Context) {
	var pageInfo pkgTest2Req.TestStruct2Search
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testStruct2Service.GetTestStruct2InfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
