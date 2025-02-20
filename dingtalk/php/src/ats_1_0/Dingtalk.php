<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vats_1_0;

use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\AddApplicationRegFormTemplateHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\AddApplicationRegFormTemplateRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\AddApplicationRegFormTemplateResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\AddFileHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\AddFileRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\AddFileResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetApplicationRegFormByFlowIdHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetApplicationRegFormByFlowIdRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetApplicationRegFormByFlowIdResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetCandidateByPhoneNumberHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetCandidateByPhoneNumberRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetCandidateByPhoneNumberResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetFileUploadInfoHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetFileUploadInfoRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetFileUploadInfoResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetFlowIdByRelationEntityIdHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetFlowIdByRelationEntityIdRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetFlowIdByRelationEntityIdResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetJobAuthHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetJobAuthRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\GetJobAuthResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\QueryInterviewsHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\QueryInterviewsRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\QueryInterviewsResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\UpdateApplicationRegFormHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\UpdateApplicationRegFormRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\UpdateApplicationRegFormResponse;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\UpdateInterviewSignInInfoHeaders;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\UpdateInterviewSignInInfoRequest;
use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\UpdateInterviewSignInInfoResponse;
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
     * @param AddFileRequest $request
     *
     * @return AddFileResponse
     */
    public function addFile($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new AddFileHeaders([]);

        return $this->addFileWithOptions($request, $headers, $runtime);
    }

    /**
     * @param AddFileRequest $request
     * @param AddFileHeaders $headers
     * @param RuntimeOptions $runtime
     *
     * @return AddFileResponse
     */
    public function addFileWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
        }
        if (!Utils::isUnset($request->opUserId)) {
            @$query['opUserId'] = $request->opUserId;
        }
        $body = [];
        if (!Utils::isUnset($request->mediaId)) {
            @$body['mediaId'] = $request->mediaId;
        }
        if (!Utils::isUnset($request->fileName)) {
            @$body['fileName'] = $request->fileName;
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
            'query'   => OpenApiUtilClient::query($query),
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return AddFileResponse::fromMap($this->doROARequest('AddFile', 'ats_1.0', 'HTTP', 'POST', 'AK', '/v1.0/ats/files', 'json', $req, $runtime));
    }

    /**
     * @param GetFlowIdByRelationEntityIdRequest $request
     *
     * @return GetFlowIdByRelationEntityIdResponse
     */
    public function getFlowIdByRelationEntityId($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new GetFlowIdByRelationEntityIdHeaders([]);

        return $this->getFlowIdByRelationEntityIdWithOptions($request, $headers, $runtime);
    }

    /**
     * @param GetFlowIdByRelationEntityIdRequest $request
     * @param GetFlowIdByRelationEntityIdHeaders $headers
     * @param RuntimeOptions                     $runtime
     *
     * @return GetFlowIdByRelationEntityIdResponse
     */
    public function getFlowIdByRelationEntityIdWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
        }
        if (!Utils::isUnset($request->relationEntity)) {
            @$query['relationEntity'] = $request->relationEntity;
        }
        if (!Utils::isUnset($request->relationEntityId)) {
            @$query['relationEntityId'] = $request->relationEntityId;
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
            'query'   => OpenApiUtilClient::query($query),
        ]);

        return GetFlowIdByRelationEntityIdResponse::fromMap($this->doROARequest('GetFlowIdByRelationEntityId', 'ats_1.0', 'HTTP', 'GET', 'AK', '/v1.0/ats/flows/ids', 'json', $req, $runtime));
    }

    /**
     * @param AddApplicationRegFormTemplateRequest $request
     *
     * @return AddApplicationRegFormTemplateResponse
     */
    public function addApplicationRegFormTemplate($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new AddApplicationRegFormTemplateHeaders([]);

        return $this->addApplicationRegFormTemplateWithOptions($request, $headers, $runtime);
    }

    /**
     * @param AddApplicationRegFormTemplateRequest $request
     * @param AddApplicationRegFormTemplateHeaders $headers
     * @param RuntimeOptions                       $runtime
     *
     * @return AddApplicationRegFormTemplateResponse
     */
    public function addApplicationRegFormTemplateWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
        }
        if (!Utils::isUnset($request->opUserId)) {
            @$query['opUserId'] = $request->opUserId;
        }
        $body = [];
        if (!Utils::isUnset($request->name)) {
            @$body['name'] = $request->name;
        }
        if (!Utils::isUnset($request->content)) {
            @$body['content'] = $request->content;
        }
        if (!Utils::isUnset($request->outerId)) {
            @$body['outerId'] = $request->outerId;
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
            'query'   => OpenApiUtilClient::query($query),
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return AddApplicationRegFormTemplateResponse::fromMap($this->doROARequest('AddApplicationRegFormTemplate', 'ats_1.0', 'HTTP', 'POST', 'AK', '/v1.0/ats/flows/applicationRegForms/templates', 'json', $req, $runtime));
    }

    /**
     * @param QueryInterviewsRequest $request
     *
     * @return QueryInterviewsResponse
     */
    public function queryInterviews($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new QueryInterviewsHeaders([]);

        return $this->queryInterviewsWithOptions($request, $headers, $runtime);
    }

    /**
     * @param QueryInterviewsRequest $request
     * @param QueryInterviewsHeaders $headers
     * @param RuntimeOptions         $runtime
     *
     * @return QueryInterviewsResponse
     */
    public function queryInterviewsWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
        }
        if (!Utils::isUnset($request->nextToken)) {
            @$query['nextToken'] = $request->nextToken;
        }
        if (!Utils::isUnset($request->size)) {
            @$query['size'] = $request->size;
        }
        $body = [];
        if (!Utils::isUnset($request->startTimeBeginMillis)) {
            @$body['startTimeBeginMillis'] = $request->startTimeBeginMillis;
        }
        if (!Utils::isUnset($request->startTimeEndMillis)) {
            @$body['startTimeEndMillis'] = $request->startTimeEndMillis;
        }
        if (!Utils::isUnset($request->candidateId)) {
            @$body['candidateId'] = $request->candidateId;
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
            'query'   => OpenApiUtilClient::query($query),
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return QueryInterviewsResponse::fromMap($this->doROARequest('QueryInterviews', 'ats_1.0', 'HTTP', 'POST', 'AK', '/v1.0/ats/interviews/query', 'json', $req, $runtime));
    }

    /**
     * @param GetFileUploadInfoRequest $request
     *
     * @return GetFileUploadInfoResponse
     */
    public function getFileUploadInfo($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new GetFileUploadInfoHeaders([]);

        return $this->getFileUploadInfoWithOptions($request, $headers, $runtime);
    }

    /**
     * @param GetFileUploadInfoRequest $request
     * @param GetFileUploadInfoHeaders $headers
     * @param RuntimeOptions           $runtime
     *
     * @return GetFileUploadInfoResponse
     */
    public function getFileUploadInfoWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
        }
        if (!Utils::isUnset($request->fileName)) {
            @$query['fileName'] = $request->fileName;
        }
        if (!Utils::isUnset($request->fileSize)) {
            @$query['fileSize'] = $request->fileSize;
        }
        if (!Utils::isUnset($request->md5)) {
            @$query['md5'] = $request->md5;
        }
        if (!Utils::isUnset($request->opUserId)) {
            @$query['opUserId'] = $request->opUserId;
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
            'query'   => OpenApiUtilClient::query($query),
        ]);

        return GetFileUploadInfoResponse::fromMap($this->doROARequest('GetFileUploadInfo', 'ats_1.0', 'HTTP', 'GET', 'AK', '/v1.0/ats/files/uploadInfos', 'json', $req, $runtime));
    }

    /**
     * @param string                               $flowId
     * @param GetApplicationRegFormByFlowIdRequest $request
     *
     * @return GetApplicationRegFormByFlowIdResponse
     */
    public function getApplicationRegFormByFlowId($flowId, $request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new GetApplicationRegFormByFlowIdHeaders([]);

        return $this->getApplicationRegFormByFlowIdWithOptions($flowId, $request, $headers, $runtime);
    }

    /**
     * @param string                               $flowId
     * @param GetApplicationRegFormByFlowIdRequest $request
     * @param GetApplicationRegFormByFlowIdHeaders $headers
     * @param RuntimeOptions                       $runtime
     *
     * @return GetApplicationRegFormByFlowIdResponse
     */
    public function getApplicationRegFormByFlowIdWithOptions($flowId, $request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
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
            'query'   => OpenApiUtilClient::query($query),
        ]);

        return GetApplicationRegFormByFlowIdResponse::fromMap($this->doROARequest('GetApplicationRegFormByFlowId', 'ats_1.0', 'HTTP', 'GET', 'AK', '/v1.0/ats/flows/' . $flowId . '/applicationRegForms', 'json', $req, $runtime));
    }

    /**
     * @param string                          $flowId
     * @param UpdateApplicationRegFormRequest $request
     *
     * @return UpdateApplicationRegFormResponse
     */
    public function updateApplicationRegForm($flowId, $request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new UpdateApplicationRegFormHeaders([]);

        return $this->updateApplicationRegFormWithOptions($flowId, $request, $headers, $runtime);
    }

    /**
     * @param string                          $flowId
     * @param UpdateApplicationRegFormRequest $request
     * @param UpdateApplicationRegFormHeaders $headers
     * @param RuntimeOptions                  $runtime
     *
     * @return UpdateApplicationRegFormResponse
     */
    public function updateApplicationRegFormWithOptions($flowId, $request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
        }
        $body = [];
        if (!Utils::isUnset($request->content)) {
            @$body['content'] = $request->content;
        }
        if (!Utils::isUnset($request->dingPanFile)) {
            @$body['dingPanFile'] = $request->dingPanFile;
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
            'query'   => OpenApiUtilClient::query($query),
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return UpdateApplicationRegFormResponse::fromMap($this->doROARequest('UpdateApplicationRegForm', 'ats_1.0', 'HTTP', 'PUT', 'AK', '/v1.0/ats/flows/' . $flowId . '/applicationRegForms', 'json', $req, $runtime));
    }

    /**
     * @param GetCandidateByPhoneNumberRequest $request
     *
     * @return GetCandidateByPhoneNumberResponse
     */
    public function getCandidateByPhoneNumber($request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new GetCandidateByPhoneNumberHeaders([]);

        return $this->getCandidateByPhoneNumberWithOptions($request, $headers, $runtime);
    }

    /**
     * @param GetCandidateByPhoneNumberRequest $request
     * @param GetCandidateByPhoneNumberHeaders $headers
     * @param RuntimeOptions                   $runtime
     *
     * @return GetCandidateByPhoneNumberResponse
     */
    public function getCandidateByPhoneNumberWithOptions($request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
        }
        if (!Utils::isUnset($request->phoneNumber)) {
            @$query['phoneNumber'] = $request->phoneNumber;
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
            'query'   => OpenApiUtilClient::query($query),
        ]);

        return GetCandidateByPhoneNumberResponse::fromMap($this->doROARequest('GetCandidateByPhoneNumber', 'ats_1.0', 'HTTP', 'GET', 'AK', '/v1.0/ats/candidates', 'json', $req, $runtime));
    }

    /**
     * @param string                           $interviewId
     * @param UpdateInterviewSignInInfoRequest $request
     *
     * @return UpdateInterviewSignInInfoResponse
     */
    public function updateInterviewSignInInfo($interviewId, $request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new UpdateInterviewSignInInfoHeaders([]);

        return $this->updateInterviewSignInInfoWithOptions($interviewId, $request, $headers, $runtime);
    }

    /**
     * @param string                           $interviewId
     * @param UpdateInterviewSignInInfoRequest $request
     * @param UpdateInterviewSignInInfoHeaders $headers
     * @param RuntimeOptions                   $runtime
     *
     * @return UpdateInterviewSignInInfoResponse
     */
    public function updateInterviewSignInInfoWithOptions($interviewId, $request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->bizCode)) {
            @$query['bizCode'] = $request->bizCode;
        }
        $body = [];
        if (!Utils::isUnset($request->signInTimeMillis)) {
            @$body['signInTimeMillis'] = $request->signInTimeMillis;
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
            'query'   => OpenApiUtilClient::query($query),
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return UpdateInterviewSignInInfoResponse::fromMap($this->doROARequest('UpdateInterviewSignInInfo', 'ats_1.0', 'HTTP', 'PUT', 'AK', '/v1.0/ats/interviews/' . $interviewId . '/signInInfos', 'none', $req, $runtime));
    }

    /**
     * @param string            $jobId
     * @param GetJobAuthRequest $request
     *
     * @return GetJobAuthResponse
     */
    public function getJobAuth($jobId, $request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = new GetJobAuthHeaders([]);

        return $this->getJobAuthWithOptions($jobId, $request, $headers, $runtime);
    }

    /**
     * @param string            $jobId
     * @param GetJobAuthRequest $request
     * @param GetJobAuthHeaders $headers
     * @param RuntimeOptions    $runtime
     *
     * @return GetJobAuthResponse
     */
    public function getJobAuthWithOptions($jobId, $request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $query = [];
        if (!Utils::isUnset($request->opUserId)) {
            @$query['opUserId'] = $request->opUserId;
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
            'query'   => OpenApiUtilClient::query($query),
        ]);

        return GetJobAuthResponse::fromMap($this->doROARequest('GetJobAuth', 'ats_1.0', 'HTTP', 'GET', 'AK', '/v1.0/ats/auths/jobs/' . $jobId . '', 'json', $req, $runtime));
    }
}
