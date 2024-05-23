package main

// before around 110 seconds
//after 1.5 seconds

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

func main() {
	rand.Seed(time.Now().Unix())
	startTime := time.Now()
	usersCount, workersCount := 100, 100

	wg := &sync.WaitGroup{}

	jobs := make(chan User, usersCount)

	// 10 goroutines will generate 10 users each
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			generateUsers(10*i, 10*(i+1), jobs)
		}()
	}
	wg.Wait()
	close(jobs)

	for i := 1; i <= workersCount; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			saveUserInfo(i, jobs)
		}()
	}
	wg.Wait()

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUserInfo(id int, jobs <-chan User) {
	fmt.Printf("Worker #%d started working\n", id)

	//reading from channel each user and writing to a file
	for user := range jobs {
		filename := fmt.Sprintf("users/uid%d.txt", user.id)
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}

		file.WriteString(user.getActivityInfo())
		time.Sleep(time.Millisecond * 100) //change to Second
	}
}

func generateUsers(fromIndex, tillIndex int, jobs chan<- User) {
	//each generated user is written to channel
	for i := fromIndex; i < tillIndex; i++ {
		jobs <- User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@company.com", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
		fmt.Printf("generated user %d\n", i+1)
		time.Sleep(time.Millisecond * 100)
	}
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}
