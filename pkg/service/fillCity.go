package service

func (s *Service) FillCities() (err error) {
	err = s.Repository.AddCity("Душанбе", "РРП")

	if err != nil {
		return
	}

	err = s.Repository.AddCity("Худжанд", "Согдийская область")

	if err != nil {
		return
	}

	err = s.Repository.AddCity("Истаравшан", "Согдийская область")

	if err != nil {
		return
	}

	err = s.Repository.AddCity("Куляб", "Хатлонская область")

	if err != nil {
		return
	}

	return
}
