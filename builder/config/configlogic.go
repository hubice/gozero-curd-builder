func (l *AddConfigLogic) AddConfig(req types.AddConfigReq) (*types.AddConfigResp, error) {
	configId, err := l.svcCtx.ConfigModel.Insert(model.Config{
		BrandId:req.BrandId
        Channel:req.Channel
        Key:req.Key
        Val:req.Val
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddConfigResp{
		Id: configId,
	}, nil
}

func (l *DelConfigLogic) DelConfig(req types.DelConfigReq) (*types.DelConfigResp, error) {
	err := l.svcCtx.ConfigModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelConfigResp{
		Id: req.Id,
	}, nil
}

func (l *InfoConfigLogic) InfoConfig(req types.InfoConfigReq) (*types.InfoConfigResp, error) {
	Info, err := l.svcCtx.ConfigModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoConfigResp{
		Id:Info.Id
        BrandId:Info.BrandId
        Channel:Info.Channel
        Key:Info.Key
        Val:Info.Val
        CreateTime:Info.CreateTime
        
	}, nil
}

func (l *ListConfigLogic) ListConfig(req types.ListConfigReq) (*types.ListConfigResp, error) {
	list, total, err := l.svcCtx.ConfigModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	configList := make([]types.Config, 0, len(list))
	for _, v := range list {
		configList = append(configList, types.Config{
			Id:v.Id
            BrandId:v.BrandId
            Channel:v.Channel
            Key:v.Key
            Val:v.Val
            CreateTime:v.CreateTime
            
		})
	}
	return &types.ListConfigResp {
		ConfigList: configList,
		Total:    total,
	}, nil
}


func (l *UpdateConfigLogic) UpdateConfig(req types.UpdateConfigReq) (*types.UpdateConfigResp, error) {
	configInfo, err := l.svcCtx.ConfigModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.ConfigModel.Update(*configInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateConfigResp{
		Id: req.Id,
	}, nil
}