// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkedu_1_0.models;

import com.aliyun.tea.*;

public class QueryTeachSubjectsResponseBody extends TeaModel {
    // 结果
    @NameInMap("result")
    public java.util.List<QueryTeachSubjectsResponseBodyResult> result;

    public static QueryTeachSubjectsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        QueryTeachSubjectsResponseBody self = new QueryTeachSubjectsResponseBody();
        return TeaModel.build(map, self);
    }

    public QueryTeachSubjectsResponseBody setResult(java.util.List<QueryTeachSubjectsResponseBodyResult> result) {
        this.result = result;
        return this;
    }
    public java.util.List<QueryTeachSubjectsResponseBodyResult> getResult() {
        return this.result;
    }

    public static class QueryTeachSubjectsResponseBodyResult extends TeaModel {
        // 老师名称
        @NameInMap("teacherName")
        public String teacherName;

        // 学科名称
        @NameInMap("subjectName")
        public String subjectName;

        // 学科code
        @NameInMap("subjectCode")
        public String subjectCode;

        // 学段code
        @NameInMap("periodCode")
        public String periodCode;

        // 组织id
        @NameInMap("orgId")
        public Long orgId;

        // 老师uid
        @NameInMap("teacherUid")
        public Long teacherUid;

        // 班级id
        @NameInMap("classId")
        public Long classId;

        public static QueryTeachSubjectsResponseBodyResult build(java.util.Map<String, ?> map) throws Exception {
            QueryTeachSubjectsResponseBodyResult self = new QueryTeachSubjectsResponseBodyResult();
            return TeaModel.build(map, self);
        }

        public QueryTeachSubjectsResponseBodyResult setTeacherName(String teacherName) {
            this.teacherName = teacherName;
            return this;
        }
        public String getTeacherName() {
            return this.teacherName;
        }

        public QueryTeachSubjectsResponseBodyResult setSubjectName(String subjectName) {
            this.subjectName = subjectName;
            return this;
        }
        public String getSubjectName() {
            return this.subjectName;
        }

        public QueryTeachSubjectsResponseBodyResult setSubjectCode(String subjectCode) {
            this.subjectCode = subjectCode;
            return this;
        }
        public String getSubjectCode() {
            return this.subjectCode;
        }

        public QueryTeachSubjectsResponseBodyResult setPeriodCode(String periodCode) {
            this.periodCode = periodCode;
            return this;
        }
        public String getPeriodCode() {
            return this.periodCode;
        }

        public QueryTeachSubjectsResponseBodyResult setOrgId(Long orgId) {
            this.orgId = orgId;
            return this;
        }
        public Long getOrgId() {
            return this.orgId;
        }

        public QueryTeachSubjectsResponseBodyResult setTeacherUid(Long teacherUid) {
            this.teacherUid = teacherUid;
            return this;
        }
        public Long getTeacherUid() {
            return this.teacherUid;
        }

        public QueryTeachSubjectsResponseBodyResult setClassId(Long classId) {
            this.classId = classId;
            return this;
        }
        public Long getClassId() {
            return this.classId;
        }

    }

}
