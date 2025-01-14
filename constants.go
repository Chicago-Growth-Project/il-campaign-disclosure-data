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
	CandidateCommittees = Table{
		Name: "candidate_committees",
		URL:  "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTV7DLOml0VEM8NtNSWZdDSnMVuzSKhxsytLIaRg2XNOt%2bmc6EtQzx68Qs1cXUpNe18YUPsqlkdzE%3d",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "candidate_id", RawName: "CandidateID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
		},
		IndexedColumns: []string{"id", "candidate_id", "committee_id"},
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
			{Name: "creation_amount", RawName: "CreationAmount", Type: ColumnTypeString}, // TODO: Could be a decimal
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
	CommitteeOfficers = Table{
		Name: "committee_officers",
		URL:  "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTV7DLOml0VEM8NtNSWZdDSnMVuzSKhxsytLIaRg2XNOt%2bmc6EtQzx68Qs1cXUpNe18YUPsqlkdzE%3d",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "officer_id", RawName: "OfficerID", Type: ColumnTypeInt, NotNullable: true},
		},
		IndexedColumns: []string{"id", "committee_id", "officer_id"},
	}

	Officers = Table{
		Name: "officers",
		URL:  "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTaRZmFX4Fy2COa3u2oOoxB%2bnPaGWLhslX%2b3THZo7F5LA%3d",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "last_name", RawName: "LastName", Type: ColumnTypeString},
			{Name: "first_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "title", RawName: "Title", Type: ColumnTypeString},
			{Name: "phone", RawName: "Phone", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "last_name", "first_name"},
	}

	PreviousOfficers = Table{
		Name: "previous_officers",
		URL:  "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTbjIfKm4g5vRkk3oDLGo4QAmi7hKUVNhW8jnpJCZDhjCg51o8G7vWmg%3d%3d",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "last_name", RawName: "LastName", Type: ColumnTypeString},
			{Name: "first_name", RawName: "FirstName", Type: ColumnTypeString},
			{Name: "address1", RawName: "Address1", Type: ColumnTypeString},
			{Name: "address2", RawName: "Address2", Type: ColumnTypeString},
			{Name: "city", RawName: "City", Type: ColumnTypeString},
			{Name: "state", RawName: "State", Type: ColumnTypeString},
			{Name: "zip", RawName: "Zip", Type: ColumnTypeString},
			{Name: "title", RawName: "Title", Type: ColumnTypeString},
			{Name: "resign_date", RawName: "ResignDate", Type: ColumnTypeString},
			{Name: "redaction_requested", RawName: "RedactionRequested", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "committee_id", "last_name", "first_name"},
	}

	D2Totals = Table{
		Name: "d2_totals",
		URL:  "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTNkbSnZzI0yeKmXS%2fgJ83cT19aS7YfngCrsBD%2feW6Rio%3d",
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "committee_id", RawName: "CommitteeID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "filed_doc_id", RawName: "FiledDocID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "beg_funds_avail", RawName: "BegFundsAvail", Type: ColumnTypeDecimal},
			{Name: "indiv_contrib_i", RawName: "IndivContribI", Type: ColumnTypeDecimal},
			{Name: "indiv_contrib_ni", RawName: "IndivContribNI", Type: ColumnTypeDecimal},
			{Name: "xfer_in_i", RawName: "XferInI", Type: ColumnTypeDecimal},
			{Name: "xfer_in_ni", RawName: "XferInNI", Type: ColumnTypeDecimal},
			{Name: "loan_rcv_i", RawName: "LoanRcvI", Type: ColumnTypeDecimal},
			{Name: "loan_rcv_ni", RawName: "LoanRcvNI", Type: ColumnTypeDecimal},
			{Name: "other_rct_i", RawName: "OtherRctI", Type: ColumnTypeDecimal},
			{Name: "other_rct_ni", RawName: "OtherRctNI", Type: ColumnTypeDecimal},
			{Name: "total_receipts", RawName: "TotalReceipts", Type: ColumnTypeDecimal},
			{Name: "in_kind_i", RawName: "InKindI", Type: ColumnTypeString},
			{Name: "in_kind_ni", RawName: "InKindNI", Type: ColumnTypeDecimal},
			{Name: "total_in_kind", RawName: "TotalInKind", Type: ColumnTypeDecimal},
			{Name: "xfer_out_i", RawName: "XferOutI", Type: ColumnTypeDecimal},
			{Name: "xfer_out_ni", RawName: "XferOutNI", Type: ColumnTypeDecimal},
			{Name: "loan_made_i", RawName: "LoanMadeI", Type: ColumnTypeString},
			{Name: "loan_made_ni", RawName: "LoanMadeNI", Type: ColumnTypeDecimal},
			{Name: "expend_i", RawName: "ExpendI", Type: ColumnTypeDecimal},
			{Name: "expend_ni", RawName: "ExpendNI", Type: ColumnTypeDecimal},
			{Name: "independent_exp_i", RawName: "IndependentExpI", Type: ColumnTypeDecimal},
			{Name: "independent_exp_ni", RawName: "IndependentExpNI", Type: ColumnTypeDecimal},
			{Name: "total_expend", RawName: "TotalExpend", Type: ColumnTypeDecimal},
			{Name: "debts_i", RawName: "DebtsI", Type: ColumnTypeDecimal},
			{Name: "debts_ni", RawName: "DebtsNI", Type: ColumnTypeDecimal},
			{Name: "total_debts", RawName: "TotalDebts", Type: ColumnTypeDecimal},
			{Name: "total_invest", RawName: "TotalInvest", Type: ColumnTypeDecimal},
			{Name: "end_funds_avail", RawName: "EndFundsAvail", Type: ColumnTypeDecimal},
			{Name: "archived", RawName: "Archived", Type: ColumnTypeBool},
		},
		IndexedColumns: []string{"id", "committee_id", "filed_doc_id"},
	}

	AllTables = []Table{D2Totals}
)

/* sources := map[string]string{
	"Expenditures":                     "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTD%2bvpWmKc2VZ%2bJCbBCxl4cx7xzlFO5F9SQoP59eSuEcBsowHAAMouEA%3d%3d",
	"FiledDocs":                        "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTU%2fFrCSWNdYSLrN5i8qj5Mm5PBztveemQ9yvfZPsAhcs%3d",
	"Investments":                      "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTW8OecYgEwdA%2bMQgVwpOb7fUjvgTu6YWEmHNhcPt5DgTZFEAQ%3d%3d",
	"Receipts":                         "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkT7AEd3icGKOoijWRSj8iPgvRbQD3nKWoOP7H2ulEOEbY%3d",
}
*/
