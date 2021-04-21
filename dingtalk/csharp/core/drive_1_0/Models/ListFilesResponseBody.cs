// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.SDK.Dingtalkdrive_1_0.Models
{
    public class ListFilesResponseBody : TeaModel {
        /// <summary>
        /// 文件列表
        /// </summary>
        [NameInMap("files")]
        [Validation(Required=false)]
        public List<ListFilesResponseBodyFiles> Files { get; set; }
        public class ListFilesResponseBodyFiles : TeaModel {
            /// <summary>
            /// 空间id
            /// </summary>
            [NameInMap("spaceId")]
            [Validation(Required=false)]
            public string SpaceId { get; set; }

            /// <summary>
            /// 文件id
            /// </summary>
            [NameInMap("fileId")]
            [Validation(Required=false)]
            public string FileId { get; set; }

            /// <summary>
            /// 文件名称
            /// </summary>
            [NameInMap("fileName")]
            [Validation(Required=false)]
            public string FileName { get; set; }

            /// <summary>
            /// 文件路径
            /// </summary>
            [NameInMap("filePath")]
            [Validation(Required=false)]
            public string FilePath { get; set; }

            /// <summary>
            /// 文件类型
            /// </summary>
            [NameInMap("fileType")]
            [Validation(Required=false)]
            public string FileType { get; set; }

            /// <summary>
            /// 文件后缀
            /// </summary>
            [NameInMap("fileExtension")]
            [Validation(Required=false)]
            public string FileExtension { get; set; }

        }

        /// <summary>
        /// 分页加载锚点, nextToken不为空表示有更多数据
        /// </summary>
        [NameInMap("nextToken")]
        [Validation(Required=false)]
        public string NextToken { get; set; }

    }

}
