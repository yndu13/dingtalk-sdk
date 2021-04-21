# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.core import TeaCore

from alibabacloud_tea_openapi.client import Client as OpenApiClient
from alibabacloud_tea_openapi import models as open_api_models
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_dingtalkservice_group_1_0 import models as dingtalkservice_group__1__0_models
from alibabacloud_tea_util import models as util_models
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient


class Client(OpenApiClient):
    """
    *\
    """
    def __init__(
        self, 
        config: open_api_models.Config,
    ):
        super().__init__(config)
        self._endpoint_rule = ''
        if UtilClient.empty(self._endpoint):
            self._endpoint = 'api.dingtalk.com'

    def delete_knowledge(
        self,
        request: dingtalkservice_group__1__0_models.DeleteKnowledgeRequest,
    ) -> dingtalkservice_group__1__0_models.DeleteKnowledgeResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.DeleteKnowledgeHeaders()
        return self.delete_knowledge_with_options(request, headers, runtime)

    async def delete_knowledge_async(
        self,
        request: dingtalkservice_group__1__0_models.DeleteKnowledgeRequest,
    ) -> dingtalkservice_group__1__0_models.DeleteKnowledgeResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.DeleteKnowledgeHeaders()
        return await self.delete_knowledge_with_options_async(request, headers, runtime)

    def delete_knowledge_with_options(
        self,
        request: dingtalkservice_group__1__0_models.DeleteKnowledgeRequest,
        headers: dingtalkservice_group__1__0_models.DeleteKnowledgeHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.DeleteKnowledgeResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        if not UtilClient.is_unset(request.open_team_id):
            body['openTeamId'] = request.open_team_id
        if not UtilClient.is_unset(request.library_key):
            body['libraryKey'] = request.library_key
        if not UtilClient.is_unset(request.source):
            body['source'] = request.source
        if not UtilClient.is_unset(request.source_primary_key):
            body['sourcePrimaryKey'] = request.source_primary_key
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.DeleteKnowledgeResponse(),
            self.do_roarequest('DeleteKnowledge', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/knowledges/batchDelete', 'json', req, runtime)
        )

    async def delete_knowledge_with_options_async(
        self,
        request: dingtalkservice_group__1__0_models.DeleteKnowledgeRequest,
        headers: dingtalkservice_group__1__0_models.DeleteKnowledgeHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.DeleteKnowledgeResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        if not UtilClient.is_unset(request.open_team_id):
            body['openTeamId'] = request.open_team_id
        if not UtilClient.is_unset(request.library_key):
            body['libraryKey'] = request.library_key
        if not UtilClient.is_unset(request.source):
            body['source'] = request.source
        if not UtilClient.is_unset(request.source_primary_key):
            body['sourcePrimaryKey'] = request.source_primary_key
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.DeleteKnowledgeResponse(),
            await self.do_roarequest_async('DeleteKnowledge', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/knowledges/batchDelete', 'json', req, runtime)
        )

    def create_group(
        self,
        request: dingtalkservice_group__1__0_models.CreateGroupRequest,
    ) -> dingtalkservice_group__1__0_models.CreateGroupResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.CreateGroupHeaders()
        return self.create_group_with_options(request, headers, runtime)

    async def create_group_async(
        self,
        request: dingtalkservice_group__1__0_models.CreateGroupRequest,
    ) -> dingtalkservice_group__1__0_models.CreateGroupResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.CreateGroupHeaders()
        return await self.create_group_with_options_async(request, headers, runtime)

    def create_group_with_options(
        self,
        request: dingtalkservice_group__1__0_models.CreateGroupRequest,
        headers: dingtalkservice_group__1__0_models.CreateGroupHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.CreateGroupResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.group_biz_id):
            body['groupBizId'] = request.group_biz_id
        if not UtilClient.is_unset(request.open_team_id):
            body['openTeamId'] = request.open_team_id
        if not UtilClient.is_unset(request.open_group_set_id):
            body['openGroupSetId'] = request.open_group_set_id
        if not UtilClient.is_unset(request.group_name):
            body['groupName'] = request.group_name
        if not UtilClient.is_unset(request.owner_staff_id):
            body['ownerStaffId'] = request.owner_staff_id
        if not UtilClient.is_unset(request.member_staff_ids):
            body['memberStaffIds'] = request.member_staff_ids
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.CreateGroupResponse(),
            self.do_roarequest('CreateGroup', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/groups', 'json', req, runtime)
        )

    async def create_group_with_options_async(
        self,
        request: dingtalkservice_group__1__0_models.CreateGroupRequest,
        headers: dingtalkservice_group__1__0_models.CreateGroupHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.CreateGroupResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.group_biz_id):
            body['groupBizId'] = request.group_biz_id
        if not UtilClient.is_unset(request.open_team_id):
            body['openTeamId'] = request.open_team_id
        if not UtilClient.is_unset(request.open_group_set_id):
            body['openGroupSetId'] = request.open_group_set_id
        if not UtilClient.is_unset(request.group_name):
            body['groupName'] = request.group_name
        if not UtilClient.is_unset(request.owner_staff_id):
            body['ownerStaffId'] = request.owner_staff_id
        if not UtilClient.is_unset(request.member_staff_ids):
            body['memberStaffIds'] = request.member_staff_ids
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.CreateGroupResponse(),
            await self.do_roarequest_async('CreateGroup', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/groups', 'json', req, runtime)
        )

    def send_service_group_message(
        self,
        request: dingtalkservice_group__1__0_models.SendServiceGroupMessageRequest,
    ) -> dingtalkservice_group__1__0_models.SendServiceGroupMessageResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.SendServiceGroupMessageHeaders()
        return self.send_service_group_message_with_options(request, headers, runtime)

    async def send_service_group_message_async(
        self,
        request: dingtalkservice_group__1__0_models.SendServiceGroupMessageRequest,
    ) -> dingtalkservice_group__1__0_models.SendServiceGroupMessageResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.SendServiceGroupMessageHeaders()
        return await self.send_service_group_message_with_options_async(request, headers, runtime)

    def send_service_group_message_with_options(
        self,
        request: dingtalkservice_group__1__0_models.SendServiceGroupMessageRequest,
        headers: dingtalkservice_group__1__0_models.SendServiceGroupMessageHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.SendServiceGroupMessageResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.target_open_conversation_id):
            body['targetOpenConversationId'] = request.target_open_conversation_id
        if not UtilClient.is_unset(request.title):
            body['title'] = request.title
        if not UtilClient.is_unset(request.content):
            body['content'] = request.content
        if not UtilClient.is_unset(request.at_mobiles):
            body['atMobiles'] = request.at_mobiles
        if not UtilClient.is_unset(request.at_dingtalk_ids):
            body['atDingtalkIds'] = request.at_dingtalk_ids
        if not UtilClient.is_unset(request.at_union_ids):
            body['atUnionIds'] = request.at_union_ids
        if not UtilClient.is_unset(request.receiver_mobiles):
            body['receiverMobiles'] = request.receiver_mobiles
        if not UtilClient.is_unset(request.receiver_dingtalk_ids):
            body['receiverDingtalkIds'] = request.receiver_dingtalk_ids
        if not UtilClient.is_unset(request.receiver_union_ids):
            body['receiverUnionIds'] = request.receiver_union_ids
        if not UtilClient.is_unset(request.message_type):
            body['messageType'] = request.message_type
        if not UtilClient.is_unset(request.btn_orientation):
            body['btnOrientation'] = request.btn_orientation
        if not UtilClient.is_unset(request.btns):
            body['btns'] = request.btns
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.SendServiceGroupMessageResponse(),
            self.do_roarequest('SendServiceGroupMessage', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/messages/send', 'json', req, runtime)
        )

    async def send_service_group_message_with_options_async(
        self,
        request: dingtalkservice_group__1__0_models.SendServiceGroupMessageRequest,
        headers: dingtalkservice_group__1__0_models.SendServiceGroupMessageHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.SendServiceGroupMessageResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.target_open_conversation_id):
            body['targetOpenConversationId'] = request.target_open_conversation_id
        if not UtilClient.is_unset(request.title):
            body['title'] = request.title
        if not UtilClient.is_unset(request.content):
            body['content'] = request.content
        if not UtilClient.is_unset(request.at_mobiles):
            body['atMobiles'] = request.at_mobiles
        if not UtilClient.is_unset(request.at_dingtalk_ids):
            body['atDingtalkIds'] = request.at_dingtalk_ids
        if not UtilClient.is_unset(request.at_union_ids):
            body['atUnionIds'] = request.at_union_ids
        if not UtilClient.is_unset(request.receiver_mobiles):
            body['receiverMobiles'] = request.receiver_mobiles
        if not UtilClient.is_unset(request.receiver_dingtalk_ids):
            body['receiverDingtalkIds'] = request.receiver_dingtalk_ids
        if not UtilClient.is_unset(request.receiver_union_ids):
            body['receiverUnionIds'] = request.receiver_union_ids
        if not UtilClient.is_unset(request.message_type):
            body['messageType'] = request.message_type
        if not UtilClient.is_unset(request.btn_orientation):
            body['btnOrientation'] = request.btn_orientation
        if not UtilClient.is_unset(request.btns):
            body['btns'] = request.btns
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.SendServiceGroupMessageResponse(),
            await self.do_roarequest_async('SendServiceGroupMessage', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/messages/send', 'json', req, runtime)
        )

    def add_knowledge(
        self,
        request: dingtalkservice_group__1__0_models.AddKnowledgeRequest,
    ) -> dingtalkservice_group__1__0_models.AddKnowledgeResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.AddKnowledgeHeaders()
        return self.add_knowledge_with_options(request, headers, runtime)

    async def add_knowledge_async(
        self,
        request: dingtalkservice_group__1__0_models.AddKnowledgeRequest,
    ) -> dingtalkservice_group__1__0_models.AddKnowledgeResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.AddKnowledgeHeaders()
        return await self.add_knowledge_with_options_async(request, headers, runtime)

    def add_knowledge_with_options(
        self,
        request: dingtalkservice_group__1__0_models.AddKnowledgeRequest,
        headers: dingtalkservice_group__1__0_models.AddKnowledgeHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.AddKnowledgeResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        if not UtilClient.is_unset(request.open_team_id):
            body['openTeamId'] = request.open_team_id
        if not UtilClient.is_unset(request.library_key):
            body['libraryKey'] = request.library_key
        if not UtilClient.is_unset(request.source):
            body['source'] = request.source
        if not UtilClient.is_unset(request.source_primary_key):
            body['sourcePrimaryKey'] = request.source_primary_key
        if not UtilClient.is_unset(request.type):
            body['type'] = request.type
        if not UtilClient.is_unset(request.title):
            body['title'] = request.title
        if not UtilClient.is_unset(request.content):
            body['content'] = request.content
        if not UtilClient.is_unset(request.link_url):
            body['linkUrl'] = request.link_url
        if not UtilClient.is_unset(request.version):
            body['version'] = request.version
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.AddKnowledgeResponse(),
            self.do_roarequest('AddKnowledge', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/knowledges', 'json', req, runtime)
        )

    async def add_knowledge_with_options_async(
        self,
        request: dingtalkservice_group__1__0_models.AddKnowledgeRequest,
        headers: dingtalkservice_group__1__0_models.AddKnowledgeHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.AddKnowledgeResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        if not UtilClient.is_unset(request.open_team_id):
            body['openTeamId'] = request.open_team_id
        if not UtilClient.is_unset(request.library_key):
            body['libraryKey'] = request.library_key
        if not UtilClient.is_unset(request.source):
            body['source'] = request.source
        if not UtilClient.is_unset(request.source_primary_key):
            body['sourcePrimaryKey'] = request.source_primary_key
        if not UtilClient.is_unset(request.type):
            body['type'] = request.type
        if not UtilClient.is_unset(request.title):
            body['title'] = request.title
        if not UtilClient.is_unset(request.content):
            body['content'] = request.content
        if not UtilClient.is_unset(request.link_url):
            body['linkUrl'] = request.link_url
        if not UtilClient.is_unset(request.version):
            body['version'] = request.version
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.AddKnowledgeResponse(),
            await self.do_roarequest_async('AddKnowledge', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/knowledges', 'json', req, runtime)
        )

    def add_library(
        self,
        request: dingtalkservice_group__1__0_models.AddLibraryRequest,
    ) -> dingtalkservice_group__1__0_models.AddLibraryResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.AddLibraryHeaders()
        return self.add_library_with_options(request, headers, runtime)

    async def add_library_async(
        self,
        request: dingtalkservice_group__1__0_models.AddLibraryRequest,
    ) -> dingtalkservice_group__1__0_models.AddLibraryResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.AddLibraryHeaders()
        return await self.add_library_with_options_async(request, headers, runtime)

    def add_library_with_options(
        self,
        request: dingtalkservice_group__1__0_models.AddLibraryRequest,
        headers: dingtalkservice_group__1__0_models.AddLibraryHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.AddLibraryResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.open_team_ids):
            body['openTeamIds'] = request.open_team_ids
        if not UtilClient.is_unset(request.title):
            body['title'] = request.title
        if not UtilClient.is_unset(request.description):
            body['description'] = request.description
        if not UtilClient.is_unset(request.type):
            body['type'] = request.type
        if not UtilClient.is_unset(request.source):
            body['source'] = request.source
        if not UtilClient.is_unset(request.source_primary_key):
            body['sourcePrimaryKey'] = request.source_primary_key
        if not UtilClient.is_unset(request.user_id):
            body['userId'] = request.user_id
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.AddLibraryResponse(),
            self.do_roarequest('AddLibrary', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/librarys', 'json', req, runtime)
        )

    async def add_library_with_options_async(
        self,
        request: dingtalkservice_group__1__0_models.AddLibraryRequest,
        headers: dingtalkservice_group__1__0_models.AddLibraryHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.AddLibraryResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.ding_token_grant_type):
            body['dingTokenGrantType'] = request.ding_token_grant_type
        if not UtilClient.is_unset(request.ding_isv_org_id):
            body['dingIsvOrgId'] = request.ding_isv_org_id
        if not UtilClient.is_unset(request.ding_suite_key):
            body['dingSuiteKey'] = request.ding_suite_key
        if not UtilClient.is_unset(request.ding_org_id):
            body['dingOrgId'] = request.ding_org_id
        if not UtilClient.is_unset(request.open_team_ids):
            body['openTeamIds'] = request.open_team_ids
        if not UtilClient.is_unset(request.title):
            body['title'] = request.title
        if not UtilClient.is_unset(request.description):
            body['description'] = request.description
        if not UtilClient.is_unset(request.type):
            body['type'] = request.type
        if not UtilClient.is_unset(request.source):
            body['source'] = request.source
        if not UtilClient.is_unset(request.source_primary_key):
            body['sourcePrimaryKey'] = request.source_primary_key
        if not UtilClient.is_unset(request.user_id):
            body['userId'] = request.user_id
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.AddLibraryResponse(),
            await self.do_roarequest_async('AddLibrary', 'serviceGroup_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/serviceGroup/librarys', 'json', req, runtime)
        )

    def list_user_teams(
        self,
        user_id: str,
    ) -> dingtalkservice_group__1__0_models.ListUserTeamsResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.ListUserTeamsHeaders()
        return self.list_user_teams_with_options(user_id, headers, runtime)

    async def list_user_teams_async(
        self,
        user_id: str,
    ) -> dingtalkservice_group__1__0_models.ListUserTeamsResponse:
        runtime = util_models.RuntimeOptions()
        headers = dingtalkservice_group__1__0_models.ListUserTeamsHeaders()
        return await self.list_user_teams_with_options_async(user_id, headers, runtime)

    def list_user_teams_with_options(
        self,
        user_id: str,
        headers: dingtalkservice_group__1__0_models.ListUserTeamsHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.ListUserTeamsResponse:
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.ListUserTeamsResponse(),
            self.do_roarequest('ListUserTeams', 'serviceGroup_1.0', 'HTTP', 'GET', 'AK', f'/v1.0/serviceGroup/users/{user_id}/teams', 'json', req, runtime)
        )

    async def list_user_teams_with_options_async(
        self,
        user_id: str,
        headers: dingtalkservice_group__1__0_models.ListUserTeamsHeaders,
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkservice_group__1__0_models.ListUserTeamsResponse:
        real_headers = {}
        if not UtilClient.is_unset(headers.common_headers):
            real_headers = headers.common_headers
        if not UtilClient.is_unset(headers.x_acs_dingtalk_access_token):
            real_headers['x-acs-dingtalk-access-token'] = headers.x_acs_dingtalk_access_token
        req = open_api_models.OpenApiRequest(
            headers=real_headers
        )
        return TeaCore.from_map(
            dingtalkservice_group__1__0_models.ListUserTeamsResponse(),
            await self.do_roarequest_async('ListUserTeams', 'serviceGroup_1.0', 'HTTP', 'GET', 'AK', f'/v1.0/serviceGroup/users/{user_id}/teams', 'json', req, runtime)
        )
