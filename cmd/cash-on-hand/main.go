// Standalone tool to compute cash on hand for a committee by name.
//
// Cash on hand = latest D2 end balance + net receipts/expenditures filed
// since the end of that D2 reporting period.
//
// This matches the methodology used by IL Sunshine.
//
// Usage:
//
//	go run ./cmd/cash-on-hand/ -name "Citizens for Waguespack"
package main

import (
	"flag"
	"fmt"
	"os"

	disclosure "chicagogrowthproject.org/il-campaign-disclosure-data"
)

func main() {
	name := flag.String("name", "", "Committee name (partial match, case-insensitive)")
	dbPath := flag.String("db", disclosure.DefaultDatabasePath, "Path to DuckDB database file")
	flag.Parse()

	if *name == "" {
		fmt.Fprintln(os.Stderr, "Usage: cash-on-hand -name <committee name>")
		os.Exit(1)
	}

	db, err := disclosure.ConnectDbReadOnly(*dbPath)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	query := `
		WITH latest_d2 AS (
			SELECT DISTINCT ON (d.committee_id)
				d.committee_id,
				d.end_funds_available,
				f.report_period_end_date
			FROM d2_totals d
			JOIN filed_docs f ON d.filed_doc_id = f.id
			ORDER BY d.committee_id, f.report_period_end_date DESC
		),
		post_d2_receipts AS (
			SELECT r.committee_id, SUM(r.amount) AS amount
			FROM receipts r
			JOIN latest_d2 ld ON r.committee_id = ld.committee_id
			WHERE r.receive_date > ld.report_period_end_date
			GROUP BY r.committee_id
		),
		post_d2_expenditures AS (
			SELECT e.committee_id, SUM(e.amount) AS amount
			FROM expenditures e
			JOIN latest_d2 ld ON e.committee_id = ld.committee_id
			WHERE e.expended_date > ld.report_period_end_date
			GROUP BY e.committee_id
		)
		SELECT
			c.name,
			ld.report_period_end_date                AS last_d2_period_end,
			ld.end_funds_available                   AS last_d2_balance,
			COALESCE(pr.amount, 0)                   AS receipts_since_d2,
			COALESCE(pe.amount, 0)                   AS expenditures_since_d2,
			ld.end_funds_available
				+ COALESCE(pr.amount, 0)
				- COALESCE(pe.amount, 0)             AS cash_on_hand
		FROM committees c
		JOIN latest_d2 ld ON c.id = ld.committee_id
		LEFT JOIN post_d2_receipts pr ON c.id = pr.committee_id
		LEFT JOIN post_d2_expenditures pe ON c.id = pe.committee_id
		WHERE c.name ILIKE '%' || ? || '%'
		ORDER BY cash_on_hand DESC`

	rows, err := db.Query(query, *name)
	if err != nil {
		fmt.Println("Query error:", err)
		os.Exit(1)
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		found = true
		var committeeName, lastD2PeriodEnd string
		var lastD2Balance, receiptsSinceD2, expendituresSinceD2, cashOnHand float64
		if err := rows.Scan(&committeeName, &lastD2PeriodEnd, &lastD2Balance, &receiptsSinceD2, &expendituresSinceD2, &cashOnHand); err != nil {
			fmt.Println("Scan error:", err)
			os.Exit(1)
		}
		fmt.Printf("Committee:             %s\n", committeeName)
		fmt.Printf("Last D2 period end:    %s\n", lastD2PeriodEnd)
		fmt.Printf("Last D2 balance:       $%.2f\n", lastD2Balance)
		fmt.Printf("Receipts since D2:     $%.2f\n", receiptsSinceD2)
		fmt.Printf("Expenditures since D2: $%.2f\n", expendituresSinceD2)
		fmt.Printf("Cash on hand:          $%.2f\n\n", cashOnHand)
	}
	if !found {
		fmt.Printf("No committees found matching %q\n", *name)
	}
}
