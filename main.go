package main

import (
	"go-siteBJ/core"
	"go-siteBJ/entities"
	"go-siteBJ/siteServices"
)

var logger = core.NewFileLogger()

func main() {
	logger.Info().Msg("Entered the main")

	siteBJ := entities.Site{
		DomainName:         "www.SiteBJ.com",
		Menu:               []string{"Courses", "Guide", "Company"},
		Search:             []string{"JavaCourse", "GoCourse", "PhpCourse", "PythonCourse"},
		CompanyDescription: []string{"CompanyName: BJ", "Numberphone: 09142889711", "CompanyAddress: Iran-Tehran"},
		InternalLinks:      []string{"https://Java.com", "https://Go.com", "https://Php.com", "https://Python.com"},
	}
	course := entities.Course{
		NameCourse:  []string{"JavaCourse", "GoCourse", "PhpCourse", "PythonCourse"},
		NameTeacher: []string{"Ali Jabbary", "Hamed Naeemaei", "Reza Khani", "Vali Naseri"},
		PriceCourse: []int64{5500000, 7900000, 10000000, 6300000},
	}

	err := siteServices.SiteBJ(siteBJ, course)
	if err != nil {
		logger.Error().Stack().Err(err).Str("main", "SiteBJ").Msg("Error in main")
	}

	logger.Info().Msg("Out of main")
}
