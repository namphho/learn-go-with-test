package _map


const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it already doesn't exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}


type Dictionary map[string]string

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

func (d Dictionary) SearchDictionary(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.SearchDictionary(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error{
	_, err := d.SearchDictionary(key)
	switch err {
	case nil:
		d[key]=value
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}


func (d Dictionary) Delete(word string) {
	delete(d, word)
}