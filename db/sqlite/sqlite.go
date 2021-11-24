package sqlite

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"lbc/db"
	"lbc/model"
)

var _ db.Store = &SQlite{}

type SQlite struct {
	Conn *gorm.DB
}

func New(dbName string) db.Store {
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	conn.AutoMigrate(&model.User{})

	return &SQlite{
		Conn: conn,
	}
}

func (s *SQlite) GetUsers() ([]*model.User, error) {
	var us []*model.User
	return us, s.Conn.Find(&us).Error
}

func (s *SQlite) GetUserByID(id string) (*model.User, error) {
	u := new(model.User)
	return u, s.Conn.Where("uuid = ?", id).First(u).Error
}

func (s *SQlite) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	err := s.Conn.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *SQlite) DeleteUser(id string) error {
	return s.Conn.Delete("uuid = ?", id).Error
}

func (s *SQlite) AddUser(u *model.User) error {
	u.ID = uuid.NewString()
	return s.Conn.Create(u).Error
}

func (s *SQlite) SponeUser() {

	listUser := []*model.User{
		{
			ID:          "a0d09b91-0dcb-4aae-a0de-3a797891666c",
			FirstName:   "Bob",
			LastName:    "Picke",
			Email:       "bob.picke@domain.fr",
			Password:    "47625ed74cab8fbc0a8348f3df1feb07f87601e34d62bd12eb0d51616566fab5", // 123password
			AccessLevel: 0,
		},
		{
			ID:          "2c3759ce-1c08-47f4-9dcc-bc5fb61f4d66",
			FirstName:   "René",
			LastName:    "Leblanc",
			Email:       "r@google.com",
			Password:    "83671681bd1a452d3e5c88bb03595a336582928f41664755e8ac16ece230bbd4", // monmdp
			AccessLevel: 1,
		},
		{
			ID:          "05eca63d-c6e5-4a38-ba4f-65ff100c17bc",
			FirstName:   "Dennis",
			LastName:    "Richie",
			Email:       "dennis.richie@domain.fr",
			Password:    "47625ed74cab8fbc0a8340f3df1feb07f87601e34d62bd12eb0d51616566fab5", // ???
			AccessLevel: 1,
		},
	}
	s.Conn.Save(listUser)
}
