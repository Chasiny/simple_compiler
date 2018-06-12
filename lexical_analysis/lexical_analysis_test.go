package lexical_analysis

import (
	"testing"
)

func TestExplain(t *testing.T) {
	l:=NewLexicalAnalysisDemo()
	words := l.Explain(`int i=0;
	int sum=0;
	float average=0.0;
	while (i<rand(100){
	    sum=sum+ rand(999);
	    i=i+1;
	}
	if(i>0) {
	    average=sum*1.0/i;
	}else{
	}`)

	for i := range words {
		t.Log(words[i].Word)
	}
}
