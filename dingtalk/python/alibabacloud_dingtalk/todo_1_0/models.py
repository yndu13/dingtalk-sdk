# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.model import TeaModel
from typing import Dict, Any, List


class GetTodoTypeConfigHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class GetTodoTypeConfigResponseBodyContentFieldList(TeaModel):
    def __init__(
        self,
        field_key: str = None,
        field_type: str = None,
        name_i18n: Dict[str, Any] = None,
    ):
        # 字段唯一标识
        self.field_key = field_key
        # 字段类型（取值为：text-文本，url-链接）
        self.field_type = field_type
        # 字段的显示名称（支持国际化）
        self.name_i18n = name_i18n

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.field_key is not None:
            result['fieldKey'] = self.field_key
        if self.field_type is not None:
            result['fieldType'] = self.field_type
        if self.name_i18n is not None:
            result['nameI18n'] = self.name_i18n
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('fieldKey') is not None:
            self.field_key = m.get('fieldKey')
        if m.get('fieldType') is not None:
            self.field_type = m.get('fieldType')
        if m.get('nameI18n') is not None:
            self.name_i18n = m.get('nameI18n')
        return self


class GetTodoTypeConfigResponseBodyActionList(TeaModel):
    def __init__(
        self,
        action_key: str = None,
        button_style_type: int = None,
        action_type: int = None,
        url: str = None,
        name_i18n: Dict[str, Any] = None,
    ):
        # 操作按钮的唯一标识
        self.action_key = action_key
        # 按钮样式类型（101：蓝色线型主按钮样式，例如「同意」，102：黑色线型副按钮样式，例如「拒绝」）
        self.button_style_type = button_style_type
        # 按钮类型（1：有操作的，2：直接跳转）
        self.action_type = action_type
        # 按钮类型为直接跳转时，对应的跳转url
        self.url = url
        # 按钮操作的显示名称（支持国际化）
        self.name_i18n = name_i18n

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action_key is not None:
            result['actionKey'] = self.action_key
        if self.button_style_type is not None:
            result['buttonStyleType'] = self.button_style_type
        if self.action_type is not None:
            result['actionType'] = self.action_type
        if self.url is not None:
            result['url'] = self.url
        if self.name_i18n is not None:
            result['nameI18n'] = self.name_i18n
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('actionKey') is not None:
            self.action_key = m.get('actionKey')
        if m.get('buttonStyleType') is not None:
            self.button_style_type = m.get('buttonStyleType')
        if m.get('actionType') is not None:
            self.action_type = m.get('actionType')
        if m.get('url') is not None:
            self.url = m.get('url')
        if m.get('nameI18n') is not None:
            self.name_i18n = m.get('nameI18n')
        return self


class GetTodoTypeConfigResponseBody(TeaModel):
    def __init__(
        self,
        id: str = None,
        created_time: int = None,
        modified_time: int = None,
        creator_id: str = None,
        modifier_id: str = None,
        biz_tag: str = None,
        request_id: str = None,
        card_type: int = None,
        icon: str = None,
        description: str = None,
        pc_detail_url_open_mode: str = None,
        content_field_list: List[GetTodoTypeConfigResponseBodyContentFieldList] = None,
        action_list: List[GetTodoTypeConfigResponseBodyActionList] = None,
    ):
        # id
        self.id = id
        # 创建时间
        self.created_time = created_time
        # 更新时间
        self.modified_time = modified_time
        # 创建者（用户的unionId）
        self.creator_id = creator_id
        # 更新者（用户的unionId）
        self.modifier_id = modifier_id
        # 接入应用标识
        self.biz_tag = biz_tag
        # requestId
        self.request_id = request_id
        # 卡片类型（取值为：1-标准卡片，2-自定义卡片）
        self.card_type = card_type
        # 卡片类型icon，用于在待办列表展示
        self.icon = icon
        # 待办卡片类型描述
        self.description = description
        # 详情页链接在PC端的打开方式，取值为：「PC_SLIDE」-PC端侧边栏打开，「PC_BROWSER」-浏览器打开
        self.pc_detail_url_open_mode = pc_detail_url_open_mode
        # 待办卡片内容区表单自定义字段配置
        self.content_field_list = content_field_list
        # 待办卡片操作区按钮配置
        self.action_list = action_list

    def validate(self):
        if self.content_field_list:
            for k in self.content_field_list:
                if k:
                    k.validate()
        if self.action_list:
            for k in self.action_list:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.id is not None:
            result['id'] = self.id
        if self.created_time is not None:
            result['createdTime'] = self.created_time
        if self.modified_time is not None:
            result['modifiedTime'] = self.modified_time
        if self.creator_id is not None:
            result['creatorId'] = self.creator_id
        if self.modifier_id is not None:
            result['modifierId'] = self.modifier_id
        if self.biz_tag is not None:
            result['bizTag'] = self.biz_tag
        if self.request_id is not None:
            result['requestId'] = self.request_id
        if self.card_type is not None:
            result['cardType'] = self.card_type
        if self.icon is not None:
            result['icon'] = self.icon
        if self.description is not None:
            result['description'] = self.description
        if self.pc_detail_url_open_mode is not None:
            result['pcDetailUrlOpenMode'] = self.pc_detail_url_open_mode
        result['contentFieldList'] = []
        if self.content_field_list is not None:
            for k in self.content_field_list:
                result['contentFieldList'].append(k.to_map() if k else None)
        result['actionList'] = []
        if self.action_list is not None:
            for k in self.action_list:
                result['actionList'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('id') is not None:
            self.id = m.get('id')
        if m.get('createdTime') is not None:
            self.created_time = m.get('createdTime')
        if m.get('modifiedTime') is not None:
            self.modified_time = m.get('modifiedTime')
        if m.get('creatorId') is not None:
            self.creator_id = m.get('creatorId')
        if m.get('modifierId') is not None:
            self.modifier_id = m.get('modifierId')
        if m.get('bizTag') is not None:
            self.biz_tag = m.get('bizTag')
        if m.get('requestId') is not None:
            self.request_id = m.get('requestId')
        if m.get('cardType') is not None:
            self.card_type = m.get('cardType')
        if m.get('icon') is not None:
            self.icon = m.get('icon')
        if m.get('description') is not None:
            self.description = m.get('description')
        if m.get('pcDetailUrlOpenMode') is not None:
            self.pc_detail_url_open_mode = m.get('pcDetailUrlOpenMode')
        self.content_field_list = []
        if m.get('contentFieldList') is not None:
            for k in m.get('contentFieldList'):
                temp_model = GetTodoTypeConfigResponseBodyContentFieldList()
                self.content_field_list.append(temp_model.from_map(k))
        self.action_list = []
        if m.get('actionList') is not None:
            for k in m.get('actionList'):
                temp_model = GetTodoTypeConfigResponseBodyActionList()
                self.action_list.append(temp_model.from_map(k))
        return self


class GetTodoTypeConfigResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: GetTodoTypeConfigResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = GetTodoTypeConfigResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class QueryTodoTasksHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class QueryTodoTasksRequest(TeaModel):
    def __init__(
        self,
        next_token: str = None,
        order_by: str = None,
        order_direction: str = None,
        is_done: bool = None,
        role_types: List[List[str]] = None,
        from_due_time: int = None,
        to_due_time: int = None,
        category: str = None,
        is_recycled: bool = None,
    ):
        # 分页游标。如果一个查询条件一次无法全部返回结果，会返回分页token，下次查询带上该token后会返回后续数据，直到分页token为null表示数据已经全部查询完毕。
        self.next_token = next_token
        # 排序字段。枚举值 默认为截止时间 dueTime。created | modified | finished | startTime | dueTime 创建时间 | 更新时间 | 完成时间 | 开始时间 | 截止时间
        self.order_by = order_by
        # 排序方向。枚举值asc | desc 默认 asc
        self.order_direction = order_direction
        # 待办完成状态。
        self.is_done = is_done
        # 查询目标用户角色类型，执行人 | 创建人 | 参与人，可以同时传多个值。如：[["executor"], ["creator"],["participant"]] 或 [["executor", "creator"]]
        self.role_types = role_types
        # 查询从计划完成时间开始
        self.from_due_time = from_due_time
        # 查询到计划完成时间结束
        self.to_due_time = to_due_time
        # 所属分类
        self.category = category
        # 待办回收状态
        self.is_recycled = is_recycled

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.next_token is not None:
            result['nextToken'] = self.next_token
        if self.order_by is not None:
            result['orderBy'] = self.order_by
        if self.order_direction is not None:
            result['orderDirection'] = self.order_direction
        if self.is_done is not None:
            result['isDone'] = self.is_done
        if self.role_types is not None:
            result['roleTypes'] = self.role_types
        if self.from_due_time is not None:
            result['fromDueTime'] = self.from_due_time
        if self.to_due_time is not None:
            result['toDueTime'] = self.to_due_time
        if self.category is not None:
            result['category'] = self.category
        if self.is_recycled is not None:
            result['isRecycled'] = self.is_recycled
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('nextToken') is not None:
            self.next_token = m.get('nextToken')
        if m.get('orderBy') is not None:
            self.order_by = m.get('orderBy')
        if m.get('orderDirection') is not None:
            self.order_direction = m.get('orderDirection')
        if m.get('isDone') is not None:
            self.is_done = m.get('isDone')
        if m.get('roleTypes') is not None:
            self.role_types = m.get('roleTypes')
        if m.get('fromDueTime') is not None:
            self.from_due_time = m.get('fromDueTime')
        if m.get('toDueTime') is not None:
            self.to_due_time = m.get('toDueTime')
        if m.get('category') is not None:
            self.category = m.get('category')
        if m.get('isRecycled') is not None:
            self.is_recycled = m.get('isRecycled')
        return self


class QueryTodoTasksResponseBodyTodoCardsDetailUrl(TeaModel):
    def __init__(
        self,
        app_url: str = None,
        pc_url: str = None,
    ):
        # 移动端url地址
        self.app_url = app_url
        # pc端url地址
        self.pc_url = pc_url

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.app_url is not None:
            result['appUrl'] = self.app_url
        if self.pc_url is not None:
            result['pcUrl'] = self.pc_url
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('appUrl') is not None:
            self.app_url = m.get('appUrl')
        if m.get('pcUrl') is not None:
            self.pc_url = m.get('pcUrl')
        return self


class QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList(TeaModel):
    def __init__(
        self,
        name: str = None,
        value: str = None,
    ):
        # 自定义表单内容名字
        self.name = name
        # 自定义表单内容值
        self.value = value

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.name is not None:
            result['name'] = self.name
        if self.value is not None:
            result['value'] = self.value
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('name') is not None:
            self.name = m.get('name')
        if m.get('value') is not None:
            self.value = m.get('value')
        return self


class QueryTodoTasksResponseBodyTodoCardsTodoCardView(TeaModel):
    def __init__(
        self,
        card_type: str = None,
        circle_eltype: str = None,
        content_type: str = None,
        action_type: str = None,
        icon: str = None,
        todo_card_title: str = None,
        todo_card_content_list: List[QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList] = None,
    ):
        # 卡片类型
        self.card_type = card_type
        # 卡片左上角 区域类型是 icon, 或者checkbox 类型的
        self.circle_eltype = circle_eltype
        # icon, name ,内容区域类型是 icon+value, 或者name+value 类型的
        self.content_type = content_type
        # link, button, 操作区类型，是链接类型，或者按钮类型
        self.action_type = action_type
        # 卡片icon
        self.icon = icon
        # 卡片标题
        self.todo_card_title = todo_card_title
        self.todo_card_content_list = todo_card_content_list

    def validate(self):
        if self.todo_card_content_list:
            for k in self.todo_card_content_list:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.card_type is not None:
            result['cardType'] = self.card_type
        if self.circle_eltype is not None:
            result['circleELType'] = self.circle_eltype
        if self.content_type is not None:
            result['contentType'] = self.content_type
        if self.action_type is not None:
            result['actionType'] = self.action_type
        if self.icon is not None:
            result['icon'] = self.icon
        if self.todo_card_title is not None:
            result['todoCardTitle'] = self.todo_card_title
        result['todoCardContentList'] = []
        if self.todo_card_content_list is not None:
            for k in self.todo_card_content_list:
                result['todoCardContentList'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('cardType') is not None:
            self.card_type = m.get('cardType')
        if m.get('circleELType') is not None:
            self.circle_eltype = m.get('circleELType')
        if m.get('contentType') is not None:
            self.content_type = m.get('contentType')
        if m.get('actionType') is not None:
            self.action_type = m.get('actionType')
        if m.get('icon') is not None:
            self.icon = m.get('icon')
        if m.get('todoCardTitle') is not None:
            self.todo_card_title = m.get('todoCardTitle')
        self.todo_card_content_list = []
        if m.get('todoCardContentList') is not None:
            for k in m.get('todoCardContentList'):
                temp_model = QueryTodoTasksResponseBodyTodoCardsTodoCardViewTodoCardContentList()
                self.todo_card_content_list.append(temp_model.from_map(k))
        return self


class QueryTodoTasksResponseBodyTodoCardsOriginalSource(TeaModel):
    def __init__(
        self,
        source_title: str = None,
    ):
        # 业务来源展示名称
        self.source_title = source_title

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.source_title is not None:
            result['sourceTitle'] = self.source_title
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('sourceTitle') is not None:
            self.source_title = m.get('sourceTitle')
        return self


class QueryTodoTasksResponseBodyTodoCardsOrgInfo(TeaModel):
    def __init__(
        self,
        corp_id: str = None,
        name: str = None,
    ):
        # 组织corpId
        self.corp_id = corp_id
        # 组织名称
        self.name = name

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.corp_id is not None:
            result['corpId'] = self.corp_id
        if self.name is not None:
            result['name'] = self.name
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('corpId') is not None:
            self.corp_id = m.get('corpId')
        if m.get('name') is not None:
            self.name = m.get('name')
        return self


class QueryTodoTasksResponseBodyTodoCards(TeaModel):
    def __init__(
        self,
        task_id: str = None,
        subject: str = None,
        due_time: int = None,
        detail_url: QueryTodoTasksResponseBodyTodoCardsDetailUrl = None,
        todo_card_view: QueryTodoTasksResponseBodyTodoCardsTodoCardView = None,
        priority: int = None,
        created_time: int = None,
        modified_time: int = None,
        todo_status: str = None,
        creator_id: str = None,
        source_id: str = None,
        category: str = None,
        biz_tag: str = None,
        original_source: QueryTodoTasksResponseBodyTodoCardsOriginalSource = None,
        is_done: bool = None,
        org_info: QueryTodoTasksResponseBodyTodoCardsOrgInfo = None,
    ):
        # 待办id
        self.task_id = task_id
        # 待办标题
        self.subject = subject
        # 待办截止时间
        self.due_time = due_time
        # 详情页链接
        self.detail_url = detail_url
        # 待办卡片视图模型
        self.todo_card_view = todo_card_view
        # 优先级
        self.priority = priority
        # 创建时间
        self.created_time = created_time
        # 更新时间
        self.modified_time = modified_time
        # 待办状态
        self.todo_status = todo_status
        # 创建者id
        self.creator_id = creator_id
        # 来源id
        self.source_id = source_id
        # 所属分类
        self.category = category
        # 所属应用
        self.biz_tag = biz_tag
        # 业务来源信息
        self.original_source = original_source
        # 待办完成状态
        self.is_done = is_done
        # 所属组织信息
        self.org_info = org_info

    def validate(self):
        if self.detail_url:
            self.detail_url.validate()
        if self.todo_card_view:
            self.todo_card_view.validate()
        if self.original_source:
            self.original_source.validate()
        if self.org_info:
            self.org_info.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.task_id is not None:
            result['taskId'] = self.task_id
        if self.subject is not None:
            result['subject'] = self.subject
        if self.due_time is not None:
            result['dueTime'] = self.due_time
        if self.detail_url is not None:
            result['detailUrl'] = self.detail_url.to_map()
        if self.todo_card_view is not None:
            result['todoCardView'] = self.todo_card_view.to_map()
        if self.priority is not None:
            result['priority'] = self.priority
        if self.created_time is not None:
            result['createdTime'] = self.created_time
        if self.modified_time is not None:
            result['modifiedTime'] = self.modified_time
        if self.todo_status is not None:
            result['todoStatus'] = self.todo_status
        if self.creator_id is not None:
            result['creatorId'] = self.creator_id
        if self.source_id is not None:
            result['sourceId'] = self.source_id
        if self.category is not None:
            result['category'] = self.category
        if self.biz_tag is not None:
            result['bizTag'] = self.biz_tag
        if self.original_source is not None:
            result['originalSource'] = self.original_source.to_map()
        if self.is_done is not None:
            result['isDone'] = self.is_done
        if self.org_info is not None:
            result['orgInfo'] = self.org_info.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('taskId') is not None:
            self.task_id = m.get('taskId')
        if m.get('subject') is not None:
            self.subject = m.get('subject')
        if m.get('dueTime') is not None:
            self.due_time = m.get('dueTime')
        if m.get('detailUrl') is not None:
            temp_model = QueryTodoTasksResponseBodyTodoCardsDetailUrl()
            self.detail_url = temp_model.from_map(m['detailUrl'])
        if m.get('todoCardView') is not None:
            temp_model = QueryTodoTasksResponseBodyTodoCardsTodoCardView()
            self.todo_card_view = temp_model.from_map(m['todoCardView'])
        if m.get('priority') is not None:
            self.priority = m.get('priority')
        if m.get('createdTime') is not None:
            self.created_time = m.get('createdTime')
        if m.get('modifiedTime') is not None:
            self.modified_time = m.get('modifiedTime')
        if m.get('todoStatus') is not None:
            self.todo_status = m.get('todoStatus')
        if m.get('creatorId') is not None:
            self.creator_id = m.get('creatorId')
        if m.get('sourceId') is not None:
            self.source_id = m.get('sourceId')
        if m.get('category') is not None:
            self.category = m.get('category')
        if m.get('bizTag') is not None:
            self.biz_tag = m.get('bizTag')
        if m.get('originalSource') is not None:
            temp_model = QueryTodoTasksResponseBodyTodoCardsOriginalSource()
            self.original_source = temp_model.from_map(m['originalSource'])
        if m.get('isDone') is not None:
            self.is_done = m.get('isDone')
        if m.get('orgInfo') is not None:
            temp_model = QueryTodoTasksResponseBodyTodoCardsOrgInfo()
            self.org_info = temp_model.from_map(m['orgInfo'])
        return self


class QueryTodoTasksResponseBody(TeaModel):
    def __init__(
        self,
        next_token: str = None,
        todo_cards: List[QueryTodoTasksResponseBodyTodoCards] = None,
        total_count: int = None,
    ):
        # 翻页token
        self.next_token = next_token
        # 待办卡片列表
        self.todo_cards = todo_cards
        # 数据总量
        self.total_count = total_count

    def validate(self):
        if self.todo_cards:
            for k in self.todo_cards:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.next_token is not None:
            result['nextToken'] = self.next_token
        result['todoCards'] = []
        if self.todo_cards is not None:
            for k in self.todo_cards:
                result['todoCards'].append(k.to_map() if k else None)
        if self.total_count is not None:
            result['totalCount'] = self.total_count
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('nextToken') is not None:
            self.next_token = m.get('nextToken')
        self.todo_cards = []
        if m.get('todoCards') is not None:
            for k in m.get('todoCards'):
                temp_model = QueryTodoTasksResponseBodyTodoCards()
                self.todo_cards.append(temp_model.from_map(k))
        if m.get('totalCount') is not None:
            self.total_count = m.get('totalCount')
        return self


class QueryTodoTasksResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: QueryTodoTasksResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = QueryTodoTasksResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class UpdateTodoTypeConfigHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class UpdateTodoTypeConfigRequestContentFieldList(TeaModel):
    def __init__(
        self,
        field_key: str = None,
        field_type: str = None,
        name_i18n: Dict[str, Any] = None,
    ):
        # 字段唯一标识
        self.field_key = field_key
        # 字段类型（取值为：text-文本，url-链接）
        self.field_type = field_type
        # 字段显示名称(需支持国际化)
        self.name_i18n = name_i18n

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.field_key is not None:
            result['fieldKey'] = self.field_key
        if self.field_type is not None:
            result['fieldType'] = self.field_type
        if self.name_i18n is not None:
            result['nameI18n'] = self.name_i18n
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('fieldKey') is not None:
            self.field_key = m.get('fieldKey')
        if m.get('fieldType') is not None:
            self.field_type = m.get('fieldType')
        if m.get('nameI18n') is not None:
            self.name_i18n = m.get('nameI18n')
        return self


class UpdateTodoTypeConfigRequestActionList(TeaModel):
    def __init__(
        self,
        action_key: str = None,
        button_style_type: int = None,
        action_type: int = None,
        url: str = None,
        name_i18n: Dict[str, Any] = None,
    ):
        # 操作按钮的唯一标识
        self.action_key = action_key
        # 按钮样式类型（101：蓝色线型主按钮样式，例如「同意」，102：黑色线型副按钮样式，例如「拒绝」）
        self.button_style_type = button_style_type
        # 按钮类型（1：有操作的，2：直接跳转）
        self.action_type = action_type
        # 按钮类型为直接跳转时，对应的跳转url
        self.url = url
        # 按钮操作的显示名称（需支持国际化）
        self.name_i18n = name_i18n

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action_key is not None:
            result['actionKey'] = self.action_key
        if self.button_style_type is not None:
            result['buttonStyleType'] = self.button_style_type
        if self.action_type is not None:
            result['actionType'] = self.action_type
        if self.url is not None:
            result['url'] = self.url
        if self.name_i18n is not None:
            result['nameI18n'] = self.name_i18n
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('actionKey') is not None:
            self.action_key = m.get('actionKey')
        if m.get('buttonStyleType') is not None:
            self.button_style_type = m.get('buttonStyleType')
        if m.get('actionType') is not None:
            self.action_type = m.get('actionType')
        if m.get('url') is not None:
            self.url = m.get('url')
        if m.get('nameI18n') is not None:
            self.name_i18n = m.get('nameI18n')
        return self


class UpdateTodoTypeConfigRequest(TeaModel):
    def __init__(
        self,
        card_type: int = None,
        icon: str = None,
        description: str = None,
        pc_detail_url_open_mode: str = None,
        content_field_list: List[UpdateTodoTypeConfigRequestContentFieldList] = None,
        action_list: List[UpdateTodoTypeConfigRequestActionList] = None,
        operator_id: str = None,
    ):
        # 卡片类型（取值为：1-标准卡片，2-自定义卡片）
        self.card_type = card_type
        # 卡片类型icon，用于在待办列表展示
        self.icon = icon
        # 待办卡片类型描述
        self.description = description
        # 详情页链接在PC端的打开方式，取值为：「PC_SLIDE」-PC端侧边栏打开，「PC_BROWSER」-浏览器打开
        self.pc_detail_url_open_mode = pc_detail_url_open_mode
        # 待办卡片内容区表单自定义字段配置
        self.content_field_list = content_field_list
        # 待办卡片操作区按钮配置
        self.action_list = action_list
        # 当前操作者id，需传用户的unionId
        self.operator_id = operator_id

    def validate(self):
        if self.content_field_list:
            for k in self.content_field_list:
                if k:
                    k.validate()
        if self.action_list:
            for k in self.action_list:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.card_type is not None:
            result['cardType'] = self.card_type
        if self.icon is not None:
            result['icon'] = self.icon
        if self.description is not None:
            result['description'] = self.description
        if self.pc_detail_url_open_mode is not None:
            result['pcDetailUrlOpenMode'] = self.pc_detail_url_open_mode
        result['contentFieldList'] = []
        if self.content_field_list is not None:
            for k in self.content_field_list:
                result['contentFieldList'].append(k.to_map() if k else None)
        result['actionList'] = []
        if self.action_list is not None:
            for k in self.action_list:
                result['actionList'].append(k.to_map() if k else None)
        if self.operator_id is not None:
            result['operatorId'] = self.operator_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('cardType') is not None:
            self.card_type = m.get('cardType')
        if m.get('icon') is not None:
            self.icon = m.get('icon')
        if m.get('description') is not None:
            self.description = m.get('description')
        if m.get('pcDetailUrlOpenMode') is not None:
            self.pc_detail_url_open_mode = m.get('pcDetailUrlOpenMode')
        self.content_field_list = []
        if m.get('contentFieldList') is not None:
            for k in m.get('contentFieldList'):
                temp_model = UpdateTodoTypeConfigRequestContentFieldList()
                self.content_field_list.append(temp_model.from_map(k))
        self.action_list = []
        if m.get('actionList') is not None:
            for k in m.get('actionList'):
                temp_model = UpdateTodoTypeConfigRequestActionList()
                self.action_list.append(temp_model.from_map(k))
        if m.get('operatorId') is not None:
            self.operator_id = m.get('operatorId')
        return self


class UpdateTodoTypeConfigResponseBody(TeaModel):
    def __init__(
        self,
        result: bool = None,
    ):
        # 更新结果
        self.result = result

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.result is not None:
            result['result'] = self.result
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('result') is not None:
            self.result = m.get('result')
        return self


class UpdateTodoTypeConfigResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: UpdateTodoTypeConfigResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = UpdateTodoTypeConfigResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class GetTodoTaskHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class GetTodoTaskResponseBodyDetailUrl(TeaModel):
    def __init__(
        self,
        pc_url: str = None,
        app_url: str = None,
    ):
        # pc端详情页地址
        self.pc_url = pc_url
        # app端详情页地址
        self.app_url = app_url

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.pc_url is not None:
            result['pcUrl'] = self.pc_url
        if self.app_url is not None:
            result['appUrl'] = self.app_url
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('pcUrl') is not None:
            self.pc_url = m.get('pcUrl')
        if m.get('appUrl') is not None:
            self.app_url = m.get('appUrl')
        return self


class GetTodoTaskResponseBody(TeaModel):
    def __init__(
        self,
        id: str = None,
        subject: str = None,
        description: str = None,
        start_time: int = None,
        due_time: int = None,
        finish_time: int = None,
        done: bool = None,
        executor_ids: List[str] = None,
        participant_ids: List[str] = None,
        detail_url: GetTodoTaskResponseBodyDetailUrl = None,
        source_id: str = None,
        source: str = None,
        created_time: int = None,
        modified_time: int = None,
        creator_id: str = None,
        modifier_id: str = None,
        tenant_id: str = None,
        tenant_type: str = None,
        biz_tag: str = None,
        request_id: str = None,
        card_type_id: str = None,
    ):
        # id
        self.id = id
        # 标题
        self.subject = subject
        # 描述
        self.description = description
        # 开始时间
        self.start_time = start_time
        # 截止时间
        self.due_time = due_time
        # 完成时间
        self.finish_time = finish_time
        # 完成状态
        self.done = done
        # 执行者列表（用户的unionId）
        self.executor_ids = executor_ids
        # 参与者列表（用户的unionId）
        self.participant_ids = participant_ids
        # 自定义详情页跳转配置
        self.detail_url = detail_url
        # 业务来源id
        self.source_id = source_id
        # 业务来源
        self.source = source
        # 创建时间
        self.created_time = created_time
        # 更新时间
        self.modified_time = modified_time
        # 创建者id（用户的unionId）
        self.creator_id = creator_id
        # 更新者id（用户的unionId）
        self.modifier_id = modifier_id
        # 租户id(unionId/orgId/groupId)
        self.tenant_id = tenant_id
        # 租户类型（user/org/group）
        self.tenant_type = tenant_type
        # 接入业务应用标识
        self.biz_tag = biz_tag
        # requestId
        self.request_id = request_id
        # 待办卡片类型id
        self.card_type_id = card_type_id

    def validate(self):
        if self.detail_url:
            self.detail_url.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.id is not None:
            result['id'] = self.id
        if self.subject is not None:
            result['subject'] = self.subject
        if self.description is not None:
            result['description'] = self.description
        if self.start_time is not None:
            result['startTime'] = self.start_time
        if self.due_time is not None:
            result['dueTime'] = self.due_time
        if self.finish_time is not None:
            result['finishTime'] = self.finish_time
        if self.done is not None:
            result['done'] = self.done
        if self.executor_ids is not None:
            result['executorIds'] = self.executor_ids
        if self.participant_ids is not None:
            result['participantIds'] = self.participant_ids
        if self.detail_url is not None:
            result['detailUrl'] = self.detail_url.to_map()
        if self.source_id is not None:
            result['sourceId'] = self.source_id
        if self.source is not None:
            result['source'] = self.source
        if self.created_time is not None:
            result['createdTime'] = self.created_time
        if self.modified_time is not None:
            result['modifiedTime'] = self.modified_time
        if self.creator_id is not None:
            result['creatorId'] = self.creator_id
        if self.modifier_id is not None:
            result['modifierId'] = self.modifier_id
        if self.tenant_id is not None:
            result['tenantId'] = self.tenant_id
        if self.tenant_type is not None:
            result['tenantType'] = self.tenant_type
        if self.biz_tag is not None:
            result['bizTag'] = self.biz_tag
        if self.request_id is not None:
            result['requestId'] = self.request_id
        if self.card_type_id is not None:
            result['cardTypeId'] = self.card_type_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('id') is not None:
            self.id = m.get('id')
        if m.get('subject') is not None:
            self.subject = m.get('subject')
        if m.get('description') is not None:
            self.description = m.get('description')
        if m.get('startTime') is not None:
            self.start_time = m.get('startTime')
        if m.get('dueTime') is not None:
            self.due_time = m.get('dueTime')
        if m.get('finishTime') is not None:
            self.finish_time = m.get('finishTime')
        if m.get('done') is not None:
            self.done = m.get('done')
        if m.get('executorIds') is not None:
            self.executor_ids = m.get('executorIds')
        if m.get('participantIds') is not None:
            self.participant_ids = m.get('participantIds')
        if m.get('detailUrl') is not None:
            temp_model = GetTodoTaskResponseBodyDetailUrl()
            self.detail_url = temp_model.from_map(m['detailUrl'])
        if m.get('sourceId') is not None:
            self.source_id = m.get('sourceId')
        if m.get('source') is not None:
            self.source = m.get('source')
        if m.get('createdTime') is not None:
            self.created_time = m.get('createdTime')
        if m.get('modifiedTime') is not None:
            self.modified_time = m.get('modifiedTime')
        if m.get('creatorId') is not None:
            self.creator_id = m.get('creatorId')
        if m.get('modifierId') is not None:
            self.modifier_id = m.get('modifierId')
        if m.get('tenantId') is not None:
            self.tenant_id = m.get('tenantId')
        if m.get('tenantType') is not None:
            self.tenant_type = m.get('tenantType')
        if m.get('bizTag') is not None:
            self.biz_tag = m.get('bizTag')
        if m.get('requestId') is not None:
            self.request_id = m.get('requestId')
        if m.get('cardTypeId') is not None:
            self.card_type_id = m.get('cardTypeId')
        return self


class GetTodoTaskResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: GetTodoTaskResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = GetTodoTaskResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class DeleteTodoTaskHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class DeleteTodoTaskRequest(TeaModel):
    def __init__(
        self,
        operator_id: str = None,
    ):
        # 操作者id，需传用户的unionId
        self.operator_id = operator_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.operator_id is not None:
            result['operatorId'] = self.operator_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('operatorId') is not None:
            self.operator_id = m.get('operatorId')
        return self


class DeleteTodoTaskResponseBody(TeaModel):
    def __init__(
        self,
        result: bool = None,
        request_id: str = None,
    ):
        # 删除结果
        self.result = result
        # requestId
        self.request_id = request_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.result is not None:
            result['result'] = self.result
        if self.request_id is not None:
            result['requestId'] = self.request_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('result') is not None:
            self.result = m.get('result')
        if m.get('requestId') is not None:
            self.request_id = m.get('requestId')
        return self


class DeleteTodoTaskResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: DeleteTodoTaskResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = DeleteTodoTaskResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class UpdateTodoTaskExecutorStatusHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class UpdateTodoTaskExecutorStatusRequestExecutorStatusList(TeaModel):
    def __init__(
        self,
        id: str = None,
        is_done: bool = None,
    ):
        # 执行者id，需传用户的unionId
        self.id = id
        # 执行者完成状态
        self.is_done = is_done

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.id is not None:
            result['id'] = self.id
        if self.is_done is not None:
            result['isDone'] = self.is_done
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('id') is not None:
            self.id = m.get('id')
        if m.get('isDone') is not None:
            self.is_done = m.get('isDone')
        return self


class UpdateTodoTaskExecutorStatusRequest(TeaModel):
    def __init__(
        self,
        executor_status_list: List[UpdateTodoTaskExecutorStatusRequestExecutorStatusList] = None,
        operator_id: str = None,
    ):
        # 执行者状态列表，id需传用户的unionId
        self.executor_status_list = executor_status_list
        # 当前操作者id，需传用户的unionId
        self.operator_id = operator_id

    def validate(self):
        if self.executor_status_list:
            for k in self.executor_status_list:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        result['executorStatusList'] = []
        if self.executor_status_list is not None:
            for k in self.executor_status_list:
                result['executorStatusList'].append(k.to_map() if k else None)
        if self.operator_id is not None:
            result['operatorId'] = self.operator_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.executor_status_list = []
        if m.get('executorStatusList') is not None:
            for k in m.get('executorStatusList'):
                temp_model = UpdateTodoTaskExecutorStatusRequestExecutorStatusList()
                self.executor_status_list.append(temp_model.from_map(k))
        if m.get('operatorId') is not None:
            self.operator_id = m.get('operatorId')
        return self


class UpdateTodoTaskExecutorStatusResponseBody(TeaModel):
    def __init__(
        self,
        result: bool = None,
    ):
        # 更新结果
        self.result = result

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.result is not None:
            result['result'] = self.result
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('result') is not None:
            self.result = m.get('result')
        return self


class UpdateTodoTaskExecutorStatusResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: UpdateTodoTaskExecutorStatusResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = UpdateTodoTaskExecutorStatusResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CreateTodoTypeConfigHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class CreateTodoTypeConfigRequestContentFieldList(TeaModel):
    def __init__(
        self,
        field_key: str = None,
        field_type: str = None,
        name_i18n: Dict[str, Any] = None,
    ):
        # 字段唯一标识
        self.field_key = field_key
        # 字段类型（取值为：text-文本，url-链接）
        self.field_type = field_type
        # 字段显示名称(需支持国际化)
        self.name_i18n = name_i18n

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.field_key is not None:
            result['fieldKey'] = self.field_key
        if self.field_type is not None:
            result['fieldType'] = self.field_type
        if self.name_i18n is not None:
            result['nameI18n'] = self.name_i18n
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('fieldKey') is not None:
            self.field_key = m.get('fieldKey')
        if m.get('fieldType') is not None:
            self.field_type = m.get('fieldType')
        if m.get('nameI18n') is not None:
            self.name_i18n = m.get('nameI18n')
        return self


class CreateTodoTypeConfigRequestActionList(TeaModel):
    def __init__(
        self,
        action_key: str = None,
        button_style_type: int = None,
        action_type: int = None,
        url: str = None,
        name_i18n: Dict[str, Any] = None,
    ):
        # 操作按钮的唯一标识
        self.action_key = action_key
        # 按钮样式类型（101：蓝色线型主按钮样式，例如「同意」，102：黑色线型副按钮样式，例如「拒绝」）
        self.button_style_type = button_style_type
        # 按钮类型（1：有操作的，2：直接跳转）
        self.action_type = action_type
        # 按钮类型为直接跳转时，对应的跳转url
        self.url = url
        # 按钮操作的显示名称（需支持国际化）
        self.name_i18n = name_i18n

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action_key is not None:
            result['actionKey'] = self.action_key
        if self.button_style_type is not None:
            result['buttonStyleType'] = self.button_style_type
        if self.action_type is not None:
            result['actionType'] = self.action_type
        if self.url is not None:
            result['url'] = self.url
        if self.name_i18n is not None:
            result['nameI18n'] = self.name_i18n
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('actionKey') is not None:
            self.action_key = m.get('actionKey')
        if m.get('buttonStyleType') is not None:
            self.button_style_type = m.get('buttonStyleType')
        if m.get('actionType') is not None:
            self.action_type = m.get('actionType')
        if m.get('url') is not None:
            self.url = m.get('url')
        if m.get('nameI18n') is not None:
            self.name_i18n = m.get('nameI18n')
        return self


class CreateTodoTypeConfigRequest(TeaModel):
    def __init__(
        self,
        card_type: int = None,
        icon: str = None,
        description: str = None,
        pc_detail_url_open_mode: str = None,
        content_field_list: List[CreateTodoTypeConfigRequestContentFieldList] = None,
        action_list: List[CreateTodoTypeConfigRequestActionList] = None,
        operator_id: str = None,
    ):
        # 卡片类型（取值为：1-标准卡片，2-自定义卡片）
        self.card_type = card_type
        # 卡片类型icon，用于在待办列表展示
        self.icon = icon
        # 待办卡片类型描述
        self.description = description
        # 详情页链接在PC端的打开方式，取值为：「PC_SLIDE」-PC端侧边栏打开，「PC_BROWSER」-浏览器打开
        self.pc_detail_url_open_mode = pc_detail_url_open_mode
        # 待办卡片内容区表单自定义字段配置
        self.content_field_list = content_field_list
        # 待办卡片操作区按钮配置
        self.action_list = action_list
        # 当前操作者id，需传用户的unionId
        self.operator_id = operator_id

    def validate(self):
        if self.content_field_list:
            for k in self.content_field_list:
                if k:
                    k.validate()
        if self.action_list:
            for k in self.action_list:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.card_type is not None:
            result['cardType'] = self.card_type
        if self.icon is not None:
            result['icon'] = self.icon
        if self.description is not None:
            result['description'] = self.description
        if self.pc_detail_url_open_mode is not None:
            result['pcDetailUrlOpenMode'] = self.pc_detail_url_open_mode
        result['contentFieldList'] = []
        if self.content_field_list is not None:
            for k in self.content_field_list:
                result['contentFieldList'].append(k.to_map() if k else None)
        result['actionList'] = []
        if self.action_list is not None:
            for k in self.action_list:
                result['actionList'].append(k.to_map() if k else None)
        if self.operator_id is not None:
            result['operatorId'] = self.operator_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('cardType') is not None:
            self.card_type = m.get('cardType')
        if m.get('icon') is not None:
            self.icon = m.get('icon')
        if m.get('description') is not None:
            self.description = m.get('description')
        if m.get('pcDetailUrlOpenMode') is not None:
            self.pc_detail_url_open_mode = m.get('pcDetailUrlOpenMode')
        self.content_field_list = []
        if m.get('contentFieldList') is not None:
            for k in m.get('contentFieldList'):
                temp_model = CreateTodoTypeConfigRequestContentFieldList()
                self.content_field_list.append(temp_model.from_map(k))
        self.action_list = []
        if m.get('actionList') is not None:
            for k in m.get('actionList'):
                temp_model = CreateTodoTypeConfigRequestActionList()
                self.action_list.append(temp_model.from_map(k))
        if m.get('operatorId') is not None:
            self.operator_id = m.get('operatorId')
        return self


class CreateTodoTypeConfigResponseBodyContentFieldList(TeaModel):
    def __init__(
        self,
        field_key: str = None,
        field_type: str = None,
        name_i18n: Dict[str, Any] = None,
    ):
        # 字段唯一标识
        self.field_key = field_key
        # 字段类型（取值为：text-文本，url-链接）
        self.field_type = field_type
        # 字段的显示名称（支持国际化）
        self.name_i18n = name_i18n

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.field_key is not None:
            result['fieldKey'] = self.field_key
        if self.field_type is not None:
            result['fieldType'] = self.field_type
        if self.name_i18n is not None:
            result['nameI18n'] = self.name_i18n
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('fieldKey') is not None:
            self.field_key = m.get('fieldKey')
        if m.get('fieldType') is not None:
            self.field_type = m.get('fieldType')
        if m.get('nameI18n') is not None:
            self.name_i18n = m.get('nameI18n')
        return self


class CreateTodoTypeConfigResponseBodyActionList(TeaModel):
    def __init__(
        self,
        action_key: str = None,
        button_style_type: int = None,
        action_type: int = None,
        url: str = None,
        name_i18n: Dict[str, Any] = None,
    ):
        # 操作按钮的唯一标识
        self.action_key = action_key
        # 按钮样式类型（101：蓝色线型主按钮样式，例如「同意」，102：黑色线型副按钮样式，例如「拒绝」）
        self.button_style_type = button_style_type
        # 按钮类型（1：有操作的，2：直接跳转）
        self.action_type = action_type
        # 按钮类型为直接跳转时，对应的跳转url
        self.url = url
        # 按钮操作的显示名称（支持国际化）
        self.name_i18n = name_i18n

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.action_key is not None:
            result['actionKey'] = self.action_key
        if self.button_style_type is not None:
            result['buttonStyleType'] = self.button_style_type
        if self.action_type is not None:
            result['actionType'] = self.action_type
        if self.url is not None:
            result['url'] = self.url
        if self.name_i18n is not None:
            result['nameI18n'] = self.name_i18n
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('actionKey') is not None:
            self.action_key = m.get('actionKey')
        if m.get('buttonStyleType') is not None:
            self.button_style_type = m.get('buttonStyleType')
        if m.get('actionType') is not None:
            self.action_type = m.get('actionType')
        if m.get('url') is not None:
            self.url = m.get('url')
        if m.get('nameI18n') is not None:
            self.name_i18n = m.get('nameI18n')
        return self


class CreateTodoTypeConfigResponseBody(TeaModel):
    def __init__(
        self,
        id: str = None,
        created_time: int = None,
        modified_time: int = None,
        creator_id: str = None,
        modifier_id: str = None,
        biz_tag: str = None,
        request_id: str = None,
        card_type: int = None,
        icon: str = None,
        description: str = None,
        pc_detail_url_open_mode: str = None,
        content_field_list: List[CreateTodoTypeConfigResponseBodyContentFieldList] = None,
        action_list: List[CreateTodoTypeConfigResponseBodyActionList] = None,
    ):
        # id
        self.id = id
        # 创建时间
        self.created_time = created_time
        # 更新时间
        self.modified_time = modified_time
        # 创建者（用户的unionId）
        self.creator_id = creator_id
        # 更新者（用户的unionId）
        self.modifier_id = modifier_id
        # 接入应用标识
        self.biz_tag = biz_tag
        # requestId
        self.request_id = request_id
        # 卡片类型（取值为：1-标准卡片，2-自定义卡片）
        self.card_type = card_type
        # 卡片类型icon，用于在待办列表展示
        self.icon = icon
        # 待办卡片类型描述
        self.description = description
        # 详情页链接在PC端的打开方式，取值为：「PC_SLIDE」-PC端侧边栏打开，「PC_BROWSER」-浏览器打开
        self.pc_detail_url_open_mode = pc_detail_url_open_mode
        # 待办卡片内容区表单自定义字段配置
        self.content_field_list = content_field_list
        # 待办卡片操作区按钮配置
        self.action_list = action_list

    def validate(self):
        if self.content_field_list:
            for k in self.content_field_list:
                if k:
                    k.validate()
        if self.action_list:
            for k in self.action_list:
                if k:
                    k.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.id is not None:
            result['id'] = self.id
        if self.created_time is not None:
            result['createdTime'] = self.created_time
        if self.modified_time is not None:
            result['modifiedTime'] = self.modified_time
        if self.creator_id is not None:
            result['creatorId'] = self.creator_id
        if self.modifier_id is not None:
            result['modifierId'] = self.modifier_id
        if self.biz_tag is not None:
            result['bizTag'] = self.biz_tag
        if self.request_id is not None:
            result['requestId'] = self.request_id
        if self.card_type is not None:
            result['cardType'] = self.card_type
        if self.icon is not None:
            result['icon'] = self.icon
        if self.description is not None:
            result['description'] = self.description
        if self.pc_detail_url_open_mode is not None:
            result['pcDetailUrlOpenMode'] = self.pc_detail_url_open_mode
        result['contentFieldList'] = []
        if self.content_field_list is not None:
            for k in self.content_field_list:
                result['contentFieldList'].append(k.to_map() if k else None)
        result['actionList'] = []
        if self.action_list is not None:
            for k in self.action_list:
                result['actionList'].append(k.to_map() if k else None)
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('id') is not None:
            self.id = m.get('id')
        if m.get('createdTime') is not None:
            self.created_time = m.get('createdTime')
        if m.get('modifiedTime') is not None:
            self.modified_time = m.get('modifiedTime')
        if m.get('creatorId') is not None:
            self.creator_id = m.get('creatorId')
        if m.get('modifierId') is not None:
            self.modifier_id = m.get('modifierId')
        if m.get('bizTag') is not None:
            self.biz_tag = m.get('bizTag')
        if m.get('requestId') is not None:
            self.request_id = m.get('requestId')
        if m.get('cardType') is not None:
            self.card_type = m.get('cardType')
        if m.get('icon') is not None:
            self.icon = m.get('icon')
        if m.get('description') is not None:
            self.description = m.get('description')
        if m.get('pcDetailUrlOpenMode') is not None:
            self.pc_detail_url_open_mode = m.get('pcDetailUrlOpenMode')
        self.content_field_list = []
        if m.get('contentFieldList') is not None:
            for k in m.get('contentFieldList'):
                temp_model = CreateTodoTypeConfigResponseBodyContentFieldList()
                self.content_field_list.append(temp_model.from_map(k))
        self.action_list = []
        if m.get('actionList') is not None:
            for k in m.get('actionList'):
                temp_model = CreateTodoTypeConfigResponseBodyActionList()
                self.action_list.append(temp_model.from_map(k))
        return self


class CreateTodoTypeConfigResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: CreateTodoTypeConfigResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = CreateTodoTypeConfigResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CountTodoTasksHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class CountTodoTasksRequest(TeaModel):
    def __init__(
        self,
        is_done: bool = None,
        role_types: List[List[str]] = None,
        from_due_time: int = None,
        to_due_time: int = None,
        category: str = None,
        is_recycled: bool = None,
    ):
        # 待办完成状态。
        self.is_done = is_done
        # 查询目标用户角色类型，执行人 | 创建人 | 参与人，可以同时传多个值。如：[["executor"], ["creator"],["participant"]] 或 [["executor", "creator"]]
        self.role_types = role_types
        # 查询从计划完成时间开始
        self.from_due_time = from_due_time
        # 查询到计划完成时间结束
        self.to_due_time = to_due_time
        # 所属分类
        self.category = category
        # 待办回收状态
        self.is_recycled = is_recycled

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.is_done is not None:
            result['isDone'] = self.is_done
        if self.role_types is not None:
            result['roleTypes'] = self.role_types
        if self.from_due_time is not None:
            result['fromDueTime'] = self.from_due_time
        if self.to_due_time is not None:
            result['toDueTime'] = self.to_due_time
        if self.category is not None:
            result['category'] = self.category
        if self.is_recycled is not None:
            result['isRecycled'] = self.is_recycled
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('isDone') is not None:
            self.is_done = m.get('isDone')
        if m.get('roleTypes') is not None:
            self.role_types = m.get('roleTypes')
        if m.get('fromDueTime') is not None:
            self.from_due_time = m.get('fromDueTime')
        if m.get('toDueTime') is not None:
            self.to_due_time = m.get('toDueTime')
        if m.get('category') is not None:
            self.category = m.get('category')
        if m.get('isRecycled') is not None:
            self.is_recycled = m.get('isRecycled')
        return self


class CountTodoTasksResponseBody(TeaModel):
    def __init__(
        self,
        result: int = None,
    ):
        # 待办数量
        self.result = result

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.result is not None:
            result['result'] = self.result
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('result') is not None:
            self.result = m.get('result')
        return self


class CountTodoTasksResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: CountTodoTasksResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = CountTodoTasksResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class UpdateTodoTaskHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class UpdateTodoTaskRequest(TeaModel):
    def __init__(
        self,
        subject: str = None,
        description: str = None,
        due_time: int = None,
        done: bool = None,
        executor_ids: List[str] = None,
        participant_ids: List[str] = None,
        operator_id: str = None,
    ):
        # 待办标题
        self.subject = subject
        # 待办描述备注
        self.description = description
        # 截止时间
        self.due_time = due_time
        # 完成状态
        self.done = done
        # 执行者列表，需传用户的unionId
        self.executor_ids = executor_ids
        # 参与者列表，需传用户的unionId
        self.participant_ids = participant_ids
        # 当前操作者id，需传用户的unionId
        self.operator_id = operator_id

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.subject is not None:
            result['subject'] = self.subject
        if self.description is not None:
            result['description'] = self.description
        if self.due_time is not None:
            result['dueTime'] = self.due_time
        if self.done is not None:
            result['done'] = self.done
        if self.executor_ids is not None:
            result['executorIds'] = self.executor_ids
        if self.participant_ids is not None:
            result['participantIds'] = self.participant_ids
        if self.operator_id is not None:
            result['operatorId'] = self.operator_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('subject') is not None:
            self.subject = m.get('subject')
        if m.get('description') is not None:
            self.description = m.get('description')
        if m.get('dueTime') is not None:
            self.due_time = m.get('dueTime')
        if m.get('done') is not None:
            self.done = m.get('done')
        if m.get('executorIds') is not None:
            self.executor_ids = m.get('executorIds')
        if m.get('participantIds') is not None:
            self.participant_ids = m.get('participantIds')
        if m.get('operatorId') is not None:
            self.operator_id = m.get('operatorId')
        return self


class UpdateTodoTaskResponseBody(TeaModel):
    def __init__(
        self,
        result: bool = None,
    ):
        # 更新结果
        self.result = result

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.result is not None:
            result['result'] = self.result
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('result') is not None:
            self.result = m.get('result')
        return self


class UpdateTodoTaskResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: UpdateTodoTaskResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = UpdateTodoTaskResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


class CreateTodoTaskHeaders(TeaModel):
    def __init__(
        self,
        common_headers: Dict[str, str] = None,
        x_acs_dingtalk_access_token: str = None,
    ):
        self.common_headers = common_headers
        self.x_acs_dingtalk_access_token = x_acs_dingtalk_access_token

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.common_headers is not None:
            result['commonHeaders'] = self.common_headers
        if self.x_acs_dingtalk_access_token is not None:
            result['x-acs-dingtalk-access-token'] = self.x_acs_dingtalk_access_token
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('commonHeaders') is not None:
            self.common_headers = m.get('commonHeaders')
        if m.get('x-acs-dingtalk-access-token') is not None:
            self.x_acs_dingtalk_access_token = m.get('x-acs-dingtalk-access-token')
        return self


class CreateTodoTaskRequestDetailUrl(TeaModel):
    def __init__(
        self,
        app_url: str = None,
        pc_url: str = None,
    ):
        # app端详情页url
        self.app_url = app_url
        # pc端详情页url
        self.pc_url = pc_url

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.app_url is not None:
            result['appUrl'] = self.app_url
        if self.pc_url is not None:
            result['pcUrl'] = self.pc_url
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('appUrl') is not None:
            self.app_url = m.get('appUrl')
        if m.get('pcUrl') is not None:
            self.pc_url = m.get('pcUrl')
        return self


class CreateTodoTaskRequest(TeaModel):
    def __init__(
        self,
        source_id: str = None,
        subject: str = None,
        creator_id: str = None,
        description: str = None,
        due_time: int = None,
        executor_ids: List[str] = None,
        participant_ids: List[str] = None,
        detail_url: CreateTodoTaskRequestDetailUrl = None,
        operator_id: str = None,
    ):
        # 来源id，接入业务系统侧的唯一标识id
        self.source_id = source_id
        # 待办标题
        self.subject = subject
        # 创建者id，需传用户的unionId
        self.creator_id = creator_id
        # 待办备注描述
        self.description = description
        # 截止时间
        self.due_time = due_time
        # 执行者列表，需传用户的unionId
        self.executor_ids = executor_ids
        # 参与者列表，需传用户的unionId
        self.participant_ids = participant_ids
        # 详情页url跳转地址
        self.detail_url = detail_url
        # 当前操作者id，需传用户的unionId
        self.operator_id = operator_id

    def validate(self):
        if self.detail_url:
            self.detail_url.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.source_id is not None:
            result['sourceId'] = self.source_id
        if self.subject is not None:
            result['subject'] = self.subject
        if self.creator_id is not None:
            result['creatorId'] = self.creator_id
        if self.description is not None:
            result['description'] = self.description
        if self.due_time is not None:
            result['dueTime'] = self.due_time
        if self.executor_ids is not None:
            result['executorIds'] = self.executor_ids
        if self.participant_ids is not None:
            result['participantIds'] = self.participant_ids
        if self.detail_url is not None:
            result['detailUrl'] = self.detail_url.to_map()
        if self.operator_id is not None:
            result['operatorId'] = self.operator_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('sourceId') is not None:
            self.source_id = m.get('sourceId')
        if m.get('subject') is not None:
            self.subject = m.get('subject')
        if m.get('creatorId') is not None:
            self.creator_id = m.get('creatorId')
        if m.get('description') is not None:
            self.description = m.get('description')
        if m.get('dueTime') is not None:
            self.due_time = m.get('dueTime')
        if m.get('executorIds') is not None:
            self.executor_ids = m.get('executorIds')
        if m.get('participantIds') is not None:
            self.participant_ids = m.get('participantIds')
        if m.get('detailUrl') is not None:
            temp_model = CreateTodoTaskRequestDetailUrl()
            self.detail_url = temp_model.from_map(m['detailUrl'])
        if m.get('operatorId') is not None:
            self.operator_id = m.get('operatorId')
        return self


class CreateTodoTaskResponseBodyDetailUrl(TeaModel):
    def __init__(
        self,
        pc_url: str = None,
        app_url: str = None,
    ):
        # pc端详情页地址
        self.pc_url = pc_url
        # app端详情页地址
        self.app_url = app_url

    def validate(self):
        pass

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.pc_url is not None:
            result['pcUrl'] = self.pc_url
        if self.app_url is not None:
            result['appUrl'] = self.app_url
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('pcUrl') is not None:
            self.pc_url = m.get('pcUrl')
        if m.get('appUrl') is not None:
            self.app_url = m.get('appUrl')
        return self


class CreateTodoTaskResponseBody(TeaModel):
    def __init__(
        self,
        id: str = None,
        subject: str = None,
        description: str = None,
        start_time: int = None,
        due_time: int = None,
        finish_time: int = None,
        done: bool = None,
        executor_ids: List[str] = None,
        participant_ids: List[str] = None,
        detail_url: CreateTodoTaskResponseBodyDetailUrl = None,
        source: str = None,
        source_id: str = None,
        created_time: int = None,
        modified_time: int = None,
        creator_id: str = None,
        modifier_id: str = None,
        biz_tag: str = None,
        request_id: str = None,
    ):
        # id
        self.id = id
        # 标题
        self.subject = subject
        # 描述
        self.description = description
        # 开始时间
        self.start_time = start_time
        # 截止时间
        self.due_time = due_time
        # 完成时间
        self.finish_time = finish_time
        # 完成状态
        self.done = done
        # 执行者列表（用户的unionId）
        self.executor_ids = executor_ids
        # 参与者列表（用户的unionId）
        self.participant_ids = participant_ids
        # 自定义详情页跳转配置
        self.detail_url = detail_url
        # 业务来源
        self.source = source
        # 业务来源id
        self.source_id = source_id
        # 创建时间
        self.created_time = created_time
        # 更新时间
        self.modified_time = modified_time
        # 创建者（用户的unionId）
        self.creator_id = creator_id
        # 更新者（用户的unionId）
        self.modifier_id = modifier_id
        # 接入应用标识
        self.biz_tag = biz_tag
        # requestId
        self.request_id = request_id

    def validate(self):
        if self.detail_url:
            self.detail_url.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.id is not None:
            result['id'] = self.id
        if self.subject is not None:
            result['subject'] = self.subject
        if self.description is not None:
            result['description'] = self.description
        if self.start_time is not None:
            result['startTime'] = self.start_time
        if self.due_time is not None:
            result['dueTime'] = self.due_time
        if self.finish_time is not None:
            result['finishTime'] = self.finish_time
        if self.done is not None:
            result['done'] = self.done
        if self.executor_ids is not None:
            result['executorIds'] = self.executor_ids
        if self.participant_ids is not None:
            result['participantIds'] = self.participant_ids
        if self.detail_url is not None:
            result['detailUrl'] = self.detail_url.to_map()
        if self.source is not None:
            result['source'] = self.source
        if self.source_id is not None:
            result['sourceId'] = self.source_id
        if self.created_time is not None:
            result['createdTime'] = self.created_time
        if self.modified_time is not None:
            result['modifiedTime'] = self.modified_time
        if self.creator_id is not None:
            result['creatorId'] = self.creator_id
        if self.modifier_id is not None:
            result['modifierId'] = self.modifier_id
        if self.biz_tag is not None:
            result['bizTag'] = self.biz_tag
        if self.request_id is not None:
            result['requestId'] = self.request_id
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('id') is not None:
            self.id = m.get('id')
        if m.get('subject') is not None:
            self.subject = m.get('subject')
        if m.get('description') is not None:
            self.description = m.get('description')
        if m.get('startTime') is not None:
            self.start_time = m.get('startTime')
        if m.get('dueTime') is not None:
            self.due_time = m.get('dueTime')
        if m.get('finishTime') is not None:
            self.finish_time = m.get('finishTime')
        if m.get('done') is not None:
            self.done = m.get('done')
        if m.get('executorIds') is not None:
            self.executor_ids = m.get('executorIds')
        if m.get('participantIds') is not None:
            self.participant_ids = m.get('participantIds')
        if m.get('detailUrl') is not None:
            temp_model = CreateTodoTaskResponseBodyDetailUrl()
            self.detail_url = temp_model.from_map(m['detailUrl'])
        if m.get('source') is not None:
            self.source = m.get('source')
        if m.get('sourceId') is not None:
            self.source_id = m.get('sourceId')
        if m.get('createdTime') is not None:
            self.created_time = m.get('createdTime')
        if m.get('modifiedTime') is not None:
            self.modified_time = m.get('modifiedTime')
        if m.get('creatorId') is not None:
            self.creator_id = m.get('creatorId')
        if m.get('modifierId') is not None:
            self.modifier_id = m.get('modifierId')
        if m.get('bizTag') is not None:
            self.biz_tag = m.get('bizTag')
        if m.get('requestId') is not None:
            self.request_id = m.get('requestId')
        return self


class CreateTodoTaskResponse(TeaModel):
    def __init__(
        self,
        headers: Dict[str, str] = None,
        body: CreateTodoTaskResponseBody = None,
    ):
        self.headers = headers
        self.body = body

    def validate(self):
        self.validate_required(self.headers, 'headers')
        self.validate_required(self.body, 'body')
        if self.body:
            self.body.validate()

    def to_map(self):
        _map = super().to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body.to_map()
        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            temp_model = CreateTodoTaskResponseBody()
            self.body = temp_model.from_map(m['body'])
        return self


