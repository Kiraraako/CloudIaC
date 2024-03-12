// Copyright (c) 2015-2023 CloudJ Technology Co., Ltd.

package models

import (
	"database/sql/driver"
)

type ResAttrs map[string]interface{}

func (v ResAttrs) Value() (driver.Value, error) {
	return MarshalValue(v)
}

func (v *ResAttrs) Scan(value interface{}) error {
	return UnmarshalValue(value, v)
}

type Resource struct {
	BaseModel

	OrgId     Id `json:"orgId" gorm:"size:32;not null"`
	ProjectId Id `json:"projectId" gorm:"size:32;not null"`
	EnvId     Id `json:"envId" gorm:"index;size:32;not null"`
	TaskId    Id `json:"taskId" gorm:"index;size:32;not null"`

	ResId         Id       `json:"resId" gorm:"index;size:255;not null;default:''"`
	ResName       string   `json:"resName" gorm:"not null;default:''"`
	Provider      string   `json:"provider" gorm:"not null"`
	Module        string   `json:"module,omitempty" gorm:"not null;default:''"`
	Address       string   `json:"address" gorm:"not null"`
	Mode          string   `json:"mode" gorm:"not null"`
	Type          string   `json:"type" gorm:"not null"`
	Name          string   `json:"name" gorm:"not null"`
	Index         string   `json:"index" gorm:"not null;default:''"`
	Attrs         ResAttrs `json:"attrs,omitempty" gorm:"type:text"`
	SensitiveKeys StrSlice `json:"sensitiveKeys,omitempty" gorm:"type:text"`
	Dependencies  StrSlice `json:"dependencies,omitempty" gorm:"type:text"`

	AppliedAt Time `json:"appliedAt" gorm:"column:applied_at;default:null"`
}

func (Resource) TableName() string {
	return "iac_resource"
}

type ResFields []ResField

type ResField struct {
	ResId     Id   `json:"resId"`
	AppliedAt Time `json:"appliedAt"`
}

type ResourceMapping struct {
	BaseModel
	Provider string `json:"provider" gorm:"not null"` // 资源所属平台
	Type     string `json:"type" gorm:"not null"`     // 资源类型
	Code     string `json:"name" gorm:"not null"`     // 属性标识
	Express  string `json:"express" gorm:"not null"`  // 值表达式
}

func (ResourceMapping) TableName() string {
	return "iac_resource_mapping"
}

type ResourceMappingCondition struct {
	Provider string `json:"provider" gorm:"not null"` // 资源所属平台
	Type     string `json:"type" gorm:"not null"`     // 资源类型
	Code     string `json:"name" gorm:"not null"`     // 属性标识
}
