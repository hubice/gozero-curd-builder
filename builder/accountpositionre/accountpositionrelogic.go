func (l *AddAccountPositionReLogic) AddAccountPositionRe(req types.AddAccountPositionReReq) (*types.AddAccountPositionReResp, error) {
	accountPositionReId, err := l.svcCtx.AccountPositionReModel.Insert(model.AccountPositionRe{
		BrandId:req.BrandId
        PositionId:req.PositionId
        AccountId:req.AccountId
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddAccountPositionReResp{
		Id: accountPositionReId,
	}, nil
}

func (l *DelAccountPositionReLogic) DelAccountPositionRe(req types.DelAccountPositionReReq) (*types.DelAccountPositionReResp, error) {
	err := l.svcCtx.AccountPositionReModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelAccountPositionReResp{
		Id: req.Id,
	}, nil
}

func (l *InfoAccountPositionReLogic) InfoAccountPositionRe(req types.InfoAccountPositionReReq) (*types.InfoAccountPositionReResp, error) {
	Info, err := l.svcCtx.AccountPositionReModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoAccountPositionReResp{
		Id:Info.Id
        BrandId:Info.BrandId
        PositionId:Info.PositionId
        AccountId:Info.AccountId
        CreateTime:Info.CreateTime
        
	}, nil
}

func (l *ListAccountPositionReLogic) ListAccountPositionRe(req types.ListAccountPositionReReq) (*types.ListAccountPositionReResp, error) {
	list, total, err := l.svcCtx.AccountPositionReModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	accountPositionReList := make([]types.AccountPositionRe, 0, len(list))
	for _, v := range list {
		accountPositionReList = append(accountPositionReList, types.AccountPositionRe{
			Id:v.Id
            BrandId:v.BrandId
            PositionId:v.PositionId
            AccountId:v.AccountId
            CreateTime:v.CreateTime
            
		})
	}
	return &types.ListAccountPositionReResp {
		AccountPositionReList: accountPositionReList,
		Total:    total,
	}, nil
}


func (l *UpdateAccountPositionReLogic) UpdateAccountPositionRe(req types.UpdateAccountPositionReReq) (*types.UpdateAccountPositionReResp, error) {
	accountPositionReInfo, err := l.svcCtx.AccountPositionReModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.AccountPositionReModel.Update(*accountPositionReInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateAccountPositionReResp{
		Id: req.Id,
	}, nil
}