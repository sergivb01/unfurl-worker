package main

func main() {
	// bench(true)
	// time.Sleep(time.Second)
	// bench2(true)
	test()
}
//
// func bench(compress bool) {
// 	reader, err := getReaderFromURL(context.TODO(), "https://sergivos.dev", compress)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	info, err := extractFromReader(reader)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("\n%v\n", info)
//
// 	// fmt.Printf("%d\n", len(b))
// }
//
// func bench2(compress bool) {
// 	reader, err := getReaderFromURL(context.TODO(), "https://sergivos.dev", compress)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	info, err := meta.ExtractInfoFromReader(reader)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// doc, err := goquery.NewDocumentFromReader(reader)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// defer BenchmarkFunction(time.Now(), "bench2")
// 	//
// 	// info := PageInfo{}
// 	// if err := getPageData(doc, &info); err != nil {
// 	// 	panic(err)
// 	// }
//
// 	fmt.Printf("\n%v\n", info)
// }
