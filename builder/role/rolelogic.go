func (l *AddRoleLogic) AddRole(req types.AddRoleReq) (*types.AddRoleResp, error) {
	roleId, err := l.svcCtx.RoleModel.Insert(model.Role{
		BrandId:req.BrandId
        Name:req.Name
        Introduce:req.Introduce
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddRoleResp{
		Id: roleId,
	}, nil
}

func (l *DelRoleLogic) DelRole(req types.DelRoleReq) (*types.DelRoleResp, error) {
	err := l.svcCtx.RoleModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelRoleResp{
		Id: req.Id,
	}, nil
}

func (l *InfoRoleLogic) InfoRole(req types.InfoRoleReq) (*types.InfoRoleResp, error) {
	Info, err := l.svcCtx.RoleModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoRoleResp{
		Id:Info.Id
        BrandId:Info.BrandId
        Name:Info.Name
        Introduce:Info.Introduce
        DeleteTime:Info.DeleteTime
        CreateTime:Info.CreateTime
        UpdateTime:Info.UpdateTime
        
	}, nil
}

func (l *ListRoleLogic) ListRole(req types.ListRoleReq) (*types.ListRoleResp, error) {
	list, total, err := l.svcCtx.RoleModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	roleList := make([]types.Role, 0, len(list))
	for _, v := range list {
		roleList = append(roleList, types.Role{
			Id:v.Id
            BrandId:v.BrandId
            Name:v.Name
            Introduce:v.Introduce
            DeleteTime:v.DeleteTime
            CreateTime:v.CreateTime
            UpdateTime:v.UpdateTime
            
		})
	}
	return &types.ListRoleResp {
		RoleList: roleList,
		Total:    total,
	}, nil
}


func (l *UpdateRoleLogic) UpdateRole(req types.UpdateRoleReq) (*types.UpdateRoleResp, error) {
	roleInfo, err := l.svcCtx.RoleModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RoleModel.Update(*roleInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateRoleResp{
		Id: req.Id,
	}, nil
}