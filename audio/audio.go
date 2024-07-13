package audio

type AudioMessage struct {
	Role    string      `json:"role"`
	Content []*AContent `json:"content"`
}

type AContent struct {
	Audio string `json:"audio"` // 语音URL
	Text  string `json:"text"`
	Type  string `json:"type"` // text, url
}
