func (l *AddAccountLogic) AddAccount(req types.AddAccountReq) (*types.AddAccountResp, error) {
	accountId, err := l.svcCtx.AccountModel.Insert(model.Account{
		BrandId:req.BrandId
        Name:req.Name
        Avatar:req.Avatar
        Gender:req.Gender
        Email:req.Email
        Mobile:req.Mobile
        No:req.No
        ShortNo:req.ShortNo
        EntryTime:req.EntryTime
        Introduce:req.Introduce
        Level:req.Level
        DepartmentIds:req.DepartmentIds
        RoleIds:req.RoleIds
        PositionIds:req.PositionIds
        ShopIds:req.ShopIds
        Password:req.Password
        DingTalkId:req.DingTalkId
        WeiXinId:req.WeiXinId
        Status:req.Status
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddAccountResp{
		Id: accountId,
	}, nil
}

func (l *DelAccountLogic) DelAccount(req types.DelAccountReq) (*types.DelAccountResp, error) {
	err := l.svcCtx.AccountModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelAccountResp{
		Id: req.Id,
	}, nil
}

func (l *InfoAccountLogic) InfoAccount(req types.InfoAccountReq) (*types.InfoAccountResp, error) {
	Info, err := l.svcCtx.AccountModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoAccountResp{
		Id:Info.Id
        BrandId:Info.BrandId
        Name:Info.Name
        Avatar:Info.Avatar
        Gender:Info.Gender
        Email:Info.Email
        Mobile:Info.Mobile
        No:Info.No
        ShortNo:Info.ShortNo
        EntryTime:Info.EntryTime
        Introduce:Info.Introduce
        Level:Info.Level
        DepartmentIds:Info.DepartmentIds
        RoleIds:Info.RoleIds
        PositionIds:Info.PositionIds
        ShopIds:Info.ShopIds
        Password:Info.Password
        DingTalkId:Info.DingTalkId
        WeiXinId:Info.WeiXinId
        Status:Info.Status
        DeleteTime:Info.DeleteTime
        CreateTime:Info.CreateTime
        UpdateTime:Info.UpdateTime
        
	}, nil
}

func (l *ListAccountLogic) ListAccount(req types.ListAccountReq) (*types.ListAccountResp, error) {
	list, total, err := l.svcCtx.AccountModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	accountList := make([]types.Account, 0, len(list))
	for _, v := range list {
		accountList = append(accountList, types.Account{
			Id:v.Id
            BrandId:v.BrandId
            Name:v.Name
            Avatar:v.Avatar
            Gender:v.Gender
            Email:v.Email
            Mobile:v.Mobile
            No:v.No
            ShortNo:v.ShortNo
            EntryTime:v.EntryTime
            Introduce:v.Introduce
            Level:v.Level
            DepartmentIds:v.DepartmentIds
            RoleIds:v.RoleIds
            PositionIds:v.PositionIds
            ShopIds:v.ShopIds
            Password:v.Password
            DingTalkId:v.DingTalkId
            WeiXinId:v.WeiXinId
            Status:v.Status
            DeleteTime:v.DeleteTime
            CreateTime:v.CreateTime
            UpdateTime:v.UpdateTime
            
		})
	}
	return &types.ListAccountResp {
		AccountList: accountList,
		Total:    total,
	}, nil
}


func (l *UpdateAccountLogic) UpdateAccount(req types.UpdateAccountReq) (*types.UpdateAccountResp, error) {
	accountInfo, err := l.svcCtx.AccountModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.AccountModel.Update(*accountInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateAccountResp{
		Id: req.Id,
	}, nil
}