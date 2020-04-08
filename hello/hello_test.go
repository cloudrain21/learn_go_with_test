package hello

import (
    "testing"
)

func TestHello(t *testing.T) {
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Edie", "Spanish")
        want := "Hola, Edie"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in English", func(t *testing.T) {
        got := Hello("Edie", "")
        want := "Hello, Edie"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Japan", func(t *testing.T) {
        got := Hello("Nakasone", "Japan")
        want := "Gonnichiwa, Nakasone"
        assertCorrectMessage(t, got, want)
    })
}

// assertion 을 function 으로 분리하여 처리 - 중복 제
//func TestHello(t *testing.T) {
//    assertCorrectMessage := func(t *testing.T, got, want string) {
//        // 이 함수는 helper 라는걸 알려줌. -> fail 시 fail line 을 function call 하는 부분 기준으로 찍
//        // t.Helper()
//        if got != want {
//            t.Errorf("got %q want %q", got, want)
//        }
//    }
//
//    t.Run("saying hello to people", func(t *testing.T) {
//        got := Hello("Chris")
//        want := "Hello, Chris"
//        assertCorrectMessage(t, got, want)
//    })
//
//    t.Run("empty string defaults to 'World'", func(t *testing.T) {
//        got := Hello("")
//        want := "Hello, World"
//        assertCorrectMessage(t, got, want)
//    })
//
//}

//func TestHello(t *testing.T) {
//    got := Hello("dplee")
//    want := "Hello, dplee"
//
//    if got != want {
//        t.Errorf("got %q want %q", got, want)
//    }
//}

// 아래와 같이 TestHello 내에 sub test 들을 여러개 둘 수 있다.
// t.Run 에 넘겨주는 인자들은 fail 시 함께 보여준다.
//func TestHello(t *testing.T) {
//
//    t.Run("saying hello to people", func(t *testing.T) {
//        got := Hello("Chris")
//        want := "Hello, Chris"
//
//        if got != want {
//            t.Errorf("got %q want %q", got, want)
//        }
//    })
//
//    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
//        got := Hello("")
//        want := "Hello, World"
//
//        if got != want {
//            t.Errorf("got %q want %q", got, want)
//        }
//    })
//}