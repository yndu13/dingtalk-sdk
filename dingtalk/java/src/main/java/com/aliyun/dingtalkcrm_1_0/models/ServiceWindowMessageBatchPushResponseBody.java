// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkcrm_1_0.models;

import com.aliyun.tea.*;

public class ServiceWindowMessageBatchPushResponseBody extends TeaModel {
    // result
    @NameInMap("result")
    public ServiceWindowMessageBatchPushResponseBodyResult result;

    @NameInMap("requestId")
    public String requestId;

    public static ServiceWindowMessageBatchPushResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ServiceWindowMessageBatchPushResponseBody self = new ServiceWindowMessageBatchPushResponseBody();
        return TeaModel.build(map, self);
    }

    public ServiceWindowMessageBatchPushResponseBody setResult(ServiceWindowMessageBatchPushResponseBodyResult result) {
        this.result = result;
        return this;
    }
    public ServiceWindowMessageBatchPushResponseBodyResult getResult() {
        return this.result;
    }

    public ServiceWindowMessageBatchPushResponseBody setRequestId(String requestId) {
        this.requestId = requestId;
        return this;
    }
    public String getRequestId() {
        return this.requestId;
    }

    public static class ServiceWindowMessageBatchPushResponseBodyResult extends TeaModel {
        @NameInMap("openPushId")
        public String openPushId;

        public static ServiceWindowMessageBatchPushResponseBodyResult build(java.util.Map<String, ?> map) throws Exception {
            ServiceWindowMessageBatchPushResponseBodyResult self = new ServiceWindowMessageBatchPushResponseBodyResult();
            return TeaModel.build(map, self);
        }

        public ServiceWindowMessageBatchPushResponseBodyResult setOpenPushId(String openPushId) {
            this.openPushId = openPushId;
            return this;
        }
        public String getOpenPushId() {
            return this.openPushId;
        }

    }

}
