func (m *RoleMenuReModel) Insert(data RoleMenuRe) (int64, error) {
	insert := builder.Eq{
        "id": data.Id,
        "brand_id": data.BrandId,
        "role_id": data.RoleId,
        "menu_id": data.MenuId,
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

func (m *RoleMenuReModel) List(page, pageSize int) ([]*RoleMenuRe, int64, error) {
	where := builder.Eq{}

	var total int64
	resp := make([]*RoleMenuRe, 0)
	err := mr.Finish(func() error {
		offset := (page - 1) * pageSize
		query, args, err := builder.Dialect(builder.MYSQL).Select(roleMenuReRows).From(m.table).Where(where).Limit(pageSize, offset).ToSQL()
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


func (m *RoleMenuReModel) InsertAll(list []*RoleMenuRe) error {
	insert := builder.Eq{
        "id": "",
        "brand_id": "",
        "role_id": "",
        "menu_id": "",
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
            "role_id": data.RoleId,
            "menu_id": data.MenuId,
            "create_time": data.CreateTime,
            
		}
		_, args, err := builder.Insert(insert).Into(m.table).ToSQL()
		err = bulk.Insert(args...)
		if err != nil {
			logx.Errorf("called RoleMenuReModel InsertAll exec: %s", err)
			continue
		}
		if i != 0 && i%1000 == 0 {
			bulk.Flush()
		}
	}
	bulk.Flush()
	return nil
}

func (m *RoleMenuReModel) Update(data RoleMenuRe) error {
	update := builder.Eq{
        "id": data.Id,
        "brand_id": data.BrandId,
        "role_id": data.RoleId,
        "menu_id": data.MenuId,
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