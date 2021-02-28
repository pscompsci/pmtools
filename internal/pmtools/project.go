package pmtools

import "fmt"

// Project ...
type Project struct {
	ProjectName        string
	ProjectCode        string
	ClientName         string
	ClientAbbreviation string
	ProjectManager     string
	BaseFolder         string
	*SOW
	*PCalc
}

// RepositoryName ...
func (p *Project) RepositoryName() string {
	return fmt.Sprintf("%s - %s %s", p.ProjectCode, p.ClientAbbreviation, p.ProjectName)
}
