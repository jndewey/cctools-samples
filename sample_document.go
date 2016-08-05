//Sample Document - Convertible Promissory Note

package main 

import "fmt"
import "github.com/jndewey/cctools" // replace this import path with the correct path to cctools on your system

func main() {
	NoParty := cctools.Party{}
	Company := cctools.Party {"1", "ABC Corporation", "Maker", true, false, "Delaware", "corporation", "123", "joe.smith@texas.org", "2100 Lone Drive", "Suite 3100", "Dallas", "Texas", "33123"}
	Holder := cctools.Party {"2", "First Bank of Texas", "Payee", true, true, "Texas", "banking corporation", "456", "tom.jones@banks.org", "2113 Lone Lane", "Suite 100", "Dallas", "Texas", "33123"}
	parties := []cctools.Party {Company, Holder}
	CompanyName := cctools.Clause {"Company Name", "none", "ABC Corporation", "Definition", false, NoParty}
    HolderName := cctools.Clause {"Holder Name", "none", "First Bank of Texas", "Definition", false, NoParty}
	preamble := cctools.Clause {"Preamble", "For value received [[Company Name]], a Delaware corporation (the “Company”), promises to pay to [[Holder Name]] or its assigns (“Holder”) the principal sum of $[[Principle Sum]] together with accrued and unpaid interest thereon, each due and payable on the date and in the manner set forth below. This convertible promissory note (the “Note”) is issued as part of a series of similar convertible promissory notes (collectively, the “Notes”) pursuant to the terms of that certain Convertible Promissory Note Purchase Agreement (as amended, the “Agreement”) dated as of [[Effective Date Convertible Promissory Note Purchase Agreement]]  to the persons and entities listed on the Schedule of Purchasers attached to the Agreement (collectively, the “Holders”). Capitalized terms used herein without definition shall have the meanings given to such terms in the Agreement.", "This is the preamble language for a convertible promissory .  It defines the company, holder, and the principle sum of the note.", "Covenant", false, NoParty}
	repayment := cctools.Clause {"Repayment", "All payments of interest and principal shall be in lawful money of the United States of America and shall be made pro rata among all Holders.  All payments shall be applied first to accrued interest, and thereafter to principal.  The outstanding principal amount of the Loan shall be due and payable on [[Maturity Date]] (the “Maturity Date”).", "This is the Repayment paragraph for a convertible promissory note", "Covenant", true, Holder}
	clauses := []cctools.Clause {CompanyName, HolderName, preamble, repayment} // Balance of remaining clauses to be added 
	convertibleNote := cctools.Document {"123", "Convertible Promissory Note", "This is a convertible promissory note based on the document open sourced by the incubator Tech Stars", 1, "Form Convertible Promissory Note", clauses, parties}
	var converted = cctools.DocumentToJSON(convertibleNote)
	fmt.Printf("%s\n", converted)
}