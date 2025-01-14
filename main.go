package main

import "fmt"

func main() {
	err := Candidates.Create()
	if err != nil {
		fmt.Println(err)
	}
}

/*
func main2() {
	fmt.Println("vim-go")

	sources := map[string]string{
		"CampaignDisclosureDataDictionary": "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTObHG93G%2bsgeH0QlvCDD%2bth2Vwaup9ji5PemZw7CSyMppvjHiwiMma2NETViZoHnLI2cLMBOcnLuov7xCX8Hn2MCZEJE5WyjUiflev7RHqd0%3d",
		"Candidates":                       "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTObHG93G%2bsgfaoI6R%2fraCOioDY6C7ntO4jZ3bgsgRc69EFlgrw5244A%3d%3d",
		"CanElections":                     "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTObHG93G%2bsgdslDDZtUVFcJiezaXUwsYOP5VwDs%2b%2f5AV0aLwo%2fAxGsw%3d%3d",
		"CmteCandidateLinks":               "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTV7DLOml0VEM8NtNSWZdDSnMVuzSKhxsytLIaRg2XNOt%2bmc6EtQzx68Qs1cXUpNe18YUPsqlkdzE%3d",
		"CmteOfficerLinks":                 "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTV7DLOml0VEM8NtNSWZdDSnMVuzSKhxsytLIaRg2XNOt%2bmc6EtQzx68Qs1cXUpNe18YUPsqlkdzE%3d",
		"Committees":                       "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTi2h8HDZ4OVd3tQEKCDCs9C1ZMQMrpj6CP8Ie6AlL%2byGH1WAwCdQ40Q%3d%3d",
		"D2Totals":                         "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTNkbSnZzI0yeKmXS%2fgJ83cT19aS7YfngCrsBD%2feW6Rio%3d",
		"Expenditures":                     "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTD%2bvpWmKc2VZ%2bJCbBCxl4cx7xzlFO5F9SQoP59eSuEcBsowHAAMouEA%3d%3d",
		"FiledDocs":                        "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTU%2fFrCSWNdYSLrN5i8qj5Mm5PBztveemQ9yvfZPsAhcs%3d",
		"Investments":                      "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTW8OecYgEwdA%2bMQgVwpOb7fUjvgTu6YWEmHNhcPt5DgTZFEAQ%3d%3d",
		"Officers":                         "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTaRZmFX4Fy2COa3u2oOoxB%2bnPaGWLhslX%2b3THZo7F5LA%3d",
		"PrevOfficers":                     "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkTbjIfKm4g5vRkk3oDLGo4QAmi7hKUVNhW8jnpJCZDhjCg51o8G7vWmg%3d%3d",
		"Receipts":                         "https://www.elections.il.gov/NewDocDisplay.aspx?khDtbt6dhc80IP6TQjNtMmEtKVxRXnDbQIqMvttCDL2B03HzqN0NedeM1l1mPpM%2fLor95xcMnwsP6tzjwAh%2fE%2f0fWy2j%2byVCCoJfGoAETz8JVtDTd1fxr8vvpWUdqx9uvIMji%2bZ5aMgzhvnDX4JRo4tJ867wjNkT7AEd3icGKOoijWRSj8iPgvRbQD3nKWoOP7H2ulEOEbY%3d",
	}

	files := make(map[string]string)

	fmt.Println("Begin Downloads")

	os.Mkdir("tmp", 0755)

	for key, value := range sources {
		fmt.Println("Downloading: " + key)
		fileName := "tmp/" + key + ".tsv"

		files[key] = fileName
		downloadFile(fileName, value)
	}

	fmt.Println("Downloads Complete")

	db, err := sql.Open("sqlite3", "./election_finances.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

func NewDB(driverName, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
*/
