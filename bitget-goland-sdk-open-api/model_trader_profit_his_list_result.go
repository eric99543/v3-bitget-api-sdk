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

// TraderProfitHisListResult struct for TraderProfitHisListResult
type TraderProfitHisListResult struct {
	NextFlag             *bool                   `json:"nextFlag,omitempty"`
	ResultList           []TraderProfitHisResult `json:"resultList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TraderProfitHisListResult TraderProfitHisListResult

// NewTraderProfitHisListResult instantiates a new TraderProfitHisListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraderProfitHisListResult() *TraderProfitHisListResult {
	this := TraderProfitHisListResult{}
	return &this
}

// NewTraderProfitHisListResultWithDefaults instantiates a new TraderProfitHisListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraderProfitHisListResultWithDefaults() *TraderProfitHisListResult {
	this := TraderProfitHisListResult{}
	return &this
}

// GetNextFlag returns the NextFlag field value if set, zero value otherwise.
func (o *TraderProfitHisListResult) GetNextFlag() bool {
	if o == nil || isNil(o.NextFlag) {
		var ret bool
		return ret
	}
	return *o.NextFlag
}

// GetNextFlagOk returns a tuple with the NextFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderProfitHisListResult) GetNextFlagOk() (*bool, bool) {
	if o == nil || isNil(o.NextFlag) {
		return nil, false
	}
	return o.NextFlag, true
}

// HasNextFlag returns a boolean if a field has been set.
func (o *TraderProfitHisListResult) HasNextFlag() bool {
	if o != nil && !isNil(o.NextFlag) {
		return true
	}

	return false
}

// SetNextFlag gets a reference to the given bool and assigns it to the NextFlag field.
func (o *TraderProfitHisListResult) SetNextFlag(v bool) {
	o.NextFlag = &v
}

// GetResultList returns the ResultList field value if set, zero value otherwise.
func (o *TraderProfitHisListResult) GetResultList() []TraderProfitHisResult {
	if o == nil || isNil(o.ResultList) {
		var ret []TraderProfitHisResult
		return ret
	}
	return o.ResultList
}

// GetResultListOk returns a tuple with the ResultList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraderProfitHisListResult) GetResultListOk() ([]TraderProfitHisResult, bool) {
	if o == nil || isNil(o.ResultList) {
		return nil, false
	}
	return o.ResultList, true
}

// HasResultList returns a boolean if a field has been set.
func (o *TraderProfitHisListResult) HasResultList() bool {
	if o != nil && !isNil(o.ResultList) {
		return true
	}

	return false
}

// SetResultList gets a reference to the given []TraderProfitHisResult and assigns it to the ResultList field.
func (o *TraderProfitHisListResult) SetResultList(v []TraderProfitHisResult) {
	o.ResultList = v
}

func (o TraderProfitHisListResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextFlag) {
		toSerialize["nextFlag"] = o.NextFlag
	}
	if !isNil(o.ResultList) {
		toSerialize["resultList"] = o.ResultList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TraderProfitHisListResult) UnmarshalJSON(bytes []byte) (err error) {
	varTraderProfitHisListResult := _TraderProfitHisListResult{}

	if err = json.Unmarshal(bytes, &varTraderProfitHisListResult); err == nil {
		*o = TraderProfitHisListResult(varTraderProfitHisListResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nextFlag")
		delete(additionalProperties, "resultList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTraderProfitHisListResult struct {
	value *TraderProfitHisListResult
	isSet bool
}

func (v NullableTraderProfitHisListResult) Get() *TraderProfitHisListResult {
	return v.value
}

func (v *NullableTraderProfitHisListResult) Set(val *TraderProfitHisListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTraderProfitHisListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTraderProfitHisListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraderProfitHisListResult(val *TraderProfitHisListResult) *NullableTraderProfitHisListResult {
	return &NullableTraderProfitHisListResult{value: val, isSet: true}
}

func (v NullableTraderProfitHisListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraderProfitHisListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
