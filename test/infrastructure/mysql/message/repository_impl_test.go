package message

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	infra "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/message"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

type repositoryImplTester struct {
	repositoryImpl domain.IRepository
	db             *gorm.DB
	mock           sqlmock.Sqlmock
}

// TestCreate メッセージを永続化させる処理のテスト
func TestSave(t *testing.T) {
	t.Parallel()

	// モックの作成
	tester := repositoryImplTester{}
	tester.setupTest(t)

	tester.mock.ExpectBegin()
	tester.mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `messages` (`id`,`room_id`,`body`)"),
	).WithArgs(tdString.Message.ID.Valid, tdString.Room.ID.Valid, tdString.Message.Body.Valid).WillReturnResult(sqlmock.NewResult(1, 1))
	tester.mock.ExpectCommit()

	// 実行
	err := tester.repositoryImpl.Create(&tdMessageDomain.Message.Entity)
	assert.NoError(t, err)

	err = tester.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	tester.TeardownTest(t)
}

// TestFindAll 該当トークルームにあるメッセージの全件検索+再構築を行う処理のテスト
func TestFindAll(t *testing.T) {
	t.Parallel()

	// モックの作成
	test := repositoryImplTester{}
	test.setupTest(t)

	test.mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "room_id", "body"}).AddRow(tdString.Message.ID.Valid, tdString.Room.ID.Valid, tdString.Message.Body.Valid))

	// 実行
	output, err := test.repositoryImpl.FindAll(&tdRoomDomain.Room.ID)
	assert.NoError(t, err)

	assert.Equal(t, &[]domain.Message{tdMessageDomain.Message.Entity}, output)

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
