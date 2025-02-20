// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkats_1_0.models;

import com.aliyun.tea.*;

public class GetFileUploadInfoResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public GetFileUploadInfoResponseBody body;

    public static GetFileUploadInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        GetFileUploadInfoResponse self = new GetFileUploadInfoResponse();
        return TeaModel.build(map, self);
    }

    public GetFileUploadInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetFileUploadInfoResponse setBody(GetFileUploadInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public GetFileUploadInfoResponseBody getBody() {
        return this.body;
    }

}
