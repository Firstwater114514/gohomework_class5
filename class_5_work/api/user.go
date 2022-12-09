package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"lanshan_homework/go1.19.2/go_homework/class_5_work/api/middleware"
	"lanshan_homework/go1.19.2/go_homework/class_5_work/dao"
	"lanshan_homework/go1.19.2/go_homework/class_5_work/model"
	"lanshan_homework/go1.19.2/go_homework/class_5_work/utils"
	"strconv"
	"time"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.Register{}); err != nil {
		utils.RespSuccess(c, "部分数据未输入，请检查")
		return
	}
	IfLogin, _ := dao.IfLogin()
	if IfLogin == "yes" {
		utils.RespFail(c, "已有账号登录，请退出在线账号后登录")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	username := c.PostForm("username")
	flag, _ := dao.SelectUser(username)
	if flag == "yes" {
		utils.RespFail(c, "用户名已被使用")
		return
	} else if flag == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	password := c.PostForm("password")
	checkQuestion := c.PostForm("check question")
	checkAnswer := c.PostForm("check answer")
	dao.AddUser(username, password, checkQuestion, checkAnswer)
	utils.RespSuccess(c, "注册成功!快去登录吧~")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.Login{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	IfLogin, _ := dao.IfLogin()
	if IfLogin == "yes" {
		utils.RespFail(c, "已有账号登录，请退出在线账号后登录")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag, uid := dao.SelectUser(username)
	if flag == "no" {
		utils.RespFail(c, "该用户不存在")
		return
	} else if flag == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	selectPassword := dao.SelectPasswordFromId(uid)
	if selectPassword != password && selectPassword != "" {
		utils.RespFail(c, "密码错误！")
		return
	} else if selectPassword == "" {
		utils.RespFail(c, "数据错误！")
		return
	}
	dao.Login(uid)
	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "sqy",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.Forget(c, tokenString, "记得关闭窗口前一定要退出账号哦!!!!!用quit")
}

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

func changePassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Change{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	IfLogin, uid := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	password := dao.SelectPasswordFromId(uid)
	newPassword := c.PostForm("new password")
	if password == newPassword {
		utils.RespFail(c, "新密码与旧密码相同~")
		return
	} else if password == "" {
		utils.RespFail(c, "数据错误！")
		return
	}
	dao.ChangePassword(uid, newPassword)
	dao.Quit(uid)
	utils.RespSuccess(c, "更改密码成功，请重新登录~")
	return
}
func forgetPassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Forget{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	IfLogin, _ := dao.IfLogin()
	if IfLogin == "yes" {
		utils.RespFail(c, "该操作需要无帐号在线哦，退出后再使用吧")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	username := c.PostForm("username")
	flag, uid := dao.SelectUser(username)
	if flag == "no" {
		utils.RespFail(c, "该用户不存在")
		return
	} else if flag == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	if dao.CheckQuestion(uid) == "" {
		utils.RespFail(c, "数据错误！")
		return
	}
	utils.Forget(c, dao.CheckQuestion(uid), "把请求中的forget password改成answer并且输入answer、username")
}
func answer(c *gin.Context) {
	if err := c.ShouldBind(&model.Question{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	username := c.PostForm("username")
	answer := c.PostForm("answer")
	check, uid := dao.CheckAnswer(answer, username)
	if !check && uid == "false" {
		utils.RespFail(c, "验证答案失败，请检查后重新输入")
		return
	} else if !check && uid == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	dao.IfChange(uid)
	utils.Forget(c, "验证成功！", "将请求中的answer改成update password并且输入new password")
}
func updatePassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Change{}); err != nil {
		utils.RespFail(c, "请输入新的密码")
		return
	}
	uid := dao.WhoChange()
	if uid == "" {
		utils.RespFail(c, "数据错误！")
		return
	}
	newPassword := c.PostForm("new password")
	dao.ChangePassword(uid, newPassword)
	dao.ChangeOver(uid)
	utils.RespSuccess(c, "更改密码成功！可以登录了哦")
	return
}
func addComment(c *gin.Context) {
	if err := c.ShouldBind(&model.AddComment{}); err != nil {
		utils.RespFail(c, "请输入你的留言")
		return
	}
	IfLogin, uid := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据错误！")
		return
	}
	username := dao.WhoLogin(uid)
	comment := c.PostForm("comment")
	realComment := username + ":" + comment
	dao.AddComment(realComment)
	utils.RespSuccess(c, "留言成功，把请求中的add comment改成scan comments查看留言板")
}
func scanComments(c *gin.Context) {
	IfLogin, _ := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	f := dao.GetFloorNumber()
	if f == 0 {
		utils.RespSuccess(c, "暂时还没有留言啦，第一个留言就交给你咯~")
		return
	}
	utils.RespSuccess(c, "留言板：")
	for floor := 1; floor <= f; floor++ {
		rc := dao.ReadComment(floor)
		if rc == "" {
			utils.RespFail(c, "数据错误！")
			return
		}
		utils.Comment(c, floor, rc)
	}
}
func deleteComment(c *gin.Context) {
	if err := c.ShouldBind(&model.Num{}); err != nil {
		utils.RespFail(c, "请输入想删除的留言序号哦")
		return
	}
	IfLogin, _ := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	num := c.PostForm("num")
	fm := dao.GetFloorNumber()
	if fm == 0 {
		utils.RespSuccess(c, "还没有人留言哦，没有留言可以删呢~")
		return
	} else if fm == -1 {
		utils.RespFail(c, "数据错误！")
		return
	}
	numI, errN := strconv.Atoi(num)
	if errN != nil {
		panic(errN)
	}
	if numI > fm {
		utils.RespFail(c, "没有该序号的留言")
		return
	}
	dao.DeleteComment(numI)
	fm = dao.GetFloorNumber()
	if fm == 0 {
		utils.RespSuccess(c, "成功删除留言，留言板暂无留言")
		return
	}
	utils.RespSuccess(c, "留言板：")
	for floor := 1; floor <= fm; floor++ {
		rc := dao.ReadComment(floor)
		if rc == "" {
			utils.RespFail(c, "数据错误！")
		}
		utils.Comment(c, floor, rc)
	}
}
func clearComments(c *gin.Context) {
	IfLogin, _ := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	if dao.GetFloorNumber() == 0 {
		utils.RespFail(c, "留言板本来就是空的哦~")
		return
	}
	dao.ClearComments()
	utils.RespSuccess(c, "清除留言板成功")
}
func quit(c *gin.Context) {
	IfLogin, uid := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	dao.Quit(uid)
	utils.RespSuccess(c, "成功退出账号")
	return
}
func unsubscribe(c *gin.Context) {
	IfLogin, uid := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	dao.Unsubscribe(uid)
	utils.RespSuccess(c, "注销账户成功")
}
func clearAll(c *gin.Context) {
	IfLogin, _ := dao.IfLogin()
	if IfLogin == "yes" {
		utils.RespFail(c, "该操作不可有帐号在线")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	dao.ClearAll()
	utils.RespSuccess(c, "成功初始化该系统！")
}
func like(c *gin.Context) {
	if err := c.ShouldBind(&model.Like{}); err != nil {
		utils.RespFail(c, "请输入想点赞的用户哦")
		return
	}
	IfLogin, uid := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	username := dao.WhoLogin(uid)
	like := c.PostForm("like")
	flag, lid := dao.SelectUser(like)
	if lid == uid {
		utils.RespFail(c, "不可以给自己点赞哦")
		return
	}
	if flag == "return" {
		utils.RespFail(c, "数据错误！")
		return
	} else if flag == "no" {
		utils.RespFail(c, "被点赞的用户不存在")
		return
	}
	ok := dao.LikeSomeone(username, like)
	if ok == 0 {
		utils.RespFail(c, "点赞过了哦")
		return
	} else if ok == -1 {
		utils.RespFail(c, "数据错误！")
	}
	utils.RespSuccess(c, "点赞成功！")
}
func cancelLike(c *gin.Context) {
	if err := c.ShouldBind(&model.CancelLike{}); err != nil {
		utils.RespFail(c, "请输入想取消点赞的用户哦")
		return
	}
	IfLogin, uid := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	username := dao.WhoLogin(uid)
	cancelLike := c.PostForm("cancel like")
	flag, lid := dao.SelectUser(cancelLike)
	if lid == uid {
		utils.RespFail(c, "不可以对自己操作哦")
		return
	}
	if flag == "return" {
		utils.RespFail(c, "数据错误！")
		return
	} else if flag == "no" {
		utils.RespFail(c, "被取消点赞的用户不存在")
		return
	}
	n := dao.CancelLike(username, cancelLike)
	if n == -1 {
		utils.RespFail(c, "数据错误！")
		return
	} else if n == 0 {
		utils.RespFail(c, "本来就没有赞哦")
		return
	}
	utils.RespSuccess(c, "取消点赞啦~")
}
func myLikes(c *gin.Context) {
	IfLogin, uid := dao.IfLogin()
	if IfLogin == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if IfLogin == "return" {
		utils.RespFail(c, "数据据错误！")
		return
	}
	username := dao.WhoLogin(uid)
	n := dao.MyLikes(username)
	if n == -1 {
		utils.RespFail(c, "数据错误！")
		return
	}
	num := "您有" + strconv.Itoa(int(n)) + "个赞哦"
	utils.RespSuccess(c, num)
}
