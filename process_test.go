package main

import (
    "testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
    name string
    in   string
    resultWanted string
}{
    {
        name: "Hexadécimal",
        in:   "1E (hex) files",
        resultWanted: "30 files",
    },
    {
        name: "Majuscule (up)",
        in:   "go (up) !",
        resultWanted: "GO!",
    },
	{
        name:         "Transformation binaire",
        in:           "101 (bin) books",
        resultWanted: "5 books",
    },
    {
        name:         "Ponctuations collées",
        in:           "I was thinking ... You , my friend !",
        resultWanted: "I was thinking... You, my friend!",
    },
    {
        name:         "Règle a/an devant voyelle",
        in:           "There is a apple on a table",
        resultWanted: "There is an apple on a table",
    },
	{
    name:         "Capitalise les 2 mots précédents",
    in:           "welcome to tokyo (cap, 2) !",
    resultWanted: "welcome To Tokyo!",
	},
	{
        name:         "Apostrophes",
        in:           "You are ' awesome '",
        resultWanted: "You are 'awesome'",
    },
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        got := Process(tt.in)
        if got != tt.resultWanted {
            t.Errorf("Process(%q) = %q; attendu %q", tt.in, got, tt.resultWanted)
        }
    })
}
}
