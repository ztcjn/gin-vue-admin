import service from '@/utils/request'

// @Tags ShopConfig
// @Summary 创建shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopConfig true "创建shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopConfig/createShopConfig [post]
export const createShopConfig = (data) => {
  return service({
    url: '/shopConfig/createShopConfig',
    method: 'post',
    data
  })
}

// @Tags ShopConfig
// @Summary 删除shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopConfig true "删除shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopConfig/deleteShopConfig [delete]
export const deleteShopConfig = (data) => {
  return service({
    url: '/shopConfig/deleteShopConfig',
    method: 'delete',
    data
  })
}

// @Tags ShopConfig
// @Summary 批量删除shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopConfig/deleteShopConfig [delete]
export const deleteShopConfigByIds = (data) => {
  return service({
    url: '/shopConfig/deleteShopConfigByIds',
    method: 'delete',
    data
  })
}

// @Tags ShopConfig
// @Summary 更新shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShopConfig true "更新shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopConfig/updateShopConfig [put]
export const updateShopConfig = (data) => {
  return service({
    url: '/shopConfig/updateShopConfig',
    method: 'put',
    data
  })
}

// @Tags ShopConfig
// @Summary 用id查询shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShopConfig true "用id查询shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopConfig/findShopConfig [get]
export const findShopConfig = (params) => {
  return service({
    url: '/shopConfig/findShopConfig',
    method: 'get',
    params
  })
}

// @Tags ShopConfig
// @Summary 分页获取shopConfig表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取shopConfig表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopConfig/getShopConfigList [get]
export const getShopConfigList = (params) => {
  return service({
    url: '/shopConfig/getShopConfigList',
    method: 'get',
    params
  })
}
