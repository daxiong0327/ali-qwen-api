package image

// VLMessage 定义的图像分析消息
type VLMessage struct {
	Role    string       `json:"role"`
	Content []*VLContent `json:"content"`
}

type VLContent struct {
	Image string `json:"image,omitempty"`
	Text  string `json:"text"`
	Type  string `json:"type"`
}
