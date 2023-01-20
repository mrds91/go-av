package model

type SysAddressInfo struct {
	Addressline1      string `json:"addressline1"`
	Addressline2      string `json:"addressline2"`
	Cityname          string `json:"cityname"`
	Statecode         string `json:"statecode"`
	Postalcode        string `json:"postalcode"`
	Countrycode       string `json:"countrycode"`
	Sourcetype        string `json:"sourcetype"`
	AddressStatusCode string `json:"addressStatusCode"`
	AddressType       string `json:"addressType"`
}
