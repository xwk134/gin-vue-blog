package routers

import "gvb_server/api"

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", app.ArticleCreateView)
	// router.GET("articles", app.ArticleListView)
	// router.PUT("articles/:id", app.ArticleUpdateView)
	// router.DELETE("articles", app.ArticleRemoveView)

}
