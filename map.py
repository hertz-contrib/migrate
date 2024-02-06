# Copyright 2022 CloudWeGo Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


package_map = {
    "github.com/valyala/fasthttp": [
        {
            "name": "RequestCtx",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
            "afterName": "RequestContext",
        },
        {
            "name": "HeaderAuthorization",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/consts",
        },
        {
            "name": "MethodGet",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/consts",
        },
        {
            "name": "StatusContinue",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/consts",
        },
        {
            "name": "AcquireTimer",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/timer",
        },
        {
            "name": "AppendGunzipBytes",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/compress",
        },
        {
            "name": "AppendGzipBytes",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/compress",
        },
        {
            "name": "AppendGzipBytesLevel",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/compress",
        },
        {
            "name": "AppendHTTPDate",
            "pkgName": "github.com/cloudwego/hertz/internal/bytesconv",
        },
        {
            "name": "AppendQuotedArg",
            "pkgName": "github.com/cloudwego/hertz/internal/bytesconv",
        },
        {
            "name": "AppendUint",
            "pkgName": "github.com/cloudwego/hertz/internal/bytesconv",
        },
        {
            "name": "DialTimeout",
            "pkgName": "github.com/cloudwego/hertz/pkg/network/dialer",
        },
        {
            "name": "Do",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/client",
        },
        {
            "name": "DoDeadline",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/client",
        },
        {
            "name": "DoRedirects",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/client",
        },
        {
            "name": "DoTimeout",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/client",
        },
        {
            "name": "Get",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/bytebufferpool",
        },
        {
            "name": "GetDeadline",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/client",
        },
        {
            "name": "GetTimeout",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/client",
        },
        {
            "name": "ParseHTTPDate",
            "pkgName": "github.com/cloudwego/hertz/internal/bytesconv",
        },
        {
            "name": "ParseUint",
            "pkgName": "github.com/cloudwego/hertz/internal/bytesconv",
        },
        {
            "name": "Post",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/client",
        },
        {
            "name": "ReleaseCookie",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "ReleaseRequest",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "ReleaseResponse",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "ReleaseTimer",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/timer",
        },
        {
            "name": "ReleaseURI",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "ServeFile",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "ServeFileUncompressed",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "StatusCodeIsRedirect",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/client",
        },
        {
            "name": "StatusMessage",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/consts",
        },
        {
            "name": "WriteGunzip",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/compress",
        },
        {
            "name": "WriteGzipLevel",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/compress",
        },
        {
            "name": "WriteMultipartForm",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "Args",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "Client",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/client",
        },
        {
            "name": "Cookie",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "AcquireCookie",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "CookieSameSite",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "CookieSameSiteDisabled",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "ErrBodyStreamWritePanic",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/http1/resp",
        },
        {
            "name": "FS",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "HijackHandler",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "HostClient",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/client",
        },
        {
            "name": "PathRewriteFunc",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "NewPathSlashesStripper",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "NewVHostPathRewriter",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "Request",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "AcquireRequest",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "RequestHeader",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "Resolver",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/client/discovery",
        },
        {
            "name": "Response",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "AcquireResponse",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "ResponseHeader",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "RetryIfFunc",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/client",
        },
        {
            "name": "Server",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol/http1",
        },
        {
            "name": "URI",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
        {
            "name": "AcquireURI",
            "pkgName": "github.com/cloudwego/hertz/pkg/protocol",
        },
    ],
    "github.com/valyala/fasthttp/prefork": [
        {
            "name": "Logger",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/hlog",
        },
        {
            "name": "New",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/errors",
        },
    ],
    "github.com/valyala/fasthttp/stackless": [
        {
            "name": "NewFunc",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/stackless",
        },
        {
            "name": "NewWriterFunc",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/stackless",
        },
        {
            "name": "Writer",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/stackless",
        },
        {
            "name": "NewWriter",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/stackless",
        },
    ],
    "github.com/fasthttp/router": [
        {
            "name": "New",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server",
            "afterName": "Default",
        }
    ],
    "github.com/buaazp/fasthttprouter": [
        {
            "name": "New",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server",
            "afterName": "Default",
        }
    ],
    "github.com/gin-gonic/gin": [
        {
            "name": "Context",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
            "afterName": "RequestContext",
        },
        {
            "name": "Version",
            "pkgName": "github.com/cloudwego/hertz/cmd/hz/meta",
        },
        {
            "name": "Mode",
            "pkgName": "github.com/cloudwego/hertz/cmd/hz/meta",
        },
        {
            "name": "Accounts",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth",
        },
        {
            "name": "Engine",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server",
            "afterName": "Hertz",
        },
        {
            "name": "New",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server",
        },
        {
            "name": "Error",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/errors",
        },
        {
            "name": "ErrorType",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/errors",
        },
        {
            "name": "ErrorTypeBind",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/errors",
        },
        {
            "name": "H",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/utils",
        },
        {
            "name": "HandlerFunc",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "BasicAuth",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth",
        },
        {
            "name": "BasicAuthForRealm",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth",
        },
        {
            "name": "Bind",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/binding",
        },
        {
            "name": "Logger",
            "pkgName": "github.com/cloudwego/hertz/pkg/common/hlog",
        },
        {
            "name": "Recovery",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery",
        },
        {
            "name": "HandlersChain",
            "pkgName": "github.com/cloudwego/hertz/pkg/app",
        },
        {
            "name": "IRouter",
            "pkgName": "github.com/cloudwego/hertz/pkg/route",
        },
        {
            "name": "IRoutes",
            "pkgName": "github.com/cloudwego/hertz/pkg/route",
        },
        {
            "name": "Param",
            "pkgName": "github.com/cloudwego/hertz/pkg/route/param",
        },
        {
            "name": "Params",
            "pkgName": "github.com/cloudwego/hertz/pkg/route/param",
        },
        {
            "name": "RouteInfo",
            "pkgName": "github.com/cloudwego/hertz/pkg/route",
        },
        {
            "name": "RouterGroup",
            "pkgName": "github.com/cloudwego/hertz/pkg/route",
        },
        {
            "name": "RoutesInfo",
            "pkgName": "github.com/cloudwego/hertz/pkg/route",
        },
        {
            "name": "Default",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server",
        },
    ],
    "github.com/gin-gonic/gin/render": [
        {
            "name": "Data",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "Delims",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "HTML",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "HTMLDebug",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "HTMLProduction",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "HTMLRender",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "IndentedJSON",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "ProtoBuf",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "PureJSON",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "Reader",
            "pkgName": "github.com/cloudwego/hertz/pkg/network",
        },
        {
            "name": "Render",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "String",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
        {
            "name": "XML",
            "pkgName": "github.com/cloudwego/hertz/pkg/app/server/render",
        },
    ],
    "github.com/gin-contrib/cors": [
        {
            "name": "Default",
            "pkgName": "github.com/hertz-contrib/cors",
        },
        {"name": "New", "pkgName": "github.com/hertz-contrib/cors"},
        {
            "name": "Config",
            "pkgName": "github.com/hertz-contrib/cors",
        },
        {
            "name": "DefaultConfig",
            "pkgName": "github.com/hertz-contrib/cors",
        },
    ],
    "github.com/appleboy/gin-jwt/v2": [
        {
            "name": "GetToken",
            "pkgName": "github.com/hertz-contrib/jwt",
        },
        {"name": "New", "pkgName": "github.com/hertz-contrib/jwt"},
        {
            "name": "MapClaims",
            "pkgName": "github.com/hertz-contrib/jwt",
        },
        {
            "name": "ExtractClaims",
            "pkgName": "github.com/hertz-contrib/jwt",
        },
        {
            "name": "ExtractClaimsFromToken",
            "pkgName": "github.com/hertz-contrib/jwt",
        },
        {
            "name": "GinJWTMiddleware",
            "pkgName": "github.com/hertz-contrib/jwt",
            "afterName": "HertzJWTMiddleware",
        },
    ],
    "github.com/swaggo/gin-swagger": [
        {
            "name": "CustomWrapHandler",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "DeepLinking",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "DefaultModelsExpandDepth",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "DocExpansion",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "InstanceName",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "Oauth2DefaultClientID",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "PersistAuthorization",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "URL",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "WrapHandler",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
        {
            "name": "Config",
            "pkgName": "github.com/hertz-contrib/swagger",
        },
    ],
    "github.com/gin-contrib/pprof": [
        {
            "name": "DefaultPrefix",
            "pkgName": "github.com/hertz-contrib/pprof",
        },
        {
            "name": "Register",
            "pkgName": "github.com/hertz-contrib/pprof",
        },
        {
            "name": "RouteRegister",
            "pkgName": "github.com/hertz-contrib/pprof",
        },
    ],
    "github.com/gin-contrib/requestid": [
        {
            "name": "Get",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
        {
            "name": "New",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
        {
            "name": "Generator",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
        {
            "name": "Handler",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
        {
            "name": "HeaderStrKey",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
        {
            "name": "Option",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
        {
            "name": "WithCustomHeaderStrKey",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
        {
            "name": "WithGenerator",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
        {
            "name": "WithHandler",
            "pkgName": "github.com/hertz-contrib/requestid",
        },
    ],
}
