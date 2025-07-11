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
		WardShapes,
		WardPrecinctShapes,
	}

	db, err := ConnectDb()

	defer db.Close()

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	for _, table := range allTables {
		fmt.Println("Creating table:", table.Name)
		err := table.Create(db)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("")
	}
}
