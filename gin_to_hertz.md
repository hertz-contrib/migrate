# Gin -> Hertz conversion table

## [gin.Context](https://pkg.go.dev/github.com/gin-gonic/gin#Context)

- All the pseudocode below assumes ctx have these types:

```Go
func handler(ctx *gin.Context){..}
->
func handler(
c context.Context
ctx *app.RequestContext
){..}
```

### Need Change Function

- ctx.Cookie -> [ctx.Cookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Cookie)

- ctx.SetCookie -> [ctx.SetCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetCookie)

- ctx.SetSameSite -> use [ctx.SetCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetCookie) directly

- ctx.RemoteIP -> net.SplitHostPort(strings.TrimSpace([ctx.RemoteAddr()](https://pkg.go.dev/github.com/cloudwego/hertz/pkg/app#RequestContext.RemoteAddr).String()))

- ctx.Stream -> use [ctx.SetBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.SetBodyStream)

### Unimplemented function

- ctx.AddParam

- ctx.AsciiJSON

- ctx.BindHeader

- ctx.BindJSON

- ctx.BindQuery

- ctx.BindTOML

- ctx.BindUri

- ctx.BindWith

- ctx.BindXML

- ctx.BindYAML

- ctx.DataFromReader

- ctx.GetPostFormArray

- ctx.GetPostFormMap

- ctx.GetQueryArray

- ctx.GetQueryMap

- ctx.IsWebsocket

- ctx.JSONP

- ctx.MustBindWith

- ctx.Negotiate

- ctx.NegotiateFormat

- ctx.PostFormArray

- ctx.PostFormMap

- ctx.QueryArray

- ctx.QueryMap

- ctx.SSEvent

- ctx.SecureJSON

- ctx.SetAccepted

- ctx.ShouldBind

- ctx.ShouldBindBodyWith

- ctx.ShouldBindHeader

- ctx.ShouldBindJSON

- ctx.ShouldBindQuery

- ctx.ShouldBindTOML

- ctx.ShouldBindUri

- ctx.ShouldBindWith

- ctx.ShouldBindXML

- ctx.ShouldBindYAML

- ctx.TOML

- ctx.YAML

## [http.Request](https://pkg.go.dev/net/http#Request)

- r.Method -> [r.Method](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.Method) & [r.SetMethod](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetMethod)

- r.URL -> [r.URL](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.URI)

- r.Proto -> [r.Header.GetProtocol](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.GetProtocol)

- r.Header -> [r.Header](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader)

- r.Body -> [r.BodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.BodyStream) & [r.SetBodyStream](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetBodyStream)

- r.ContentLength -> [r.Header.ContentLength](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.ContentLength)

- r.Close -> [r.Header.ConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.ConnectionClose) & [r.Header.SetConnectionClose](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.SetConnectionClose)

- r.Host -> [r.Header.Host](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Host)

- r.Form -> [r.URL().QueryArgs](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#URI.QueryArgs) & [r.PostArgs](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.PostArgs)

- r.MultipartForm -> [r.MultipartForm](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.MultipartForm)

- r.RemoteAddr -> [r.RemoteAddr](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Response.RemoteAddr)

- r.RequestURI -> [r.Header.RequestURI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.RequestURI) & [r.SetRequestURI](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetRequestURI)

- r.AddCookie -> [r.SetCookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#Request.SetCookie)

- r.Cookie -> [r.Header.Cookie](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.Cookie)

- r.UserAgent -> [r.Header.UserAgent](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/protocol#RequestHeader.UserAgent)


### No Changed Function

- ctx.Abort -> [ctx.Abort](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Abort)

- ctx.AbortWithError -> [ctx.AbortWithError](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.AbortWithError)

- ctx.AbortWithStatus -> [ctx.AbortWithStatus](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.AbortWithStatus)

- ctx.AbortWithStatusJSON -> [ctx.AbortWithStatusJSON](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.AbortWithStatusJSON)

- ctx.Bind -> [ctx.Bind](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Bind)

- ctx.ClientIP -> [ctx.ClientIP](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.ClientIP)

- ctx.ContentType -> [ctx.ContentType](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.ContentType)

- ctx.Copy -> [ctx.Copy](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Copy)

- ctx.Data -> [ctx.Data](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Data)

- ctx.DefaultPostForm -> [ctx.DefaultPostForm](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.DefaultPostForm)

- ctx.DefaultQuery -> [ctx.DefaultQuery](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.DefaultQuery)

- ctx.Error -> [ctx.Error](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Error)

- ctx.File -> [ctx.File](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.File)

- ctx.FileAttachment -> [ctx.FileAttachment](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.FileAttachment)

- ctx.FileFromFS -> [ctx.FileFormFS](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.FileFromFS)

- ctx.FormFile -> [ctx.FormFile](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.FormFile)

- ctx.FullPath -> [ctx.FullPath](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.FullPath)

- ctx.Get -> [ctx.Get](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Get)

- ctx.GetBool -> [ctx.GetBool](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetBool)

- ctx.GetDuration -> [ctx.GetDuartion](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetDuration)

- ctx.GetFloat64 -> [ctx.GetFloat64](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetFloat64)

- ctx.GetHeader -> [ctx.GetHeader](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetHeader)

- ctx.GetInt -> [ctx.GetInt](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetInt)

- ctx.GetInt64 -> [ctx.GetInt64](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetInt64)

- ctx.GetPostForm -> [ctx.GetPostForm](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetPostForm)

- ctx.GetQuery -> [ctx.GetQuery](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetQuery)

- ctx.GetRawData -> [ctx.GetRawData](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetRawData)

- ctx.GetString -> [ctx.GetString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetString)

- ctx.GetStringMap -> [ctx.GetStringMap](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetStringMap)

- ctx.GetStringMapString -> [ctx.GetStringMapString](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetStringMapString)

- ctx.GetStringMapStringSlice -> [ctx.GetStringMapStringSlice](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetStringMapStringSlice)

- ctx.GetStringSlice -> [ctx.GetStringSlice](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetStringSlice)

- ctx.GetTime -> [ctx.GetTime](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetTime)

- ctx.GetUint -> [ctx.GetUint](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetUint)

- ctx.GetUint64 -> [ctx.GetUint64](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.GetUint64)

- ctx.HTML -> [ctx.HTML](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.HTML)

- ctx.Handler -> [ctx.Handler](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Handler)

- ctx.HandlerName -> [ctx.HandlerName](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.HandlerName)

- ctx.Header -> [ctx.Header](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Header)

- ctx.IndentedJSON -> [ctx.IndentedJSON](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.IndentedJSON)

- ctx.IsAborted -> [ctx.IsAborted](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.IsAborted)

- ctx.JSON -> [ctx.JSON](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.JSON)

- ctx.MultipartForm -> [ctx.MultipartForm](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.MultipartForm)

- ctx.MustGet -> [ctx.MustGet](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.MustGet)

- ctx.Next -> [ctx.Next](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Next)

- ctx.Param -> [ctx.Param](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Param)

- ctx.PostForm -> [ctx.PostForm](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.PostForm)

- ctx.ProtoBuf -> [ctx.ProtoBuf](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.ProtoBuf)

- ctx.PureJSON -> [ctx.PureJSON](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.PureJSON)

- ctx.Query -> [ctx.Query](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Query)

- ctx.Redirect -> [ctx.Redirect](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Redirect)

- ctx.Render -> [ctx.Render](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Render)

- ctx.Set -> [ctx.Set](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Set)

- ctx.Status -> [ctx.Status](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Status)

- ctx.String -> [ctx.String](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.String)

- ctx.Value -> [ctx.Value](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.Value)

- ctx.XML -> [ctx.XML](https://pkg.go.dev/github.com/cloudwego/hertz@v0.4.1/pkg/app#RequestContext.XML)

### context.Context Interface

- ctx.Deadline -> c.Deadline

- ctx.Done -> c.Done

- ctx.Err -> c.Err

- ctx.Value -> c.Value
