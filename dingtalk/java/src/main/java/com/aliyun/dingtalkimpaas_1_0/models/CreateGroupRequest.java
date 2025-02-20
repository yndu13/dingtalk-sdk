// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkimpaas_1_0.models;

import com.aliyun.tea.*;

public class CreateGroupRequest extends TeaModel {
    @NameInMap("uuid")
    public String uuid;

    @NameInMap("creatorUid")
    public String creatorUid;

    @NameInMap("name")
    public String name;

    @NameInMap("iconMediaId")
    public String iconMediaId;

    @NameInMap("channel")
    public String channel;

    @NameInMap("properties")
    public java.util.Map<String, String> properties;

    @NameInMap("members")
    public java.util.List<CreateGroupRequestMembers> members;

    public static CreateGroupRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateGroupRequest self = new CreateGroupRequest();
        return TeaModel.build(map, self);
    }

    public CreateGroupRequest setUuid(String uuid) {
        this.uuid = uuid;
        return this;
    }
    public String getUuid() {
        return this.uuid;
    }

    public CreateGroupRequest setCreatorUid(String creatorUid) {
        this.creatorUid = creatorUid;
        return this;
    }
    public String getCreatorUid() {
        return this.creatorUid;
    }

    public CreateGroupRequest setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public CreateGroupRequest setIconMediaId(String iconMediaId) {
        this.iconMediaId = iconMediaId;
        return this;
    }
    public String getIconMediaId() {
        return this.iconMediaId;
    }

    public CreateGroupRequest setChannel(String channel) {
        this.channel = channel;
        return this;
    }
    public String getChannel() {
        return this.channel;
    }

    public CreateGroupRequest setProperties(java.util.Map<String, String> properties) {
        this.properties = properties;
        return this;
    }
    public java.util.Map<String, String> getProperties() {
        return this.properties;
    }

    public CreateGroupRequest setMembers(java.util.List<CreateGroupRequestMembers> members) {
        this.members = members;
        return this;
    }
    public java.util.List<CreateGroupRequestMembers> getMembers() {
        return this.members;
    }

    public static class CreateGroupRequestMembers extends TeaModel {
        @NameInMap("uid")
        public String uid;

        @NameInMap("nick")
        public String nick;

        public static CreateGroupRequestMembers build(java.util.Map<String, ?> map) throws Exception {
            CreateGroupRequestMembers self = new CreateGroupRequestMembers();
            return TeaModel.build(map, self);
        }

        public CreateGroupRequestMembers setUid(String uid) {
            this.uid = uid;
            return this;
        }
        public String getUid() {
            return this.uid;
        }

        public CreateGroupRequestMembers setNick(String nick) {
            this.nick = nick;
            return this;
        }
        public String getNick() {
            return this.nick;
        }

    }

}
