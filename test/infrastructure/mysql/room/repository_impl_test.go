package room

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	infra "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

type repositoryImplTester struct {
	repositoryImpl domainModel.IRepository
	db             *gorm.DB
	mock           sqlmock.Sqlmock
}

// TestSave トークルームを永続化させる処理のテスト
func TestSave(t *testing.T) {
	// モックの作成
	tester := repositoryImplTester{}
	tester.setupTest(t)

	tester.mock.ExpectBegin()
	tester.mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `rooms` (`id`,`title`)"),
	).WithArgs(tdString.Room.ID.Valid, tdString.Room.Title.Valid).WillReturnResult(sqlmock.NewResult(1, 1))
	tester.mock.ExpectCommit()

	// 実行
	err := tester.repositoryImpl.Save(&tdDomain.Room.Entity)
	assert.NoError(t, err)

	err = tester.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	tester.TeardownTest(t)
}

// TestFindAll トークルームの全件検索+再構築を行う処理のテスト
func TestFindAll(t *testing.T) {
	// モックの作成
	test := repositoryImplTester{}
	test.setupTest(t)

	test.mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(tdString.Room.ID.Valid, tdString.Room.Title.Valid))

	// 実行
	output, err := test.repositoryImpl.FindAll()
	assert.NoError(t, err)

	assert.Equal(t, &[]domainModel.Room{tdDomain.Room.Entity}, output)

	err = test.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	test.TeardownTest(t)
}

func (r *repositoryImplTester) setupTest(t *testing.T) {
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

func (r *repositoryImplTester) TeardownTest(t *testing.T) {
	r.db.Close()
}
