package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShopConfigRouter struct {
}

// InitShopConfigRouter 初始化 shopConfig表 路由信息
func (s *ShopConfigRouter) InitShopConfigRouter(Router *gin.RouterGroup) {
	shopConfigRouter := Router.Group("shopConfig").Use(middleware.OperationRecord())
	shopConfigRouterWithoutRecord := Router.Group("shopConfig")
	var shopConfigApi = v1.ApiGroupApp.ShopApiGroup.ShopConfigApi
	{
		shopConfigRouter.POST("createShopConfig", shopConfigApi.CreateShopConfig)             // 新建shopConfig表
		shopConfigRouter.DELETE("deleteShopConfig", shopConfigApi.DeleteShopConfig)           // 删除shopConfig表
		shopConfigRouter.DELETE("deleteShopConfigByIds", shopConfigApi.DeleteShopConfigByIds) // 批量删除shopConfig表
		shopConfigRouter.PUT("updateShopConfig", shopConfigApi.UpdateShopConfig)              // 更新shopConfig表
	}
	{
		shopConfigRouterWithoutRecord.GET("findShopConfig", shopConfigApi.FindShopConfig)       // 根据ID获取shopConfig表
		shopConfigRouterWithoutRecord.GET("getShopConfigList", shopConfigApi.GetShopConfigList) // 获取shopConfig表列表
	}
}
