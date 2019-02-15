package datastore

import (
	"examps/router-fw/connection"
	. "examps/router-fw/model"
)

func GetMember(id string) (Member, error){
	var member Member
	row := connection.Db.QueryRow("SELECT members.name, members.address, members.email FROM members where members.id = ?", id)
	err := row.Scan(&member.Name, &member.Address, &member.Email)
	if err != nil {
		return member, err
	}

	return member, nil
}

func GetMembers() (Members, error){
	var members Members

	rows, err := connection.Db.Query("SELECT * FROM members WHERE state=?", 1)
	if err != nil {
		return members, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for(rows.Next()){
		var member Member
		err := rows.Scan(
			&member.ID,
			&member.Name,
			&member.Address,
			&member.Email,
			)
		if err == nil {
			members.Data = append(members.Data, member)
		}
	}

	return members, nil
}

func AddMember(member Member)(int64, error){
	stmt, err := connection.Db.Prepare("INSERT INTO members(members.name, members.address, members.email) " +
		"VALUES(?,?,?")
	if err != nil{
		return -1, err
	}

	res, err := stmt.Exec(member.Name, member.Address, member.Email)
	if err != nil{
		return -1, err
	}

	return res.LastInsertId()
}
