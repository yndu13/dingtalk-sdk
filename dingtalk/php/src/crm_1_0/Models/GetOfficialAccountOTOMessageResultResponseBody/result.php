<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vcrm_1_0\Models\GetOfficialAccountOTOMessageResultResponseBody;

use AlibabaCloud\Tea\Model;

class result extends Model
{
    /**
     * @description 执行状态： 0：未开始  1：处理中  2：处理完毕
     *
     * @var int
     */
    public $status;

    /**
     * @description 已读消息的userid列表
     *
     * @var string[]
     */
    public $readUserIdList;
    protected $_name = [
        'status'         => 'status',
        'readUserIdList' => 'readUserIdList',
    ];

    public function validate()
    {
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->status) {
            $res['status'] = $this->status;
        }
        if (null !== $this->readUserIdList) {
            $res['readUserIdList'] = $this->readUserIdList;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return result
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['status'])) {
            $model->status = $map['status'];
        }
        if (isset($map['readUserIdList'])) {
            if (!empty($map['readUserIdList'])) {
                $model->readUserIdList = $map['readUserIdList'];
            }
        }

        return $model;
    }
}
