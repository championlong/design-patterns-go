package isp

type UserService interface {
	register()
	login()
}

type RestrictedUserService interface {
	deleteUserById(userId int64)
}

type Statistics struct {
}

func (receiver Statistics) count()  {}

type Updater interface {
	update()
}

type Viewer interface {
	outPut()
}

type MysqlConfig struct {}

func (self *MysqlConfig)update()  {}

func (self *MysqlConfig)outPut()  {}

type PgSqlConfig struct {}

func (self *PgSqlConfig)update()  {}