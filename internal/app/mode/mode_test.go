package mode_test

import (
	"hangman/internal/app/mode"
	"testing"

	"hangman/internal/domain/session"

	"hangman/internal/app/mode/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInteractive_Run_WinScenario(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUI := mocks.NewMockUserInterface(ctrl)

	sess, err := session.NewSession("go", 2, session.WithHint("lang"))
	assert.NoError(t, err)

	gomock.InOrder(
		mockUI.EXPECT().Clear().Times(1),
		mockUI.EXPECT().AskRune().Return('g').Times(1),
		mockUI.EXPECT().PrintHangman(2).Times(1),
		mockUI.EXPECT().PrintGuessResult(session.Correct, 2).Times(1),

		mockUI.EXPECT().Clear().Times(1),
		mockUI.EXPECT().AskRune().Return('o').Times(1),
		mockUI.EXPECT().PrintHangman(2).Times(1),
		mockUI.EXPECT().PrintGuessResult(session.Correct, 2).Times(1),

		mockUI.EXPECT().PrintGameStatus(session.Win).Times(1),
	)

	interactiveMode := mode.NewInteractive(sess, mockUI)
	err = interactiveMode.Run()

	assert.NoError(t, err)
	assert.Equal(t, session.Win, sess.GetStatus())
}

func TestInteractive_Run_LoseWithHintScenario(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUI := mocks.NewMockUserInterface(ctrl)

	sess, err := session.NewSession("go", 4, session.WithHint("lang"))
	assert.NoError(t, err)

	gomock.InOrder(
		mockUI.EXPECT().Clear().Times(1),
		mockUI.EXPECT().AskRune().Return('z').Times(1),
		mockUI.EXPECT().PrintHangman(3).Times(1),
		mockUI.EXPECT().PrintGuessResult(session.Incorrect, 3).Times(1),
		mockUI.EXPECT().PrintHint("lang").Times(1), // Проверка вашей логики хинта

		mockUI.EXPECT().Clear().Times(1), mockUI.EXPECT().AskRune().Return('x').Times(1),
		mockUI.EXPECT().PrintHangman(2).Times(1), mockUI.EXPECT().PrintGuessResult(session.Incorrect, 2).Times(1),

		mockUI.EXPECT().Clear().Times(1), mockUI.EXPECT().AskRune().Return('y').Times(1),
		mockUI.EXPECT().PrintHangman(1).Times(1), mockUI.EXPECT().PrintGuessResult(session.Incorrect, 1).Times(1),

		mockUI.EXPECT().Clear().Times(1), mockUI.EXPECT().AskRune().Return('w').Times(1),
		mockUI.EXPECT().PrintHangman(0).Times(1), mockUI.EXPECT().PrintGuessResult(session.Incorrect, 0).Times(1),

		mockUI.EXPECT().PrintGameStatus(session.Lose).Times(1),
	)

	interactiveMode := mode.NewInteractive(sess, mockUI)
	err = interactiveMode.Run()

	assert.NoError(t, err)
	assert.Equal(t, session.Lose, sess.GetStatus())
}

func TestNonInteractive_Run(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUI := mocks.NewMockUserInterface(ctrl)

	sess, err := session.NewSession("go", 2)
	assert.NoError(t, err)

	mockUI.EXPECT().
		Println(gomock.Any()).
		Times(1)

	nonInteractiveMode := mode.NewNonInteractive(sess, "go", mockUI)
	err = nonInteractiveMode.Run()

	assert.NoError(t, err)
	assert.Equal(t, session.Win, sess.GetStatus())
}
