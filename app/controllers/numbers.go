package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/abhi11/numbers/set"
	"github.com/abhi11/numbers/util"
)

type Number struct {
	Numbers []int `json:"numbers"`
}

// controller to fetch all urls and give a sorted result
// Does processing parallely, GET calss and sorting happens parallely
func SortNumbers(w http.ResponseWriter, r *http.Request) {
	s := set.NewSet()
	n := Number{}

	vals := util.GetParamValues(r, "u")
	if len(vals) < 1 {
		n.Numbers = []int{}
		data, _ := util.ToJSONBytes(n)
		util.WriteResponse(w, http.StatusOK, data)
		return
	}

	log.Println("URLs: ", vals)
	vals = util.GetValidURLs(vals)
	log.Println("Valid URLs: ", vals)

	// declare channel and wait group, timer
	var wg sync.WaitGroup
	num := make(chan []int)
	doneWait := make(chan bool)
	doneSort := make(chan bool)
	tick := time.NewTimer(500 * time.Millisecond)

	for _, v := range vals {
		wg.Add(1)
		go getNumbers(v, num, &wg)
	}

	go wait(&wg, doneWait)
	go sortNums(num, s, doneSort)

	// wait for the sorting to take place;
	// and get calls carry on parallely once
	// one get call is succesfull the channel num is updated
	for {
		toBreak := false
		select {
		case <-doneWait:
			log.Println("get calls completed in time... closing channel num")
			close(num)

		case <-doneSort:
			log.Println("done sorting before time")
			toBreak = true

		case <-tick.C:
			log.Println("timer up")
			toBreak = true
		}

		// break from the loop show result
		if toBreak {
			break
		}
	}

	/* Write numbers */
	n.Numbers = s.Sort()
	data, _ := util.ToJSONBytes(n)
	util.WriteResponse(w, http.StatusOK, data)
}

// signals once all get calls are done
func wait(w *sync.WaitGroup, d chan bool) {
	w.Wait()
	d <- true
}

// Makes GET Call to fetch the numbers from URLs
func getNumbers(u string, num chan []int, w *sync.WaitGroup) {
	n := &Number{}
	_, err := util.HttpGetAndRead(u, n)
	if err != nil {
		log.Println("Error getting numbers from url: ", u, err)
		w.Done()
		return
	}
	num <- n.Numbers
	w.Done()
}

// Takes a num channel and keeps sorting by appending the
// latest []int in num channel in Number struct
func sortNums(num chan []int, s *set.Set, d chan bool) {
	for {
		numFromChan, more := <-num
		if more {
			s.AddList(numFromChan)
			continue
		}
		// else break
		break
	}
	// signal sortDone
	d <- true
}
