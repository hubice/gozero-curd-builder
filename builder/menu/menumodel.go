func (m *MenuModel) Insert(data Menu) (int64, error) {
	insert := builder.Eq{
        "id": data.Id,
        "channel": data.Channel,
        "parent_id": data.ParentId,
        "name": data.Name,
        "icon": data.Icon,
        "icon_color": data.IconColor,
        "url_type": data.UrlType,
        "url_path": data.UrlPath,
        "remark": data.Remark,
        "sort": data.Sort,
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

func (m *MenuModel) List(page, pageSize int) ([]*Menu, int64, error) {
	where := builder.Eq{}

	var total int64
	resp := make([]*Menu, 0)
	err := mr.Finish(func() error {
		offset := (page - 1) * pageSize
		query, args, err := builder.Dialect(builder.MYSQL).Select(menuRows).From(m.table).Where(where).Limit(pageSize, offset).ToSQL()
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


func (m *MenuModel) InsertAll(list []*Menu) error {
	insert := builder.Eq{
        "id": "",
        "channel": "",
        "parent_id": "",
        "name": "",
        "icon": "",
        "icon_color": "",
        "url_type": "",
        "url_path": "",
        "remark": "",
        "sort": "",
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
            "channel": data.Channel,
            "parent_id": data.ParentId,
            "name": data.Name,
            "icon": data.Icon,
            "icon_color": data.IconColor,
            "url_type": data.UrlType,
            "url_path": data.UrlPath,
            "remark": data.Remark,
            "sort": data.Sort,
            "delete_time": data.DeleteTime,
            "create_time": data.CreateTime,
            "update_time": data.UpdateTime,
            
		}
		_, args, err := builder.Insert(insert).Into(m.table).ToSQL()
		err = bulk.Insert(args...)
		if err != nil {
			logx.Errorf("called MenuModel InsertAll exec: %s", err)
			continue
		}
		if i != 0 && i%1000 == 0 {
			bulk.Flush()
		}
	}
	bulk.Flush()
	return nil
}

func (m *MenuModel) Update(data Menu) error {
	update := builder.Eq{
        "id": data.Id,
        "channel": data.Channel,
        "parent_id": data.ParentId,
        "name": data.Name,
        "icon": data.Icon,
        "icon_color": data.IconColor,
        "url_type": data.UrlType,
        "url_path": data.UrlPath,
        "remark": data.Remark,
        "sort": data.Sort,
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