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

// ApiResponseResultOfMarginBatchCancelOrderResult struct for ApiResponseResultOfMarginBatchCancelOrderResult
type ApiResponseResultOfMarginBatchCancelOrderResult struct {
	// code
	Code *string                       `json:"code,omitempty"`
	Data *MarginBatchCancelOrderResult `json:"data,omitempty"`
	// msg
	Msg *string `json:"msg,omitempty"`
	// requestTime
	RequestTime          *int64 `json:"requestTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiResponseResultOfMarginBatchCancelOrderResult ApiResponseResultOfMarginBatchCancelOrderResult

// NewApiResponseResultOfMarginBatchCancelOrderResult instantiates a new ApiResponseResultOfMarginBatchCancelOrderResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponseResultOfMarginBatchCancelOrderResult() *ApiResponseResultOfMarginBatchCancelOrderResult {
	this := ApiResponseResultOfMarginBatchCancelOrderResult{}
	return &this
}

// NewApiResponseResultOfMarginBatchCancelOrderResultWithDefaults instantiates a new ApiResponseResultOfMarginBatchCancelOrderResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponseResultOfMarginBatchCancelOrderResultWithDefaults() *ApiResponseResultOfMarginBatchCancelOrderResult {
	this := ApiResponseResultOfMarginBatchCancelOrderResult{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) GetData() MarginBatchCancelOrderResult {
	if o == nil || isNil(o.Data) {
		var ret MarginBatchCancelOrderResult
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) GetDataOk() (*MarginBatchCancelOrderResult, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MarginBatchCancelOrderResult and assigns it to the Data field.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) SetData(v MarginBatchCancelOrderResult) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) GetMsg() string {
	if o == nil || isNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) GetMsgOk() (*string, bool) {
	if o == nil || isNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) HasMsg() bool {
	if o != nil && !isNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) SetMsg(v string) {
	o.Msg = &v
}

// GetRequestTime returns the RequestTime field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) GetRequestTime() int64 {
	if o == nil || isNil(o.RequestTime) {
		var ret int64
		return ret
	}
	return *o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) GetRequestTimeOk() (*int64, bool) {
	if o == nil || isNil(o.RequestTime) {
		return nil, false
	}
	return o.RequestTime, true
}

// HasRequestTime returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) HasRequestTime() bool {
	if o != nil && !isNil(o.RequestTime) {
		return true
	}

	return false
}

// SetRequestTime gets a reference to the given int64 and assigns it to the RequestTime field.
func (o *ApiResponseResultOfMarginBatchCancelOrderResult) SetRequestTime(v int64) {
	o.RequestTime = &v
}

func (o ApiResponseResultOfMarginBatchCancelOrderResult) MarshalJSON() ([]byte, error) {
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

func (o *ApiResponseResultOfMarginBatchCancelOrderResult) UnmarshalJSON(bytes []byte) (err error) {
	varApiResponseResultOfMarginBatchCancelOrderResult := _ApiResponseResultOfMarginBatchCancelOrderResult{}

	if err = json.Unmarshal(bytes, &varApiResponseResultOfMarginBatchCancelOrderResult); err == nil {
		*o = ApiResponseResultOfMarginBatchCancelOrderResult(varApiResponseResultOfMarginBatchCancelOrderResult)
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

type NullableApiResponseResultOfMarginBatchCancelOrderResult struct {
	value *ApiResponseResultOfMarginBatchCancelOrderResult
	isSet bool
}

func (v NullableApiResponseResultOfMarginBatchCancelOrderResult) Get() *ApiResponseResultOfMarginBatchCancelOrderResult {
	return v.value
}

func (v *NullableApiResponseResultOfMarginBatchCancelOrderResult) Set(val *ApiResponseResultOfMarginBatchCancelOrderResult) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponseResultOfMarginBatchCancelOrderResult) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponseResultOfMarginBatchCancelOrderResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponseResultOfMarginBatchCancelOrderResult(val *ApiResponseResultOfMarginBatchCancelOrderResult) *NullableApiResponseResultOfMarginBatchCancelOrderResult {
	return &NullableApiResponseResultOfMarginBatchCancelOrderResult{value: val, isSet: true}
}

func (v NullableApiResponseResultOfMarginBatchCancelOrderResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponseResultOfMarginBatchCancelOrderResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
