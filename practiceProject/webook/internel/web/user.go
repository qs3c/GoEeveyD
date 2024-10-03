package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"practiceProject/webook/internel/domain"
	"practiceProject/webook/internel/service"
)

const (
	emailPatten    = `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
	passwordPatten = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$"
)

// 把所有和某个功能下的所有路由方法绑定到一个结构体上（包括注册这些方法的方法也绑定到这个结构体上）
// 比如：用户注册、登陆等相关的 “路由方法” 都绑定到 “用户” 结构体上

type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	// service 层的服务
	svc *service.UserService
}

// 定义一个初始化的方法来进行预编译这俩个正则并存在结构体实力的变量中

// 正则表达式预编译函数

//func (u *UserHandler) PreCompile() {
//
//	u.emailExp = regexp.MustCompile(emailPatten, regexp.None)
//	u.passwordExp = regexp.MustCompile(passwordPatten, regexp.None)
//}

// 老师写的是这个，并非结构体方法，而是一个普通的方法但是返回带有这两个参数的结构体
// 老师的这个写法其实是一种初始化的模式，这种写法感觉更好
// 初始化函数写法：

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailExp:    regexp.MustCompile(emailPatten, regexp.None),
		passwordExp: regexp.MustCompile(passwordPatten, regexp.None),
		//
		svc: svc,
	}
}

// 路由方法：

func (u *UserHandler) Signup(c *gin.Context) {
	// 接收http数据并处理
	// 定义结构体来接收
	// 这个结构体定义在内部的好处就是别的方法不能用，就在这里用
	type SignupReq struct {
		// 意思是对应到 json 格式的时候名字是 email
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	// 声明一个实例 req 然后用 Bind 来自动解析和填入 req
	// Bind 是根据 http head 中的 content-type 类型来决定如何处理数据的
	// 我们这里 content-type 设置的是 application/json
	// 解析过程出问题会直接写回 400 错误
	var req SignupReq
	if err := c.Bind(&req); err != nil {
		return
	}

	// 校验数据——正则表达式
	// 1 校验合法邮箱
	//const emailPatten = `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
	//ok, err := regexp.Match(emailPatten, []byte(req.Email))
	// 将正则表达式匹配变成了两个阶段 预编译和匹配
	// 预编译的作用是为了提升正则匹配的执行效率
	// 所以要利用好这个预编译，比如把编译和匹配这两句写在一起就起不到任何作用
	// 如果我把预编译放在外部，提前做好并只做一次（实例化 handler 的时候）
	// 然后每次校验就只是用编译好的正则进行匹配
	// 效率就提高了
	emailExp := regexp.MustCompile(emailPatten, regexp.None)
	ok, err := emailExp.MatchString(req.Email)

	if err != nil {
		// 正则表达式写的不对
		c.String(http.StatusOK, "系统错误！")
		return
	}
	if !ok {
		// 邮箱格式输入不对
		c.String(http.StatusOK, "邮箱格式不对！")
		return
	}

	// 2 校验密码符合规律
	//const passwordPatten = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]{8,}$"
	//ok, err = regexp.Match(passwordPatten, []byte(req.Password))
	passwordExp := regexp.MustCompile(passwordPatten, regexp.None)
	ok, err = passwordExp.MatchString(req.Password)

	if err != nil {
		// 正则表达式写的不对
		c.String(http.StatusOK, "系统错误！")
		return
	}
	if !ok {
		// 密码格式输入不对
		c.String(http.StatusOK, "密码需大于8位，包含大小写字母、数字和特殊字符")
		return
	}

	// 3 校验两次密码输入相等

	if req.ConfirmPassword != req.Password {
		c.String(http.StatusOK, "两次输入密码不一致！")
		return
	}

	// 调用 service 的方法（因为现在handler中已经有service结构体了）
	// 这里上下文给的是 gin 的 context c
	err = u.svc.SignUp(c, domain.User{Email: req.Email, Password: req.Password})
	if err != nil {
		c.String(http.StatusOK, "注册服务异常")
		return
	}
	// errors.Is(err, service.ErrDuplicateEmail)
	if err == service.ErrDuplicateEmail {
		c.String(http.StatusOK, "邮箱已被注册")
		return
	}
	c.String(http.StatusOK, "signup success")
	//fmt.Printf("%v", req)

}
func (u *UserHandler) Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return
	}

	user, err := u.svc.Login(c, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		c.String(http.StatusOK, "账号/密码错误")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}

	// 登录成功后把 seesion 设置一下
	sess := sessions.Default(c)
	sess.Set("user_id", user.Id)
	_ = sess.Save()

	c.String(http.StatusOK, "登陆成功")
	return

}
func (u *UserHandler) Edit(c *gin.Context)    {}
func (u *UserHandler) Profile(c *gin.Context) {}

// 路由注册：

func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	server.POST("/users/signup", u.Signup)
	server.POST("/users/login", u.Login)
	server.POST("/users/edit", u.Edit)
	server.GET("/users/profile", u.Profile)

}

// 分组路由：

func (u *UserHandler) RegisterRoutersWithGroup(ug *gin.RouterGroup) {
	ug.PUT("/users/signup", u.Signup)
	ug.POST("/users/login", u.Login)
	ug.GET("/users/edit", u.Edit)
	ug.GET("/users/profile", u.Profile)
}
