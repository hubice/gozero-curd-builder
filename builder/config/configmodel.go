func (m *ConfigModel) Insert(data Config) (int64, error) {
	insert := builder.Eq{
        "id": data.Id,
        "brand_id": data.BrandId,
        "channel": data.Channel,
        "key": data.Key,
        "val": data.Val,
        "create_time": data.CreateTime,
        
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

func (m *ConfigModel) List(page, pageSize int) ([]*Config, int64, error) {
	where := builder.Eq{}

	var total int64
	resp := make([]*Config, 0)
	err := mr.Finish(func() error {
		offset := (page - 1) * pageSize
		query, args, err := builder.Dialect(builder.MYSQL).Select(configRows).From(m.table).Where(where).Limit(pageSize, offset).ToSQL()
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


func (m *ConfigModel) InsertAll(list []*Config) error {
	insert := builder.Eq{
        "id": "",
        "brand_id": "",
        "channel": "",
        "key": "",
        "val": "",
        "create_time": "",
        
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
            "channel": data.Channel,
            "key": data.Key,
            "val": data.Val,
            "create_time": data.CreateTime,
            
		}
		_, args, err := builder.Insert(insert).Into(m.table).ToSQL()
		err = bulk.Insert(args...)
		if err != nil {
			logx.Errorf("called ConfigModel InsertAll exec: %s", err)
			continue
		}
		if i != 0 && i%1000 == 0 {
			bulk.Flush()
		}
	}
	bulk.Flush()
	return nil
}

func (m *ConfigModel) Update(data Config) error {
	update := builder.Eq{
        "id": data.Id,
        "brand_id": data.BrandId,
        "channel": data.Channel,
        "key": data.Key,
        "val": data.Val,
        "create_time": data.CreateTime,
        
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