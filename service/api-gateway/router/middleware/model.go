package middleware

type HTTPLog struct {
	Path      string                 `json:"path"`
	Method    string                 `json:"method"`
	Status    int                    `json:"status"`
	Latency   int64                  `json:"latency_ms"`
	UserID    int64                  `json:"user_id,omitempty"`
	Ts        int64                  `json:"ts"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Body      interface{}            `json:"body,omitempty"`
	Timestamp string                 `json:"@timestamp"`
}
