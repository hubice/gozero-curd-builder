func (l *AddAccountDepartmentReLogic) AddAccountDepartmentRe(req types.AddAccountDepartmentReReq) (*types.AddAccountDepartmentReResp, error) {
	accountDepartmentReId, err := l.svcCtx.AccountDepartmentReModel.Insert(model.AccountDepartmentRe{
		BrandId:req.BrandId
        DepartmentId:req.DepartmentId
        AccountId:req.AccountId
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddAccountDepartmentReResp{
		Id: accountDepartmentReId,
	}, nil
}

func (l *DelAccountDepartmentReLogic) DelAccountDepartmentRe(req types.DelAccountDepartmentReReq) (*types.DelAccountDepartmentReResp, error) {
	err := l.svcCtx.AccountDepartmentReModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelAccountDepartmentReResp{
		Id: req.Id,
	}, nil
}

func (l *InfoAccountDepartmentReLogic) InfoAccountDepartmentRe(req types.InfoAccountDepartmentReReq) (*types.InfoAccountDepartmentReResp, error) {
	Info, err := l.svcCtx.AccountDepartmentReModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoAccountDepartmentReResp{
		Id:Info.Id
        BrandId:Info.BrandId
        DepartmentId:Info.DepartmentId
        AccountId:Info.AccountId
        CreateTime:Info.CreateTime
        
	}, nil
}

func (l *ListAccountDepartmentReLogic) ListAccountDepartmentRe(req types.ListAccountDepartmentReReq) (*types.ListAccountDepartmentReResp, error) {
	list, total, err := l.svcCtx.AccountDepartmentReModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	accountDepartmentReList := make([]types.AccountDepartmentRe, 0, len(list))
	for _, v := range list {
		accountDepartmentReList = append(accountDepartmentReList, types.AccountDepartmentRe{
			Id:v.Id
            BrandId:v.BrandId
            DepartmentId:v.DepartmentId
            AccountId:v.AccountId
            CreateTime:v.CreateTime
            
		})
	}
	return &types.ListAccountDepartmentReResp {
		AccountDepartmentReList: accountDepartmentReList,
		Total:    total,
	}, nil
}


func (l *UpdateAccountDepartmentReLogic) UpdateAccountDepartmentRe(req types.UpdateAccountDepartmentReReq) (*types.UpdateAccountDepartmentReResp, error) {
	accountDepartmentReInfo, err := l.svcCtx.AccountDepartmentReModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.AccountDepartmentReModel.Update(*accountDepartmentReInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateAccountDepartmentReResp{
		Id: req.Id,
	}, nil
}