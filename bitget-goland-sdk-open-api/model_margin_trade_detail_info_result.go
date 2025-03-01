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

// MarginTradeDetailInfoResult struct for MarginTradeDetailInfoResult
type MarginTradeDetailInfoResult struct {
	Fills                []MarginTradeDetailInfo `json:"fills,omitempty"`
	MaxId                *string                 `json:"maxId,omitempty"`
	MinId                *string                 `json:"minId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginTradeDetailInfoResult MarginTradeDetailInfoResult

// NewMarginTradeDetailInfoResult instantiates a new MarginTradeDetailInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginTradeDetailInfoResult() *MarginTradeDetailInfoResult {
	this := MarginTradeDetailInfoResult{}
	return &this
}

// NewMarginTradeDetailInfoResultWithDefaults instantiates a new MarginTradeDetailInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginTradeDetailInfoResultWithDefaults() *MarginTradeDetailInfoResult {
	this := MarginTradeDetailInfoResult{}
	return &this
}

// GetFills returns the Fills field value if set, zero value otherwise.
func (o *MarginTradeDetailInfoResult) GetFills() []MarginTradeDetailInfo {
	if o == nil || isNil(o.Fills) {
		var ret []MarginTradeDetailInfo
		return ret
	}
	return o.Fills
}

// GetFillsOk returns a tuple with the Fills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfoResult) GetFillsOk() ([]MarginTradeDetailInfo, bool) {
	if o == nil || isNil(o.Fills) {
		return nil, false
	}
	return o.Fills, true
}

// HasFills returns a boolean if a field has been set.
func (o *MarginTradeDetailInfoResult) HasFills() bool {
	if o != nil && !isNil(o.Fills) {
		return true
	}

	return false
}

// SetFills gets a reference to the given []MarginTradeDetailInfo and assigns it to the Fills field.
func (o *MarginTradeDetailInfoResult) SetFills(v []MarginTradeDetailInfo) {
	o.Fills = v
}

// GetMaxId returns the MaxId field value if set, zero value otherwise.
func (o *MarginTradeDetailInfoResult) GetMaxId() string {
	if o == nil || isNil(o.MaxId) {
		var ret string
		return ret
	}
	return *o.MaxId
}

// GetMaxIdOk returns a tuple with the MaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfoResult) GetMaxIdOk() (*string, bool) {
	if o == nil || isNil(o.MaxId) {
		return nil, false
	}
	return o.MaxId, true
}

// HasMaxId returns a boolean if a field has been set.
func (o *MarginTradeDetailInfoResult) HasMaxId() bool {
	if o != nil && !isNil(o.MaxId) {
		return true
	}

	return false
}

// SetMaxId gets a reference to the given string and assigns it to the MaxId field.
func (o *MarginTradeDetailInfoResult) SetMaxId(v string) {
	o.MaxId = &v
}

// GetMinId returns the MinId field value if set, zero value otherwise.
func (o *MarginTradeDetailInfoResult) GetMinId() string {
	if o == nil || isNil(o.MinId) {
		var ret string
		return ret
	}
	return *o.MinId
}

// GetMinIdOk returns a tuple with the MinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfoResult) GetMinIdOk() (*string, bool) {
	if o == nil || isNil(o.MinId) {
		return nil, false
	}
	return o.MinId, true
}

// HasMinId returns a boolean if a field has been set.
func (o *MarginTradeDetailInfoResult) HasMinId() bool {
	if o != nil && !isNil(o.MinId) {
		return true
	}

	return false
}

// SetMinId gets a reference to the given string and assigns it to the MinId field.
func (o *MarginTradeDetailInfoResult) SetMinId(v string) {
	o.MinId = &v
}

func (o MarginTradeDetailInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Fills) {
		toSerialize["fills"] = o.Fills
	}
	if !isNil(o.MaxId) {
		toSerialize["maxId"] = o.MaxId
	}
	if !isNil(o.MinId) {
		toSerialize["minId"] = o.MinId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginTradeDetailInfoResult) UnmarshalJSON(bytes []byte) (err error) {
	varMarginTradeDetailInfoResult := _MarginTradeDetailInfoResult{}

	if err = json.Unmarshal(bytes, &varMarginTradeDetailInfoResult); err == nil {
		*o = MarginTradeDetailInfoResult(varMarginTradeDetailInfoResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fills")
		delete(additionalProperties, "maxId")
		delete(additionalProperties, "minId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginTradeDetailInfoResult struct {
	value *MarginTradeDetailInfoResult
	isSet bool
}

func (v NullableMarginTradeDetailInfoResult) Get() *MarginTradeDetailInfoResult {
	return v.value
}

func (v *NullableMarginTradeDetailInfoResult) Set(val *MarginTradeDetailInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginTradeDetailInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginTradeDetailInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginTradeDetailInfoResult(val *MarginTradeDetailInfoResult) *NullableMarginTradeDetailInfoResult {
	return &NullableMarginTradeDetailInfoResult{value: val, isSet: true}
}

func (v NullableMarginTradeDetailInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginTradeDetailInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
