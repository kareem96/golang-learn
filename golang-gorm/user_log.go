package golanggorm

type UserLog struct {
	ID int `gorm:"primary_key;column:id;autoIncrement"`
	UserID string `gorm:"column:user_id"`
	Action string `gorm:"column:action"`
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"column:updated_at;autoCreateTime;autoUpdateTime:milli"`
}

func (user *UserLog) TableName()string  {
	return "users_logs"
}