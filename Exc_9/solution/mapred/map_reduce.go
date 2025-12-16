package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce
func (mr *MapReduce) Run(input []string) map[string]int {
	// input processed by mapper -> keyvalue pairs
	mapOut := make(chan []KeyValue, len(input))
	var wgMap sync.WaitGroup
	// start go routine
	for _, line := range input {
		wgMap.Add(1)
		go func(t string) {
			defer wgMap.Done()
			mapOut <- mr.wordCountMapper(t)
		}(line)
	}
	//close map output channel when done
	go func() {
		wgMap.Wait()
		close(mapOut)
	}()
	// all keyvalues from map phase
	var allKVs []KeyValue
	for kvs := range mapOut {
		allKVs = append(allKVs, kvs...)
	}
	// group values by keys
	grouped := make(map[string][]int)
	for _, kv := range allKVs {
		grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
	}
	type redRes struct {
		kv KeyValue
	}
	// reduce key by summing values
	redOut := make(chan KeyValue, len(grouped))
	var wgRed sync.WaitGroup
	for k, vals := range grouped {
		wgRed.Add(1)
		// start goroutine per key
		go func(key string, values []int) {
			defer wgRed.Done()
			redOut <- mr.wordCountReducer(key, values)
		}(k, vals)
	}
	// close channel
	go func() {
		wgRed.Wait()
		close(redOut)
	}()
	// result map
	results := make(map[string]int)
	for kv := range redOut {
		results[kv.Key] = kv.Value
	}
	return results
}

// filter out all special chars and numericals (RegEx)
var nonLetters = regexp.MustCompile(`[^a-zA-Z]+`)

// map step for word counting, filter words
func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	s := strings.ToLower(text)
	s = nonLetters.ReplaceAllString(s, " ")
	words := strings.Fields(s)
	out := make([]KeyValue, 0, len(words))
	for _, w := range words {
		out = append(out, KeyValue{Key: w, Value: 1})
	}
	return out
}

// reduce step, sum values w key then return total freq
func (mr *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return KeyValue{Key: key, Value: sum}
}
