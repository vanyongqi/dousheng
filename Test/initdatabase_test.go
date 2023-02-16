package Test

import (
	"dousheng-backend/Databases"
	"testing"
)

func TestDatabaseSessions(t *testing.T) {
	t.Run("", func(t *testing.T) {
		Databases.InitDatabase()
	})

}

/*
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()//t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	   //
	   //在这里，我们将在我们的测试库中引入另一个工具 —— 子测试。有时，对一个 "事情" 进行分组测试是很有用的，然后进行描述不同场景的子测试。
	   //这种方法的好处是，你可以设置在其他测试中也能够使用的共享代码。
	   //当我们检查信息是否符合预期时，会有重复的代码。

	t.Run("to a person", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "Bonjour, Lauren"
		assertCorrectMessage(t, got, want)
	})
}
*/
