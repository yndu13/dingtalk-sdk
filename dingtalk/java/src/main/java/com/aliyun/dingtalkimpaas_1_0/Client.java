// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkimpaas_1_0;

import com.aliyun.tea.*;
import com.aliyun.dingtalkimpaas_1_0.models.*;
import com.aliyun.teautil.*;
import com.aliyun.teautil.models.*;
import com.aliyun.teaopenapi.*;
import com.aliyun.teaopenapi.models.*;
import com.aliyun.openapiutil.*;

public class Client extends com.aliyun.teaopenapi.Client {

    public Client(Config config) throws Exception {
        super(config);
        this._endpointRule = "";
        if (com.aliyun.teautil.Common.empty(_endpoint)) {
            this._endpoint = "api.dingtalk.com";
        }

    }


    public GetConversationIdResponse getConversationId(GetConversationIdRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        GetConversationIdHeaders headers = new GetConversationIdHeaders();
        return this.getConversationIdWithOptions(request, headers, runtime);
    }

    public GetConversationIdResponse getConversationIdWithOptions(GetConversationIdRequest request, GetConversationIdHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.userId)) {
            body.put("userId", request.userId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.appUid)) {
            body.put("appUid", request.appUid);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("GetConversationId", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/conversations", "json", req, runtime), new GetConversationIdResponse());
    }

    public RecallMessageResponse recallMessage(RecallMessageRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        RecallMessageHeaders headers = new RecallMessageHeaders();
        return this.recallMessageWithOptions(request, headers, runtime);
    }

    public RecallMessageResponse recallMessageWithOptions(RecallMessageRequest request, RecallMessageHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.operatorUid)) {
            body.put("operatorUid", request.operatorUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.messageId)) {
            body.put("messageId", request.messageId);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("RecallMessage", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/messages/recall", "none", req, runtime), new RecallMessageResponse());
    }

    public UpdateGroupNameResponse updateGroupName(UpdateGroupNameRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        UpdateGroupNameHeaders headers = new UpdateGroupNameHeaders();
        return this.updateGroupNameWithOptions(request, headers, runtime);
    }

    public UpdateGroupNameResponse updateGroupNameWithOptions(UpdateGroupNameRequest request, UpdateGroupNameHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.operatorUid)) {
            body.put("operatorUid", request.operatorUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.conversationId)) {
            body.put("conversationId", request.conversationId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.name)) {
            body.put("name", request.name);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("UpdateGroupName", "impaas_1.0", "HTTP", "PUT", "AK", "/v1.0/impaas/interconnections/groups/names", "none", req, runtime), new UpdateGroupNameResponse());
    }

    public CreateGroupResponse createGroup(CreateGroupRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        CreateGroupHeaders headers = new CreateGroupHeaders();
        return this.createGroupWithOptions(request, headers, runtime);
    }

    public CreateGroupResponse createGroupWithOptions(CreateGroupRequest request, CreateGroupHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.uuid)) {
            body.put("uuid", request.uuid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.creatorUid)) {
            body.put("creatorUid", request.creatorUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.name)) {
            body.put("name", request.name);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.iconMediaId)) {
            body.put("iconMediaId", request.iconMediaId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.channel)) {
            body.put("channel", request.channel);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.properties)) {
            body.put("properties", request.properties);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.members)) {
            body.put("members", request.members);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("CreateGroup", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/groups", "json", req, runtime), new CreateGroupResponse());
    }

    public RemoveGroupMembersResponse removeGroupMembers(RemoveGroupMembersRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        RemoveGroupMembersHeaders headers = new RemoveGroupMembersHeaders();
        return this.removeGroupMembersWithOptions(request, headers, runtime);
    }

    public RemoveGroupMembersResponse removeGroupMembersWithOptions(RemoveGroupMembersRequest request, RemoveGroupMembersHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.operatorUid)) {
            body.put("operatorUid", request.operatorUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.conversationId)) {
            body.put("conversationId", request.conversationId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.memberUids)) {
            body.put("memberUids", request.memberUids);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("RemoveGroupMembers", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/groups/members/batchRemove", "none", req, runtime), new RemoveGroupMembersResponse());
    }

    public GetMediaUrlResponse getMediaUrl(GetMediaUrlRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        GetMediaUrlHeaders headers = new GetMediaUrlHeaders();
        return this.getMediaUrlWithOptions(request, headers, runtime);
    }

    public GetMediaUrlResponse getMediaUrlWithOptions(GetMediaUrlRequest request, GetMediaUrlHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.mediaId)) {
            body.put("mediaId", request.mediaId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.urlExpireTime)) {
            body.put("urlExpireTime", request.urlExpireTime);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("GetMediaUrl", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/medium/urls", "json", req, runtime), new GetMediaUrlResponse());
    }

    public ReadMessageResponse readMessage(ReadMessageRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        ReadMessageHeaders headers = new ReadMessageHeaders();
        return this.readMessageWithOptions(request, headers, runtime);
    }

    public ReadMessageResponse readMessageWithOptions(ReadMessageRequest request, ReadMessageHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.operatorUid)) {
            body.put("operatorUid", request.operatorUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.messageId)) {
            body.put("messageId", request.messageId);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("ReadMessage", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/messages/read", "none", req, runtime), new ReadMessageResponse());
    }

    public DismissGroupResponse dismissGroup(DismissGroupRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        DismissGroupHeaders headers = new DismissGroupHeaders();
        return this.dismissGroupWithOptions(request, headers, runtime);
    }

    public DismissGroupResponse dismissGroupWithOptions(DismissGroupRequest request, DismissGroupHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.operatorUid)) {
            body.put("operatorUid", request.operatorUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.conversationId)) {
            body.put("conversationId", request.conversationId);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("DismissGroup", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/groups/dismiss", "none", req, runtime), new DismissGroupResponse());
    }

    public AddProfileResponse addProfile(AddProfileRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        AddProfileHeaders headers = new AddProfileHeaders();
        return this.addProfileWithOptions(request, headers, runtime);
    }

    public AddProfileResponse addProfileWithOptions(AddProfileRequest request, AddProfileHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.nick)) {
            body.put("nick", request.nick);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.avatarMediaId)) {
            body.put("avatarMediaId", request.avatarMediaId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.appUid)) {
            body.put("appUid", request.appUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.mobileNumber)) {
            body.put("mobileNumber", request.mobileNumber);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("AddProfile", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/users/profiles", "none", req, runtime), new AddProfileResponse());
    }

    public AddGroupMembersResponse addGroupMembers(AddGroupMembersRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        AddGroupMembersHeaders headers = new AddGroupMembersHeaders();
        return this.addGroupMembersWithOptions(request, headers, runtime);
    }

    public AddGroupMembersResponse addGroupMembersWithOptions(AddGroupMembersRequest request, AddGroupMembersHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.operatorUid)) {
            body.put("operatorUid", request.operatorUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.conversationId)) {
            body.put("conversationId", request.conversationId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.members)) {
            body.put("members", request.members);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("AddGroupMembers", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/groups/members/batchAdd", "none", req, runtime), new AddGroupMembersResponse());
    }

    public BatchSendResponse batchSend(BatchSendRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        BatchSendHeaders headers = new BatchSendHeaders();
        return this.batchSendWithOptions(request, headers, runtime);
    }

    public BatchSendResponse batchSendWithOptions(BatchSendRequest request, BatchSendHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.userId)) {
            body.put("userId", request.userId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.appUids)) {
            body.put("appUids", request.appUids);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.content)) {
            body.put("content", request.content);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("BatchSend", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/messages/batchSend", "json", req, runtime), new BatchSendResponse());
    }

    public ListGroupStaffMembersResponse listGroupStaffMembers(ListGroupStaffMembersRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        ListGroupStaffMembersHeaders headers = new ListGroupStaffMembersHeaders();
        return this.listGroupStaffMembersWithOptions(request, headers, runtime);
    }

    public ListGroupStaffMembersResponse listGroupStaffMembersWithOptions(ListGroupStaffMembersRequest request, ListGroupStaffMembersHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.conversationId)) {
            body.put("conversationId", request.conversationId);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("ListGroupStaffMembers", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/groups/staffMemers/query", "json", req, runtime), new ListGroupStaffMembersResponse());
    }

    public SendMessageResponse sendMessage(SendMessageRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        SendMessageHeaders headers = new SendMessageHeaders();
        return this.sendMessageWithOptions(request, headers, runtime);
    }

    public SendMessageResponse sendMessageWithOptions(SendMessageRequest request, SendMessageHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> body = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.senderUid)) {
            body.put("senderUid", request.senderUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.receiverUid)) {
            body.put("receiverUid", request.receiverUid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.conversationId)) {
            body.put("conversationId", request.conversationId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.content)) {
            body.put("content", request.content);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.uuid)) {
            body.put("uuid", request.uuid);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.createTime)) {
            body.put("createTime", request.createTime);
        }

        java.util.Map<String, String> realHeaders = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(headers.commonHeaders)) {
            realHeaders = headers.commonHeaders;
        }

        if (!com.aliyun.teautil.Common.isUnset(headers.xAcsDingtalkAccessToken)) {
            realHeaders.put("x-acs-dingtalk-access-token", headers.xAcsDingtalkAccessToken);
        }

        OpenApiRequest req = OpenApiRequest.build(TeaConverter.buildMap(
            new TeaPair("headers", realHeaders),
            new TeaPair("body", com.aliyun.openapiutil.Client.parseToMap(body))
        ));
        return TeaModel.toModel(this.doROARequest("SendMessage", "impaas_1.0", "HTTP", "POST", "AK", "/v1.0/impaas/interconnections/messages/send", "json", req, runtime), new SendMessageResponse());
    }
}
