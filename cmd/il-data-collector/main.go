package main

import (
	"fmt"

	disclosure "chicagogrowthproject.org/il-campaign-disclosure-data"
)

func main() {
	allTables := []disclosure.Table{
		disclosure.Candidates,
		disclosure.CandidateCommittees,
		disclosure.CandidateElections,
		disclosure.Committees,
		disclosure.CommitteeOfficers,
		disclosure.D2Totals,
		disclosure.Expenditures,
		disclosure.FiledDocs,
		disclosure.Investments,
		disclosure.Officers,
		disclosure.PreviousOfficers,
		disclosure.Receipts,
		disclosure.WardShapes,
		disclosure.WardPrecinctShapes,
	}

	db, err := disclosure.ConnectDb(disclosure.DefaultDatabasePath)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	for _, table := range allTables {
		fmt.Println("Creating table:", table.Name)
		if err := table.Create(db); err != nil {
			fmt.Println(err)
		}
		fmt.Println("")
	}
}
