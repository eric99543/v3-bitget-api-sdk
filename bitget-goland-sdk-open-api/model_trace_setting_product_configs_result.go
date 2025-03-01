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

// TraceSettingProductConfigsResult struct for TraceSettingProductConfigsResult
type TraceSettingProductConfigsResult struct {
	BusinessLine             *string `json:"businessLine,omitempty"`
	MaxStopLossRation        *string `json:"maxStopLossRation,omitempty"`
	MaxStopProfitRation      *string `json:"maxStopProfitRation,omitempty"`
	MaxTraceAmount           *string `json:"maxTraceAmount,omitempty"`
	MaxTraceAmountSystem     *string `json:"maxTraceAmountSystem,omitempty"`
	MaxTraceCount            *string `json:"maxTraceCount,omitempty"`
	MaxTraceRation           *string `json:"maxTraceRation,omitempty"`
	MinStopLossRation        *string `json:"minStopLossRation,omitempty"`
	MinStopProfitRation      *string `json:"minStopProfitRation,omitempty"`
	MinTraceAmount           *string `json:"minTraceAmount,omitempty"`
	MinTraceCount            *string `json:"minTraceCount,omitempty"`
	MinTraceRation           *string `json:"minTraceRation,omitempty"`
	SliderMaxStopLossRatio   *string `json:"sliderMaxStopLossRatio,omitempty"`
	SliderMaxStopProfitRatio *string `json:"sliderMaxStopProfitRatio,omitempty"`
	SymbolId                 *string `json:"symbolId,omitempty"`
	SymbolName               *string `json:"symbolName,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _TraceSettingProductConfigsResult TraceSettingProductConfigsResult

// NewTraceSettingProductConfigsResult instantiates a new TraceSettingProductConfigsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceSettingProductConfigsResult() *TraceSettingProductConfigsResult {
	this := TraceSettingProductConfigsResult{}
	return &this
}

// NewTraceSettingProductConfigsResultWithDefaults instantiates a new TraceSettingProductConfigsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceSettingProductConfigsResultWithDefaults() *TraceSettingProductConfigsResult {
	this := TraceSettingProductConfigsResult{}
	return &this
}

// GetBusinessLine returns the BusinessLine field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetBusinessLine() string {
	if o == nil || isNil(o.BusinessLine) {
		var ret string
		return ret
	}
	return *o.BusinessLine
}

// GetBusinessLineOk returns a tuple with the BusinessLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetBusinessLineOk() (*string, bool) {
	if o == nil || isNil(o.BusinessLine) {
		return nil, false
	}
	return o.BusinessLine, true
}

// HasBusinessLine returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasBusinessLine() bool {
	if o != nil && !isNil(o.BusinessLine) {
		return true
	}

	return false
}

// SetBusinessLine gets a reference to the given string and assigns it to the BusinessLine field.
func (o *TraceSettingProductConfigsResult) SetBusinessLine(v string) {
	o.BusinessLine = &v
}

// GetMaxStopLossRation returns the MaxStopLossRation field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMaxStopLossRation() string {
	if o == nil || isNil(o.MaxStopLossRation) {
		var ret string
		return ret
	}
	return *o.MaxStopLossRation
}

// GetMaxStopLossRationOk returns a tuple with the MaxStopLossRation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMaxStopLossRationOk() (*string, bool) {
	if o == nil || isNil(o.MaxStopLossRation) {
		return nil, false
	}
	return o.MaxStopLossRation, true
}

// HasMaxStopLossRation returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMaxStopLossRation() bool {
	if o != nil && !isNil(o.MaxStopLossRation) {
		return true
	}

	return false
}

// SetMaxStopLossRation gets a reference to the given string and assigns it to the MaxStopLossRation field.
func (o *TraceSettingProductConfigsResult) SetMaxStopLossRation(v string) {
	o.MaxStopLossRation = &v
}

// GetMaxStopProfitRation returns the MaxStopProfitRation field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMaxStopProfitRation() string {
	if o == nil || isNil(o.MaxStopProfitRation) {
		var ret string
		return ret
	}
	return *o.MaxStopProfitRation
}

// GetMaxStopProfitRationOk returns a tuple with the MaxStopProfitRation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMaxStopProfitRationOk() (*string, bool) {
	if o == nil || isNil(o.MaxStopProfitRation) {
		return nil, false
	}
	return o.MaxStopProfitRation, true
}

// HasMaxStopProfitRation returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMaxStopProfitRation() bool {
	if o != nil && !isNil(o.MaxStopProfitRation) {
		return true
	}

	return false
}

// SetMaxStopProfitRation gets a reference to the given string and assigns it to the MaxStopProfitRation field.
func (o *TraceSettingProductConfigsResult) SetMaxStopProfitRation(v string) {
	o.MaxStopProfitRation = &v
}

// GetMaxTraceAmount returns the MaxTraceAmount field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMaxTraceAmount() string {
	if o == nil || isNil(o.MaxTraceAmount) {
		var ret string
		return ret
	}
	return *o.MaxTraceAmount
}

// GetMaxTraceAmountOk returns a tuple with the MaxTraceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMaxTraceAmountOk() (*string, bool) {
	if o == nil || isNil(o.MaxTraceAmount) {
		return nil, false
	}
	return o.MaxTraceAmount, true
}

// HasMaxTraceAmount returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMaxTraceAmount() bool {
	if o != nil && !isNil(o.MaxTraceAmount) {
		return true
	}

	return false
}

// SetMaxTraceAmount gets a reference to the given string and assigns it to the MaxTraceAmount field.
func (o *TraceSettingProductConfigsResult) SetMaxTraceAmount(v string) {
	o.MaxTraceAmount = &v
}

// GetMaxTraceAmountSystem returns the MaxTraceAmountSystem field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMaxTraceAmountSystem() string {
	if o == nil || isNil(o.MaxTraceAmountSystem) {
		var ret string
		return ret
	}
	return *o.MaxTraceAmountSystem
}

// GetMaxTraceAmountSystemOk returns a tuple with the MaxTraceAmountSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMaxTraceAmountSystemOk() (*string, bool) {
	if o == nil || isNil(o.MaxTraceAmountSystem) {
		return nil, false
	}
	return o.MaxTraceAmountSystem, true
}

// HasMaxTraceAmountSystem returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMaxTraceAmountSystem() bool {
	if o != nil && !isNil(o.MaxTraceAmountSystem) {
		return true
	}

	return false
}

// SetMaxTraceAmountSystem gets a reference to the given string and assigns it to the MaxTraceAmountSystem field.
func (o *TraceSettingProductConfigsResult) SetMaxTraceAmountSystem(v string) {
	o.MaxTraceAmountSystem = &v
}

// GetMaxTraceCount returns the MaxTraceCount field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMaxTraceCount() string {
	if o == nil || isNil(o.MaxTraceCount) {
		var ret string
		return ret
	}
	return *o.MaxTraceCount
}

// GetMaxTraceCountOk returns a tuple with the MaxTraceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMaxTraceCountOk() (*string, bool) {
	if o == nil || isNil(o.MaxTraceCount) {
		return nil, false
	}
	return o.MaxTraceCount, true
}

// HasMaxTraceCount returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMaxTraceCount() bool {
	if o != nil && !isNil(o.MaxTraceCount) {
		return true
	}

	return false
}

// SetMaxTraceCount gets a reference to the given string and assigns it to the MaxTraceCount field.
func (o *TraceSettingProductConfigsResult) SetMaxTraceCount(v string) {
	o.MaxTraceCount = &v
}

// GetMaxTraceRation returns the MaxTraceRation field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMaxTraceRation() string {
	if o == nil || isNil(o.MaxTraceRation) {
		var ret string
		return ret
	}
	return *o.MaxTraceRation
}

// GetMaxTraceRationOk returns a tuple with the MaxTraceRation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMaxTraceRationOk() (*string, bool) {
	if o == nil || isNil(o.MaxTraceRation) {
		return nil, false
	}
	return o.MaxTraceRation, true
}

// HasMaxTraceRation returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMaxTraceRation() bool {
	if o != nil && !isNil(o.MaxTraceRation) {
		return true
	}

	return false
}

// SetMaxTraceRation gets a reference to the given string and assigns it to the MaxTraceRation field.
func (o *TraceSettingProductConfigsResult) SetMaxTraceRation(v string) {
	o.MaxTraceRation = &v
}

// GetMinStopLossRation returns the MinStopLossRation field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMinStopLossRation() string {
	if o == nil || isNil(o.MinStopLossRation) {
		var ret string
		return ret
	}
	return *o.MinStopLossRation
}

// GetMinStopLossRationOk returns a tuple with the MinStopLossRation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMinStopLossRationOk() (*string, bool) {
	if o == nil || isNil(o.MinStopLossRation) {
		return nil, false
	}
	return o.MinStopLossRation, true
}

// HasMinStopLossRation returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMinStopLossRation() bool {
	if o != nil && !isNil(o.MinStopLossRation) {
		return true
	}

	return false
}

// SetMinStopLossRation gets a reference to the given string and assigns it to the MinStopLossRation field.
func (o *TraceSettingProductConfigsResult) SetMinStopLossRation(v string) {
	o.MinStopLossRation = &v
}

// GetMinStopProfitRation returns the MinStopProfitRation field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMinStopProfitRation() string {
	if o == nil || isNil(o.MinStopProfitRation) {
		var ret string
		return ret
	}
	return *o.MinStopProfitRation
}

// GetMinStopProfitRationOk returns a tuple with the MinStopProfitRation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMinStopProfitRationOk() (*string, bool) {
	if o == nil || isNil(o.MinStopProfitRation) {
		return nil, false
	}
	return o.MinStopProfitRation, true
}

// HasMinStopProfitRation returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMinStopProfitRation() bool {
	if o != nil && !isNil(o.MinStopProfitRation) {
		return true
	}

	return false
}

// SetMinStopProfitRation gets a reference to the given string and assigns it to the MinStopProfitRation field.
func (o *TraceSettingProductConfigsResult) SetMinStopProfitRation(v string) {
	o.MinStopProfitRation = &v
}

// GetMinTraceAmount returns the MinTraceAmount field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMinTraceAmount() string {
	if o == nil || isNil(o.MinTraceAmount) {
		var ret string
		return ret
	}
	return *o.MinTraceAmount
}

// GetMinTraceAmountOk returns a tuple with the MinTraceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMinTraceAmountOk() (*string, bool) {
	if o == nil || isNil(o.MinTraceAmount) {
		return nil, false
	}
	return o.MinTraceAmount, true
}

// HasMinTraceAmount returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMinTraceAmount() bool {
	if o != nil && !isNil(o.MinTraceAmount) {
		return true
	}

	return false
}

// SetMinTraceAmount gets a reference to the given string and assigns it to the MinTraceAmount field.
func (o *TraceSettingProductConfigsResult) SetMinTraceAmount(v string) {
	o.MinTraceAmount = &v
}

// GetMinTraceCount returns the MinTraceCount field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMinTraceCount() string {
	if o == nil || isNil(o.MinTraceCount) {
		var ret string
		return ret
	}
	return *o.MinTraceCount
}

// GetMinTraceCountOk returns a tuple with the MinTraceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMinTraceCountOk() (*string, bool) {
	if o == nil || isNil(o.MinTraceCount) {
		return nil, false
	}
	return o.MinTraceCount, true
}

// HasMinTraceCount returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMinTraceCount() bool {
	if o != nil && !isNil(o.MinTraceCount) {
		return true
	}

	return false
}

// SetMinTraceCount gets a reference to the given string and assigns it to the MinTraceCount field.
func (o *TraceSettingProductConfigsResult) SetMinTraceCount(v string) {
	o.MinTraceCount = &v
}

// GetMinTraceRation returns the MinTraceRation field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetMinTraceRation() string {
	if o == nil || isNil(o.MinTraceRation) {
		var ret string
		return ret
	}
	return *o.MinTraceRation
}

// GetMinTraceRationOk returns a tuple with the MinTraceRation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetMinTraceRationOk() (*string, bool) {
	if o == nil || isNil(o.MinTraceRation) {
		return nil, false
	}
	return o.MinTraceRation, true
}

// HasMinTraceRation returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasMinTraceRation() bool {
	if o != nil && !isNil(o.MinTraceRation) {
		return true
	}

	return false
}

// SetMinTraceRation gets a reference to the given string and assigns it to the MinTraceRation field.
func (o *TraceSettingProductConfigsResult) SetMinTraceRation(v string) {
	o.MinTraceRation = &v
}

// GetSliderMaxStopLossRatio returns the SliderMaxStopLossRatio field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetSliderMaxStopLossRatio() string {
	if o == nil || isNil(o.SliderMaxStopLossRatio) {
		var ret string
		return ret
	}
	return *o.SliderMaxStopLossRatio
}

// GetSliderMaxStopLossRatioOk returns a tuple with the SliderMaxStopLossRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetSliderMaxStopLossRatioOk() (*string, bool) {
	if o == nil || isNil(o.SliderMaxStopLossRatio) {
		return nil, false
	}
	return o.SliderMaxStopLossRatio, true
}

// HasSliderMaxStopLossRatio returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasSliderMaxStopLossRatio() bool {
	if o != nil && !isNil(o.SliderMaxStopLossRatio) {
		return true
	}

	return false
}

// SetSliderMaxStopLossRatio gets a reference to the given string and assigns it to the SliderMaxStopLossRatio field.
func (o *TraceSettingProductConfigsResult) SetSliderMaxStopLossRatio(v string) {
	o.SliderMaxStopLossRatio = &v
}

// GetSliderMaxStopProfitRatio returns the SliderMaxStopProfitRatio field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetSliderMaxStopProfitRatio() string {
	if o == nil || isNil(o.SliderMaxStopProfitRatio) {
		var ret string
		return ret
	}
	return *o.SliderMaxStopProfitRatio
}

// GetSliderMaxStopProfitRatioOk returns a tuple with the SliderMaxStopProfitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetSliderMaxStopProfitRatioOk() (*string, bool) {
	if o == nil || isNil(o.SliderMaxStopProfitRatio) {
		return nil, false
	}
	return o.SliderMaxStopProfitRatio, true
}

// HasSliderMaxStopProfitRatio returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasSliderMaxStopProfitRatio() bool {
	if o != nil && !isNil(o.SliderMaxStopProfitRatio) {
		return true
	}

	return false
}

// SetSliderMaxStopProfitRatio gets a reference to the given string and assigns it to the SliderMaxStopProfitRatio field.
func (o *TraceSettingProductConfigsResult) SetSliderMaxStopProfitRatio(v string) {
	o.SliderMaxStopProfitRatio = &v
}

// GetSymbolId returns the SymbolId field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetSymbolId() string {
	if o == nil || isNil(o.SymbolId) {
		var ret string
		return ret
	}
	return *o.SymbolId
}

// GetSymbolIdOk returns a tuple with the SymbolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetSymbolIdOk() (*string, bool) {
	if o == nil || isNil(o.SymbolId) {
		return nil, false
	}
	return o.SymbolId, true
}

// HasSymbolId returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasSymbolId() bool {
	if o != nil && !isNil(o.SymbolId) {
		return true
	}

	return false
}

// SetSymbolId gets a reference to the given string and assigns it to the SymbolId field.
func (o *TraceSettingProductConfigsResult) SetSymbolId(v string) {
	o.SymbolId = &v
}

// GetSymbolName returns the SymbolName field value if set, zero value otherwise.
func (o *TraceSettingProductConfigsResult) GetSymbolName() string {
	if o == nil || isNil(o.SymbolName) {
		var ret string
		return ret
	}
	return *o.SymbolName
}

// GetSymbolNameOk returns a tuple with the SymbolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSettingProductConfigsResult) GetSymbolNameOk() (*string, bool) {
	if o == nil || isNil(o.SymbolName) {
		return nil, false
	}
	return o.SymbolName, true
}

// HasSymbolName returns a boolean if a field has been set.
func (o *TraceSettingProductConfigsResult) HasSymbolName() bool {
	if o != nil && !isNil(o.SymbolName) {
		return true
	}

	return false
}

// SetSymbolName gets a reference to the given string and assigns it to the SymbolName field.
func (o *TraceSettingProductConfigsResult) SetSymbolName(v string) {
	o.SymbolName = &v
}

func (o TraceSettingProductConfigsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BusinessLine) {
		toSerialize["businessLine"] = o.BusinessLine
	}
	if !isNil(o.MaxStopLossRation) {
		toSerialize["maxStopLossRation"] = o.MaxStopLossRation
	}
	if !isNil(o.MaxStopProfitRation) {
		toSerialize["maxStopProfitRation"] = o.MaxStopProfitRation
	}
	if !isNil(o.MaxTraceAmount) {
		toSerialize["maxTraceAmount"] = o.MaxTraceAmount
	}
	if !isNil(o.MaxTraceAmountSystem) {
		toSerialize["maxTraceAmountSystem"] = o.MaxTraceAmountSystem
	}
	if !isNil(o.MaxTraceCount) {
		toSerialize["maxTraceCount"] = o.MaxTraceCount
	}
	if !isNil(o.MaxTraceRation) {
		toSerialize["maxTraceRation"] = o.MaxTraceRation
	}
	if !isNil(o.MinStopLossRation) {
		toSerialize["minStopLossRation"] = o.MinStopLossRation
	}
	if !isNil(o.MinStopProfitRation) {
		toSerialize["minStopProfitRation"] = o.MinStopProfitRation
	}
	if !isNil(o.MinTraceAmount) {
		toSerialize["minTraceAmount"] = o.MinTraceAmount
	}
	if !isNil(o.MinTraceCount) {
		toSerialize["minTraceCount"] = o.MinTraceCount
	}
	if !isNil(o.MinTraceRation) {
		toSerialize["minTraceRation"] = o.MinTraceRation
	}
	if !isNil(o.SliderMaxStopLossRatio) {
		toSerialize["sliderMaxStopLossRatio"] = o.SliderMaxStopLossRatio
	}
	if !isNil(o.SliderMaxStopProfitRatio) {
		toSerialize["sliderMaxStopProfitRatio"] = o.SliderMaxStopProfitRatio
	}
	if !isNil(o.SymbolId) {
		toSerialize["symbolId"] = o.SymbolId
	}
	if !isNil(o.SymbolName) {
		toSerialize["symbolName"] = o.SymbolName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TraceSettingProductConfigsResult) UnmarshalJSON(bytes []byte) (err error) {
	varTraceSettingProductConfigsResult := _TraceSettingProductConfigsResult{}

	if err = json.Unmarshal(bytes, &varTraceSettingProductConfigsResult); err == nil {
		*o = TraceSettingProductConfigsResult(varTraceSettingProductConfigsResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "businessLine")
		delete(additionalProperties, "maxStopLossRation")
		delete(additionalProperties, "maxStopProfitRation")
		delete(additionalProperties, "maxTraceAmount")
		delete(additionalProperties, "maxTraceAmountSystem")
		delete(additionalProperties, "maxTraceCount")
		delete(additionalProperties, "maxTraceRation")
		delete(additionalProperties, "minStopLossRation")
		delete(additionalProperties, "minStopProfitRation")
		delete(additionalProperties, "minTraceAmount")
		delete(additionalProperties, "minTraceCount")
		delete(additionalProperties, "minTraceRation")
		delete(additionalProperties, "sliderMaxStopLossRatio")
		delete(additionalProperties, "sliderMaxStopProfitRatio")
		delete(additionalProperties, "symbolId")
		delete(additionalProperties, "symbolName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTraceSettingProductConfigsResult struct {
	value *TraceSettingProductConfigsResult
	isSet bool
}

func (v NullableTraceSettingProductConfigsResult) Get() *TraceSettingProductConfigsResult {
	return v.value
}

func (v *NullableTraceSettingProductConfigsResult) Set(val *TraceSettingProductConfigsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceSettingProductConfigsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceSettingProductConfigsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceSettingProductConfigsResult(val *TraceSettingProductConfigsResult) *NullableTraceSettingProductConfigsResult {
	return &NullableTraceSettingProductConfigsResult{value: val, isSet: true}
}

func (v NullableTraceSettingProductConfigsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceSettingProductConfigsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
