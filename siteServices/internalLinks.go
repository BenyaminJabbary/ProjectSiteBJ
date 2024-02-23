package siteServices

import (
	"errors"
	"fmt"
	"go-siteBJ/entities"
)

func LinksSite(site *entities.Site, course *entities.Course, link string) error {
	logger.Info().Str("InternalLinks", "LinksSite").Msg("Entered the LinksSite")

	for link != "exit" {
	    if site.InternalLinks [0] == link {
	    	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", course.NameCourse[0], course.NameTeacher[0], course.PriceCourse[0])
			break
	    } else if site.InternalLinks[1] == link {
	    	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", course.NameCourse[1], course.NameTeacher[1], course.PriceCourse[1])
			break
	    } else if site.InternalLinks[2] == link {
			fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", course.NameCourse[2], course.NameTeacher[2], course.PriceCourse[2])
			break
	    } else if site.InternalLinks[3] == link {
	    	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", course.NameCourse[3], course.NameTeacher[3], course.PriceCourse[3])
			break
	    } else if site.InternalLinks[0] != link && site.InternalLinks[1] != link && site.InternalLinks[2] != link && site.InternalLinks[3] != link {
			return errors.New("this link does not exist")
		}
    }

	logger.Info().Str("internalLinks", "LinksSite").Msg("Out of LinksSite")

	return nil
}