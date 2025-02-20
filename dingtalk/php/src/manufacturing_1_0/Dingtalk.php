<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vmanufacturing_1_0;

use AlibabaCloud\OpenApiUtil\OpenApiUtilClient;
use AlibabaCloud\SDK\Dingtalk\Vmanufacturing_1_0\Models\IndustrializeManufactureJobBookRequest;
use AlibabaCloud\SDK\Dingtalk\Vmanufacturing_1_0\Models\IndustrializeManufactureJobBookResponse;
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
     * @param string                                 $userId
     * @param IndustrializeManufactureJobBookRequest $request
     *
     * @return IndustrializeManufactureJobBookResponse
     */
    public function industrializeManufactureJobBook($userId, $request)
    {
        $runtime = new RuntimeOptions([]);
        $headers = [];

        return $this->industrializeManufactureJobBookWithOptions($userId, $request, $headers, $runtime);
    }

    /**
     * @param string                                 $userId
     * @param IndustrializeManufactureJobBookRequest $request
     * @param string[]                               $headers
     * @param RuntimeOptions                         $runtime
     *
     * @return IndustrializeManufactureJobBookResponse
     */
    public function industrializeManufactureJobBookWithOptions($userId, $request, $headers, $runtime)
    {
        Utils::validateModel($request);
        $body = [];
        if (!Utils::isUnset($request->scrappedQuantity)) {
            @$body['scrappedQuantity'] = $request->scrappedQuantity;
        }
        if (!Utils::isUnset($request->productSpecification)) {
            @$body['productSpecification'] = $request->productSpecification;
        }
        if (!Utils::isUnset($request->qualifiedQuantity)) {
            @$body['qualifiedQuantity'] = $request->qualifiedQuantity;
        }
        if (!Utils::isUnset($request->reworkableQuantity)) {
            @$body['reworkableQuantity'] = $request->reworkableQuantity;
        }
        if (!Utils::isUnset($request->userName)) {
            @$body['userName'] = $request->userName;
        }
        if (!Utils::isUnset($request->uuid)) {
            @$body['uuid'] = $request->uuid;
        }
        if (!Utils::isUnset($request->productName)) {
            @$body['productName'] = $request->productName;
        }
        if (!Utils::isUnset($request->productEnName)) {
            @$body['productEnName'] = $request->productEnName;
        }
        if (!Utils::isUnset($request->extend)) {
            @$body['extend'] = $request->extend;
        }
        if (!Utils::isUnset($request->productCode)) {
            @$body['productCode'] = $request->productCode;
        }
        if (!Utils::isUnset($request->processName)) {
            @$body['processName'] = $request->processName;
        }
        if (!Utils::isUnset($request->processEnName)) {
            @$body['processEnName'] = $request->processEnName;
        }
        if (!Utils::isUnset($request->mesAppKey)) {
            @$body['mesAppKey'] = $request->mesAppKey;
        }
        if (!Utils::isUnset($request->instNo)) {
            @$body['instNo'] = $request->instNo;
        }
        if (!Utils::isUnset($request->manufactureDate)) {
            @$body['manufactureDate'] = $request->manufactureDate;
        }
        if (!Utils::isUnset($request->dingCorpId)) {
            @$body['dingCorpId'] = $request->dingCorpId;
        }
        $req = new OpenApiRequest([
            'headers' => $headers,
            'body'    => OpenApiUtilClient::parseToMap($body),
        ]);

        return IndustrializeManufactureJobBookResponse::fromMap($this->doROARequest('IndustrializeManufactureJobBook', 'manufacturing_1.0', 'HTTP', 'POST', 'AK', '/v1.0/manufacturing/users/' . $userId . '/jobs', 'json', $req, $runtime));
    }
}
