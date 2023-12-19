package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
)

type ShopConfigService struct {
}

// CreateShopConfig 创建shopConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopConfigService *ShopConfigService) CreateShopConfig(shopConfig *shop.ShopConfig) (err error) {
	err = global.GVA_DB.Create(shopConfig).Error
	return err
}

// DeleteShopConfig 删除shopConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopConfigService *ShopConfigService) DeleteShopConfig(shopConfig shop.ShopConfig) (err error) {
	err = global.GVA_DB.Delete(&shopConfig).Error
	return err
}

// DeleteShopConfigByIds 批量删除shopConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopConfigService *ShopConfigService) DeleteShopConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]shop.ShopConfig{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateShopConfig 更新shopConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopConfigService *ShopConfigService) UpdateShopConfig(shopConfig shop.ShopConfig) (err error) {
	err = global.GVA_DB.Save(&shopConfig).Error
	return err
}

// GetShopConfig 根据id获取shopConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopConfigService *ShopConfigService) GetShopConfig(id uint) (shopConfig shop.ShopConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&shopConfig).Error
	return
}

// GetShopConfigInfoList 分页获取shopConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (shopConfigService *ShopConfigService) GetShopConfigInfoList(info shopReq.ShopConfigSearch) (list []shop.ShopConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&shop.ShopConfig{})
	var shopConfigs []shop.ShopConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Mode != nil {
		db = db.Where("mode = ?", info.Mode)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&shopConfigs).Error
	return shopConfigs, total, err
}
