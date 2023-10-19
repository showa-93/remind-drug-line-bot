package bot

import (
	"context"
	"net"
	"net/http"
	"time"
)

type (
	requestIDKey struct{}
	requestKey   struct{}
	responseKey  struct{}
)

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

func GetRequestID(ctx context.Context) string {
	value := ctx.Value(requestIDKey{})
	if value == nil {
		return ""
	}
	return value.(string)
}

type RequestRecord struct {
	Host      string
	URI       string
	Method    string
	RemoteIP  string
	UserAgent string
	Referrer  string
	Time      time.Time
}

func SetRequestRecord(ctx context.Context, r *http.Request) context.Context {

	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		Warn(ctx, err.Error())
		remoteIP = ""
	}
	rr := RequestRecord{
		Host:      r.URL.Host,
		URI:       r.RequestURI,
		Method:    r.Method,
		RemoteIP:  remoteIP,
		UserAgent: r.UserAgent(),
		Referrer:  r.Header.Get("Referer"),
		Time:      time.Now(),
	}
	return context.WithValue(ctx, requestKey{}, rr)
}

func GetRequestRecord(ctx context.Context) RequestRecord {
	value := ctx.Value(requestKey{})
	if value == nil {
		return RequestRecord{}
	}
	return value.(RequestRecord)
}

type ResponseRecord struct {
	StatusCode int
	Time       time.Time
}

func SetResponseRecord(ctx context.Context, statusCode int) context.Context {
	rs := ResponseRecord{
		StatusCode: statusCode,
		Time:       time.Now(),
	}
	return context.WithValue(ctx, responseKey{}, rs)
}

func GetResponseRecord(ctx context.Context) ResponseRecord {
	value := ctx.Value(responseKey{})
	if value == nil {
		return ResponseRecord{}
	}
	return value.(ResponseRecord)
}
