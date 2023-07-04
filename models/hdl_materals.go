package models

type HdlMaterals struct {
	Id        string `json:"id" gorm:"primaryKey"`               // ID
	Name      string `json:"name,omitempty" gorm:"size:500"`     // 物料名称
	Dosage    int64  `json:"dosage,omitempty"`                   // 用量
	Unit      string `json:"unit,omitempty" gorm:"size:255"`     // 单位
	WaterLine int64  `json:"water_line,omitempty"`               // 加汤水位标准
	Station   string `json:"station,omitempty"  gorm:"size:255"` //工位 （鲜料工位、传锅工位、所有工位、不显示）
	Resource  string `json:"resource,omitempty" gorm:"size:255"` //物料来源 material-锅底 taste-口味
	CreateAt  int64  `json:"create_at,omitempty"`
	Remark    string `json:"remark,omitempty" gorm:"size:255"`
}

func (HdlMaterals) TableName() string {
	return "hdl_materals"
}
