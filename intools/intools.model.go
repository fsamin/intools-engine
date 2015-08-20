package intools

import (
    . "gopkg.in/redis.v3"
    "github.com/robfig/cron"
	"github.com/samalba/dockerclient"
	"time"
)

type IntoolsEngine interface {
	GetDockerClient() dockerclient.Client
	GetDockerHost() string
	GetRedisClient() RedisWrapper
	GetCron() CronWrapper
}

type CronWrapper interface {
	AddFunc(spec string, cmd func()) error
	AddJob(spec string, cmd cron.Job) error
	Schedule(schedule cron.Schedule, cmd cron.Job)
	Entries() []*cron.Entry
	Start()
	Stop()
}

type RedisWrapper interface {
	Process(cmd Cmder)
	Auth(password string) *StatusCmd
	Echo(message string) *StringCmd
	Ping() *StatusCmd
	Quit() *StatusCmd
	Select(index int64) *StatusCmd
	Del(keys ...string) *IntCmd
	Dump(key string) *StringCmd
	Exists(key string) *BoolCmd
	Expire(key string, expiration time.Duration) *BoolCmd
	ExpireAt(key string, tm time.Time) *BoolCmd
	Keys(pattern string) *StringSliceCmd
	Migrate(host, port, key string, db int64, timeout time.Duration) *StatusCmd
	Move(key string, db int64) *BoolCmd
	Multi() *Multi
	ObjectRefCount(keys ...string) *IntCmd
	ObjectEncoding(keys ...string) *StringCmd
	ObjectIdleTime(keys ...string) *DurationCmd
	Persist(key string) *BoolCmd
	PExpire(key string, expiration time.Duration) *BoolCmd
	PExpireAt(key string, tm time.Time) *BoolCmd
	PTTL(key string) *DurationCmd
	RandomKey() *StringCmd
	Rename(key, newkey string) *StatusCmd
	RenameNX(key, newkey string) *BoolCmd
	Restore(key string, ttl time.Duration, value string) *StatusCmd
	Sort(key string, sort Sort) *StringSliceCmd
	TTL(key string) *DurationCmd
	Type(key string) *StatusCmd
	Scan(cursor int64, match string, count int64) *ScanCmd
	SScan(key string, cursor int64, match string, count int64) *ScanCmd
	HScan(key string, cursor int64, match string, count int64) *ScanCmd
	ZScan(key string, cursor int64, match string, count int64) *ScanCmd
	Append(key, value string) *IntCmd
	BitCount(key string, bitCount *BitCount) *IntCmd
	BitOpAnd(destKey string, keys ...string) *IntCmd
	BitOpOr(destKey string, keys ...string) *IntCmd
	BitOpXor(destKey string, keys ...string) *IntCmd
	BitOpNot(destKey string, key string) *IntCmd
	BitPos(key string, bit int64, pos ...int64) *IntCmd
	Decr(key string) *IntCmd
	DecrBy(key string, decrement int64) *IntCmd
	Get(key string) *StringCmd
	GetBit(key string, offset int64) *IntCmd
	GetRange(key string, start, end int64) *StringCmd
	GetSet(key string, value interface{}) *StringCmd
	Incr(key string) *IntCmd
	IncrBy(key string, value int64) *IntCmd
	IncrByFloat(key string, value float64) *FloatCmd
	MGet(keys ...string) *SliceCmd
	MSet(pairs ...string) *StatusCmd
	MSetNX(pairs ...string) *BoolCmd
	Set(key string, value interface{}, expiration time.Duration) *StatusCmd
	SetBit(key string, offset int64, value int) *IntCmd
	SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd
	SetXX(key string, value interface{}, expiration time.Duration) *BoolCmd
	SetRange(key string, offset int64, value string) *IntCmd
	StrLen(key string) *IntCmd
	HDel(key string, fields ...string) *IntCmd
	HExists(key, field string) *BoolCmd
	HGet(key, field string) *StringCmd
	HGetAll(key string) *StringSliceCmd
	HGetAllMap(key string) *StringStringMapCmd
	HIncrBy(key, field string, incr int64) *IntCmd
	HIncrByFloat(key, field string, incr float64) *FloatCmd
	HKeys(key string) *StringSliceCmd
	HLen(key string) *IntCmd
	HMGet(key string, fields ...string) *SliceCmd
	HMSet(key, field, value string, pairs ...string) *StatusCmd
	HSet(key, field, value string) *BoolCmd
	HSetNX(key, field, value string) *BoolCmd
	HVals(key string) *StringSliceCmd
	BLPop(timeout time.Duration, keys ...string) *StringSliceCmd
	BRPop(timeout time.Duration, keys ...string) *StringSliceCmd
	BRPopLPush(source, destination string, timeout time.Duration) *StringCmd
	LIndex(key string, index int64) *StringCmd
	LInsert(key, op, pivot, value string) *IntCmd
	LLen(key string) *IntCmd
	LPop(key string) *StringCmd
	LPush(key string, values ...string) *IntCmd
	LPushX(key, value string) *IntCmd
	LRange(key string, start, stop int64) *StringSliceCmd
	LRem(key string, count int64, value string) *IntCmd
	LSet(key string, index int64, value string) *StatusCmd
	LTrim(key string, start, stop int64) *StatusCmd
	RPop(key string) *StringCmd
	RPopLPush(source, destination string) *StringCmd
	RPush(key string, values ...string) *IntCmd
	RPushX(key string, value string) *IntCmd
	SAdd(key string, members ...string) *IntCmd
	SCard(key string) *IntCmd
	SDiff(keys ...string) *StringSliceCmd
	SDiffStore(destination string, keys ...string) *IntCmd
	SInter(keys ...string) *StringSliceCmd
	SInterStore(destination string, keys ...string) *IntCmd
	SIsMember(key, member string) *BoolCmd
	SMembers(key string) *StringSliceCmd
	SMove(source, destination, member string) *BoolCmd
	SPop(key string) *StringCmd
	SRandMember(key string) *StringCmd
	SRem(key string, members ...string) *IntCmd
	SUnion(keys ...string) *StringSliceCmd
	SUnionStore(destination string, keys ...string) *IntCmd
	ZAdd(key string, members ...Z) *IntCmd
	ZCard(key string) *IntCmd
	ZCount(key, min, max string) *IntCmd
	ZIncrBy(key string, increment float64, member string) *FloatCmd
	ZInterStore(destination string, store ZStore, keys ...string) *IntCmd
	ZRange(key string, start, stop int64) *StringSliceCmd
	ZRangeWithScores(key string, start, stop int64) *ZSliceCmd
	ZRangeByScore(key string, opt ZRangeByScore) *StringSliceCmd
	ZRangeByScoreWithScores(key string, opt ZRangeByScore) *ZSliceCmd
	ZRank(key, member string) *IntCmd
	ZRem(key string, members ...string) *IntCmd
	ZRemRangeByRank(key string, start, stop int64) *IntCmd
	ZRemRangeByScore(key, min, max string) *IntCmd
	ZRevRange(key string, start, stop int64) *StringSliceCmd
	ZRevRangeWithScores(key string, start, stop int64) *ZSliceCmd
	ZRevRangeByScore(key string, opt ZRangeByScore) *StringSliceCmd
	ZRevRangeByScoreWithScores(key string, opt ZRangeByScore) *ZSliceCmd
	ZRevRank(key, member string) *IntCmd
	ZScore(key, member string) *FloatCmd
	ZUnionStore(dest string, store ZStore, keys ...string) *IntCmd
	BgRewriteAOF() *StatusCmd
	BgSave() *StatusCmd
	ClientKill(ipPort string) *StatusCmd
	ClientList() *StringCmd
	ClientPause(dur time.Duration) *BoolCmd
	ConfigGet(parameter string) *SliceCmd
	ConfigResetStat() *StatusCmd
	ConfigSet(parameter, value string) *StatusCmd
	DbSize() *IntCmd
	FlushAll() *StatusCmd
	FlushDb() *StatusCmd
	Info() *StringCmd
	LastSave() *IntCmd
	Save() *StatusCmd
	Shutdown() *StatusCmd
	ShutdownSave() *StatusCmd
	ShutdownNoSave() *StatusCmd
	SlaveOf(host, port string) *StatusCmd
	SlowLog()
	Sync()
	Time() *StringSliceCmd
	Eval(script string, keys []string, args []string) *Cmd
	EvalSha(sha1 string, keys []string, args []string) *Cmd
	ScriptExists(scripts ...string) *BoolSliceCmd
	ScriptFlush() *StatusCmd
	ScriptKill() *StatusCmd
	ScriptLoad(script string) *StringCmd
	DebugObject(key string) *StringCmd
	PubSubChannels(pattern string) *StringSliceCmd
	PubSubNumSub(channels ...string) *StringIntMapCmd
	PubSubNumPat() *IntCmd
	ClusterSlots() *ClusterSlotCmd
	ClusterNodes() *StringCmd
	ClusterMeet(host, port string) *StatusCmd
	ClusterReplicate(nodeID string) *StatusCmd
	ClusterInfo() *StringCmd
	ClusterFailover() *StatusCmd
	ClusterAddSlots(slots ...int) *StatusCmd
	ClusterAddSlotsRange(min, max int) *StatusCmd
}

type IntoolsEngineImpl struct {
	DockerClient dockerclient.Client
	DockerHost   string
	RedisClient  RedisWrapper
	Cron         CronWrapper
}

func (e *IntoolsEngineImpl) GetDockerClient() dockerclient.Client {
	return e.DockerClient
}

func (e *IntoolsEngineImpl) GetDockerHost() string {
	return e.DockerHost
}

func (e *IntoolsEngineImpl) GetRedisClient() RedisWrapper {
	return e.RedisClient
}

func (e *IntoolsEngineImpl) GetCron() CronWrapper {
	return e.Cron
}

var (
	Engine IntoolsEngine
)
