// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkservice_group_1_0.models;

import com.aliyun.tea.*;

public class QueryGroupResponseBody extends TeaModel {
    // 开放群ID
    @NameInMap("openConversationId")
    public String openConversationId;

    // 群名称
    @NameInMap("groupName")
    public String groupName;

    // 开放团队ID
    @NameInMap("openTeamId")
    public String openTeamId;

    // 开放群组ID
    @NameInMap("openGroupSetId")
    public String openGroupSetId;

    // 入群URL
    @NameInMap("groupUrl")
    public String groupUrl;

    // 服务群机器人code
    @NameInMap("robotCode")
    public String robotCode;

    // 服务群机器人名称
    @NameInMap("robotName")
    public String robotName;

    // 群bizId
    @NameInMap("bizId")
    public String bizId;

    public static QueryGroupResponseBody build(java.util.Map<String, ?> map) throws Exception {
        QueryGroupResponseBody self = new QueryGroupResponseBody();
        return TeaModel.build(map, self);
    }

    public QueryGroupResponseBody setOpenConversationId(String openConversationId) {
        this.openConversationId = openConversationId;
        return this;
    }
    public String getOpenConversationId() {
        return this.openConversationId;
    }

    public QueryGroupResponseBody setGroupName(String groupName) {
        this.groupName = groupName;
        return this;
    }
    public String getGroupName() {
        return this.groupName;
    }

    public QueryGroupResponseBody setOpenTeamId(String openTeamId) {
        this.openTeamId = openTeamId;
        return this;
    }
    public String getOpenTeamId() {
        return this.openTeamId;
    }

    public QueryGroupResponseBody setOpenGroupSetId(String openGroupSetId) {
        this.openGroupSetId = openGroupSetId;
        return this;
    }
    public String getOpenGroupSetId() {
        return this.openGroupSetId;
    }

    public QueryGroupResponseBody setGroupUrl(String groupUrl) {
        this.groupUrl = groupUrl;
        return this;
    }
    public String getGroupUrl() {
        return this.groupUrl;
    }

    public QueryGroupResponseBody setRobotCode(String robotCode) {
        this.robotCode = robotCode;
        return this;
    }
    public String getRobotCode() {
        return this.robotCode;
    }

    public QueryGroupResponseBody setRobotName(String robotName) {
        this.robotName = robotName;
        return this;
    }
    public String getRobotName() {
        return this.robotName;
    }

    public QueryGroupResponseBody setBizId(String bizId) {
        this.bizId = bizId;
        return this;
    }
    public String getBizId() {
        return this.bizId;
    }

}
