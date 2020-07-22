package metrics

import "testing"

func Test_AddPiece(t *testing.T) {
	TasksInQueue.WithLabelValues("t01000", "1", "Add_Piece", "Worker1").Inc()
	t.Errorf("fucked")
}
