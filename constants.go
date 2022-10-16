package log

const (
	GcpJsonLoggingJsonPayloadKey = "jsonPayload"
	GcpJsonLoggingRawMessageKey  = "rawMessage"
	GcpJsonLoggingMessageKey     = "message"
	GcpJsonLoggingTraceKey       = "logging.googleapis.com/trace"
	GcpJsonLoggingSpanIdKey      = "logging.googleapis.com/spanId"
)

const (
	HttpRequestHeaderTraceKey  = "trace"
	HttpRequestHeaderSpanIdKey = "spanId"
)
