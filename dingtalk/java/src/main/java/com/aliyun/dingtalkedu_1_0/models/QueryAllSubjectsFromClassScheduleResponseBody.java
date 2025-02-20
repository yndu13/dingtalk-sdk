// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkedu_1_0.models;

import com.aliyun.tea.*;

public class QueryAllSubjectsFromClassScheduleResponseBody extends TeaModel {
    // 查询结果。
    @NameInMap("result")
    public java.util.List<QueryAllSubjectsFromClassScheduleResponseBodyResult> result;

    public static QueryAllSubjectsFromClassScheduleResponseBody build(java.util.Map<String, ?> map) throws Exception {
        QueryAllSubjectsFromClassScheduleResponseBody self = new QueryAllSubjectsFromClassScheduleResponseBody();
        return TeaModel.build(map, self);
    }

    public QueryAllSubjectsFromClassScheduleResponseBody setResult(java.util.List<QueryAllSubjectsFromClassScheduleResponseBodyResult> result) {
        this.result = result;
        return this;
    }
    public java.util.List<QueryAllSubjectsFromClassScheduleResponseBodyResult> getResult() {
        return this.result;
    }

    public static class QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList extends TeaModel {
        // 老师的userid。
        @NameInMap("userId")
        public String userId;

        // 老师名称。
        @NameInMap("name")
        public String name;

        // 老师的头像url
        @NameInMap("avator")
        public String avator;

        // 老师的uid。
        @NameInMap("uid")
        public Long uid;

        public static QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList build(java.util.Map<String, ?> map) throws Exception {
            QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList self = new QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList();
            return TeaModel.build(map, self);
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList setUserId(String userId) {
            this.userId = userId;
            return this;
        }
        public String getUserId() {
            return this.userId;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList setAvator(String avator) {
            this.avator = avator;
            return this;
        }
        public String getAvator() {
            return this.avator;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList setUid(Long uid) {
            this.uid = uid;
            return this;
        }
        public Long getUid() {
            return this.uid;
        }

    }

    public static class QueryAllSubjectsFromClassScheduleResponseBodyResultExt extends TeaModel {
        // 学科字体颜色
        @NameInMap("fontColor")
        public String fontColor;

        // 学科背景颜色
        @NameInMap("backgroundColor")
        public String backgroundColor;

        // 班级id。
        @NameInMap("classId")
        public Long classId;

        // 老师列表
        @NameInMap("teacherList")
        public java.util.List<QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList> teacherList;

        public static QueryAllSubjectsFromClassScheduleResponseBodyResultExt build(java.util.Map<String, ?> map) throws Exception {
            QueryAllSubjectsFromClassScheduleResponseBodyResultExt self = new QueryAllSubjectsFromClassScheduleResponseBodyResultExt();
            return TeaModel.build(map, self);
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResultExt setFontColor(String fontColor) {
            this.fontColor = fontColor;
            return this;
        }
        public String getFontColor() {
            return this.fontColor;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResultExt setBackgroundColor(String backgroundColor) {
            this.backgroundColor = backgroundColor;
            return this;
        }
        public String getBackgroundColor() {
            return this.backgroundColor;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResultExt setClassId(Long classId) {
            this.classId = classId;
            return this;
        }
        public Long getClassId() {
            return this.classId;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResultExt setTeacherList(java.util.List<QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList> teacherList) {
            this.teacherList = teacherList;
            return this;
        }
        public java.util.List<QueryAllSubjectsFromClassScheduleResponseBodyResultExtTeacherList> getTeacherList() {
            return this.teacherList;
        }

    }

    public static class QueryAllSubjectsFromClassScheduleResponseBodyResult extends TeaModel {
        // 学科名称。
        @NameInMap("subjectName")
        public String subjectName;

        // 学科code。
        @NameInMap("subjectCode")
        public String subjectCode;

        // 学段编码：  KINDERGARTEN：小学 PRIMARY_SCHOOL：小学 MODDLE_SCHOOL： 初中 HIGH_SCHOOL： 高中 UNIVERSITY：大学 NOT_SCHOOL：无学段
        @NameInMap("periodCode")
        public String periodCode;

        // creatorOrgId
        @NameInMap("creatorOrgId")
        public Long creatorOrgId;

        // 拓展信息
        @NameInMap("ext")
        public QueryAllSubjectsFromClassScheduleResponseBodyResultExt ext;

        public static QueryAllSubjectsFromClassScheduleResponseBodyResult build(java.util.Map<String, ?> map) throws Exception {
            QueryAllSubjectsFromClassScheduleResponseBodyResult self = new QueryAllSubjectsFromClassScheduleResponseBodyResult();
            return TeaModel.build(map, self);
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResult setSubjectName(String subjectName) {
            this.subjectName = subjectName;
            return this;
        }
        public String getSubjectName() {
            return this.subjectName;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResult setSubjectCode(String subjectCode) {
            this.subjectCode = subjectCode;
            return this;
        }
        public String getSubjectCode() {
            return this.subjectCode;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResult setPeriodCode(String periodCode) {
            this.periodCode = periodCode;
            return this;
        }
        public String getPeriodCode() {
            return this.periodCode;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResult setCreatorOrgId(Long creatorOrgId) {
            this.creatorOrgId = creatorOrgId;
            return this;
        }
        public Long getCreatorOrgId() {
            return this.creatorOrgId;
        }

        public QueryAllSubjectsFromClassScheduleResponseBodyResult setExt(QueryAllSubjectsFromClassScheduleResponseBodyResultExt ext) {
            this.ext = ext;
            return this;
        }
        public QueryAllSubjectsFromClassScheduleResponseBodyResultExt getExt() {
            return this.ext;
        }

    }

}
