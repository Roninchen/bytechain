package blockchian

type ProofOfWork struct {
	Block *Block
}

// create new ProofOfWork obj
func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{block}
}

func (proofOfWork *ProofOfWork) Run()([]byte,int64) {

	return nil,0
}
