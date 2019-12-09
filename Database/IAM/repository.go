package IAM

type Repository interface {
	ValidateCred(userId string, password string) (int,string)
	ResetPassword(userId int, password string, newPassword string) string
}
