package main

//	func main() {
//		router := newRoute()
//		svr := http.Server{
//			Addr:         fmt.Sprintf(":%d", 8080),
//			Handler:      router,
//			IdleTimeout:  1 * time.Minute,
//			ReadTimeout:  10 * time.Second,
//			WriteTimeout: 30 * time.Second,
//		}
//		svr.ListenAndServe()
//	}
//func newRoute() *chi.Mux {
//	router := chi.NewRouter()
//	router.Get("/api/v1/health",
//		func(writer http.ResponseWriter, request *http.Request) {
//			if request.URL.Query().Get("name") == "" {
//				http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
//				return
//			}
//			writer.Write([]byte("Hello World!"))
//		},
//	)
//	router.Post("/api/v1/books", _sayHello)
//	router.Method(http.MethodGet, "/api/v1/books", sayHello())
//	return router
//}

//
//func sayHello() http.Handler {
//	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		writer.Write([]byte("Hello World!"))
//	})
//}
//
//func _sayHello(writer http.ResponseWriter, request *http.Request) {
//	writer.Write([]byte("Hello World!"))
//}

//	func newRoutes() *chi.Mux {
//		router := chi.NewRouter()
//
//		// TODO:
//		// - do we need some grouping, based on use case?
//		// - start with v1
//		router.Method(http.MethodPost, "/v1/discussion/delete/talk/user/{userID}", mw.HandlerFunc(discHandler.DeleteQuestionByUserID))
//
//		// Inbox Handler
//		router.Method(http.MethodGet, "/v1/inbox/isWhitelisted", mw.HandlerFunc(discHandler.IsShopWhitelisted))
//
//		router.Method(http.MethodGet, "/v1/discussion/ban/user/{userID}", mw.HandlerFunc(discHandler.CheckBanDiscussion))
//		router.Method(http.MethodPost, "/v1/discussion/ban/user", mw.HandlerFunc(discHandler.SetBanDiscussion))
//
//		router.HandleFunc("/healthcheck", discussionhandler.Health)
//		return router
//	}
//func (h *DiscussionHandler) IsShopWhitelisted(w http.ResponseWriter, r *http.Request) {
//	shopID := r.URL.Query().Get("shop_id")
//	if shopID == "" {
//		if _, err := response.WriteJSONAPIError(w, r, http.StatusBadRequest, "parameter shop_id shouldn't be empty"); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//		}
//		return
//	}
//
//	whitelisted := h.inboxUC.IsShopWhitelisted(r.Context(), shopID)
//	resp := isShopWhitelisted{IsWhitelisted: whitelisted}
//	if _, err := response.WriteJSONAPIData(w, r, http.StatusOK, resp); err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//}
