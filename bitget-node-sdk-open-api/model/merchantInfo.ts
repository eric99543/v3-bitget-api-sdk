/**
 * Bitget Open API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class MerchantInfo {
    'averagePayment'?: string;
    'averageRealese'?: string;
    'isOnline'?: string;
    'merchantId'?: string;
    'nickName'?: string;
    'registerTime'?: string;
    'thirtyBuy'?: string;
    'thirtyCompletionRate'?: string;
    'thirtySell'?: string;
    'thirtyTrades'?: string;
    'totalBuy'?: string;
    'totalCompletionRate'?: string;
    'totalSell'?: string;
    'totalTrades'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "averagePayment",
            "baseName": "averagePayment",
            "type": "string"
        },
        {
            "name": "averageRealese",
            "baseName": "averageRealese",
            "type": "string"
        },
        {
            "name": "isOnline",
            "baseName": "isOnline",
            "type": "string"
        },
        {
            "name": "merchantId",
            "baseName": "merchantId",
            "type": "string"
        },
        {
            "name": "nickName",
            "baseName": "nickName",
            "type": "string"
        },
        {
            "name": "registerTime",
            "baseName": "registerTime",
            "type": "string"
        },
        {
            "name": "thirtyBuy",
            "baseName": "thirtyBuy",
            "type": "string"
        },
        {
            "name": "thirtyCompletionRate",
            "baseName": "thirtyCompletionRate",
            "type": "string"
        },
        {
            "name": "thirtySell",
            "baseName": "thirtySell",
            "type": "string"
        },
        {
            "name": "thirtyTrades",
            "baseName": "thirtyTrades",
            "type": "string"
        },
        {
            "name": "totalBuy",
            "baseName": "totalBuy",
            "type": "string"
        },
        {
            "name": "totalCompletionRate",
            "baseName": "totalCompletionRate",
            "type": "string"
        },
        {
            "name": "totalSell",
            "baseName": "totalSell",
            "type": "string"
        },
        {
            "name": "totalTrades",
            "baseName": "totalTrades",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return MerchantInfo.attributeTypeMap;
    }
}

