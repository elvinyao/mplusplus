package models

type ServerConfig struct {
	Port              int                      `yaml:"port"`
	LogLevel          string                   `yaml:"log_level"`
	WebSocketPath     string                   `yaml:"websocket_path"`
	ConfluenceConfig  []string                 `yaml:"confluence_configs"`
	WebSocketHandlers []WebSocketHandlerConfig `yaml:"websocket_handlers"`
}

type WebSocketHandlerConfig struct {
	Action   string   `yaml:"action"`
	Handlers []string `yaml:"handlers"`
}

type ConfluenceConfig struct {
	ID            string        `yaml:"id"`
	HandlerID     string        `yaml:"handler_id"`
	URL           string        `yaml:"url"`
	CacheDuration int           `yaml:"cache_duration"`
	Tables        []TableConfig `yaml:"tables"`
}

type TableConfig struct {
	Name    string   `yaml:"name"`
	Columns []string `yaml:"columns"`
}
