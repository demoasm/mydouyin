package consts

const (
	UserTableName       = "user"
	VideoTableName      = "video"
	RelationTableName   = "relation"
	FavoriteTableName   = "favorite"
	SecretKey           = "secret key"
	IdentityKey         = "id"
	Total               = "total"
	ApiServiceName      = "douyinapi"
	UserServiceName     = "douyinuser"
	VideoServiceName    = "douyinvideo"
	RelationServiceName = "relation"
	FavoriteServiceName = "douyinfavorite"
	MySQLDefaultDSN     = "douyin:douyin@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = ":9000"
	VideoServiceAddr    = ":10000"
	RelationServiceAddr = ":12000"
	FavoriteServiceAddr = ":11000"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:2379"
	DefaultLimit        = 10
	CDNURL              = "http://192.168.251.93:8080/"
	StaticRoot          = "/home/jszfree/mydouyin/"
)
