package main

import (
	"flag"
	"time"
)

/* flags */
var (
	// stringvar := flag.String("optionname", "defaultvalue", "description of the flag")
	cEntry      = flag.Bool("ce", false, "create a new entry")
	deleteEntry = flag.Int("de", -1, "delete an existing entry; default is -1")
	deleteProj  = flag.Int("dp", -1, "delete an existing project; default is -1")
	editProj    = flag.Int("ep", -1, "rename an existing project; default is empty string")
	markdown    = flag.Bool("md", false, "output all entries to markdown file")
	pdf         = flag.Bool("pdf", false, "output all entries to pdf file")
	start       = flag.String("s", "", "start date for date range")
	end         = flag.String("e", time.Now().Format("2006-01-02"), "end date for date range")
)
