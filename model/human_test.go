package model

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mock "github.com/7kaji/interface_sample/interface/mock"

	"github.com/golang/mock/gomock"
)

// TODO: 引数に interface の型をしていしてるので gomock 使ってみた
func TestHuman(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAnimal := mock.NewMockAnimal(mockCtrl)
	mockAnimal.EXPECT().Cry().Return("Bow-wow")
	mockAnimal.EXPECT().GetName().Return("Pochi")

	// TODO: table driven test
	taro := NewHuman("Taro")
	actual := taro.Touch(mockAnimal)
	expected := "Taro touched the Pochi then Pochi cried Bow-wow."
	assert.Equal(t, expected, actual)
}
