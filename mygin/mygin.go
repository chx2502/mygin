package mygin

import (
	"log"
	"net/http"
)

type HandleFunc func(c *Context)

type RouterGroup struct {
	prefix string	// 通过前缀来确定分组
	middlewares []HandleFunc	// 不同的分组拥使用不同的中间件
	parent *RouterGroup	// 子分组拥有父分组的功能
	engine *Engine	// 通过持有一个 Engine 来获得与路由有关的功能
}

type Engine struct {
	*RouterGroup	// 嵌套匿名结构体，类似于子类继承父类
	router *router
	groups []*RouterGroup	// Engine 作为整个服务的入口，持有所有 group 信息
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) NewGroup(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, path string, handler HandleFunc) {
	pattern := group.prefix + path
	log.Printf("Route %6s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
