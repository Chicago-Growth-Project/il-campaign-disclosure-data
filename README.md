# IL Campaign Disclosure Database

A set of scripts for downloading IL's campaign disclosure database and other data sources into a useable DuckDB database.

## Caveats

At the moment there's an issue with the Windows executable due to permissions on the DuckDB server.
A [github issue](https://github.com/marcboeker/go-duckdb/issues/478) has been opened to track the issue.
The issue should hopefully resolve once they fix the access to the extension.

## Usage

To create a new database file from scratch, Download the executable for your machine and run it

`./il-data-collector-macos-arm` for Macs using the M chips
`./il-data-collector-linux-amd64` for Linux machines
`il-data-collector-windows-amd64.exe` for Windows machines

This will download the latest data from the IL campaign disclosure website as well as other sources into a new DuckDB
database named `il-campaign-disclosures.db` in the curent directory. The process takes about 30+ minutes. The majority
of the time is taken in loading the data. Expect this time to increase as more data sources are added.

Right now a few records in the d2_totals and receipts table are invalid. Some error messages are to be expected relating to these.
Affected records are not imported but otherwise the database is fine. If you need to have the data, you can manually insert it
or reach out to contact@chicagogrowthproject.org

## Contributing

### Repo setup

PR requests are welcome!

```
git clone https://github.com/chicago-growth-project/il-campaign-disclosure-data.git
cd il-campaign-disclosure-data
go run .
```

If you want to skip some tables, modify the `AllTables` array in [main.go](main.go) to only include the tables you want to download.

Development on Windows requires the installation of gcc or clang. Instructions for setting up
with MinGW-w64 can be found [here](https://code.visualstudio.com/docs/cpp/config-mingw#_installing-the-mingww64-toolchain)

## Schema

The state website provides a full [data dictionary](https://elections.il.gov/campaigndisclosuredatafiles/CampaignDisclosureDataDictionary.txt) with field names and descriptions. The database schema matches this dictionary, with a few changes for improved usability:

* Field and table names are changed to snake_case.
* Join tables are renamed to match the `{first_table}_{second_table}` pattern rather than the mixed pattern in the dictionary.
* Many field abbreviations are expanded. For example, on the D2Totals table, `XferInNI` is expanded to `transfers_in_not_itemized`.
* The dictionary does not include data types, but we infer some of them and create the database columns accordingly.

You can see the schema in the [constants.go](constants.go) file.
