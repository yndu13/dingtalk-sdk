<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0;

use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\AddGroupMembersHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\AddGroupMembersRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\AddGroupMembersResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\AddProfileHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\AddProfileRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\AddProfileResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\BatchSendHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\BatchSendRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\BatchSendResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\CreateGroupHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\CreateGroupRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\CreateGroupResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\DismissGroupHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\DismissGroupRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\DismissGroupResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\GetConversationIdHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\GetConversationIdRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\GetConversationIdResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\GetMediaUrlHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\GetMediaUrlRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\GetMediaUrlResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\ListGroupStaffMembersHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\ListGroupStaffMembersRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\ListGroupStaffMembersResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\ReadMessageHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\ReadMessageRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\ReadMessageResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\RecallMessageHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\RecallMessageRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\RecallMessageResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\RemoveGroupMembersHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\RemoveGroupMembersRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\RemoveGroupMembersResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\SendMessageHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\SendMessageRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\SendMessageResponse;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\UpdateGroupNameHeaders;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\UpdateGroupNameRequest;
use AlibabaCloud\SDK\Dingtalk\Vimpaas_1_0\Models\UpdateGroupNameResponse;
use AlibabaCloud\Tea\Utils\Utils;
use AlibabaCloud\Tea\Utils\Utils\RuntimeOptions;
use Darabonba\OpenApi\Models\OpenApiRequest;
use Darabonba\OpenApi\OpenApiClient;

class Dingtalk extends OpenApiClient
{
    public function __construct($config)
    {
        parent::__construct($config);
        $this->_endpointRule = '';
        if (Utils::empty_($this->_endpoint)) {
            $this->_endpoint = 'api.dingtalk.com';
        }
    }

    /**
     * @param GetConversationIdRequest $request
     *
     * @return GetConversationIdResponse
     */
    public function getConversationId($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new GetConversationIdHeaders([]);

        return $this->getConversationIdWithOptions($request, $headers, $runtime);
    }

    /**
     * @param GetConversationIdRequest $request
     * @param GetConversationIdHeaders $headers
     * @param RuntimeOptions           $runtime
     *
     * @return GetConversationIdResponse
     */
    public function getConversationIdWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->userId)) {
            @$body['userId'] = $request->userId;
        }
        if (!Utils::isUnset($request->appUid)) {
            @$body['appUid'] = $request->appUid;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return GetConversationIdResponse::fromMap($this->doROARequest('GetConversationId', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/conversations', 'json', $req, $runtime));
    }

    /**
     * @param RecallMessageRequest $request
     *
     * @return RecallMessageResponse
     */
    public function recallMessage($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new RecallMessageHeaders([]);

        return $this->recallMessageWithOptions($request, $headers, $runtime);
    }

    /**
     * @param RecallMessageRequest $request
     * @param RecallMessageHeaders $headers
     * @param RuntimeOptions       $runtime
     *
     * @return RecallMessageResponse
     */
    public function recallMessageWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->operatorUid)) {
            @$body['operatorUid'] = $request->operatorUid;
        }
        if (!Utils::isUnset($request->messageId)) {
            @$body['messageId'] = $request->messageId;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return RecallMessageResponse::fromMap($this->doROARequest('RecallMessage', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/messages/recall', 'none', $req, $runtime));
    }

    /**
     * @param UpdateGroupNameRequest $request
     *
     * @return UpdateGroupNameResponse
     */
    public function updateGroupName($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new UpdateGroupNameHeaders([]);

        return $this->updateGroupNameWithOptions($request, $headers, $runtime);
    }

    /**
     * @param UpdateGroupNameRequest $request
     * @param UpdateGroupNameHeaders $headers
     * @param RuntimeOptions         $runtime
     *
     * @return UpdateGroupNameResponse
     */
    public function updateGroupNameWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->operatorUid)) {
            @$body['operatorUid'] = $request->operatorUid;
        }
        if (!Utils::isUnset($request->conversationId)) {
            @$body['conversationId'] = $request->conversationId;
        }
        if (!Utils::isUnset($request->name)) {
            @$body['name'] = $request->name;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return UpdateGroupNameResponse::fromMap($this->doROARequest('UpdateGroupName', 'impaas_1.0', 'HTTP', 'PUT', 'AK', '/v1.0/impaas/interconnections/groups/names', 'none', $req, $runtime));
    }

    /**
     * @param CreateGroupRequest $request
     *
     * @return CreateGroupResponse
     */
    public function createGroup($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new CreateGroupHeaders([]);

        return $this->createGroupWithOptions($request, $headers, $runtime);
    }

    /**
     * @param CreateGroupRequest $request
     * @param CreateGroupHeaders $headers
     * @param RuntimeOptions     $runtime
     *
     * @return CreateGroupResponse
     */
    public function createGroupWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->uuid)) {
            @$body['uuid'] = $request->uuid;
        }
        if (!Utils::isUnset($request->creatorUid)) {
            @$body['creatorUid'] = $request->creatorUid;
        }
        if (!Utils::isUnset($request->name)) {
            @$body['name'] = $request->name;
        }
        if (!Utils::isUnset($request->iconMediaId)) {
            @$body['iconMediaId'] = $request->iconMediaId;
        }
        if (!Utils::isUnset($request->channel)) {
            @$body['channel'] = $request->channel;
        }
        if (!Utils::isUnset($request->properties)) {
            @$body['properties'] = $request->properties;
        }
        if (!Utils::isUnset($request->members)) {
            @$body['members'] = $request->members;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return CreateGroupResponse::fromMap($this->doROARequest('CreateGroup', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/groups', 'json', $req, $runtime));
    }

    /**
     * @param RemoveGroupMembersRequest $request
     *
     * @return RemoveGroupMembersResponse
     */
    public function removeGroupMembers($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new RemoveGroupMembersHeaders([]);

        return $this->removeGroupMembersWithOptions($request, $headers, $runtime);
    }

    /**
     * @param RemoveGroupMembersRequest $request
     * @param RemoveGroupMembersHeaders $headers
     * @param RuntimeOptions            $runtime
     *
     * @return RemoveGroupMembersResponse
     */
    public function removeGroupMembersWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->operatorUid)) {
            @$body['operatorUid'] = $request->operatorUid;
        }
        if (!Utils::isUnset($request->conversationId)) {
            @$body['conversationId'] = $request->conversationId;
        }
        if (!Utils::isUnset($request->memberUids)) {
            @$body['memberUids'] = $request->memberUids;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return RemoveGroupMembersResponse::fromMap($this->doROARequest('RemoveGroupMembers', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/groups/members/batchRemove', 'none', $req, $runtime));
    }

    /**
     * @param GetMediaUrlRequest $request
     *
     * @return GetMediaUrlResponse
     */
    public function getMediaUrl($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new GetMediaUrlHeaders([]);

        return $this->getMediaUrlWithOptions($request, $headers, $runtime);
    }

    /**
     * @param GetMediaUrlRequest $request
     * @param GetMediaUrlHeaders $headers
     * @param RuntimeOptions     $runtime
     *
     * @return GetMediaUrlResponse
     */
    public function getMediaUrlWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->mediaId)) {
            @$body['mediaId'] = $request->mediaId;
        }
        if (!Utils::isUnset($request->urlExpireTime)) {
            @$body['urlExpireTime'] = $request->urlExpireTime;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return GetMediaUrlResponse::fromMap($this->doROARequest('GetMediaUrl', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/medium/urls', 'json', $req, $runtime));
    }

    /**
     * @param ReadMessageRequest $request
     *
     * @return ReadMessageResponse
     */
    public function readMessage($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new ReadMessageHeaders([]);

        return $this->readMessageWithOptions($request, $headers, $runtime);
    }

    /**
     * @param ReadMessageRequest $request
     * @param ReadMessageHeaders $headers
     * @param RuntimeOptions     $runtime
     *
     * @return ReadMessageResponse
     */
    public function readMessageWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->operatorUid)) {
            @$body['operatorUid'] = $request->operatorUid;
        }
        if (!Utils::isUnset($request->messageId)) {
            @$body['messageId'] = $request->messageId;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return ReadMessageResponse::fromMap($this->doROARequest('ReadMessage', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/messages/read', 'none', $req, $runtime));
    }

    /**
     * @param DismissGroupRequest $request
     *
     * @return DismissGroupResponse
     */
    public function dismissGroup($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new DismissGroupHeaders([]);

        return $this->dismissGroupWithOptions($request, $headers, $runtime);
    }

    /**
     * @param DismissGroupRequest $request
     * @param DismissGroupHeaders $headers
     * @param RuntimeOptions      $runtime
     *
     * @return DismissGroupResponse
     */
    public function dismissGroupWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->operatorUid)) {
            @$body['operatorUid'] = $request->operatorUid;
        }
        if (!Utils::isUnset($request->conversationId)) {
            @$body['conversationId'] = $request->conversationId;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return DismissGroupResponse::fromMap($this->doROARequest('DismissGroup', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/groups/dismiss', 'none', $req, $runtime));
    }

    /**
     * @param AddProfileRequest $request
     *
     * @return AddProfileResponse
     */
    public function addProfile($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new AddProfileHeaders([]);

        return $this->addProfileWithOptions($request, $headers, $runtime);
    }

    /**
     * @param AddProfileRequest $request
     * @param AddProfileHeaders $headers
     * @param RuntimeOptions    $runtime
     *
     * @return AddProfileResponse
     */
    public function addProfileWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->nick)) {
            @$body['nick'] = $request->nick;
        }
        if (!Utils::isUnset($request->avatarMediaId)) {
            @$body['avatarMediaId'] = $request->avatarMediaId;
        }
        if (!Utils::isUnset($request->appUid)) {
            @$body['appUid'] = $request->appUid;
        }
        if (!Utils::isUnset($request->mobileNumber)) {
            @$body['mobileNumber'] = $request->mobileNumber;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return AddProfileResponse::fromMap($this->doROARequest('AddProfile', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/users/profiles', 'none', $req, $runtime));
    }

    /**
     * @param AddGroupMembersRequest $request
     *
     * @return AddGroupMembersResponse
     */
    public function addGroupMembers($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new AddGroupMembersHeaders([]);

        return $this->addGroupMembersWithOptions($request, $headers, $runtime);
    }

    /**
     * @param AddGroupMembersRequest $request
     * @param AddGroupMembersHeaders $headers
     * @param RuntimeOptions         $runtime
     *
     * @return AddGroupMembersResponse
     */
    public function addGroupMembersWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->operatorUid)) {
            @$body['operatorUid'] = $request->operatorUid;
        }
        if (!Utils::isUnset($request->conversationId)) {
            @$body['conversationId'] = $request->conversationId;
        }
        if (!Utils::isUnset($request->members)) {
            @$body['members'] = $request->members;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return AddGroupMembersResponse::fromMap($this->doROARequest('AddGroupMembers', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/groups/members/batchAdd', 'none', $req, $runtime));
    }

    /**
     * @param BatchSendRequest $request
     *
     * @return BatchSendResponse
     */
    public function batchSend($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new BatchSendHeaders([]);

        return $this->batchSendWithOptions($request, $headers, $runtime);
    }

    /**
     * @param BatchSendRequest $request
     * @param BatchSendHeaders $headers
     * @param RuntimeOptions   $runtime
     *
     * @return BatchSendResponse
     */
    public function batchSendWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->userId)) {
            @$body['userId'] = $request->userId;
        }
        if (!Utils::isUnset($request->appUids)) {
            @$body['appUids'] = $request->appUids;
        }
        if (!Utils::isUnset($request->content)) {
            @$body['content'] = $request->content;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return BatchSendResponse::fromMap($this->doROARequest('BatchSend', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/messages/batchSend', 'json', $req, $runtime));
    }

    /**
     * @param ListGroupStaffMembersRequest $request
     *
     * @return ListGroupStaffMembersResponse
     */
    public function listGroupStaffMembers($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new ListGroupStaffMembersHeaders([]);

        return $this->listGroupStaffMembersWithOptions($request, $headers, $runtime);
    }

    /**
     * @param ListGroupStaffMembersRequest $request
     * @param ListGroupStaffMembersHeaders $headers
     * @param RuntimeOptions               $runtime
     *
     * @return ListGroupStaffMembersResponse
     */
    public function listGroupStaffMembersWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->conversationId)) {
            @$body['conversationId'] = $request->conversationId;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return ListGroupStaffMembersResponse::fromMap($this->doROARequest('ListGroupStaffMembers', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/groups/staffMemers/query', 'json', $req, $runtime));
    }

    /**
     * @param SendMessageRequest $request
     *
     * @return SendMessageResponse
     */
    public function sendMessage($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new SendMessageHeaders([]);

        return $this->sendMessageWithOptions($request, $headers, $runtime);
    }

    /**
     * @param SendMessageRequest $request
     * @param SendMessageHeaders $headers
     * @param RuntimeOptions     $runtime
     *
     * @return SendMessageResponse
     */
    public function sendMessageWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->senderUid)) {
            @$body['senderUid'] = $request->senderUid;
        }
        if (!Utils::isUnset($request->receiverUid)) {
            @$body['receiverUid'] = $request->receiverUid;
        }
        if (!Utils::isUnset($request->conversationId)) {
            @$body['conversationId'] = $request->conversationId;
        }
        if (!Utils::isUnset($request->content)) {
            @$body['content'] = $request->content;
        }
        if (!Utils::isUnset($request->uuid)) {
            @$body['uuid'] = $request->uuid;
        }
        if (!Utils::isUnset($request->createTime)) {
            @$body['createTime'] = $request->createTime;
        }
        $realHeaders = [];
        if (!Utils::isUnset($headers->commonHeaders)) {
            $realHeaders = $headers->commonHeaders;
        }
        if (!Utils::isUnset($headers->xAcsDingtalkAccessToken)) {
            @$realHeaders['x-acs-dingtalk-access-token'] = $headers->xAcsDingtalkAccessToken;
        }
        $req = new OpenApiRequest([
            'headers' => $realHeaders,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return SendMessageResponse::fromMap($this->doROARequest('SendMessage', 'impaas_1.0', 'HTTP', 'POST', 'AK', '/v1.0/impaas/interconnections/messages/send', 'json', $req, $runtime));
    }
}
