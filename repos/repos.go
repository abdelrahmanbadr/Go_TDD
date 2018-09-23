package repos

import (
	"tdd/models"
)

type ReposClient struct {
	Repos []models.Repo
	Err   error
}

// Fake repository data
func (c *ReposClient) Init() {
	c.Repos = []models.Repo{
		{"user1", "user1 description"},
		{"user2", "user2 description"},
		{"user3", "user3 description"},
	}

}

//func (c *ReposClient) Get(user string) (models.Repo, error) {
//	c.Init()
//	for _,v := range c.Repos{
//		if v.Name == user{
//			return v,nil
//		}
//	}
//	c.Err = errors.New("User isn't exit")
//	return models.Repo{}, c.Err
//}
func (c ReposClient) Get(user string) ([]models.Repo, error) {
	c.Init()

	return c.Repos, c.Err
}
