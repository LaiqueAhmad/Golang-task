package main

import (
	"fmt"
	"bd_test_task_three/service"
	"bd_test_task_three/repository"
)

// main file which calls the repo layer and service layer
// prints all the names and sum of all forks
func main() {

	dataset := repository.RepoLayer()
	formattedDataSet := service.JoinNames(dataset)
	fmt.Println("**************")
	fmt.Println("All names joined with comma separated delimeter: ", formattedDataSet.Name)
	fmt.Println("Sum of all forkCounts                          : ", formattedDataSet.ForksCount)
	fmt.Println("**************")
}