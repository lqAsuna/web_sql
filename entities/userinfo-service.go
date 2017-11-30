package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(u)
	checkErr(err)
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	res := make([]UserInfo, 0)
	err := myengine.Find(&res)
	checkErr(err)
	return res
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	u := new(UserInfo)
	_, err := engine.ID(id).Get(u)
	checkErr(err)
	return u
}
