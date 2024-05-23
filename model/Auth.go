package model

// 权限管理/控制
type Auth struct {
	Id          int    `gorm:"column:id" json:"id"`
	ModuleName  string `gorm:"column:module_name" json:"module_name"`
	ActionName  string `gorm:"column:action_name" json:"action_name"`
	Type        int8   `gorm:"column:type" json:"type"`
	Url         string `gorm:"column:url" json:"url"`
	ModuleID    int    `gorm:"column:module_id" json:"module_id"`
	Sort        int    `gorm:"column:sort" json:"sort"`
	Description string `gorm:"column:description" json:"description"`
	Status      int8   `gorm:"column:status" json:"status"`
	AddTime     int    `gorm:"column:add_time" json:"add_time"`
	Checked     int8   `gorm:"column:checked" json:"checked"`
}

func (Auth) TableName() string {
	return "auth"
}
