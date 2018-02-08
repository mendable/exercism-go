// Letter provides methods for calculating frequency of letters in text
// using either sequential or concurrent processing.
package letter

import "sync"

// FreqMap represents unique runes and the number of occurances.
type FreqMap map[rune]int

// Frequency processes a string using serial computation, and returns a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// Counter is used for 3rd possible solution containing shared mutex and waitgroup
// for concurrent processing.
type Counter struct {
	result FreqMap
	mutex  *sync.Mutex
	wg     *sync.WaitGroup
}

// ConcurrentFrequency processes a slice of strings using concurrent computation,
// and returns a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	result := FreqMap{}

	//   // Solution 1 - inline closure
	//   for _, text := range texts {
	//     wg.Add(1)
	//     go countFreq(text, result)
	//       defer wg.Done()
	//       for _, r := range text {
	//         mutex.Lock()
	//         f[r]++
	//         mutex.Unlock()
	//       }
	//     }(text, result)
	//   }

	//  // Solution 2 - defer to goroutine, passing in multiple vars.
	//  for _, text := range texts {
	//    wg.Add(1)
	//    go countFreq(text, result, &wg, &mutex)
	//  }

	// Solution 3 - use new Counter type to do the counting work.
	counter := Counter{wg: &wg, mutex: &mutex, result: result}
	for _, text := range texts {
		wg.Add(1)
		go counter.count(text)
	}

	wg.Wait()

	return result
}

// // Soution 2 - func that can be run deferred
// func countFreq(text string, f FreqMap, wg *sync.WaitGroup, mutex *sync.Mutex) {
//   defer wg.Done()
//   for _, r := range text {
//     mutex.Lock()
//     f[r]++
//     mutex.Unlock()
//   }
// }

func (w *Counter) count(text string) {
	defer w.wg.Done()
	for _, r := range text {
		w.mutex.Lock()
		w.result[r]++
		w.mutex.Unlock()
	}
}
