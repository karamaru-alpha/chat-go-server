package message

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	repoImpl "github.com/karamaru-alpha/chat-go-server/infrastructure/repository/message"

	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdCommonString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/common"
	tdMessageString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/message"
)

type testHandler struct {
	repositoryImpl domain.IRepository

	db          *gorm.DB
	redisClient *redis.Client
	mock        sqlmock.Sqlmock
}

// TestSave メッセージを永続化させる処理のテスト
func TestSave(t *testing.T) {
	t.Parallel()

	// モックの作成
	var tester testHandler
	tester.setupTest(t)

	tester.mock.ExpectBegin()
	tester.mock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO `messages` (`id`,`room_id`,`body`)"),
	).WithArgs(tdCommonString.ULID.Valid, tdCommonString.ULID.Valid, tdMessageString.Body.Valid).WillReturnResult(sqlmock.NewResult(1, 1))
	tester.mock.ExpectCommit()

	// 実行
	err := tester.repositoryImpl.Save(context.TODO(), tdMessageDomain.Entity)
	assert.NoError(t, err)

	err = tester.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	tester.TeardownTest(t)
}

// TestFindAll 該当トークルームにあるメッセージの全件検索+再構築を行う処理のテスト
func TestFindAll(t *testing.T) {
	t.Parallel()

	var tester testHandler
	tester.setupTest(t)

	tester.mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "room_id", "body"}).AddRow(tdCommonString.ULID.Valid, tdCommonString.ULID.Valid, tdMessageString.Body.Valid))

	// 実行
	output, err := tester.repositoryImpl.FindAll(tdRoomDomain.ID)
	assert.NoError(t, err)

	assert.Equal(t, []domain.Message{tdMessageDomain.Entity}, output)

	err = tester.mock.ExpectationsWereMet()
	assert.NoError(t, err)

	tester.TeardownTest(t)
}

// TODO
// TestSubscribe メッセージの監視処理をテスト
func TestSubscribe(t *testing.T) {
	t.Parallel()

	assert.Equal(t, true, true)
}

func (r *testHandler) setupTest(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	r.mock = mock

	gormDB, err := gorm.Open("mysql", db)
	assert.NoError(t, err)
	gormDB.LogMode(true)
	r.db = gormDB

	mockRedis, err := miniredis.Run()
	assert.NoError(t, err)
	r.redisClient = redis.NewClient(
		&redis.Options{
			Addr: mockRedis.Addr(),
		},
	)

	repositoryImpl := repoImpl.NewRepositoryImpl(
		gormDB,
		r.redisClient,
	)
	r.repositoryImpl = repositoryImpl
}

func (r *testHandler) TeardownTest(t *testing.T) {
	r.db.Close()
	r.redisClient.Close()
}
