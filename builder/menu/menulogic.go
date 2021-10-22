func (l *AddMenuLogic) AddMenu(req types.AddMenuReq) (*types.AddMenuResp, error) {
	menuId, err := l.svcCtx.MenuModel.Insert(model.Menu{
		Channel:req.Channel
        ParentId:req.ParentId
        Name:req.Name
        Icon:req.Icon
        IconColor:req.IconColor
        UrlType:req.UrlType
        UrlPath:req.UrlPath
        Remark:req.Remark
        Sort:req.Sort
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddMenuResp{
		Id: menuId,
	}, nil
}

func (l *DelMenuLogic) DelMenu(req types.DelMenuReq) (*types.DelMenuResp, error) {
	err := l.svcCtx.MenuModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelMenuResp{
		Id: req.Id,
	}, nil
}

func (l *InfoMenuLogic) InfoMenu(req types.InfoMenuReq) (*types.InfoMenuResp, error) {
	Info, err := l.svcCtx.MenuModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoMenuResp{
		Id:Info.Id
        Channel:Info.Channel
        ParentId:Info.ParentId
        Name:Info.Name
        Icon:Info.Icon
        IconColor:Info.IconColor
        UrlType:Info.UrlType
        UrlPath:Info.UrlPath
        Remark:Info.Remark
        Sort:Info.Sort
        DeleteTime:Info.DeleteTime
        CreateTime:Info.CreateTime
        UpdateTime:Info.UpdateTime
        
	}, nil
}

func (l *ListMenuLogic) ListMenu(req types.ListMenuReq) (*types.ListMenuResp, error) {
	list, total, err := l.svcCtx.MenuModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	menuList := make([]types.Menu, 0, len(list))
	for _, v := range list {
		menuList = append(menuList, types.Menu{
			Id:v.Id
            Channel:v.Channel
            ParentId:v.ParentId
            Name:v.Name
            Icon:v.Icon
            IconColor:v.IconColor
            UrlType:v.UrlType
            UrlPath:v.UrlPath
            Remark:v.Remark
            Sort:v.Sort
            DeleteTime:v.DeleteTime
            CreateTime:v.CreateTime
            UpdateTime:v.UpdateTime
            
		})
	}
	return &types.ListMenuResp {
		MenuList: menuList,
		Total:    total,
	}, nil
}


func (l *UpdateMenuLogic) UpdateMenu(req types.UpdateMenuReq) (*types.UpdateMenuResp, error) {
	menuInfo, err := l.svcCtx.MenuModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.MenuModel.Update(*menuInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateMenuResp{
		Id: req.Id,
	}, nil
}