package siteServices

import (
	"errors"
	"fmt"
	"go-siteBJ/entities"
)

func SearchSite(site *entities.Site , course *entities.Course, search string) error {
	logger.Info().Str("searchSite", "SearchSite").Msg("Entered the SearchSite")

	for search != "exit" {
	    if site.Search[0] == search {
	    	JavaCourse(*course)
			break
	    } else if site.Search[1] == search {
	    	GoCourse(*course)
			break
	    } else if site.Search[2] == search {
	    	PhpCourse(*course)
			break
	    } else if site.Search[3] == search {
	    	PythonCourse(*course)
			break
		} else if site.Search[0] != search && site.Search[1] != search && site.Search[2] != search && site.Search[3] != search {
			return errors.New("this search does not exist")
		}
	}

	logger.Info().Str("searchSite", "SearchSite").Msg("Out of SearchSite")

	return nil
}

func JavaCourse(search entities.Course) {
	logger.Info().Str("searchSite", "JavaCourse").Msg("Entered the JavaCourse")
	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", search.NameCourse[0], search.NameTeacher[0], search.PriceCourse[0])
	logger.Info().Str("searchSite", "JavaCourse").Msg("Out of JavaCourse")
}

func GoCourse(search entities.Course) {
	logger.Info().Str("searchSite", "JavaCourse").Msg("Entered the GoCourse")
	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", search.NameCourse[1], search.NameTeacher[1], search.PriceCourse[1])
	logger.Info().Str("searchSite", "JavaCourse").Msg("Out of GoCourse")
}

func PhpCourse(search entities.Course) {
	logger.Info().Str("searchSite", "JavaCourse").Msg("Entered the PhpCourse")
	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", search.NameCourse[2], search.NameTeacher[2], search.PriceCourse[2])
	logger.Info().Str("searchSite", "JavaCourse").Msg("Out of PhpCourse")
}

func PythonCourse(search entities.Course) {
	logger.Info().Str("searchSite", "JavaCourse").Msg("Entered the PythonCourse")
	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", search.NameCourse[3], search.NameTeacher[3], search.PriceCourse[3])
	logger.Info().Str("searchSite", "JavaCourse").Msg("Out of PythonCourse")
}