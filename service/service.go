package service

type NodeData struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ForksCount  int `json:"forksCount"`
}
type Node struct {
	MasterNodeData []NodeData `json:"nodes"`
}
type Project struct {
	ProjectData Node `json:"projects"`
}
type Data struct {
	MasterData Project `json:"data"`
}


// this is the service layer
// it gets the dataset from the main file
// processes it and returns sum of all fork counts and names with a comma delimeter in NodeData object
func JoinNames(dataset Data) NodeData {

	allNamesJoined := ""
	sumOfForkCounts := 0

	for i, j := range dataset.MasterData.ProjectData.MasterNodeData {
		if i == 0 {
			allNamesJoined = allNamesJoined + j.Name
		} else {
			name := ", " + j.Name
			allNamesJoined = allNamesJoined + name
		}
		sumOfForkCounts = sumOfForkCounts + j.ForksCount
	}

	return NodeData {
		Name: allNamesJoined,
		ForksCount: sumOfForkCounts,
	}
}