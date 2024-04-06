package backendapi

import (
	"cockroachai/config"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	group := s.Group("/")
	group.ALL("/backend-api/*any", ProxyBackendApi)
	group.GET("/backend-api/prompt_library/", PromptLibrary)
}

func Init(ctx g.Ctx) {
	g.Log().Info(ctx, "BackendApi module initialized")
}

// /backend-api/prompt_library/
func PromptLibrary(r *ghttp.Request) {
	ctx := r.Context()
	limit := r.Get("limit").String()
	offset := r.Get("offset").String()
	ProxyClient := gclient.New().Proxy(config.Ja3Proxy.String()).SetBrowserMode(true).SetHeaderMap(g.MapStrStr{
		"Origin":     "https://chat.openai.com",
		"Referer":    "https://chat.openai.com/",
		"Host":       "chat.openai.com",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
	})
	res, err := ProxyClient.Get(ctx, "https://chat.openai.com/backend-api/prompt_library/", g.Map{"limit": limit, "offset": offset})
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.WriteStatus(500, "")
		return
	}
	defer res.Close()
	// res.RawDump()
	r.Response.Status = res.StatusCode
	r.Response.Write(res.ReadAll())
}