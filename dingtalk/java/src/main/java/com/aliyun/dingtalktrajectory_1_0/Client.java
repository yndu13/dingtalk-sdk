// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalktrajectory_1_0;

import com.aliyun.tea.*;
import com.aliyun.dingtalktrajectory_1_0.models.*;
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


    public QueryCollectingTraceTaskResponse queryCollectingTraceTask(QueryCollectingTraceTaskRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        QueryCollectingTraceTaskHeaders headers = new QueryCollectingTraceTaskHeaders();
        return this.queryCollectingTraceTaskWithOptions(request, headers, runtime);
    }

    public QueryCollectingTraceTaskResponse queryCollectingTraceTaskWithOptions(QueryCollectingTraceTaskRequest tmpReq, QueryCollectingTraceTaskHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(tmpReq);
        QueryCollectingTraceTaskShrinkRequest request = new QueryCollectingTraceTaskShrinkRequest();
        com.aliyun.openapiutil.Client.convert(tmpReq, request);
        if (!com.aliyun.teautil.Common.isUnset(tmpReq.staffIds)) {
            request.staffIdsShrink = com.aliyun.openapiutil.Client.arrayToStringWithSpecifiedStyle(tmpReq.staffIds, "staffIds", "json");
        }

        java.util.Map<String, Object> query = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.staffIdsShrink)) {
            query.put("staffIds", request.staffIdsShrink);
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
            new TeaPair("query", com.aliyun.openapiutil.Client.query(query))
        ));
        return TeaModel.toModel(this.doROARequest("QueryCollectingTraceTask", "trajectory_1.0", "HTTP", "GET", "AK", "/v1.0/trajectory/currentTasks", "json", req, runtime), new QueryCollectingTraceTaskResponse());
    }

    public QueryPageTraceDataResponse queryPageTraceData(QueryPageTraceDataRequest request) throws Exception {
        RuntimeOptions runtime = new RuntimeOptions();
        QueryPageTraceDataHeaders headers = new QueryPageTraceDataHeaders();
        return this.queryPageTraceDataWithOptions(request, headers, runtime);
    }

    public QueryPageTraceDataResponse queryPageTraceDataWithOptions(QueryPageTraceDataRequest request, QueryPageTraceDataHeaders headers, RuntimeOptions runtime) throws Exception {
        com.aliyun.teautil.Common.validateModel(request);
        java.util.Map<String, Object> query = new java.util.HashMap<>();
        if (!com.aliyun.teautil.Common.isUnset(request.traceId)) {
            query.put("traceId", request.traceId);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.nextToken)) {
            query.put("nextToken", request.nextToken);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.maxResults)) {
            query.put("maxResults", request.maxResults);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.startTime)) {
            query.put("startTime", request.startTime);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.endTime)) {
            query.put("endTime", request.endTime);
        }

        if (!com.aliyun.teautil.Common.isUnset(request.staffId)) {
            query.put("staffId", request.staffId);
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
            new TeaPair("query", com.aliyun.openapiutil.Client.query(query))
        ));
        return TeaModel.toModel(this.doROARequest("QueryPageTraceData", "trajectory_1.0", "HTTP", "GET", "AK", "/v1.0/trajectory/data", "json", req, runtime), new QueryPageTraceDataResponse());
    }
}
