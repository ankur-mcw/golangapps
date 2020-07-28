package server

import (
    "encoding/json"
    "fmt"
    "github.com/buaazp/fasthttprouter"
    "github.com/ankur-mcw/golangapps/services/labelGeneration/processor"
    "github.com/narvar/NarvarGolangApps/nlog"
    "github.com/narvar/NarvarGolangApps/ntime"
    "github.com/valyala/fasthttp"
)

type errorResponse struct {
    Message string `json:"message"`
}

func health(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("up")
}

func generateLabel(ctx *fasthttp.RequestCtx) {
    startupSeqBeginTime := ntime.NtimeNow()
    tracerID := string(ctx.Request.Header.Peek("X-Narvar-Tracer-ID"))
    nlog.Debugf("tracerID=%s", tracerID)

    mockResponse := string(ctx.Request.Header.Peek("X-Mock-Response"))
    nlog.Debugf("mockResponse=%s", mockResponse)

    experimental := string(ctx.Request.Header.Peek("X-Narvar-Experimental"))
    nlog.Debugf("experimental=%s", experimental)
    
    var reqBody processor.LabelGenerationRequest
    err := json.Unmarshal(ctx.PostBody(), &reqBody)
    if err != nil {
        ctx.SetStatusCode(400)
    } else {
        response, err := processor.GenerateLabel(reqBody)
        ctx.SetContentType("application/json")
        if err != nil {
            ctx.SetStatusCode(400)
            json.NewEncoder(ctx).Encode(errorResponse{err.Error()})
        } else {
            ctx.SetStatusCode(201)
            json.NewEncoder(ctx).Encode(response)
        }
    }
    nlog.Debugf("Responded in %v", ntime.Since(startupSeqBeginTime))
}

// Start the server and registers endpoints
func Start(port int) error {
	router := fasthttprouter.New()
	router.GET("/health", health)
	router.POST("/api/v1/carrier/label", generateLabel)

	nlog.Infof("Server running on port %d", port)
    return fasthttp.ListenAndServe(fmt.Sprint(":", port), router.Handler)
}
