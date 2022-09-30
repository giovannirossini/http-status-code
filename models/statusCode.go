package models

type httpStatus struct {
	Id          string        `json:"id"`
	Description string        `json:"description"`
	Codes       []statusCodes `json:"codes"`
}

type statusCodes struct {
	Id          string    `json:"id"`
	Description string `json:"description"`
}

var code1xx = []statusCodes{
	{Id: "100", Description: "Continue"},
	{Id: "101", Description: "Switching Protocols"},
	{Id: "102", Description: "Processing"},
}

var code2xx = []statusCodes{
	{Id: "201", Description: "Created"},
	{Id: "202", Description: "Accepted"},
	{Id: "203", Description: "Non-authoritative Information"},
	{Id: "204", Description: "No Content"},
	{Id: "205", Description: "Reset Content"},
	{Id: "206", Description: "Partial Content"},
	{Id: "207", Description: "Multi-Status"},
	{Id: "208", Description: "Already Reported"},
	{Id: "226", Description: "IM Used"},
}

var code3xx = []statusCodes{
	{Id: "300", Description: "Multiple Choices"},
	{Id: "301", Description: "Moved Permanently"},
	{Id: "302", Description: "Found"},
	{Id: "303", Description: "See Other"},
	{Id: "304", Description: "Not Modified"},
	{Id: "305", Description: "Use Proxy"},
	{Id: "307", Description: "Temporary Redirect"},
	{Id: "308", Description: "Permanent Redirect"},
}

var code4xx = []statusCodes{
	{Id: "400", Description: "Bad Request"},
	{Id: "401", Description: "Unauthorized"},
	{Id: "402", Description: "Payment Required"},
	{Id: "403", Description: "Forbidden"},
	{Id: "404", Description: "Not Found"},
	{Id: "405", Description: "Method Not Allowed"},
	{Id: "406", Description: "Not Acceptable"},
	{Id: "407", Description: "Proxy Authentication Required"},
	{Id: "408", Description: "Request Timeout"},
	{Id: "409", Description: "Conflict"},
	{Id: "410", Description: "Gone"},
	{Id: "411", Description: "Length Required"},
	{Id: "412", Description: "Precondition Failed"},
	{Id: "413", Description: "Payload Too Large"},
	{Id: "414", Description: "Request-URI Too Long"},
	{Id: "415", Description: "Unsupported Media Type"},
	{Id: "416", Description: "Requested Range Not Satisfiable"},
	{Id: "417", Description: "Expectation Failed"},
	{Id: "418", Description: "I’m a teapot"},
	{Id: "421", Description: "Misdirected Request"},
	{Id: "422", Description: "Unprocessable Entity"},
	{Id: "423", Description: "Locked"},
	{Id: "424", Description: "Failed Dependency"},
	{Id: "426", Description: "Upgrade Required"},
	{Id: "428", Description: "Precondition Required"},
	{Id: "429", Description: "Too Many Requests"},
	{Id: "431", Description: "Request Header Fields Too Large"},
	{Id: "444", Description: "Connection Closed Without Response"},
	{Id: "451", Description: "Unavailable For Legal Reasons"},
	{Id: "499", Description: "Client Closed Request"},
}

var code5xx = []statusCodes{
	{Id: "500", Description: "Internal Server Error"},
	{Id: "501", Description: "Not Implemented"},
	{Id: "502", Description: "Bad Gateway"},
	{Id: "503", Description: "Service Unavailable"},
	{Id: "504", Description: "Gateway Timeout"},
	{Id: "505", Description: "HTTP Version Not Supported"},
	{Id: "506", Description: "Variant Also Negotiates"},
	{Id: "507", Description: "Insufficient Storage"},
	{Id: "508", Description: "Loop Detected"},
	{Id: "510", Description: "Not Extended"},
	{Id: "511", Description: "Network Authentication Required"},
	{Id: "599", Description: "Network Connect Timeout Error"},
}

var HttpStatusCodes = []httpStatus{
	{Id: "1xx", Description: "Informational", Codes: code1xx},
	{Id: "2xx", Description: "Success", Codes: code2xx},
	{Id: "3xx", Description: "Redirection", Codes: code3xx},
	{Id: "4xx", Description: "Client Error", Codes: code4xx},
	{Id: "5xx", Description: "Server Error", Codes: code5xx},
}
