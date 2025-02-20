// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkats_1_0.models;

import com.aliyun.tea.*;

public class GetCandidateByPhoneNumberResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public GetCandidateByPhoneNumberResponseBody body;

    public static GetCandidateByPhoneNumberResponse build(java.util.Map<String, ?> map) throws Exception {
        GetCandidateByPhoneNumberResponse self = new GetCandidateByPhoneNumberResponse();
        return TeaModel.build(map, self);
    }

    public GetCandidateByPhoneNumberResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public GetCandidateByPhoneNumberResponse setBody(GetCandidateByPhoneNumberResponseBody body) {
        this.body = body;
        return this;
    }
    public GetCandidateByPhoneNumberResponseBody getBody() {
        return this.body;
    }

}
