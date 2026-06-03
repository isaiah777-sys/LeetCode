package leetcode

func printInOrder(first, second, third func()) {
	done := make(chan struct{}, 3)
	go func() {
		first()
		done <- struct{}{}
	}()
	go func() {
		<-done
		second()
		done <- struct{}{}
	}()
	go func() {
		<-done
		third()
		done <- struct{}{}
	}()
	<-done
}
