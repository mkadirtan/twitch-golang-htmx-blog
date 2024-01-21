package templates

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"time"
)

// watchThemeChanges intentionally does not panic, this is a development feature
func watchThemeChanges(themesRoot string) {
	w := watcher.New()
	w.SetMaxEvents(1)

	go func() {
		for {
			select {
			case _ = <-w.Event:
				err := registerThemes(themesRoot)
				if err != nil {
					fmt.Println(err.Error())
				}

			case err := <-w.Error:
				fmt.Println(err)
			case <-w.Closed:
				fmt.Println("closed...")
				return
			}
		}
	}()

	err := w.AddRecursive(themesRoot)
	if err != nil {
		panic(err)
	}

	go func() {
		err = w.Start(time.Millisecond * 100)
		if err != nil {
			panic(err)
		}
	}()
}
