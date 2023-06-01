package pojo

// User 用户表
type User struct {
	ID       int32  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;type:int"`      //用户id 主键 自增 int
	Account  string `json:"account,omitempty" gorm:"not null;unique;index;type:varchar(45)"` //账号 非空 唯一 索引 varchar
	Password string `json:"password,omitempty" gorm:"not null;type:varchar(255)"`            //密码 非空 varchar
	Avatar   string `json:"avatar,omitempty" gorm:"null;type:varchar(255)"`                  //头像 空 varchar
}

// Email 邮箱表
type Email struct {
	UserID int32  `json:"user_id,omitempty" gorm:"primaryKey;autoIncrement:false;not null;type:int"` //用户id 主键 非自增 非空 int
	Email  string `json:"email,omitempty" gorm:"not null;unique;type:varchar(255)"`                  //用户邮箱 非空 唯一 varchar(255)
}

// Login 登录表 记录用户登录信息
type Login struct {
	UserID        int32  `json:"user_id,omitempty" gorm:"primaryKey;autoIncrement:false;not null;type:int"` //用户id 主键 非自增 非空 int
	Token         string `json:"token,omitempty" gorm:"not null;unique;type:char(64)"`                      //token 非空 char(64)
	LastLoginTime string `json:"last_login_time,omitempty" gorm:"not null;unique;type:char(64)"`            //最后登录时间
	LastLoginIP   string `json:"last_login_ip,omitempty" gorm:"not null;unique;type:varchar(15)"`           //最后登录ip
}

// Photo 图片表
type Photo struct {
	ID       int32  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;type:int"`       //相册id 主键 自增 int
	UserID   int32  `json:"user_id,omitempty" gorm:"primaryKey;autoIncrement:false;type:int"` //所属用户id 主键 int
	Path     string `json:"desc,omitempty" gorm:"not null;type:varchar(255)"`                 //图片路径 非空 varchar
	CreateAt int64  `json:"create_at,omitempty"`                                              //创建时间
	UpdateAt int64  `json:"update_at,omitempty"`                                              //创建时间
}

// Album 相册表
type Album struct {
	ID       int32  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;type:int"`       //相册id 主键 自增 int
	UserID   int32  `json:"user_id,omitempty" gorm:"primaryKey;autoIncrement:false;type:int"` //所属用户id 主键 int
	Title    string `json:"title,omitempty" gorm:"not null;type:varchar(28)"`                 //相册标题 非空 varchar
	Desc     string `json:"desc,omitempty" gorm:"null;type:varchar(255)"`                     //相册描述 空 varchar
	CreateAt int64  `json:"create_at,omitempty"`                                              //创建时间
	UpdateAt int64  `json:"update_at,omitempty"`                                              //创建时间
}

// PhotosInAlbum 照片相册关联表
type PhotosInAlbum struct {
}

// Link 外链表
type Link struct {
}
