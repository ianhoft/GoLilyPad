package minecraft

type IdMapPacket struct {
	id int
}

func (this *IdMapPacket) Id() int {
	return this.id
}

func (this *IdMapPacket) IdFrom(idMap *IdMap) {
	panic("")
}
