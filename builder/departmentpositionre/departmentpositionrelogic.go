func (l *AddDepartmentPositionReLogic) AddDepartmentPositionRe(req types.AddDepartmentPositionReReq) (*types.AddDepartmentPositionReResp, error) {
	departmentPositionReId, err := l.svcCtx.DepartmentPositionReModel.Insert(model.DepartmentPositionRe{
		BrandId:req.BrandId
        DepartmentId:req.DepartmentId
        PositionId:req.PositionId
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddDepartmentPositionReResp{
		Id: departmentPositionReId,
	}, nil
}

func (l *DelDepartmentPositionReLogic) DelDepartmentPositionRe(req types.DelDepartmentPositionReReq) (*types.DelDepartmentPositionReResp, error) {
	err := l.svcCtx.DepartmentPositionReModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelDepartmentPositionReResp{
		Id: req.Id,
	}, nil
}

func (l *InfoDepartmentPositionReLogic) InfoDepartmentPositionRe(req types.InfoDepartmentPositionReReq) (*types.InfoDepartmentPositionReResp, error) {
	Info, err := l.svcCtx.DepartmentPositionReModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoDepartmentPositionReResp{
		Id:Info.Id
        BrandId:Info.BrandId
        DepartmentId:Info.DepartmentId
        PositionId:Info.PositionId
        CreateTime:Info.CreateTime
        
	}, nil
}

func (l *ListDepartmentPositionReLogic) ListDepartmentPositionRe(req types.ListDepartmentPositionReReq) (*types.ListDepartmentPositionReResp, error) {
	list, total, err := l.svcCtx.DepartmentPositionReModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	departmentPositionReList := make([]types.DepartmentPositionRe, 0, len(list))
	for _, v := range list {
		departmentPositionReList = append(departmentPositionReList, types.DepartmentPositionRe{
			Id:v.Id
            BrandId:v.BrandId
            DepartmentId:v.DepartmentId
            PositionId:v.PositionId
            CreateTime:v.CreateTime
            
		})
	}
	return &types.ListDepartmentPositionReResp {
		DepartmentPositionReList: departmentPositionReList,
		Total:    total,
	}, nil
}


func (l *UpdateDepartmentPositionReLogic) UpdateDepartmentPositionRe(req types.UpdateDepartmentPositionReReq) (*types.UpdateDepartmentPositionReResp, error) {
	departmentPositionReInfo, err := l.svcCtx.DepartmentPositionReModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DepartmentPositionReModel.Update(*departmentPositionReInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateDepartmentPositionReResp{
		Id: req.Id,
	}, nil
}