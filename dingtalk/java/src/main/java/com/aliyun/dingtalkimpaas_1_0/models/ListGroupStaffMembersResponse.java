// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkimpaas_1_0.models;

import com.aliyun.tea.*;

public class ListGroupStaffMembersResponse extends TeaModel {
    @NameInMap("headers")
    @Validation(required = true)
    public java.util.Map<String, String> headers;

    @NameInMap("body")
    @Validation(required = true)
    public ListGroupStaffMembersResponseBody body;

    public static ListGroupStaffMembersResponse build(java.util.Map<String, ?> map) throws Exception {
        ListGroupStaffMembersResponse self = new ListGroupStaffMembersResponse();
        return TeaModel.build(map, self);
    }

    public ListGroupStaffMembersResponse setHeaders(java.util.Map<String, String> headers) {
        this.headers = headers;
        return this;
    }
    public java.util.Map<String, String> getHeaders() {
        return this.headers;
    }

    public ListGroupStaffMembersResponse setBody(ListGroupStaffMembersResponseBody body) {
        this.body = body;
        return this;
    }
    public ListGroupStaffMembersResponseBody getBody() {
        return this.body;
    }

}
