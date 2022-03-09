package models

//go:generate reform

// Sysdiagrams represents a row in sysdiagrams table.
//reform:sysdiagrams
type Sysdiagrams struct {
	Name        string `reform:"name"`
	PrincipalID int32  `reform:"principal_id"`
	DiagramID   int32  `reform:"diagram_id,pk"`
	Version     *int32 `reform:"version"`
	Definition  []byte `reform:"definition"`
}
