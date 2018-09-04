package service

import (
	"fmt"
	"sync"

	"github.com/scaat/apiserver/model"
	"github.com/scaat/apiserver/util"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)

	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, user := range users {
		wg.Add(1)
		go func(user *model.UserModel) {
			defer wg.Done()

			shortID, err := util.GenShortID()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[user.ID] = &model.UserInfo{
				ID:        user.ID,
				Username:  user.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortID),
				Password:  user.Password,
				CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(user)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}
