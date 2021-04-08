package room

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	domainService "github.com/karamaru-alpha/chat-go-server/domain/service/room"
	mockDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
)

type domainServiceTester struct {
	domainService domainService.IDomainService
	repository    *mockDomain.MockIRepository
}

// TestExists トークルームの重複チェックを担うドメインサービスのテスト
func TestExists(t *testing.T) {
	t.Parallel()

	var tester domainServiceTester
	tester.setupTest(t)

	tests := []struct {
		title     string
		before    func()
		input     *domain.Room
		expected1 bool
		expected2 error
	}{
		{
			title: "【正常系】該当タイトルのトークルームが存在しない",
			before: func() {
				tester.repository.EXPECT().FindByTitle(&tdDomain.Room.Title).Return(nil, nil)
			},
			input:     &tdDomain.Room.Entity,
			expected1: false,
			expected2: nil,
		},
		{
			title: "【正常系】該当タイトルのトークルームが存在する",
			before: func() {
				tester.repository.EXPECT().FindByTitle(&tdDomain.Room.Title).Return(&tdDomain.Room.Entity, nil)
			},
			input:     &tdDomain.Room.Entity,
			expected1: true,
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Exists:"+td.title, func(t *testing.T) {
			t.Parallel()

			td.before()

			output1, output2 := tester.domainService.Exists(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func (d *domainServiceTester) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	d.repository = mockDomain.NewMockIRepository(ctrl)
	d.domainService = domainService.NewDomainService(d.repository)
}
