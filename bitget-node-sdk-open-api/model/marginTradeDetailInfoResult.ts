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
import { MarginTradeDetailInfo } from './marginTradeDetailInfo';

export class MarginTradeDetailInfoResult {
    'fills'?: Array<MarginTradeDetailInfo>;
    'maxId'?: string;
    'minId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "fills",
            "baseName": "fills",
            "type": "Array<MarginTradeDetailInfo>"
        },
        {
            "name": "maxId",
            "baseName": "maxId",
            "type": "string"
        },
        {
            "name": "minId",
            "baseName": "minId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return MarginTradeDetailInfoResult.attributeTypeMap;
    }
}

