package lib

type StatJS struct {
	Json   string
	Addr   string
	Uagent string
}

type Configure struct {
	Redis struct {
		Password string `json:"password"`
		Address  string `json:"address"`
	} `json:"redis"`
	RedisIp struct{
		Password string `json:"password"`
		Address string `json:"address"`
	} `json:"redis_ip"`
	Clickhouse string `json:"clickhouse"`
}

type RawJS struct {
	Point      int             `json:"point"`
	Statistics [][]interface{} `json:"statistics"`
}

type ValidJS struct {
	Point    int
	Datetime int64
	Md5      string
	Len      int
}