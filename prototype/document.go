package prototype

type Document interface {
	Print()
	Clone() Document
}
