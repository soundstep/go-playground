package viewmodels

func GetHome() ViewModel {
	result := ViewModel{
		Title:    "Stories",
		SignedIn: false,
		Active:   "home",
	}
	return result
}
