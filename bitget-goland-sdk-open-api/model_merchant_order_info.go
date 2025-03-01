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

// MerchantOrderInfo struct for MerchantOrderInfo
type MerchantOrderInfo struct {
	AdvNo                *string                   `json:"advNo,omitempty"`
	Amount               *string                   `json:"amount,omitempty"`
	BuyerRealName        *string                   `json:"buyerRealName,omitempty"`
	Coin                 *string                   `json:"coin,omitempty"`
	Count                *string                   `json:"count,omitempty"`
	Ctime                *string                   `json:"ctime,omitempty"`
	Fiat                 *string                   `json:"fiat,omitempty"`
	OrderId              *string                   `json:"orderId,omitempty"`
	OrderNo              *string                   `json:"orderNo,omitempty"`
	PaymentInfo          *MerchantOrderPaymentInfo `json:"paymentInfo,omitempty"`
	PaymentTime          *string                   `json:"paymentTime,omitempty"`
	Price                *string                   `json:"price,omitempty"`
	ReleaseCoinTime      *string                   `json:"releaseCoinTime,omitempty"`
	RepresentTime        *string                   `json:"representTime,omitempty"`
	SellerRealName       *string                   `json:"sellerRealName,omitempty"`
	Status               *string                   `json:"status,omitempty"`
	Type                 *string                   `json:"type,omitempty"`
	WithdrawTime         *string                   `json:"withdrawTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MerchantOrderInfo MerchantOrderInfo

// NewMerchantOrderInfo instantiates a new MerchantOrderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantOrderInfo() *MerchantOrderInfo {
	this := MerchantOrderInfo{}
	return &this
}

// NewMerchantOrderInfoWithDefaults instantiates a new MerchantOrderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantOrderInfoWithDefaults() *MerchantOrderInfo {
	this := MerchantOrderInfo{}
	return &this
}

// GetAdvNo returns the AdvNo field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetAdvNo() string {
	if o == nil || isNil(o.AdvNo) {
		var ret string
		return ret
	}
	return *o.AdvNo
}

// GetAdvNoOk returns a tuple with the AdvNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetAdvNoOk() (*string, bool) {
	if o == nil || isNil(o.AdvNo) {
		return nil, false
	}
	return o.AdvNo, true
}

// HasAdvNo returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasAdvNo() bool {
	if o != nil && !isNil(o.AdvNo) {
		return true
	}

	return false
}

// SetAdvNo gets a reference to the given string and assigns it to the AdvNo field.
func (o *MerchantOrderInfo) SetAdvNo(v string) {
	o.AdvNo = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetAmount() string {
	if o == nil || isNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetAmountOk() (*string, bool) {
	if o == nil || isNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MerchantOrderInfo) SetAmount(v string) {
	o.Amount = &v
}

// GetBuyerRealName returns the BuyerRealName field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetBuyerRealName() string {
	if o == nil || isNil(o.BuyerRealName) {
		var ret string
		return ret
	}
	return *o.BuyerRealName
}

// GetBuyerRealNameOk returns a tuple with the BuyerRealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetBuyerRealNameOk() (*string, bool) {
	if o == nil || isNil(o.BuyerRealName) {
		return nil, false
	}
	return o.BuyerRealName, true
}

// HasBuyerRealName returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasBuyerRealName() bool {
	if o != nil && !isNil(o.BuyerRealName) {
		return true
	}

	return false
}

// SetBuyerRealName gets a reference to the given string and assigns it to the BuyerRealName field.
func (o *MerchantOrderInfo) SetBuyerRealName(v string) {
	o.BuyerRealName = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetCoin() string {
	if o == nil || isNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetCoinOk() (*string, bool) {
	if o == nil || isNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasCoin() bool {
	if o != nil && !isNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *MerchantOrderInfo) SetCoin(v string) {
	o.Coin = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetCount() string {
	if o == nil || isNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetCountOk() (*string, bool) {
	if o == nil || isNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *MerchantOrderInfo) SetCount(v string) {
	o.Count = &v
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetCtime() string {
	if o == nil || isNil(o.Ctime) {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetCtimeOk() (*string, bool) {
	if o == nil || isNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasCtime() bool {
	if o != nil && !isNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *MerchantOrderInfo) SetCtime(v string) {
	o.Ctime = &v
}

// GetFiat returns the Fiat field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetFiat() string {
	if o == nil || isNil(o.Fiat) {
		var ret string
		return ret
	}
	return *o.Fiat
}

// GetFiatOk returns a tuple with the Fiat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetFiatOk() (*string, bool) {
	if o == nil || isNil(o.Fiat) {
		return nil, false
	}
	return o.Fiat, true
}

// HasFiat returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasFiat() bool {
	if o != nil && !isNil(o.Fiat) {
		return true
	}

	return false
}

// SetFiat gets a reference to the given string and assigns it to the Fiat field.
func (o *MerchantOrderInfo) SetFiat(v string) {
	o.Fiat = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *MerchantOrderInfo) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetOrderNo() string {
	if o == nil || isNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetOrderNoOk() (*string, bool) {
	if o == nil || isNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasOrderNo() bool {
	if o != nil && !isNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *MerchantOrderInfo) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetPaymentInfo returns the PaymentInfo field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetPaymentInfo() MerchantOrderPaymentInfo {
	if o == nil || isNil(o.PaymentInfo) {
		var ret MerchantOrderPaymentInfo
		return ret
	}
	return *o.PaymentInfo
}

// GetPaymentInfoOk returns a tuple with the PaymentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetPaymentInfoOk() (*MerchantOrderPaymentInfo, bool) {
	if o == nil || isNil(o.PaymentInfo) {
		return nil, false
	}
	return o.PaymentInfo, true
}

// HasPaymentInfo returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasPaymentInfo() bool {
	if o != nil && !isNil(o.PaymentInfo) {
		return true
	}

	return false
}

// SetPaymentInfo gets a reference to the given MerchantOrderPaymentInfo and assigns it to the PaymentInfo field.
func (o *MerchantOrderInfo) SetPaymentInfo(v MerchantOrderPaymentInfo) {
	o.PaymentInfo = &v
}

// GetPaymentTime returns the PaymentTime field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetPaymentTime() string {
	if o == nil || isNil(o.PaymentTime) {
		var ret string
		return ret
	}
	return *o.PaymentTime
}

// GetPaymentTimeOk returns a tuple with the PaymentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetPaymentTimeOk() (*string, bool) {
	if o == nil || isNil(o.PaymentTime) {
		return nil, false
	}
	return o.PaymentTime, true
}

// HasPaymentTime returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasPaymentTime() bool {
	if o != nil && !isNil(o.PaymentTime) {
		return true
	}

	return false
}

// SetPaymentTime gets a reference to the given string and assigns it to the PaymentTime field.
func (o *MerchantOrderInfo) SetPaymentTime(v string) {
	o.PaymentTime = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetPrice() string {
	if o == nil || isNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetPriceOk() (*string, bool) {
	if o == nil || isNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MerchantOrderInfo) SetPrice(v string) {
	o.Price = &v
}

// GetReleaseCoinTime returns the ReleaseCoinTime field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetReleaseCoinTime() string {
	if o == nil || isNil(o.ReleaseCoinTime) {
		var ret string
		return ret
	}
	return *o.ReleaseCoinTime
}

// GetReleaseCoinTimeOk returns a tuple with the ReleaseCoinTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetReleaseCoinTimeOk() (*string, bool) {
	if o == nil || isNil(o.ReleaseCoinTime) {
		return nil, false
	}
	return o.ReleaseCoinTime, true
}

// HasReleaseCoinTime returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasReleaseCoinTime() bool {
	if o != nil && !isNil(o.ReleaseCoinTime) {
		return true
	}

	return false
}

// SetReleaseCoinTime gets a reference to the given string and assigns it to the ReleaseCoinTime field.
func (o *MerchantOrderInfo) SetReleaseCoinTime(v string) {
	o.ReleaseCoinTime = &v
}

// GetRepresentTime returns the RepresentTime field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetRepresentTime() string {
	if o == nil || isNil(o.RepresentTime) {
		var ret string
		return ret
	}
	return *o.RepresentTime
}

// GetRepresentTimeOk returns a tuple with the RepresentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetRepresentTimeOk() (*string, bool) {
	if o == nil || isNil(o.RepresentTime) {
		return nil, false
	}
	return o.RepresentTime, true
}

// HasRepresentTime returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasRepresentTime() bool {
	if o != nil && !isNil(o.RepresentTime) {
		return true
	}

	return false
}

// SetRepresentTime gets a reference to the given string and assigns it to the RepresentTime field.
func (o *MerchantOrderInfo) SetRepresentTime(v string) {
	o.RepresentTime = &v
}

// GetSellerRealName returns the SellerRealName field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetSellerRealName() string {
	if o == nil || isNil(o.SellerRealName) {
		var ret string
		return ret
	}
	return *o.SellerRealName
}

// GetSellerRealNameOk returns a tuple with the SellerRealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetSellerRealNameOk() (*string, bool) {
	if o == nil || isNil(o.SellerRealName) {
		return nil, false
	}
	return o.SellerRealName, true
}

// HasSellerRealName returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasSellerRealName() bool {
	if o != nil && !isNil(o.SellerRealName) {
		return true
	}

	return false
}

// SetSellerRealName gets a reference to the given string and assigns it to the SellerRealName field.
func (o *MerchantOrderInfo) SetSellerRealName(v string) {
	o.SellerRealName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MerchantOrderInfo) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MerchantOrderInfo) SetType(v string) {
	o.Type = &v
}

// GetWithdrawTime returns the WithdrawTime field value if set, zero value otherwise.
func (o *MerchantOrderInfo) GetWithdrawTime() string {
	if o == nil || isNil(o.WithdrawTime) {
		var ret string
		return ret
	}
	return *o.WithdrawTime
}

// GetWithdrawTimeOk returns a tuple with the WithdrawTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantOrderInfo) GetWithdrawTimeOk() (*string, bool) {
	if o == nil || isNil(o.WithdrawTime) {
		return nil, false
	}
	return o.WithdrawTime, true
}

// HasWithdrawTime returns a boolean if a field has been set.
func (o *MerchantOrderInfo) HasWithdrawTime() bool {
	if o != nil && !isNil(o.WithdrawTime) {
		return true
	}

	return false
}

// SetWithdrawTime gets a reference to the given string and assigns it to the WithdrawTime field.
func (o *MerchantOrderInfo) SetWithdrawTime(v string) {
	o.WithdrawTime = &v
}

func (o MerchantOrderInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdvNo) {
		toSerialize["advNo"] = o.AdvNo
	}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.BuyerRealName) {
		toSerialize["buyerRealName"] = o.BuyerRealName
	}
	if !isNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Ctime) {
		toSerialize["ctime"] = o.Ctime
	}
	if !isNil(o.Fiat) {
		toSerialize["fiat"] = o.Fiat
	}
	if !isNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !isNil(o.OrderNo) {
		toSerialize["orderNo"] = o.OrderNo
	}
	if !isNil(o.PaymentInfo) {
		toSerialize["paymentInfo"] = o.PaymentInfo
	}
	if !isNil(o.PaymentTime) {
		toSerialize["paymentTime"] = o.PaymentTime
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.ReleaseCoinTime) {
		toSerialize["releaseCoinTime"] = o.ReleaseCoinTime
	}
	if !isNil(o.RepresentTime) {
		toSerialize["representTime"] = o.RepresentTime
	}
	if !isNil(o.SellerRealName) {
		toSerialize["sellerRealName"] = o.SellerRealName
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.WithdrawTime) {
		toSerialize["withdrawTime"] = o.WithdrawTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MerchantOrderInfo) UnmarshalJSON(bytes []byte) (err error) {
	varMerchantOrderInfo := _MerchantOrderInfo{}

	if err = json.Unmarshal(bytes, &varMerchantOrderInfo); err == nil {
		*o = MerchantOrderInfo(varMerchantOrderInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "advNo")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "buyerRealName")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "count")
		delete(additionalProperties, "ctime")
		delete(additionalProperties, "fiat")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderNo")
		delete(additionalProperties, "paymentInfo")
		delete(additionalProperties, "paymentTime")
		delete(additionalProperties, "price")
		delete(additionalProperties, "releaseCoinTime")
		delete(additionalProperties, "representTime")
		delete(additionalProperties, "sellerRealName")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "withdrawTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMerchantOrderInfo struct {
	value *MerchantOrderInfo
	isSet bool
}

func (v NullableMerchantOrderInfo) Get() *MerchantOrderInfo {
	return v.value
}

func (v *NullableMerchantOrderInfo) Set(val *MerchantOrderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantOrderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantOrderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantOrderInfo(val *MerchantOrderInfo) *NullableMerchantOrderInfo {
	return &NullableMerchantOrderInfo{value: val, isSet: true}
}

func (v NullableMerchantOrderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantOrderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
