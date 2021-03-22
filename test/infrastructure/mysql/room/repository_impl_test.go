package room

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	infra "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
	testdata "github.com/karamaru-alpha/chat-go-server/test/testdata"
)

type repositoryImplTest struct {
	repositoryImpl domainModel.IRepository
	db             *gorm.DB
	mock           sqlmock.Sqlmock
}

// TestSave トークルームを永続化させる処理のテスト
func TestSave(t *testing.T) {

	// 永続化したいトークルームの準備
	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)
	room, err := domainModel.Create(roomTitle)
	assert.NoError(t, err)

	// モックの作成
	test := repositoryImplTest{}
	test.setupTest(t)

	test.mock.ExpectBegin()
	test.mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `rooms` (`id`,`title`)"),
	).WithArgs(ulid.ULID(room.ID).String(), string(room.Title)).WillReturnResult(sqlmock.NewResult(1, 1))
	test.mock.ExpectCommit()

	// 実行
	err = test.repositoryImpl.Save(room)
	assert.NoError(t, err)

	err = test.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	test.TeardownTest(t)
}

// TestFindAll トークルームの全件検索+再構築を行う処理のテスト
func TestFindAll(t *testing.T) {

	// 再構築したいトークルームの準備
	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)
	room, err := domainModel.Create(roomTitle)
	assert.NoError(t, err)

	// モックの作成
	test := repositoryImplTest{}
	test.setupTest(t)

	test.mock.ExpectQuery(("SELECT")).WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(ulid.ULID(room.ID).String(), string(room.Title)))

	// 実行
	output, err := test.repositoryImpl.FindAll()
	assert.NoError(t, err)

	assert.Equal(t, &[]domainModel.Room{*room}, output)

	err = test.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	test.TeardownTest(t)
}

func (r *repositoryImplTest) setupTest(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gormDB, err := gorm.Open("mysql", db)
	assert.NoError(t, err)
	gormDB.LogMode(true)

	repositoryImpl := infra.NewRepositoryImpl(gormDB)

	r.db = gormDB
	r.mock = mock
	r.repositoryImpl = repositoryImpl
}

func (r *repositoryImplTest) TeardownTest(t *testing.T) {
	r.db.Close()
}
