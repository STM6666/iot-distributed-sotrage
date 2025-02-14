package File_Index

import (
	"fmt"
	"testing"
	"time"
)

func TestBuildTraverser(t *testing.T) {
	filepath := string("../Block_Chain/backup.txt")
	tree := BuildTraverser(filepath)

	if tree.Empty() == true {
		fmt.Errorf("Fail to create tree!\n")
	}
	fmt.Println(tree)
	data := QueryData(tree, "MacBook", time.Now(), time.Now())
	fmt.Println(data)
}
