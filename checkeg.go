func CheckUrL(url string, scode []int) (string, string) {

	for {
		resp, err := http.Get(url)
		if err != nil {
			print(err.Error())
		} else {
			for _, num := range scode {

				if num == resp.StatusCode {

					return "check is healthy", resp.Status

				} else {
					
					return "check is un healthy", resp.Status
				}
			}

		}
		/*t1 := time.NewTimer(5 * time.Second)
		<-t1.C*/
	}

}
