**Code to test it:**


	cache := cache.New()

	cache.Set("userId", 42)
	userId, err := cache.Get("userId")

	if err == nil {
		fmt.Println(userId)
	}

	cache.Delete("userId")
	userId, err = cache.Get("userId")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(userId)
