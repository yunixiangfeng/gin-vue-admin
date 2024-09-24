import service from '@/utils/request'

// @Tags TestStruct2
// @Summary 创建TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct2 true "创建TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStruct2/createTestStruct2 [post]
export const createTestStruct2 = (data) => {
  return service({
    url: '/testStruct2/createTestStruct2',
    method: 'post',
    data
  })
}

// @Tags TestStruct2
// @Summary 删除TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct2 true "删除TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStruct2/deleteTestStruct2 [delete]
export const deleteTestStruct2 = (data) => {
  return service({
    url: '/testStruct2/deleteTestStruct2',
    method: 'delete',
    data
  })
}

// @Tags TestStruct2
// @Summary 删除TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testStruct2/deleteTestStruct2 [delete]
export const deleteTestStruct2ByIds = (data) => {
  return service({
    url: '/testStruct2/deleteTestStruct2ByIds',
    method: 'delete',
    data
  })
}

// @Tags TestStruct2
// @Summary 更新TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct2 true "更新TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testStruct2/updateTestStruct2 [put]
export const updateTestStruct2 = (data) => {
  return service({
    url: '/testStruct2/updateTestStruct2',
    method: 'put',
    data
  })
}

// @Tags TestStruct2
// @Summary 用id查询TestStruct2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestStruct2 true "用id查询TestStruct2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testStruct2/findTestStruct2 [get]
export const findTestStruct2 = (params) => {
  return service({
    url: '/testStruct2/findTestStruct2',
    method: 'get',
    params
  })
}

// @Tags TestStruct2
// @Summary 分页获取TestStruct2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TestStruct2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testStruct2/getTestStruct2List [get]
export const getTestStruct2List = (params) => {
  return service({
    url: '/testStruct2/getTestStruct2List',
    method: 'get',
    params
  })
}
