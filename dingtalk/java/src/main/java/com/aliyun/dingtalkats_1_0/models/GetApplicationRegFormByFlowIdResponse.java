// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkats_1_0.models;

import com.aliyun.tea.*;

public class GetApplicationRegFormByFlowIdResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public GetApplicationRegFormByFlowIdResponseBody body;

    public static GetApplicationRegFormByFlowIdResponse build(java.util.Map<String, ?> map) throws Exception {
        GetApplicationRegFormByFlowIdResponse self = new GetApplicationRegFormByFlowIdResponse();
        return TeaModel.build(map, self);
    }

    public GetApplicationRegFormByFlowIdResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetApplicationRegFormByFlowIdResponse setBody(GetApplicationRegFormByFlowIdResponseBody body) {
        this.body = body;
        return this;
    }
    public GetApplicationRegFormByFlowIdResponseBody getBody() {
        return this.body;
    }

}
