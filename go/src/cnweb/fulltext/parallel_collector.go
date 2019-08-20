//
// Functions for retrieving text text matches in parallel from text that are 
// either the file or in a remote object store
//
package fulltext

import (
	"github.com/alexamies/cnweb/applog"
)

type Job struct {
	key string
	results chan<- Result
}

type Result struct {
	key string
	dm DocMatch
}

 // A long operation, needs to be done in parallel
func (job Job) Do(loader TextLoader, queryTerms []string) {
	dm := DocMatch{
		PlainTextFile: job.key,
	}
	mt, err := loader.GetMatching(job.key, queryTerms)
	if err != nil {
		// log and move on
		applog.Info("job.Do error ", job.key, err)
	} else {
		dm.MT = mt
	}
	result := Result{
		key: job.key,
		dm: dm,
	}
	job.results <- result
}

func addJobs(jobs chan<- Job, keys []string, results chan<- Result) {
	for _, key := range keys {
		job := Job{
			key: key,
			results: results,
		}
		jobs <- job
	}
	close(jobs)
}

func collectDocs(done <-chan struct{}, results chan Result, keys []string) map[string]DocMatch {
	applog.Info("fulltext.collectDocs\n")
	matches := map[string]DocMatch{}
	workers := len(keys)
	for working := workers; working > 0; {
		select {
		case result := <- results:
			matches[result.key] = result.dm
			applog.Info("fulltext.collectDocs", result.key, result.dm)
		case <-done:
			working--
		}
	}
DONE:
	for {
		select {
			case result := <- results:
				matches[result.key] = result.dm
				applog.Info("fulltext.collectDocs", result.key, result.dm)
			default:
				break DONE 
		}
	}
	return matches
}

func getDoc(done chan struct{}, loader TextLoader, key string, queryTerms []string, jobs <-chan Job) {
	//applog.Info("fulltext.getDoc:", key)
	for job := range jobs {
		job.Do(loader, queryTerms)
	}
	done <- struct{}{}
}

func GetMatches(keys []string, queryTerms []string) map[string]DocMatch {
	applog.Info("GetMatches")
	loader := getLoader()
	jobs := make(chan Job, len(keys))
	results := make(chan Result, len(keys))
	done := make(chan struct{}, len(keys))
	addJobs(jobs, keys, results)
	for _, key := range keys {
		go getDoc(done, loader, key, queryTerms, jobs)
	}
	return collectDocs(done, results, keys)
}
