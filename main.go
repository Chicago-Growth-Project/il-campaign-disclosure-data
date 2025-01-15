package main

import "fmt"

const DatabasePath = "il-campaign-disclosures.db"

func main() {
	allTables := []Table{
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

	for _, table := range allTables {
		err := table.Create()
		if err != nil {
			fmt.Println(err)
		}
	}
}
