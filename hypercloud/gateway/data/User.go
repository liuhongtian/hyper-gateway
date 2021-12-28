package data

type User struct {
	Name      string
	Gender    string
	LoginName string
	Email     string
	CellPhone string
	Wechat    string
}

func (u User) ToString() string {
	return "Name:" + u.Name + "; Gender:" + u.Gender + "; LoginName:" + u.LoginName + "; Email:" + u.Email + "; CellPhone:" + u.CellPhone + "; WeChat:" + u.Wechat
}
