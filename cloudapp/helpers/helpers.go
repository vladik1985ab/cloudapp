package helpers

import "cloudapp/cloudapp/model"
import (
	"cloudapp/cloudapp/service"
	"math/rand"
)

func Shuffle(buildpack *[]model.BuildpackCustom) {
	for i := 1; i < len(*buildpack); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			(*buildpack)[r], (*buildpack)[i] = (*buildpack)[i], (*buildpack)[r]
		}
	}
}

func BuildpackCustomObjects() []model.BuildpackCustom {
	client := service.GetClient();
	buildpack, _ := client.ListBuildpacks()
	count := len(buildpack);
	b := make([]model.BuildpackCustom, count)
	//make JSON similar to the example.
	//in normal case I would do just serializable without copy the object , here I'm not sure how to do it
	for i := 0; i < count; i++ {
		b[i].Enabled = buildpack[i].Enabled
		b[i].Filename = buildpack[i].Filename
		b[i].Locked = buildpack[i].Locked
		b[i].Name = buildpack[i].Name
	}

	return b;
}


// SortByName sorts planets by axis.
type SortByName []model.BuildpackCustom

func (a SortByName) Len() int {
	return len(a)
}
func (a SortByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

