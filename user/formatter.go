package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Token:    token,
		ImageURL: user.AvatarFileName,
	}
	return formatter
}

func FormatUserDetail(user User) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		ImageURL: user.AvatarFileName,
	}
	return formatter
}

func FormatUserEdit(user User) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		ImageURL: user.AvatarFileName,
	}
	return formatter
}

type UserUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageURL string `json:"image_url"`
}

func FormatUserUpdate(user User) UserUpdate {
	formatter := UserUpdate{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		ImageURL: user.AvatarFileName,
	}
	return formatter
}
