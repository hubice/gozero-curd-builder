func (l *AddLevelLogic) AddLevel(req types.AddLevelReq) (*types.AddLevelResp, error) {
	levelId, err := l.svcCtx.LevelModel.Insert(model.Level{
		BrandId:req.BrandId
        Code:req.Code
        Name:req.Name
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddLevelResp{
		Id: levelId,
	}, nil
}

func (l *DelLevelLogic) DelLevel(req types.DelLevelReq) (*types.DelLevelResp, error) {
	err := l.svcCtx.LevelModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelLevelResp{
		Id: req.Id,
	}, nil
}

func (l *InfoLevelLogic) InfoLevel(req types.InfoLevelReq) (*types.InfoLevelResp, error) {
	Info, err := l.svcCtx.LevelModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoLevelResp{
		Id:Info.Id
        BrandId:Info.BrandId
        Code:Info.Code
        Name:Info.Name
        DeleteTime:Info.DeleteTime
        CreateTime:Info.CreateTime
        UpdateTime:Info.UpdateTime
        
	}, nil
}

func (l *ListLevelLogic) ListLevel(req types.ListLevelReq) (*types.ListLevelResp, error) {
	list, total, err := l.svcCtx.LevelModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	levelList := make([]types.Level, 0, len(list))
	for _, v := range list {
		levelList = append(levelList, types.Level{
			Id:v.Id
            BrandId:v.BrandId
            Code:v.Code
            Name:v.Name
            DeleteTime:v.DeleteTime
            CreateTime:v.CreateTime
            UpdateTime:v.UpdateTime
            
		})
	}
	return &types.ListLevelResp {
		LevelList: levelList,
		Total:    total,
	}, nil
}


func (l *UpdateLevelLogic) UpdateLevel(req types.UpdateLevelReq) (*types.UpdateLevelResp, error) {
	levelInfo, err := l.svcCtx.LevelModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.LevelModel.Update(*levelInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateLevelResp{
		Id: req.Id,
	}, nil
}