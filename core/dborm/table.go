package dborm

// 用户

type User struct {
	Id          uint   `gorm:"primaryKey"`
	Username    string `gorm:"index,unique"`
	Password    string
	Description string `gorm:"default:什么也没有"`
	Secrets     []Secret
	Sessions    []Session
	CreatedAt   int64
	UpdatedAt   int64
}

// 用户会话

type Session struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint `gorm:"index"`
	User      User
	Token     string `gorm:"index,unique"`
	CreatedAt int64
	UpdatedAt int64
}

// CAM 密钥

type Secret struct {
	Id          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"index"`
	SecretId    string `gorm:"index,unique"`
	SecretKey   string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// SSH 主机

type SSHHost struct {
	Id          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"index"`
	Address     string `gorm:"index,unique"`
	Username    string
	Password    string
	Description string
	SSHKeyId    uint
	SSHKey      SSHKey
	CreatedAt   int64
	UpdatedAt   int64
}

// SSH 密钥

type SSHKey struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	PublicKey   string
	PrivateKey  string
	Description string
	SSHHost     []SSHHost
	CreatedAt   int64
	UpdatedAt   int64
}

// 自动化助手 脚本

type TATScript struct {
	Id               uint `gorm:"primaryKey"`
	UserId           uint `gorm:"index"`
	Name             string
	Username         string
	Content          string
	Description      string
	CommandType      string
	WorkingDirectory string
	Timeout          uint
	CreatedAt        int64
	UpdatedAt        int64
}

// 自动化助手 历史记录

type TATHistory struct {
	Id                   uint `gorm:"primaryKey"`
	UserId               uint `gorm:"index"`
	KeyId                uint
	Name                 string
	Region               string
	InvocationId         string
	InvocationStatus     string
	InvocationResultJson string
}
