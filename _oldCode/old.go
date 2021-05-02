package _oldCode

//
//func (d IssueRepositorySql) Save(i Issue) (*Issue, *errs.AppError) {
//	sqlInsert := "INSERT INTO issues (name, description, createdAt, status, account_id) values (?,?,?,?,?)"
//
//	result, err := d.client.Exec(sqlInsert, i.Name, i.Description, i.CreatedAt, i.Status, i.AccountId)
//
//	if err != nil {
//		logger.Error("Error while creating issue: " + err.Error())
//		return nil, errs.NewUnexpectedError("Unexpected error from database")
//	}
//
//	id, err := result.LastInsertId()
//
//	i.IssueId = strconv.FormatInt(id, 10)
//	return &i, nil
//}

//
//
//func (d IssueRepositorySql) FindAll() (*[]Issue, *errs.AppError) {
//	//var rows *sql.Rows
//	var err error
//	issues := make([]Issue, 0)
//
//	findAllSql := "select * from issues"
//	err = d.client.Select(&issues, findAllSql)
//
//	if err != nil {
//		logger.Error("Error while quering issues from database " + err.Error())
//		return nil, errs.NewUnexpectedError("Unexpected database error")
//	}
//
//	//err = sqlx.StructScan(rows, &customers)
//
//	return &issues, nil
//}
