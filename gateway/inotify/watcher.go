package inotify

import (
	notify "bytetrade.io/web3os/fs-lib/jfsnotify"
	"fmt"
	"io/fs"
	"math"
	"path/filepath"
	"sync"
	"time"
	"wzinc/dify"

	//notify "github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
)

var WatchDir string

// var watcher *jfsnotify.Watcher
var watcher *notify.Watcher = nil

func WatchPath(pathMap map[string]PathComparison) {
	// Create a new watcher.
	var err error
	if watcher == nil {
		watcher, err = notify.NewWatcher("DifyWatch")
		//watcher, err = notify.NewWatcher()
		if err != nil {
			panic(err)
		}

		// Start listening for events.
		go dedupLoop(watcher)
		//log.Info().Msgf("watching path %s", strings.Join(addPaths, ","))
	}

	for keyPath, pathComparison := range pathMap {
		// 遍历 Delete 列表中的每个值，并进行一些操作
		for _, deleteDatasetID := range pathComparison.Delete {
			// 执行操作，你可以在这里添加你的操作逻辑
			fmt.Println(keyPath, "Performing operation on Delete value:", deleteDatasetID)

			err = filepath.Walk(keyPath, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					if pathComparison.Op == "delete" {
						err = watcher.Remove(path)
						if err != nil {
							fmt.Println("watcher remove error:", err)
							return err
						}
					}
				} else {
					log.Info().Msgf("push indexer task delete %s", path)
					dify.DatasetsDeleteDocument(path, deleteDatasetID)
				}
				return nil
			})
			if err != nil {
				panic(err)
			}
		}

		// 遍历 Add 列表中的每个值，并进行一些操作
		for _, addDatasetID := range pathComparison.Add {
			// 执行操作，你可以在这里添加你的操作逻辑
			fmt.Println(keyPath, "Performing operation on Add value:", addDatasetID)

			err = filepath.Walk(keyPath, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					if pathComparison.Op == "add" {
						err = watcher.Add(path)
						if err != nil {
							fmt.Println("watcher add error:", err)
							return err
						}
					}
				} else {
					err = dify.DatasetsAddDocument(path, addDatasetID)
					if err != nil {
						log.Error().Msgf("udpate or input doc err %v", err)
					}
				}
				return nil
			})
			if err != nil {
				panic(err)
			}
		}
	}

	printTime("ready; press ^C to exit")
}

func dedupLoop(w *notify.Watcher) {
	var (
		// Wait 1000ms for new events; each new event resets the timer.
		waitFor = 1000 * time.Millisecond

		// Keep track of the timers, as path → timer.
		mu           sync.Mutex
		timers       = make(map[string]*time.Timer)
		pendingEvent = make(map[string]notify.Event)

		// Callback we run.
		printEvent = func(e notify.Event) {
			log.Info().Msgf("handle event %v %v", e.Op.String(), e.Name)

			// Don't need to remove the timer if you don't have a lot of files.
			mu.Lock()
			delete(pendingEvent, e.Name)
			delete(timers, e.Name)
			mu.Unlock()
		}
	)

	for {
		select {
		// Read from Errors.
		case err, ok := <-w.Errors:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}
			printTime("ERROR: %s", err)
		// Read from Events.
		case e, ok := <-w.Events:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				log.Warn().Msg("watcher event channel closed")
				return
			}
			if e.Has(notify.Chmod) {
				continue
			}
			log.Debug().Msgf("pending event %v", e)
			// Get timer.
			mu.Lock()
			_, ok = pendingEvent[e.Name]
			if !ok {
				pendingEvent[e.Name] = e
			} else {
				var temp notify.Event
				temp.Name = e.Name
				temp.Op = pendingEvent[e.Name].Op | e.Op
				pendingEvent[e.Name] = temp
			}
			t, ok := timers[e.Name]
			mu.Unlock()

			// No timer yet, so create one.
			if !ok {
				t = time.AfterFunc(math.MaxInt64, func() {
					mu.Lock()
					ev := pendingEvent[e.Name]
					mu.Unlock()
					printEvent(ev)
					err := handleEvent(ev)
					if err != nil {
						log.Error().Msgf("handle watch file event error %s", err.Error())
					}
				})
				t.Stop()

				mu.Lock()
				timers[e.Name] = t
				mu.Unlock()
			}

			// Reset the timer for this path, so it will start from 100ms again.
			t.Reset(waitFor)
		}
	}
}

func handleEvent(e notify.Event) error {
	fmt.Println(e.Name, e.Op, e)
	targetDatasetIDs := FindBaseFromPath(e.Name)
	if e.Has(notify.Remove) || e.Has(notify.Rename) {
		for _, targetDatasetID := range targetDatasetIDs {
			dify.DatasetsDeleteDocument(e.Name, targetDatasetID)
		}
	}
	if e.Has(notify.Create) || e.Has(notify.Write) {
		for _, targetDatasetID := range targetDatasetIDs {
			dify.DatasetsAddDocument(e.Name, targetDatasetID)
		}
	}

	if e.Has(notify.Create) {
		err := filepath.Walk(e.Name, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				//add dir to watch list
				err = watcher.Add(path)
				if err != nil {
					log.Error().Msgf("watcher add error:%v", err)
				}
			} else {
				for _, targetDatasetID := range targetDatasetIDs {
					err = dify.DatasetsAddDocument(path, targetDatasetID)
					if err != nil {
						log.Error().Msgf("update or input doc error %v", err)
					}
				}
			}
			return nil
		})
		if err != nil {
			log.Error().Msgf("handle create file error %v", err)
		}
		return nil
	}

	if e.Has(notify.Write) { // || e.Has(notify.Chmod) {
		for _, targetDatasetID := range targetDatasetIDs {
			dify.DatasetsAddDocument(e.Name, targetDatasetID)
		}
	}
	return nil
}

func printTime(s string, args ...interface{}) {
	log.Info().Msgf(time.Now().Format("15:04:05.0000")+" "+s+"\n", args...)
}
