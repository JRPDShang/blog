package models

import (
	"github.com/beego/beego/v2/client/orm"
	context2 "github.com/beego/beego/v2/server/web/context"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
	"time"
)

const (
	dbName       = "data/blog.db"
	sqliteDriver = "sqlite3"
)

type Category struct {
	Id         int64
	Title      string
	Created    time.Time `orm:"index"`
	Views      int64     `orm:"index"`
	TopicTime  time.Time `orm:"index"`
	TopicCount int64
}
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(50000)"`
	Label           string `orm:"size(1000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Category        string
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}
type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}
type User struct {
	Id int64
	NickName string
	Password string
	Replies int64
}

func RegisterDb() {
	//if !com.IsExit(_DB_NAME){
	//	os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
	//	os.Create(_DB_NAME)
	//}
	orm.RegisterModel(new(Category), new(Topic), new(Comment),new(User))
	err := orm.RegisterDriver(sqliteDriver, orm.DRSqlite)
	if err != nil {
		return
	}
	err = orm.RegisterDataBase("default", sqliteDriver, dbName)
	if err != nil {
		return
	}

}
func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	Cats := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&Cats)
	return Cats, err
}
func GetAllTopics(cate ,label string, IsDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	Topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if IsDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label)>0{
			qs=qs.Filter("label__contains","#"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&Topics)
	} else {
		_, err = qs.All(&Topics)
	}

	return Topics, err
}
func AddTopic(title, content, category, label string) error {
	o := orm.NewOrm()
	label="#"+strings.Join(strings.Split(label," "),"#")+"#"
	topic := &Topic{
		Title:     title,
		Content:   content,
		Category:  category,
		Label:     label,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	UpdateCategory()
	return err
}
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	topic.Label=strings.Replace(
		strings.Replace(topic.Label,"#"," ",-1)," ","",1)
	return topic, err
}
func ModifyTopic(tid, title, content, category,label string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	label="#"+strings.Join(strings.Split(label," "),"#")+"#"
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Label=label
		topic.Updated = time.Now()
		topic.Category = category
		_, err := o.Update(topic)
		if err != nil {
			return err
		}
	}
	UpdateCategory()
	return nil
}
func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	UpdateCategory()
	return err
}
func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	UpdateReply()
	UpdateUser()
	return err
}
func GetAllReplies(tid string) ([]*Comment, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	replies := make([]*Comment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}
func DeleteReply(rid string) {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	o := orm.NewOrm()
	reply := &Comment{Id: ridNum}
	_, err = o.Delete(reply)
	if err != nil {
		panic(err)
	}
	UpdateReply()
	UpdateUser()
}
func UpdateCategory() {
	o := orm.NewOrm()
	categories, _ := GetAllCategories()
	for _, c := range categories {
		var err error
		topics, _ := GetAllTopics(c.Title,"", true)
		c.TopicCount = int64(len(topics))
		_, err = o.Update(c)
		if err != nil {
			panic(err)
		}
	}

}
func UpdateReply() {
	o := orm.NewOrm()
	topics, _ := GetAllTopics("", "",false)
	for _, t := range topics {
		var err error
		replies, _ := GetAllReplies(strconv.FormatInt(t.Id, 10))
		t.ReplyCount = int64(len(replies))
		_, err = o.Update(t)
		if err != nil {
			panic(err)
		}
	}

}
func UpdateUser(){
	o:=orm.NewOrm()
	qs:=o.QueryTable("comment")
	replies:=make([]*Comment, 0)
	users,_:=GetAllUsers()
	for _,u:=range users{
		_,err:=qs.Filter("name",u.NickName).All(&replies)
		if err != nil {
			panic(err)
		}
		u.Replies=int64(len(replies))
		_, err = o.Update(u)
	}
}
func AddUser(nickname,password string) error {
	o := orm.NewOrm()
	user:=&User{
		NickName: nickname,
		Password: password,
		Replies: 0,
	}
	qs := o.QueryTable("user")
	err := qs.Filter("nickname", nickname).One(user)
	if err == nil {
		return err
	}
	_, err = o.Insert(user)
	return nil
}
func UserLogin(nickname, password string) (*User,bool) {
	o := orm.NewOrm()
	user:=new(User)
	qs:=o.QueryTable("user")
	err:=qs.Filter("nickname",nickname).One(user)
	if err==nil{
		if password==user.Password{
			return user,true
		}
	}
	return user,false

}
func GetAllUsers()([]*User,error){
	o := orm.NewOrm()
	Users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.All(&Users)
	return Users, err
}
func GetUser(id string)(*User,error){
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	user:= new(User)
	qs := o.QueryTable("user")
	err = qs.Filter("id", idNum).One(user)
	if err != nil {
		return nil, err
	}
	return user, err
}
func CheckUser(ctx *context2.Context)(*User,bool){
		user:=new(User)
		ck,err:=ctx.Request.Cookie("nickname")
		if err!=nil{
			return user,false
		}
		nickname:=ck.Value
		ck,err=ctx.Request.Cookie("pwd")
		if err!=nil{
			return user,false
		}
		pwd:=ck.Value
		o:=orm.NewOrm()
		qs:=o.QueryTable("user")
		err=qs.Filter("nickname",nickname).One(user)
		if err != nil {
			return user,false
		}
		return user,nickname==user.NickName&&pwd==user.Password
}
