package service

/**
 * 消息体
 * 示例{"type":1, "money":100.00, "toToken":"12312", "fromToken":"2131"}
 */
type Message struct {
	Type      int     `json:"type"`
	Money     float32 `json:"money"`
	ToToken   string  `json:"toToken"`
	FromToken string  `json:"fromToken"`
}
