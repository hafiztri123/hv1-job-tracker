package applications

func (s *ApplicationService) CreateApplication(req *CreateApplicationDto, userId string) error {

	if req.Status == nil {
		req.Status = new(string)
		*req.Status = "Wishlist"
	}

	if err := s.repo.InsertApplication(req, userId); err != nil {
		return err
	}

	return nil
}

func (s *ApplicationService) GetApplications(userId string, queryParams ApplicationQueryParams) ([]Application, error) {
	applications, err := s.repo.FindApplicationsById(userId, queryParams)

	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (s *ApplicationService) DeleteApplications(userId, applicationId string) error {
	err := s.repo.DeleteApplications(userId, applicationId)

	return err
}

func (s *ApplicationService) GetApplicationOptions(queryParams ApplicationOptionQueryParams) ApplicationOptions {
	options := ApplicationOptions{
		StatusOption: []string{},
	}

	if queryParams.StatusOption {
		options.StatusOption = []string{"Wishlist", "Applied", "Interviewing", "Offer", "Rejected"}
	}

	return options
}

func (s *ApplicationService) UpdateApplication(body UpdateApplicationDto, userId, applicationId string) error {

	return s.repo.UpdateApplications(userId, applicationId, &body)

}

func (s *ApplicationService) BatchDeleteApplications(userId string, req *BatchDeleteDto) error {
	return s.repo.BatchDeleteApplications(userId, req.ApplicationIds)
}

func (s *ApplicationService) BatchUpdateStatusApplications(userId string, req *BatchUpdateStatusDto) error {
	return s.repo.BatchUpdateStatusApplications(userId, req.ApplicationIds, req.Status)
}
