package IAM

func (s *SQLRepo)ValidateCred(userId string, password string) (int,string) {
	results,err:=s.db.Query("Select role from users where userId = ? and password = ?",userId,password)
	if err!= nil{
		panic(err.Error())
	}
	re:=0
	var role string
	for results.Next() {
		re++
		results.Scan(&role)
	}
	return re, role
}

func (s *SQLRepo)ResetPassword(userId int, password string, newPassword string) string {
	results,err:=s.db.Exec("UPDATE users set password = ? where userId = ? and password = ?",newPassword,userId,password)
	if err!= nil{
		panic(err.Error())
	}
	row, err:= results.RowsAffected()
	if row == 0{
		return "Invalid credentials"
	}
	return ""
}
