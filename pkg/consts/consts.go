package consts

const (
	UserTableName       = "user"
	VideoTableName      = "video"
	RelationTableName   = "relation"
	SecretKey           = "secret key"
	IdentityKey         = "id"
	Total               = "total"
	ApiServiceName      = "douyinapi"
	UserServiceName     = "douyinuser"
	VideoServiceName    = "douyinvideo"
	RelationServiceName = "relation"
	MySQLDefaultDSN     = "douyin:douyin@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = ":9000"
	VideoServiceAddr    = ":10000"
	RelationServiceAddr = ":11000"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:2379"
	DefaultLimit        = 10
	CDNURL              = "http://192.168.56.102:8080/"
)
