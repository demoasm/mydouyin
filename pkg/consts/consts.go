package consts

const (
	UserTableName       = "user"
	VideoTableName      = "video"
	CommentTableName    = "comment"
	MessageTableName    = "message"
	SecretKey           = "secret key"
	IdentityKey         = "id"
	Total               = "total"
	ApiServiceName      = "douyinapi"
	UserServiceName     = "douyinuser"
	VideoServiceName    = "douyinvideo"
	CommentServiceName  = "douyincomment"
	MessageServiceName  = "message"
	RelationTableName   = "relation"
	FavoriteTableName   = "favorite"
	RelationServiceName = "relation"
	FavoriteServiceName = "douyinfavorite"
	MySQLDefaultDSN     = "douyin:douyin@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = ":9000"
	VideoServiceAddr    = ":10000"
	RelationServiceAddr = ":12000"
	CommentServiceAddr  = ":13000"
	FavoriteServiceAddr = ":11000"
	MessageServiceAddr  = ":14000"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:10079"
	DefaultLimit        = 10

	//oss相关信息
	Endpoint = "oss-cn-beijing.aliyuncs.com"
	AKID     = "LTAI5tQ4x1ACnZo5brw92kxo"
	AKS      = "SmEavhOQDQ2lBXBaiognBiLuS9N3K9"
	Bucket   = "douyin-video-9567"
	CDNURL   = "http://aliyun.maomint.cn/"
)
