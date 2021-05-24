// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalktodo_1_0.models;

import com.aliyun.tea.*;

public class CreateTodoTaskRequest extends TeaModel {
    // 来源id，接入业务系统侧的唯一标识id
    @NameInMap("sourceId")
    public String sourceId;

    // 待办标题
    @NameInMap("subject")
    public String subject;

    // 创建者id，需传用户的unionId
    @NameInMap("creatorId")
    public String creatorId;

    // 待办备注描述
    @NameInMap("description")
    public String description;

    // 截止时间
    @NameInMap("dueTime")
    public Long dueTime;

    // 执行者列表，需传用户的unionId
    @NameInMap("executorIds")
    public java.util.List<String> executorIds;

    // 参与者列表，需传用户的unionId
    @NameInMap("participantIds")
    public java.util.List<String> participantIds;

    // 详情页url跳转地址
    @NameInMap("detailUrl")
    public CreateTodoTaskRequestDetailUrl detailUrl;

    // 待办卡片类型id
    @NameInMap("cardTypeId")
    public String cardTypeId;

    // 待办卡片内容区表单自定义字段列表
    @NameInMap("contentFieldList")
    public java.util.List<CreateTodoTaskRequestContentFieldList> contentFieldList;

    // 当前操作者id，需传用户的unionId
    @NameInMap("operatorId")
    public String operatorId;

    public static CreateTodoTaskRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateTodoTaskRequest self = new CreateTodoTaskRequest();
        return TeaModel.build(map, self);
    }

    public CreateTodoTaskRequest setSourceId(String sourceId) {
        this.sourceId = sourceId;
        return this;
    }
    public String getSourceId() {
        return this.sourceId;
    }

    public CreateTodoTaskRequest setSubject(String subject) {
        this.subject = subject;
        return this;
    }
    public String getSubject() {
        return this.subject;
    }

    public CreateTodoTaskRequest setCreatorId(String creatorId) {
        this.creatorId = creatorId;
        return this;
    }
    public String getCreatorId() {
        return this.creatorId;
    }

    public CreateTodoTaskRequest setDescription(String description) {
        this.description = description;
        return this;
    }
    public String getDescription() {
        return this.description;
    }

    public CreateTodoTaskRequest setDueTime(Long dueTime) {
        this.dueTime = dueTime;
        return this;
    }
    public Long getDueTime() {
        return this.dueTime;
    }

    public CreateTodoTaskRequest setExecutorIds(java.util.List<String> executorIds) {
        this.executorIds = executorIds;
        return this;
    }
    public java.util.List<String> getExecutorIds() {
        return this.executorIds;
    }

    public CreateTodoTaskRequest setParticipantIds(java.util.List<String> participantIds) {
        this.participantIds = participantIds;
        return this;
    }
    public java.util.List<String> getParticipantIds() {
        return this.participantIds;
    }

    public CreateTodoTaskRequest setDetailUrl(CreateTodoTaskRequestDetailUrl detailUrl) {
        this.detailUrl = detailUrl;
        return this;
    }
    public CreateTodoTaskRequestDetailUrl getDetailUrl() {
        return this.detailUrl;
    }

    public CreateTodoTaskRequest setCardTypeId(String cardTypeId) {
        this.cardTypeId = cardTypeId;
        return this;
    }
    public String getCardTypeId() {
        return this.cardTypeId;
    }

    public CreateTodoTaskRequest setContentFieldList(java.util.List<CreateTodoTaskRequestContentFieldList> contentFieldList) {
        this.contentFieldList = contentFieldList;
        return this;
    }
    public java.util.List<CreateTodoTaskRequestContentFieldList> getContentFieldList() {
        return this.contentFieldList;
    }

    public CreateTodoTaskRequest setOperatorId(String operatorId) {
        this.operatorId = operatorId;
        return this;
    }
    public String getOperatorId() {
        return this.operatorId;
    }

    public static class CreateTodoTaskRequestDetailUrl extends TeaModel {
        // app端详情页url
        @NameInMap("appUrl")
        public String appUrl;

        // pc端详情页url
        @NameInMap("pcUrl")
        public String pcUrl;

        public static CreateTodoTaskRequestDetailUrl build(java.util.Map<String, ?> map) throws Exception {
            CreateTodoTaskRequestDetailUrl self = new CreateTodoTaskRequestDetailUrl();
            return TeaModel.build(map, self);
        }

        public CreateTodoTaskRequestDetailUrl setAppUrl(String appUrl) {
            this.appUrl = appUrl;
            return this;
        }
        public String getAppUrl() {
            return this.appUrl;
        }

        public CreateTodoTaskRequestDetailUrl setPcUrl(String pcUrl) {
            this.pcUrl = pcUrl;
            return this;
        }
        public String getPcUrl() {
            return this.pcUrl;
        }

    }

    public static class CreateTodoTaskRequestContentFieldList extends TeaModel {
        // 字段唯一标识
        @NameInMap("fieldKey")
        public String fieldKey;

        // 字段值
        @NameInMap("fieldValue")
        public String fieldValue;

        // 字段内容链接
        @NameInMap("fieldLink")
        public String fieldLink;

        public static CreateTodoTaskRequestContentFieldList build(java.util.Map<String, ?> map) throws Exception {
            CreateTodoTaskRequestContentFieldList self = new CreateTodoTaskRequestContentFieldList();
            return TeaModel.build(map, self);
        }

        public CreateTodoTaskRequestContentFieldList setFieldKey(String fieldKey) {
            this.fieldKey = fieldKey;
            return this;
        }
        public String getFieldKey() {
            return this.fieldKey;
        }

        public CreateTodoTaskRequestContentFieldList setFieldValue(String fieldValue) {
            this.fieldValue = fieldValue;
            return this;
        }
        public String getFieldValue() {
            return this.fieldValue;
        }

        public CreateTodoTaskRequestContentFieldList setFieldLink(String fieldLink) {
            this.fieldLink = fieldLink;
            return this;
        }
        public String getFieldLink() {
            return this.fieldLink;
        }

    }

}
