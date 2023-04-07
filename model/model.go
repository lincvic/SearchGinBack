package model

import "sync"

type RequestData struct {
	Keyword string `json:"keyword"`
}

type OAIData struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Index     int       `json:"index"`
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

type QueryPayload struct {
	TopK        int       `json:"top"`
	Vector      []float64 `json:"vector"`
	Space       string    `json:"space"`
	WithPayLoad bool      `json:"with_payload"`
}

type QdrantResp struct {
	Result []struct {
		Version int     `json:"version"`
		Score   float64 `json:"score"`
		Payload struct {
			Answer   string `json:"answer"`
			Question string `json:"question"`
		} `json:"payload"`
		Vector any `json:"vector"`
	} `json:"result"`
	Status string  `json:"status"`
	Time   float64 `json:"time"`
}

type ServerRespData struct {
	sync.RWMutex
	CognitiveSearchResp string `json:"cognitive_search_resp"`
	QdrantResp          string `json:"qdrant_resp"`
}

func (s *ServerRespData) SetCognitiveSearchRespWithLock(resp string) {
	s.Lock()
	defer s.Unlock()
	s.CognitiveSearchResp = resp
}

func (s *ServerRespData) GetCognitiveSearchRespWithLock() string {
	s.RLock()
	defer s.RUnlock()
	return s.CognitiveSearchResp
}

func (s *ServerRespData) SetQdrantRespWithLock(resp string) {
	s.Lock()
	defer s.Unlock()
	s.QdrantResp = resp
}

func (s *ServerRespData) GetQdrantRespWithLock() string {
	s.RLock()
	defer s.RUnlock()
	return s.QdrantResp
}
