// This file is auto-generated, don't edit it
/**
 *
 */
import Util, * as $Util from '@alicloud/tea-util';
import OpenApi, * as $OpenApi from '@alicloud/openapi-client';
import OpenApiUtil from '@alicloud/openapi-util';
import * as $tea from '@alicloud/tea-typescript';

export class CreateProcessHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateProcessRequest extends $tea.Model {
  dingCorpId?: string;
  initiatorUserId?: string;
  taskName?: string;
  signEndTime?: number;
  redirectUrl?: string;
  files?: CreateProcessRequestFiles[];
  participants?: CreateProcessRequestParticipants[];
  ccs?: CreateProcessRequestCcs[];
  sourceInfo?: CreateProcessRequestSourceInfo;
  static names(): { [key: string]: string } {
    return {
      dingCorpId: 'dingCorpId',
      initiatorUserId: 'initiatorUserId',
      taskName: 'taskName',
      signEndTime: 'signEndTime',
      redirectUrl: 'redirectUrl',
      files: 'files',
      participants: 'participants',
      ccs: 'ccs',
      sourceInfo: 'sourceInfo',
    };
  }

  static types(): { [key: string]: any } {
    return {
      dingCorpId: 'string',
      initiatorUserId: 'string',
      taskName: 'string',
      signEndTime: 'number',
      redirectUrl: 'string',
      files: { 'type': 'array', 'itemType': CreateProcessRequestFiles },
      participants: { 'type': 'array', 'itemType': CreateProcessRequestParticipants },
      ccs: { 'type': 'array', 'itemType': CreateProcessRequestCcs },
      sourceInfo: CreateProcessRequestSourceInfo,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateProcessResponseBody extends $tea.Model {
  taskId?: string;
  static names(): { [key: string]: string } {
    return {
      taskId: 'taskId',
    };
  }

  static types(): { [key: string]: any } {
    return {
      taskId: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateProcessResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: CreateProcessResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: CreateProcessResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetSignDetailHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetSignDetailResponseBody extends $tea.Model {
  businessScene?: string;
  flowStatus?: number;
  signers?: GetSignDetailResponseBodySigners[];
  static names(): { [key: string]: string } {
    return {
      businessScene: 'businessScene',
      flowStatus: 'flowStatus',
      signers: 'signers',
    };
  }

  static types(): { [key: string]: any } {
    return {
      businessScene: 'string',
      flowStatus: 'number',
      signers: { 'type': 'array', 'itemType': GetSignDetailResponseBodySigners },
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetSignDetailResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetSignDetailResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetSignDetailResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAttachsApprovalHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  tsignOpenAppId?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      tsignOpenAppId: 'tsignOpenAppId',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      tsignOpenAppId: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAttachsApprovalResponseBody extends $tea.Model {
  data?: GetAttachsApprovalResponseBodyData[];
  static names(): { [key: string]: string } {
    return {
      data: 'data',
    };
  }

  static types(): { [key: string]: any } {
    return {
      data: { 'type': 'array', 'itemType': GetAttachsApprovalResponseBodyData },
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAttachsApprovalResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetAttachsApprovalResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetAttachsApprovalResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ProcessStartHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ProcessStartRequest extends $tea.Model {
  autoStart?: string;
  initiatorUserId?: string;
  dingCorpId?: string;
  taskName?: string;
  redirectUrl?: string;
  files?: ProcessStartRequestFiles[];
  participants?: ProcessStartRequestParticipants[];
  ccs?: ProcessStartRequestCcs[];
  sourceInfo?: ProcessStartRequestSourceInfo;
  static names(): { [key: string]: string } {
    return {
      autoStart: 'autoStart',
      initiatorUserId: 'initiatorUserId',
      dingCorpId: 'dingCorpId',
      taskName: 'taskName',
      redirectUrl: 'redirectUrl',
      files: 'files',
      participants: 'participants',
      ccs: 'ccs',
      sourceInfo: 'sourceInfo',
    };
  }

  static types(): { [key: string]: any } {
    return {
      autoStart: 'string',
      initiatorUserId: 'string',
      dingCorpId: 'string',
      taskName: 'string',
      redirectUrl: 'string',
      files: { 'type': 'array', 'itemType': ProcessStartRequestFiles },
      participants: { 'type': 'array', 'itemType': ProcessStartRequestParticipants },
      ccs: { 'type': 'array', 'itemType': ProcessStartRequestCcs },
      sourceInfo: ProcessStartRequestSourceInfo,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ProcessStartResponseBody extends $tea.Model {
  taskId?: string;
  pcUrl?: string;
  mobileUrl?: string;
  static names(): { [key: string]: string } {
    return {
      taskId: 'taskId',
      pcUrl: 'pcUrl',
      mobileUrl: 'mobileUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      taskId: 'string',
      pcUrl: 'string',
      mobileUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ProcessStartResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: ProcessStartResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: ProcessStartResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ApprovalListHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ApprovalListResponseBody extends $tea.Model {
  data?: ApprovalListResponseBodyData[];
  static names(): { [key: string]: string } {
    return {
      data: 'data',
    };
  }

  static types(): { [key: string]: any } {
    return {
      data: { 'type': 'array', 'itemType': ApprovalListResponseBodyData },
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ApprovalListResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: ApprovalListResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: ApprovalListResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAuthUrlHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAuthUrlRequest extends $tea.Model {
  dingCorpId?: string;
  redirectUrl?: string;
  static names(): { [key: string]: string } {
    return {
      dingCorpId: 'dingCorpId',
      redirectUrl: 'redirectUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      dingCorpId: 'string',
      redirectUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAuthUrlResponseBody extends $tea.Model {
  taskId?: string;
  pcUrl?: string;
  mobileUrl?: string;
  static names(): { [key: string]: string } {
    return {
      taskId: 'taskId',
      pcUrl: 'pcUrl',
      mobileUrl: 'mobileUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      taskId: 'string',
      pcUrl: 'string',
      mobileUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAuthUrlResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetAuthUrlResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetAuthUrlResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetCorpConsoleHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetCorpConsoleResponseBody extends $tea.Model {
  orgConsoleUrl?: string;
  static names(): { [key: string]: string } {
    return {
      orgConsoleUrl: 'orgConsoleUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      orgConsoleUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetCorpConsoleResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetCorpConsoleResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetCorpConsoleResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFileInfoHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFileInfoResponseBody extends $tea.Model {
  fileId?: string;
  name?: string;
  downloadUrl?: string;
  size?: number;
  status?: number;
  pdfTotalPages?: number;
  static names(): { [key: string]: string } {
    return {
      fileId: 'fileId',
      name: 'name',
      downloadUrl: 'downloadUrl',
      size: 'size',
      status: 'status',
      pdfTotalPages: 'pdfTotalPages',
    };
  }

  static types(): { [key: string]: any } {
    return {
      fileId: 'string',
      name: 'string',
      downloadUrl: 'string',
      size: 'number',
      status: 'number',
      pdfTotalPages: 'number',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFileInfoResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetFileInfoResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetFileInfoResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ChannelOrdersHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ChannelOrdersRequest extends $tea.Model {
  dingCorpId?: string;
  orderId?: string;
  itemCode?: string;
  itemName?: string;
  quantity?: number;
  payFee?: number;
  orderCreateTime?: number;
  static names(): { [key: string]: string } {
    return {
      dingCorpId: 'dingCorpId',
      orderId: 'orderId',
      itemCode: 'itemCode',
      itemName: 'itemName',
      quantity: 'quantity',
      payFee: 'payFee',
      orderCreateTime: 'orderCreateTime',
    };
  }

  static types(): { [key: string]: any } {
    return {
      dingCorpId: 'string',
      orderId: 'string',
      itemCode: 'string',
      itemName: 'string',
      quantity: 'number',
      payFee: 'number',
      orderCreateTime: 'number',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ChannelOrdersResponseBody extends $tea.Model {
  esignOrderId?: string;
  static names(): { [key: string]: string } {
    return {
      esignOrderId: 'esignOrderId',
    };
  }

  static types(): { [key: string]: any } {
    return {
      esignOrderId: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ChannelOrdersResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: ChannelOrdersResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: ChannelOrdersResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ResaleOrderHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ResaleOrderRequest extends $tea.Model {
  dingCorpId?: string;
  orderId?: string;
  quantity?: number;
  orderCreateTime?: number;
  serviceStartTime?: number;
  serviceStopTime?: number;
  static names(): { [key: string]: string } {
    return {
      dingCorpId: 'dingCorpId',
      orderId: 'orderId',
      quantity: 'quantity',
      orderCreateTime: 'orderCreateTime',
      serviceStartTime: 'serviceStartTime',
      serviceStopTime: 'serviceStopTime',
    };
  }

  static types(): { [key: string]: any } {
    return {
      dingCorpId: 'string',
      orderId: 'string',
      quantity: 'number',
      orderCreateTime: 'number',
      serviceStartTime: 'number',
      serviceStopTime: 'number',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ResaleOrderResponseBody extends $tea.Model {
  esignOrderId?: string;
  static names(): { [key: string]: string } {
    return {
      esignOrderId: 'esignOrderId',
    };
  }

  static types(): { [key: string]: any } {
    return {
      esignOrderId: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ResaleOrderResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: ResaleOrderResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: ResaleOrderResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CancelCorpAuthHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CancelCorpAuthRequest extends $tea.Model {
  dingCorpId?: string;
  static names(): { [key: string]: string } {
    return {
      dingCorpId: 'dingCorpId',
    };
  }

  static types(): { [key: string]: any } {
    return {
      dingCorpId: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CancelCorpAuthResponseBody extends $tea.Model {
  result?: boolean;
  static names(): { [key: string]: string } {
    return {
      result: 'result',
    };
  }

  static types(): { [key: string]: any } {
    return {
      result: 'boolean',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CancelCorpAuthResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: CancelCorpAuthResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: CancelCorpAuthResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFileUploadUrlHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFileUploadUrlRequest extends $tea.Model {
  dingCorpId?: string;
  contentMd5?: string;
  contentType?: string;
  fileName?: string;
  fileSize?: number;
  convert2Pdf?: boolean;
  static names(): { [key: string]: string } {
    return {
      dingCorpId: 'dingCorpId',
      contentMd5: 'contentMd5',
      contentType: 'contentType',
      fileName: 'fileName',
      fileSize: 'fileSize',
      convert2Pdf: 'convert2Pdf',
    };
  }

  static types(): { [key: string]: any } {
    return {
      dingCorpId: 'string',
      contentMd5: 'string',
      contentType: 'string',
      fileName: 'string',
      fileSize: 'number',
      convert2Pdf: 'boolean',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFileUploadUrlResponseBody extends $tea.Model {
  fileId?: string;
  uploadUrl?: string;
  static names(): { [key: string]: string } {
    return {
      fileId: 'fileId',
      uploadUrl: 'uploadUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      fileId: 'string',
      uploadUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFileUploadUrlResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetFileUploadUrlResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetFileUploadUrlResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetIsvStatusHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetIsvStatusResponseBody extends $tea.Model {
  installStatus?: string;
  authStatus?: string;
  static names(): { [key: string]: string } {
    return {
      installStatus: 'installStatus',
      authStatus: 'authStatus',
    };
  }

  static types(): { [key: string]: any } {
    return {
      installStatus: 'string',
      authStatus: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetIsvStatusResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetIsvStatusResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetIsvStatusResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFlowDocsHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFlowDocsResponseBody extends $tea.Model {
  data?: GetFlowDocsResponseBodyData[];
  static names(): { [key: string]: string } {
    return {
      data: 'data',
    };
  }

  static types(): { [key: string]: any } {
    return {
      data: { 'type': 'array', 'itemType': GetFlowDocsResponseBodyData },
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFlowDocsResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetFlowDocsResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetFlowDocsResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class UsersRealnameHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class UsersRealnameRequest extends $tea.Model {
  userId?: string;
  redirectUrl?: string;
  dingCorpId?: string;
  static names(): { [key: string]: string } {
    return {
      userId: 'userId',
      redirectUrl: 'redirectUrl',
      dingCorpId: 'dingCorpId',
    };
  }

  static types(): { [key: string]: any } {
    return {
      userId: 'string',
      redirectUrl: 'string',
      dingCorpId: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class UsersRealnameResponseBody extends $tea.Model {
  taskId?: string;
  pcUrl?: string;
  mobileUrl?: string;
  static names(): { [key: string]: string } {
    return {
      taskId: 'taskId',
      pcUrl: 'pcUrl',
      mobileUrl: 'mobileUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      taskId: 'string',
      pcUrl: 'string',
      mobileUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class UsersRealnameResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: UsersRealnameResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: UsersRealnameResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFlowDetailHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFlowDetailResponseBody extends $tea.Model {
  businessScene?: string;
  flowStatus?: number;
  initiatorAuthorizedName?: string;
  initiatorName?: string;
  logs?: GetFlowDetailResponseBodyLogs[];
  static names(): { [key: string]: string } {
    return {
      businessScene: 'businessScene',
      flowStatus: 'flowStatus',
      initiatorAuthorizedName: 'initiatorAuthorizedName',
      initiatorName: 'initiatorName',
      logs: 'logs',
    };
  }

  static types(): { [key: string]: any } {
    return {
      businessScene: 'string',
      flowStatus: 'number',
      initiatorAuthorizedName: 'string',
      initiatorName: 'string',
      logs: { 'type': 'array', 'itemType': GetFlowDetailResponseBodyLogs },
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFlowDetailResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetFlowDetailResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetFlowDetailResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetCorpInfoHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetCorpInfoResponseBody extends $tea.Model {
  isRealName?: string;
  orgRealName?: string;
  static names(): { [key: string]: string } {
    return {
      isRealName: 'isRealName',
      orgRealName: 'orgRealName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      isRealName: 'string',
      orgRealName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetCorpInfoResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetCorpInfoResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetCorpInfoResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetUserInfoHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetUserInfoResponseBody extends $tea.Model {
  isRealName?: string;
  userRealName?: string;
  static names(): { [key: string]: string } {
    return {
      isRealName: 'isRealName',
      userRealName: 'userRealName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      isRealName: 'string',
      userRealName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetUserInfoResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetUserInfoResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetUserInfoResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetExecuteUrlHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetExecuteUrlRequest extends $tea.Model {
  taskId?: string;
  signContainer?: number;
  dingCorpId?: string;
  account?: string;
  static names(): { [key: string]: string } {
    return {
      taskId: 'taskId',
      signContainer: 'signContainer',
      dingCorpId: 'dingCorpId',
      account: 'account',
    };
  }

  static types(): { [key: string]: any } {
    return {
      taskId: 'string',
      signContainer: 'number',
      dingCorpId: 'string',
      account: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetExecuteUrlResponseBody extends $tea.Model {
  mobileUrl?: string;
  pcUrl?: string;
  longUrl?: string;
  shortUrl?: string;
  static names(): { [key: string]: string } {
    return {
      mobileUrl: 'mobileUrl',
      pcUrl: 'pcUrl',
      longUrl: 'longUrl',
      shortUrl: 'shortUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      mobileUrl: 'string',
      pcUrl: 'string',
      longUrl: 'string',
      shortUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetExecuteUrlResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetExecuteUrlResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetExecuteUrlResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetContractMarginHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetContractMarginResponseBody extends $tea.Model {
  margin?: number;
  static names(): { [key: string]: string } {
    return {
      margin: 'margin',
    };
  }

  static types(): { [key: string]: any } {
    return {
      margin: 'number',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetContractMarginResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: GetContractMarginResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: GetContractMarginResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateDevelopersHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateDevelopersRequest extends $tea.Model {
  dingCorpId?: string;
  noticeUrl?: string;
  static names(): { [key: string]: string } {
    return {
      dingCorpId: 'dingCorpId',
      noticeUrl: 'noticeUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      dingCorpId: 'string',
      noticeUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateDevelopersResponseBody extends $tea.Model {
  data?: boolean;
  static names(): { [key: string]: string } {
    return {
      data: 'data',
    };
  }

  static types(): { [key: string]: any } {
    return {
      data: 'boolean',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateDevelopersResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: CreateDevelopersResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: CreateDevelopersResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CorpRealnameHeaders extends $tea.Model {
  commonHeaders?: { [key: string]: string };
  serviceGroup?: string;
  xAcsDingtalkAccessToken?: string;
  static names(): { [key: string]: string } {
    return {
      commonHeaders: 'commonHeaders',
      serviceGroup: 'serviceGroup',
      xAcsDingtalkAccessToken: 'x-acs-dingtalk-access-token',
    };
  }

  static types(): { [key: string]: any } {
    return {
      commonHeaders: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      serviceGroup: 'string',
      xAcsDingtalkAccessToken: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CorpRealnameRequest extends $tea.Model {
  dingCorpId?: string;
  userId?: string;
  redirectUrl?: string;
  static names(): { [key: string]: string } {
    return {
      dingCorpId: 'dingCorpId',
      userId: 'userId',
      redirectUrl: 'redirectUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      dingCorpId: 'string',
      userId: 'string',
      redirectUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CorpRealnameResponseBody extends $tea.Model {
  taskId?: string;
  pcUrl?: string;
  mobileUrl?: string;
  static names(): { [key: string]: string } {
    return {
      taskId: 'taskId',
      pcUrl: 'pcUrl',
      mobileUrl: 'mobileUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      taskId: 'string',
      pcUrl: 'string',
      mobileUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CorpRealnameResponse extends $tea.Model {
  headers: { [key: string]: string };
  body: CorpRealnameResponseBody;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      body: 'body',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: CorpRealnameResponseBody,
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateProcessRequestFiles extends $tea.Model {
  fileId?: string;
  fileType?: number;
  fileName?: string;
  static names(): { [key: string]: string } {
    return {
      fileId: 'fileId',
      fileType: 'fileType',
      fileName: 'fileName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      fileId: 'string',
      fileType: 'number',
      fileName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateProcessRequestParticipants extends $tea.Model {
  signRequirements?: string;
  signOrder?: number;
  accountType?: string;
  account?: string;
  dingCorpId?: string;
  userId?: string;
  accountName?: string;
  orgName?: string;
  static names(): { [key: string]: string } {
    return {
      signRequirements: 'signRequirements',
      signOrder: 'signOrder',
      accountType: 'accountType',
      account: 'account',
      dingCorpId: 'dingCorpId',
      userId: 'userId',
      accountName: 'accountName',
      orgName: 'orgName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      signRequirements: 'string',
      signOrder: 'number',
      accountType: 'string',
      account: 'string',
      dingCorpId: 'string',
      userId: 'string',
      accountName: 'string',
      orgName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateProcessRequestCcs extends $tea.Model {
  accountType?: string;
  account?: string;
  dingCorpId?: string;
  userId?: string;
  accountName?: string;
  orgName?: string;
  static names(): { [key: string]: string } {
    return {
      accountType: 'accountType',
      account: 'account',
      dingCorpId: 'dingCorpId',
      userId: 'userId',
      accountName: 'accountName',
      orgName: 'orgName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      accountType: 'string',
      account: 'string',
      dingCorpId: 'string',
      userId: 'string',
      accountName: 'string',
      orgName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class CreateProcessRequestSourceInfo extends $tea.Model {
  showText?: string;
  pcUrl?: string;
  mobileUrl?: string;
  static names(): { [key: string]: string } {
    return {
      showText: 'showText',
      pcUrl: 'pcUrl',
      mobileUrl: 'mobileUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      showText: 'string',
      pcUrl: 'string',
      mobileUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetSignDetailResponseBodySigners extends $tea.Model {
  signStatus?: number;
  signerName?: string;
  static names(): { [key: string]: string } {
    return {
      signStatus: 'signStatus',
      signerName: 'signerName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      signStatus: 'number',
      signerName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAttachsApprovalResponseBodyDataFiles extends $tea.Model {
  fileName?: string;
  originalFileUrl?: string;
  signFinishFileUrl?: string;
  static names(): { [key: string]: string } {
    return {
      fileName: 'fileName',
      originalFileUrl: 'originalFileUrl',
      signFinishFileUrl: 'signFinishFileUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      fileName: 'string',
      originalFileUrl: 'string',
      signFinishFileUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetAttachsApprovalResponseBodyData extends $tea.Model {
  flowId?: string;
  status?: string;
  files?: GetAttachsApprovalResponseBodyDataFiles[];
  static names(): { [key: string]: string } {
    return {
      flowId: 'flowId',
      status: 'status',
      files: 'files',
    };
  }

  static types(): { [key: string]: any } {
    return {
      flowId: 'string',
      status: 'string',
      files: { 'type': 'array', 'itemType': GetAttachsApprovalResponseBodyDataFiles },
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ProcessStartRequestFiles extends $tea.Model {
  fileId?: string;
  fileName?: string;
  static names(): { [key: string]: string } {
    return {
      fileId: 'fileId',
      fileName: 'fileName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      fileId: 'string',
      fileName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ProcessStartRequestParticipants extends $tea.Model {
  accountType?: string;
  signRequirements?: string;
  dingCorpId?: string;
  userId?: string;
  account?: string;
  accountName?: string;
  orgName?: string;
  static names(): { [key: string]: string } {
    return {
      accountType: 'accountType',
      signRequirements: 'signRequirements',
      dingCorpId: 'dingCorpId',
      userId: 'userId',
      account: 'account',
      accountName: 'accountName',
      orgName: 'orgName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      accountType: 'string',
      signRequirements: 'string',
      dingCorpId: 'string',
      userId: 'string',
      account: 'string',
      accountName: 'string',
      orgName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ProcessStartRequestCcs extends $tea.Model {
  accountType?: string;
  dingCorpId?: string;
  userId?: string;
  account?: string;
  accountName?: string;
  orgName?: string;
  static names(): { [key: string]: string } {
    return {
      accountType: 'accountType',
      dingCorpId: 'dingCorpId',
      userId: 'userId',
      account: 'account',
      accountName: 'accountName',
      orgName: 'orgName',
    };
  }

  static types(): { [key: string]: any } {
    return {
      accountType: 'string',
      dingCorpId: 'string',
      userId: 'string',
      account: 'string',
      accountName: 'string',
      orgName: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ProcessStartRequestSourceInfo extends $tea.Model {
  showText?: string;
  pcUrl?: string;
  mobileUrl?: string;
  static names(): { [key: string]: string } {
    return {
      showText: 'showText',
      pcUrl: 'pcUrl',
      mobileUrl: 'mobileUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      showText: 'string',
      pcUrl: 'string',
      mobileUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ApprovalListResponseBodyDataApprovalNodes extends $tea.Model {
  approverName?: string;
  status?: string;
  startTime?: number;
  approvalTime?: string;
  static names(): { [key: string]: string } {
    return {
      approverName: 'approverName',
      status: 'status',
      startTime: 'startTime',
      approvalTime: 'approvalTime',
    };
  }

  static types(): { [key: string]: any } {
    return {
      approverName: 'string',
      status: 'string',
      startTime: 'number',
      approvalTime: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class ApprovalListResponseBodyData extends $tea.Model {
  approvalName?: string;
  status?: string;
  refuseReason?: string;
  sponsorAccountName?: string;
  startTime?: number;
  endTime?: number;
  sealIdImg?: string;
  approvalNodes?: ApprovalListResponseBodyDataApprovalNodes[];
  static names(): { [key: string]: string } {
    return {
      approvalName: 'approvalName',
      status: 'status',
      refuseReason: 'refuseReason',
      sponsorAccountName: 'sponsorAccountName',
      startTime: 'startTime',
      endTime: 'endTime',
      sealIdImg: 'sealIdImg',
      approvalNodes: 'approvalNodes',
    };
  }

  static types(): { [key: string]: any } {
    return {
      approvalName: 'string',
      status: 'string',
      refuseReason: 'string',
      sponsorAccountName: 'string',
      startTime: 'number',
      endTime: 'number',
      sealIdImg: 'string',
      approvalNodes: { 'type': 'array', 'itemType': ApprovalListResponseBodyDataApprovalNodes },
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFlowDocsResponseBodyData extends $tea.Model {
  fileId?: string;
  fileName?: string;
  fileUrl?: string;
  static names(): { [key: string]: string } {
    return {
      fileId: 'fileId',
      fileName: 'fileName',
      fileUrl: 'fileUrl',
    };
  }

  static types(): { [key: string]: any } {
    return {
      fileId: 'string',
      fileName: 'string',
      fileUrl: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class GetFlowDetailResponseBodyLogs extends $tea.Model {
  operatorAccountName?: string;
  logType?: string;
  operateDescription?: string;
  operateTime?: number;
  static names(): { [key: string]: string } {
    return {
      operatorAccountName: 'operatorAccountName',
      logType: 'logType',
      operateDescription: 'operateDescription',
      operateTime: 'operateTime',
    };
  }

  static types(): { [key: string]: any } {
    return {
      operatorAccountName: 'string',
      logType: 'string',
      operateDescription: 'string',
      operateTime: 'number',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}


export default class Client extends OpenApi {

  constructor(config: $OpenApi.Config) {
    super(config);
    this._endpointRule = "";
    if (Util.empty(this._endpoint)) {
      this._endpoint = "api.dingtalk.com";
    }

  }


  async createProcess(request: CreateProcessRequest): Promise<CreateProcessResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new CreateProcessHeaders({ });
    return await this.createProcessWithOptions(request, headers, runtime);
  }

  async createProcessWithOptions(request: CreateProcessRequest, headers: CreateProcessHeaders, runtime: $Util.RuntimeOptions): Promise<CreateProcessResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.initiatorUserId)) {
      body["initiatorUserId"] = request.initiatorUserId;
    }

    if (!Util.isUnset(request.taskName)) {
      body["taskName"] = request.taskName;
    }

    if (!Util.isUnset(request.signEndTime)) {
      body["signEndTime"] = request.signEndTime;
    }

    if (!Util.isUnset(request.redirectUrl)) {
      body["redirectUrl"] = request.redirectUrl;
    }

    if (!Util.isUnset(request.files)) {
      body["files"] = request.files;
    }

    if (!Util.isUnset(request.participants)) {
      body["participants"] = request.participants;
    }

    if (!Util.isUnset(request.ccs)) {
      body["ccs"] = request.ccs;
    }

    if (!Util.isUnset($tea.toMap(request.sourceInfo))) {
      body["sourceInfo"] = request.sourceInfo;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<CreateProcessResponse>(await this.doROARequest("CreateProcess", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/process/startAtOnce`, "json", req, runtime), new CreateProcessResponse({}));
  }

  async getSignDetail(taskId: string): Promise<GetSignDetailResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetSignDetailHeaders({ });
    return await this.getSignDetailWithOptions(taskId, headers, runtime);
  }

  async getSignDetailWithOptions(taskId: string, headers: GetSignDetailHeaders, runtime: $Util.RuntimeOptions): Promise<GetSignDetailResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetSignDetailResponse>(await this.doROARequest("GetSignDetail", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/signTasks/${taskId}`, "json", req, runtime), new GetSignDetailResponse({}));
  }

  async getAttachsApproval(instanceId: string): Promise<GetAttachsApprovalResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetAttachsApprovalHeaders({ });
    return await this.getAttachsApprovalWithOptions(instanceId, headers, runtime);
  }

  async getAttachsApprovalWithOptions(instanceId: string, headers: GetAttachsApprovalHeaders, runtime: $Util.RuntimeOptions): Promise<GetAttachsApprovalResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.tsignOpenAppId)) {
      realHeaders["tsignOpenAppId"] = headers.tsignOpenAppId;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetAttachsApprovalResponse>(await this.doROARequest("GetAttachsApproval", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/dingInstances/${instanceId}/attachments`, "json", req, runtime), new GetAttachsApprovalResponse({}));
  }

  async processStart(request: ProcessStartRequest): Promise<ProcessStartResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new ProcessStartHeaders({ });
    return await this.processStartWithOptions(request, headers, runtime);
  }

  async processStartWithOptions(request: ProcessStartRequest, headers: ProcessStartHeaders, runtime: $Util.RuntimeOptions): Promise<ProcessStartResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.autoStart)) {
      body["autoStart"] = request.autoStart;
    }

    if (!Util.isUnset(request.initiatorUserId)) {
      body["initiatorUserId"] = request.initiatorUserId;
    }

    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.taskName)) {
      body["taskName"] = request.taskName;
    }

    if (!Util.isUnset(request.redirectUrl)) {
      body["redirectUrl"] = request.redirectUrl;
    }

    if (!Util.isUnset(request.files)) {
      body["files"] = request.files;
    }

    if (!Util.isUnset(request.participants)) {
      body["participants"] = request.participants;
    }

    if (!Util.isUnset(request.ccs)) {
      body["ccs"] = request.ccs;
    }

    if (!Util.isUnset($tea.toMap(request.sourceInfo))) {
      body["sourceInfo"] = request.sourceInfo;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<ProcessStartResponse>(await this.doROARequest("ProcessStart", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/processes/startUrls`, "json", req, runtime), new ProcessStartResponse({}));
  }

  async approvalList(taskId: string): Promise<ApprovalListResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new ApprovalListHeaders({ });
    return await this.approvalListWithOptions(taskId, headers, runtime);
  }

  async approvalListWithOptions(taskId: string, headers: ApprovalListHeaders, runtime: $Util.RuntimeOptions): Promise<ApprovalListResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<ApprovalListResponse>(await this.doROARequest("ApprovalList", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/approvals/${taskId}`, "json", req, runtime), new ApprovalListResponse({}));
  }

  async getAuthUrl(request: GetAuthUrlRequest): Promise<GetAuthUrlResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetAuthUrlHeaders({ });
    return await this.getAuthUrlWithOptions(request, headers, runtime);
  }

  async getAuthUrlWithOptions(request: GetAuthUrlRequest, headers: GetAuthUrlHeaders, runtime: $Util.RuntimeOptions): Promise<GetAuthUrlResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.redirectUrl)) {
      body["redirectUrl"] = request.redirectUrl;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<GetAuthUrlResponse>(await this.doROARequest("GetAuthUrl", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/auths/urls`, "json", req, runtime), new GetAuthUrlResponse({}));
  }

  async getCorpConsole(): Promise<GetCorpConsoleResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetCorpConsoleHeaders({ });
    return await this.getCorpConsoleWithOptions(headers, runtime);
  }

  async getCorpConsoleWithOptions(headers: GetCorpConsoleHeaders, runtime: $Util.RuntimeOptions): Promise<GetCorpConsoleResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetCorpConsoleResponse>(await this.doROARequest("GetCorpConsole", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/corps/consoles`, "json", req, runtime), new GetCorpConsoleResponse({}));
  }

  async getFileInfo(fileId: string): Promise<GetFileInfoResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetFileInfoHeaders({ });
    return await this.getFileInfoWithOptions(fileId, headers, runtime);
  }

  async getFileInfoWithOptions(fileId: string, headers: GetFileInfoHeaders, runtime: $Util.RuntimeOptions): Promise<GetFileInfoResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetFileInfoResponse>(await this.doROARequest("GetFileInfo", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/files/${fileId}`, "json", req, runtime), new GetFileInfoResponse({}));
  }

  async channelOrders(request: ChannelOrdersRequest): Promise<ChannelOrdersResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new ChannelOrdersHeaders({ });
    return await this.channelOrdersWithOptions(request, headers, runtime);
  }

  async channelOrdersWithOptions(request: ChannelOrdersRequest, headers: ChannelOrdersHeaders, runtime: $Util.RuntimeOptions): Promise<ChannelOrdersResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.orderId)) {
      body["orderId"] = request.orderId;
    }

    if (!Util.isUnset(request.itemCode)) {
      body["itemCode"] = request.itemCode;
    }

    if (!Util.isUnset(request.itemName)) {
      body["itemName"] = request.itemName;
    }

    if (!Util.isUnset(request.quantity)) {
      body["quantity"] = request.quantity;
    }

    if (!Util.isUnset(request.payFee)) {
      body["payFee"] = request.payFee;
    }

    if (!Util.isUnset(request.orderCreateTime)) {
      body["orderCreateTime"] = request.orderCreateTime;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<ChannelOrdersResponse>(await this.doROARequest("ChannelOrders", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/orders/channel`, "json", req, runtime), new ChannelOrdersResponse({}));
  }

  async resaleOrder(request: ResaleOrderRequest): Promise<ResaleOrderResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new ResaleOrderHeaders({ });
    return await this.resaleOrderWithOptions(request, headers, runtime);
  }

  async resaleOrderWithOptions(request: ResaleOrderRequest, headers: ResaleOrderHeaders, runtime: $Util.RuntimeOptions): Promise<ResaleOrderResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.orderId)) {
      body["orderId"] = request.orderId;
    }

    if (!Util.isUnset(request.quantity)) {
      body["quantity"] = request.quantity;
    }

    if (!Util.isUnset(request.orderCreateTime)) {
      body["orderCreateTime"] = request.orderCreateTime;
    }

    if (!Util.isUnset(request.serviceStartTime)) {
      body["serviceStartTime"] = request.serviceStartTime;
    }

    if (!Util.isUnset(request.serviceStopTime)) {
      body["serviceStopTime"] = request.serviceStopTime;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<ResaleOrderResponse>(await this.doROARequest("ResaleOrder", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/orders/resale`, "json", req, runtime), new ResaleOrderResponse({}));
  }

  async cancelCorpAuth(request: CancelCorpAuthRequest): Promise<CancelCorpAuthResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new CancelCorpAuthHeaders({ });
    return await this.cancelCorpAuthWithOptions(request, headers, runtime);
  }

  async cancelCorpAuthWithOptions(request: CancelCorpAuthRequest, headers: CancelCorpAuthHeaders, runtime: $Util.RuntimeOptions): Promise<CancelCorpAuthResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<CancelCorpAuthResponse>(await this.doROARequest("CancelCorpAuth", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/auths/cancel`, "json", req, runtime), new CancelCorpAuthResponse({}));
  }

  async getFileUploadUrl(request: GetFileUploadUrlRequest): Promise<GetFileUploadUrlResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetFileUploadUrlHeaders({ });
    return await this.getFileUploadUrlWithOptions(request, headers, runtime);
  }

  async getFileUploadUrlWithOptions(request: GetFileUploadUrlRequest, headers: GetFileUploadUrlHeaders, runtime: $Util.RuntimeOptions): Promise<GetFileUploadUrlResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.contentMd5)) {
      body["contentMd5"] = request.contentMd5;
    }

    if (!Util.isUnset(request.contentType)) {
      body["contentType"] = request.contentType;
    }

    if (!Util.isUnset(request.fileName)) {
      body["fileName"] = request.fileName;
    }

    if (!Util.isUnset(request.fileSize)) {
      body["fileSize"] = request.fileSize;
    }

    if (!Util.isUnset(request.convert2Pdf)) {
      body["convert2Pdf"] = request.convert2Pdf;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<GetFileUploadUrlResponse>(await this.doROARequest("GetFileUploadUrl", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/files/uploadUrls`, "json", req, runtime), new GetFileUploadUrlResponse({}));
  }

  async getIsvStatus(): Promise<GetIsvStatusResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetIsvStatusHeaders({ });
    return await this.getIsvStatusWithOptions(headers, runtime);
  }

  async getIsvStatusWithOptions(headers: GetIsvStatusHeaders, runtime: $Util.RuntimeOptions): Promise<GetIsvStatusResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetIsvStatusResponse>(await this.doROARequest("GetIsvStatus", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/corps/appStatus`, "json", req, runtime), new GetIsvStatusResponse({}));
  }

  async getFlowDocs(taskId: string): Promise<GetFlowDocsResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetFlowDocsHeaders({ });
    return await this.getFlowDocsWithOptions(taskId, headers, runtime);
  }

  async getFlowDocsWithOptions(taskId: string, headers: GetFlowDocsHeaders, runtime: $Util.RuntimeOptions): Promise<GetFlowDocsResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetFlowDocsResponse>(await this.doROARequest("GetFlowDocs", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/flowTasks/${taskId}/docs`, "json", req, runtime), new GetFlowDocsResponse({}));
  }

  async usersRealname(request: UsersRealnameRequest): Promise<UsersRealnameResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new UsersRealnameHeaders({ });
    return await this.usersRealnameWithOptions(request, headers, runtime);
  }

  async usersRealnameWithOptions(request: UsersRealnameRequest, headers: UsersRealnameHeaders, runtime: $Util.RuntimeOptions): Promise<UsersRealnameResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.userId)) {
      body["userId"] = request.userId;
    }

    if (!Util.isUnset(request.redirectUrl)) {
      body["redirectUrl"] = request.redirectUrl;
    }

    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<UsersRealnameResponse>(await this.doROARequest("UsersRealname", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/users/realnames`, "json", req, runtime), new UsersRealnameResponse({}));
  }

  async getFlowDetail(taskId: string): Promise<GetFlowDetailResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetFlowDetailHeaders({ });
    return await this.getFlowDetailWithOptions(taskId, headers, runtime);
  }

  async getFlowDetailWithOptions(taskId: string, headers: GetFlowDetailHeaders, runtime: $Util.RuntimeOptions): Promise<GetFlowDetailResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetFlowDetailResponse>(await this.doROARequest("GetFlowDetail", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/flowTasks/${taskId}`, "json", req, runtime), new GetFlowDetailResponse({}));
  }

  async getCorpInfo(): Promise<GetCorpInfoResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetCorpInfoHeaders({ });
    return await this.getCorpInfoWithOptions(headers, runtime);
  }

  async getCorpInfoWithOptions(headers: GetCorpInfoHeaders, runtime: $Util.RuntimeOptions): Promise<GetCorpInfoResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetCorpInfoResponse>(await this.doROARequest("GetCorpInfo", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/corps/infos`, "json", req, runtime), new GetCorpInfoResponse({}));
  }

  async getUserInfo(userId: string): Promise<GetUserInfoResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetUserInfoHeaders({ });
    return await this.getUserInfoWithOptions(userId, headers, runtime);
  }

  async getUserInfoWithOptions(userId: string, headers: GetUserInfoHeaders, runtime: $Util.RuntimeOptions): Promise<GetUserInfoResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetUserInfoResponse>(await this.doROARequest("GetUserInfo", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/users/${userId}`, "json", req, runtime), new GetUserInfoResponse({}));
  }

  async getExecuteUrl(request: GetExecuteUrlRequest): Promise<GetExecuteUrlResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetExecuteUrlHeaders({ });
    return await this.getExecuteUrlWithOptions(request, headers, runtime);
  }

  async getExecuteUrlWithOptions(request: GetExecuteUrlRequest, headers: GetExecuteUrlHeaders, runtime: $Util.RuntimeOptions): Promise<GetExecuteUrlResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.taskId)) {
      body["taskId"] = request.taskId;
    }

    if (!Util.isUnset(request.signContainer)) {
      body["signContainer"] = request.signContainer;
    }

    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.account)) {
      body["account"] = request.account;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<GetExecuteUrlResponse>(await this.doROARequest("GetExecuteUrl", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/process/executeUrls`, "json", req, runtime), new GetExecuteUrlResponse({}));
  }

  async getContractMargin(): Promise<GetContractMarginResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new GetContractMarginHeaders({ });
    return await this.getContractMarginWithOptions(headers, runtime);
  }

  async getContractMarginWithOptions(headers: GetContractMarginHeaders, runtime: $Util.RuntimeOptions): Promise<GetContractMarginResponse> {
    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
    });
    return $tea.cast<GetContractMarginResponse>(await this.doROARequest("GetContractMargin", "esign_2.0", "HTTP", "GET", "AK", `/v2.0/esign/margins`, "json", req, runtime), new GetContractMarginResponse({}));
  }

  async createDevelopers(request: CreateDevelopersRequest): Promise<CreateDevelopersResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new CreateDevelopersHeaders({ });
    return await this.createDevelopersWithOptions(request, headers, runtime);
  }

  async createDevelopersWithOptions(request: CreateDevelopersRequest, headers: CreateDevelopersHeaders, runtime: $Util.RuntimeOptions): Promise<CreateDevelopersResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.noticeUrl)) {
      body["noticeUrl"] = request.noticeUrl;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<CreateDevelopersResponse>(await this.doROARequest("CreateDevelopers", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/developers`, "json", req, runtime), new CreateDevelopersResponse({}));
  }

  async corpRealname(request: CorpRealnameRequest): Promise<CorpRealnameResponse> {
    let runtime = new $Util.RuntimeOptions({ });
    let headers = new CorpRealnameHeaders({ });
    return await this.corpRealnameWithOptions(request, headers, runtime);
  }

  async corpRealnameWithOptions(request: CorpRealnameRequest, headers: CorpRealnameHeaders, runtime: $Util.RuntimeOptions): Promise<CorpRealnameResponse> {
    Util.validateModel(request);
    let body : {[key: string ]: any} = { };
    if (!Util.isUnset(request.dingCorpId)) {
      body["dingCorpId"] = request.dingCorpId;
    }

    if (!Util.isUnset(request.userId)) {
      body["userId"] = request.userId;
    }

    if (!Util.isUnset(request.redirectUrl)) {
      body["redirectUrl"] = request.redirectUrl;
    }

    let realHeaders : {[key: string ]: string} = { };
    if (!Util.isUnset(headers.commonHeaders)) {
      realHeaders = headers.commonHeaders;
    }

    if (!Util.isUnset(headers.serviceGroup)) {
      realHeaders["serviceGroup"] = headers.serviceGroup;
    }

    if (!Util.isUnset(headers.xAcsDingtalkAccessToken)) {
      realHeaders["x-acs-dingtalk-access-token"] = headers.xAcsDingtalkAccessToken;
    }

    let req = new $OpenApi.OpenApiRequest({
      headers: realHeaders,
      body: OpenApiUtil.parseToMap(body),
    });
    return $tea.cast<CorpRealnameResponse>(await this.doROARequest("CorpRealname", "esign_2.0", "HTTP", "POST", "AK", `/v2.0/esign/corps/realnames`, "json", req, runtime), new CorpRealnameResponse({}));
  }

}
