package models

import (
	"fmt"

	"gorm.io/gorm"
)

const notFound uint = 0

type Project struct {
	gorm.Model
	Name string
}

// TODO: change to PrintAllEntries
func PrintAll(project Project, db *gorm.DB) {
	var entries []Entry
	db.Where("project_id = ?", project.ID).Find(&entries) // note to self: queries should be snakecase
	for _, entry := range entries {
		fmt.Printf(Format, entry.ID, entry.Message)
	}
}

func PrintProjects(db *gorm.DB) {
	if hasProjects(db) {
		projects := getAllProjects(db)
		for _, project := range projects {
			fmt.Printf(Format, project.ID, project.Name)
		}
	} else {
		fmt.Printf("There are no projects available")
	}
}

// error handling in case no projects are found
func hasProjects(db *gorm.DB) bool {
	var projects []Project
	if err := db.Find(&projects).Error; err != nil {
		return false
	}
	return true
}

// countProjects: return the number of projects
func countProjects(db *gorm.DB) int {
	var projects []Project
	db.Find(&projects) // note to self: queries should be snakecase
	return len(projects)
}

// getProject: return a project by id
func getProject(projectId int, db *gorm.DB) Project {
	var project Project
	db.Where("id = ?", projectId).Find(&project)
	return project
}

// getAllProjects: return all projects
func getAllProjects(db *gorm.DB) []Project {
	var projects []Project
	if hasProjects(db) {
		db.Find(&projects)
	}
	return projects
}

// DeleteProject: delete a project by id
func DeleteProject(pKey int, db *gorm.DB) {
	// what if pKey does not exist?
	db.Where("project_id = ?", pKey).Delete(&Entry{})
	db.Delete(&Project{}, pKey)
}

// TODO: move to a more appropriate place
func newProjectPrompt() string {
	var name string
	fmt.Println("what would you like to name your project?")
	fmt.Scanf("%s", &name)
	return name
}

/* other */
// TODO: want to centralize prompts
func CreateProject(db *gorm.DB) Project {
	name := newProjectPrompt()
	PrintProjects(db)
	proj := Project{Name: name}
	db.Create(&proj)
	return proj
}

func GetOrCreateProject(pKey int, db *gorm.DB) Project {
	proj := getProject(pKey, db)
	if proj.ID == notFound {
		return CreateProject(db)
	}
	return proj
}

func RenameProject(pKey int, db *gorm.DB) {
	name := newProjectPrompt()
	var project Project
	db.Where("id = ?", pKey).First(&project)
	project.Name = name
	db.Save(&project)
}
