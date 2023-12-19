package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShopConfigApi struct {
}

var shopConfigService = service.ServiceGroupApp.ShopServiceGroup.ShopConfigService

// CreateShopConfig 创建shopConfig表
// @Tags ShopConfig
// @Summary 创建shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopConfig true "创建shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shopConfig/createShopConfig [post]
func (shopConfigApi *ShopConfigApi) CreateShopConfig(c *gin.Context) {
	var shopConfig shop.ShopConfig
	err := c.ShouldBindJSON(&shopConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopConfigService.CreateShopConfig(&shopConfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShopConfig 删除shopConfig表
// @Tags ShopConfig
// @Summary 删除shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopConfig true "删除shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shopConfig/deleteShopConfig [delete]
func (shopConfigApi *ShopConfigApi) DeleteShopConfig(c *gin.Context) {
	var shopConfig shop.ShopConfig
	err := c.ShouldBindJSON(&shopConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopConfigService.DeleteShopConfig(shopConfig); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShopConfigByIds 批量删除shopConfig表
// @Tags ShopConfig
// @Summary 批量删除shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shopConfig/deleteShopConfigByIds [delete]
func (shopConfigApi *ShopConfigApi) DeleteShopConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopConfigService.DeleteShopConfigByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShopConfig 更新shopConfig表
// @Tags ShopConfig
// @Summary 更新shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shop.ShopConfig true "更新shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shopConfig/updateShopConfig [put]
func (shopConfigApi *ShopConfigApi) UpdateShopConfig(c *gin.Context) {
	var shopConfig shop.ShopConfig
	err := c.ShouldBindJSON(&shopConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shopConfigService.UpdateShopConfig(shopConfig); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShopConfig 用id查询shopConfig表
// @Tags ShopConfig
// @Summary 用id查询shopConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shop.ShopConfig true "用id查询shopConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shopConfig/findShopConfig [get]
func (shopConfigApi *ShopConfigApi) FindShopConfig(c *gin.Context) {
	var shopConfig shop.ShopConfig
	err := c.ShouldBindQuery(&shopConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reshopConfig, err := shopConfigService.GetShopConfig(shopConfig.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshopConfig": reshopConfig}, c)
	}
}

// GetShopConfigList 分页获取shopConfig表列表
// @Tags ShopConfig
// @Summary 分页获取shopConfig表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shopReq.ShopConfigSearch true "分页获取shopConfig表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shopConfig/getShopConfigList [get]
func (shopConfigApi *ShopConfigApi) GetShopConfigList(c *gin.Context) {
	var pageInfo shopReq.ShopConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shopConfigService.GetShopConfigInfoList(pageInfo); err != nil {
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
