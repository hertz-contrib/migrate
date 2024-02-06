net/http 已经适配完成的 api:
- srv.ListenAndServe -> h.Spin
- http.NewServeMux -> server.Default
- http.Redirect -> c.Redirect
```go
// net/http
func ping(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/pong", 301)
}

// hertz
func ping(ctx context.Context, c *app.RequestContext) {
    c.Redirect(301, []byte("/pong"))
}
```
- http.NotFound -> c.NotFound
- func(w ResponseWriter, r *Request) -> func(ctx context.Context, c *app.RequestContext)
```go
// net/http
http.HandleFunc("/ping", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
})

func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}

http.HandleFunc("/ping", ping)

// hertz
server.Default().Any("/ping", func (ctx context.Context, c *app.RequestContext) {
    c.Response.Write([]byte("pong"))
})

func ping(ctx context.Context, c *app.RequestContext) {
    c.Response.Write([]byte("pong"))
}

server.Default().Any("/ping", ping)

// 高级场景 
// net/http

func ping(w http.ResponseWriter, r *http.Request, data any) {
    w.Header().Set("Content-Type", "application/json")
    b, _ := json.Marshal(data)
    w.Write(b)
}

func ping2(w http.ResponseWriter, r *http.Request, data any) error{
    w.Header().Set("Content-Type", "application/json")
    b, _ := json.Marshal(data)
    _, err := w.Write(b)
	return err
}

ping(w, r, map[string]string{"msg": "pong"})
if err := ping2(w, r, map[string]string{"msg": "pong"}); err != nil {
    log.Println(err)
}

// hertz
func ping(c *app.RequestContext, data any) {
    c.Response.Header.Set("Content-Type", "application/json")
    b, _ := json.Marshal(data)
    c.SetStatusCode(200)
    c.Response.SetBody(b)
}

func ping2(c *app.RequestContext, data any) error {
    w.Header().Set("Content-Type", "application/json")
    b, _ := json.Marshal(data)
    c.SetStatusCode(200)
    c.Response.SetBody(b)
}

ping(c, map[string]string{"msg": "pong"})
if err := ping2(c, map[string]string{"msg": "pong"}); err != nil {
    log.Println(err)
}

```
- req.Header -> c.Request.Header
- file, fileHeader, err := r.FormFile("s") -> fileHeader, err := c.Request.FormFile("s")
- req.Header.Get/Del -> c.Request.Header.Get/Del
- req.Header.Set/Del -> c.Request.Header.Set/Del
- req.Host -> string(c.Request.Host())
```go
// net/http
func ping(w http.ResponseWriter, r *http.Request) {
    host := r.Host
	if r.Host == "localhost:8080" {
		
    }
}

// hertz
func ping(ctx context.Context, c *app.RequestContext) {
    host := string(c.Host())
    if string(c.Host()) == "localhost:8080" {
        
    }
}
```

- req.Method -> string(c.Method())
```go
// net/http
func ping(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        w.Write([]byte("pong"))
    }
	
	switch r.Method {
  
    }
	
	method := r.Method
}

// hertz
func ping(ctx context.Context, c *app.RequestContext) {
    if string(c.Method()) == "GET" {
        c.Response.Write([]byte("pong"))
    }
	switch string(c.Method()) {
  
    }
	method := string(c.Method())
}
```
- req.RequestURI -> c.Request.URI().String()
```go
// net/http
func ping(w http.ResponseWriter, r *http.Request) {
    uri := r.RequestURI
	if r.RequestURI == "/ping" {
        
    }
}

// hertz
func ping(ctx context.Context, c *app.RequestContext) {
    uri := c.Request.URI().String()
	if string(c.Request.RequestURI()) == "/ping" {
        
    }
}
```
- req.URL.String() -> c.URL.String()
- req.URL.Query().Get -> c.Query
- req.Cookie -> c.Cookie
```go
// net/http
func ping(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("s")
	if err != nil {}
	s := cookie.Value
	if cookie.Value == "s" {}
}

// hertz
func ping(ctx context.Context, c *app.RequestContext) {
    cookie, err := c.Cookie("s")
    if err != nil {}
    s := string(cookie)
    if string(cookie) == "s" {}
}
```
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
- req.FormValue -> string(c.FormValue)
- http.HandlerFunc-> h.Any
- resp.Header() -> c.Response.Header
- resp.Header().Get/Del -> c.Response.Header.Get/Del
- resp.Header().Set/Del -> c.Response.Header.Set/Del
- fmt.Fprintf(w, s) -> c.SetBodyString(s)
```go
// net/http
func ping(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "pong")
}

// 在 存在 w.WriteHeader() 的情况下, 不会自动设置状态码
func ping(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(400)
    w.Write([]byte("pong"))
}

// hertz
func ping(ctx context.Context, c *app.RequestContext) {
	c.SetStatusCode(200)
    c.SetBodyString("pong")
}

func ping(ctx context.Context, c *app.RequestContext) {
    c.SetStatusCode(400)
    c.Response.Write([]byte("pong"))
}
```
- resp.Write -> c.Response.Write
```go
// net/http
func ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
}
// 在 存在 w.WriteHeader() 的情况下, 不会自动设置状态码
func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
    w.Write([]byte("pong"))
}

// hertz
func ping(ctx context.Context, c *app.RequestContext) {
	c.SetStatusCode(200)
    c.Response.Write([]byte("pong"))
}

func ping(ctx context.Context, c *app.RequestContext) {
    c.SetStatusCode(400)
    c.Response.Write([]byte("pong"))
}
```
- 从 http.Server struct 中获取的字段(~~fmt.Sprintf() 等函数组成的属性尚未取到~~ 已支持) -> hertz server option 中的配置函数
```go
// net/http
srv := &http.Server{
	// 类似填入 cfg.Addr 的复杂度太高, 无法直接适配
    Addr:    ":8080", // 现可填入 fmt.Sprintf(":8080") 
	ReadTimeout: 10 * time.Second,
}

// hertz
srv := server.Default(
    server.WithHostPorts(":8080"),
    server.WithReadTimeout(10),
)
```
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
r.Get("/ping/{id}", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
})

// hertz
r := server.Default()
r.GET("/ping/:id", func(ctx context.Context, c *app.RequestContext) {
    c.String(200, "pong")
})
```
- r.Method -> h.Handle
```go
// chi
r := chi.NewRouter()
r.Method("GET", "/ping/{id}", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
})
// hertz
r := server.Default()
r.Handle("GET", "/ping/{id}", func(ctx context.Context, c *app.RequestContext) {
   c.Response.Write([]byte("pong"))
})
```
