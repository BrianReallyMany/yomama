package ui

import (
	"fmt"
	"github.com/BrianReallyMany/yomama/dozens"
)

type MamaController struct {
}

// Instantiate a new yomama controller
func MakeMamaController() *MamaController {
    return &MamaController{}
}

func (c *MamaController) Dozens() string {
    return dozens.RandomDozens()
}

func (c *MamaController) PrepFiles(args []string) {
	fmt.Println("Prepping files ...")
	fmt.Println("Files prepped. No, for real.")
}
