// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package micro_app_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *CreateInnerAppHeaders) SetCommonHeaders(v map[string]*string) *CreateInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInnerAppRequest struct {
	// 创建人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 关联组织corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用首页地址
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// 应用PC端地址
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// 应用管理后台地址
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// 服务器出口ip白名单
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
	// 权限类型
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s CreateInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppRequest) GoString() string {
	return s.String()
}

func (s *CreateInnerAppRequest) SetOpUnionId(v string) *CreateInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *CreateInnerAppRequest) SetEcologicalCorpId(v string) *CreateInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *CreateInnerAppRequest) SetName(v string) *CreateInnerAppRequest {
	s.Name = &v
	return s
}

func (s *CreateInnerAppRequest) SetDesc(v string) *CreateInnerAppRequest {
	s.Desc = &v
	return s
}

func (s *CreateInnerAppRequest) SetIcon(v string) *CreateInnerAppRequest {
	s.Icon = &v
	return s
}

func (s *CreateInnerAppRequest) SetHomepageLink(v string) *CreateInnerAppRequest {
	s.HomepageLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetPcHomepageLink(v string) *CreateInnerAppRequest {
	s.PcHomepageLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetOmpLink(v string) *CreateInnerAppRequest {
	s.OmpLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetIpWhiteList(v []*string) *CreateInnerAppRequest {
	s.IpWhiteList = v
	return s
}

func (s *CreateInnerAppRequest) SetScopeType(v string) *CreateInnerAppRequest {
	s.ScopeType = &v
	return s
}

type CreateInnerAppResponseBody struct {
	// 应用id
	AgentId   *int64  `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AppKey    *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
}

func (s CreateInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInnerAppResponseBody) SetAgentId(v int64) *CreateInnerAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *CreateInnerAppResponseBody) SetAppKey(v string) *CreateInnerAppResponseBody {
	s.AppKey = &v
	return s
}

func (s *CreateInnerAppResponseBody) SetAppSecret(v string) *CreateInnerAppResponseBody {
	s.AppSecret = &v
	return s
}

type CreateInnerAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppResponse) GoString() string {
	return s.String()
}

func (s *CreateInnerAppResponse) SetHeaders(v map[string]*string) *CreateInnerAppResponse {
	s.Headers = v
	return s
}

func (s *CreateInnerAppResponse) SetBody(v *CreateInnerAppResponseBody) *CreateInnerAppResponse {
	s.Body = v
	return s
}

type GetInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *GetInnerAppHeaders) SetCommonHeaders(v map[string]*string) *GetInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *GetInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInnerAppRequest struct {
	// 操作人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 关联组织corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
}

func (s GetInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppRequest) GoString() string {
	return s.String()
}

func (s *GetInnerAppRequest) SetOpUnionId(v string) *GetInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *GetInnerAppRequest) SetEcologicalCorpId(v string) *GetInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

type GetInnerAppResponseBody struct {
	// 应用id
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用移动端首页地址
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// 应用PC端首页地址
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// 应用管理后台地址
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// 应用的appkey
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// 应用的secret
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
	// 服务器出口ip
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
}

func (s GetInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetInnerAppResponseBody) SetAgentId(v int64) *GetInnerAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *GetInnerAppResponseBody) SetName(v string) *GetInnerAppResponseBody {
	s.Name = &v
	return s
}

func (s *GetInnerAppResponseBody) SetDesc(v string) *GetInnerAppResponseBody {
	s.Desc = &v
	return s
}

func (s *GetInnerAppResponseBody) SetIcon(v string) *GetInnerAppResponseBody {
	s.Icon = &v
	return s
}

func (s *GetInnerAppResponseBody) SetHomepageLink(v string) *GetInnerAppResponseBody {
	s.HomepageLink = &v
	return s
}

func (s *GetInnerAppResponseBody) SetPcHomepageLink(v string) *GetInnerAppResponseBody {
	s.PcHomepageLink = &v
	return s
}

func (s *GetInnerAppResponseBody) SetOmpLink(v string) *GetInnerAppResponseBody {
	s.OmpLink = &v
	return s
}

func (s *GetInnerAppResponseBody) SetAppKey(v string) *GetInnerAppResponseBody {
	s.AppKey = &v
	return s
}

func (s *GetInnerAppResponseBody) SetAppSecret(v string) *GetInnerAppResponseBody {
	s.AppSecret = &v
	return s
}

func (s *GetInnerAppResponseBody) SetIpWhiteList(v []*string) *GetInnerAppResponseBody {
	s.IpWhiteList = v
	return s
}

type GetInnerAppResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppResponse) GoString() string {
	return s.String()
}

func (s *GetInnerAppResponse) SetHeaders(v map[string]*string) *GetInnerAppResponse {
	s.Headers = v
	return s
}

func (s *GetInnerAppResponse) SetBody(v *GetInnerAppResponseBody) *GetInnerAppResponse {
	s.Body = v
	return s
}

type RegisterCustomAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterCustomAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *RegisterCustomAppRoleHeaders) SetCommonHeaders(v map[string]*string) *RegisterCustomAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterCustomAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterCustomAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterCustomAppRoleRequest struct {
	// 执行用户userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 是否拥有管理角色的权限，可不传，默认false
	CanManageRole *bool `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
}

func (s RegisterCustomAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomAppRoleRequest) GoString() string {
	return s.String()
}

func (s *RegisterCustomAppRoleRequest) SetOpUserId(v string) *RegisterCustomAppRoleRequest {
	s.OpUserId = &v
	return s
}

func (s *RegisterCustomAppRoleRequest) SetRoleName(v string) *RegisterCustomAppRoleRequest {
	s.RoleName = &v
	return s
}

func (s *RegisterCustomAppRoleRequest) SetCanManageRole(v bool) *RegisterCustomAppRoleRequest {
	s.CanManageRole = &v
	return s
}

type RegisterCustomAppRoleResponseBody struct {
	// 角色id
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 角色版本号
	ScopeVersion *int64 `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
}

func (s RegisterCustomAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCustomAppRoleResponseBody) SetRoleId(v int64) *RegisterCustomAppRoleResponseBody {
	s.RoleId = &v
	return s
}

func (s *RegisterCustomAppRoleResponseBody) SetScopeVersion(v int64) *RegisterCustomAppRoleResponseBody {
	s.ScopeVersion = &v
	return s
}

type RegisterCustomAppRoleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterCustomAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterCustomAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomAppRoleResponse) GoString() string {
	return s.String()
}

func (s *RegisterCustomAppRoleResponse) SetHeaders(v map[string]*string) *RegisterCustomAppRoleResponse {
	s.Headers = v
	return s
}

func (s *RegisterCustomAppRoleResponse) SetBody(v *RegisterCustomAppRoleResponseBody) *RegisterCustomAppRoleResponse {
	s.Body = v
	return s
}

type UpdateApaasAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateApaasAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateApaasAppHeaders) GoString() string {
	return s.String()
}

func (s *UpdateApaasAppHeaders) SetCommonHeaders(v map[string]*string) *UpdateApaasAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateApaasAppHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateApaasAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateApaasAppRequest struct {
	AppName   *string `json:"appName,omitempty" xml:"appName,omitempty"`
	AppIcon   *string `json:"appIcon,omitempty" xml:"appIcon,omitempty"`
	AppStatus *int32  `json:"appStatus,omitempty" xml:"appStatus,omitempty"`
	OpUserId  *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	BizAppId  *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
}

func (s UpdateApaasAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApaasAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateApaasAppRequest) SetAppName(v string) *UpdateApaasAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateApaasAppRequest) SetAppIcon(v string) *UpdateApaasAppRequest {
	s.AppIcon = &v
	return s
}

func (s *UpdateApaasAppRequest) SetAppStatus(v int32) *UpdateApaasAppRequest {
	s.AppStatus = &v
	return s
}

func (s *UpdateApaasAppRequest) SetOpUserId(v string) *UpdateApaasAppRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateApaasAppRequest) SetBizAppId(v string) *UpdateApaasAppRequest {
	s.BizAppId = &v
	return s
}

type UpdateApaasAppResponseBody struct {
	// 钉钉侧应用id
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// ISV侧应用id
	BizAppId *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
}

func (s UpdateApaasAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApaasAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApaasAppResponseBody) SetAgentId(v int64) *UpdateApaasAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *UpdateApaasAppResponseBody) SetBizAppId(v string) *UpdateApaasAppResponseBody {
	s.BizAppId = &v
	return s
}

type UpdateApaasAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateApaasAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateApaasAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApaasAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateApaasAppResponse) SetHeaders(v map[string]*string) *UpdateApaasAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateApaasAppResponse) SetBody(v *UpdateApaasAppResponseBody) *UpdateApaasAppResponse {
	s.Body = v
	return s
}

type AddAppRolesToMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddAppRolesToMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberHeaders) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberHeaders) SetCommonHeaders(v map[string]*string) *AddAppRolesToMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAppRolesToMemberHeaders) SetXAcsDingtalkAccessToken(v string) *AddAppRolesToMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddAppRolesToMemberRequest struct {
	// 执行用户userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 人员id
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 人员类型，“DEPT”表示部门，“USER”表示员工
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// 角色Id列表
	RoleList []*AddAppRolesToMemberRequestRoleList `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
}

func (s AddAppRolesToMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberRequest) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberRequest) SetOpUserId(v string) *AddAppRolesToMemberRequest {
	s.OpUserId = &v
	return s
}

func (s *AddAppRolesToMemberRequest) SetMemberId(v string) *AddAppRolesToMemberRequest {
	s.MemberId = &v
	return s
}

func (s *AddAppRolesToMemberRequest) SetMemberType(v string) *AddAppRolesToMemberRequest {
	s.MemberType = &v
	return s
}

func (s *AddAppRolesToMemberRequest) SetRoleList(v []*AddAppRolesToMemberRequestRoleList) *AddAppRolesToMemberRequest {
	s.RoleList = v
	return s
}

type AddAppRolesToMemberRequestRoleList struct {
	// 角色ID
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 角色范围版本号
	ScopeVersion *int64 `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
}

func (s AddAppRolesToMemberRequestRoleList) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberRequestRoleList) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberRequestRoleList) SetRoleId(v int64) *AddAppRolesToMemberRequestRoleList {
	s.RoleId = &v
	return s
}

func (s *AddAppRolesToMemberRequestRoleList) SetScopeVersion(v int64) *AddAppRolesToMemberRequestRoleList {
	s.ScopeVersion = &v
	return s
}

type AddAppRolesToMemberResponseBody struct {
	Result []*AddAppRolesToMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s AddAppRolesToMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberResponseBody) SetResult(v []*AddAppRolesToMemberResponseBodyResult) *AddAppRolesToMemberResponseBody {
	s.Result = v
	return s
}

type AddAppRolesToMemberResponseBodyResult struct {
	// 角色id
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 角色范围最新版本号
	LatestScopeVersion *int64 `json:"latestScopeVersion,omitempty" xml:"latestScopeVersion,omitempty"`
	// 角色添加结果，true: 成功，false: 失败
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
	SubErrorCode *string `json:"subErrorCode,omitempty" xml:"subErrorCode,omitempty"`
	SubErrorMsg  *string `json:"subErrorMsg,omitempty" xml:"subErrorMsg,omitempty"`
}

func (s AddAppRolesToMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberResponseBodyResult) SetRoleId(v int64) *AddAppRolesToMemberResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *AddAppRolesToMemberResponseBodyResult) SetLatestScopeVersion(v int64) *AddAppRolesToMemberResponseBodyResult {
	s.LatestScopeVersion = &v
	return s
}

func (s *AddAppRolesToMemberResponseBodyResult) SetSuccess(v bool) *AddAppRolesToMemberResponseBodyResult {
	s.Success = &v
	return s
}

func (s *AddAppRolesToMemberResponseBodyResult) SetSubErrorCode(v string) *AddAppRolesToMemberResponseBodyResult {
	s.SubErrorCode = &v
	return s
}

func (s *AddAppRolesToMemberResponseBodyResult) SetSubErrorMsg(v string) *AddAppRolesToMemberResponseBodyResult {
	s.SubErrorMsg = &v
	return s
}

type AddAppRolesToMemberResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAppRolesToMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAppRolesToMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberResponse) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberResponse) SetHeaders(v map[string]*string) *AddAppRolesToMemberResponse {
	s.Headers = v
	return s
}

func (s *AddAppRolesToMemberResponse) SetBody(v *AddAppRolesToMemberResponseBody) *AddAppRolesToMemberResponse {
	s.Body = v
	return s
}

type GetAppRoleScopeByRoleIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAppRoleScopeByRoleIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAppRoleScopeByRoleIdHeaders) GoString() string {
	return s.String()
}

func (s *GetAppRoleScopeByRoleIdHeaders) SetCommonHeaders(v map[string]*string) *GetAppRoleScopeByRoleIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAppRoleScopeByRoleIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetAppRoleScopeByRoleIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAppRoleScopeByRoleIdResponseBody struct {
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 角色id
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 角色范围类型，“ALL_VISIBLE”表示全员，“PART_VISIBLE”表示部分
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 部门id列表
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// 员工userId列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	// 角色范围版本号
	ScopeVersion *string `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	// 是否拥有角色管理权限，默认false
	CanManageRole *bool `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
}

func (s GetAppRoleScopeByRoleIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppRoleScopeByRoleIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetRoleName(v string) *GetAppRoleScopeByRoleIdResponseBody {
	s.RoleName = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetRoleId(v int64) *GetAppRoleScopeByRoleIdResponseBody {
	s.RoleId = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetScopeType(v string) *GetAppRoleScopeByRoleIdResponseBody {
	s.ScopeType = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetDeptIdList(v []*int64) *GetAppRoleScopeByRoleIdResponseBody {
	s.DeptIdList = v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetUserIdList(v []*string) *GetAppRoleScopeByRoleIdResponseBody {
	s.UserIdList = v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetScopeVersion(v string) *GetAppRoleScopeByRoleIdResponseBody {
	s.ScopeVersion = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetCanManageRole(v bool) *GetAppRoleScopeByRoleIdResponseBody {
	s.CanManageRole = &v
	return s
}

type GetAppRoleScopeByRoleIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppRoleScopeByRoleIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppRoleScopeByRoleIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppRoleScopeByRoleIdResponse) GoString() string {
	return s.String()
}

func (s *GetAppRoleScopeByRoleIdResponse) SetHeaders(v map[string]*string) *GetAppRoleScopeByRoleIdResponse {
	s.Headers = v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponse) SetBody(v *GetAppRoleScopeByRoleIdResponseBody) *GetAppRoleScopeByRoleIdResponse {
	s.Body = v
	return s
}

type ListRoleInfoByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListRoleInfoByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRoleInfoByUserHeaders) GoString() string {
	return s.String()
}

func (s *ListRoleInfoByUserHeaders) SetCommonHeaders(v map[string]*string) *ListRoleInfoByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRoleInfoByUserHeaders) SetXAcsDingtalkAccessToken(v string) *ListRoleInfoByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListRoleInfoByUserResponseBody struct {
	Result []*ListRoleInfoByUserResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListRoleInfoByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoleInfoByUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoleInfoByUserResponseBody) SetResult(v []*ListRoleInfoByUserResponseBodyResult) *ListRoleInfoByUserResponseBody {
	s.Result = v
	return s
}

type ListRoleInfoByUserResponseBodyResult struct {
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 角色id
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 是否拥有角色管理权限，默认false
	CanManageRole *bool `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
}

func (s ListRoleInfoByUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoleInfoByUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoleInfoByUserResponseBodyResult) SetRoleName(v string) *ListRoleInfoByUserResponseBodyResult {
	s.RoleName = &v
	return s
}

func (s *ListRoleInfoByUserResponseBodyResult) SetRoleId(v int64) *ListRoleInfoByUserResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *ListRoleInfoByUserResponseBodyResult) SetCanManageRole(v bool) *ListRoleInfoByUserResponseBodyResult {
	s.CanManageRole = &v
	return s
}

type ListRoleInfoByUserResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRoleInfoByUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoleInfoByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoleInfoByUserResponse) GoString() string {
	return s.String()
}

func (s *ListRoleInfoByUserResponse) SetHeaders(v map[string]*string) *ListRoleInfoByUserResponse {
	s.Headers = v
	return s
}

func (s *ListRoleInfoByUserResponse) SetBody(v *ListRoleInfoByUserResponseBody) *ListRoleInfoByUserResponse {
	s.Body = v
	return s
}

type ListInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *ListInnerAppHeaders) SetCommonHeaders(v map[string]*string) *ListInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *ListInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListInnerAppRequest struct {
	// 合作空间corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
}

func (s ListInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppRequest) GoString() string {
	return s.String()
}

func (s *ListInnerAppRequest) SetEcologicalCorpId(v string) *ListInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

type ListInnerAppResponseBody struct {
	// 应用列表
	AppList []*ListInnerAppResponseBodyAppList `json:"appList,omitempty" xml:"appList,omitempty" type:"Repeated"`
}

func (s ListInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponseBody) SetAppList(v []*ListInnerAppResponseBodyAppList) *ListInnerAppResponseBody {
	s.AppList = v
	return s
}

type ListInnerAppResponseBodyAppList struct {
	// 应用id
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用移动端首页地址
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// 应用PC端首页地址
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// 应用管理后台地址
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
}

func (s ListInnerAppResponseBodyAppList) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponseBodyAppList) SetAgentId(v int64) *ListInnerAppResponseBodyAppList {
	s.AgentId = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetName(v string) *ListInnerAppResponseBodyAppList {
	s.Name = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetDesc(v string) *ListInnerAppResponseBodyAppList {
	s.Desc = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetIcon(v string) *ListInnerAppResponseBodyAppList {
	s.Icon = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetHomepageLink(v string) *ListInnerAppResponseBodyAppList {
	s.HomepageLink = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetPcHomepageLink(v string) *ListInnerAppResponseBodyAppList {
	s.PcHomepageLink = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetOmpLink(v string) *ListInnerAppResponseBodyAppList {
	s.OmpLink = &v
	return s
}

type ListInnerAppResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponse) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponse) SetHeaders(v map[string]*string) *ListInnerAppResponse {
	s.Headers = v
	return s
}

func (s *ListInnerAppResponse) SetBody(v *ListInnerAppResponseBody) *ListInnerAppResponse {
	s.Body = v
	return s
}

type RemoveMemberForAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveMemberForAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberForAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *RemoveMemberForAppRoleHeaders) SetCommonHeaders(v map[string]*string) *RemoveMemberForAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveMemberForAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveMemberForAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveMemberForAppRoleRequest struct {
	// 执行用户userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 角色范围版本号
	ScopeVersion *int64 `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	// 部门id列表
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// 员工userId列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s RemoveMemberForAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberForAppRoleRequest) GoString() string {
	return s.String()
}

func (s *RemoveMemberForAppRoleRequest) SetOpUserId(v string) *RemoveMemberForAppRoleRequest {
	s.OpUserId = &v
	return s
}

func (s *RemoveMemberForAppRoleRequest) SetScopeVersion(v int64) *RemoveMemberForAppRoleRequest {
	s.ScopeVersion = &v
	return s
}

func (s *RemoveMemberForAppRoleRequest) SetDeptIdList(v []*int64) *RemoveMemberForAppRoleRequest {
	s.DeptIdList = v
	return s
}

func (s *RemoveMemberForAppRoleRequest) SetUserIdList(v []*string) *RemoveMemberForAppRoleRequest {
	s.UserIdList = v
	return s
}

type RemoveMemberForAppRoleResponseBody struct {
	// 角色最新版本号
	LatestScopeVersion *int64 `json:"latestScopeVersion,omitempty" xml:"latestScopeVersion,omitempty"`
}

func (s RemoveMemberForAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberForAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMemberForAppRoleResponseBody) SetLatestScopeVersion(v int64) *RemoveMemberForAppRoleResponseBody {
	s.LatestScopeVersion = &v
	return s
}

type RemoveMemberForAppRoleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveMemberForAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveMemberForAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberForAppRoleResponse) GoString() string {
	return s.String()
}

func (s *RemoveMemberForAppRoleResponse) SetHeaders(v map[string]*string) *RemoveMemberForAppRoleResponse {
	s.Headers = v
	return s
}

func (s *RemoveMemberForAppRoleResponse) SetBody(v *RemoveMemberForAppRoleResponseBody) *RemoveMemberForAppRoleResponse {
	s.Body = v
	return s
}

type UpdateInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppHeaders) SetCommonHeaders(v map[string]*string) *UpdateInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInnerAppRequest struct {
	// 创建人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 关联组织corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// 应用名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 应用描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 应用图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 应用首页地址
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// 应用PC端地址
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// 应用管理后台地址
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// 服务器出口ip白名单
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
}

func (s UpdateInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppRequest) SetOpUnionId(v string) *UpdateInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *UpdateInnerAppRequest) SetEcologicalCorpId(v string) *UpdateInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *UpdateInnerAppRequest) SetName(v string) *UpdateInnerAppRequest {
	s.Name = &v
	return s
}

func (s *UpdateInnerAppRequest) SetDesc(v string) *UpdateInnerAppRequest {
	s.Desc = &v
	return s
}

func (s *UpdateInnerAppRequest) SetIcon(v string) *UpdateInnerAppRequest {
	s.Icon = &v
	return s
}

func (s *UpdateInnerAppRequest) SetHomepageLink(v string) *UpdateInnerAppRequest {
	s.HomepageLink = &v
	return s
}

func (s *UpdateInnerAppRequest) SetPcHomepageLink(v string) *UpdateInnerAppRequest {
	s.PcHomepageLink = &v
	return s
}

func (s *UpdateInnerAppRequest) SetOmpLink(v string) *UpdateInnerAppRequest {
	s.OmpLink = &v
	return s
}

func (s *UpdateInnerAppRequest) SetIpWhiteList(v []*string) *UpdateInnerAppRequest {
	s.IpWhiteList = v
	return s
}

type UpdateInnerAppResponseBody struct {
	// 更新结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppResponseBody) SetResult(v bool) *UpdateInnerAppResponseBody {
	s.Result = &v
	return s
}

type UpdateInnerAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppResponse) SetHeaders(v map[string]*string) *UpdateInnerAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateInnerAppResponse) SetBody(v *UpdateInnerAppResponseBody) *UpdateInnerAppResponse {
	s.Body = v
	return s
}

type AddMemberToAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddMemberToAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *AddMemberToAppRoleHeaders) SetCommonHeaders(v map[string]*string) *AddMemberToAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMemberToAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *AddMemberToAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddMemberToAppRoleRequest struct {
	// 执行用户userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 角色范围版本号
	ScopeVersion *int64 `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	// 部门id列表
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// 员工userId列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s AddMemberToAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToAppRoleRequest) GoString() string {
	return s.String()
}

func (s *AddMemberToAppRoleRequest) SetOpUserId(v string) *AddMemberToAppRoleRequest {
	s.OpUserId = &v
	return s
}

func (s *AddMemberToAppRoleRequest) SetScopeVersion(v int64) *AddMemberToAppRoleRequest {
	s.ScopeVersion = &v
	return s
}

func (s *AddMemberToAppRoleRequest) SetDeptIdList(v []*int64) *AddMemberToAppRoleRequest {
	s.DeptIdList = v
	return s
}

func (s *AddMemberToAppRoleRequest) SetUserIdList(v []*string) *AddMemberToAppRoleRequest {
	s.UserIdList = v
	return s
}

type AddMemberToAppRoleResponseBody struct {
	// 角色范围最新版本号
	LatestScopeVersion *int64 `json:"latestScopeVersion,omitempty" xml:"latestScopeVersion,omitempty"`
}

func (s AddMemberToAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberToAppRoleResponseBody) SetLatestScopeVersion(v int64) *AddMemberToAppRoleResponseBody {
	s.LatestScopeVersion = &v
	return s
}

type AddMemberToAppRoleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddMemberToAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddMemberToAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToAppRoleResponse) GoString() string {
	return s.String()
}

func (s *AddMemberToAppRoleResponse) SetHeaders(v map[string]*string) *AddMemberToAppRoleResponse {
	s.Headers = v
	return s
}

func (s *AddMemberToAppRoleResponse) SetBody(v *AddMemberToAppRoleResponseBody) *AddMemberToAppRoleResponse {
	s.Body = v
	return s
}

type ListAppRoleScopesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAppRoleScopesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesHeaders) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesHeaders) SetCommonHeaders(v map[string]*string) *ListAppRoleScopesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAppRoleScopesHeaders) SetXAcsDingtalkAccessToken(v string) *ListAppRoleScopesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAppRoleScopesRequest struct {
	// 起始点，默认0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 数据量，默认20，最大50
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAppRoleScopesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesRequest) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesRequest) SetNextToken(v int64) *ListAppRoleScopesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppRoleScopesRequest) SetSize(v int64) *ListAppRoleScopesRequest {
	s.Size = &v
	return s
}

type ListAppRoleScopesResponseBody struct {
	// 是否还有数据，true: 还有；false: 已经全部拉取完成
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下一次请求的起始点
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 数据列表
	DataList []*ListAppRoleScopesResponseBodyDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
}

func (s ListAppRoleScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesResponseBody) SetHasMore(v bool) *ListAppRoleScopesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListAppRoleScopesResponseBody) SetNextToken(v int64) *ListAppRoleScopesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppRoleScopesResponseBody) SetDataList(v []*ListAppRoleScopesResponseBodyDataList) *ListAppRoleScopesResponseBody {
	s.DataList = v
	return s
}

type ListAppRoleScopesResponseBodyDataList struct {
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// 角色Id
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// 角色范围类型，“ALL_VISIBLE”表示全员，“PART_VISIBLE”表示部分
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 部门id列表
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// 员工userId列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	// 角色范围最新版本号
	ScopeVersion *int64 `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	// 是否拥有角色管理权限，默认false
	CanManageRole *bool `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
}

func (s ListAppRoleScopesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesResponseBodyDataList) SetRoleName(v string) *ListAppRoleScopesResponseBodyDataList {
	s.RoleName = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetRoleId(v int64) *ListAppRoleScopesResponseBodyDataList {
	s.RoleId = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetScopeType(v string) *ListAppRoleScopesResponseBodyDataList {
	s.ScopeType = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetDeptIdList(v []*int64) *ListAppRoleScopesResponseBodyDataList {
	s.DeptIdList = v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetUserIdList(v []*string) *ListAppRoleScopesResponseBodyDataList {
	s.UserIdList = v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetScopeVersion(v int64) *ListAppRoleScopesResponseBodyDataList {
	s.ScopeVersion = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetCanManageRole(v bool) *ListAppRoleScopesResponseBodyDataList {
	s.CanManageRole = &v
	return s
}

type ListAppRoleScopesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppRoleScopesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppRoleScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesResponse) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesResponse) SetHeaders(v map[string]*string) *ListAppRoleScopesResponse {
	s.Headers = v
	return s
}

func (s *ListAppRoleScopesResponse) SetBody(v *ListAppRoleScopesResponseBody) *ListAppRoleScopesResponse {
	s.Body = v
	return s
}

type AddAppToWorkBenchGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddAppToWorkBenchGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupHeaders) SetCommonHeaders(v map[string]*string) *AddAppToWorkBenchGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAppToWorkBenchGroupHeaders) SetXAcsDingtalkAccessToken(v string) *AddAppToWorkBenchGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddAppToWorkBenchGroupRequest struct {
	// 创建人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 关联组织corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// 工作台分组id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
}

func (s AddAppToWorkBenchGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupRequest) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupRequest) SetOpUnionId(v string) *AddAppToWorkBenchGroupRequest {
	s.OpUnionId = &v
	return s
}

func (s *AddAppToWorkBenchGroupRequest) SetEcologicalCorpId(v string) *AddAppToWorkBenchGroupRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *AddAppToWorkBenchGroupRequest) SetComponentId(v string) *AddAppToWorkBenchGroupRequest {
	s.ComponentId = &v
	return s
}

type AddAppToWorkBenchGroupResponseBody struct {
	// 更新结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddAppToWorkBenchGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupResponseBody) SetResult(v bool) *AddAppToWorkBenchGroupResponseBody {
	s.Result = &v
	return s
}

type AddAppToWorkBenchGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAppToWorkBenchGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAppToWorkBenchGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupResponse) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupResponse) SetHeaders(v map[string]*string) *AddAppToWorkBenchGroupResponse {
	s.Headers = v
	return s
}

func (s *AddAppToWorkBenchGroupResponse) SetBody(v *AddAppToWorkBenchGroupResponseBody) *AddAppToWorkBenchGroupResponse {
	s.Body = v
	return s
}

type RebuildRoleScopeForAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RebuildRoleScopeForAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s RebuildRoleScopeForAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *RebuildRoleScopeForAppRoleHeaders) SetCommonHeaders(v map[string]*string) *RebuildRoleScopeForAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RebuildRoleScopeForAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *RebuildRoleScopeForAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RebuildRoleScopeForAppRoleRequest struct {
	// 执行用户userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 角色范围最新版本号
	ScopeVersion *int64 `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	// 角色范围类型，“ALL_VISIBLE”表示全员，“PART_VISIBLE”表示部分
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// 部门id列表
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// 员工userId列表
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s RebuildRoleScopeForAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildRoleScopeForAppRoleRequest) GoString() string {
	return s.String()
}

func (s *RebuildRoleScopeForAppRoleRequest) SetOpUserId(v string) *RebuildRoleScopeForAppRoleRequest {
	s.OpUserId = &v
	return s
}

func (s *RebuildRoleScopeForAppRoleRequest) SetScopeVersion(v int64) *RebuildRoleScopeForAppRoleRequest {
	s.ScopeVersion = &v
	return s
}

func (s *RebuildRoleScopeForAppRoleRequest) SetScopeType(v string) *RebuildRoleScopeForAppRoleRequest {
	s.ScopeType = &v
	return s
}

func (s *RebuildRoleScopeForAppRoleRequest) SetDeptIdList(v []*int64) *RebuildRoleScopeForAppRoleRequest {
	s.DeptIdList = v
	return s
}

func (s *RebuildRoleScopeForAppRoleRequest) SetUserIdList(v []*string) *RebuildRoleScopeForAppRoleRequest {
	s.UserIdList = v
	return s
}

type RebuildRoleScopeForAppRoleResponseBody struct {
	// 角色范围最新版本号
	LatestScopeVersion *int64 `json:"latestScopeVersion,omitempty" xml:"latestScopeVersion,omitempty"`
}

func (s RebuildRoleScopeForAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebuildRoleScopeForAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildRoleScopeForAppRoleResponseBody) SetLatestScopeVersion(v int64) *RebuildRoleScopeForAppRoleResponseBody {
	s.LatestScopeVersion = &v
	return s
}

type RebuildRoleScopeForAppRoleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RebuildRoleScopeForAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebuildRoleScopeForAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildRoleScopeForAppRoleResponse) GoString() string {
	return s.String()
}

func (s *RebuildRoleScopeForAppRoleResponse) SetHeaders(v map[string]*string) *RebuildRoleScopeForAppRoleResponse {
	s.Headers = v
	return s
}

func (s *RebuildRoleScopeForAppRoleResponse) SetBody(v *RebuildRoleScopeForAppRoleResponseBody) *RebuildRoleScopeForAppRoleResponse {
	s.Body = v
	return s
}

type RemoveApaasAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveApaasAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveApaasAppHeaders) GoString() string {
	return s.String()
}

func (s *RemoveApaasAppHeaders) SetCommonHeaders(v map[string]*string) *RemoveApaasAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveApaasAppHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveApaasAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveApaasAppRequest struct {
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	BizAppId *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
}

func (s RemoveApaasAppRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveApaasAppRequest) GoString() string {
	return s.String()
}

func (s *RemoveApaasAppRequest) SetOpUserId(v string) *RemoveApaasAppRequest {
	s.OpUserId = &v
	return s
}

func (s *RemoveApaasAppRequest) SetBizAppId(v string) *RemoveApaasAppRequest {
	s.BizAppId = &v
	return s
}

type RemoveApaasAppResponseBody struct {
	// 结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveApaasAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveApaasAppResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApaasAppResponseBody) SetResult(v bool) *RemoveApaasAppResponseBody {
	s.Result = &v
	return s
}

type RemoveApaasAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveApaasAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveApaasAppResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveApaasAppResponse) GoString() string {
	return s.String()
}

func (s *RemoveApaasAppResponse) SetHeaders(v map[string]*string) *RemoveApaasAppResponse {
	s.Headers = v
	return s
}

func (s *RemoveApaasAppResponse) SetBody(v *RemoveApaasAppResponseBody) *RemoveApaasAppResponse {
	s.Body = v
	return s
}

type DeleteAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAppRoleHeaders) SetCommonHeaders(v map[string]*string) *DeleteAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteAppRoleRequest struct {
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s DeleteAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRoleRequest) SetOpUserId(v string) *DeleteAppRoleRequest {
	s.OpUserId = &v
	return s
}

type DeleteAppRoleResponseBody struct {
	// 删除结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppRoleResponseBody) SetResult(v bool) *DeleteAppRoleResponseBody {
	s.Result = &v
	return s
}

type DeleteAppRoleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppRoleResponse) SetHeaders(v map[string]*string) *DeleteAppRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppRoleResponse) SetBody(v *DeleteAppRoleResponseBody) *DeleteAppRoleResponse {
	s.Body = v
	return s
}

type CreateApaasAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateApaasAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateApaasAppHeaders) GoString() string {
	return s.String()
}

func (s *CreateApaasAppHeaders) SetCommonHeaders(v map[string]*string) *CreateApaasAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateApaasAppHeaders) SetXAcsDingtalkAccessToken(v string) *CreateApaasAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateApaasAppRequest struct {
	AppName            *string `json:"appName,omitempty" xml:"appName,omitempty"`
	AppDesc            *string `json:"appDesc,omitempty" xml:"appDesc,omitempty"`
	AppIcon            *string `json:"appIcon,omitempty" xml:"appIcon,omitempty"`
	HomepageLink       *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	PcHomepageLink     *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	OmpLink            *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	HomepageEditLink   *string `json:"homepageEditLink,omitempty" xml:"homepageEditLink,omitempty"`
	PcHomepageEditLink *string `json:"pcHomepageEditLink,omitempty" xml:"pcHomepageEditLink,omitempty"`
	OpUserId           *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	BizAppId           *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
	TemplateKey        *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
	IsShortCut         *int32  `json:"isShortCut,omitempty" xml:"isShortCut,omitempty"`
}

func (s CreateApaasAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApaasAppRequest) GoString() string {
	return s.String()
}

func (s *CreateApaasAppRequest) SetAppName(v string) *CreateApaasAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateApaasAppRequest) SetAppDesc(v string) *CreateApaasAppRequest {
	s.AppDesc = &v
	return s
}

func (s *CreateApaasAppRequest) SetAppIcon(v string) *CreateApaasAppRequest {
	s.AppIcon = &v
	return s
}

func (s *CreateApaasAppRequest) SetHomepageLink(v string) *CreateApaasAppRequest {
	s.HomepageLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetPcHomepageLink(v string) *CreateApaasAppRequest {
	s.PcHomepageLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetOmpLink(v string) *CreateApaasAppRequest {
	s.OmpLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetHomepageEditLink(v string) *CreateApaasAppRequest {
	s.HomepageEditLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetPcHomepageEditLink(v string) *CreateApaasAppRequest {
	s.PcHomepageEditLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetOpUserId(v string) *CreateApaasAppRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateApaasAppRequest) SetBizAppId(v string) *CreateApaasAppRequest {
	s.BizAppId = &v
	return s
}

func (s *CreateApaasAppRequest) SetTemplateKey(v string) *CreateApaasAppRequest {
	s.TemplateKey = &v
	return s
}

func (s *CreateApaasAppRequest) SetIsShortCut(v int32) *CreateApaasAppRequest {
	s.IsShortCut = &v
	return s
}

type CreateApaasAppResponseBody struct {
	// 钉钉侧应用id
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// ISV侧应用id
	BizAppId *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
}

func (s CreateApaasAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApaasAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApaasAppResponseBody) SetAgentId(v int64) *CreateApaasAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *CreateApaasAppResponseBody) SetBizAppId(v string) *CreateApaasAppResponseBody {
	s.BizAppId = &v
	return s
}

type CreateApaasAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApaasAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApaasAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApaasAppResponse) GoString() string {
	return s.String()
}

func (s *CreateApaasAppResponse) SetHeaders(v map[string]*string) *CreateApaasAppResponse {
	s.Headers = v
	return s
}

func (s *CreateApaasAppResponse) SetBody(v *CreateApaasAppResponseBody) *CreateApaasAppResponse {
	s.Body = v
	return s
}

type DeleteInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppHeaders) SetCommonHeaders(v map[string]*string) *DeleteInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteInnerAppRequest struct {
	// 操作人unionId
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// 合作空间corpId
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
}

func (s DeleteInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppRequest) SetOpUnionId(v string) *DeleteInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *DeleteInnerAppRequest) SetEcologicalCorpId(v string) *DeleteInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

type DeleteInnerAppResponseBody struct {
	// 删除结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppResponseBody) SetResult(v bool) *DeleteInnerAppResponseBody {
	s.Result = &v
	return s
}

type DeleteInnerAppResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppResponse) SetHeaders(v map[string]*string) *DeleteInnerAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteInnerAppResponse) SetBody(v *DeleteInnerAppResponseBody) *DeleteInnerAppResponse {
	s.Body = v
	return s
}

type UpdateAppRoleInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateAppRoleInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRoleInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAppRoleInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateAppRoleInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAppRoleInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateAppRoleInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateAppRoleInfoRequest struct {
	// 执行用户userId
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// 变更角色名称，可不传，不传则不变
	NewRoleName *string `json:"newRoleName,omitempty" xml:"newRoleName,omitempty"`
	// 变更角色管理权限，可不传，不传则不变
	CanManageRole *bool `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
}

func (s UpdateAppRoleInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRoleInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRoleInfoRequest) SetOpUserId(v string) *UpdateAppRoleInfoRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateAppRoleInfoRequest) SetNewRoleName(v string) *UpdateAppRoleInfoRequest {
	s.NewRoleName = &v
	return s
}

func (s *UpdateAppRoleInfoRequest) SetCanManageRole(v bool) *UpdateAppRoleInfoRequest {
	s.CanManageRole = &v
	return s
}

type UpdateAppRoleInfoResponseBody struct {
	// 更新结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateAppRoleInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRoleInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppRoleInfoResponseBody) SetResult(v bool) *UpdateAppRoleInfoResponseBody {
	s.Result = &v
	return s
}

type UpdateAppRoleInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppRoleInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppRoleInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRoleInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppRoleInfoResponse) SetHeaders(v map[string]*string) *UpdateAppRoleInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppRoleInfoResponse) SetBody(v *UpdateAppRoleInfoResponseBody) *UpdateAppRoleInfoResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) CreateInnerApp(request *CreateInnerAppRequest) (_result *CreateInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInnerAppHeaders{}
	_result = &CreateInnerAppResponse{}
	_body, _err := client.CreateInnerAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInnerAppWithOptions(request *CreateInnerAppRequest, headers *CreateInnerAppHeaders, runtime *util.RuntimeOptions) (_result *CreateInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		body["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageLink)) {
		body["homepageLink"] = request.HomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageLink)) {
		body["pcHomepageLink"] = request.PcHomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.OmpLink)) {
		body["ompLink"] = request.OmpLink
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		body["ipWhiteList"] = request.IpWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		body["scopeType"] = request.ScopeType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInnerApp(agentId *string, request *GetInnerAppRequest) (_result *GetInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInnerAppHeaders{}
	_result = &GetInnerAppResponse{}
	_body, _err := client.GetInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInnerAppWithOptions(agentId *string, request *GetInnerAppRequest, headers *GetInnerAppHeaders, runtime *util.RuntimeOptions) (_result *GetInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		query["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterCustomAppRole(agentId *string, request *RegisterCustomAppRoleRequest) (_result *RegisterCustomAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterCustomAppRoleHeaders{}
	_result = &RegisterCustomAppRoleResponse{}
	_body, _err := client.RegisterCustomAppRoleWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterCustomAppRoleWithOptions(agentId *string, request *RegisterCustomAppRoleRequest, headers *RegisterCustomAppRoleHeaders, runtime *util.RuntimeOptions) (_result *RegisterCustomAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["roleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CanManageRole)) {
		body["canManageRole"] = request.CanManageRole
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RegisterCustomAppRoleResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterCustomAppRole"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApaasApp(request *UpdateApaasAppRequest) (_result *UpdateApaasAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateApaasAppHeaders{}
	_result = &UpdateApaasAppResponse{}
	_body, _err := client.UpdateApaasAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApaasAppWithOptions(request *UpdateApaasAppRequest, headers *UpdateApaasAppHeaders, runtime *util.RuntimeOptions) (_result *UpdateApaasAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppIcon)) {
		body["appIcon"] = request.AppIcon
	}

	if !tea.BoolValue(util.IsUnset(request.AppStatus)) {
		body["appStatus"] = request.AppStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BizAppId)) {
		body["bizAppId"] = request.BizAppId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateApaasAppResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateApaasApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/microApp/apaasApps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAppRolesToMember(agentId *string, request *AddAppRolesToMemberRequest) (_result *AddAppRolesToMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAppRolesToMemberHeaders{}
	_result = &AddAppRolesToMemberResponse{}
	_body, _err := client.AddAppRolesToMemberWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAppRolesToMemberWithOptions(agentId *string, request *AddAppRolesToMemberRequest, headers *AddAppRolesToMemberHeaders, runtime *util.RuntimeOptions) (_result *AddAppRolesToMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["memberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberType)) {
		body["memberType"] = request.MemberType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleList)) {
		body["roleList"] = request.RoleList
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddAppRolesToMemberResponse{}
	_body, _err := client.DoROARequest(tea.String("AddAppRolesToMember"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/members/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppRoleScopeByRoleId(agentId *string, roleId *string) (_result *GetAppRoleScopeByRoleIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAppRoleScopeByRoleIdHeaders{}
	_result = &GetAppRoleScopeByRoleIdResponse{}
	_body, _err := client.GetAppRoleScopeByRoleIdWithOptions(agentId, roleId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppRoleScopeByRoleIdWithOptions(agentId *string, roleId *string, headers *GetAppRoleScopeByRoleIdHeaders, runtime *util.RuntimeOptions) (_result *GetAppRoleScopeByRoleIdResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetAppRoleScopeByRoleIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAppRoleScopeByRoleId"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/roles/"+tea.StringValue(roleId)+"/scopes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoleInfoByUser(agentId *string, userId *string) (_result *ListRoleInfoByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRoleInfoByUserHeaders{}
	_result = &ListRoleInfoByUserResponse{}
	_body, _err := client.ListRoleInfoByUserWithOptions(agentId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRoleInfoByUserWithOptions(agentId *string, userId *string, headers *ListRoleInfoByUserHeaders, runtime *util.RuntimeOptions) (_result *ListRoleInfoByUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &ListRoleInfoByUserResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRoleInfoByUser"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/users/"+tea.StringValue(userId)+"/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInnerApp(request *ListInnerAppRequest) (_result *ListInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListInnerAppHeaders{}
	_result = &ListInnerAppResponse{}
	_body, _err := client.ListInnerAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInnerAppWithOptions(request *ListInnerAppRequest, headers *ListInnerAppHeaders, runtime *util.RuntimeOptions) (_result *ListInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/microApp/apps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveMemberForAppRole(agentId *string, roleId *string, request *RemoveMemberForAppRoleRequest) (_result *RemoveMemberForAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveMemberForAppRoleHeaders{}
	_result = &RemoveMemberForAppRoleResponse{}
	_body, _err := client.RemoveMemberForAppRoleWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveMemberForAppRoleWithOptions(agentId *string, roleId *string, request *RemoveMemberForAppRoleRequest, headers *RemoveMemberForAppRoleHeaders, runtime *util.RuntimeOptions) (_result *RemoveMemberForAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeVersion)) {
		body["scopeVersion"] = request.ScopeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DeptIdList)) {
		body["deptIdList"] = request.DeptIdList
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RemoveMemberForAppRoleResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveMemberForAppRole"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/roles/"+tea.StringValue(roleId)+"/members/batchRemove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInnerApp(agentId *string, request *UpdateInnerAppRequest) (_result *UpdateInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInnerAppHeaders{}
	_result = &UpdateInnerAppResponse{}
	_body, _err := client.UpdateInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInnerAppWithOptions(agentId *string, request *UpdateInnerAppRequest, headers *UpdateInnerAppHeaders, runtime *util.RuntimeOptions) (_result *UpdateInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		body["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageLink)) {
		body["homepageLink"] = request.HomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageLink)) {
		body["pcHomepageLink"] = request.PcHomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.OmpLink)) {
		body["ompLink"] = request.OmpLink
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		body["ipWhiteList"] = request.IpWhiteList
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMemberToAppRole(agentId *string, roleId *string, request *AddMemberToAppRoleRequest) (_result *AddMemberToAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddMemberToAppRoleHeaders{}
	_result = &AddMemberToAppRoleResponse{}
	_body, _err := client.AddMemberToAppRoleWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddMemberToAppRoleWithOptions(agentId *string, roleId *string, request *AddMemberToAppRoleRequest, headers *AddMemberToAppRoleHeaders, runtime *util.RuntimeOptions) (_result *AddMemberToAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeVersion)) {
		body["scopeVersion"] = request.ScopeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DeptIdList)) {
		body["deptIdList"] = request.DeptIdList
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddMemberToAppRoleResponse{}
	_body, _err := client.DoROARequest(tea.String("AddMemberToAppRole"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/roles/"+tea.StringValue(roleId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppRoleScopes(agentId *string, request *ListAppRoleScopesRequest) (_result *ListAppRoleScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAppRoleScopesHeaders{}
	_result = &ListAppRoleScopesResponse{}
	_body, _err := client.ListAppRoleScopesWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppRoleScopesWithOptions(agentId *string, request *ListAppRoleScopesRequest, headers *ListAppRoleScopesHeaders, runtime *util.RuntimeOptions) (_result *ListAppRoleScopesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ListAppRoleScopesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListAppRoleScopes"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/roles"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAppToWorkBenchGroup(agentId *string, request *AddAppToWorkBenchGroupRequest) (_result *AddAppToWorkBenchGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAppToWorkBenchGroupHeaders{}
	_result = &AddAppToWorkBenchGroupResponse{}
	_body, _err := client.AddAppToWorkBenchGroupWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAppToWorkBenchGroupWithOptions(agentId *string, request *AddAppToWorkBenchGroupRequest, headers *AddAppToWorkBenchGroupHeaders, runtime *util.RuntimeOptions) (_result *AddAppToWorkBenchGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		body["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentId)) {
		body["componentId"] = request.ComponentId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddAppToWorkBenchGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("AddAppToWorkBenchGroup"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/addToWorkBenchGroup"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebuildRoleScopeForAppRole(agentId *string, roleId *string, request *RebuildRoleScopeForAppRoleRequest) (_result *RebuildRoleScopeForAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RebuildRoleScopeForAppRoleHeaders{}
	_result = &RebuildRoleScopeForAppRoleResponse{}
	_body, _err := client.RebuildRoleScopeForAppRoleWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebuildRoleScopeForAppRoleWithOptions(agentId *string, roleId *string, request *RebuildRoleScopeForAppRoleRequest, headers *RebuildRoleScopeForAppRoleHeaders, runtime *util.RuntimeOptions) (_result *RebuildRoleScopeForAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeVersion)) {
		body["scopeVersion"] = request.ScopeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		body["scopeType"] = request.ScopeType
	}

	if !tea.BoolValue(util.IsUnset(request.DeptIdList)) {
		body["deptIdList"] = request.DeptIdList
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RebuildRoleScopeForAppRoleResponse{}
	_body, _err := client.DoROARequest(tea.String("RebuildRoleScopeForAppRole"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/roles/"+tea.StringValue(roleId)+"/scopes/rebuild"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveApaasApp(request *RemoveApaasAppRequest) (_result *RemoveApaasAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveApaasAppHeaders{}
	_result = &RemoveApaasAppResponse{}
	_body, _err := client.RemoveApaasAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveApaasAppWithOptions(request *RemoveApaasAppRequest, headers *RemoveApaasAppHeaders, runtime *util.RuntimeOptions) (_result *RemoveApaasAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BizAppId)) {
		body["bizAppId"] = request.BizAppId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RemoveApaasAppResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveApaasApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apaasApps/remove"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppRole(agentId *string, roleId *string, request *DeleteAppRoleRequest) (_result *DeleteAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAppRoleHeaders{}
	_result = &DeleteAppRoleResponse{}
	_body, _err := client.DeleteAppRoleWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppRoleWithOptions(agentId *string, roleId *string, request *DeleteAppRoleRequest, headers *DeleteAppRoleHeaders, runtime *util.RuntimeOptions) (_result *DeleteAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteAppRoleResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteAppRole"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/roles/"+tea.StringValue(roleId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApaasApp(request *CreateApaasAppRequest) (_result *CreateApaasAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateApaasAppHeaders{}
	_result = &CreateApaasAppResponse{}
	_body, _err := client.CreateApaasAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApaasAppWithOptions(request *CreateApaasAppRequest, headers *CreateApaasAppHeaders, runtime *util.RuntimeOptions) (_result *CreateApaasAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppDesc)) {
		body["appDesc"] = request.AppDesc
	}

	if !tea.BoolValue(util.IsUnset(request.AppIcon)) {
		body["appIcon"] = request.AppIcon
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageLink)) {
		body["homepageLink"] = request.HomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageLink)) {
		body["pcHomepageLink"] = request.PcHomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.OmpLink)) {
		body["ompLink"] = request.OmpLink
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageEditLink)) {
		body["homepageEditLink"] = request.HomepageEditLink
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageEditLink)) {
		body["pcHomepageEditLink"] = request.PcHomepageEditLink
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BizAppId)) {
		body["bizAppId"] = request.BizAppId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateKey)) {
		body["templateKey"] = request.TemplateKey
	}

	if !tea.BoolValue(util.IsUnset(request.IsShortCut)) {
		body["isShortCut"] = request.IsShortCut
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateApaasAppResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateApaasApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/microApp/apaasApps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInnerApp(agentId *string, request *DeleteInnerAppRequest) (_result *DeleteInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteInnerAppHeaders{}
	_result = &DeleteInnerAppResponse{}
	_body, _err := client.DeleteInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInnerAppWithOptions(agentId *string, request *DeleteInnerAppRequest, headers *DeleteInnerAppHeaders, runtime *util.RuntimeOptions) (_result *DeleteInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		query["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteInnerAppResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInnerApp"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppRoleInfo(agentId *string, roleId *string, request *UpdateAppRoleInfoRequest) (_result *UpdateAppRoleInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAppRoleInfoHeaders{}
	_result = &UpdateAppRoleInfoResponse{}
	_body, _err := client.UpdateAppRoleInfoWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppRoleInfoWithOptions(agentId *string, roleId *string, request *UpdateAppRoleInfoRequest, headers *UpdateAppRoleInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateAppRoleInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.NewRoleName)) {
		body["newRoleName"] = request.NewRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.CanManageRole)) {
		body["canManageRole"] = request.CanManageRole
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateAppRoleInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateAppRoleInfo"), tea.String("microApp_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/microApp/apps/"+tea.StringValue(agentId)+"/roles/"+tea.StringValue(roleId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
