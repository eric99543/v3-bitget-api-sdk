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

// MerchantPersonInfo struct for MerchantPersonInfo
type MerchantPersonInfo struct {
	AveragePayment       *string `json:"averagePayment,omitempty"`
	AverageRealese       *string `json:"averageRealese,omitempty"`
	Email                *string `json:"email,omitempty"`
	EmailBindFlag        *bool   `json:"emailBindFlag,omitempty"`
	KycFlag              *bool   `json:"kycFlag,omitempty"`
	MerchantId           *string `json:"merchantId,omitempty"`
	Mobile               *string `json:"mobile,omitempty"`
	MobileBindFlag       *bool   `json:"mobileBindFlag,omitempty"`
	NickName             *string `json:"nickName,omitempty"`
	RealName             *string `json:"realName,omitempty"`
	RegisterTime         *string `json:"registerTime,omitempty"`
	ThirtyBuy            *string `json:"thirtyBuy,omitempty"`
	ThirtyCompletionRate *string `json:"thirtyCompletionRate,omitempty"`
	ThirtySell           *string `json:"thirtySell,omitempty"`
	ThirtyTrades         *string `json:"thirtyTrades,omitempty"`
	TotalBuy             *string `json:"totalBuy,omitempty"`
	TotalCompletionRate  *string `json:"totalCompletionRate,omitempty"`
	TotalSell            *string `json:"totalSell,omitempty"`
	TotalTrades          *string `json:"totalTrades,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MerchantPersonInfo MerchantPersonInfo

// NewMerchantPersonInfo instantiates a new MerchantPersonInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPersonInfo() *MerchantPersonInfo {
	this := MerchantPersonInfo{}
	return &this
}

// NewMerchantPersonInfoWithDefaults instantiates a new MerchantPersonInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPersonInfoWithDefaults() *MerchantPersonInfo {
	this := MerchantPersonInfo{}
	return &this
}

// GetAveragePayment returns the AveragePayment field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetAveragePayment() string {
	if o == nil || isNil(o.AveragePayment) {
		var ret string
		return ret
	}
	return *o.AveragePayment
}

// GetAveragePaymentOk returns a tuple with the AveragePayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetAveragePaymentOk() (*string, bool) {
	if o == nil || isNil(o.AveragePayment) {
		return nil, false
	}
	return o.AveragePayment, true
}

// HasAveragePayment returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasAveragePayment() bool {
	if o != nil && !isNil(o.AveragePayment) {
		return true
	}

	return false
}

// SetAveragePayment gets a reference to the given string and assigns it to the AveragePayment field.
func (o *MerchantPersonInfo) SetAveragePayment(v string) {
	o.AveragePayment = &v
}

// GetAverageRealese returns the AverageRealese field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetAverageRealese() string {
	if o == nil || isNil(o.AverageRealese) {
		var ret string
		return ret
	}
	return *o.AverageRealese
}

// GetAverageRealeseOk returns a tuple with the AverageRealese field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetAverageRealeseOk() (*string, bool) {
	if o == nil || isNil(o.AverageRealese) {
		return nil, false
	}
	return o.AverageRealese, true
}

// HasAverageRealese returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasAverageRealese() bool {
	if o != nil && !isNil(o.AverageRealese) {
		return true
	}

	return false
}

// SetAverageRealese gets a reference to the given string and assigns it to the AverageRealese field.
func (o *MerchantPersonInfo) SetAverageRealese(v string) {
	o.AverageRealese = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MerchantPersonInfo) SetEmail(v string) {
	o.Email = &v
}

// GetEmailBindFlag returns the EmailBindFlag field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetEmailBindFlag() bool {
	if o == nil || isNil(o.EmailBindFlag) {
		var ret bool
		return ret
	}
	return *o.EmailBindFlag
}

// GetEmailBindFlagOk returns a tuple with the EmailBindFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetEmailBindFlagOk() (*bool, bool) {
	if o == nil || isNil(o.EmailBindFlag) {
		return nil, false
	}
	return o.EmailBindFlag, true
}

// HasEmailBindFlag returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasEmailBindFlag() bool {
	if o != nil && !isNil(o.EmailBindFlag) {
		return true
	}

	return false
}

// SetEmailBindFlag gets a reference to the given bool and assigns it to the EmailBindFlag field.
func (o *MerchantPersonInfo) SetEmailBindFlag(v bool) {
	o.EmailBindFlag = &v
}

// GetKycFlag returns the KycFlag field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetKycFlag() bool {
	if o == nil || isNil(o.KycFlag) {
		var ret bool
		return ret
	}
	return *o.KycFlag
}

// GetKycFlagOk returns a tuple with the KycFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetKycFlagOk() (*bool, bool) {
	if o == nil || isNil(o.KycFlag) {
		return nil, false
	}
	return o.KycFlag, true
}

// HasKycFlag returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasKycFlag() bool {
	if o != nil && !isNil(o.KycFlag) {
		return true
	}

	return false
}

// SetKycFlag gets a reference to the given bool and assigns it to the KycFlag field.
func (o *MerchantPersonInfo) SetKycFlag(v bool) {
	o.KycFlag = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetMerchantId() string {
	if o == nil || isNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetMerchantIdOk() (*string, bool) {
	if o == nil || isNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasMerchantId() bool {
	if o != nil && !isNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *MerchantPersonInfo) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetMobile() string {
	if o == nil || isNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetMobileOk() (*string, bool) {
	if o == nil || isNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasMobile() bool {
	if o != nil && !isNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *MerchantPersonInfo) SetMobile(v string) {
	o.Mobile = &v
}

// GetMobileBindFlag returns the MobileBindFlag field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetMobileBindFlag() bool {
	if o == nil || isNil(o.MobileBindFlag) {
		var ret bool
		return ret
	}
	return *o.MobileBindFlag
}

// GetMobileBindFlagOk returns a tuple with the MobileBindFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetMobileBindFlagOk() (*bool, bool) {
	if o == nil || isNil(o.MobileBindFlag) {
		return nil, false
	}
	return o.MobileBindFlag, true
}

// HasMobileBindFlag returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasMobileBindFlag() bool {
	if o != nil && !isNil(o.MobileBindFlag) {
		return true
	}

	return false
}

// SetMobileBindFlag gets a reference to the given bool and assigns it to the MobileBindFlag field.
func (o *MerchantPersonInfo) SetMobileBindFlag(v bool) {
	o.MobileBindFlag = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetNickName() string {
	if o == nil || isNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetNickNameOk() (*string, bool) {
	if o == nil || isNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasNickName() bool {
	if o != nil && !isNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *MerchantPersonInfo) SetNickName(v string) {
	o.NickName = &v
}

// GetRealName returns the RealName field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetRealName() string {
	if o == nil || isNil(o.RealName) {
		var ret string
		return ret
	}
	return *o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetRealNameOk() (*string, bool) {
	if o == nil || isNil(o.RealName) {
		return nil, false
	}
	return o.RealName, true
}

// HasRealName returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasRealName() bool {
	if o != nil && !isNil(o.RealName) {
		return true
	}

	return false
}

// SetRealName gets a reference to the given string and assigns it to the RealName field.
func (o *MerchantPersonInfo) SetRealName(v string) {
	o.RealName = &v
}

// GetRegisterTime returns the RegisterTime field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetRegisterTime() string {
	if o == nil || isNil(o.RegisterTime) {
		var ret string
		return ret
	}
	return *o.RegisterTime
}

// GetRegisterTimeOk returns a tuple with the RegisterTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetRegisterTimeOk() (*string, bool) {
	if o == nil || isNil(o.RegisterTime) {
		return nil, false
	}
	return o.RegisterTime, true
}

// HasRegisterTime returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasRegisterTime() bool {
	if o != nil && !isNil(o.RegisterTime) {
		return true
	}

	return false
}

// SetRegisterTime gets a reference to the given string and assigns it to the RegisterTime field.
func (o *MerchantPersonInfo) SetRegisterTime(v string) {
	o.RegisterTime = &v
}

// GetThirtyBuy returns the ThirtyBuy field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetThirtyBuy() string {
	if o == nil || isNil(o.ThirtyBuy) {
		var ret string
		return ret
	}
	return *o.ThirtyBuy
}

// GetThirtyBuyOk returns a tuple with the ThirtyBuy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetThirtyBuyOk() (*string, bool) {
	if o == nil || isNil(o.ThirtyBuy) {
		return nil, false
	}
	return o.ThirtyBuy, true
}

// HasThirtyBuy returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasThirtyBuy() bool {
	if o != nil && !isNil(o.ThirtyBuy) {
		return true
	}

	return false
}

// SetThirtyBuy gets a reference to the given string and assigns it to the ThirtyBuy field.
func (o *MerchantPersonInfo) SetThirtyBuy(v string) {
	o.ThirtyBuy = &v
}

// GetThirtyCompletionRate returns the ThirtyCompletionRate field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetThirtyCompletionRate() string {
	if o == nil || isNil(o.ThirtyCompletionRate) {
		var ret string
		return ret
	}
	return *o.ThirtyCompletionRate
}

// GetThirtyCompletionRateOk returns a tuple with the ThirtyCompletionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetThirtyCompletionRateOk() (*string, bool) {
	if o == nil || isNil(o.ThirtyCompletionRate) {
		return nil, false
	}
	return o.ThirtyCompletionRate, true
}

// HasThirtyCompletionRate returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasThirtyCompletionRate() bool {
	if o != nil && !isNil(o.ThirtyCompletionRate) {
		return true
	}

	return false
}

// SetThirtyCompletionRate gets a reference to the given string and assigns it to the ThirtyCompletionRate field.
func (o *MerchantPersonInfo) SetThirtyCompletionRate(v string) {
	o.ThirtyCompletionRate = &v
}

// GetThirtySell returns the ThirtySell field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetThirtySell() string {
	if o == nil || isNil(o.ThirtySell) {
		var ret string
		return ret
	}
	return *o.ThirtySell
}

// GetThirtySellOk returns a tuple with the ThirtySell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetThirtySellOk() (*string, bool) {
	if o == nil || isNil(o.ThirtySell) {
		return nil, false
	}
	return o.ThirtySell, true
}

// HasThirtySell returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasThirtySell() bool {
	if o != nil && !isNil(o.ThirtySell) {
		return true
	}

	return false
}

// SetThirtySell gets a reference to the given string and assigns it to the ThirtySell field.
func (o *MerchantPersonInfo) SetThirtySell(v string) {
	o.ThirtySell = &v
}

// GetThirtyTrades returns the ThirtyTrades field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetThirtyTrades() string {
	if o == nil || isNil(o.ThirtyTrades) {
		var ret string
		return ret
	}
	return *o.ThirtyTrades
}

// GetThirtyTradesOk returns a tuple with the ThirtyTrades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetThirtyTradesOk() (*string, bool) {
	if o == nil || isNil(o.ThirtyTrades) {
		return nil, false
	}
	return o.ThirtyTrades, true
}

// HasThirtyTrades returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasThirtyTrades() bool {
	if o != nil && !isNil(o.ThirtyTrades) {
		return true
	}

	return false
}

// SetThirtyTrades gets a reference to the given string and assigns it to the ThirtyTrades field.
func (o *MerchantPersonInfo) SetThirtyTrades(v string) {
	o.ThirtyTrades = &v
}

// GetTotalBuy returns the TotalBuy field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetTotalBuy() string {
	if o == nil || isNil(o.TotalBuy) {
		var ret string
		return ret
	}
	return *o.TotalBuy
}

// GetTotalBuyOk returns a tuple with the TotalBuy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetTotalBuyOk() (*string, bool) {
	if o == nil || isNil(o.TotalBuy) {
		return nil, false
	}
	return o.TotalBuy, true
}

// HasTotalBuy returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasTotalBuy() bool {
	if o != nil && !isNil(o.TotalBuy) {
		return true
	}

	return false
}

// SetTotalBuy gets a reference to the given string and assigns it to the TotalBuy field.
func (o *MerchantPersonInfo) SetTotalBuy(v string) {
	o.TotalBuy = &v
}

// GetTotalCompletionRate returns the TotalCompletionRate field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetTotalCompletionRate() string {
	if o == nil || isNil(o.TotalCompletionRate) {
		var ret string
		return ret
	}
	return *o.TotalCompletionRate
}

// GetTotalCompletionRateOk returns a tuple with the TotalCompletionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetTotalCompletionRateOk() (*string, bool) {
	if o == nil || isNil(o.TotalCompletionRate) {
		return nil, false
	}
	return o.TotalCompletionRate, true
}

// HasTotalCompletionRate returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasTotalCompletionRate() bool {
	if o != nil && !isNil(o.TotalCompletionRate) {
		return true
	}

	return false
}

// SetTotalCompletionRate gets a reference to the given string and assigns it to the TotalCompletionRate field.
func (o *MerchantPersonInfo) SetTotalCompletionRate(v string) {
	o.TotalCompletionRate = &v
}

// GetTotalSell returns the TotalSell field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetTotalSell() string {
	if o == nil || isNil(o.TotalSell) {
		var ret string
		return ret
	}
	return *o.TotalSell
}

// GetTotalSellOk returns a tuple with the TotalSell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetTotalSellOk() (*string, bool) {
	if o == nil || isNil(o.TotalSell) {
		return nil, false
	}
	return o.TotalSell, true
}

// HasTotalSell returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasTotalSell() bool {
	if o != nil && !isNil(o.TotalSell) {
		return true
	}

	return false
}

// SetTotalSell gets a reference to the given string and assigns it to the TotalSell field.
func (o *MerchantPersonInfo) SetTotalSell(v string) {
	o.TotalSell = &v
}

// GetTotalTrades returns the TotalTrades field value if set, zero value otherwise.
func (o *MerchantPersonInfo) GetTotalTrades() string {
	if o == nil || isNil(o.TotalTrades) {
		var ret string
		return ret
	}
	return *o.TotalTrades
}

// GetTotalTradesOk returns a tuple with the TotalTrades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPersonInfo) GetTotalTradesOk() (*string, bool) {
	if o == nil || isNil(o.TotalTrades) {
		return nil, false
	}
	return o.TotalTrades, true
}

// HasTotalTrades returns a boolean if a field has been set.
func (o *MerchantPersonInfo) HasTotalTrades() bool {
	if o != nil && !isNil(o.TotalTrades) {
		return true
	}

	return false
}

// SetTotalTrades gets a reference to the given string and assigns it to the TotalTrades field.
func (o *MerchantPersonInfo) SetTotalTrades(v string) {
	o.TotalTrades = &v
}

func (o MerchantPersonInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AveragePayment) {
		toSerialize["averagePayment"] = o.AveragePayment
	}
	if !isNil(o.AverageRealese) {
		toSerialize["averageRealese"] = o.AverageRealese
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.EmailBindFlag) {
		toSerialize["emailBindFlag"] = o.EmailBindFlag
	}
	if !isNil(o.KycFlag) {
		toSerialize["kycFlag"] = o.KycFlag
	}
	if !isNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !isNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !isNil(o.MobileBindFlag) {
		toSerialize["mobileBindFlag"] = o.MobileBindFlag
	}
	if !isNil(o.NickName) {
		toSerialize["nickName"] = o.NickName
	}
	if !isNil(o.RealName) {
		toSerialize["realName"] = o.RealName
	}
	if !isNil(o.RegisterTime) {
		toSerialize["registerTime"] = o.RegisterTime
	}
	if !isNil(o.ThirtyBuy) {
		toSerialize["thirtyBuy"] = o.ThirtyBuy
	}
	if !isNil(o.ThirtyCompletionRate) {
		toSerialize["thirtyCompletionRate"] = o.ThirtyCompletionRate
	}
	if !isNil(o.ThirtySell) {
		toSerialize["thirtySell"] = o.ThirtySell
	}
	if !isNil(o.ThirtyTrades) {
		toSerialize["thirtyTrades"] = o.ThirtyTrades
	}
	if !isNil(o.TotalBuy) {
		toSerialize["totalBuy"] = o.TotalBuy
	}
	if !isNil(o.TotalCompletionRate) {
		toSerialize["totalCompletionRate"] = o.TotalCompletionRate
	}
	if !isNil(o.TotalSell) {
		toSerialize["totalSell"] = o.TotalSell
	}
	if !isNil(o.TotalTrades) {
		toSerialize["totalTrades"] = o.TotalTrades
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MerchantPersonInfo) UnmarshalJSON(bytes []byte) (err error) {
	varMerchantPersonInfo := _MerchantPersonInfo{}

	if err = json.Unmarshal(bytes, &varMerchantPersonInfo); err == nil {
		*o = MerchantPersonInfo(varMerchantPersonInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "averagePayment")
		delete(additionalProperties, "averageRealese")
		delete(additionalProperties, "email")
		delete(additionalProperties, "emailBindFlag")
		delete(additionalProperties, "kycFlag")
		delete(additionalProperties, "merchantId")
		delete(additionalProperties, "mobile")
		delete(additionalProperties, "mobileBindFlag")
		delete(additionalProperties, "nickName")
		delete(additionalProperties, "realName")
		delete(additionalProperties, "registerTime")
		delete(additionalProperties, "thirtyBuy")
		delete(additionalProperties, "thirtyCompletionRate")
		delete(additionalProperties, "thirtySell")
		delete(additionalProperties, "thirtyTrades")
		delete(additionalProperties, "totalBuy")
		delete(additionalProperties, "totalCompletionRate")
		delete(additionalProperties, "totalSell")
		delete(additionalProperties, "totalTrades")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMerchantPersonInfo struct {
	value *MerchantPersonInfo
	isSet bool
}

func (v NullableMerchantPersonInfo) Get() *MerchantPersonInfo {
	return v.value
}

func (v *NullableMerchantPersonInfo) Set(val *MerchantPersonInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPersonInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPersonInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPersonInfo(val *MerchantPersonInfo) *NullableMerchantPersonInfo {
	return &NullableMerchantPersonInfo{value: val, isSet: true}
}

func (v NullableMerchantPersonInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPersonInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
