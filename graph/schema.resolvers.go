package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bd_test_task_three/graph/generated"
	"bd_test_task_three/graph/model"
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"context"
	"fmt"
)

// you can run this query to create a new project
// mutation {
// 	createProject (input: {name: "hello", description: "hello", forksCount: 1}){
// 	  node {
// 		name
// 		description
// 		forksCount
// 	  }
// 	 }
//  }
func (r *mutationResolver) CreateProject(ctx context.Context, input *model.NewProject) (*model.Project, error) {

	user, password, dbName := "postgres", "postgres", "postgres"
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error connecting to DB: ", err)
		panic(err)
	}
	
	name := *input.Name
	description := *input.Description
	forksCount := *input.ForksCount
	nameScan := ""

	errQueryRow := db.QueryRow(
        "INSERT INTO Nodes(name, description, forkscount) VALUES($1, $2, $3) RETURNING name",
        name, description, forksCount).Scan(&nameScan)
	
	if errQueryRow != nil {
		
		fmt.Println("Error inserting record to DB: ", errQueryRow)
		return nil, nil
	}

	finalNode := &model.Node{
		Name:       &name,
		Description: &description,
		ForksCount: &forksCount,
	}
	project := &model.Project{
		Node: finalNode,
	}

	return project, nil
}

// you can run this query to get all the projects in the database
// this endpoint will return the names of all projects and sum of all forks count
// {
//   projects{
//     node{
//       name
//       description
//       forksCount
//     }
//   }
// }
func (r *queryResolver) Projects(ctx context.Context) (*model.Project, error) {

	var allProjects *model.Project // returned at end with allNames and sumForksCounts
	allNames := "" // will contain all the names with , delimeter
	sumForkCounts := 0 // will contain sum of all the forkCounts

	// connects with the database
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error connecting to DB: ", err)
		panic(err)
	}

	rows, errRows := db.Query("select * from nodes")
	if errRows != nil {
		fmt.Println("Error connecting to DB: ", err)
		panic(err)
	}

	// scanning and returning all names with , delimiter and sum of all fork counts
	i := 0
	for rows.Next() {
		name, description, forksCount := "", "", 0
		if errScan := rows.Scan(&name, &description, &forksCount); errScan == nil {
			if i == 0 {
				allNames = allNames + name
				sumForkCounts = sumForkCounts + forksCount
			} else {
				name = ", " + name
				allNames = allNames + name
				sumForkCounts = sumForkCounts + forksCount
			}
			i += 1
        } else {
			fmt.Println("Error ", errScan)
		}
	}

	finalNode := &model.Node{
		Name:       &allNames,
		ForksCount: &sumForkCounts,
	}
	allProjects = &model.Project{
		Node: finalNode,
	}

	return allProjects, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
