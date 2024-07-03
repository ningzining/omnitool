package async

import (
	"log"
	"sync"
)

type function func()

func SaveGo(f function) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		f()
	}()
}

func SaveGos(fs ...function) {
	for _, f := range fs {
		SaveGo(f)
	}
}

func WaitGroup(fs ...function) {
	var wg sync.WaitGroup
	wg.Add(len(fs))
	for _, f := range fs {
		SaveGo(func() {
			defer wg.Done()
			f()
		})
	}

	wg.Wait()
}
