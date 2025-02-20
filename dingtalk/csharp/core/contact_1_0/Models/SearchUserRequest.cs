// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.Dingtalkcontact_1_0.Models
{
    public class SearchUserRequest : TeaModel {
        [NameInMap("dingOrgId")]
        [Validation(Required=false)]
        public long? DingOrgId { get; set; }

        /// <summary>
        /// 用户名称、名称拼音或英文名称
        /// </summary>
        [NameInMap("queryWord")]
        [Validation(Required=false)]
        public string QueryWord { get; set; }

        /// <summary>
        /// 分页查询锚点
        /// </summary>
        [NameInMap("offset")]
        [Validation(Required=false)]
        public int? Offset { get; set; }

        /// <summary>
        /// 分页长度
        /// </summary>
        [NameInMap("size")]
        [Validation(Required=false)]
        public int? Size { get; set; }

    }

}
