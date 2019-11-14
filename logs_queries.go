package datadog

import "encoding/json"

const (
	logsQueriesPath = "/v1/logs-queries/list"
)

type LogsTime struct {
	From     *json.Number `json:"from"`
	To       *json.Number `json:"to"`
	Timezone *string      `json:"timezone,omitempty"`
	Offset   *string      `json:"offset,omitempty"`
}

type LogsQueries struct {
	Query   *string      `json:"query"`
	Time    *LogsTime    `json:"time"`
	StartAt *json.Number `json:"startAt,omitempty"`
	Sort    *string      `json:"sort,omitempty"`
	Limit   *json.Number `json:"limit,omitempty"`
	Index   *string      `json:"index,omitempty"`
}

type Logs struct {
	Logs      *[]Log  `json:"logs"`
	NextLogId *string `json:"nextLogId"`
	Status    *string `json:"status"`
	RequestId *string `json:"requestId"`
}

type Log struct {
	Id      *string  `json:"id"`
	Content *Content `json:"content"`
	//	Content *json.RawMessage
}

type Content struct {
	Timestamp  *string          `json:"timestamp"`
	Tags       *[]string        `json:"tags"`
	Attributes *json.RawMessage `json:"attributes"`
	Host       *string          `json:"host"`
	Service    *string          `json:"service"`
	Message    *string          `json:"message"`
}

func (client *Client) SearchLogsQueries(logsQueries *LogsQueries) (*Logs, error) {
	var out Logs
	if err := client.doJsonRequest("POST", logsQueriesPath, logsQueries, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
