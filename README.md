# IL Campaign Disclosure Database

A set of scripts for downloading IL's campaign disclosure database and other data sources into a useable DuckDB database.

## Usage

To create a new database file from scratch, load the repository and run the main file:

```
git clone https://github.com/lionelbarrow/il-campaign-disclosure-data.git
cd il-campaign-disclosure-data
go run .
```

This will download the latest data from the IL campaign disclosure website and create a new DuckDb database named `il-campaign-disclosures.db` in the current directory. On my Internet connection, this takes about 5 minutes, most of which is spent downloading the files.

Right now a few records in the d2_totals table are invalid as well as a bunch of records in the receipts table. Some error messages are to be expected relating to these. Affected records are not imported but otherwise the database is fine.

If you want to skip some tables, modify the `AllTables` array in [main.go](main.go) to only include the tables you want to download.

As of now this script takes about 30 minutes to run, but this may change as more data sources are added.

## Schema

The state website provides a full [data dictionary](https://elections.il.gov/campaigndisclosuredatafiles/CampaignDisclosureDataDictionary.txt) with field names and descriptions. The database schema matches this dictionary, with a few changes for improved usability:

* Field and table names are changed to snake_case.
* Join tables are renamed to match the `{first_table}_{second_table}` pattern rather than the mixed pattern in the dictionary.
* Many field abbreviations are expanded. For example, on the D2Totals table, `XferInNI` is expanded to `transfers_in_not_itemized`.
* The dictionary does not include data types, but we infer some of them and create the database columns accordingly.

You can see the schema in the [constants.go](constants.go) file.
