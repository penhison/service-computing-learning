package service

// 实现了使用GET方法模拟登录和注册的方法

import (
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
    "fmt"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    mx.HandleFunc("/hello/{id}", helloHandler(formatter)).Methods("GET")
    mx.HandleFunc("/register", registerHandler(formatter)).Methods("GET")
    mx.HandleFunc("/login", loginHandler(formatter)).Methods("GET")
}

func helloHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        id := vars["id"]
        // formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
        formatter.Text(w, http.StatusOK, "hello, " + id)
    }
}

type user struct{
    name string
    password string
}

// 模拟数据库
var users = make(map[string]user)


// 解析url参数为user
func parseUser(req *http.Request) user{
    req.ParseForm()  //解析参数，默认是不会解析的
    // fmt.Println(req.Form)  //这些信息是输出到服务器端的打印信息
    n, ok1 := req.Form["name"]
    p, ok2 := req.Form["password"]
    if ok1 && ok2 {
        return user{n[0], p[0]}
    } else {
        return user{}
    }
}

// 注册器
func registerHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        // vars := mux.Vars(req)
        // name := vars["name"]
        // password := vars["password"]
        // formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
        u := parseUser(req)
        fmt.Println(u)
        if u.name != "" && u.password != "" {
            _, ok := users[u.name]
            if ok {
                formatter.Text(w, http.StatusOK, u.name + " has already been register!")
                return
            }
        }
        users[u.name] = u
        formatter.Text(w, http.StatusOK, "user " + u.name + " register success!")

    }
}

// 登录器
func loginHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        // vars := mux.Vars(req)
        // name := vars["name"]
        // password := vars["password"]
        // formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
        u := parseUser(req)
        fmt.Println(u)
        if u.name != "" && u.password != "" {
            u2, ok := users[u.name]
            if ok && u.password == u2.password {
                formatter.Text(w, http.StatusOK, "user " + u.name + " login success!")
                return
            }
        }
        formatter.Text(w, http.StatusOK, "login fail! username or password incorrect")
    }
}
