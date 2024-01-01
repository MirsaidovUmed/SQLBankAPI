package repository

func (repo *Repository) GetPercent() float64 {
	return repo.Database.Percent
}
