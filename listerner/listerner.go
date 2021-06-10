
package listener

type JackListener struct {
	*parser.BaseJackListener
}


func NewJackListener() *JackListener {
	return new(JackListener)
}
