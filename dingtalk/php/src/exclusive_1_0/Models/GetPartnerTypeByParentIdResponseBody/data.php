<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vexclusive_1_0\Models\GetPartnerTypeByParentIdResponseBody;

use AlibabaCloud\Tea\Model;

class data extends Model
{
    /**
     * @description 自标签id
     *
     * @var float
     */
    public $typeId;

    /**
     * @description 子标签名
     *
     * @var string
     */
    public $typeName;
    protected $_name = [
        'typeId'   => 'typeId',
        'typeName' => 'typeName',
    ];

    public function validate()
    {
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->typeId) {
            $res['typeId'] = $this->typeId;
        }
        if (null !== $this->typeName) {
            $res['typeName'] = $this->typeName;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return data
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['typeId'])) {
            $model->typeId = $map['typeId'];
        }
        if (isset($map['typeName'])) {
            $model->typeName = $map['typeName'];
        }

        return $model;
    }
}
