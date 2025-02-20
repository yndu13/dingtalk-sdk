<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\QueryInterviewsResponseBody;

use AlibabaCloud\SDK\Dingtalk\Vats_1_0\Models\QueryInterviewsResponseBody\list_\interviewers;
use AlibabaCloud\Tea\Model;

class list_ extends Model
{
    /**
     * @description 面试标识
     *
     * @var string
     */
    public $interviewId;

    /**
     * @description 职位标识
     *
     * @var string
     */
    public $jobId;

    /**
     * @description 面试开始时间（单位：毫秒）
     *
     * @var int
     */
    public $startTimeMillis;

    /**
     * @description 面试结束时间（单位：毫秒）
     *
     * @var int
     */
    public $endTimeMillis;

    /**
     * @description 面试创建人员工标识
     *
     * @var string
     */
    public $creatorUserId;

    /**
     * @description 面试官列表
     *
     * @var interviewers[]
     */
    public $interviewers;
    protected $_name = [
        'interviewId'     => 'interviewId',
        'jobId'           => 'jobId',
        'startTimeMillis' => 'startTimeMillis',
        'endTimeMillis'   => 'endTimeMillis',
        'creatorUserId'   => 'creatorUserId',
        'interviewers'    => 'interviewers',
    ];

    public function validate()
    {
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->interviewId) {
            $res['interviewId'] = $this->interviewId;
        }
        if (null !== $this->jobId) {
            $res['jobId'] = $this->jobId;
        }
        if (null !== $this->startTimeMillis) {
            $res['startTimeMillis'] = $this->startTimeMillis;
        }
        if (null !== $this->endTimeMillis) {
            $res['endTimeMillis'] = $this->endTimeMillis;
        }
        if (null !== $this->creatorUserId) {
            $res['creatorUserId'] = $this->creatorUserId;
        }
        if (null !== $this->interviewers) {
            $res['interviewers'] = [];
            if (null !== $this->interviewers && \is_array($this->interviewers)) {
                $n = 0;
                foreach ($this->interviewers as $item) {
                    $res['interviewers'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return list_
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['interviewId'])) {
            $model->interviewId = $map['interviewId'];
        }
        if (isset($map['jobId'])) {
            $model->jobId = $map['jobId'];
        }
        if (isset($map['startTimeMillis'])) {
            $model->startTimeMillis = $map['startTimeMillis'];
        }
        if (isset($map['endTimeMillis'])) {
            $model->endTimeMillis = $map['endTimeMillis'];
        }
        if (isset($map['creatorUserId'])) {
            $model->creatorUserId = $map['creatorUserId'];
        }
        if (isset($map['interviewers'])) {
            if (!empty($map['interviewers'])) {
                $model->interviewers = [];
                $n                   = 0;
                foreach ($map['interviewers'] as $item) {
                    $model->interviewers[$n++] = null !== $item ? interviewers::fromMap($item) : $item;
                }
            }
        }

        return $model;
    }
}
