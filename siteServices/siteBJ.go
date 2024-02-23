package siteServices

import (
	"errors"
	"fmt"
	"go-siteBJ/core"
	"go-siteBJ/entities"
)

var logger = core.NewFileLogger()

func SiteBJ(site entities.Site, course entities.Course) error {
	logger.Info().Str("siteBJ", "SiteBJ").Msg("Entered the SiteBJ")

	domainName := ""
	fmt.Print("Google: ")
	fmt.Scanln(&domainName)
	if site.DomainName == domainName {
		fmt.Println("Welcome to siteBJ")
		err := Site(site, course, domainName)
		if err != nil {
			logger.Error().Stack().Err(err).Str("siteBJ", "Site").Msg("Error in SiteBJ")
		}
	} else if site.DomainName != domainName {
		return errors.New("this site does not exist")
	}

	logger.Info().Str("siteBJ", "SiteBJ").Msg("Out of SiteBJ")

	return nil
}

func Site(site entities.Site, course entities.Course, domainName string) error {
	logger.Info().Str("siteBJ", "Site").Msg("Entered the Site")

	for site.DomainName == domainName {
	    siteBJ := ""
		fmt.Println("SiteBJ: Menu, Search, InternalLinks")
	    fmt.Print("SiteBJ: ")
	    fmt.Scanln(&siteBJ)
	    if siteBJ == "Menu" {
			err := Menu(site, course, domainName) 
			if err != nil {
				logger.Error().Stack().Err(err).Str("siteBJ", "Menu").Msg("Error in Menu")
			}
	    } else if siteBJ == "Search" {
		    err := Search(site, course, domainName)
			if err != nil {
				logger.Error().Stack().Err(err).Str("siteBJ", "Search").Msg("Error in Search")
			}
	    } else if siteBJ == "InternalLinks" {
			err := Link(site, course, domainName)
			if err != nil {
				logger.Error().Stack().Err(err).Str("siteBJ", "Link").Msg("Error in Link")
			}
		} else if siteBJ == "exit" {
			fmt.Println("Exiting...")
			break
		} else if siteBJ != "Menu" && siteBJ != "Search" && siteBJ != "InternalLinks" {
			return errors.New("this option does not exist")
		}
    }

	logger.Info().Str("siteBJ", "Site").Msg("Out of Site")

	return nil
}

func Menu(site entities.Site, course entities.Course, domainName string) error {
	logger.Info().Str("siteBJ", "Menu").Msg("Entered the Menu")

	for site.DomainName == domainName {
		menu := ""
	    fmt.Println("Menu: Courses, Guide, Company")
	    fmt.Print("Select menu: ")
	    fmt.Scanln(&menu)
	    if menu != "exit" {
	    	err := MenuSite(&site, &course, menu)
			if err != nil {
				return errors.New("error in site menu")
			}
	    } else if menu == "exit" {
	    	fmt.Println("Exit Menu!")
	    	break
	    }
	}

	logger.Info().Str("siteBJ", "Menu").Msg("Out of Menu")

	return nil
}

func Search(site entities.Site, course entities.Course, domainName string) error {
	logger.Info().Str("siteBJ", "Search").Msg("Entered the Search")

	for site.DomainName == domainName {
		search := ""
	    fmt.Println("Searches: JavaCourse, GoCourse, PhpCourse, PythonCourse")
		fmt.Print("Search: ")
	    fmt.Scanln(&search)
		if search != "exit" {
	    	err := SearchSite(&site, &course, search)
			if err != nil {
				return errors.New("error in site search")
			}
	    } else if search == "exit" {
	    	fmt.Println("Exit Search!")
	    	break
	    }
	}

	logger.Info().Str("siteBJ", "Search").Msg("Out of Search")

	return nil
}

func Link(site entities.Site, course entities.Course, domainName string) error {
	logger.Info().Str("siteBJ", "Link").Msg("Entered the Link")

	for site.DomainName == domainName {
	    link := ""
	    fmt.Println("InternalLinks: https://Java.com, https://Go.com, https://Php.com, https://Python.com")
	    fmt.Print("link: ")
	    fmt.Scanln(&link)
		if link != "exit" {
			err := LinksSite(&site, &course, link)
			if err != nil {
				return errors.New("error in site links")
			}
		} else if link == "exit" {
			fmt.Println("Exit InternalLinks!")
			break
		}
	}

	logger.Info().Str("siteBJ", "Link").Msg("Out of Link")

	return nil
}