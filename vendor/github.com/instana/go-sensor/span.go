// (c) Copyright IBM Corp. 2021
// (c) Copyright Instana Inc. 2016

package instana

import (
	"bytes"
	"strings"
	"sync"
	"time"

	ot "github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
)

type spanS struct {
	Service     string
	Operation   string
	Start       time.Time
	Duration    time.Duration
	Correlation EUMCorrelationData
	Tags        ot.Tags
	Logs        []ot.LogRecord
	ErrorCount  int

	tracer *tracerS
	mu     sync.Mutex

	context SpanContext
}

func (r *spanS) BaggageItem(key string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.context.Baggage[key]
}

func (r *spanS) SetBaggageItem(key, val string) ot.Span {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.context = r.context.WithBaggageItem(key, val)

	return r
}

func (r *spanS) Context() ot.SpanContext {
	return r.context
}

func (r *spanS) Finish() {
	r.FinishWithOptions(ot.FinishOptions{})
}

func (r *spanS) FinishWithOptions(opts ot.FinishOptions) {
	finishTime := opts.FinishTime
	if finishTime.IsZero() {
		finishTime = time.Now()
	}

	duration := finishTime.Sub(r.Start)

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, lr := range opts.LogRecords {
		r.appendLog(lr)
	}

	for _, ld := range opts.BulkLogData {
		r.appendLog(ld.ToLogRecord())
	}

	r.Duration = duration
	if !r.context.Suppressed {
		r.tracer.recorder.RecordSpan(r)
		r.sendOpenTracingLogRecords()
	}
}

func (r *spanS) appendLog(lr ot.LogRecord) {
	maxLogs := r.tracer.Options().MaxLogsPerSpan
	if maxLogs == 0 || len(r.Logs) < maxLogs {
		r.Logs = append(r.Logs, lr)
	}
}

func (r *spanS) Log(ld ot.LogData) {
	if r.tracer.Options().DropAllLogs {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if ld.Timestamp.IsZero() {
		ld.Timestamp = time.Now()
	}

	r.appendLog(ld.ToLogRecord())
}

func (r *spanS) LogEvent(event string) {
	r.Log(ot.LogData{
		Event: event})
}

func (r *spanS) LogEventWithPayload(event string, payload interface{}) {
	r.Log(ot.LogData{
		Event:   event,
		Payload: payload})
}

func (r *spanS) LogFields(fields ...otlog.Field) {

	for _, v := range fields {
		// If this tag indicates an error, increase the error count
		if isErrorLogField(v) {
			r.ErrorCount++
		}
	}

	lr := ot.LogRecord{
		Fields: fields,
	}

	if r.tracer.Options().DropAllLogs {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if lr.Timestamp.IsZero() {
		lr.Timestamp = time.Now()
	}

	r.appendLog(lr)
}

func (r *spanS) LogKV(keyValues ...interface{}) {
	fields, err := otlog.InterleavedKVToFields(keyValues...)
	if err != nil {
		r.LogFields(otlog.Error(err), otlog.String("function", "LogKV"))

		return
	}

	r.LogFields(fields...)
}

func (r *spanS) SetOperationName(operationName string) ot.Span {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Operation = operationName

	return r
}

func (r *spanS) SetTag(key string, value interface{}) ot.Span {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Tags == nil {
		r.Tags = ot.Tags{}
	}

	// If this tag indicates an error, increase the error count
	if key == "error" {
		r.ErrorCount++
	}

	if key == suppressTracingTag {
		r.context.Suppressed = true
		return r
	}

	r.Tags[key] = value

	return r
}

func (r *spanS) Tracer() ot.Tracer {
	return r.tracer
}

// sendOpenTracingLogRecords converts OpenTracing log records that contain errors
// to Instana log spans and sends them to the agent
func (r *spanS) sendOpenTracingLogRecords() {
	for _, lr := range r.Logs {
		for _, lf := range lr.Fields {
			if isErrorLogField(lf) {
				r.sendOpenTracingLogRecord(lr)
				break
			}
		}
	}
}

func (r *spanS) sendOpenTracingLogRecord(lr ot.LogRecord) {
	buf := bytes.NewBuffer(nil)

	enc := newOpenTracingLogEncoder(buf)
	for _, lf := range lr.Fields {
		lf.Marshal(enc)
		buf.WriteByte(' ')
	}

	r.tracer.StartSpan(
		"log.go",
		ot.ChildOf(r.context),
		ot.StartTime(lr.Timestamp),
		ot.Tags{
			"log.level":   "ERROR",
			"log.message": strings.TrimSpace(buf.String()),
		},
	).FinishWithOptions(
		ot.FinishOptions{
			FinishTime: lr.Timestamp,
		},
	)
}

func isErrorLogField(lf otlog.Field) bool {
	return lf.Key() == "error" || lf.Key() == "error.object"
}
