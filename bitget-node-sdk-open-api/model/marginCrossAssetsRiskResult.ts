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

export class MarginCrossAssetsRiskResult {
    'riskRate'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "riskRate",
            "baseName": "riskRate",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return MarginCrossAssetsRiskResult.attributeTypeMap;
    }
}

