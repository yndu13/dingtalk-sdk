// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.dingtalkedu_1_0.models;

import com.aliyun.tea.*;

public class QuerySubjectTeachersResponseBody extends TeaModel {
    // 结果
    @NameInMap("result")
    public java.util.List<QuerySubjectTeachersResponseBodyResult> result;

    public static QuerySubjectTeachersResponseBody build(java.util.Map<String, ?> map) throws Exception {
        QuerySubjectTeachersResponseBody self = new QuerySubjectTeachersResponseBody();
        return TeaModel.build(map, self);
    }

    public QuerySubjectTeachersResponseBody setResult(java.util.List<QuerySubjectTeachersResponseBodyResult> result) {
        this.result = result;
        return this;
    }
    public java.util.List<QuerySubjectTeachersResponseBodyResult> getResult() {
        return this.result;
    }

    public static class QuerySubjectTeachersResponseBodyResult extends TeaModel {
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

        // 老师Uid
        @NameInMap("teacherUid")
        public Long teacherUid;

        // 班级id
        @NameInMap("classId")
        public Long classId;

        public static QuerySubjectTeachersResponseBodyResult build(java.util.Map<String, ?> map) throws Exception {
            QuerySubjectTeachersResponseBodyResult self = new QuerySubjectTeachersResponseBodyResult();
            return TeaModel.build(map, self);
        }

        public QuerySubjectTeachersResponseBodyResult setTeacherName(String teacherName) {
            this.teacherName = teacherName;
            return this;
        }
        public String getTeacherName() {
            return this.teacherName;
        }

        public QuerySubjectTeachersResponseBodyResult setSubjectName(String subjectName) {
            this.subjectName = subjectName;
            return this;
        }
        public String getSubjectName() {
            return this.subjectName;
        }

        public QuerySubjectTeachersResponseBodyResult setSubjectCode(String subjectCode) {
            this.subjectCode = subjectCode;
            return this;
        }
        public String getSubjectCode() {
            return this.subjectCode;
        }

        public QuerySubjectTeachersResponseBodyResult setPeriodCode(String periodCode) {
            this.periodCode = periodCode;
            return this;
        }
        public String getPeriodCode() {
            return this.periodCode;
        }

        public QuerySubjectTeachersResponseBodyResult setOrgId(Long orgId) {
            this.orgId = orgId;
            return this;
        }
        public Long getOrgId() {
            return this.orgId;
        }

        public QuerySubjectTeachersResponseBodyResult setTeacherUid(Long teacherUid) {
            this.teacherUid = teacherUid;
            return this;
        }
        public Long getTeacherUid() {
            return this.teacherUid;
        }

        public QuerySubjectTeachersResponseBodyResult setClassId(Long classId) {
            this.classId = classId;
            return this;
        }
        public Long getClassId() {
            return this.classId;
        }

    }

}
