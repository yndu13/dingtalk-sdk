<?php

// This file is auto-generated, don't edit it. Thanks.

namespace AlibabaCloud\SDK\Dingtalk\Vcrm_1_0\Models;

use AlibabaCloud\SDK\Dingtalk\Vcrm_1_0\Models\CreateCustomerResponseBody\contacts;
use AlibabaCloud\Tea\Model;

class CreateCustomerResponseBody extends Model
{
    /**
     * @description 客户实例id
     *
     * @var string
     */
    public $customerInstanceId;

    /**
     * @description 保存客户类型
     *
     * @var string
     */
    public $objectType;

    /**
     * @description 联系人保存结果
     *
     * @var contacts[]
     */
    public $contacts;
    protected $_name = [
        'customerInstanceId' => 'customerInstanceId',
        'objectType'         => 'objectType',
        'contacts'           => 'contacts',
    ];

    public function validate()
    {
    }

    public function toMap()
    {
        $res = [];
        if (null !== $this->customerInstanceId) {
            $res['customerInstanceId'] = $this->customerInstanceId;
        }
        if (null !== $this->objectType) {
            $res['objectType'] = $this->objectType;
        }
        if (null !== $this->contacts) {
            $res['contacts'] = [];
            if (null !== $this->contacts && \is_array($this->contacts)) {
                $n = 0;
                foreach ($this->contacts as $item) {
                    $res['contacts'][$n++] = null !== $item ? $item->toMap() : $item;
                }
            }
        }

        return $res;
    }

    /**
     * @param array $map
     *
     * @return CreateCustomerResponseBody
     */
    public static function fromMap($map = [])
    {
        $model = new self();
        if (isset($map['customerInstanceId'])) {
            $model->customerInstanceId = $map['customerInstanceId'];
        }
        if (isset($map['objectType'])) {
            $model->objectType = $map['objectType'];
        }
        if (isset($map['contacts'])) {
            if (!empty($map['contacts'])) {
                $model->contacts = [];
                $n               = 0;
                foreach ($map['contacts'] as $item) {
                    $model->contacts[$n++] = null !== $item ? contacts::fromMap($item) : $item;
                }
            }
        }

        return $model;
    }
}
