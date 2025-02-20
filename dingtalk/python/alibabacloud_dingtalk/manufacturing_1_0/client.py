# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from typing import Dict
from Tea.core import TeaCore

from alibabacloud_tea_openapi.client import Client as OpenApiClient
from alibabacloud_tea_openapi import models as open_api_models
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_dingtalk.manufacturing_1_0 import models as dingtalkmanufacturing__1__0_models
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

    def industrialize_manufacture_job_book(
        self,
        user_id: str,
        request: dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookRequest,
    ) -> dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookResponse:
        runtime = util_models.RuntimeOptions()
        headers = {}
        return self.industrialize_manufacture_job_book_with_options(user_id, request, headers, runtime)

    async def industrialize_manufacture_job_book_async(
        self,
        user_id: str,
        request: dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookRequest,
    ) -> dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookResponse:
        runtime = util_models.RuntimeOptions()
        headers = {}
        return await self.industrialize_manufacture_job_book_with_options_async(user_id, request, headers, runtime)

    def industrialize_manufacture_job_book_with_options(
        self,
        user_id: str,
        request: dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookRequest,
        headers: Dict[str, str],
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.scrapped_quantity):
            body['scrappedQuantity'] = request.scrapped_quantity
        if not UtilClient.is_unset(request.product_specification):
            body['productSpecification'] = request.product_specification
        if not UtilClient.is_unset(request.qualified_quantity):
            body['qualifiedQuantity'] = request.qualified_quantity
        if not UtilClient.is_unset(request.reworkable_quantity):
            body['reworkableQuantity'] = request.reworkable_quantity
        if not UtilClient.is_unset(request.user_name):
            body['userName'] = request.user_name
        if not UtilClient.is_unset(request.uuid):
            body['uuid'] = request.uuid
        if not UtilClient.is_unset(request.product_name):
            body['productName'] = request.product_name
        if not UtilClient.is_unset(request.product_en_name):
            body['productEnName'] = request.product_en_name
        if not UtilClient.is_unset(request.extend):
            body['extend'] = request.extend
        if not UtilClient.is_unset(request.product_code):
            body['productCode'] = request.product_code
        if not UtilClient.is_unset(request.process_name):
            body['processName'] = request.process_name
        if not UtilClient.is_unset(request.process_en_name):
            body['processEnName'] = request.process_en_name
        if not UtilClient.is_unset(request.mes_app_key):
            body['mesAppKey'] = request.mes_app_key
        if not UtilClient.is_unset(request.inst_no):
            body['instNo'] = request.inst_no
        if not UtilClient.is_unset(request.manufacture_date):
            body['manufactureDate'] = request.manufacture_date
        if not UtilClient.is_unset(request.ding_corp_id):
            body['dingCorpId'] = request.ding_corp_id
        req = open_api_models.OpenApiRequest(
            headers=headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookResponse(),
            self.do_roarequest('IndustrializeManufactureJobBook', 'manufacturing_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/manufacturing/users/{user_id}/jobs', 'json', req, runtime)
        )

    async def industrialize_manufacture_job_book_with_options_async(
        self,
        user_id: str,
        request: dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookRequest,
        headers: Dict[str, str],
        runtime: util_models.RuntimeOptions,
    ) -> dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookResponse:
        UtilClient.validate_model(request)
        body = {}
        if not UtilClient.is_unset(request.scrapped_quantity):
            body['scrappedQuantity'] = request.scrapped_quantity
        if not UtilClient.is_unset(request.product_specification):
            body['productSpecification'] = request.product_specification
        if not UtilClient.is_unset(request.qualified_quantity):
            body['qualifiedQuantity'] = request.qualified_quantity
        if not UtilClient.is_unset(request.reworkable_quantity):
            body['reworkableQuantity'] = request.reworkable_quantity
        if not UtilClient.is_unset(request.user_name):
            body['userName'] = request.user_name
        if not UtilClient.is_unset(request.uuid):
            body['uuid'] = request.uuid
        if not UtilClient.is_unset(request.product_name):
            body['productName'] = request.product_name
        if not UtilClient.is_unset(request.product_en_name):
            body['productEnName'] = request.product_en_name
        if not UtilClient.is_unset(request.extend):
            body['extend'] = request.extend
        if not UtilClient.is_unset(request.product_code):
            body['productCode'] = request.product_code
        if not UtilClient.is_unset(request.process_name):
            body['processName'] = request.process_name
        if not UtilClient.is_unset(request.process_en_name):
            body['processEnName'] = request.process_en_name
        if not UtilClient.is_unset(request.mes_app_key):
            body['mesAppKey'] = request.mes_app_key
        if not UtilClient.is_unset(request.inst_no):
            body['instNo'] = request.inst_no
        if not UtilClient.is_unset(request.manufacture_date):
            body['manufactureDate'] = request.manufacture_date
        if not UtilClient.is_unset(request.ding_corp_id):
            body['dingCorpId'] = request.ding_corp_id
        req = open_api_models.OpenApiRequest(
            headers=headers,
            body=OpenApiUtilClient.parse_to_map(body)
        )
        return TeaCore.from_map(
            dingtalkmanufacturing__1__0_models.IndustrializeManufactureJobBookResponse(),
            await self.do_roarequest_async('IndustrializeManufactureJobBook', 'manufacturing_1.0', 'HTTP', 'POST', 'AK', f'/v1.0/manufacturing/users/{user_id}/jobs', 'json', req, runtime)
        )
