package excutor

import "encoding/json"

type Selector struct {
	Using string `json:"using"`
	Value string `json:"value"`
}

const (
	CSS_SELECTOR               = "css selector"
	XPATH_SELECTOR             = "xpath"
	LINK_TEXT_SELECTOR         = "link text"
	PARTIAL_LINK_TEXT_SELECTOR = "partial link text"
	TAG_NAME_SELECTOR          = "Tag name"
)

func CreateSelector(using, value string) Selector {
	return Selector{Using: using, Value: value}
}

func (s *Selector) ToString() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}
