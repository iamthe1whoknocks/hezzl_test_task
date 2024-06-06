package usecase_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

var (
	errInternalServer = errors.New("internal server error")
	errRepo           = errors.New("save item repo error")
	errBrokerPublish  = errors.New("broker publish error")
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func item(t *testing.T) (usecase.ItemUseCase, *MockItemsRepo, *MockCacher, *MockBroker) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockItemsRepo(mockCtl)
	cache := NewMockCacher(mockCtl)
	broker := NewMockBroker(mockCtl)

	itemUseCase := usecase.New(repo, cache, broker)

	return *itemUseCase, repo, cache, broker
}

// test get method.
func TestGet(t *testing.T) {
	t.Parallel()

	item, repo, _, _ := item(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetItems(context.Background()).Return(nil, nil)
			},
			res: []models.Item(nil),
			err: nil,
		},
		{
			name: "internal server error",
			mock: func() {
				repo.EXPECT().GetItems(context.Background()).Return(nil, errInternalServer)
			},
			res: []models.Item(nil),
			err: errInternalServer,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := item.Get(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

// test get method.
func TestSave(t *testing.T) {
	t.Parallel()

	item, repo, _, broker := item(t)

	data, _ := json.Marshal(models.Item{})

	nilItem := &models.Item{}
	nilItem = nil

	tests := []test{
		{
			name: "ok result",
			mock: func() {
				repo.EXPECT().SaveItem(context.Background(), &models.Item{}).Return(&models.Item{}, nil)
				broker.EXPECT().GetSubject().Return("test")
				broker.EXPECT().Publish("test", data).Return(nil)
			},
			res: &models.Item{},
			err: nil,
		},
		{
			name: "repo error",
			mock: func() {
				repo.EXPECT().SaveItem(context.Background(), &models.Item{}).Return(nilItem, errRepo)
			},
			res: nilItem,
			err: errRepo,
		},
		{
			name: "broker error",
			mock: func() {
				repo.EXPECT().SaveItem(context.Background(), &models.Item{}).Return(&models.Item{}, nil)
				broker.EXPECT().GetSubject().Return("test")
				broker.EXPECT().Publish("test", data).Return(errBrokerPublish)
			},
			res: nilItem,
			err: errBrokerPublish,
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := item.Save(context.Background(), &models.Item{})

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

// test delete method.
func TestDelete(t *testing.T) {
	t.Parallel()

	item, repo, _, broker := item(t)

	testItem := models.Item{
		Removed: true,
	}

	data, _ := json.Marshal(testItem)

	testID, testCompaignID := 0, 0

	// nilItem := &models.Item{}
	// nilItem = nil

	tests := []test{
		{
			name: "ok result",
			mock: func() {
				repo.EXPECT().DeleteItem(context.Background(), testID, testCompaignID).Return(true, nil)
				broker.EXPECT().GetSubject().Return("test")
				broker.EXPECT().Publish("test", data).Return(nil)
			},
			res: true,
			err: nil,
		},
		{
			name: "no rows err",
			mock: func() {
				repo.EXPECT().DeleteItem(context.Background(), testID, testCompaignID).Return(true, nil)
				broker.EXPECT().GetSubject().Return("test")
				broker.EXPECT().Publish("test", data).Return(nil)
			},
			res: true,
			err: nil,
		},
		// {
		// 	name: "repo error",
		// 	mock: func() {
		// 		repo.EXPECT().SaveItem(context.Background(), &models.Item{}).Return(nilItem, errRepo)
		// 	},
		// 	res: nilItem,
		// 	err: errRepo,
		// },
		// {
		// 	name: "broker error",
		// 	mock: func() {
		// 		repo.EXPECT().SaveItem(context.Background(), &models.Item{}).Return(&models.Item{}, nil)
		// 		broker.EXPECT().GetSubject().Return("test")
		// 		broker.EXPECT().Publish("test", data).Return(errBrokerPublish)
		// 	},
		// 	res: nilItem,
		// 	err: errBrokerPublish,
		// },
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := item.Delete(context.Background(), testID, testCompaignID)

			require.EqualValues(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
