package repostory

import "1-loyiha/model"



type UserRepostory struct {
	users []model.User
}




func NewUserRepostroy() *UserRepostory {
	return  &UserRepostory{}
}


func (r *UserRepostory) GetAll() []model.User {
	return r.users
}

func (r *UserRepostory) GetByID(id int) *model.User {
	for _, user := range r.users {
		if user.ID == id {
			return  &user
		}
	}

	return nil
}



func (r *UserRepostory) Create(user model.User) model.User{
	user.ID = len(r.users ) + 1
     r.users = append(r.users, user)
	 return user

}

func (r *UserRepostory) Updated(id int, updated model.User) *model.User{
	for i, user := range r.users {
		if user.ID == id {
			r.users[i].Name = updated.Name
			r.users[i].Email = updated.Email
		}
	}
	return  nil
}

func (r *UserRepostory) Delete(id int) bool {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return  true
		}
	}
	return  false
}