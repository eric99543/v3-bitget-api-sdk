/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MarginLiquidationInfoResult struct for MarginLiquidationInfoResult
type MarginLiquidationInfoResult struct {
	MaxId                *string                 `json:"maxId,omitempty"`
	MinId                *string                 `json:"minId,omitempty"`
	ResultList           []MarginLiquidationInfo `json:"resultList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginLiquidationInfoResult MarginLiquidationInfoResult

// NewMarginLiquidationInfoResult instantiates a new MarginLiquidationInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginLiquidationInfoResult() *MarginLiquidationInfoResult {
	this := MarginLiquidationInfoResult{}
	return &this
}

// NewMarginLiquidationInfoResultWithDefaults instantiates a new MarginLiquidationInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginLiquidationInfoResultWithDefaults() *MarginLiquidationInfoResult {
	this := MarginLiquidationInfoResult{}
	return &this
}

// GetMaxId returns the MaxId field value if set, zero value otherwise.
func (o *MarginLiquidationInfoResult) GetMaxId() string {
	if o == nil || isNil(o.MaxId) {
		var ret string
		return ret
	}
	return *o.MaxId
}

// GetMaxIdOk returns a tuple with the MaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfoResult) GetMaxIdOk() (*string, bool) {
	if o == nil || isNil(o.MaxId) {
		return nil, false
	}
	return o.MaxId, true
}

// HasMaxId returns a boolean if a field has been set.
func (o *MarginLiquidationInfoResult) HasMaxId() bool {
	if o != nil && !isNil(o.MaxId) {
		return true
	}

	return false
}

// SetMaxId gets a reference to the given string and assigns it to the MaxId field.
func (o *MarginLiquidationInfoResult) SetMaxId(v string) {
	o.MaxId = &v
}

// GetMinId returns the MinId field value if set, zero value otherwise.
func (o *MarginLiquidationInfoResult) GetMinId() string {
	if o == nil || isNil(o.MinId) {
		var ret string
		return ret
	}
	return *o.MinId
}

// GetMinIdOk returns a tuple with the MinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfoResult) GetMinIdOk() (*string, bool) {
	if o == nil || isNil(o.MinId) {
		return nil, false
	}
	return o.MinId, true
}

// HasMinId returns a boolean if a field has been set.
func (o *MarginLiquidationInfoResult) HasMinId() bool {
	if o != nil && !isNil(o.MinId) {
		return true
	}

	return false
}

// SetMinId gets a reference to the given string and assigns it to the MinId field.
func (o *MarginLiquidationInfoResult) SetMinId(v string) {
	o.MinId = &v
}

// GetResultList returns the ResultList field value if set, zero value otherwise.
func (o *MarginLiquidationInfoResult) GetResultList() []MarginLiquidationInfo {
	if o == nil || isNil(o.ResultList) {
		var ret []MarginLiquidationInfo
		return ret
	}
	return o.ResultList
}

// GetResultListOk returns a tuple with the ResultList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginLiquidationInfoResult) GetResultListOk() ([]MarginLiquidationInfo, bool) {
	if o == nil || isNil(o.ResultList) {
		return nil, false
	}
	return o.ResultList, true
}

// HasResultList returns a boolean if a field has been set.
func (o *MarginLiquidationInfoResult) HasResultList() bool {
	if o != nil && !isNil(o.ResultList) {
		return true
	}

	return false
}

// SetResultList gets a reference to the given []MarginLiquidationInfo and assigns it to the ResultList field.
func (o *MarginLiquidationInfoResult) SetResultList(v []MarginLiquidationInfo) {
	o.ResultList = v
}

func (o MarginLiquidationInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxId) {
		toSerialize["maxId"] = o.MaxId
	}
	if !isNil(o.MinId) {
		toSerialize["minId"] = o.MinId
	}
	if !isNil(o.ResultList) {
		toSerialize["resultList"] = o.ResultList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginLiquidationInfoResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginLiquidationInfoResult := _MarginLiquidationInfoResult{}

	if err = json.Unmarshal(bytes, &varMarginLiquidationInfoResult); err == nil {
		*o = MarginLiquidationInfoResult(varMarginLiquidationInfoResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "maxId")
		delete(additionalProperties, "minId")
		delete(additionalProperties, "resultList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginLiquidationInfoResult struct {
	value *MarginLiquidationInfoResult
	isSet bool
}

func (v NullableMarginLiquidationInfoResult) Get() *MarginLiquidationInfoResult {
	return v.value
}

func (v *NullableMarginLiquidationInfoResult) Set(val *MarginLiquidationInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginLiquidationInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginLiquidationInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginLiquidationInfoResult(val *MarginLiquidationInfoResult) *NullableMarginLiquidationInfoResult {
	return &NullableMarginLiquidationInfoResult{value: val, isSet: true}
}

func (v NullableMarginLiquidationInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginLiquidationInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
