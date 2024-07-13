package chat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxySendRequest(t *testing.T) {
	args := &RequestArgs{
		BaseUrl: "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
		ApiKey:  "sk-2b7jhk870ac98jd949sf1dec7ef9kdjsh", // 这里填入自己的key
		Request: &Request{
			Model: "qwen-turbo", // 这里填入自己想要的模型版本
			Input: &Input{
				Messages: []*Message{
					{Role: "user", Content: "你好"},
				},
			},
		},
	}
	res, err := SendRequest(args)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, res)
	assert.NotEqual(t, "", res.RequestId)
	fmt.Println(res.RequestId)
	assert.NotEqual(t, nil, res.Usage)
	fmt.Println(res.Usage)
	assert.NotEqual(t, nil, res.Output)
	fmt.Println(res.Output)
}
