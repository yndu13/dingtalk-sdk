// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.Dingtalkworkbench_1_0.Models
{
    public class QueryComponentScopesResponseBody : TeaModel {
        /// <summary>
        /// scopes
        /// </summary>
        [NameInMap("userVisibleScopes")]
        [Validation(Required=false)]
        public List<string> UserVisibleScopes { get; set; }

        [NameInMap("deptVisibleScopes")]
        [Validation(Required=false)]
        public List<string> DeptVisibleScopes { get; set; }

    }

}
