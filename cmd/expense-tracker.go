package cmd   // указываем название нашего пакета
import "time" // импортируем стандартный пакет времени Go

func addExpense(zone string) (string, error) {
	loc, err := time.LoadLocation(zone) // узнаем текущую локацию

	// проверяем на ошибку
	if err != nil {
		return "", err // возвращаем пустой результат с данными об ошибке
	}

	timeNow := time.Now().In(loc)            // получаем текущее время на основе локации
	return timeNow.Format(time.RFC1123), nil // возвращаем отформатированный результат без данных об ошибке
}
