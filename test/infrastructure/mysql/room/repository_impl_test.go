package room

// TODO implment repository impl test

// import (
// 	"regexp"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/jinzhu/gorm"
// 	"github.com/oklog/ulid"
// 	"github.com/stretchr/testify/assert"

// 	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
// 	infra "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
// 	testdata "github.com/karamaru-alpha/chat-go-server/test/testdata"
// )

// type repositoryImplTest struct {
// 	repositoryImpl domainModel.IRepository
// 	db             *gorm.DB
// 	mock           sqlmock.Sqlmock
// }

// func TestSave(t *testing.T) {
// 	t.Parallel()

// 	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
// 	assert.NoError(t, err)
// 	room, err := domainModel.Create(roomTitle)
// 	assert.NoError(t, err)

// 	test := repositoryImplTest{}
// 	test.setupTest(t)

// 	test.mock.ExpectBegin()

// 	test.mock.ExpectExec(
// 		regexp.QuoteMeta("INSERT INTO \"rooms\" (\"id\", \"title\") VALUES (?, ?)"),
// 	).WithArgs(ulid.ULID(room.ID).String(), string(room.Title))

// 	test.mock.ExpectCommit()

// 	tests := []struct {
// 		title   string
// 		input   *domainModel.Room
// 		isError bool
// 	}{
// 		{
// 			title:   "【正常系】",
// 			input:   room,
// 			isError: false,
// 		},
// 	}

// 	for _, td := range tests {
// 		td := td

// 		err := test.repositoryImpl.Save(td.input)

// 		if td.isError {
// 			assert.Error(t, err)
// 		} else {
// 			assert.NoError(t, err)
// 		}
// 	}

// 	test.TeardownTest(t)
// }

// func (r *repositoryImplTest) setupTest(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)

// 	gormDB, err := gorm.Open("mysql", db)
// 	assert.NoError(t, err)

// 	repositoryImpl := infra.NewRepositoryImpl(gormDB)

// 	r.db = gormDB
// 	r.mock = mock
// 	r.repositoryImpl = repositoryImpl
// }

// func (r *repositoryImplTest) TeardownTest(t *testing.T) {
// 	err := r.db.Close()
// 	assert.NoError(t, err)
// }
