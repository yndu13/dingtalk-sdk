// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkimpaas_1_0.models;

import com.aliyun.tea.*;

public class BatchSendResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public BatchSendResponseBody body;

    public static BatchSendResponse build(java.util.Map<String, ?> map) throws Exception {
        BatchSendResponse self = new BatchSendResponse();
        return TeaModel.build(map, self);
    }

    public BatchSendResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public BatchSendResponse setBody(BatchSendResponseBody body) {
        this.body = body;
        return this;
    }
    public BatchSendResponseBody getBody() {
        return this.body;
    }

}
