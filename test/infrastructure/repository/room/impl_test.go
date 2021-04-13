package room

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	repoImpl "github.com/karamaru-alpha/chat-go-server/infrastructure/repository/room"

	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdCommonString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/common"
	tdRoomString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/room"
)

type testHandler struct {
	repositoryImpl domain.IRepository
	db             *gorm.DB
	mock           sqlmock.Sqlmock
}

// TestSave トークルームを永続化させる処理のテスト
func TestSave(t *testing.T) {
	t.Parallel()

	var tester testHandler
	tester.setupTest(t)

	tester.mock.ExpectBegin()
	tester.mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `rooms` (`id`,`title`)"),
	).WithArgs(tdCommonString.ULID.Valid, tdRoomString.Title.Valid).WillReturnResult(sqlmock.NewResult(1, 1))
	tester.mock.ExpectCommit()

	// 実行
	err := tester.repositoryImpl.Save(tdDomain.Entity)
	assert.NoError(t, err)

	err = tester.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	tester.teardownTest(t)
}

// TestFindAll トークルームの全件検索+再構築を行う処理のテスト
func TestFindAll(t *testing.T) {
	t.Parallel()

	var tester testHandler
	tester.setupTest(t)

	tester.mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(tdCommonString.ULID.Valid, tdRoomString.Title.Valid))

	// 実行
	output, err := tester.repositoryImpl.FindAll()
	assert.NoError(t, err)

	assert.Equal(t, []domain.Room{tdDomain.Entity}, output)

	err = tester.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	tester.teardownTest(t)
}

// TestFind トークルームをIDから一件取得・再構築する処理のテスト
func TestFind(t *testing.T) {
	t.Parallel()

	// モックの作成
	test := testHandler{}
	test.setupTest(t)

	test.mock.ExpectQuery("SELECT").WithArgs(tdCommonString.ULID.Valid).WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(tdCommonString.ULID.Valid, tdRoomString.Title.Valid))

	// 実行
	output, err := test.repositoryImpl.Find(tdDomain.ID)
	assert.NoError(t, err)

	assert.Equal(t, tdDomain.Entity, output)

	err = test.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	test.teardownTest(t)
}

// TestFindByTitle トークルームをタイトルから一件取得・再構築する処理のテスト
func TestFindByTitle(t *testing.T) {
	t.Parallel()

	var tester testHandler
	tester.setupTest(t)

	tester.mock.ExpectQuery("SELECT").WithArgs(tdRoomString.Title.Valid).WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(tdCommonString.ULID.Valid, tdRoomString.Title.Valid))

	// 実行
	output, err := tester.repositoryImpl.FindByTitle(tdDomain.Title)
	assert.NoError(t, err)

	assert.Equal(t, tdDomain.Entity, output)

	err = tester.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	tester.teardownTest(t)
}

func (r *testHandler) setupTest(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	gormDB, err := gorm.Open("mysql", db)
	assert.NoError(t, err)
	gormDB.LogMode(true)

	repositoryImpl := repoImpl.NewRepositoryImpl(gormDB)

	r.db = gormDB
	r.mock = mock
	r.repositoryImpl = repositoryImpl
}

func (r *testHandler) teardownTest(t *testing.T) {
	r.db.Close()
}
