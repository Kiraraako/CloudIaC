package forms

import (
	"cloudiac/portal/models"
	"time"
)

type CreatePolicyForm struct {
	BaseForm

	Name          string `json:"name" binding:"required" example:"ECS分配公网IP"`                                                          // 策略名称
	FixSuggestion string `json:"fixSuggestion" binding:"" example:"1. 设置 internet_max_bandwidth_out = 0\n 2. 取消设置 allocate_public_ip"` // 修复建议
	Severity      string `json:"severity" binding:"" enums:"'high','medium','low'" example:"medium"`                                   // 严重性

	Rego string `json:"rego" binding:"required"` // rego脚本
}

type SearchPolicyForm struct {
	PageForm

	Q        string      `form:"q" json:"q" binding:""` // 策略组名称，支持模糊搜索
	Severity string      `json:"severity" form:"severity" enums:"'high','medium','low','none'" example:"medium"`
	GroupId  []models.Id `json:"groupId" form:"groupId" `
}

type UpdatePolicyForm struct {
	BaseForm

	Id            models.Id `uri:"id"`
	Name          string    `json:"name" binding:"required" example:"ECS分配公网IP"`                                                          // 策略名称
	FixSuggestion string    `json:"fixSuggestion" binding:"" example:"1. 设置 internet_max_bandwidth_out = 0\n 2. 取消设置 allocate_public_ip"` // 修复建议
	Severity      string    `json:"severity" binding:"" enums:"'high','medium','low','none'" example:"medium"`                            // 严重性

	Rego   string `json:"rego" binding:"required"` // rego脚本
	Status string `json:"status" form:"status" enums:"'enable', 'disable'" `
}

type DeletePolicyForm struct {
	BaseForm

	Id models.Id `uri:"id"`
}

type DetailPolicyForm struct {
	BaseForm

	Id models.Id `uri:"id"`
}

type CreatePolicyGroupForm struct {
	BaseForm

	Name        string `json:"name" binding:"required" example:"安全合规策略组"`
	Description string `json:"description" binding:"" example:"本组包含对于安全合规的检查策略"`

	//PolicyIds []string `json:"policyIds" binding:"" example:"[\"po-c3ek0co6n88ldvq1n6ag\"]"`
}

type SearchPolicyGroupForm struct {
	PageForm

	Q string `form:"q" json:"q" binding:""` // 策略组名称，支持模糊搜索
}

type UpdatePolicyGroupForm struct {
	BaseForm

	Id          models.Id `uri:"id"`
	Name        string    `json:"name" form:"name" `
	Description string    `json:"description" binding:"" example:"本组包含对于安全合规的检查策略"`
	Status      string    `json:"status" form:"status" enums:"'enable', 'disable'" `
}

type DeletePolicyGroupForm struct {
	BaseForm

	Id models.Id `uri:"id"`
}

type DetailPolicyGroupForm struct {
	BaseForm

	Id models.Id `uri:"id"`
}

type CreatePolicyRelForm struct {
	BaseForm

	PolicyGroupIds []models.Id `json:"policyGroupIds" binding:"required" example:"[\"pog-c3ek0co6n88ldvq1n6ag\"]"`
	EnvId          models.Id   `json:"envId" binding:"" example:"env-c3ek0co6n88ldvq1n6ag"`
	TplId          models.Id   `json:"tplId" binding:"" example:"tpl-c3ek0co6n88ldvq1n6ag"`
}

type ScanTemplateForm struct {
	BaseForm

	Id    models.Id `uri:"id" binding:"" example:"tpl-c3ek0co6n88ldvq1n6ag"`      // 云模板Id
	Parse bool      `json:"parse" binding:""  enums:"true,false" example:"false"` // 是否只执行解析
}

type ScanEnvironmentForm struct {
	BaseForm

	Id    models.Id `uri:"id" binding:"" example:"env-c3ek0co6n88ldvq1n6ag"`      // 环境Id
	Parse bool      `json:"parse" binding:""  enums:"true,false" example:"false"` // 是否只执行解析
}

type PolicyParseForm struct {
	BaseForm

	TemplateId models.Id `form:"tplId" json:"tplId" binding:"" example:"tpl-c3ek0co6n88ldvq1n6ag"` // 云模板Id
	EnvId      models.Id `form:"envId" json:"envId" binding:"" example:"env-c3ek0co6n88ldvq1n6ag"` // 云模板Id
}

type OpnPolicyAndPolicyGroupRelForm struct {
	BaseForm

	PolicyGroupId models.Id `uri:"id" json:"policyGroupId" form:"policyGroupId" `
	RmPolicyIds   []string  `json:"rmPolicyIds" binding:"" example:"[\"po-c3ek0co6n88ldvq1n6ag\"]"`
	AddPolicyIds  []string  `json:"addPolicyIds" binding:"" example:"[\"po-c3ek0co6n88ldvq1n6ag\"]"`
}

type CreatePolicyShieldForm struct {
	BaseForm

	CreatorId models.Id   `json:"creatorId" `
	TplId     models.Id   `json:"tplId" `
	EnvId     []models.Id `json:"envId" `
	PolicyId  models.Id   `json:"policyId" `

	Reason string `json:"reason" form:"reason" `
	Type   string `json:"type" form:"type" enums:"'strategy','source'"`
}

type SearchPolicySuppressForm struct {
	BaseForm

	PolicyId models.Id `uri:"policyId"`
}

type DeletePolicySuppressForm struct {
	BaseForm

	Id models.Id `uri:"id"`
}

type SearchPolicyTplForm struct {
	PageForm
	Q string `form:"q" json:"q" binding:""` // 模糊搜索
}

type UpdatePolicyTplForm struct {
	BaseForm
	Scope   string      `json:"scope"`
	TplId   models.Id   `json:"tplId" form:"tplId" `
	GroupId []models.Id `json:"groupId" form:"groupId" `
}

type DetailPolicyTplForm struct {
	BaseForm
	Id models.Id `json:"id" form:"id" `
}

type SearchPolicyEnvForm struct {
	PageForm
	Q string `form:"q" json:"q" binding:""` // 模糊搜索
}

type UpdatePolicyEnvForm struct {
	BaseForm
	Scope   string      `json:"scope"`
	EnvId   models.Id   `json:"envId" form:"envId" `
	GroupId []models.Id `json:"groupId" form:"groupId" `
}

type EnvOfPolicyForm struct {
	PageForm

	Id       models.Id `json:"id" form:"id" `
	Q        string    `form:"q" json:"q" binding:""` // 策略组名称，支持模糊搜索
	Severity string    `json:"severity" form:"severity" enums:"'high','medium','low','none'" example:"medium"`
	GroupId  models.Id `json:"groupId" form:"groupId" `
}

type PolicyErrorForm struct {
	BaseForm
	Id models.Id `json:"id" form:"id" `
}

type PolicyReferenceForm struct {
	BaseForm
	Id models.Id `json:"id" form:"id" `
}

type PolicyRepoForm struct {
	BaseForm
	Id models.Id `json:"id" form:"id" `
}

type PolicyScanResultForm struct {
	PageForm

	Id    models.Id `uri:"id" `
	Scope string    `json:"-"`
}

type PolicyScanReportForm struct {
	BaseForm

	Id    models.Id `uri:"id" `
	Scope string    `json:"-"`
	From  time.Time `json:"from" form:"from" example:"2006-01-02T15:04:05Z07:00"` //  开始日期
	To    time.Time `json:"to" form:"to" example:"2006-01-02T15:04:05Z07:00"`     // 结束日期
}

type PolicyTestForm struct {
	BaseForm

	Input string `form:"input" json:"input" binding:"" example:"{\n\"alicloud_instance\": [\n\n{\t\n\"id\": \"alicloud_instance.instance\"..."` // 脚本验证源数据
	Rego  string `form:"rego" json:"rego" binding:"" example:"package accurics\ninstanceWithNoVpc[retVal] {..."`                                // rego脚本内容
}

type PolicyLastTasksForm struct {
	PageForm

	Id    models.Id `uri:"id" `
	Scope string    `json:"-"`
}
