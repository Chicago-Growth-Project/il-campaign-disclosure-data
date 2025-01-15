package main

import "fmt"

var AllTables = []Table{
	Candidates,
	CandidateCommittees,
	CandidateElections,
	Committees,
	CommitteeOfficers,
	D2Totals,
	Expenditures,
	FiledDocs,
	Investments,
	Officers,
	PreviousOfficers,
	Receipts,
}

func main() {

	for _, table := range AllTables {
		err := table.Create()
		if err != nil {
			fmt.Println(err)
		}
	}
}
