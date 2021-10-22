func (l *AddPositionLogic) AddPosition(req types.AddPositionReq) (*types.AddPositionResp, error) {
	positionId, err := l.svcCtx.PositionModel.Insert(model.Position{
		BrandId:req.BrandId
        Name:req.Name
        Remark:req.Remark
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddPositionResp{
		Id: positionId,
	}, nil
}

func (l *DelPositionLogic) DelPosition(req types.DelPositionReq) (*types.DelPositionResp, error) {
	err := l.svcCtx.PositionModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelPositionResp{
		Id: req.Id,
	}, nil
}

func (l *InfoPositionLogic) InfoPosition(req types.InfoPositionReq) (*types.InfoPositionResp, error) {
	Info, err := l.svcCtx.PositionModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoPositionResp{
		Id:Info.Id
        BrandId:Info.BrandId
        Name:Info.Name
        Remark:Info.Remark
        DeleteTime:Info.DeleteTime
        CreateTime:Info.CreateTime
        UpdateTime:Info.UpdateTime
        
	}, nil
}

func (l *ListPositionLogic) ListPosition(req types.ListPositionReq) (*types.ListPositionResp, error) {
	list, total, err := l.svcCtx.PositionModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	positionList := make([]types.Position, 0, len(list))
	for _, v := range list {
		positionList = append(positionList, types.Position{
			Id:v.Id
            BrandId:v.BrandId
            Name:v.Name
            Remark:v.Remark
            DeleteTime:v.DeleteTime
            CreateTime:v.CreateTime
            UpdateTime:v.UpdateTime
            
		})
	}
	return &types.ListPositionResp {
		PositionList: positionList,
		Total:    total,
	}, nil
}


func (l *UpdatePositionLogic) UpdatePosition(req types.UpdatePositionReq) (*types.UpdatePositionResp, error) {
	positionInfo, err := l.svcCtx.PositionModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.PositionModel.Update(*positionInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdatePositionResp{
		Id: req.Id,
	}, nil
}