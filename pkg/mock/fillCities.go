package mock

import (
	"bankCLI/pkg"
	"bankCLI/pkg/models"
)

func FillCities() {
	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Душанбе",
		Region: "РРП",
	})

	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Худжанд",
		Region: "Согдийская область",
	})

	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Истаравшан",
		Region: "Согдийская область",
	})

	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Куляб",
		Region: "Хатлонская область",
	})
}
