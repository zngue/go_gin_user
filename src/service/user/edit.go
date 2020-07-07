package user
import "github.com/zngue/go_gin_user/src/model"
func (Service) Edit (user model.User) error  {
	if user.ID>0 {
		return user.Update(user)
	}else{
		return user.Add(user)
	}
}
