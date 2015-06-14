package raven

// http://sentry.readthedocs.org/en/latest/developer/interfaces/index.html#sentry.interfaces.Message
type Message struct {
	// Required
	Message string `json:"message"`

	// Optional
	Params []interface{} `json:"params,omitempty"`
}

func (m *Message) Class() string { return "sentry.interfaces.Message" }

// http://sentry.readthedocs.org/en/latest/developer/interfaces/index.html#sentry.interfaces.Template
type Template struct {
	// Required
	Filename    string `json:"filename"`
	Lineno      int    `json:"lineno"`
	ContextLine string `json:"context_line"`

	// Optional
	PreContext   []string `json:"pre_context,omitempty"`
	PostContext  []string `json:"post_context,omitempty"`
	AbsolutePath string   `json:"abs_path,omitempty"`
}

func (t *Template) Class() string { return "sentry.interfaces.Template" }

// http://sentry.readthedocs.org/en/latest/developer/interfaces/index.html#sentry.interfaces.User
type User struct {
	ID       string `json:"id"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (h *User) Class() string { return "sentry.interfaces.User" }

// http://sentry.readthedocs.org/en/latest/developer/interfaces/index.html#sentry.interfaces.Query
type Query struct {
	// Required
	Query string `json:"query"`

	// Optional
	Engine string `json:"engine,omitempty"`
}

func (q *Query) Class() string { return "sentry.interfaces.Query" }

type Log struct {
	CommitID string      `json:"commit_id"`
	Version  string      `json:"version"`
	Assets   []string    `json:"assets"`
	Entries  []*LogEntry `json:"entries"`
}

func (l *Log) Class() { return "sentry_log.Log" }

type LogEntry struct {
	Date       int64       `json:"date"`
	Type       string      `json:"type"`
	Message    string      `json:"message"`
	Objects    interface{} `json:"objects"`
	Stacktrace []*LogFrame `json:"stacktrace"`
}

type LogFrame struct {
	Filename    string   `json:"filename"`
	Name        string   `json:"name"`
	LineNo      int      `json:"line_no"`
	ColNo       int      `json:"col_no"`
	InApp       bool     `json:"in_app"`
	ContextPre  []string `json:"context_pre,omitempty"`
	ContextLine string   `json:"context_line,omitempty"`
	ContextPost []string `json:"context_post,omitempty"`
	AbsPath     string   `json:"abs_path"`
	StartLineNo int      `json:"start_line_no"`
}
