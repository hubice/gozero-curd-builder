func (l *AddDepartmentLogic) AddDepartment(req types.AddDepartmentReq) (*types.AddDepartmentResp, error) {
	departmentId, err := l.svcCtx.DepartmentModel.Insert(model.Department{
		BrandId:req.BrandId
        ParentId:req.ParentId
        DingTalkId:req.DingTalkId
        Type:req.Type
        Name:req.Name
        
	})
	if err != nil {
		return nil, err
	}
	return &types.AddDepartmentResp{
		Id: departmentId,
	}, nil
}

func (l *DelDepartmentLogic) DelDepartment(req types.DelDepartmentReq) (*types.DelDepartmentResp, error) {
	err := l.svcCtx.DepartmentModel.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DelDepartmentResp{
		Id: req.Id,
	}, nil
}

func (l *InfoDepartmentLogic) InfoDepartment(req types.InfoDepartmentReq) (*types.InfoDepartmentResp, error) {
	Info, err := l.svcCtx.DepartmentModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.InfoDepartmentResp{
		Id:Info.Id
        BrandId:Info.BrandId
        ParentId:Info.ParentId
        DingTalkId:Info.DingTalkId
        Type:Info.Type
        Name:Info.Name
        DeleteTime:Info.DeleteTime
        CreateTime:Info.CreateTime
        UpdateTime:Info.UpdateTime
        
	}, nil
}

func (l *ListDepartmentLogic) ListDepartment(req types.ListDepartmentReq) (*types.ListDepartmentResp, error) {
	list, total, err := l.svcCtx.DepartmentModel.List(req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	departmentList := make([]types.Department, 0, len(list))
	for _, v := range list {
		departmentList = append(departmentList, types.Department{
			Id:v.Id
            BrandId:v.BrandId
            ParentId:v.ParentId
            DingTalkId:v.DingTalkId
            Type:v.Type
            Name:v.Name
            DeleteTime:v.DeleteTime
            CreateTime:v.CreateTime
            UpdateTime:v.UpdateTime
            
		})
	}
	return &types.ListDepartmentResp {
		DepartmentList: departmentList,
		Total:    total,
	}, nil
}


func (l *UpdateDepartmentLogic) UpdateDepartment(req types.UpdateDepartmentReq) (*types.UpdateDepartmentResp, error) {
	departmentInfo, err := l.svcCtx.DepartmentModel.FindOne(req.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DepartmentModel.Update(*departmentInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdateDepartmentResp{
		Id: req.Id,
	}, nil
}