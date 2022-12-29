# FastHTTP -> Hertz conversion table

## [fasthttp.RequestCtx](https://pkg.go.dev/github.com/valyala/fasthttp#RequestCtx)

- All the pseudocode below assumes ctx have these types:

```Go
func handler(ctx *fasthttp.RequestCtx){..}
->
func handler(
    c context.Context
    ctx *app.RequestContext
){..}
```

### No Changed Function

- ctx.FormFile -> [ctx.FormFile](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.FormFile)

- ctx.FormValue -> [ctx.FormValue](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.FormValue)

- ctx.Hijack -> [ctx.Hijack](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Hijack)

- ctx.Hijacked -> [ctx.Hijacked](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Hijacked)

- ctx.Host -> [ctx.Host](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Host)

- ctx.IfModifiedSince -> [ctx.IfModifiedSince](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.IfModifiedSince)

- ctx.IsGet -> [ctx.IsGet](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.IsGet)

- ctx.IsHead -> [ctx.IsHead](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.IsHead)

- ctx.IsPost -> [ctx.IsPost](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.IsPost)

- ctx.Method -> [ctx.Method](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Method)

- ctx.MultipartForm -> [ctx.MultipartForm](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.MultipartForm)

- ctx.NotFound -> [ctx.NotFound](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.NotFound)

- ctx.NotModified -> [ctx.NotModified](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.NotModified)

- ctx.Path -> [ctx.Path](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Path)

- ctx.PostArgs -> [ctx.PostArgs](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.PostArgs)

- ctx.PostBody -> [ctx.Request.Body](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.Body)

- ctx.QueryArgs -> [ctx.QueryArgs](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.QueryArgs)

- ctx.RequestBodyStream -> [ctx.RequestBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.RequestBodyStream)

- ctx.RemoteAddr -> [ctx.RemoteAddr](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.RemoteAddr)

- ctx.SetBodyStream -> [ctx.SetBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetBodyStream)

- ctx.SetBodyString -> [ctx.SetBodyString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetBodyString)

- ctx.SetConnectionClose -> [ctx.SetConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetConnectionClose)

- ctx.SetContentType -> [ctx.SetContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetContentType)

- ctx.SetContentTypeBytes -> [ctx.SetContentTypeBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetContentTypeBytes)

- ctx.SetStatusCode -> [ctx.SetStatusCode](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetStatusCode)

- ctx.URI -> [ctx.URI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.URI)

- ctx.UserAgent -> [ctx.UserAgent](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.UserAgent)

- ctx.Write -> [ctx.Write](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Write)

- ctx.WriteString -> [ctx.WriteString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.WriteString)

### context.Context Interface

- ctx.Deadline -> c.Deadline

- ctx.Done -> c.Done

- ctx.Err -> c.Err

- ctx.Value -> c.Value

### Need Change Function

- ctx.Conn -> ctx.[GetConn()](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetConn)

- ctx.Error -> [ctx.AboutWithMsg](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.AbortWithMsg)

- ctx.IsBodyStream -> [ctx.Request.IsBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.IsBodyStream)

- ctx.IsConnect -> [ctx.Request.Header.IsConnect](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsConnect)

- ctx.IsDelete -> [ctx.Request.Header.IsDelete](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsDelete)

- ctx.IsOptions -> [ctx.Request.Header.IsOptions](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsOptions)

- ctx.IsPatch -> ctx.Method() == "patch"

- ctx.IsPut -> [ctx.Request.Header.IsPut](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsPut)

- ctx.IsTrace -> [ctx.Request.Header.IsTrace](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsTrace)

- ctx.LocalAddr -> ctx.GetConn().LocalAddr

- ctx.LocalIP -> ctx.GetConn().LocalAddr

- ctx.Redirect -> [ctx.Redirect](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Redirect)

- ctx.RedirectBytes -> [ctx.Redirect](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Redirect)

- ctx.Referer -> [ctx.Request.Header.Get("referer")](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Get)

- ctx.RemoteIP -> [ctx.ClientIP](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.ClientIP)

- ctx.RemoveUserValue -> use ctx.Keys directly

- ctx.RemoveUserValueBytes -> use ctx.Keys directly

- ctx.RequestURI -> [ctx.Request.Header.RequestURI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.RequestURI)

- ctx.ResetBody -> [ctx.Response.ResetBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.ResetBody)

- ctx.ResetUserValues -> use ctx.Keys directly

- ctx.SendFile -> [app.ServeFile(c, path)](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#ServeFile)

- ctx.SendFileBytes -> [app.ServeFile(c, path)](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#ServeFile)

- ctx.SetBody -> [ctx.Response.SetBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.SetBody)

- ctx.SetUserValue -> [ctx.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Set)

- ctx.SetUserValueBytes -> [ctx.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Set)

- ctx.Success -> [ctx.SetContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetContentType) + [ctx.SetBodyString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetBodyString)

- ctx.SuccessString -> [ctx.SetContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetContentType) + [ctx.SetBodyString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetBodyString)

- ctx.UserValue -> [ctx.Value](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Value)

- ctx.UserValueBytes -> [ctx.Value](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Value)

- ctx.VisitUserValues -> use ctx.Keys directly

- ctx.VisitUserValuesAll -> use ctx.Keys directly

### Need Another Library

- ctx.Logger -> [pkg/common/hlog](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/common/hlog#Logger)

- ctx.ID -> [RequestID](https://github.com/hertz-contrib/requestid)

### Unimplemented function

- ctx.ConnID

- ctx.ConnRequestNum

- ctx.ConnTime

- ctx.HijackSetNoResponse

- ctx.Init

- ctx.Init2

- ctx.LastTimeoutErrorResponse

- ctx.IsTLS

- ctx.SetBodyStreamWriter

- ctx.SetRemoteAddr

- ctx.TLSConnectionState

- ctx.Time

- ctx.TimeoutError

- ctx.TimeoutErrorWithCode

- ctx.TimeoutErrorWithResponse

## [fasthttp.Request](https://pkg.go.dev/github.com/valyala/fasthttp#Request)

### No Changed Function

- req.AppendBody -> [req.AppendBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.AppendBody)

- req.AppendBodyString -> [req.AppendBodyString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.AppendBodyString)

- req.Body -> [req.Body](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.Body)

- req.BodyWriteTo -> [req.BodyWriteTo](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.BodyWriteTo)

- req.BodyWriter -> [req.BodyWriter](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.BodyWriter)

- req.ConnectionClose -> [req.ConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.ConnectionClose)

- req.CopyTo -> [req.CopyTo](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.CopyTo)

- req.Host -> [req.Host](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.Host)

- req.IsBodyStream -> [req.IsBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.IsBodyStream)

- req.MayContinue -> [req.MayContinue](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.MayContinue)

- req.MultipartForm -> [req.MultipartForm](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.MultipartForm)

- req.PostArgs -> [req.PostArgs](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.PostArgs)

- req.RemoveMultipartFormFiles -> [req.RemoveMultipartFormFiles](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.RemoveMultipartFormFiles)

- req.RequestURI -> [req.RequestURI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.RequestURI)

- req.Reset -> [req.Reset](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.Reset)

- req.ResetBody -> [req.ResetBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.ResetBody)

- req.SetBody -> [req.SetBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetBody)

- req.SetBodyRaw -> [req.SetBodyRaw](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetBodyRaw)

- req.SetBodyStream -> [req.SetBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetBodyStream)

- req.SetBodyString -> [req.SetBodyString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetBodyString)

- req.SetHost -> [req.SetHost](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetHost)

- req.SetRequestURI -> [req.SetRequestURI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetRequestURI)

- req.SwapBody -> [req.SwapBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SwapBody)

- req.URI -> [req.URI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.URI)

### Need Change Function

- req.SetTimeout -> use [DoTimeout](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol/client#DoTimeout) directly

- req.SetHostBytes -> [req.SetHost](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetHost)

- req.SetRequestURIBytes -> [req.SetRequestURI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetRequestURI)

### Need Another Library

- req.BodyGunzip -> [gzip](https://github.com/hertz-contrib/gzip)

### Unimplement Function

- req.BodyInflate ->

- req.BodyUnbrotli ->

- req.BodyUncompressed ->

- req.ContinueReadBody ->

- req.ContinueReadBodyStream ->

- req.Read ->

- req.ReadBody ->

- req.ReadLimitBody ->

- req.ReleaseBody ->

- req.SetBodyStreamWriter ->

- req.SetURI ->

- req.Write ->

- req.WriteTo ->

## [fasthttp.RequestHeader](https://pkg.go.dev/github.com/valyala/fasthttp#RequestHeader)

### No Changed Function

- h.Add -> [h.Add](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Add)

- h.AppendBytes -> [h.AppendBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.AppendBytes)

- h.ConnectionClose -> [h.ConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.ConnectionClose)

- h.ContentLength -> [h.ContentLength](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.ContentLength)

- h.ContentType -> [h.ContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.ContentType)

- h.Cookie -> [h.Cookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Cookie)

- h.CopyTo -> [h.CopyTo](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.CopyTo)

- h.DelAllCookies -> [h.DelAllCookies](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.DelAllCookies)

- h.DelBytes -> [h.DelBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.DelBytes)

- h.DelCookie -> [h.DelCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.DelCookie)

- h.DisableNormalizing -> [h.DisableNormalizing](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.DisableNormalizing)

- h.HasAcceptEncoding -> [h.HasAcceptEncodingBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.HasAcceptEncodingBytes)

- h.HasAcceptEncodingBytes -> [h.HasAcceptEncodingBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.HasAcceptEncodingBytes)

- h.Header -> [h.Header](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Header)

- h.Host -> [h.Host](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Host)

- h.IsConnect -> [h.IsConnect](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsConnect)

- h.IsDelete -> [h.IsDelete](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsDelete)

- h.IsGet -> [h.IsGet](https://github.com/cloudwego/hertz/blob/v0.4.1/pkg/protocol/header.go#L1077)

- h.IsHTTP11 -> [h.IsHTTP11](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsHTTP11)

- h.IsHead -> [h.IsHead](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsHead)

- h.IsOptions -> [h.IsOptions](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsOptions)

- h.IsPost -> [h.IsPost](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsPost)

- h.IsPut -> [h.IsPut](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsPut)

- h.IsTrace -> [h.IsTrace](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.IsTrace)

- h.Len -> [h.Len](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Len)

- h.Method -> [h.Method](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Method)

- h.MultipartFormBoundary -> [h.MultipartFormBoundary](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.MultipartFormBoundary)

- h.Peek -> [h.Peek](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Peek)

- h.RawHeaders -> [h.RawHeaders](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.RawHeaders)

- h.RequestURI -> [h.RequestURI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.RequestURI)

- h.Reset -> [h.Reset](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Reset)

- h.ResetConnectionClose -> [h.ResetConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.ResetConnectionClose)

- h.Set -> [h.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Set)

- h.SetByteRange -> [h.SetByteRange](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetByteRange)

- h.SetBytesKV -> [h.SetBytesKV](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetBytesKV)

- h.SetCanonical -> [h.SetCanonical](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetCanonical)

- h.SetConnectionClose -> [h.SetConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetConnectionClose)

- h.SetContentLength -> [h.SetContentLength](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetContentLength)

- h.SetContentTypeBytes -> [h.SetContentTypeBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetContentTypeBytes)

- h.SetCookie -> [h.SetCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetCookie)

- h.SetHost -> [h.SetHost](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetHost)

- h.SetHostBytes -> [h.SetHostBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetHostBytes)

- h.SetMethod -> [h.SetMethod](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetMethod)

- h.SetMethodBytes -> [h.SetMethodBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetMethodBytes)

- h.SetMultipartFormBoundary -> [h.SetMultipartFormBoundary](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetMultipartFormBoundary)

- h.SetMultipartFormBoundaryBytes -> [h.SetMultipartFormBoundary](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetMultipartFormBoundary)

- h.SetNoDefaultContentType -> [h.SetNoDefaultContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetNoDefaultContentType)

- h.SetProtocol -> [h.SetProtocol](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetProtocol)

- h.SetRequestURI -> [h.SetRequestURI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetRequestURI)

- h.SetRequestURIBytes -> [h.SetRequestURIBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetRequestURIBytes)

- h.SetUserAgentBytes -> [h.SetUserAgentBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetUserAgentBytes)

- h.UserAgent -> [h.UserAgent](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.UserAgent)

- h.VisitAll -> [h.VisitAll](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.VisitAll)

- h.VisitAllCookie -> [h.VisitAllCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.VisitAllCookie)

### Need Change Function

- h.AddBytesK -> [h.Add](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Add)

- h.AddBytesKV -> [h.Add](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Add)

- h.AddBytesV -> [h.Add](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Add)

- h.CookieBytes -> [h.Cookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Cookie)

- h.Del -> [h.DelBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.DelBytes)

- h.DelCookieBytes -> [h.DelCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.DelCookie)

- h.IsPatch -> h.Method() == "patch"

- h.PeekBytes -> [h.Peek](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Peek)

- h.Protocol -> [h.GetProtocol](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.GetProtocol)

- h.SetBytesK -> [h.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Set)

- h.SetBytesV -> [h.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Set)

- h.SetContentType -> [h.SetContentTypeBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetContentTypeBytes)

- h.SetCookieBytesK -> [h.SetCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetCookie)

- h.SetCookieBytesKV -> [h.SetCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetCookie)

- h.SetProtocolBytes -> [h.SetProtocol](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetProtocol)

- h.SetUserAgent -> [h.SetUserAgentBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetUserAgentBytes)

### Unimplement Function

- h.AddTrailer

- h.AddTrailerBytes

- h.ConnectionUpgrade

- h.ContentEncoding

- h.EnableNormalizing

- h.PeekAll

- h.PeekKeys

- h.PeekTrailerKeys

- h.Read

- h.ReadTrailer

- h.Referer

- h.SetContentEncoding

- h.SetContentEncodingBytes

- h.SetRefererBytes

- h.SetTrailer

- h.SetTrailerBytes

- h.TrailerHeader

- h.VisitAllInOrder

- h.VisitAllTrailer

- h.Write

- h.WriteTo

## [fasthttp.Response](https://pkg.go.dev/github.com/valyala/fasthttp#Response)

### No Changed Function

- resp.AppendBody -> [resp.AppendBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.AppendBody)

- resp.AppendBodyString -> [resp.AppendBodyString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.AppendBodyString)

- resp.Body -> [resp.Body](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.Body)

- resp.BodyGunzip -> [resp.BodyGunzip](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.BodyGunzip)

- resp.BodyWriteTo -> [resp.BodyWriteTo](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.BodyWriteTo)

- resp.BodyWriter -> [resp.BodyWriter](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.BodyWriter)

- resp.ConnectionClose -> [resp.ConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.ConnectionClose)

- resp.CopyTo -> [resp.CopyTo](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.CopyTo)

- resp.IsBodyStream -> [resp.IsBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.IsBodyStream)

- resp.LocalAddr -> [resp.LocalAddr](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.LocalAddr)

- resp.RemoteAddr -> [resp.RemoteAddr](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.RemoteAddr)

- resp.Reset -> [resp.Reset](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.Reset)

- resp.ResetBody -> [resp.ResetBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.ResetBody)

- resp.SetBody -> [resp.SetBody](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.SetBody)

- resp.SetBodyRaw -> [resp.SetBodyRaw](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.SetBodyRaw)

- resp.SetBodyStream -> [resp.SetBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.SetBodyStream)

- resp.SetBodyString -> [resp.SetBodyString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.SetBodyString)

- resp.SetConnectionClose -> [resp.SetConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.SetConnectionClose)

- resp.SetStatusCode -> [resp.SetStatusCode](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.SetStatusCode)

- resp.StatusCode -> [resp.StatusCode](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.StatusCode)

### Unimplement Function

- resp.BodyInflate ->

- resp.BodyUnbrotli ->

- resp.BodyUncompressed ->

- resp.Read ->

- resp.ReadBody ->

- resp.ReadLimitBody ->

- resp.ReleaseBody ->

- resp.SendFile ->

- resp.SetBodyStreamWriter ->

- resp.SwapBody ->

- resp.Write ->

- resp.WriteDeflate ->

- resp.WriteDeflateLevel ->

- resp.WriteGzip ->

- resp.WriteGzipLevel ->

- resp.WriteTo ->

## [fasthttp.ResponseHeader](https://pkg.go.dev/github.com/valyala/fasthttp#ResponseHeader)

### No Changed Function

- h.Add -> [h.Add](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Add)

- h.AppendBytes -> [h.AppendBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.AppendBytes)

- h.ConnectionClose -> [h.ConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.ConnectionClose)

- h.ContentLength -> [h.ContentLength](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.ContentLength)

- h.ContentType -> [h.ContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.ContentType)

- h.Cookie -> [h.Cookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Cookie)

- h.CopyTo -> [h.CopyTo](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.CopyTo)

- h.Del -> [h.Del](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Del)

- h.DelAllCookies -> [h.DelAllCookies](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.DelAllCookies)

- h.DelBytes -> [h.DelBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.DelBytes)

- h.DelClientCookie -> [h.DelClientCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.DelClientCookie)

- h.DelClientCookieBytes -> [h.DelClientCookieBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.DelClientCookieBytes)

- h.DelCookie -> [h.DelCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.DelCookie)

- h.DelCookieBytes -> [h.DelCookieBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.DelCookieBytes)

- h.DisableNormalizing -> [h.DisableNormalizing](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.DisableNormalizing)

- h.Header -> [h.Header](https://pkg.go.dev/github.com/valyala/fasthttp#ResponseHeader.Header)

- h.IsHTTP11 -> [h.IsHTTP11](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.IsHTTP11)

- h.Len -> [h.Len](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Len)

- h.Peek -> [h.Peek](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Peek)

- h.Reset -> [h.Reset](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Reset)

- h.ResetConnectionClose -> [h.ResetConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.ResetConnectionClose)

- h.Server -> [h.Server](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Server)

- h.Set -> [h.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Set)

- h.SetBytesV -> [h.SetBytesV](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetBytesV)

- h.SetCanonical -> [h.SetCanonical](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetCanonical)

- h.SetConnectionClose -> [h.SetConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetConnectionClose)

- h.SetContentEncoding

- h.SetContentEncodingBytes

- h.SetContentLength -> [h.SetContentLength](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetContentLength)

- h.SetContentRange -> [h.SetContentRange](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetContentRange)

- h.SetContentType -> [h.SetContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetContentType)

- h.SetContentTypeBytes -> [h.SetContentTypeBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetContentTypeBytes)

- h.SetCookie -> [h.SetCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetCookie)

- h.SetNoDefaultContentType -> [h.SetNoDefaultContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetNoDefaultContentType)

- h.SetServerBytes -> [h.SetServerBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetServerBytes)

- h.SetStatusCode -> [h.SetStatusCode](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetStatusCode)

- h.StatusCode -> [h.StatusCode](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.StatusCode)

- h.VisitAll -> [h.VisitAll](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.VisitAll)

- h.VisitAllCookie -> [h.VisitAllCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.VisitAllCookie)

### Need Change Function

- h.AddBytesK -> [h.Add](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Add)

- h.AddBytesKV -> [h.Add](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Add)

- h.addBytesV -> [h.Add](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Add)

- h.PeekBytes -> [h.Peek](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Peek)

- h.SetBytesK -> [h.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Set)

- h.SetBytesKV -> [h.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Set)

- h.SetServer -> [h.SetServerBytes](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.SetServerBytes)

- h.SetLastModified -> [h.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Set)

- h.SetProtocol -> [h.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#ResponseHeader.Set)

### Unimplement Function

- h.AddTrailer

- h.AddTrailerBytes

- h.ConnectionUpgrade

- h.ContentEncoding

- h.EnableNormalizing

- h.PeekAll

- h.PeekKeys

- h.PeekTrailerKeys

- h.Protocol

- h.Read

- h.ReadTrailer

- h.SetStatusMessage

- h.SetTrailer

- h.SetTrailerBytes

- h.StatusMessage

- h.TrailerHeader

- h.VisitAllTrailer

- h.Write

- h.WriteTo
