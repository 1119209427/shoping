package dao

import (
	"database/sql"
	"log"
	"shoping/model"
)

type CommonDao struct {
	*sql.DB
}
func(cd *CommonDao)QueryById(id int64)([]model.Comment,error) {

	var commentSlice []model.Comment
	//var commentTime time.Time
	stmt, err := cd.Prepare("select good_id,user_id,value,time,likes from where good_id = ?")
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var commentModel model.Comment
		err := rows.Scan( &commentModel.GoodId, &commentModel.UserId, &commentModel.Value, &commentModel.Time, &commentModel.Likes)
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}
		//commentModel.Time = commentTime.Format("2006-01-02 15:04:05")

		commentSlice = append(commentSlice, commentModel)


	}
	return commentSlice, nil

}
func(cd *CommonDao)InsertComment(comment model.Comment)(int64,error){
	//comment:=model.Comment{}
	stmt,err:=cd.Prepare("insert into comment(good_id,user_id,value,time,likes) values (?,?,?,?)")
	if err!=nil{
		log.Fatal(err.Error())
		return 0,err
	}
	defer stmt.Close()
	result,err:=stmt.Exec(&comment.GoodId,&comment.UserId,&comment.Value,&comment.Time,&comment.Likes)
	if err!=nil{
		log.Fatal(err.Error())
		return 0,err
	}
	id,err:=result.LastInsertId()
	if err!=nil{
		log.Fatal(err.Error())
		return 0,err
	}
	return id,nil
}
