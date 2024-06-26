package utils

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
  "time"
)

func EnsureDir(dir string) error {
  if _, err := os.Stat(dir); os.IsNotExist(err) {
    err := os.MkdirAll(dir, os.ModePerm)
    if err != nil {
      return fmt.Errorf("failed to create directory: %w", err)
    }
  }
  return nil
}

func SaveTransactionToFile(transaction interface{}, dir string) error {

  err := EnsureDir(dir)
  if err != nil {
    return err
  }

  data, err := json.MarshalIndent(transaction, "", "  ")
  if err != nil {
    return fmt.Errorf("failed to marshal transaction: %w", err)
  }

  filename := fmt.Sprintf("transaction_%s.json", time.Now().Format("2006-01-02T15-04-05"))

  filepath := filepath.Join(dir, filename)

  file, err := os.Create(filepath)
  if err != nil {
    return fmt.Errorf("failed to create file: %w", err)
  }
  defer file.Close()

  _, err = file.Write(data)
  if err != nil {
    return fmt.Errorf("failed to write data to file: %w", err)
  }

  return nil
}
