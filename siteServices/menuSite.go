package siteServices

import (
	"errors"
	"fmt"
	"go-siteBJ/entities"
)

func MenuSite(site *entities.Site, course *entities.Course, menu string) error {
	logger.Info().Str("menuSite", "MenuSite").Msg("Entered the MenuSite")

	for menu != "exit" {
	    if site.Menu[0] == menu {

	    	inCourse := ""
	     	fmt.Println("Courses: JavaCourse, GoCourse, PhpCourse, PythonCourse")
	        fmt.Print("Course: ")
	        fmt.Scanln(&inCourse)

		    if course.NameCourse[0] == inCourse {
			    fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", course.NameCourse[0], course.NameTeacher[0], course.PriceCourse[0])
		    } else if course.NameCourse[1] == inCourse {
		    	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", course.NameCourse[1], course.NameTeacher[1], course.PriceCourse[1])
		    } else if course.NameCourse[2] == inCourse {
		    	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", course.NameCourse[2], course.NameTeacher[2], course.PriceCourse[2])
		    } else if course.NameCourse[3] == inCourse {
		    	fmt.Printf("NameCourse: %s , NameTeacher: %s , PriceCourse: %d\n", course.NameCourse[3], course.NameTeacher[3], course.PriceCourse[3])
		    } else if inCourse == "exit" {
		    	fmt.Println("Exit Courses!")
		    	break
		    } else {
			    return errors.New("this course does not exist")
		    }

        } else if site.Menu[1] == menu {
	    	fmt.Print("Contact us through the links below:\n Telegram\n Instagram\n Twitter\n Number\n")
			break
		} else if site.Menu[2] == menu {
	    	fmt.Printf(" %s\n %s\n %s\n", site.CompanyDescription[0], site.CompanyDescription[1], site.CompanyDescription[2])
			break
		} else if site.Menu[0] != menu && site.Menu[1] != menu && site.Menu[2] != menu {
			 return errors.New("this menu does not exist")
		}
	}

	logger.Info().Str("menuSite", "MenuSite").Msg("Out of MenuSite")

	return nil
}


