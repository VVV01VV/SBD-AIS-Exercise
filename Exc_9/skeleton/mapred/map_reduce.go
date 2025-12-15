package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce

// Run executes a concurrent MapReduce word count:
// 1. Map phase (concurrent): each input line becomes KeyValue(word, 1)
// 2. Shuffle phase: group values by key
// 3. Reduce phase (concurrent): sum values per key
func (mr MapReduce) Run(input []string) map[string]int {
	// Map phase (concurrent)
	mapOut := make(chan []KeyValue)

	var mapWG sync.WaitGroup
	mapWG.Add(len(input))

	for _, line := range input {
		go func(text string) {
			defer mapWG.Done()
			kvs := mr.wordCountMapper(text)
			mapOut <- kvs
		}(line)
	}

	// Close channel when all mappers are done
	go func() {
		mapWG.Wait()
		close(mapOut)
	}()

	// Shuffle / group phase
	grouped := make(map[string][]int)
	for kvs := range mapOut {
		for _, kv := range kvs {
			grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
		}
	}

	// Reduce phase (concurrent)
	results := make(map[string]int)

	type redResult struct {
		kv KeyValue
	}
	redOut := make(chan KeyValue)

	var redWG sync.WaitGroup
	redWG.Add(len(grouped))

	for key, values := range grouped {
		go func(k string, vals []int) {
			defer redWG.Done()
			kv := mr.wordCountReducer(k, vals)
			redOut <- kv
		}(key, values)
	}

	// Close reducer output when done
	go func() {
		redWG.Wait()
		close(redOut)
	}()

	for kv := range redOut {
		results[kv.Key] = kv.Value
	}

	return results
}

// wordCountMapper turns a string into KeyValue(word, 1).
// It lowercases and filters out all special chars and numerics via regex.
func (mr MapReduce) wordCountMapper(text string) []KeyValue {
	// Keep only ascii letters and whitespace, everything else becomes space
	re := regexp.MustCompile(`[^a-zA-Z\s]+`)
	clean := re.ReplaceAllString(text, " ")
	clean = strings.ToLower(clean)

	words := strings.Fields(clean)

	out := make([]KeyValue, 0, len(words))
	for _, w := range words {
		out = append(out, KeyValue{Key: w, Value: 1})
	}
	return out
}

// wordCountReducer sums the integer values for a given key
func (mr MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return KeyValue{Key: key, Value: sum}
}
