package tests

import (
	. "gopkg.in/redis.v3"
	"time"
)

type RedisClientMock struct{}

func (r *RedisClientMock) Process(cmd Cmder)                                    { }
func (r *RedisClientMock) Auth(password string) *StatusCmd                      { return nil }
func (r *RedisClientMock) Echo(message string) *StringCmd                       { return nil }
func (r *RedisClientMock) Ping() *StatusCmd                                     { return nil }
func (r *RedisClientMock) Quit() *StatusCmd                                     { return nil }
func (r *RedisClientMock) Select(index int64) *StatusCmd                        { return nil }
func (r *RedisClientMock) Del(keys ...string) *IntCmd                           { return nil }
func (r *RedisClientMock) Dump(key string) *StringCmd                           { return nil }
func (r *RedisClientMock) Exists(key string) *BoolCmd                           { return nil }
func (r *RedisClientMock) Expire(key string, expiration time.Duration) *BoolCmd { return nil }
func (r *RedisClientMock) ExpireAt(key string, tm time.Time) *BoolCmd           { return nil }
func (r *RedisClientMock) Keys(pattern string) *StringSliceCmd                  { return nil }
func (r *RedisClientMock) Migrate(host, port, key string, db int64, timeout time.Duration) *StatusCmd {
	return nil
}
func (r *RedisClientMock) Move(key string, db int64) *BoolCmd                     { return nil }
func (r *RedisClientMock) Multi() *Multi                                          { return nil }
func (r *RedisClientMock) ObjectRefCount(keys ...string) *IntCmd                  { return nil }
func (r *RedisClientMock) ObjectEncoding(keys ...string) *StringCmd               { return nil }
func (r *RedisClientMock) ObjectIdleTime(keys ...string) *DurationCmd             { return nil }
func (r *RedisClientMock) Persist(key string) *BoolCmd                            { return nil }
func (r *RedisClientMock) PExpire(key string, expiration time.Duration) *BoolCmd  { return nil }
func (r *RedisClientMock) PExpireAt(key string, tm time.Time) *BoolCmd            { return nil }
func (r *RedisClientMock) PTTL(key string) *DurationCmd                           { return nil }
func (r *RedisClientMock) RandomKey() *StringCmd                                  { return nil }
func (r *RedisClientMock) Rename(key, newkey string) *StatusCmd                   { return nil }
func (r *RedisClientMock) RenameNX(key, newkey string) *BoolCmd                   { return nil }
func (r *RedisClientMock) Restore(key string, ttl time.Duration, value string) *StatusCmd { return nil }
func (r *RedisClientMock) Sort(key string, sort Sort) *StringSliceCmd             { return nil }
func (r *RedisClientMock) TTL(key string) *DurationCmd                            { return nil }
func (r *RedisClientMock) Type(key string) *StatusCmd                             { return nil }
func (r *RedisClientMock) Scan(cursor int64, match string, count int64) *ScanCmd  { return nil }
func (r *RedisClientMock) SScan(key string, cursor int64, match string, count int64) *ScanCmd {
	return nil
}
func (r *RedisClientMock) HScan(key string, cursor int64, match string, count int64) *ScanCmd {
	return nil
}
func (r *RedisClientMock) ZScan(key string, cursor int64, match string, count int64) *ScanCmd {
	return nil
}
func (r *RedisClientMock) Append(key, value string) *IntCmd                            { return nil }
func (r *RedisClientMock) BitCount(key string, bitCount *BitCount) *IntCmd             { return nil }
func (r *RedisClientMock) BitOpAnd(destKey string, keys ...string) *IntCmd             { return nil }
func (r *RedisClientMock) BitOpOr(destKey string, keys ...string) *IntCmd              { return nil }
func (r *RedisClientMock) BitOpXor(destKey string, keys ...string) *IntCmd             { return nil }
func (r *RedisClientMock) BitOpNot(destKey string, key string) *IntCmd                 { return nil }
func (r *RedisClientMock) BitPos(key string, bit int64, pos ...int64) *IntCmd          { return nil }
func (r *RedisClientMock) Decr(key string) *IntCmd                                     { return nil }
func (r *RedisClientMock) DecrBy(key string, decrement int64) *IntCmd                  { return nil }
func (r *RedisClientMock) Get(key string) *StringCmd                                   { return nil }
func (r *RedisClientMock) GetBit(key string, offset int64) *IntCmd                     { return nil }
func (r *RedisClientMock) GetRange(key string, start, end int64) *StringCmd            { return nil }
func (r *RedisClientMock) GetSet(key string, value interface{}) *StringCmd                         { return nil }
func (r *RedisClientMock) Incr(key string) *IntCmd                                     { return nil }
func (r *RedisClientMock) IncrBy(key string, value int64) *IntCmd                      { return nil }
func (r *RedisClientMock) IncrByFloat(key string, value float64) *FloatCmd             { return nil }
func (r *RedisClientMock) MGet(keys ...string) *SliceCmd                               { return nil }
func (r *RedisClientMock) MSet(pairs ...string) *StatusCmd                             { return nil }
func (r *RedisClientMock) MSetNX(pairs ...string) *BoolCmd                             { return nil }
func (r *RedisClientMock) Set(key string, value interface{}, expiration time.Duration) *StatusCmd  { return nil }
func (r *RedisClientMock) SetBit(key string, offset int64, value int) *IntCmd          { return nil }
func (r *RedisClientMock) SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd  { return nil }
func (r *RedisClientMock) SetXX(key string, value interface{}, expiration time.Duration) *BoolCmd  { return nil }
func (r *RedisClientMock) SetRange(key string, offset int64, value string) *IntCmd     { return nil }
func (r *RedisClientMock) StrLen(key string) *IntCmd                                   { return nil }
func (r *RedisClientMock) HDel(key string, fields ...string) *IntCmd                   { return nil }
func (r *RedisClientMock) HExists(key, field string) *BoolCmd                          { return nil }
func (r *RedisClientMock) HGet(key, field string) *StringCmd                           { return nil }
func (r *RedisClientMock) HGetAll(key string) *StringSliceCmd                          { return nil }
func (r *RedisClientMock) HGetAllMap(key string) *StringStringMapCmd                   { return nil }
func (r *RedisClientMock) HIncrBy(key, field string, incr int64) *IntCmd               { return nil }
func (r *RedisClientMock) HIncrByFloat(key, field string, incr float64) *FloatCmd      { return nil }
func (r *RedisClientMock) HKeys(key string) *StringSliceCmd                            { return nil }
func (r *RedisClientMock) HLen(key string) *IntCmd                                     { return nil }
func (r *RedisClientMock) HMGet(key string, fields ...string) *SliceCmd                { return nil }
func (r *RedisClientMock) HMSet(key, field, value string, pairs ...string) *StatusCmd  { return nil }
func (r *RedisClientMock) HSet(key, field, value string) *BoolCmd                      { return nil }
func (r *RedisClientMock) HSetNX(key, field, value string) *BoolCmd                    { return nil }
func (r *RedisClientMock) HVals(key string) *StringSliceCmd                            { return nil }
func (r *RedisClientMock) BLPop(timeout time.Duration, keys ...string) *StringSliceCmd { return nil }
func (r *RedisClientMock) BRPop(timeout time.Duration, keys ...string) *StringSliceCmd { return nil }
func (r *RedisClientMock) BRPopLPush(source, destination string, timeout time.Duration) *StringCmd {
	return nil
}
func (r *RedisClientMock) LIndex(key string, index int64) *StringCmd                      { return nil }
func (r *RedisClientMock) LInsert(key, op, pivot, value string) *IntCmd                   { return nil }
func (r *RedisClientMock) LLen(key string) *IntCmd                                        { return nil }
func (r *RedisClientMock) LPop(key string) *StringCmd                                     { return nil }
func (r *RedisClientMock) LPush(key string, values ...string) *IntCmd                     { return nil }
func (r *RedisClientMock) LPushX(key, value interface{}) *IntCmd                               { return nil }
func (r *RedisClientMock) LRange(key string, start, stop int64) *StringSliceCmd           { return nil }
func (r *RedisClientMock) LRem(key string, count int64, value interface{}) *IntCmd             { return nil }
func (r *RedisClientMock) LSet(key string, index int64, value interface{}) *StatusCmd          { return nil }
func (r *RedisClientMock) LTrim(key string, start, stop int64) *StatusCmd                 { return nil }
func (r *RedisClientMock) RPop(key string) *StringCmd                                     { return nil }
func (r *RedisClientMock) RPopLPush(source, destination string) *StringCmd                { return nil }
func (r *RedisClientMock) RPush(key string, values ...string) *IntCmd                     { return nil }
func (r *RedisClientMock) RPushX(key string, value interface{}) *IntCmd                        { return nil }
func (r *RedisClientMock) SAdd(key string, members ...string) *IntCmd                     { return nil }
func (r *RedisClientMock) SCard(key string) *IntCmd                                       { return nil }
func (r *RedisClientMock) SDiff(keys ...string) *StringSliceCmd                           { return nil }
func (r *RedisClientMock) SDiffStore(destination string, keys ...string) *IntCmd          { return nil }
func (r *RedisClientMock) SInter(keys ...string) *StringSliceCmd                          { return nil }
func (r *RedisClientMock) SInterStore(destination string, keys ...string) *IntCmd         { return nil }
func (r *RedisClientMock) SIsMember(key string, member interface{}) *BoolCmd                          { return nil }
func (r *RedisClientMock) SMembers(key string) *StringSliceCmd                            { return nil }
func (r *RedisClientMock) SMove(source, destination string, member interface{}) *BoolCmd              { return nil }
func (r *RedisClientMock) SPop(key string) *StringCmd                                     { return nil }
func (r *RedisClientMock) SRandMember(key string) *StringCmd                              { return nil }
func (r *RedisClientMock) SRem(key string, members ...string) *IntCmd                     { return nil }
func (r *RedisClientMock) SUnion(keys ...string) *StringSliceCmd                          { return nil }
func (r *RedisClientMock) SUnionStore(destination string, keys ...string) *IntCmd         { return nil }
func (r *RedisClientMock) ZAdd(key string, members ...Z) *IntCmd                          { return nil }
func (r *RedisClientMock) ZCard(key string) *IntCmd                                       { return nil }
func (r *RedisClientMock) ZCount(key, min, max string) *IntCmd                            { return nil }
func (r *RedisClientMock) ZIncrBy(key string, increment float64, member string) *FloatCmd { return nil }
func (r *RedisClientMock) ZInterStore(destination string, store ZStore, keys ...string) *IntCmd {
	return nil
}
func (r *RedisClientMock) ZRange(key string, start, stop int64) *StringSliceCmd        { return nil }
func (r *RedisClientMock) ZRangeWithScores(key string, start, stop int64) *ZSliceCmd   { return nil }
func (r *RedisClientMock) ZRangeByScore(key string, opt ZRangeByScore) *StringSliceCmd { return nil }
func (r *RedisClientMock) ZRangeByScoreWithScores(key string, opt ZRangeByScore) *ZSliceCmd {
	return nil
}
func (r *RedisClientMock) ZRank(key, member string) *IntCmd                               { return nil }
func (r *RedisClientMock) ZRem(key string, members ...string) *IntCmd                     { return nil }
func (r *RedisClientMock) ZRemRangeByRank(key string, start, stop int64) *IntCmd          { return nil }
func (r *RedisClientMock) ZRemRangeByScore(key, min, max string) *IntCmd                  { return nil }
func (r *RedisClientMock) ZRevRange(key string, start, stop int64) *StringSliceCmd        { return nil }
func (r *RedisClientMock) ZRevRangeWithScores(key string, start, stop int64) *ZSliceCmd   { return nil }
func (r *RedisClientMock) ZRevRangeByScore(key string, opt ZRangeByScore) *StringSliceCmd { return nil }
func (r *RedisClientMock) ZRevRangeByScoreWithScores(key string, opt ZRangeByScore) *ZSliceCmd {
	return nil
}
func (r *RedisClientMock) ZRevRank(key, member string) *IntCmd                           { return nil }
func (r *RedisClientMock) ZScore(key, member string) *FloatCmd                           { return nil }
func (r *RedisClientMock) ZUnionStore(dest string, store ZStore, keys ...string) *IntCmd { return nil }
func (r *RedisClientMock) BgRewriteAOF() *StatusCmd                                      { return nil }
func (r *RedisClientMock) BgSave() *StatusCmd                                            { return nil }
func (r *RedisClientMock) ClientKill(ipPort string) *StatusCmd                           { return nil }
func (r *RedisClientMock) ClientList() *StringCmd                                        { return nil }
func (r *RedisClientMock) ClientPause(dur time.Duration) *BoolCmd                        { return nil }
func (r *RedisClientMock) ConfigGet(parameter string) *SliceCmd                          { return nil }
func (r *RedisClientMock) ConfigResetStat() *StatusCmd                                   { return nil }
func (r *RedisClientMock) ConfigSet(parameter, value string) *StatusCmd                  { return nil }
func (r *RedisClientMock) DbSize() *IntCmd                                               { return nil }
func (r *RedisClientMock) FlushAll() *StatusCmd                                          { return nil }
func (r *RedisClientMock) FlushDb() *StatusCmd                                           { return nil }
func (r *RedisClientMock) Info() *StringCmd                                              { return nil }
func (r *RedisClientMock) LastSave() *IntCmd                                             { return nil }
func (r *RedisClientMock) Save() *StatusCmd                                              { return nil }
func (r *RedisClientMock) Shutdown() *StatusCmd                                          { return nil }
func (r *RedisClientMock) ShutdownSave() *StatusCmd                                      { return nil }
func (r *RedisClientMock) ShutdownNoSave() *StatusCmd                                    { return nil }
func (r *RedisClientMock) SlaveOf(host, port string) *StatusCmd                          { return nil }
func (r *RedisClientMock) SlowLog()                                                      { }
func (r *RedisClientMock) Sync()                                                         { }
func (r *RedisClientMock) Time() *StringSliceCmd                                         { return nil }
func (r *RedisClientMock) Eval(script string, keys []string, args []string) *Cmd         { return nil }
func (r *RedisClientMock) EvalSha(sha1 string, keys []string, args []string) *Cmd        { return nil }
func (r *RedisClientMock) ScriptExists(scripts ...string) *BoolSliceCmd                  { return nil }
func (r *RedisClientMock) ScriptFlush() *StatusCmd                                       { return nil }
func (r *RedisClientMock) ScriptKill() *StatusCmd                                        { return nil }
func (r *RedisClientMock) ScriptLoad(script string) *StringCmd                           { return nil }
func (r *RedisClientMock) DebugObject(key string) *StringCmd                             { return nil }
func (r *RedisClientMock) PubSubChannels(pattern string) *StringSliceCmd                 { return nil }
func (r *RedisClientMock) PubSubNumSub(channels ...string) *StringIntMapCmd              { return nil }
func (r *RedisClientMock) PubSubNumPat() *IntCmd                                         { return nil }
func (r *RedisClientMock) ClusterSlots() *ClusterSlotCmd                                 { return nil }
func (r *RedisClientMock) ClusterNodes() *StringCmd                                      { return nil }
func (r *RedisClientMock) ClusterMeet(host, port string) *StatusCmd                      { return nil }
func (r *RedisClientMock) ClusterReplicate(nodeID string) *StatusCmd                     { return nil }
func (r *RedisClientMock) ClusterInfo() *StringCmd                                       { return nil }
func (r *RedisClientMock) ClusterFailover() *StatusCmd                                   { return nil }
func (r *RedisClientMock) ClusterAddSlots(slots ...int) *StatusCmd                       { return nil }
func (r *RedisClientMock) ClusterAddSlotsRange(min, max int) *StatusCmd                  { return nil }
