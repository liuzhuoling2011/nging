// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type AccessLog struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*AccessLog
	namer   func(string) string
	connID  int
	
	Id           	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	VhostId      	uint    	`db:"vhost_id" bson:"vhost_id" comment:"虚拟主机ID" json:"vhost_id" xml:"vhost_id"`
	RemoteAddr   	string  	`db:"remote_addr" bson:"remote_addr" comment:"IP地址" json:"remote_addr" xml:"remote_addr"`
	XRealIp      	string  	`db:"x_real_ip" bson:"x_real_ip" comment:"IP" json:"x_real_ip" xml:"x_real_ip"`
	XForwardFor  	string  	`db:"x_forward_for" bson:"x_forward_for" comment:"IP" json:"x_forward_for" xml:"x_forward_for"`
	LocalAddr    	string  	`db:"local_addr" bson:"local_addr" comment:"本机地址" json:"local_addr" xml:"local_addr"`
	Elapsed      	float64 	`db:"elapsed" bson:"elapsed" comment:"耗时(秒)" json:"elapsed" xml:"elapsed"`
	Host         	string  	`db:"host" bson:"host" comment:"Header中的Host, 一般会是域名" json:"host" xml:"host"`
	User         	string  	`db:"user" bson:"user" comment:"用户" json:"user" xml:"user"`
	TimeLocal    	string  	`db:"time_local" bson:"time_local" comment:"本地时间格式" json:"time_local" xml:"time_local"`
	Minute       	string  	`db:"minute" bson:"minute" comment:"H:i" json:"minute" xml:"minute"`
	Method       	string  	`db:"method" bson:"method" comment:"GET POST HEAD OPTIONS PUT DELETE..." json:"method" xml:"method"`
	Uri          	string  	`db:"uri" bson:"uri" comment:"网址" json:"uri" xml:"uri"`
	Version      	string  	`db:"version" bson:"version" comment:"HTTP/1.0, HTTP/1.1 ..." json:"version" xml:"version"`
	StatusCode   	uint    	`db:"status_code" bson:"status_code" comment:"状态码" json:"status_code" xml:"status_code"`
	BodyBytes    	uint64  	`db:"body_bytes" bson:"body_bytes" comment:"响应body字节数" json:"body_bytes" xml:"body_bytes"`
	Referer      	string  	`db:"referer" bson:"referer" comment:"来源网址" json:"referer" xml:"referer"`
	UserAgent    	string  	`db:"user_agent" bson:"user_agent" comment:"用户代理" json:"user_agent" xml:"user_agent"`
	HitStatus    	uint    	`db:"hit_status" bson:"hit_status" comment:"缓存服务器命中状态" json:"hit_status" xml:"hit_status"`
	Scheme       	string  	`db:"scheme" bson:"scheme" comment:"https/http" json:"scheme" xml:"scheme"`
	BrowerName   	string  	`db:"brower_name" bson:"brower_name" comment:"浏览器名" json:"brower_name" xml:"brower_name"`
	BrowerType   	string  	`db:"brower_type" bson:"brower_type" comment:"浏览器类型(spider/mobile/pc)" json:"brower_type" xml:"brower_type"`
	Created      	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
}

func (this *AccessLog) Trans() *factory.Transaction {
	return this.trans
}

func (this *AccessLog) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *AccessLog) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *AccessLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *AccessLog) Objects() []*AccessLog {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *AccessLog) NewObjects() *[]*AccessLog {
	this.objects = []*AccessLog{}
	return &this.objects
}

func (this *AccessLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *AccessLog) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *AccessLog) Short_() string {
	return "access_log"
}

func (this *AccessLog) Struct_() string {
	return "AccessLog"
}

func (this *AccessLog) Name_() string {
	if this.namer != nil {
		return this.namer(this.Short_())
	}
	return factory.TableNamerGet(this.Short_())(this)
}

func (this *AccessLog) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *AccessLog) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *AccessLog) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *AccessLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *AccessLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *AccessLog) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.TimeLocal) == 0 { this.TimeLocal = "1970-01-01 00:00:00" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return
}

func (this *AccessLog) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.TimeLocal) == 0 { this.TimeLocal = "1970-01-01 00:00:00" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *AccessLog) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *AccessLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *AccessLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["time_local"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["time_local"] = "1970-01-01 00:00:00" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *AccessLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.TimeLocal) == 0 { this.TimeLocal = "1970-01-01 00:00:00" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.TimeLocal) == 0 { this.TimeLocal = "1970-01-01 00:00:00" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return 
}

func (this *AccessLog) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *AccessLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *AccessLog) Reset() *AccessLog {
	this.Id = 0
	this.VhostId = 0
	this.RemoteAddr = ``
	this.XRealIp = ``
	this.XForwardFor = ``
	this.LocalAddr = ``
	this.Elapsed = 0.0
	this.Host = ``
	this.User = ``
	this.TimeLocal = ``
	this.Minute = ``
	this.Method = ``
	this.Uri = ``
	this.Version = ``
	this.StatusCode = 0
	this.BodyBytes = 0
	this.Referer = ``
	this.UserAgent = ``
	this.HitStatus = 0
	this.Scheme = ``
	this.BrowerName = ``
	this.BrowerType = ``
	this.Created = 0
	return this
}

func (this *AccessLog) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["VhostId"] = this.VhostId
	r["RemoteAddr"] = this.RemoteAddr
	r["XRealIp"] = this.XRealIp
	r["XForwardFor"] = this.XForwardFor
	r["LocalAddr"] = this.LocalAddr
	r["Elapsed"] = this.Elapsed
	r["Host"] = this.Host
	r["User"] = this.User
	r["TimeLocal"] = this.TimeLocal
	r["Minute"] = this.Minute
	r["Method"] = this.Method
	r["Uri"] = this.Uri
	r["Version"] = this.Version
	r["StatusCode"] = this.StatusCode
	r["BodyBytes"] = this.BodyBytes
	r["Referer"] = this.Referer
	r["UserAgent"] = this.UserAgent
	r["HitStatus"] = this.HitStatus
	r["Scheme"] = this.Scheme
	r["BrowerName"] = this.BrowerName
	r["BrowerType"] = this.BrowerType
	r["Created"] = this.Created
	return r
}

func (this *AccessLog) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["vhost_id"] = this.VhostId
	r["remote_addr"] = this.RemoteAddr
	r["x_real_ip"] = this.XRealIp
	r["x_forward_for"] = this.XForwardFor
	r["local_addr"] = this.LocalAddr
	r["elapsed"] = this.Elapsed
	r["host"] = this.Host
	r["user"] = this.User
	r["time_local"] = this.TimeLocal
	r["minute"] = this.Minute
	r["method"] = this.Method
	r["uri"] = this.Uri
	r["version"] = this.Version
	r["status_code"] = this.StatusCode
	r["body_bytes"] = this.BodyBytes
	r["referer"] = this.Referer
	r["user_agent"] = this.UserAgent
	r["hit_status"] = this.HitStatus
	r["scheme"] = this.Scheme
	r["brower_name"] = this.BrowerName
	r["brower_type"] = this.BrowerType
	r["created"] = this.Created
	return r
}

func (this *AccessLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *AccessLog) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

