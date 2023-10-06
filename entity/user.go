package entity

type User struct {
	Email    string
	Password string
	Name     string
}

// Có thể chứa các hàm tương tác với user như
// Hàm tạo user, ví dụ sẽ có hàm mã hóa password, ...
// Hàm kiểm tra password có đúng ko, dùng khi đăng nhập

func NewUser(email, password, name string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
}

func (u *User) CheckPassword(p string) bool {
	return u.Password == p
}