// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkesign_2_0.models;

import com.aliyun.tea.*;

public class GetAttachsApprovalResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public GetAttachsApprovalResponseBody body;

    public static GetAttachsApprovalResponse build(java.util.Map<String, ?> map) throws Exception {
        GetAttachsApprovalResponse self = new GetAttachsApprovalResponse();
        return TeaModel.build(map, self);
    }

    public GetAttachsApprovalResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetAttachsApprovalResponse setBody(GetAttachsApprovalResponseBody body) {
        this.body = body;
        return this;
    }
    public GetAttachsApprovalResponseBody getBody() {
        return this.body;
    }

}
