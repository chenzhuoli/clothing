/**
 * 主程后端与推荐模块的thrift接口和消息格式
 */

namespace go recommend

struct RecommendStruct {
  1: i32 key
  2: string value
}

service RecommendService {
  RecommendStruct getStruct(1: i32 key)
}
