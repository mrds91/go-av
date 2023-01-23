package util

import "github.com/magiconair/properties"

type contextProperty struct {
	AppHost     string
	AppPort     string
	BackendHost string
	BackendPath string
	Cachettl    int
}

var ContextProperty contextProperty

func LoadProperty() {
	p := properties.MustLoadFile("./pkg/config/app.properties", properties.UTF8)
	ContextProperty.AppHost = p.GetString("app.hostname", "localhost")
	ContextProperty.AppPort = p.GetString("app.port", "8080")
	ContextProperty.BackendHost = p.GetString("backend.addr_val.host", "localhost:9000")
	ContextProperty.BackendPath = p.GetString("backend.addr_val.path", "api/common/v1/addresses")
	ContextProperty.Cachettl = p.GetInt("cache.ttl.in.seconds", 60)
}
