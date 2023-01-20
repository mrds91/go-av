package model

type ExpAddressReq struct {
	Addresses []AddressList `json:"addresses"`
}

type AddressList struct {
	Address      AddressInfo       `json:"address"`
	AddressTypes []AddressTypeInfo `json:"addressTypes"`
}

type AddressInfo struct {
	Addressline1     string `json:"addressline1"`
	Addressline2     string `json:"addressline2"`
	City             string `json:"city"`
	State            string `json:"state"`
	ZipCode          string `json:"zipCode"`
	ZipCodeExtension string `json:"zipCodeExtension"`
}

type AddressTypeInfo struct {
	AddressUsageType string `json:"addressUsageType"`
}
