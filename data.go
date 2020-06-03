package cterm

type ContentListSourceDataType struct {
	Title string
	ID    int
	Msg   []string
}

type IndexListSourceDataType struct {
	Name        string
	ID          int
	Desc        string
	ContentList []ContentListSourceDataType
}

func FetchData() {
	indexListSourceData = readFromDataFile()
	// indexListSourceData = []IndexListSourceDataType{
	// 	{
	// 		Name: "工作",
	// 		ID:   1,
	// 		Desc: "hz的生活",
	// 		ContentList: []ContentListSourceDataType{
	// 			{"地址", 1, []string{"江干", "西湖"}},
	// 			{"公司", 2, []string{"mfe", "yz"}},
	// 		},
	// 	},
	// 	{Name: "生活", ID: 2, Desc: "yz的生活", ContentList: []ContentListSourceDataType{{"同事", 1, []string{"morgan", "fantasy"}}}},
	// 	{Name: "休闲", ID: 3, Desc: "hz的生活", ContentList: []ContentListSourceDataType{{"吃饭地", 1, []string{"大东北", "多伦多", "大渝火锅"}}}},
	// }
}
