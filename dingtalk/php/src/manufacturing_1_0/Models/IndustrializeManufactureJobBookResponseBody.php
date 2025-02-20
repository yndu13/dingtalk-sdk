<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vmanufacturing_1_0\Models;

use AlibabaCloud\Tea\Model;

class IndustrializeManufactureJobBookResponseBody extends Model
{
    /**
     * @description httpCode
     *
     * @var string
     */
    public $httpCode;

    /**
     * @description 此次报工记录的唯一标识
     *
     * @var string
     */
    public $uuid;

    /**
     * @description content
     *
     * @var string
     */
    public $content;

    /**
     * @description errorMsg
     *
     * @var string
     */
    public $errorMsg;

    /**
     * @description errorLevel
     *
     * @var int
     */
    public $errorLevel;

    /**
     * @description errorCode
     *
     * @var string
     */
    public $errorCode;

    /**
     * @description success
     *
     * @var bool
     */
    public $success;
    protected $_name = [
        'httpCode'   => 'httpCode',
        'uuid'       => 'uuid',
        'content'    => 'content',
        'errorMsg'   => 'errorMsg',
        'errorLevel' => 'errorLevel',
        'errorCode'  => 'errorCode',
        'success'    => 'success',
    ];

    public function validate()
    {
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->httpCode) {
            $res['httpCode'] = $this->httpCode;
        }
        if (null !== $this->uuid) {
            $res['uuid'] = $this->uuid;
        }
        if (null !== $this->content) {
            $res['content'] = $this->content;
        }
        if (null !== $this->errorMsg) {
            $res['errorMsg'] = $this->errorMsg;
        }
        if (null !== $this->errorLevel) {
            $res['errorLevel'] = $this->errorLevel;
        }
        if (null !== $this->errorCode) {
            $res['errorCode'] = $this->errorCode;
        }
        if (null !== $this->success) {
            $res['success'] = $this->success;
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return IndustrializeManufactureJobBookResponseBody
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['httpCode'])) {
            $model->httpCode = $map['httpCode'];
        }
        if (isset($map['uuid'])) {
            $model->uuid = $map['uuid'];
        }
        if (isset($map['content'])) {
            $model->content = $map['content'];
        }
        if (isset($map['errorMsg'])) {
            $model->errorMsg = $map['errorMsg'];
        }
        if (isset($map['errorLevel'])) {
            $model->errorLevel = $map['errorLevel'];
        }
        if (isset($map['errorCode'])) {
            $model->errorCode = $map['errorCode'];
        }
        if (isset($map['success'])) {
            $model->success = $map['success'];
        }

        return $model;
    }
}
