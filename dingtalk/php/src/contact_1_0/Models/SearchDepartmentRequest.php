<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vcontact_1_0\Models;

use AlibabaCloud\Tea\Model;

class SearchDepartmentRequest extends Model
{
    /**
     * @var int
     */
    public $dingOrgId;

    /**
     * @description 部门名称或者部门名称拼音
     *
     * @var string
     */
    public $queryWord;

    /**
     * @description 分页查询锚点
     *
     * @var int
     */
    public $offset;

    /**
     * @description 分页长度
     *
     * @var int
     */
    public $size;
    protected $_name = [
        'dingOrgId' => 'dingOrgId',
        'queryWord' => 'queryWord',
        'offset'    => 'offset',
        'size'      => 'size',
    ];

    public function validate()
    {
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->dingOrgId) {
            $res['dingOrgId'] = $this->dingOrgId;
        }
        if (null !== $this->queryWord) {
            $res['queryWord'] = $this->queryWord;
        }
        if (null !== $this->offset) {
            $res['offset'] = $this->offset;
        }
        if (null !== $this->size) {
            $res['size'] = $this->size;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return SearchDepartmentRequest
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['dingOrgId'])) {
            $model->dingOrgId = $map['dingOrgId'];
        }
        if (isset($map['queryWord'])) {
            $model->queryWord = $map['queryWord'];
        }
        if (isset($map['offset'])) {
            $model->offset = $map['offset'];
        }
        if (isset($map['size'])) {
            $model->size = $map['size'];
        }

        return $model;
    }
}
