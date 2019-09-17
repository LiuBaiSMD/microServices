package dao

var (
	defaultMysqlDriveName string = "mysql"
	defaultMysqlURL string = "root:123456@tcp(127.0.0.1:3306)/micro_user"
	defaultRedisConnType string = "tcp"
	defaultRedisUrl string = "localhost:6379"
	inited bool = false
)

type Option func(o *Options)

type Options struct {
	MysqlDriveName string
	MysqlURL string
	RedisConnType  string
	RedisUrl string
}

func Init(opts ...Option){
	if !inited{
	opt := Options{}
	defaultOptions(&opt)
	for _, o := range opts {
		o(&opt)
	}
	InitMysql(opt.MysqlDriveName, opt.MysqlURL)
	InitTokenRedis(opt.RedisConnType, opt.RedisUrl)
	}
	inited = true
}

func defaultOptions(opts *Options){
	opts.MysqlDriveName = defaultMysqlDriveName
	opts.MysqlURL = defaultMysqlURL
	opts.RedisConnType = defaultRedisConnType
	opts.RedisUrl = defaultRedisUrl
}

func SetMysqlDriveName(n string) Option {
	return func(o *Options) {
		o.MysqlDriveName = n
	}
}

func SetMysqlURL(n string) Option {
	return func(o *Options) {
		o.MysqlURL = n
	}
}

func SetRedisConnType(n string) Option {
	return func(o *Options) {
		o.RedisConnType = n
	}
}

func SetRedisUrl(n string) Option {
	return func(o *Options) {
		o.RedisUrl = n
	}
}


