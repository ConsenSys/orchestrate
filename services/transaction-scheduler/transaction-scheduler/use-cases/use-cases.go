package usecases

type UseCases interface {
	TransactionUseCases
	ScheduleUseCases
	JobUseCases
}
