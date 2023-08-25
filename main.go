package main

import (
  "log"
  "github.com/abolfazlalz/asalidrive/storage"
)

func main() {
  storage := storage.NewLocal()
  // Write "Hello World" as array byte to `hello.txt` file
  if err := storage.Write("hello.txt", []byte("Hello World")); err != nil {
    log.Fatal(err)
  }
  // Make copy `hello.txt` to `world.txt`
  if err := storage.Copy("hello.txt", "world.txt"); err != nil {
    log.Fatal(err)
  }
  // Read `world.txt` as byte and log it as string
  if data, err := storage.Read("world.txt"); err != nil {
    log.Fatal(err)
  } else {
    log.Println(string(data))
  }
  // Delete `world.txt` file
  if err := storage.Delete("world.txt"); err != nil {
    log.Fatal(err)
  }
}
