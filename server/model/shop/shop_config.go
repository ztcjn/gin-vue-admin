// 自动生成模板ShopConfig
package shop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// shopConfig表 结构体  ShopConfig
type ShopConfig struct {
	global.GVA_MODEL         //`json:"UserId" form:"UserId" gorm:"column:UserId;type:int(10);comment:用户id 关联;size:10;"`
	Mode             *string `json:"mode" form:"mode" gorm:"column:mode;type:TINYINT(1);comment:棉花糖机器支付模式 1 = 汇联模式 2 = 自建模式;"` //支付模式
}

// TableName shopConfig表 ShopConfig自定义表名 shop_config
func (ShopConfig) TableName() string {
	return "shop_config"
}
