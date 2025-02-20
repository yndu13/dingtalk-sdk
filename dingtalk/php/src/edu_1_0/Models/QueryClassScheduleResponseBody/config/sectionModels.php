<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vedu_1_0\Models\QueryClassScheduleResponseBody\config;

use AlibabaCloud\SDK\Dingtalk\Vedu_1_0\Models\QueryClassScheduleResponseBody\config\sectionModels\end;
use AlibabaCloud\SDK\Dingtalk\Vedu_1_0\Models\QueryClassScheduleResponseBody\config\sectionModels\start;
use AlibabaCloud\Tea\Model;

class sectionModels extends Model
{
    /**
     * @description 节次名称。
     *
     * @var string
     */
    public $sectionName;

    /**
     * @description 节次类型：  COURSE：上课节次 REST：休息节次
     *
     * @var string
     */
    public $sectionType;

    /**
     * @description 节次序列。
     *
     * @var int
     */
    public $sectionIndex;

    /**
     * @description 开始时间（时分级别）。
     *
     * @var start
     */
    public $start;

    /**
     * @description 结束时间（时分级别）
     *
     * @var end
     */
    public $end;
    protected $_name = [
        'sectionName'  => 'sectionName',
        'sectionType'  => 'sectionType',
        'sectionIndex' => 'sectionIndex',
        'start'        => 'start',
        'end'          => 'end',
    ];

    public function validate()
    {
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->sectionName) {
            $res['sectionName'] = $this->sectionName;
        }
        if (null !== $this->sectionType) {
            $res['sectionType'] = $this->sectionType;
        }
        if (null !== $this->sectionIndex) {
            $res['sectionIndex'] = $this->sectionIndex;
        }
        if (null !== $this->start) {
            $res['start'] = null !== $this->start ? $this->start->toMap() : null;
        }
        if (null !== $this->end) {
            $res['end'] = null !== $this->end ? $this->end->toMap() : null;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return sectionModels
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['sectionName'])) {
            $model->sectionName = $map['sectionName'];
        }
        if (isset($map['sectionType'])) {
            $model->sectionType = $map['sectionType'];
        }
        if (isset($map['sectionIndex'])) {
            $model->sectionIndex = $map['sectionIndex'];
        }
        if (isset($map['start'])) {
            $model->start = start::fromMap($map['start']);
        }
        if (isset($map['end'])) {
            $model->end = end::fromMap($map['end']);
        }

        return $model;
    }
}
