<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vedu_1_0\Models\InitCoursesOfClassRequest\sectionConfig\sectionModels;

use AlibabaCloud\Tea\Model;

class start extends Model
{
    /**
     * @description 分钟
     *
     * @var int
     */
    public $min;

    /**
     * @description 小时
     *
     * @var int
     */
    public $hour;
    protected $_name = [
        'min'  => 'min',
        'hour' => 'hour',
    ];

    public function validate()
    {
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->min) {
            $res['min'] = $this->min;
        }
        if (null !== $this->hour) {
            $res['hour'] = $this->hour;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return start
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['min'])) {
            $model->min = $map['min'];
        }
        if (isset($map['hour'])) {
            $model->hour = $map['hour'];
        }

        return $model;
    }
}
