#### 字符串操作

```shell
go get github.com/Anderson-Lu/gohelper/string
```

|所属包|方法|功能|
|-----|-----|-----|
|`string_helper`|`FindUrl(raw string) []string`|从指定字符串中正则匹配出来URL|

#### 数值操作

```shell
go get github.com/Anderson-Lu/gohelper/number
```

|所属包|方法|功能|
|-----|-----|-----|
|`number_helper`|`RandomInt(max int) int`|返回0~max-1的随机数|
|`number_helper`|`Min(a1, a2 float64) float64`|返回a1和a2的最小值|
|`number_helper`|`MinN(aN ...int) int`|返回a1到aN的最小值|
|`number_helper`|`MinFloat64N(aN ...float64) float64`|返回a1到aN的最小值|
|`number_helper`|`FloorOrCeil(n int, raw float64, isUp bool) float64 `|指定精度位数进行向上或者向取整|
|`number_helper`|`CalcDigist(n float64) int`|返回精度|
|`number_helper`|`Round(f float64, n int) float64`|指定精度位数进行向下取整|
|`number_helper`|`CalcAverage(data Avg) float64`|计算平均数|
|`number_helper`|`CalcVariance(data Variance) float64`|计算方差|

#### 类型转换

```shell
go get github.com/Anderson-Lu/gohelper/convert
```

|所属包|方法|功能|
|-----|-----|-----|
|`convert_helper`|`ConvertBs2Interface(bs []byte) (map[string]interface{}, error)`|将字节数组转化为对象|

#### 定时任务

```shell
go get github.com/Anderson-Lu/gohelper/cron
```

|所属包|方法|功能|
|-----|-----|-----|
|`cron_helper`|`AddJob(job func(), errHandler func(error), taskName string, cronTab string)`|添加cron异步定时任务|

#### 并发任务

```shell
go get github.com/Anderson-Lu/gohelper/concurrent
```

|所属包|方法|功能|
|-----|-----|-----|
|`concurrent_helper`|`AddTask(task func(...interface{}), params ...interface{})`|添加并发任务|

#### 重试任务

```shell
go get github.com/Anderson-Lu/gohelper/retry
```

|所属包|方法|功能|
|-----|-----|-----|
|`retry_helper`|`Retry(retryTimes int, errHandler func(error), job func() error) error`|指定重试运行指定任务|

#### 数据库操作

```shell
go get github.com/Anderson-Lu/gohelper/database
```

|所属包|方法|功能|
|-----|-----|-----|
|`mongo_helper`|`NewMongoClient(connStr string) (*mgo.Session,error)`|生成mongo连接客户端|
|`mysql_helper`|`NewMysqlClient(host, dbname, user, pwd string) (gorm.DB*, error)`|生成mysql连接客户端|
|`redis_helper`|`NewRedisClient(addr, pwd string) *RedisClient `|生成redis连接客户端|

#### 中间件操作

```shell
go get github.com/Anderson-Lu/gohelper/middleware
```

|所属包|方法|功能|
|-----|-----|-----|
|`kafka_helper`|`NewKafkaProducer(address []string) (*KafkaProducer, error)`|生成kafka连接客户端|

#### 日志操作

```shell
go get github.com/Anderson-Lu/gohelper/log
```

|所属包|方法|功能|
|-----|-----|-----|
|`log_helper`|`NewLogger() *logging.Logger`|返回日志操作对象|



