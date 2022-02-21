package dao

import (
	"database/sql"
	"log"
	"shoping/model"
	"time"
)

type UserDao struct {
	*sql.DB
}
func(ud *UserDao)UpdateBalance(change ,uid int64)error{
	stmt,err:=ud.Prepare("update user set balance = balance + ? where id = ?")
	defer stmt.Close()
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	_,err=stmt.Exec(change,uid)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil

}
func(ud *UserDao)QueryUserInfoById(id int64)(model.User,error){
	user:=model.User{}
	stmt,err:=ud.Prepare("select id, username, password, email, phone, salt, avatar,reg_date,statement,gender,balance from user where id = ?")
	defer stmt.Close()
	if err!=nil{
		log.Fatal(err.Error())
		return user,nil
	}
	row:=stmt.QueryRow(id)
	if row.Err()!=nil{
		log.Fatal(err.Error())
		return user,nil
	}
	var regdate time.Time
	err=row.Scan(&user.Id,&user.Username,&user.Password,&user.Email,&user.Phone,&user.Salt,&user.Avatar,&user.RegDate,&user.Statement,&user.Gender,&user.Balance)
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	user.RegDate=regdate
	return user,nil


}
func(ud *UserDao)UpdatePwd(id int64,salt string,pwd string)(model.User,error){
	user:=model.User{}
	stmt,err:=ud.Prepare("update user set password = ?,salt = ? where id = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	defer stmt.Close()
	_, err = stmt.Exec(pwd,salt,id)
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	return user,nil

}
func(ud *UserDao)SelectIdByEmail(email string)(model.User,error){

		user:=model.User{}
		stmt,err:=ud.Prepare("select id, username, password, email, phone, salt, avatar,reg_date,statement,gender,balance from user where id = ?  ")
		if err!=nil{
			log.Fatal(err.Error())
			return user,err
		}
		defer stmt.Close()
		row:=stmt.QueryRow(email)
		if row.Err()!=nil{
			log.Fatal(row.Err().Error())
			return user,err
		}
		err=row.Scan(&user.Id,&user.Username,&user.Password,&user.Email,&user.Phone,&user.Salt,&user.Avatar,&user.RegDate,&user.Statement,&user.Gender,&user.Balance)
		if err!=nil{
			log.Fatal(err.Error())
			return user,err
		}
		return user,nil


	}

func(ud *UserDao)SelectIdByPhone(phone string)(model.User,error){
	user:=model.User{}
	stmt,err:=ud.Prepare("select id, username, password, email, phone, salt, avatar,reg_date,statement,gender,banlance from user where id = ?  ")
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	defer stmt.Close()
	row:=stmt.QueryRow(phone)
	if row.Err()!=nil{
		log.Fatal(row.Err().Error())
		return user,err
	}
	err=row.Scan(&user.Id,&user.Username,&user.Password,&user.Email,&user.Phone,&user.Salt,&user.Avatar,&user.RegDate,&user.Statement,&user.Gender,&user.Balance)
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	return user,nil


}
func(ud *UserDao)ChangeUserNameById(id int64,newUserName string)error{
	stmt,err:=ud.Prepare("update user set username = ? where id = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(id,newUserName)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil



}
func (ud *UserDao)ChangeGenderById(id int64,newGender string)error{
	stmt,err:=ud.Prepare("update user set gender = ? where id = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_,err=stmt.Exec(id,newGender)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil




}
func (ud *UserDao)ChangePhoneById(id int64,newPhone string)error{

	stmt,err:=ud.Prepare("update user set phone = ? where id = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer  stmt.Close()
	_, err = stmt.Exec(id, newPhone)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil
}

// ChangeEmailById 修改邮箱
func (ud *UserDao)ChangeEmailById(id int64,newEmail string)error{

	stmt,err:=ud.Prepare("update user set email = ? where id =?")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer  stmt.Close()
	_, err = stmt.Exec(newEmail, id)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil

}


// QueryByEmail 根据邮箱查询
func (ud *UserDao)QueryByEmail(email string)(model.User,error){
	user:=model.User{}
	stmt,err:=ud.DB.Prepare("select id, username, password, email, phone, salt, avatar,reg_date,statement,gender,balance from user where email=?")
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	defer stmt.Close()
	row:=stmt.QueryRow(email)
	if row.Err()!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	err=row.Scan(&user.Id,&user.Username,&user.Password,&user.Email,&user.Phone,&user.Salt,&user.Avatar,&user.RegDate,&user.Statement,&user.Gender,&user.Balance)
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	return user,nil


}

// QueryByUsername 根据用户名查询
func(ud *UserDao)QueryByUsername(username string)(model.User,error){
	user:=model.User{}
	stmt,err:=ud.DB.Prepare("select id, username, password, email, phone, salt, avatar,reg_date,statement,gender,balance from user where username = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	defer stmt.Close()
	row:=stmt.QueryRow(username)
	if row.Err()!=nil{
		return user,row.Err()
	}
	err=row.Scan(&user.Id,&user.Username,&user.Password,&user.Email,&user.Phone,&user.Salt,&user.Avatar,&user.RegDate,&user.Statement,&user.Gender,&user.Balance)
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	return user,nil
}

// QueryByPhone 根据手机号查询
func (ud *UserDao)QueryByPhone(phone string)(model.User,error){
	//防止sql注入
	user:=model.User{}
	stmt,err:=ud.DB.Prepare("select id, username, password, email, phone, salt, avatar,reg_date,statement,gender,balance from user where phone = ?")
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	defer stmt.Close()
	row:=stmt.QueryRow(phone)
	if row.Err()!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	err=row.Scan(&user.Id,&user.Username,&user.Password,&user.Email,&user.Phone,&user.Salt,&user.Avatar,&user.RegDate,&user.Statement,&user.Gender,&user.Balance)
	if err!=nil{
		log.Fatal(err.Error())
		return user,err
	}
	return user,nil

}
func(ud *UserDao)InsertUser(user model.User)error{
	stmt,err:=ud.DB.Prepare("insert into user(id,username, password, reg_date, phone, salt,balance) VALUES (?, ?, ?, ?, ?, ?,?) ")
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id,user.Username, user.Password, user.RegDate, user.Phone, user.Salt,user.Balance)
	if err!=nil{
		log.Fatal(err.Error())
		return err
	}
	return nil

}
