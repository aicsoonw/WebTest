package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	in "github.com/richkule/prepareTestWeb/init"
)

// Обработка страницы редактирования различных элементов /edit/elem/id
func Edit(w http.ResponseWriter, req *http.Request, sessUs *in.SessUs) error {
	if sessUs.UsId == in.GuestUserId {
		http.Redirect(w, req, `/`, http.StatusFound)
		return nil
	}

	var path string
	// Функция получающая id элемента из путя с помощью регулярного выражения
	// В случае ошибки вставит название элемента nameId в ошибку
	idFunc := func(regexp *regexp.Regexp, nameId string) (int, error) {
		elemGroup := regexp.FindStringSubmatch(path)

		// Регулярные выражения построенны так, что в первой группе всегда будет необходимый id
		id, err := strconv.Atoi(elemGroup[1])
		if err != nil {
			err := fmt.Errorf("Неправильный id %s edit %s ", nameId, err.Error())
			return 0, err
		}
		return id, nil
	}
	path = req.URL.Path
	switch {
	case in.RegTestEdit.MatchString(path):
		intTestId, err := idFunc(in.RegTestEdit, "теста")
		testId := in.TestId(intTestId)
		if err != nil {
			return err
		}
		return editTest(w, req, sessUs, testId)
	case in.RegTopicEdit.MatchString(path):
		intTopicId, err := idFunc(in.RegTopicEdit, "темы")
		topicId := in.TopicId(intTopicId)
		if err != nil {
			return err
		}
		return editTopic(w, req, sessUs, topicId)
	case in.RegQuesEdit.MatchString(path):
		quesId, err := idFunc(in.RegQuesEdit, "вопроса")
		if err != nil {
			return nil
		}
		return editQuestion(w, req, sessUs, quesId)
	}
	err := errors.New("Неправильный путь для редактирования edit ")
	return err
}
