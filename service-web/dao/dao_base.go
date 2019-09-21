package dao

var (
	defaultMysqlDriveName  = "mysql"
	defaultMysqlURL  = "root:123456@tcp(127.0.0.1:3306)/micro_user"
	defaultRedisConnType  = "tcp"
	defaultRedisUrl  = "localhost:6379"
	defaultRedisPassword = ""
	defaultRedisDB = 0
	inited  = false
)

type Option func(o *Options)

type Options struct {
	MysqlDriveName string
	MysqlURL string
	RedisConnType  string
	RedisUrl string
	RedisPassword string
	RedisDB int
}

func Init(opts ...Option){
	if !inited{
	opt := Options{}
	defaultOptions(&opt)
	for _, o := range opts {
		o(&opt)
	}
	InitMysql(opt.MysqlDriveName, opt.MysqlURL)
	InitRedis(opt.RedisPassword, opt.RedisUrl, opt.RedisDB)
	}
	inited = true
}

func defaultOptions(opts *Options){
	opts.MysqlDriveName = defaultMysqlDriveName
	opts.MysqlURL = defaultMysqlURL
	opts.RedisConnType = defaultRedisConnType
	opts.RedisUrl = defaultRedisUrl
	opts.RedisPassword = defaultRedisPassword
}

func SetMysqlDriveName(s string) Option {
	return func(o *Options) {
		o.MysqlDriveName = s
	}
}

func SetMysqlURL(s string) Option {
	return func(o *Options) {
		o.MysqlURL = s
	}
}

func SetRedisConnType(s string) Option {
	return func(o *Options) {
		o.RedisConnType = s
	}
}

func SetRedisUrl(s string) Option {
	return func(o *Options) {
		o.RedisUrl = s
	}
}

func SetRedisPassword(s string) Option{
	return func(o *Options){
		o.RedisPassword = s
	}
}

func SetRedisDB(n int) Option{
	return func(o *Options){
		o.RedisDB = n
	}
}
