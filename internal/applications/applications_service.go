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

func (s *ApplicationService) GetApplications(userId string) ([]Application, error) {
	applications, err := s.repo.FindApplicationsById(userId)

	if err != nil {
		return nil, err
	}

	return applications, nil
}
