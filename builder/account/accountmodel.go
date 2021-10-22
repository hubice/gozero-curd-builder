func (m *AccountModel) Insert(data Account) (int64, error) {
	insert := builder.Eq{
        "id": data.Id,
        "brand_id": data.BrandId,
        "name": data.Name,
        "avatar": data.Avatar,
        "gender": data.Gender,
        "email": data.Email,
        "mobile": data.Mobile,
        "no": data.No,
        "short_no": data.ShortNo,
        "entry_time": data.EntryTime,
        "introduce": data.Introduce,
        "level": data.Level,
        "department_ids": data.DepartmentIds,
        "role_ids": data.RoleIds,
        "position_ids": data.PositionIds,
        "shop_ids": data.ShopIds,
        "password": data.Password,
        "ding_talk_id": data.DingTalkId,
        "wei_xin_id": data.WeiXinId,
        "status": data.Status,
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

func (m *AccountModel) List(page, pageSize int) ([]*Account, int64, error) {
	where := builder.Eq{}

	var total int64
	resp := make([]*Account, 0)
	err := mr.Finish(func() error {
		offset := (page - 1) * pageSize
		query, args, err := builder.Dialect(builder.MYSQL).Select(accountRows).From(m.table).Where(where).Limit(pageSize, offset).ToSQL()
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


func (m *AccountModel) InsertAll(list []*Account) error {
	insert := builder.Eq{
        "id": "",
        "brand_id": "",
        "name": "",
        "avatar": "",
        "gender": "",
        "email": "",
        "mobile": "",
        "no": "",
        "short_no": "",
        "entry_time": "",
        "introduce": "",
        "level": "",
        "department_ids": "",
        "role_ids": "",
        "position_ids": "",
        "shop_ids": "",
        "password": "",
        "ding_talk_id": "",
        "wei_xin_id": "",
        "status": "",
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
            "avatar": data.Avatar,
            "gender": data.Gender,
            "email": data.Email,
            "mobile": data.Mobile,
            "no": data.No,
            "short_no": data.ShortNo,
            "entry_time": data.EntryTime,
            "introduce": data.Introduce,
            "level": data.Level,
            "department_ids": data.DepartmentIds,
            "role_ids": data.RoleIds,
            "position_ids": data.PositionIds,
            "shop_ids": data.ShopIds,
            "password": data.Password,
            "ding_talk_id": data.DingTalkId,
            "wei_xin_id": data.WeiXinId,
            "status": data.Status,
            "delete_time": data.DeleteTime,
            "create_time": data.CreateTime,
            "update_time": data.UpdateTime,
            
		}
		_, args, err := builder.Insert(insert).Into(m.table).ToSQL()
		err = bulk.Insert(args...)
		if err != nil {
			logx.Errorf("called AccountModel InsertAll exec: %s", err)
			continue
		}
		if i != 0 && i%1000 == 0 {
			bulk.Flush()
		}
	}
	bulk.Flush()
	return nil
}

func (m *AccountModel) Update(data Account) error {
	update := builder.Eq{
        "id": data.Id,
        "brand_id": data.BrandId,
        "name": data.Name,
        "avatar": data.Avatar,
        "gender": data.Gender,
        "email": data.Email,
        "mobile": data.Mobile,
        "no": data.No,
        "short_no": data.ShortNo,
        "entry_time": data.EntryTime,
        "introduce": data.Introduce,
        "level": data.Level,
        "department_ids": data.DepartmentIds,
        "role_ids": data.RoleIds,
        "position_ids": data.PositionIds,
        "shop_ids": data.ShopIds,
        "password": data.Password,
        "ding_talk_id": data.DingTalkId,
        "wei_xin_id": data.WeiXinId,
        "status": data.Status,
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