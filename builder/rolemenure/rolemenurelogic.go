func (l *AddRoleMenuReLogic) AddRoleMenuRe(req types.AddRoleMenuReReq) (*types.AddRoleMenuReResp, error) {
	roleMenuReId, err := l.svcCtx.RoleMenuReModel.Insert(model.RoleMenuRe{
		BrandId:req.BrandId
        RoleId:req.RoleId
        MenuId:req.MenuId
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddRoleMenuReResp{
		Id: roleMenuReId,
	}, nil
}

func (l *DelRoleMenuReLogic) DelRoleMenuRe(req types.DelRoleMenuReReq) (*types.DelRoleMenuReResp, error) {
	err := l.svcCtx.RoleMenuReModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelRoleMenuReResp{
		Id: req.Id,
	}, nil
}

func (l *InfoRoleMenuReLogic) InfoRoleMenuRe(req types.InfoRoleMenuReReq) (*types.InfoRoleMenuReResp, error) {
	Info, err := l.svcCtx.RoleMenuReModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoRoleMenuReResp{
		Id:Info.Id
        BrandId:Info.BrandId
        RoleId:Info.RoleId
        MenuId:Info.MenuId
        CreateTime:Info.CreateTime
        
	}, nil
}

func (l *ListRoleMenuReLogic) ListRoleMenuRe(req types.ListRoleMenuReReq) (*types.ListRoleMenuReResp, error) {
	list, total, err := l.svcCtx.RoleMenuReModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	roleMenuReList := make([]types.RoleMenuRe, 0, len(list))
	for _, v := range list {
		roleMenuReList = append(roleMenuReList, types.RoleMenuRe{
			Id:v.Id
            BrandId:v.BrandId
            RoleId:v.RoleId
            MenuId:v.MenuId
            CreateTime:v.CreateTime
            
		})
	}
	return &types.ListRoleMenuReResp {
		RoleMenuReList: roleMenuReList,
		Total:    total,
	}, nil
}


func (l *UpdateRoleMenuReLogic) UpdateRoleMenuRe(req types.UpdateRoleMenuReReq) (*types.UpdateRoleMenuReResp, error) {
	roleMenuReInfo, err := l.svcCtx.RoleMenuReModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RoleMenuReModel.Update(*roleMenuReInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateRoleMenuReResp{
		Id: req.Id,
	}, nil
}