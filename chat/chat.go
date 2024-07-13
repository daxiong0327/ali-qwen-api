package chat

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"io"
)

// Message 定义的聊天消息
type Message struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	ToolCalls any    `json:"tool_calls,omitempty"`
}

type Input struct {
	Messages []*Message `json:"messages"`
}

type Parameters struct {
	ResultFormat string `json:"result_format"`
}
type Request struct {
	Model             string      `json:"model"`
	Input             *Input      `json:"input"`
	Parameters        *Parameters `json:"parameters,omitempty"`
	Seed              int32       `json:"seed,omitempty"`               //生成时使用的随机数种子，用于控制模型生成内容的随机性。seed支持无符号64位整数。
	MaxTokens         int32       `json:"max_tokens,omitempty"`         // 生成的最大token数，qwen-turbo最大值和默认值为1500 tokens。qwen-max、qwen-max-1201、qwen-max-longcontext和qwen-plus模型，最大值和默认值均为2000 tokens。
	TopP              float32     `json:"top_p,omitempty"`              // 生成过程中的核采样方法概率阈值，例如，取值为0.8时，仅保留概率加起来大于等于0.8的最可能token的最小集合作为候选集。取值范围为（0,1.0)，取值越大，生成的随机性越高；取值越低，生成的确定性越高
	TopK              int32       `json:"top_k,omitempty"`              // 生成时，采样候选集的大小。例如，取值为50时，仅将单次生成中得分最高的50个token组成随机采样的候选集。取值越大，生成的随机性越高；取值越小，生成的确定性越高。取值为None或当top_k大于100时，表示不启用top_k策略，此时，仅有top_p策略生效。
	RepetitionPenalty float32     `json:"repetition_penalty,omitempty"` // 用于控制模型生成时连续序列中的重复度。提高repetition_penalty时可以降低模型生成的重复度，1.0表示不做惩罚。没有严格的取值范围。
	PresencePenalty   float32     `json:"presence_penalty,omitempty"`   // 用户控制模型生成时整个序列中的重复度。提高presence_penalty时可以降低模型生成的重复度，取值范围[-2.0, 2.0]。
	Temperature       float32     `json:"temperature,omitempty"`        // 用于控制模型回复的随机性和多样性。具体来说，temperature值控制了生成文本时对每个候选词的概率分布进行平滑的程度。较高的temperature值会降低概率分布的峰值，使得更多的低概率词被选择，生成结果更加多样化；而较低的temperature值则会增强概率分布的峰值，使得高概率词更容易被选择，生成结果更加确定。 取值范围：[0, 2)，不建议取值为0，无意义。
	Stop              []string    `json:"stop,omitempty"`               //
	Stream            bool        `json:"stream,omitempty"`             // 是否开启流式响应，默认为false
	EnableSearch      bool        `json:"enable_search,omitempty"`      // 用于控制模型在生成文本时是否使用互联网搜索结果进行参考。
	ResultFormat      string      `json:"result_format,omitempty"`      // 用于指定返回结果的格式，默认为text，也可选择message。当设置为message时，输出格式请参考返回结果。推荐您优先使用message格式。
	IncrementalOutput bool        `json:"incremental_output,omitempty"` // 用于指定是否在每次调用中增量输出，默认为false。
	Tools             []string    `json:"tools,omitempty"`              //
	ToolChoice        string      `json:"tool_choice,omitempty"`        //
}

type Choice struct {
	FinishReason string   `json:"finish_reason"`
	Message      *Message `json:"message"`
}

type Usage struct {
	InputTokens  int32 `json:"input_tokens"`
	OutputTokens int32 `json:"output_tokens"`
	TotalTokens  int32 `json:"total_tokens"`
}

type Output struct {
	Text         string    `json:"text"`
	FinishReason string    `json:"finish_reason"`
	Choices      []*Choice `json:"choices"`
}

type Response struct {
	StatusCode int     `json:"status_code"` // 200（HTTPStatus.OK）表示请求成功，否则表示请求失败，可以通过code获取错误码，通过message字段获取错误详细信息。
	RequestId  string  `json:"request_id"`  // 请求唯一标识id
	Code       string  `json:"code"`        //表示错误码，调用成功时为空值。仅适用于Python。
	Message    string  `json:"message"`     // 错误信息,仅适用于python
	Output     *Output `json:"output"`      // 结果信息
	Usage      *Usage  `json:"usage"`       // 消耗的token
}

type RequestArgs struct {
	*Request
	BaseUrl string `json:"base_url"`
	ApiKey  string `json:"api_key"`
}

func ProxySendRequest(args *RequestArgs) (*Response, error) {
	client := req.C()
	req3 := client.Post(args.BaseUrl)
	dataBytes, err := json.Marshal(args.Request)
	if err != nil {
		return nil, err
	}
	res := req3.SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", args.ApiKey),
		"Content-Type":  "application/json",
	}).SetBody(dataBytes).Do()

	// 反序列化结果数据
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resData *Response
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return nil, err
	}
	return resData, nil
}
