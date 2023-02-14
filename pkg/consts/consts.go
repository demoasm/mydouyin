package consts

const (
	UserTableName      = "user"
	VideoTableName     = "video"
	CommentTableName   = "comment"
	SecretKey          = "secret key"
	IdentityKey        = "id"
	Total              = "total"
	ApiServiceName     = "douyinapi"
	UserServiceName    = "douyinuser"
	VideoServiceName   = "douyinvideo"
	CommentServiceName = "douyincomment"
	MySQLDefaultDSN    = "douyin:douyin@tcp(localhost:3308)/douyin?charset=utf8&parseTime=True&loc=Local"
	TCP                = "tcp"
	UserServiceAddr    = ":9000"
	VideoServiceAddr   = ":10000"
	CommentServiceAddr = ":20000"
	ExportEndpoint     = ":4317"
	ETCDAddress        = "127.0.0.1:2379"
	DefaultLimit       = 10
	CDNURL             = "http://172.31.146.48:8080/"
)
