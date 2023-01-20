package model

type ExpAddressRes struct {
	AddressInfo []addressInfo //`json:"addressInfo"`
}

type addressInfo struct {
	AddressStatusCode string      //`json:"addressStatusCode"`
	AddressInfo       AddressInfo //`json:"address"`
	AddressType       string
}
