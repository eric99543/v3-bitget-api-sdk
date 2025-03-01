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

// ApiResponseResultOfMarginCrossMaxBorrowResult struct for ApiResponseResultOfMarginCrossMaxBorrowResult
type ApiResponseResultOfMarginCrossMaxBorrowResult struct {
	// code
	Code *string                     `json:"code,omitempty"`
	Data *MarginCrossMaxBorrowResult `json:"data,omitempty"`
	// msg
	Msg *string `json:"msg,omitempty"`
	// requestTime
	RequestTime          *int64 `json:"requestTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiResponseResultOfMarginCrossMaxBorrowResult ApiResponseResultOfMarginCrossMaxBorrowResult

// NewApiResponseResultOfMarginCrossMaxBorrowResult instantiates a new ApiResponseResultOfMarginCrossMaxBorrowResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponseResultOfMarginCrossMaxBorrowResult() *ApiResponseResultOfMarginCrossMaxBorrowResult {
	this := ApiResponseResultOfMarginCrossMaxBorrowResult{}
	return &this
}

// NewApiResponseResultOfMarginCrossMaxBorrowResultWithDefaults instantiates a new ApiResponseResultOfMarginCrossMaxBorrowResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponseResultOfMarginCrossMaxBorrowResultWithDefaults() *ApiResponseResultOfMarginCrossMaxBorrowResult {
	this := ApiResponseResultOfMarginCrossMaxBorrowResult{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) GetData() MarginCrossMaxBorrowResult {
	if o == nil || isNil(o.Data) {
		var ret MarginCrossMaxBorrowResult
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) GetDataOk() (*MarginCrossMaxBorrowResult, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MarginCrossMaxBorrowResult and assigns it to the Data field.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) SetData(v MarginCrossMaxBorrowResult) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) GetMsg() string {
	if o == nil || isNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) GetMsgOk() (*string, bool) {
	if o == nil || isNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) HasMsg() bool {
	if o != nil && !isNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) SetMsg(v string) {
	o.Msg = &v
}

// GetRequestTime returns the RequestTime field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) GetRequestTime() int64 {
	if o == nil || isNil(o.RequestTime) {
		var ret int64
		return ret
	}
	return *o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) GetRequestTimeOk() (*int64, bool) {
	if o == nil || isNil(o.RequestTime) {
		return nil, false
	}
	return o.RequestTime, true
}

// HasRequestTime returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) HasRequestTime() bool {
	if o != nil && !isNil(o.RequestTime) {
		return true
	}

	return false
}

// SetRequestTime gets a reference to the given int64 and assigns it to the RequestTime field.
func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) SetRequestTime(v int64) {
	o.RequestTime = &v
}

func (o ApiResponseResultOfMarginCrossMaxBorrowResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !isNil(o.RequestTime) {
		toSerialize["requestTime"] = o.RequestTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApiResponseResultOfMarginCrossMaxBorrowResult) UnmarshalJSON(bytes []byte) (err error) {
	varApiResponseResultOfMarginCrossMaxBorrowResult := _ApiResponseResultOfMarginCrossMaxBorrowResult{}

	if err = json.Unmarshal(bytes, &varApiResponseResultOfMarginCrossMaxBorrowResult); err == nil {
		*o = ApiResponseResultOfMarginCrossMaxBorrowResult(varApiResponseResultOfMarginCrossMaxBorrowResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "data")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "requestTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiResponseResultOfMarginCrossMaxBorrowResult struct {
	value *ApiResponseResultOfMarginCrossMaxBorrowResult
	isSet bool
}

func (v NullableApiResponseResultOfMarginCrossMaxBorrowResult) Get() *ApiResponseResultOfMarginCrossMaxBorrowResult {
	return v.value
}

func (v *NullableApiResponseResultOfMarginCrossMaxBorrowResult) Set(val *ApiResponseResultOfMarginCrossMaxBorrowResult) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponseResultOfMarginCrossMaxBorrowResult) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponseResultOfMarginCrossMaxBorrowResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponseResultOfMarginCrossMaxBorrowResult(val *ApiResponseResultOfMarginCrossMaxBorrowResult) *NullableApiResponseResultOfMarginCrossMaxBorrowResult {
	return &NullableApiResponseResultOfMarginCrossMaxBorrowResult{value: val, isSet: true}
}

func (v NullableApiResponseResultOfMarginCrossMaxBorrowResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponseResultOfMarginCrossMaxBorrowResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
