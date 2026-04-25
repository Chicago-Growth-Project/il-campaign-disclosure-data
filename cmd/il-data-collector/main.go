package main

import (
	"fmt"

	"chicagogrowthproject.org/il-campaign-disclosure-data/internal"
)

func main() {
	allTables := []internal.Table{
		internal.Candidates,
		internal.CandidateCommittees,
		internal.CandidateElections,
		internal.Committees,
		internal.CommitteeOfficers,
		internal.D2Totals,
		internal.Expenditures,
		internal.FiledDocs,
		internal.Investments,
		internal.Officers,
		internal.PreviousOfficers,
		internal.Receipts,
		internal.WardShapes,
		internal.WardPrecinctShapes,
	}

	db, err := internal.ConnectDb(internal.DefaultDatabasePath)
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
