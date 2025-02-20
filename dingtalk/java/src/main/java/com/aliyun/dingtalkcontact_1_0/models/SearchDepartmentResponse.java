// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkcontact_1_0.models;

import com.aliyun.tea.*;

public class SearchDepartmentResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public SearchDepartmentResponseBody body;

    public static SearchDepartmentResponse build(java.util.Map<String, ?> map) throws Exception {
        SearchDepartmentResponse self = new SearchDepartmentResponse();
        return TeaModel.build(map, self);
    }

    public SearchDepartmentResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public SearchDepartmentResponse setBody(SearchDepartmentResponseBody body) {
        this.body = body;
        return this;
    }
    public SearchDepartmentResponseBody getBody() {
        return this.body;
    }

}
