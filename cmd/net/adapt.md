net/http 已经适配完成的 api:
- srv.ListenAndServe -> h.Spin
- http.NewServeMux -> server.Default
- func(w ResponseWriter, r *Request) -> func(ctx context.Context, c *app.RequestContext)
```go
// net/http
http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
})

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}

http.HandleFunc("/ping", ping)

// hertz
server.Default().Any("/ping", func(ctx context.Context, c *app.RequestContext) {
    c.Response.Write([]byte("pong"))
})

func ping(ctx context.Context, c *app.RequestContext) {
    c.Response.Write([]byte("pong"))
}

server.Default().Any("/ping", ping)
```
- req.Header -> c.Request.Header
- file, fileHeader, err := r.FormFile("s") -> fileHeader, err := c.Request.FormFile("s")
- req.Header.Get/Del -> c.Request.Header.Get/Del
- req.Header.Set/Del -> c.Request.Header.Set/Del
- req.Host -> string(c.Request.Host)
- req.Method -> string(c.Request.Method)
- req.RequestURI -> c.Request.URI().String()
- req.URL.String() -> c.URL.String()
- req.URL.Query().Get -> c.Query
- req.Form.Get() -> string(c.FormValue())
- form := req.MultipartForm -> form, err := c.MultipartForm()
- form := req.MultipartForm.Value/File ->  _form, err := c.MultiPartForm() form := _form.Value/File
```go
// net/http
func ping(w http.ResponseWriter, r *http.Request) {
    file := r.MultipartForm.Value
}

// hertz
func ping(ctx context.Context, c *app.RequestContext) {
    _form, err := r.MultipartForm
	form := _form.Value
}
```
- http.HandlerFunc-> h.Any
- resp.Header() -> c.Response.Header
- resp.Header().Get/Del -> c.Response.Header.Get/Del
- resp.Header().Set/Del -> c.Response.Header.Set/Del
- resp.Write -> c.Response.Write
- 从 http.Server struct 中获取的字段(~~fmt.Sprintf() 等函数组成的属性尚未取到~~ 已支持) -> hertz server option 中的配置函数
- resp.WriteHeader -> c.SetStatusCode
- http.Error -> c.AbortWithMsg/c.AbortWithStatus

chi 已经适配完成的 api:
- chi.NewRouter -> server.Default
```go
// chi
r := chi.NewRouter()

func newRouter() *chi.Mux {
    r := chi.NewRouter()
	//...
	return r
}
// hertz
r := server.Default()

func newRouter() *server.Hertz {
    r := server.Default()
    //...
    return r
}
```

- r.Get/Post/Put/Delete/Options/Head/Patch -> h.GET/POST/PUT/DELETE/OPTIONS/HEAD/PATCH
```go
// chi  
r := chi.NewRouter()
r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
})

// hertz
r := server.Default()
r.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
    c.String(200, "pong")
})
```
- r.Method -> h.Handle
```go
// chi
r := chi.NewRouter()
r.Method("GET", "/ping", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
})
// hertz
r := server.Default()
r.Handle("GET", "/ping", func(ctx context.Context, c *app.RequestContext) {
   c.Response.Write([]byte("pong"))
})
```
