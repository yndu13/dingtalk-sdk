// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkoauth2_1_0.models;

import com.aliyun.tea.*;

public class GetCorpAccessTokenRequest extends TeaModel {
    // 应用id
    @NameInMap("suiteKey")
    public String suiteKey;

    // 应用密码
    @NameInMap("suiteSecret")
    public String suiteSecret;

    // OAuth 2.0 临时授权码
    @NameInMap("authCorpId")
    public String authCorpId;

    // suiteTicket
    @NameInMap("suiteTicket")
    public String suiteTicket;

    public static GetCorpAccessTokenRequest build(java.util.Map<String, ?> map) throws Exception {
        GetCorpAccessTokenRequest self = new GetCorpAccessTokenRequest();
        return TeaModel.build(map, self);
    }

    public GetCorpAccessTokenRequest setSuiteKey(String suiteKey) {
        this.suiteKey = suiteKey;
        return this;
    }
    public String getSuiteKey() {
        return this.suiteKey;
    }

    public GetCorpAccessTokenRequest setSuiteSecret(String suiteSecret) {
        this.suiteSecret = suiteSecret;
        return this;
    }
    public String getSuiteSecret() {
        return this.suiteSecret;
    }

    public GetCorpAccessTokenRequest setAuthCorpId(String authCorpId) {
        this.authCorpId = authCorpId;
        return this;
    }
    public String getAuthCorpId() {
        return this.authCorpId;
    }

    public GetCorpAccessTokenRequest setSuiteTicket(String suiteTicket) {
        this.suiteTicket = suiteTicket;
        return this;
    }
    public String getSuiteTicket() {
        return this.suiteTicket;
    }

}
