// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkindustry_1_0.models;

import com.aliyun.tea.*;

public class UpdateUserExtendInfoRequest extends TeaModel {
    // 职称code
    @NameInMap("jobCode")
    public String jobCode;

    // 用户属性code
    @NameInMap("userProbCode")
    public String userProbCode;

    // 工作状态code
    @NameInMap("jobStatusCode")
    public java.util.List<String> jobStatusCode;

    // comments
    @NameInMap("comments")
    public String comments;

    public static UpdateUserExtendInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        UpdateUserExtendInfoRequest self = new UpdateUserExtendInfoRequest();
        return TeaModel.build(map, self);
    }

    public UpdateUserExtendInfoRequest setJobCode(String jobCode) {
        this.jobCode = jobCode;
        return this;
    }
    public String getJobCode() {
        return this.jobCode;
    }

    public UpdateUserExtendInfoRequest setUserProbCode(String userProbCode) {
        this.userProbCode = userProbCode;
        return this;
    }
    public String getUserProbCode() {
        return this.userProbCode;
    }

    public UpdateUserExtendInfoRequest setJobStatusCode(java.util.List<String> jobStatusCode) {
        this.jobStatusCode = jobStatusCode;
        return this;
    }
    public java.util.List<String> getJobStatusCode() {
        return this.jobStatusCode;
    }

    public UpdateUserExtendInfoRequest setComments(String comments) {
        this.comments = comments;
        return this;
    }
    public String getComments() {
        return this.comments;
    }

}
