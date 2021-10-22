func (m *PositionModel) Insert(data Position) (int64, error) {
	insert := builder.Eq{
        "id": data.Id,
        "brand_id": data.BrandId,
        "name": data.Name,
        "remark": data.Remark,
        "delete_time": data.DeleteTime,
        "create_time": data.CreateTime,
        "update_time": data.UpdateTime,
        
	}
	toSQL, i, err := builder.Insert(insert).Into(m.table).ToSQL()
	if err != nil {
		return 0, err
	}
	ret, err := m.conn.Exec(toSQL, i...)
	if err != nil {
		return 0, err
	}
	return ret.LastInsertId()
}

func (m *PositionModel) List(page, pageSize int) ([]*Position, int64, error) {
	where := builder.Eq{}

	var total int64
	resp := make([]*Position, 0)
	err := mr.Finish(func() error {
		offset := (page - 1) * pageSize
		query, args, err := builder.Dialect(builder.MYSQL).Select(positionRows).From(m.table).Where(where).Limit(pageSize, offset).ToSQL()
		if err != nil {
			return err
		}
		return m.conn.QueryRows(&resp, query, args...)
	}, func() error {
		query, args, err := builder.Select("COUNT(*) as total").From(m.table).Where(where).ToSQL()
		if err != nil {
			return err
		}
		return m.conn.QueryRow(&total, query, args...)
	})
	return resp, total, err
}


func (m *PositionModel) InsertAll(list []*Position) error {
	insert := builder.Eq{
        "id": "",
        "brand_id": "",
        "name": "",
        "remark": "",
        "delete_time": "",
        "create_time": "",
        "update_time": "",
        
	}
	query, _, err := builder.Insert(insert).Into(m.table).ToSQL()
	if err != nil {
		return err
	}
	var bulk *sqlx.BulkInserter
	bulk, _ = sqlx.NewBulkInserter(m.conn, query) //注意：批量插入不能重复实例化
	for i, data := range list {
		insert := builder.Eq{
	        "id": data.Id,
            "brand_id": data.BrandId,
            "name": data.Name,
            "remark": data.Remark,
            "delete_time": data.DeleteTime,
            "create_time": data.CreateTime,
            "update_time": data.UpdateTime,
            
		}
		_, args, err := builder.Insert(insert).Into(m.table).ToSQL()
		err = bulk.Insert(args...)
		if err != nil {
			logx.Errorf("called PositionModel InsertAll exec: %s", err)
			continue
		}
		if i != 0 && i%1000 == 0 {
			bulk.Flush()
		}
	}
	bulk.Flush()
	return nil
}

func (m *PositionModel) Update(data Position) error {
	update := builder.Eq{
        "id": data.Id,
        "brand_id": data.BrandId,
        "name": data.Name,
        "remark": data.Remark,
        "delete_time": data.DeleteTime,
        "create_time": data.CreateTime,
        "update_time": data.UpdateTime,
        
	}
	toSQL, i, err := builder.Update(update).From(m.table).Where(builder.Eq{"id": data.Id}).ToSQL()
	if err != nil {
		return err
	}
	_, err = m.conn.Exec(toSQL, i...)
	if err != nil {
		return err
	}
	return err
}