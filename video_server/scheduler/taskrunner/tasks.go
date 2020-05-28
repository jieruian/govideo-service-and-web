package taskrunner

import (
	"errors"
	"govideo/video_server/scheduler/dbops"
	"log"
	"os"
	"sync"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_OIR + vid + ".mp4")
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecuor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
FORLOOP:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DeleteVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break FORLOOP
		}
	}
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
