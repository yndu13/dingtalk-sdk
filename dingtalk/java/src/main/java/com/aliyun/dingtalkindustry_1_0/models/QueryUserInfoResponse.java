// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkindustry_1_0.models;

import com.aliyun.tea.*;

public class QueryUserInfoResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public QueryUserInfoResponseBody body;

    public static QueryUserInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        QueryUserInfoResponse self = new QueryUserInfoResponse();
        return TeaModel.build(map, self);
    }

    public QueryUserInfoResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public QueryUserInfoResponse setBody(QueryUserInfoResponseBody body) {
        this.body = body;
        return this;
    }
    public QueryUserInfoResponseBody getBody() {
        return this.body;
    }

}
