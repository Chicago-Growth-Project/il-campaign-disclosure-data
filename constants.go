package main

var (
	Candidates = Table{
		Name: "candidates",
		URL:  "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTObHG93G%2bsgfaoI6R%2fraCOioDY6C7ntO4jZ3bgsgRc69EFlgrw5244A%3d%3d",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "first_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "last_name", RawName: "LastName", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "office", RawName: "Office", Type: ColumnTypeString},
			{Name: "district_type", RawName: "DistrictType", Type: ColumnTypeString},
			{Name: "district", RawName: "District", Type: ColumnTypeString},
			{Name: "residence_county", RawName: "ResidenceCounty", Type: ColumnTypeString},
			{Name: "party_affiliation", RawName: "PartyAffiliation", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "first_name", "last_name"},
	}
	CandidateElections = Table{
		Name: "candidate_elections",
		URL:  "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTObHG93G%2bsgdslDDZtUVFcJiezaXUwsYOP5VwDs%2b%2f5AV0aLwo%2fAxGsw%3d%3d",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "candidate_id", RawName: "CandidateID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "election_type", RawName: "ElectionType", Type: ColumnTypeString},
			{Name: "election_year", RawName: "ElectionYear", Type: ColumnTypeInt},
			{Name: "inc_chall_open", RawName: "IncChallOpen", Type: ColumnTypeString},
			{Name: "won_lost", RawName: "WonLost", Type: ColumnTypeString},
			{Name: "fair_campaign", RawName: "FairCampaign", Type: ColumnTypeBool},
			{Name: "limits_off", RawName: "LimitsOff", Type: ColumnTypeBool},
			{Name: "limits_off_reason", RawName: "LimitsOffReason", Type: ColumnTypeString},
		},
		IndexedColumns: []string{"id", "candidate_id"},
	}
	Committees = Table{
		Name: "committees",
		URL:  "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTi2h8HDZ4OVd3tQEKCDCs9C1ZMQMrpj6CP8Ie6AlL%2byGH1WAwCdQ40Q%3d%3d",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "type_of_committee", RawName: "TypeOfCommittee", Type: ColumnTypeString},
			{Name: "state_committee", RawName: "StateCommittee", Type: ColumnTypeString},
			{Name: "state_id", RawName: "StateID", Type: ColumnTypeString},
			{Name: "local_committee", RawName: "LocalCommittee", Type: ColumnTypeString},
			{Name: "local_id", RawName: "LocalID", Type: ColumnTypeString},
			{Name: "refer_name", RawName: "ReferName", Type: ColumnTypeString},
			{Name: "name", RawName: "Name", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "address3", RawName: "Address3", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "status", RawName: "Status", Type: ColumnTypeString},
			{Name: "status_date", RawName: "StatusDate", Type: ColumnTypeString},
			{Name: "creation_date", RawName: "CreationDate", Type: ColumnTypeString},
			{Name: "creation_amount", RawName: "CreationAmount", Type: ColumnTypeString},
			{Name: "disp_funds_return", RawName: "DispFundsReturn", Type: ColumnTypeString},
			{Name: "disp_funds_pol_comm", RawName: "DispFundsPolComm", Type: ColumnTypeString},
			{Name: "disp_funds_charity", RawName: "DispFundsCharity", Type: ColumnTypeString},
			{Name: "disp_funds_95", RawName: "DispFunds95", Type: ColumnTypeBool},
			{Name: "disp_funds_descrip", RawName: "DispFundsDescrip", Type: ColumnTypeString},
			{Name: "can_supp_opp", RawName: "CanSuppOpp", Type: ColumnTypeString},
			{Name: "policy_supp_opp", RawName: "PolicySuppOpp", Type: ColumnTypeString},
			{Name: "party_affiliation", RawName: "PartyAffiliation", Type: ColumnTypeString},
			{Name: "purpose", RawName: "Purpose", Type: ColumnTypeString},
		},
		IndexedColumns: []string{"id"},
	}

	AllTables = []Table{Committees}
)

/* sources := map[string]string{
	"CmteCandidateLinks":               "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTV7DLOml0VEM8NtNSWZdDSnMVuzSKhxsytLIaRg2XNOt%2bmc6EtQzx68Qs1cXUpNe18YUPsqlkdzE%3d",
	"CmteOfficerLinks":                 "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTV7DLOml0VEM8NtNSWZdDSnMVuzSKhxsytLIaRg2XNOt%2bmc6EtQzx68Qs1cXUpNe18YUPsqlkdzE%3d",
	"D2Totals":                         "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTNkbSnZzI0yeKmXS%2fgJ83cT19aS7YfngCrsBD%2feW6Rio%3d",
	"Expenditures":                     "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTD%2bvpWmKc2VZ%2bJCbBCxl4cx7xzlFO5F9SQoP59eSuEcBsowHAAMouEA%3d%3d",
	"FiledDocs":                        "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTU%2fFrCSWNdYSLrN5i8qj5Mm5PBztveemQ9yvfZPsAhcs%3d",
	"Investments":                      "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTW8OecYgEwdA%2bMQgVwpOb7fUjvgTu6YWEmHNhcPt5DgTZFEAQ%3d%3d",
	"Officers":                         "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTaRZmFX4Fy2COa3u2oOoxB%2bnPaGWLhslX%2b3THZo7F5LA%3d",
	"PrevOfficers":                     "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTbjIfKm4g5vRkk3oDLGo4QAmi7hKUVNhW8jnpJCZDhjCg51o8G7vWmg%3d%3d",
	"Receipts":                         "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkT7AEd3icGKOoijWRSj8iPgvRbQD3nKWoOP7H2ulEOEbY%3d",
}
*/
