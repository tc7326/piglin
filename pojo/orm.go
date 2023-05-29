package pojo

// User 用户表
type User struct {
	ID       int32  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;type:int"`      //用户id 主键 自增 int
	Account  string `json:"account,omitempty" gorm:"not null;unique;index;type:varchar(45)"` //账号 非空 唯一 索引 varchar
	Password string `json:"password,omitempty" gorm:"not null;type:varchar(255)"`            //密码 非空 varchar
	Avatar   string `json:"avatar,omitempty" gorm:"null;type:varchar(255)"`                  //头像 空 varchar
}

// Album 相册表
type Album struct {
	ID     int32  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;type:int"`       //相册id 主键 自增 int
	UserID int32  `json:"user_id,omitempty" gorm:"primaryKey;autoIncrement:false;type:int"` //所属用户id 主键 int
	Title  string `json:"title,omitempty" gorm:"not null;type:varchar(28)"`                 //相册标题 非空 varchar
	Desc   string `json:"desc,omitempty" gorm:"null;type:varchar(255)"`                     //相册描述 空 varchar
}

// Photo 图片表
type Photo struct {
	ID     int32  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;type:int"`       //相册id 主键 自增 int
	UserID int32  `json:"user_id,omitempty" gorm:"primaryKey;autoIncrement:false;type:int"` //所属用户id 主键 int
	Path   string `json:"desc,omitempty" gorm:"not null;type:varchar(255)"`                 //图片路径 非空 varchar
}

// Link 外链表
type Link struct {
}
