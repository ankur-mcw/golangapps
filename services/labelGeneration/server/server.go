package server

import (
    "encoding/json"
	"log"
	"github.com/buaazp/fasthttprouter"
    "github.com/valyala/fasthttp"
    "../processor"
)

type errorResponse struct {
    Message string `json:"message"`
}

func health(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("up")
}

func generateLabel(ctx *fasthttp.RequestCtx) {

    tracerID := string(ctx.Request.Header.Peek("X-Narvar-Tracer-ID"))
    log.Printf("tracerID=%s", tracerID)

    mockResponse := string(ctx.Request.Header.Peek("X-Mock-Response"))
    log.Printf("mockResponse=%s", mockResponse)

    experimental := string(ctx.Request.Header.Peek("X-Narvar-Experimental"))
    log.Printf("experimental=%s", experimental)
    
    var reqBody processor.LabelGenerationRequest
    err := json.Unmarshal(ctx.PostBody(), &reqBody)
    if err != nil {
        ctx.SetStatusCode(400)
    } else {
        response, err := processor.GenerateLabel(reqBody)
        ctx.SetContentType("application/json")
        if err != nil {
            // fmt.Printf("error: %s", err)
            ctx.SetStatusCode(400)
            // json.NewEncoder(ctx).Encode(errorResponse{err})
        } else {
            ctx.SetStatusCode(201)
            json.NewEncoder(ctx).Encode(response)
        }
    }
}

// Start starts the server and registers endppints
func Start(port string) error {
	router := fasthttprouter.New()
	router.GET("/health", health)
	router.POST("/api/v1//carrier/label", generateLabel)

	log.Printf("Server running on port %s", port)
    return fasthttp.ListenAndServe(":" + port, router.Handler)
}
